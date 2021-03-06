package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()
	r.Any("/testing", func(c *gin.Context) {
		var person Person
		if c.ShouldBindQuery(&person) == nil {
			log.Println("====== Only Bind By Query String ======")
			log.Println(person.Name)
			log.Println(person.Address)
		}

		c.String(http.StatusOK, "Success")
	})

	r.Run(":8080")
}
