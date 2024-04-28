package main

import "fmt"

type Task struct {
	Title     string
	Completed bool
}

type TasksList map[int]*Task

// Add a new task to the list
func (tl TasksList) Add(title string) {
	key := len(tl) + 1
	tl[key] = &Task{Title: title}
	fmt.Println("Added task:", title)
}

// Delete a task by ID
func (tl TasksList) Delete(key int) {
	if _, ok := tl[key]; ok {
		delete(tl, key)
		fmt.Printf("Task %d has been deleted\n", key)
	} else {
		fmt.Printf("TaskList does not contain a Task with the specified id: %d\n", key)
	}
}

// Print all tasks in the list
func (tl TasksList) Print() {
	if len(tl) == 0 {
		fmt.Println("TaskList is empty")
		return
	}
	fmt.Println("TaskList:")
	for key, task := range tl {
		fmt.Printf("TaskID: %d Title: %s Completed: %t\n", key, task.Title, task.Completed)
	}
}

// Mark a task as completed by ID
func (tl TasksList) CompleteTask(taskID int) {
	if task, ok := tl[taskID]; ok {
		task.Completed = true
		fmt.Printf("Task %d marked as completed\n", taskID)
	} else {
		fmt.Printf("TaskList does not contain a Task with the specified id: %d\n", taskID)
	}
}

func main() {
	tasks := make(TasksList)

	for {
		fmt.Println("\n1. Add a task")
		fmt.Println("2. Remove a task")
		fmt.Println("3. View all tasks")
		fmt.Println("4. Mark a task as completed")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task title: ")
			var title string
			_, err := fmt.Scanf("%s", &title)
			if err != nil {
				fmt.Println("Error: Failed to read task title.")
				continue
			}
			tasks.Add(title)
		case 2:
			fmt.Print("Enter ID of the task to remove: ")
			var id int
			_, err := fmt.Scanf("%d", &id)
			if err != nil {
				fmt.Println("Error: Failed to read task ID.")
				continue
			}
			tasks.Delete(id)
		case 3:
			tasks.Print()
		case 4:
			fmt.Print("Enter ID of the task to mark as completed: ")
			var id int
			_, err := fmt.Scanf("%d", &id)
			if err != nil {
				fmt.Println("Error: Failed to read task ID.")
				continue
			}
			tasks.CompleteTask(id)
		case 5:
			fmt.Println("Program exited.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
