package books

type Service interface{
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
}

type service struct{
	respository Respository
}

func NewService(respository Respository) *service{
	return &service{respository}
}

func(s*service) FindAll() ([]Book, error){
	books, err := s.respository.FindAll()
	return books, err
}

func(s*service) FindByID(ID int) (Book, error){
	book, err := s.respository.FindByID(ID)
	return book, err
}

func(s*service) Create(bookRequest BookRequest) (Book, error){
	price, _ := bookRequest.Price.Int64()

	book := Book{
		Title : bookRequest.Title,
		Price : int(price),
	}

	newBook, err := s.respository.Create(book)
	return newBook, err
}