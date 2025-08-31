package main

import (
	"encoding/json"
	"os"
)

// Save all tasks to todo.json
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("todo.json", data, 0644)
}

// Load all tasks from todo.json
func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile("todo.json")
	if err != nil {
		// если файла ещё нет — вернём пустой список
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
