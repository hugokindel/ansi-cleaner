package cmd

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ansi-cleaner",
	Short: "ansi-cleaner - remove ANSI from anything",
	Long:  "ansi-cleaner is a tool to remove all ANSI escape codes from your files.",
}
