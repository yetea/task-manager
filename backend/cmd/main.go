package main

import (
	"api/internal/handlers"
	"api/pkg/router"
	"database/sql"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := router.Router()

	router.HandleFunc("/api/go/users", handlers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/api/go/users", handlers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/api/go/users/{id}", handlers.UpdatedUser(db)).Methods("PUT")
	router.HandleFunc("/api/go/users/{id}", handlers.DeleteUser(db)).Methods("DELETE")

	enhancedRouter := handlers.EnableCORS(handlers.JsonContentTypeMiddleware(router))

	log.Fatal(http.ListenAndServe(":8080", enhancedRouter))
}
