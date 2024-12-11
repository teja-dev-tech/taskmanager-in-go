package main

import (
	"fmt"
	"log"
	"os"
	"taskmaster/tasks"
	"taskmaster/api"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: taskmaster [add|list|delete|quote]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		err := tasks.AddTask(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task added successfully!")
	case "list":
		err := tasks.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task ID to delete.")
			return
		}
		err := tasks.DeleteTask(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task deleted successfully!")
	case "quote":
		quote, err := api.GetMotivationalQuote()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Motivational Quote:", quote)
	default:
		fmt.Println("Unknown command. Try [add|list|delete|quote].")
	}
}
