package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "databases.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	password := "'' OR 1=1"

	segura(db, password)

	fmt.Println("-------------------------")

	insegura(db, password)
}

func segura(db *sql.DB, password string) {

	rows, err := db.Query("SELECT * FROM users WHERE username = 'cesar' AND password = ?", password)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var u, p string

		rows.Scan(&u, &p)

		fmt.Println(u, p)
	}

}

func insegura(db *sql.DB, password string) {

	s := fmt.Sprintf("SELECT * FROM users WHERE username = 'cesar' AND password = %s", password)

	rows, err := db.Query(s)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for rows.Next() {
		var u, p string

		rows.Scan(&u, &p)

		fmt.Println(u, p)
	}
}
