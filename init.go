package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/docopt/docopt-go"
)

// RunCommandInit runs the command "init"
func RunCommandInit(args docopt.Opts) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get working directory: %w", err)
	}

	remote := strings.TrimSuffix(getCurrentRepoRemote(), ".git")

	// Compute repository name: special case for github
	var name string
	if strings.Contains(remote, "https://github.com") {
		name = remote[strings.LastIndex(remote, "/")+1:]
	} else {
		name = filepath.Dir(workingDir)
	}

	// For the manifest's content
	if remote == "" {
		remote = "# TODO"
	}

	// TODO: set user{Chrome,Content,.js} by finding their path
	// TODO: only set assets if chrome/ actually exists
	content := fmt.Sprintf(heredoc.Doc(`
		ffcss: %d

		name: %s
		download: %s

		userChrome: userChrome.css
		userContent: userContent.css
		user.js: user.js
		assets:
			- chrome/**
	`), VersionMajor, name, remote)

	err = ioutil.WriteFile(filepath.Join(workingDir, "ffcss.yaml"), []byte(content), 0700)
	if err != nil {
		return fmt.Errorf("while writing the manifest: %w", err)
	}

	return nil
}
