package books

import (
	"time"

	"github.com/kataras/iris/mvc"
)

// Book 书本表
type Book struct {
	ID      int64
	Isbn    string
	Name    string
	CardNum string    `json:"-"`
	Created time.Time `xorm:"created" json:"-" form:"-"`
	Updated time.Time `xorm:"updated" json:"-" form:"-"`
}

type BookService interface {
	GetAll() []Book
	GetByID(id int64) (Book, bool)
	DeleteByID(id int64) bool

	Update(id int64, book Book) (Book, error)

	Create(bookName string, book Book) (Book, error)
}

// BooksController 是 /books controller.
type BooksController struct {
	mvc.C

	Service BookService
}

// // 返回 books列表
// // 例子:
// // curl -i http://localhost:8080/books
// func (c *BooksController) Get() []Book {
// 	GetAll()
// 	var books []Book
// 	if ok, _ := orm.Get(&books); ok {
// 		// ctx.Writef("user found: %#v", books[])
// 	}
// 	return books
// }

// // GetBy 返回一个 movie
// // 例子:
// // curl -i http://localhost:8080/movies/1
// func (c *BooksController) GetBy(id int) Movie {
// 	return movies[id]
// }

// // PutBy 更新一个 movie
// // 例子:
// // curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
// func (c *BooksController) PutBy(id int) Movie {
// 	// 获取一个 movie
// 	m := movies[id]

// 	// 获取一个poster文件
// 	file, info, err := c.Ctx.FormFile("poster")
// 	if err != nil {
// 		c.Ctx.StatusCode(iris.StatusInternalServerError)
// 		return Movie{}
// 	}
// 	file.Close()            // 我们不需要这个文件
// 	poster := info.Filename // 比如这就是上传的文件url
// 	genre := c.Ctx.FormValue("genre")

// 	// 更新poster
// 	m.Poster = poster
// 	m.Genre = genre
// 	movies[id] = m

// 	return m
// }

// // DeleteBy 删除一个 movie
// // 例子:
// // curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
// func (c *BooksController) DeleteBy(id int) iris.Map {
// 	//从movies slice中删除索引
// 	deleted := movies[id].Name
// 	movies = append(movies[:id], movies[id+1:]...)
// 	// 返回删除movie的名称
// 	return iris.Map{"deleted": deleted}
// }
