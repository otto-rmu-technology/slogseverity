package slogseverity

import (
	"log/slog"
)

const (
	TraceLevel = "TRACE"
	DebugLevel = "DEBUG"
	InfoLevel  = "INFO"
	WarnLevel  = "WARNING"
	ErrorLevel = "ERROR"
)

func ReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	// Customize the name of the level key and the output string..
	if a.Key == slog.LevelKey {
		// Rename the level key from "level" to "severity".
		a.Key = "severity"

		// Handle custom level values.
		sev := a.Value.Any().(slog.Level)

		// Rename WARN string for WarningLevel to WARNING
		switch {
		case sev < slog.LevelDebug:
			a.Value = slog.StringValue(TraceLevel)
		case sev < slog.LevelInfo:
			a.Value = slog.StringValue(DebugLevel)
		case sev < slog.LevelWarn:
			a.Value = slog.StringValue(InfoLevel)
		case sev < slog.LevelError:
			a.Value = slog.StringValue(WarnLevel)
		default:
			a.Value = slog.StringValue(ErrorLevel)
		}

		return a
	}
	return a
}
