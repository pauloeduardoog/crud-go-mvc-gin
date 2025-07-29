package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book
var nextID = 11 // começa do 11 porque já temos 10 livros

func init() {
	books = []Book{
		{ID: 1, Title: "O Senhor dos Anéis", Author: "J.R.R. Tolkien"},
		{ID: 2, Title: "Dom Quixote", Author: "Miguel de Cervantes"},
		{ID: 3, Title: "1984", Author: "George Orwell"},
		{ID: 4, Title: "A Revolução dos Bichos", Author: "George Orwell"},
		{ID: 5, Title: "Cem Anos de Solidão", Author: "Gabriel García Márquez"},
		{ID: 6, Title: "Orgulho e Preconceito", Author: "Jane Austen"},
		{ID: 7, Title: "O Pequeno Príncipe", Author: "Antoine de Saint-Exupéry"},
		{ID: 8, Title: "O Hobbit", Author: "J.R.R. Tolkien"},
		{ID: 9, Title: "Moby Dick", Author: "Herman Melville"},
		{ID: 10, Title: "O Código Da Vinci", Author: "Dan Brown"},
	}
}

func GetBooks() []Book {
	return books
}

func GetBookByID(id int) (Book, bool) {
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}

func CreateBook(book Book) Book {
	book.ID = nextID
	nextID++
	books = append(books, book)
	return book
}

func UpdateBook(id int, updated Book) (Book, bool) {
	for i, b := range books {
		if b.ID == id {
			updated.ID = id
			books[i] = updated
			return updated, true
		}
	}
	return Book{}, false
}

func DeleteBook(id int) bool {
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}
	return false
}
