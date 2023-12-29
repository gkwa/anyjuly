package anyjuly

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	LogFormat string `long:"log-format" choice:"text" choice:"json" default:"text" required:"false"`
	Verbose   []bool `short:"v" long:"verbose" description:"Show verbose debug information, each -v bumps log level"`
	logLevel  slog.Level
}

func Execute() int {
	if err := parseFlags(); err != nil {
		return 1
	}

	if err := setLogLevel(); err != nil {
		return 1
	}

	if err := setupLogger(); err != nil {
		return 1
	}

	if err := run(); err != nil {
		slog.Error("run failed", "error", err)
		return 1
	}

	return 0
}

func parseFlags() error {
	_, err := flags.Parse(&opts)
	return err
}

func run() error {
	tempFile, err := os.CreateTemp("", "clipboard-*.txt")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	defer tempFile.Close()

	slog.Debug("temp file created", "path", tempFile.Name())

	return nil
}
