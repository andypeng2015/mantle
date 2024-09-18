package controller

import (
	"log/slog"
	"os"
)

var gLogger *slog.Logger

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	gLogger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				a.Key = "logged_at"
			case slog.LevelKey:
				a.Key = "severity"
			case slog.MessageKey:
				a.Key = "message"
			}
			return a
		},
	})).With(slog.String("utsname", hostname))
}
