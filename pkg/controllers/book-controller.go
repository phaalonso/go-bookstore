package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/phaalonso/book-store/pkg/models"
	"github.com/phaalonso/book-store/pkg/utils"
	"net/http"
)

func GetBook(w http.ResponseWriter, _ *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)

	utils.SendJson(w, http.StatusOK, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.ExtractParamId(r, "bookId")

	if err != nil {
		ErrorResponse(http.StatusPreconditionFailed, "Incorrect param", w)
		return
	}

	bookDetails, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)

	utils.SendJson(w, http.StatusOK, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)

	b := createBook.CreateBook()

	fmt.Println(b)

	res, _ := json.Marshal(b)

	utils.SendJson(w, http.StatusCreated, res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.ExtractParamId(r, "bookId")

	if err != nil {
		ErrorResponse(http.StatusPreconditionFailed, "Incorrect param", w)
		return
	}

	deletedBook := models.DeleteBook(ID)

	res, _ := json.Marshal(deletedBook)

	utils.SendJson(w, http.StatusNoContent, res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	ID, err := utils.ExtractParamId(r, "bookId")

	if err != nil {
		ErrorResponse(http.StatusPreconditionFailed, "Incorrect param", w)
		return
	}

	book, db := models.GetBookById(ID)

	// TODO existe alguma maneira de mesclar objetos?
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)

	res, _ := json.Marshal(book)

	utils.SendJson(w, http.StatusOK, res)
}

func ErrorResponse(statusCode int, error string, w http.ResponseWriter) {
	//Create a new map and fill it
	fields := make(map[string]interface{})
	fields["status"] = "error"
	fields["message"] = error
	message, err := json.Marshal(fields)

	if err != nil {
		//An error occurred processing the json
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occured internally"))

		return
	}

	utils.SendJson(w, statusCode, message)
}
