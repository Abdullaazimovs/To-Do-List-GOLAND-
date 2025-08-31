package main

import "fmt"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task
var nextID int = 1

func InitTasks() {
	loaded, err := LoadTasks()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}
	tasks = loaded

	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
}

func AddTask(title string) {
	task := Task{ID: nextID, Title: title, Done: false}
	tasks = append(tasks, task)
	nextID++

	err := SaveTasks(tasks)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	fmt.Println("Задача добавлена:", title)
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список пуст")
		return
	}
	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}
}
