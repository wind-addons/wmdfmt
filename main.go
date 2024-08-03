package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mdfmt",
		Usage: "Format Markdown files",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "in-place",
				Aliases: []string{"i"},
				Usage:   "Format file(s) in-place",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "Recursively format files in subdirectories",
			},
			&cli.StringFlag{
				Name:  "ignore",
				Usage: "Ignore files matching this pattern",
			},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		if showErr := cli.ShowAppHelp(cli.NewContext(app, nil, nil)); showErr != nil {
			os.Exit(1)
		}
		_, _ = fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	inPlace := c.Bool("in-place")
	recursive := c.Bool("recursive")
	ignorePattern := c.String("ignore")

	if c.NArg() == 0 {
		// No arguments, read from stdin
		return processStdin()
	}

	for _, path := range c.Args().Slice() {
		err := processPath(path, inPlace, recursive, ignorePattern)
		if err != nil {
			return err
		}
	}

	return nil
}

func processStdin() error {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("error reading from stdin: %v", err)
	}

	formatted := formatContent(input)
	_, err = os.Stdout.Write(formatted)
	if err != nil {
		return fmt.Errorf("error writing to stdout: %v", err)
	}

	return nil
}

func processPath(path string, inPlace, recursive bool, ignorePattern string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("error accessing path %s: %v", path, err)
	}

	if fileInfo.IsDir() {
		return processDirectory(path, inPlace, recursive, ignorePattern)
	}
	return processFile(path, inPlace, ignorePattern)
}

func processDirectory(dirPath string, inPlace, recursive bool, ignorePattern string) error {
	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if !recursive && path != dirPath {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		return processFile(path, inPlace, ignorePattern)
	}

	return filepath.WalkDir(dirPath, walkFunc)
}

func processFile(filePath string, inPlace bool, ignorePattern string) error {
	if shouldIgnore(filePath, ignorePattern) {
		return nil
	}

	input, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	formatted := formatContent(input)

	if inPlace {
		err = os.WriteFile(filePath, formatted, 0o600)
		if err != nil {
			return fmt.Errorf("error writing to file %s: %v", filePath, err)
		}
	} else {
		fmt.Printf("--- %s ---\n", filePath)
		_, err = os.Stdout.Write(formatted)
		if err != nil {
			return fmt.Errorf("error writing to stdout: %v", err)
		}
		fmt.Println()
	}

	return nil
}

func shouldIgnore(path, ignorePattern string) bool {
	if ignorePattern == "" {
		return false
	}
	return strings.Contains(filepath.Base(path), ignorePattern)
}
