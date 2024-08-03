package main

import (
	"github.com/88250/lute"
)

var formatter = lute.New()

func formatContent(input []byte) []byte {
	formatted := formatter.Format("", input)

	if len(formatted) > 0 && formatted[len(formatted)-1] != '\n' {
		formatted = append(formatted, '\n')
	}

	return formatted
}
