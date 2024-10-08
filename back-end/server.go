package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func init_server() {
	os.LoadEnv()
	fmt.Println("Initializing server...")

	err := init_db()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		log.Fatal(err)
		return
	}

	// Initialize router go get github.com/joho/godotenv
	router := mux.NewRouter()
	// Define routes
	router.HandleFunc("/api", apiHandler).Methods("GET")
	router.HandleFunc("/api/health", healthHandler).Methods("GET")
	// Start server
	http.ListenAndServe(":8080", router)

	fmt.Println("Server started on port 8080")
}

func init_db() error {
	fmt.Println("Initializing database...")
	// Database initialization code here
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DATABASE_NAME"))

	if err != nil {
		return err
	}
	defer conn.Close()

	sql := "CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY, name TEXT)"

	_, err = conn.Exec(sql)
	if err != nil {
		return err
	}

	fmt.Println("Database initialized")
	return nil
}
