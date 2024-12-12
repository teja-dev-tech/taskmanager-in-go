package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Task struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var tasksFile = "data/tasks.json"

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(tasksFile); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(tasksFile)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func AddTask(text string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	newTask := Task{ID: fmt.Sprintf("%d", len(tasks)+1), Text: text}
	tasks = append(tasks, newTask)
	return saveTasks(tasks)
}

func ListTasks() error {
	tasks, err := loadTasks()
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
	tasks, err := loadTasks()
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
	return saveTasks(updatedTasks)
}
