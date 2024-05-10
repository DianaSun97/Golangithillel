package main

import (
	"fmt"
)

func main() {
	manageEmployeeTime()
}

func manageEmployeeTime() {
	employeeManager := app.NewEmployeeManager()

	for {
		fmt.Println("\nTime Management System")
		fmt.Println("\nOptions:")
		fmt.Println("[1] Add an employee")
		fmt.Println("[2] Assign work hours")
		fmt.Println("[3] List all employees")
		fmt.Println("[4] Show weekly work hours")
		fmt.Println("[5] Exit")
		fmt.Print("\nChoose an option: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			app.InputNewEmployee(employeeManager)
		case 2:
			app.SetEmployeeWorkHours(employeeManager)
		case 3:
			app.DisplayAllEmployees(employeeManager)
		case 4:
			app.ShowEmployeeWeeklyHours(employeeManager)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
