package books

import "gorm.io/gorm"

type Respository interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book Book) (Book, error)
}

// private struct
type respository struct {
	db *gorm.DB

}

func NewRepository(db *gorm.DB) *respository{
	return &respository{db}
}

func(r*respository) FindAll() ([]Book, error){
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func(r*respository) FindByID(ID int) (Book, error){
	var book Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func(r*respository) Create(book Book) (Book, error){
	err := r.db.Create(&book).Error
	return book, err
}