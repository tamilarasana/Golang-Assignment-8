package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	  "gorm.io/gorm"
	  "log"
	  
)

func main() {

  dsn := "root:admin@tcp(127.0.0.1:3306)/tamil?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  // db, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/tamil?charset=utf8mb4&parseTime=True&loc=Local")

   if err != nil{

   	 log.Fatal("Error Conncting database")
   }

   defer db.Close()

	r := gin.Default()
	r.GET("/posts",Posts)
	r.GET("/posts/:id",Show)
	r.POST("/posts", Store)
	r.PATCH("/posts/:id", Update)
	r.DELETE("posts/:id",Delete) 
	r.Run(":123") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

 

 func Posts(c *gin.Context) {

 }

 func Show(c *gin.Context) {
 	
 }

 func Store(c *gin.Context) {
 	
 }

 func Update(c *gin.Context) {
 	
 }

 func Delete(c *gin.Context) {
 	
 }