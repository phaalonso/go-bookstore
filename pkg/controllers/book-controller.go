package controllers

import (
	"encoding/json"
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
	err := utils.ParseBody(r, createBook)

	if err != nil {
		ErrorResponse(http.StatusPreconditionFailed, "Invalid json", w)
		return
	}

	b := createBook.CreateBook()

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
	err := utils.ParseBody(r, updateBook)

	if err != nil {
		ErrorResponse(http.StatusPreconditionFailed, "Invalid json", w)
		return
	}

	ID, err := utils.ExtractParamId(r, "bookId")

	if err != nil {
		ErrorResponse(http.StatusPreconditionFailed, "Incorrect param", w)
		return
	}

	book, db := models.GetBookById(ID)

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
		utils.SendMessage(w, http.StatusInternalServerError, []byte("An error occurred internally"))
		return
	}

	utils.SendJson(w, statusCode, message)
}
