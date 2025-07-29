package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"crud-gin-mvc/models"
)

// GET /books
func GetBooks(c *gin.Context) {
	books := models.GetBooks()
	c.JSON(http.StatusOK, books)
}

// GET /books/:id
func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	book, ok := models.GetBookByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// POST /books
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	created := models.CreateBook(book)
	c.JSON(http.StatusCreated, created)
}

// PUT /books/:id
func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	updatedBook, ok := models.UpdateBook(id, book)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

// DELETE /books/:id
func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ok := models.DeleteBook(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro removido"})
}
