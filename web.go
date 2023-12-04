// web.go
package main

import (
	"myBookSystem/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define routes
	r.GET("/", showBooks)
	r.GET("/add", showAddForm)
	r.POST("/add", addBook)
	r.GET("/update/:isbn", showUpdateForm)
	r.POST("/update/:isbn", updateBook)
	r.GET("/delete/:isbn", deleteBook)

	// Run the server
	r.Run(":8080")
}

func showBooks(c *gin.Context) {
	allBooks := book.GetAllBooks()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"books": allBooks})
}

func showAddForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add.tmpl", nil)
}

func addBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	isbn := c.PostForm("isbn")

	book.AddBook(title, author, isbn)

	c.Redirect(http.StatusSeeOther, "/")
}

func showUpdateForm(c *gin.Context) {
	isbn := c.Param("isbn")
	b, err := book.GetBookByISBN(isbn)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.tmpl", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "update.tmpl", gin.H{"book": b})
}

func updateBook(c *gin.Context) {
	isbn := c.Param("isbn")
	title := c.PostForm("title")
	author := c.PostForm("author")

	err := book.UpdateBook(isbn, title, author)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func deleteBook(c *gin.Context) {
	isbn := c.Param("isbn")

	err := book.DeleteBook(isbn)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}
