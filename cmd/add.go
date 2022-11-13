/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

}

func addRun(cmd *cobra.Command, args []string) {

	for _, x := range args {
		fmt.Println(x)
	}
}
