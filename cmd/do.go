/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
	"strconv"
	// "encoding/binary"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your todo list as complete, removing it from list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("todo_list.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		for _, arg := range(args){
			temp_int, _ := strconv.Atoi(arg)

			db.Update(func(tx *bolt.Tx) error{
				b := tx.Bucket([]byte("MyTasks"))
				if(b == nil){
					fmt.Println("No tasks currently on list.")
					return nil
				}
				c := b.Cursor()
				key, val := c.First()

				for i := 1; i < temp_int; i++ {
					key, val = c.Next()
				}

				// finished := b.Get([]byte(Itob(temp_int)))
				fmt.Printf("Completed task \"%s\"\n", val)
				return b.Delete(key)
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
