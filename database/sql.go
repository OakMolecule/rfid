package sql

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// User 用户信息表
type User struct {
	ID      int64
	Name    string
	CardNum string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

// Book 书本表
type Book struct {
	ID   int64
	Isbn string
}
