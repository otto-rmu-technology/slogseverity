package slogseverity

import (
	"log/slog"
	"reflect"
	"testing"
	"time"
)

func TestReplaceAttr(t *testing.T) {
	testTime := time.Date(2025, 6, 27, 11, 28, 0, 0, time.UTC)
	testSource := slog.Source{
		Function: "test",
		File:     "test.go",
		Line:     1,
	}

	type args struct {
		groups []string
		a      slog.Attr
	}
	tests := []struct {
		name string
		args args
		want slog.Attr
	}{
		{
			name: "LevelKey",
			args: args{
				groups: []string{},
				a:      slog.Any(slog.LevelKey, slog.LevelDebug),
			},
			want: slog.Any("severity", slog.LevelDebug),
		},
		{
			name: "MessageKey",
			args: args{
				groups: []string{},
				a:      slog.String(slog.MessageKey, "test"),
			},
			want: slog.String("msg", "test"),
		},
		{
			name: "OtherKey",
			args: args{
				groups: []string{},
				a:      slog.String("test", "test"),
			},
			want: slog.String("test", "test"),
		},
		{
			name: "TimeKey",
			args: args{
				groups: []string{},
				a:      slog.Time(slog.TimeKey, testTime),
			},
			want: slog.Time("time", testTime),
		},
		{
			name: "SourceKey",
			args: args{
				groups: []string{},
				a:      slog.Any(slog.SourceKey, &testSource),
			},
			want: slog.Any("source", &testSource),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAttr(tt.args.groups, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReplaceAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}
