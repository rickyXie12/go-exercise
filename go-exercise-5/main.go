package main

import (
	"errors"
	"fmt"
	"os"
)

const studentRecordFileName = "students.txt"

func readStudentRecord() (string, error) {
	data, err := os.ReadFile(studentRecordFileName)
	if err != nil {
		return "No Data", errors.New("no Data, please fill the data first")
	}

	return string(data), nil
}

func addStudentRecord(name string, score int) string {
	data, err := readStudentRecord()
	if err != nil {
		data = fmt.Sprintf("[name: %s, score: %d]", name, score)
		os.WriteFile(studentRecordFileName, []byte(data), 0664)
		return data
	}

	data = fmt.Sprintf("%s,\n[name: %s, score: %d]", data, name, score)
	os.WriteFile(studentRecordFileName, []byte(data), 0664)
	return data

}

func main() {
	fmt.Println("WELCOME TO STUDENT MANAGEMENT SYSTEM")
	for {
		fmt.Println("Choose your action below with number")
		fmt.Println("1. Add a student record (name & number)")
		fmt.Println("2. View  all student records")
		fmt.Println("3. Exit")

		var choice int

		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var score int
			fmt.Print("Student Name: ")
			fmt.Scan(&name)
			fmt.Print("Student Score: ")
			fmt.Scan(&score)
			if score < 0 {
				err := errors.New("invalid Score. Score should not be negative")
				fmt.Println(err)
				break
			}
			studentData := addStudentRecord(name, score)
			fmt.Println(studentData)
		case 2:
			studentData, err := readStudentRecord()
			if err != nil {
				fmt.Println(studentData)
				fmt.Println(err)
				fmt.Println("_________")
				break
			}
			fmt.Println(studentData)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// const fileName = "students.txt"

// // Function to add a student record
// func addStudent() {
// 	var name string
// 	var score int

// 	fmt.Print("Enter student name: ")
// 	fmt.Scanln(&name)

// 	fmt.Print("Enter student score: ")
// 	fmt.Scanln(&score)

// 	// Open file in append mode
// 	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Write data to file
// 	_, err = file.WriteString(fmt.Sprintf("%s,%d\n", name, score))
// 	if err != nil {
// 		fmt.Println("Error writing to file:", err)
// 		return
// 	}

// 	fmt.Println("Student record added successfully!")
// }

// // Function to read and display student records
// func viewStudents() {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	fmt.Println("\n--- Student Records ---")
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		data := strings.Split(line, ",")
// 		if len(data) == 2 {
// 			score, _ := strconv.Atoi(data[1])
// 			fmt.Printf("Name: %s, Score: %d\n", data[0], score)
// 		}
// 	}
// 	fmt.Println("-----------------------")

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading file:", err)
// 	}
// }

// func main() {
// 	for {
// 		// Display menu
// 		fmt.Println("\nStudent Score Manager")
// 		fmt.Println("1. Add Student")
// 		fmt.Println("2. View Students")
// 		fmt.Println("3. Exit")
// 		fmt.Print("Enter your choice: ")

// 		var choice int
// 		fmt.Scanln(&choice)

// 		// Handle user choice using switch
// 		switch choice {
// 		case 1:
// 			addStudent()
// 		case 2:
// 			viewStudents()
// 		case 3:
// 			fmt.Println("Exiting program...")
// 			return
// 		default:
// 			fmt.Println("Invalid choice! Please enter 1, 2, or 3.")
// 		}
// 	}
// }
