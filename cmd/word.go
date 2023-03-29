package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Word",
	Run: func(cmd *cobra.Command, args []string) {
		res := "Word"
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(wordCmd)
}
