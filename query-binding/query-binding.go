package main

import (
	"log"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	// Birthday time.Time `json:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func setup(engine *gin.Engine){
	engine.Use(gin.LoggerWithFormatter(func (param gin.LogFormatterParams) string{
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
	engine.Use(gin.Recovery())
}

func main() {
	gin.ForceConsoleColor()
	router := gin.New()
	setup(router)
	router.POST("/testing", startPage)
	router.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	err := c.Bind(&person)
	if err == nil{
		log.Println(person.Name)
		log.Println(person.Address)
		c.JSON(200, gin.H{
		"name": person.Name,
		"address" : person.Address,
		// "birthday" : person.Birthday,
	})
	}else{
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"internal error" : err.Error(),
		})
	}
}