/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
	// "encoding/binary"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incompleted tasks",
	Long: `Retrieves a list of all the incompleted tasks that are
	currently on your task list. `,

	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("todo_list.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		fmt.Println("You have the following tasks: ")
		db.View(func(tx *bolt.Tx) error{
			b := tx.Bucket([]byte("MyTasks"))
			if(b == nil) {
				return nil
			}

			c := b.Cursor()
			temp_at := 1

			for k, v := c.First(); k != nil; k,v = c.Next() {
				fmt.Printf("%d : %s \n", temp_at, string(v))
				temp_at += 1 
			}

			return nil
		})

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
