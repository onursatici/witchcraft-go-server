// This file was generated by Conjure and should not be manually edited.

package logging

import (
	"encoding/json"
	"strings"
)

type AuditResult string

const (
	AuditResultSuccess      AuditResult = "SUCCESS"
	AuditResultUnauthorized AuditResult = "UNAUTHORIZED"
	AuditResultError        AuditResult = "ERROR"
	AuditResultUnknown      AuditResult = "UNKNOWN"
)

func (e *AuditResult) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch strings.ToUpper(s) {
	default:
		*e = AuditResultUnknown
	case "SUCCESS":
		*e = AuditResultSuccess
	case "UNAUTHORIZED":
		*e = AuditResultUnauthorized
	case "ERROR":
		*e = AuditResultError
	}
	return nil
}

type LogLevel string

const (
	LogLevelFatal   LogLevel = "FATAL"
	LogLevelError   LogLevel = "ERROR"
	LogLevelWarn    LogLevel = "WARN"
	LogLevelInfo    LogLevel = "INFO"
	LogLevelDebug   LogLevel = "DEBUG"
	LogLevelTrace   LogLevel = "TRACE"
	LogLevelUnknown LogLevel = "UNKNOWN"
)

func (e *LogLevel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch strings.ToUpper(s) {
	default:
		*e = LogLevelUnknown
	case "FATAL":
		*e = LogLevelFatal
	case "ERROR":
		*e = LogLevelError
	case "WARN":
		*e = LogLevelWarn
	case "INFO":
		*e = LogLevelInfo
	case "DEBUG":
		*e = LogLevelDebug
	case "TRACE":
		*e = LogLevelTrace
	}
	return nil
}
