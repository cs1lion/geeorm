package session

import (
	"database/sql"
	"fmt"
	"geeorm/dialect"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestMain(m *testing.M) {
	fmt.Println("Initing test...")
	TestDB, _ = sql.Open("sqlite3", "../gee.db")
	code := m.Run()
	fmt.Println("Tear down after tests")
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
