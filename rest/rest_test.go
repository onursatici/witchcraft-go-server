package rest_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/palantir/witchcraft-go-error"
	"github.com/palantir/witchcraft-go-server/rest"
	"github.com/stretchr/testify/require"
)

func TestWriteJSONResponse_Error(t *testing.T) {
	for _, test := range []struct {
		Name         string
		Err          error
		ExpectedCode int
		ExpectedJSON string
	}{
		{
			Name:         "standard werror",
			Err:          werror.Error("bad things happened"),
			ExpectedCode: 500,
			ExpectedJSON: "bad things happened",
		},
		{
			Name:         "rest.Error with code",
			Err:          rest.NewError(werror.Error("bad things happened"), rest.StatusCode(400)),
			ExpectedCode: 400,
			ExpectedJSON: "bad things happened",
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			server := httptest.NewServer(rest.NewJSONHandler(func(http.ResponseWriter, *http.Request) error {
				return test.Err
			}, rest.StatusCodeMapper, rest.ErrHandler))
			defer server.Close()

			resp, err := server.Client().Get(server.URL)
			require.NoError(t, err)
			require.Equal(t, resp.StatusCode, test.ExpectedCode)
			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, test.ExpectedJSON, string(body))
		})
	}
}
