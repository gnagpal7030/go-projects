package main

import (
	"fmt"
	"slices"

	"github.com/gnagpal7030/go-crud-cli/database"
	"github.com/gnagpal7030/go-crud-cli/model"
)

func createTODO() {
	var todoName string
	fmt.Println("Enter todo name")
	fmt.Scan(&todoName)

	var todoStatus string
	fmt.Println("Enter todoStatus")
	fmt.Scan(&todoStatus)

	database.TodoDB = append(database.TodoDB, model.Todo{
		TodoID:     len(database.TodoDB) + 1,
		TodoName:   todoName,
		TodoStatus: todoStatus,
	})
}

func showTODO() {
	fmt.Println("Below are the todos: ")
	for _, v := range database.TodoDB {
		fmt.Printf("\n%+v", v)
	}
}

func updateTODO() {
	var todoID int
	fmt.Println("Enter the todoID to be updated")
	fmt.Scan(&todoID)

	var todoStatus string
	fmt.Println("Enter todo new status")
	fmt.Scan(&todoStatus)

	// search the todoID first, update only when it is present
	for i, v := range database.TodoDB {
		if v.TodoID == todoID {
			database.TodoDB[i].TodoStatus = todoStatus
			fmt.Println("todo updated successfully")
			return
		}
	}
	fmt.Println("todoID not found")
}

func deleteTODO() {
	var todoID int
	fmt.Println("Enter the todoID to be deleted")
	fmt.Scan(&todoID)

	// search the todoID first, delete only when it is present
	for i, v := range database.TodoDB {
		if v.TodoID == todoID {
			database.TodoDB = slices.Delete(database.TodoDB, i, i+1)
			fmt.Println("todo deleted successfully.")
			return
		}
	}
	fmt.Println("todoID not found")
}

func main() {
	var choice int

	for true {
		fmt.Println("1. Create TODO")
		fmt.Println("2. Show TODOs")
		fmt.Println("3. Update TODO")
		fmt.Println("4. Delete TODO")
		fmt.Println("Enter any other key to exit.")
		fmt.Println("Enter your choice")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createTODO()
			fmt.Println()
		case 2:
			showTODO()
			fmt.Println()
		case 3:
			updateTODO()
			fmt.Println()
		case 4:
			deleteTODO()
			fmt.Println()
		default:
			fmt.Println("You exited the CLI tool.")
			return
		}
	}
}
