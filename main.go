package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var id int
var name string
var domain string

func main() {
	var choice int
	db := connectPostgresDB()
	fmt.Println("Welcome to the Student Database Management System")
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Insert Student Record")
		fmt.Println("2. View Student Records")
		fmt.Println("3. Update Student Record")
		fmt.Println("4. Delete Student Record")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			Insert(db)
		case 2:
			Read(db)
		case 3:
			Update(db)
		case 4:
			Delete(db)
		case 5:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// CONNECT DB

// Before connecting, you have to create a database and a table in the psql shell (just a base code; improve these codes as well as you need)

func connectPostgresDB() *sql.DB {
	connstring := "user=postgres dbname=studinfo password=1234 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// INSERT

func Insert(db *sql.DB) {
	var id int
	var name, domain string

	fmt.Println("\nEnter Student Information:")
	fmt.Print("Student ID: ")
	fmt.Scan(&id)

	fmt.Print("Student Name: ")
	fmt.Scan(&name)

	fmt.Print("Student Domain: ")
	fmt.Scan(&domain)

	insertIntoPostgres(db, id, name, domain)
	fmt.Println("Student record inserted successfully!")
}

func insertIntoPostgres(db *sql.DB, id int, name, domain string) {
	_, err := db.Exec("INSERT INTO  students(id,name,domain) VALUES($1,$2,$3)", id, name, domain)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value inserted")
	}
}

// READ

func Read(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM students ORDER BY id")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nStudent Records:")
	fmt.Println("ID  Name       Domain")
	for rows.Next() {
		rows.Scan(&id, &name, &domain)
		fmt.Printf("%-3d %-10s %s\n", id, name, domain)
	}
}

// UPDATE

func Update(db *sql.DB) {
	var id int
	var choice int

	Read(db) // Display current student records.

	fmt.Print("\nEnter the ID of the student you want to update: ")
	fmt.Scan(&id)

	fmt.Println("Choose what to update:")
	fmt.Println("1. Update Name")
	fmt.Println("2. Update Domain")
	fmt.Println("3. Cancel")
	fmt.Print("Enter your choice: ")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		updatename(db, id)
	case 2:
		updatelang(db, id)
	case 3:
		fmt.Println("Update canceled.")
	default:
		fmt.Println("Invalid choice. Update canceled.")
	}
}

func updatename(db *sql.DB, id int) {
	var newName string
	fmt.Println("Enter new name:")
	fmt.Scan(&newName)

	_, err := db.Exec("UPDATE students SET name=$1 WHERE id=$2", newName, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data updated")
	}
}

func updatelang(db *sql.DB, id int) {
	var newDomain string
	fmt.Println("Enter new domain:")
	fmt.Scan(&newDomain)

	_, err := db.Exec("UPDATE students SET domain=$1 WHERE id=$2", newDomain, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data updated")
	}
}

// DELETE

func Delete(db *sql.DB) {
	var id int

	Read(db) // Display current student records.

	fmt.Print("\nEnter the ID of the student you want to delete: ")
	fmt.Scan(&id)

	_, err := db.Exec("DELETE FROM students WHERE id=$1", id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Student record deleted successfully!")
	}
}
