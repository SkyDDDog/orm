package session

import (
	"testing"
)

type User struct {
	Name string `orm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	//engine, _ := orm.NewEngine("mysql", "root:9738faq@(127.0.0.1:3306)/test")
	//fmt.Println(engine)
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}

}
