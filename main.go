package main

import (
	"fmt"
	"time"
)

func Welcome() string {
	fmt.Println("1 - Add your task to the list")
	fmt.Println("2 - Check all the tasks")
	fmt.Println("3 - done - type id like that 'done<12>' ")
	fmt.Println("4 - delete - type delete like that 'delete<12>' ")
	fmt.Println("5 - Stop it ")
	return ""
}

func main() {
	fmt.Println("Welcome to TO_DO_LIST and manage everything")
	fmt.Println("To start your journey type 1 out of 4 ")

	Welcome()

	var choice int
	for {
		_, err := fmt.Scan(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:
			var title string
			fmt.Scan(&title)
			AddTask(title)
			time.Sleep(2 * time.Second)
			Welcome()
		case 2:
			ListTasks()
			time.Sleep(2 * time.Second)
			Welcome()
		case 3:
			AddTask("title")
			time.Sleep(2 * time.Second)
			Welcome()
		case 4:
			AddTask("title")
			time.Sleep(2 * time.Second)
			Welcome()
		case 5:
			AddTask("title")
			time.Sleep(2 * time.Second)
			Welcome()
		}

	}
}
