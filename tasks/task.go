package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"taskmaster/storage"
)

type Task struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var tasksFile = "data/tasks.json"

func AddTask(text string) error {
	tasks, err := storage.LoadTasks(tasksFile)
	if err != nil {
		return err
	}
	newTask := Task{ID: fmt.Sprintf("%d", len(tasks)+1), Text: text}
	tasks = append(tasks, newTask)
	return storage.SaveTasks(tasksFile, tasks)
}

func ListTasks() error {
	tasks, err := storage.LoadTasks(tasksFile)
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("[%s] %s\n", task.ID, task.Text)
	}
	return nil
}

func DeleteTask(id string) error {
	tasks, err := storage.LoadTasks(tasksFile)
	if err != nil {
		return err
	}
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}
	if len(tasks) == len(updatedTasks) {
		return errors.New("task ID not found")
	}
	return storage.SaveTasks(tasksFile, updatedTasks)
}
