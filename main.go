package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

var testData = []Data{
	{Key: "1", Value: "TestValue" },
}

func main() {
	  router := gin.Default()

		router.GET("/", getData)
		router.POST("/", postData)
		router.GET("/:id", getDataById)

		router.Run("localhost:8080")
}

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testData)
}

func getDataById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range testData {
			if a.Key == id {
				c.IndentedJSON(http.StatusOK, a)
			}
		}
}

func postData(c *gin.Context) {
	var newData Data

	if err := c.BindJSON(&newData); err != nil {
		return
	}

	testData = append(testData, newData)

	c.IndentedJSON(http.StatusCreated, newData)
}
