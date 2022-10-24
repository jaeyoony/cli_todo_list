/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A CLI todo-list application",
	Long: `A command line interface todo list application utilizing bolt-DB to keep track 
	of a list of tasks to do. Ya know, like a list of to-do's..... a todo list...`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
		// do i need to check if my todo-list bucket exists here, 
		// before the application does anything? 

	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("Does this run first every time..? ")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli_todo_list.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


