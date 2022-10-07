package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/books"
)

//mengakses service
type bookHandler struct{
	bookService books.Service
}

func NewBookHandler(bookService books.Service) *bookHandler{
	return &bookHandler{bookService}
}


/*POST*/
//function handler query untuk post
func (handler *bookHandler)PostBookHandler(c *gin.Context){
	//membuat variable input
	var bookInput books.BookRequest

	err := c.ShouldBindJSON(&bookInput)
	if err != nil{
		//slice err
		errorMessages := []string{}
		//validation error
		for _, e:= range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on %s, where condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			c.JSON(http.StatusBadRequest, gin.H{
				"errors" : errorMessages,
			})
		}
	}else {
		//status 201 untuk post
		book, err := handler.bookService.Create(bookInput)
		if err !=nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors" : err,
			})
		}else{
			c.JSON(http.StatusCreated, gin.H{
				"data" : book,
			})
		}
	}
}

/*GET*/
//function rooter
func (handler *bookHandler)RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Desta",
		"status" : "Welcome to My API withh Golang-Gin Library",
	})
}

//function handler read book all
func (handler *bookHandler)GetBooksHandler(c *gin.Context){
	book, err := handler.bookService.FindAll()
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
		return
	}

	//var get response bentuk json
	var bookRes []books.BooksRequestResponse
	for _, b := range book{
		booksRes := convertToBookResponse(b)
		bookRes= append(bookRes, booksRes)
	}

	c.JSON(http.StatusOK, gin.H{"data": bookRes})
}

//function handler read/get single book
func (handler *bookHandler)GetBookHandler(c *gin.Context){
	idStr := c.Param("id")
	//casting from string to int
	id, _ := strconv.Atoi(idStr)

	book, err := handler.bookService.FindByID(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
		return
	}

	bookRes := convertToBookResponse(book)
	c.JSON(http.StatusOK, gin.H{"data": bookRes})
}

/*PUT*/
func (handler *bookHandler) PutBookHandler(c *gin.Context){
	//membuat variable input
	var bookInput books.BookRequest

	err := c.ShouldBindJSON(&bookInput)
	if err != nil{
		//slice err
		errorMessages := []string{}
		//validation error
		for _, e:= range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on %s, where condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			c.JSON(http.StatusBadRequest, gin.H{
				"errors" : errorMessages,
			})
		}
	}

	//Get the ID
	idStr := c.Param("id")
	//casting from string to int
	id, _ := strconv.Atoi(idStr)

	//status 204 untuk put
	book, err := handler.bookService.Update(id, bookInput)
	if err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"data" : book,
	})
}

// func (handler *bookHandler)QueryHandler(c *gin.Context){
// 	//mengambil query id
// 	id := c.Query("id")

// 	c.JSON(http.StatusOK, gin.H{"id": id})
// }

// //function handler query untuk title dan price
// func (handler *bookHandler)QueryHandlers(c *gin.Context){
// 	//mengambil query title dan price
// 	title := c.Query("title")
// 	price := c.Query("price")
	
// 	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
// }


//private function for response json format
func convertToBookResponse(b books.Book) books.BooksRequestResponse{
	return books.BooksRequestResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
		Discount:	 b.Discount,
	}
}