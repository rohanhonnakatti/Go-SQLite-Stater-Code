package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open the SQLite database file (create it if it doesn't exist)
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	
	// Ensure the database connection is closed when done
	defer db.Close()

	// Create a new table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Students (
        ID INTEGER PRIMARY KEY,
        Name TEXT,
        Age INTEGER,
        Grade REAL
    )`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Insert a new record
	_, err = db.Exec("INSERT INTO Students (Name, Age, Grade) VALUES (?, ?, ?)", "rohanyh", 24, 8.23)
	if err != nil {
		fmt.Println("Error inserting record:", err)
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM Students")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	
	// Ensure the result set is closed when done
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		var grade float64

		if err := rows.Scan(&id, &name, &age, &grade); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		fmt.Printf("ID: %d, Name: %s, Age: %d, Grade: %.2f\n", id, name, age, grade)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over result set:", err)
		return
	}
}
