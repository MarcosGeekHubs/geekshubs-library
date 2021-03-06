package api

import (
	"strconv"

	"github.com/rubencougil/geekshubs-library/pkg/library"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// API struct
type API struct {
	Logger *logrus.Logger
}

// NewAPI return api
func NewAPI(logger *logrus.Logger) *API {
	return &API{
		Logger: logger,
	}
}

// Hello hello
func (a *API) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// GetAllBooks get all books
func (a *API) GetAllBooks(lib *library.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		books, err := lib.All()
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, books)
	}
	return gin.HandlerFunc(fn)
}

// AddBook add book
func (a *API) AddBook(lib *library.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var book library.Book
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		id, err := lib.Add(book)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(202, gin.H{
			"id": id,
		})
	}
	return gin.HandlerFunc(fn)
}

// UpdateBook update book
func (a *API) UpdateBook(lib *library.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		var book library.Book
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		id, err = lib.Modify(id, book)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(202, gin.H{
			"id": id,
		})
	}
	return gin.HandlerFunc(fn)
}

// GetBook get book
func (a *API) GetBook(library *library.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		book, err := library.Find(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, book)
	}
	return gin.HandlerFunc(fn)
}

// DeleteBook delete book
func (a *API) DeleteBook(library *library.Library) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		err := library.Delete(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(202, gin.H{
			"id": id,
		})
	}
	return gin.HandlerFunc(fn)
}
