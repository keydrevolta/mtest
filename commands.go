package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/openebs/mtest/cmd"
)

// Commands returns the mapping of CLI commands for mtest. The meta
// parameter lets you set meta options for all commands.
func Commands(metaPtr *cmd.Meta) map[string]cli.CommandFactory {
	if metaPtr == nil {
		metaPtr = new(cmd.Meta)
	}

	meta := *metaPtr
	if meta.Ui == nil {
		meta.Ui = &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		}
	}

	return map[string]cli.CommandFactory{
		"ebs": func() (cli.Command, error) {
			return &cmd.EBSCommand{
				Ui: meta.Ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			ver := Version
			rel := VersionPrerelease
			if GitDescribe != "" {
				ver = GitDescribe
				// Trim off a leading 'v', we append it anyways.
				if ver[0] == 'v' {
					ver = ver[1:]
				}
			}
			if GitDescribe == "" && rel == "" && VersionPrerelease != "" {
				rel = "dev"
			}

			return &cmd.VersionCommand{
				Revision:          GitCommit,
				Version:           ver,
				VersionPrerelease: rel,
				Ui:                meta.Ui,
			}, nil
		},
	}
}
