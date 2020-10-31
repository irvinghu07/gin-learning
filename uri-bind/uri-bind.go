package main

import "github.com/gin-gonic/gin"

type Person struct{
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main(){
	router := gin.Default()
	router.GET(":name/:id", func(c *gin.Context){
		var person Person
		if err := c.ShouldBindUri(&person); err != nil{
			c.JSON(400, gin.H{
				"status" : err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"ID" : person.ID,
			"Name" : person.Name,
		})
	})
	router.Run(":8088")
}