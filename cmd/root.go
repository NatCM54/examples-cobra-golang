package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "cmd - a simple println ",
	Long: `stringer is a super fancy CLI (kidding)

	One can use stringer to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hellow Word")

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
