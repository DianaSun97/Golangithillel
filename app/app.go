package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DailySchedule struct {
	StartTime float64
	EndTime   float64
}

type WeeklySchedule struct {
	Monday    DailySchedule
	Tuesday   DailySchedule
	Wednesday DailySchedule
	Thursday  DailySchedule
	Friday    DailySchedule
	Saturday  DailySchedule
	Sunday    DailySchedule
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Schedule  WeeklySchedule
}

type EmployeeManager struct {
	employees map[int]Employee
}

func NewEmployeeManager() *EmployeeManager {
	return &EmployeeManager{
		employees: make(map[int]Employee),
	}
}

func (em *EmployeeManager) AddEmployee(firstName, lastName string) {
	employeeID := len(em.employees) + 1
	employee := Employee{
		ID:        employeeID,
		FirstName: firstName,
		LastName:  lastName,
	}
	em.employees[employeeID] = employee
}

func InputNewEmployee(em *EmployeeManager) {
	var firstName, lastName string

	fmt.Print("Enter employee first name: ")
	reader := bufio.NewReader(os.Stdin)
	firstName, _ = reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Enter employee last name: ")
	lastName, _ = reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	if firstName == "" || lastName == "" {
		fmt.Println("\nFirst name or last name cannot be empty.")
		return
	}

	em.AddEmployee(firstName, lastName)
}

func (em *EmployeeManager) AssignWorkHours(employeeID, day int, startTime, endTime float64) {
	employee, exists := em.employees[employeeID]
	if !exists {
		fmt.Println("Employee not found")
		return
	}
	switch day {
	case 1:
		employee.Schedule.Monday = DailySchedule{StartTime: startTime, EndTime: endTime}
	case 2:
		employee.Schedule.Tuesday = DailySchedule{StartTime: startTime, EndTime: endTime}
	case 3:
		employee.Schedule.Wednesday = DailySchedule{StartTime: startTime, EndTime: endTime}
	case 4:
		employee.Schedule.Thursday = DailySchedule{StartTime: startTime, EndTime: endTime}
	case 5:
		employee.Schedule.Friday = DailySchedule{StartTime: startTime, EndTime: endTime}
	case 6:
		employee.Schedule.Saturday = DailySchedule{StartTime: startTime, EndTime: endTime}
	case 7:
		employee.Schedule.Sunday = DailySchedule{StartTime: startTime, EndTime: endTime}
	default:
		fmt.Println("Invalid day")
		return
	}
	em.employees[employeeID] = employee
}

func SetEmployeeWorkHours(em *EmployeeManager) {
	var employeeID, day int
	var startTime, endTime float64

	fmt.Print("Enter the employee ID:")
	fmt.Scanln(&employeeID)

	if _, ok := em.employees[employeeID]; !ok {
		fmt.Printf("Employee with ID %d not found.\n", employeeID)
		return
	}

	fmt.Print("Enter the day of the week (1 for Monday, etc.):")
	fmt.Scanln(&day)

	if day < 1 || day > 7 {
		fmt.Println("Invalid day")
		return
	}

	fmt.Print("Enter the time the employee starts work:")
	fmt.Scanln(&startTime)
	fmt.Print("Enter the time the employee ends work:")
	fmt.Scanln(&endTime)

	if startTime < 0 || endTime < 0 {
		fmt.Println("Negative time is not allowed")
		return
	}

	em.AssignWorkHours(employeeID, day, startTime, endTime)
}

func DisplayAllEmployees(em *EmployeeManager) {
	if len(em.employees) == 0 {
		fmt.Println("\nNo employees found")
		return
	}
	fmt.Println("\nEmployee list:")
	for _, employee := range em.employees {
		fmt.Printf("ID: %d, First Name: %s, Last Name: %s\n", employee.ID, employee.FirstName, employee.LastName)
	}
}

func ShowEmployeeWeeklyHours(em *EmployeeManager) {
	var employeeID int
	var totalHours float64

	fmt.Print("Enter the employee ID:")
	fmt.Scanln(&employeeID)

	employee, exists := em.employees[employeeID]
	if !exists {
		fmt.Println("Employee with ID", employeeID, "not found")
		return
	}

	fmt.Printf("\nWeekly schedule for %s %s\n", employee.FirstName, employee.LastName)

	totalHours = 0
	workDays := []DailySchedule{
		employee.Schedule.Monday,
		employee.Schedule.Tuesday,
		employee.Schedule.Wednesday,
		employee.Schedule.Thursday,
		employee.Schedule.Friday,
		employee.Schedule.Saturday,
		employee.Schedule.Sunday,
	}

	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i, day := range workDays {
		fmt.Printf("%s: Starts at %.2f, Ends at %.2f\n", daysOfWeek[i], day.StartTime, day.EndTime)
		if day.StartTime > day.EndTime {
			totalHours += 24 - day.StartTime + day.EndTime
		} else {
			totalHours += day.EndTime - day.StartTime
		}
	}

	fmt.Printf("Total work hours for the week: %.2f\n", totalHours)
	fmt.Println()
}
