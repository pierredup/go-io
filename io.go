package io

import (
	"github.com/fatih/color"
)

const (
	danger   = color.BgRed
	question = color.FgYellow
	info     = color.FgCyan
	comment  = color.FgMagenta
	warning  = color.FgYellow
	success  = color.FgGreen
)

func Question(q ...string) string { return printColor(question, q...) }

func Info(q ...string) string { return printColor(info, q...) }

func Success(q ...string) string { return printColor(success, q...) }

func Comment(q ...string) string { return printColor(comment, q...) }

func Warning(q ...string) string { return printColor(warning, q...) }

func Error(q ...string) string { return printColor(danger, q...) }

func printColor(a color.Attribute, strings ...string) string {
	c := color.New(a)
	r := ""

	for _, str := range strings {
		r += c.Sprint(str)
	}

	return r
}
