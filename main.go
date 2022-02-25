package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iamkakashi/db/crud"
)

const (
	username = "root"
	password = "password"
	hostname = "127.0.0.1:3306"
	dbname   = "demo"
)

func dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func Create(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS EMPLOYEE (ID INT NOT NULL, NAME VARCHAR(100), EMAIL VARCHAR(100), ROLE VARCHAR(100),PRIMARY KEY(ID))")
	if err != nil {
		fmt.Printf("Error %v while creating table", err)
	}
}

func main() {

	// CONNECTION
	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		fmt.Printf("Error %v while creating database", err)
	}
	defer db.Close()

	Create(db)

	err = crud.Insert(db, 2, "Goku", "saiyan@gmail.com", "SDE")
	fmt.Println(err)
	err = crud.Update(db, 8, "ROLE", "Intern")
	fmt.Println(err)
	res, err := crud.GetEmpById(db, 6)
	fmt.Println(res, " ", err)
	err = crud.Delete(db, 4)
	fmt.Println(err)

}
