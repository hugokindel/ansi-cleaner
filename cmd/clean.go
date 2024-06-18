package cmd

import (
	"fmt"
	"github.com/hugokindel/ansi-cleaner/pkg/ansi"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var monitorCmd = &cobra.Command{
	Use:   "clean [filename]",
	Short: "Remove all ANSI escape codes from a file",
	Args:  cobra.ExactArgs(1),
	Run:   runCleanCommand,
}

func init() {
	Cmd.AddCommand(monitorCmd)
}

func runCleanCommand(cmd *cobra.Command, args []string) {
	filename := args[0]

	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ansi.Strip(string(b)))
}
