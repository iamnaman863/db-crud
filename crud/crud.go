package crud

import (
	"database/sql"
	"errors"
	"fmt"
)

type user struct {
	id    int
	name  string
	email string
	role  string
}

func (u user) String() string {
	return fmt.Sprintf("%v %v %v %v", u.id, u.name, u.email, u.role)
}

func Insert(db *sql.DB, id int, name string, email string, role string) error {

	stmt, err := db.Prepare("INSERT INTO EMPLOYEE(ID,NAME,EMAIL,ROLE) VALUES(?,?,?,?)")

	if err != nil {
		return errors.New("fail to prepare")
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, name, email, role)
	if err != nil {
		err = errors.New("DUPLICATE ID")
	}
	return err

}
func GetEmpById(db *sql.DB, val int) (user, error) {

	row := db.QueryRow("SELECT * FROM EMPLOYEE WHERE ID = ?", val)
	use := user{}

	err := row.Scan(&use.id, &use.name, &use.email, &use.role)

	if err != nil {
		return use, err
	}
	fmt.Println("Got User :  ", use)
	return use, nil
}

func Update(db *sql.DB, id int, col string, val string) error {

	stmt, err := db.Prepare("UPDATE EMPLOYEE SET " + col + " = ? WHERE ID = ?")
	if err != nil {
		return errors.New("fail to prepare")
	}
	defer stmt.Close()
	_, err = stmt.Exec(val, id)
	if err != nil {
		err = errors.New("ID doesn't exist")
	}
	return err

}
func Delete(db *sql.DB, id int) error {

	stmt, err := db.Prepare("DELETE FROM EMPLOYEE WHERE ID = ?")
	if err != nil {
		return errors.New("fail to prepare")
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		err = errors.New("id not exist")
	}
	return err
}
