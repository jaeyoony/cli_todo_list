/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"encoding/binary"
	// "strconv"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task to your todo list ",
	Long: `Use this command to add a task to your todo list. 
	Usage is: add <task>, where <task> is the entry to be added
	to the todo list. It really isnt too complicated.`,

	Run: func(cmd *cobra.Command, args []string) {
		// connect to db, create bucket iff doesn't exist
		db, err := bolt.Open("todo_list.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte("MyTasks"))
			if err != nil {
				return fmt.Errorf("create bucket %s", err)
			}
			return nil
		})

		// add tasks to bucket 
		for _, arg := range args{
			
			db.Update(func(tx *bolt.Tx) error {
				bucket := tx.Bucket([]byte("MyTasks"))
				id, _ := bucket.NextSequence()
				return bucket.Put(Itob(int(id)), []byte(arg))

			})
			fmt.Println("Added \"" + arg + "\" to your task list.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func Itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}