package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "swapnil"
	dbName := "go_db"
	dbHost := "192.168.0.170"
	dbPort := "3306"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {

	db := dbConn()

	name := os.Args[1]
	city := os.Args[2]

	insForm, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")

	if err != nil {
		panic(err.Error())
	}

	insForm.Exec(name, city)
	log.Println("INSERT: Name - " + name + " | City - " + city + " SUCCESS")

	defer db.Close()

}
