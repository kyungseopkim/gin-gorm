package main

import (
	"fmt"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	dsn := "host=localhost user=user password=password dbname=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//db.AutoMigrate(&User{})
	r := gin.New()
	r.Use(logger.SetLogger(
		logger.WithUTC(true)))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	r.PUT("/api/v1/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login",
		})
	})
	err := r.Run()
	if err != nil {
		os.Exit(1)
	} // listen and serve on
}
