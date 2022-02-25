package crud

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetEmpById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("err while mocking")
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"ID", "NAME", "EMAIL", "ROLE"}).
		AddRow(1, "naman", "office@gmail.com", "Intern")

	testcase := []struct {
		desc   string
		input  int
		output user
		mock   interface{}
		erOut  error
	}{
		{
			desc:   "Success",
			input:  1,
			output: user{1, "naman", "office@gmail.com", "Intern"},
			mock:   mock.ExpectQuery("SELECT * FROM EMPLOYEE WHERE ID = ?").WithArgs(1).WillReturnRows(rows),
			erOut:  nil,
		},
		{
			desc:   "Fail",
			input:  4,
			output: user{0, "", "", ""},
			mock:   mock.ExpectQuery("SELECT * FROM EMPLOYEE WHERE ID = ?").WithArgs(4).WillReturnError(sql.ErrNoRows),
			erOut:  sql.ErrNoRows,
		},
	}

	for _, tc := range testcase {
		t.Run("", func(t *testing.T) {
			out, err := GetEmpById(db, tc.input)
			if err != nil && !reflect.DeepEqual(err, tc.erOut) {
				t.Errorf("Expected %v got %v", tc.erOut, err)
			}
			if !reflect.DeepEqual(out, tc.output) {
				t.Errorf("Expected %v got %v", tc.output, out)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("err while mocking")
	}
	defer db.Close()
	testcase := []struct {
		desc  string
		id    int
		mock  interface{}
		erOut error
	}{
		{
			desc:  "Success",
			id:    1,
			mock:  mock.ExpectPrepare("DELETE FROM EMPLOYEE WHERE ID = ?").ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1)),
			erOut: nil,
		},
		{
			desc:  "Failure",
			id:    1,
			mock:  mock.ExpectPrepare("DELETE FROM EMPLOYEE WHERE ID = ?").ExpectExec().WithArgs(1).WillReturnError(errors.New("id not exist")),
			erOut: errors.New("id not exist"),
		},
		{
			desc:  "Prepare Failure",
			id:    1,
			mock:  mock.ExpectPrepare("DELETE FROM EMPLOYEE WHERE ID = ?").WillReturnError(errors.New("fail to prepare")),
			erOut: errors.New("fail to prepare"),
		},
	}

	for _, tc := range testcase {
		t.Run("", func(t *testing.T) {
			err := Delete(db, tc.id)
			if err != nil && !reflect.DeepEqual(err, tc.erOut) {
				t.Errorf("Expected %v got %v", tc.erOut, err)
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("err while mocking")
	}
	defer db.Close()
	testcase := []struct {
		desc  string
		id    int
		col   string
		val   string
		mock  interface{}
		erOut error
	}{
		{
			desc:  "Success",
			id:    3,
			col:   "ROLE",
			val:   "SDE",
			mock:  mock.ExpectPrepare("UPDATE EMPLOYEE SET ROLE = ? WHERE ID = ?").ExpectExec().WithArgs("SDE", 3).WillReturnResult(sqlmock.NewResult(0, 1)),
			erOut: nil,
		},
		{
			desc:  "Fail",
			id:    2,
			col:   "ROLE",
			val:   "SDE",
			mock:  mock.ExpectPrepare("UPDATE EMPLOYEE SET ROLE = ? WHERE ID = ?").ExpectExec().WithArgs("SDE", 2).WillReturnError(errors.New("ID doesn't exist")),
			erOut: errors.New("ID doesn't exist"),
		},
		{
			desc:  "Prepare Failure",
			id:    1,
			mock:  mock.ExpectPrepare("DELETE FROM EMPLOYEE WHERE ID = ?").WillReturnError(errors.New("fail to prepare")),
			erOut: errors.New("fail to prepare"),
		},
	}
	for _, tc := range testcase {
		t.Run("", func(t *testing.T) {
			err := Update(db, tc.id, tc.col, tc.val)
			if err != nil && !reflect.DeepEqual(err, tc.erOut) {
				t.Errorf("Expected %v got %v", tc.erOut, err)
			}
		})
	}
}
func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("err while mocking")
	}
	defer db.Close()

	testcase := []struct {
		desc  string
		id    int
		name  string
		email string
		role  string
		mock  interface{}
		erOut error
	}{
		{
			desc:  "Success",
			id:    1,
			name:  "naman",
			email: "fake@gmail.com",
			role:  "SDE",
			mock:  mock.ExpectPrepare("INSERT INTO EMPLOYEE(ID,NAME,EMAIL,ROLE) VALUES(?,?,?,?)").ExpectExec().WithArgs(1, "naman", "fake@gmail.com", "SDE").WillReturnResult(sqlmock.NewResult(0, 1)),
			erOut: nil,
		},
		{
			desc:  "Failure",
			id:    3,
			name:  "naman",
			email: "fake@gmail.com",
			role:  "SDE",
			mock:  mock.ExpectPrepare("INSERT INTO EMPLOYEE(ID,NAME,EMAIL,ROLE) VALUES(?,?,?,?)").ExpectExec().WithArgs(3, "naman", "fake@gmail.com", "SDE").WillReturnError(errors.New("DUPLICATE ID")),
			erOut: errors.New("DUPLICATE ID"),
		},
		{
			desc:  "Prepare Failure",
			id:    1,
			mock:  mock.ExpectPrepare("DELETE FROM EMPLOYEE WHERE ID = ?").WillReturnError(errors.New("fail to prepare")),
			erOut: errors.New("fail to prepare"),
		},
	}
	for _, tc := range testcase {
		t.Run("", func(t *testing.T) {
			err := Insert(db, tc.id, tc.name, tc.email, tc.role)
			if err != nil && !reflect.DeepEqual(err, tc.erOut) {
				t.Errorf("Expected %v got %v", tc.erOut, err)
			}
		})
	}
}
