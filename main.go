package main

import (
	"time"

	"github.com/go-xorm/xorm"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/basicauth"

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

var orm *xorm.Engine

var basicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})

func newApp() *iris.Application {
	app := iris.New()
	// 设置debug的调试级别
	app.Logger().SetLevel("debug")
	app.Use(basicAuth)

	app.Controller("/books", new(BooksController))

	return app
}

func main() {
	app := newApp()

	var err error
	orm, err = xorm.NewEngine("mysql", "root:root@/rfid?charset=utf8")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}

	iris.RegisterOnInterrupt(func() {
		orm.Close()
	})

	err = orm.Sync2(new(User))
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}
	err = orm.Sync2(new(Book))
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized Book table: %v", err)
	}

	app.Run(iris.Addr(":8080"))
}

func h(ctx context.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}
