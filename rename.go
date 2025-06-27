package slogseverity

import (
	"log/slog"
)

func ReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		a.Key = "severity"
		return a
	}
	return a
}
