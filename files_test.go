package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveFilenames(t *testing.T) {
	var result []string

	result, err := ResolveFilenames([]FileTemplate{"chrome/**"}, UserChoices{})
	assert.Nil(t, err)
	assert.Equal(t,
		[]string{"chrome/**"},
		result,
	)

	result, err = ResolveFilenames([]FileTemplate{
		"{{ os }}/chrome-{{variant}}.css",
		"userChrome.css",
	}, UserChoices{
		Variant: Variant{Name: "rainbow"},
		OS:      "linux",
	})
	assert.Nil(t, err)
	assert.Equal(t,
		[]string{
			"linux/chrome-rainbow.css",
		},
		result,
	)
}