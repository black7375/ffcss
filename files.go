package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/bmatcuk/doublestar"
	"github.com/hoisie/mustache"
)

type UserChoices struct {
	Variant Variant
	OS      string
}

func NewUserChoices() UserChoices {
	return UserChoices{
		Variant: Variant{},
		OS:      GOOStoOS(runtime.GOOS),
	}
}

func GOOStoOS(GOOS string) string {
	switch GOOS {
	case "darwin":
		return "macos"
	case "plan9":
		return "linux"
	default:
		return GOOS
	}
}

// ResolveFilenames resolves the file names using choices made by the user (variant selected, current OS).
// It does not resolve glob patterns though.
func ResolveFilenames(files []FileTemplate, choices UserChoices) (resolved []string, err error) {
	for _, file := range files {
		var output string
		templ, err := mustache.ParseString(file)
		if err != nil {
			return resolved, fmt.Errorf("could not parse %q: %s", file, err.Error())
		}
		output = templ.Render(map[string]string{
			"os": choices.OS,
			"variant": choices.Variant.Name,
		})
		if err != nil {
			return resolved, fmt.Errorf("could not render %q: %s", file, err.Error())
		}
		resolved = append(resolved, output)
	}
	return
}

func CopyOver(config Config, files []string, toDirs []string, theme Manifest) {
	for _, glob := range files {
		for _, file := range doublestar.Glob(GetThemeDownloadPath() glob) {

		}
	}
}