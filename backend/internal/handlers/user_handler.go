package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(f http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(f).Encode(users)
	}
}
