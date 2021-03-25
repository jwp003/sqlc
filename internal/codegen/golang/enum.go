package golang

import (
	"regexp"
	"strings"

	"golang.org/x/text/unicode/runenames"
)

var IdentPattern = regexp.MustCompile("[^a-zA-Z0-9_]+")
var SeparatorPattern = regexp.MustCompile("([^a-zA-Z0-9]+)[-:/]+([^a-zA-Z0-9]+)")

type Constant struct {
	Name  string
	Type  string
	Value string
}

type Enum struct {
	Name      string
	Comment   string
	Constants []Constant
}

func EnumReplace(value string) string {

	id := SeparatorPattern.ReplaceAllString(value, "$1_$2")

	id = IdentPattern.ReplaceAllStringFunc(id, func(s string) string {
		var replacement string
		for _, v := range s {
			replacement += strings.Title(strings.ToLower(runenames.Name(v)))
		}
		return replacement
	})

	return IdentPattern.ReplaceAllString(id, "")
}

func EnumValueName(value string) string {
	name := ""
	id := strings.Replace(value, "-", "_", -1)
	id = strings.Replace(id, ":", "_", -1)
	id = strings.Replace(id, "/", "_", -1)
	id = IdentPattern.ReplaceAllString(id, "")
	for _, part := range strings.Split(id, "_") {
		name += strings.Title(part)
	}
	return name
}
