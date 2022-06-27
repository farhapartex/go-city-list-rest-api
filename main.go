package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func fetchCities(query string, RapidApiKey string, RapidAPIHost string) []byte {
	rapidAPIUrl := "https://andruxnet-world-cities-v1.p.rapidapi.com/?query=" + query + "&searchby=country"
	req, _ := http.NewRequest("GET", rapidAPIUrl, nil)
	req.Header.Add("X-RapidAPI-Key", RapidApiKey)
	req.Header.Add("X-RapidAPI-Host", RapidAPIHost)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	RapidApiKey := os.Getenv("RapidApiKey")
	RapidAPIHost := os.Getenv("RapidAPIHost")

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to City list search"})
	})

	r.GET("/cities", func(c *gin.Context) {
		query := c.DefaultQuery("country", "Bangladesh")
		cities := fetchCities(strings.ToLower(query), RapidApiKey, RapidAPIHost)
		c.Data(http.StatusOK, "application/json", cities)
	})

	r.Run()
}
