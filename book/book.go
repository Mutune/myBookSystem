// book/book.go
package book

import "fmt"

type Book struct {
	Title  string
	Author string
	ISBN   string
}

var books []Book

// Create a new book
func AddBook(title, author, isbn string) {
	newBook := Book{Title: title, Author: author, ISBN: isbn}
	books = append(books, newBook)
}

// Read all books
func GetAllBooks() []Book {
	return books
}

// Read a book by ISBN
func GetBookByISBN(isbn string) (Book, error) {
	for _, book := range books {
		if book.ISBN == isbn {
			return book, nil
		}
	}
	return Book{}, fmt.Errorf("Book not found with ISBN: %s", isbn)
}

// Update a book's information
func UpdateBook(isbn, title, author string) error {
	for i, book := range books {
		if book.ISBN == isbn {
			books[i] = Book{Title: title, Author: author, ISBN: isbn}
			return nil
		}
	}
	return fmt.Errorf("Book not found with ISBN: %s", isbn)
}

// Delete a book by ISBN
func DeleteBook(isbn string) error {
	for i, book := range books {
		if book.ISBN == isbn {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Book not found with ISBN: %s", isbn)
}
