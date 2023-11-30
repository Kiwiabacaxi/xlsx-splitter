package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Kiwiabacaxi/xlsx-splitter/internal/log"
	"github.com/Kiwiabacaxi/xlsx-splitter/internal/myapp"
	"go.uber.org/zap"

	// "playground/pkg/myapp"

	_ "go.uber.org/automaxprocs"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := log.NewAtLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return err
	}

	defer func() {
		err = logger.Sync()
	}()

	wd, err := os.Getwd()
	if err != nil {
		logger.Error("get working directory failed", zap.Error(err))
		return err
	}

	logger.Info("Current working directory", zap.String("directory", wd))

	result, err := myapp.ProcessExcelFile(wd + "/testdata/Culturas.xlsx")
	if err != nil {
		logger.Error("process excel file failed", zap.Error(err))
		return err
	}

	err = os.WriteFile(wd+"/output/output.txt", []byte(strings.Join(result, "\n\n")), 0644)
	if err != nil {
		logger.Error("write file failed", zap.Error(err))
		return err
	}

	return err
}
