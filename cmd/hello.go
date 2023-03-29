package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello",
	Run: func(cmd *cobra.Command, args []string) {
		res := "Hello"
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
