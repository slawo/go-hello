package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"
)

func main() {
	appName := filepath.Base(os.Args[0])

	err := (&cli.Command{
		Name:   appName,
		Action: runGoHello,
		Commands: []*cli.Command{
			getVersionCommand(),
		},
		Authors: []any{"Slawomir Caluch"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "level",
				Value: log.ErrorLevel.String(),
				Usage: fmt.Sprintf("error level, can be any of %s",
					strings.Join([]string{
						log.PanicLevel.String(),
						log.FatalLevel.String(),
						log.ErrorLevel.String(),
						log.WarnLevel.String(),
						log.InfoLevel.String(),
						log.DebugLevel.String(),
						log.TraceLevel.String(),
					}, "|"),
				),
				Aliases: []string{"error-level"},
			},
		},
	}).Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalf("%s ended with error: %s", appName, err.Error())
	}
}

func initErrorLevel(e string) error {
	l, err := log.ParseLevel(strings.ToLower(e))
	if err != nil {
		return err
	}
	log.SetLevel(l)
	return nil
}

func runGoHello(ctx context.Context, cli *cli.Command) error {
	if err := initErrorLevel(cli.String("level")); err != nil {
		return err
	}
	log.Tracef("starting")
	defer log.Tracef("ending")

	return nil
}
