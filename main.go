package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)

	router.GET("/adjeng", adjengHandler)
	router.GET("/quotes/:id/:title", quotesHandler) //variable id
	router.GET("/query", queryHandler)              //
	router.POST("/quote", postQuoteHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":      "TB Maulana Aghni",
		"horoscope": "Gemini",
	})
}
func adjengHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":      "Diadjeng",
		"horoscope": "Aries",
	})
}

// URL PARAMETER / PATH HANDLER
func quotesHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	year := c.Query("year")
	c.JSON(http.StatusOK, gin.H{"title": title, "year": year})
}

type QuoteInput struct {
	Quote string
	Name  string
	Year  string
	// if there an underscore , use 'json:"blabla_blabla"' after data type (keep using the '')
}

func postQuoteHandler(c *gin.Context) {
	// quote, name, year

	var quoteInput QuoteInput

	err := c.ShouldBindJSON(&quoteInput)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"quote": quoteInput.Quote,
		"name":  quoteInput.Name,
		"year":  quoteInput.Year,
	})
}
