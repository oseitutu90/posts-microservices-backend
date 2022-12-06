// Description: Hello World in Go

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql" // need to install this
	"gorm.io/gorm"
)

type Post struct {
	Id          uint
	Title       string
	Description string
}

func main() {

	// https://github.com/go-sql-driver/mysql
	dsn := "root:Geforce229!@tcp(localhost:3306)/posts_ms?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Pass the connection to the database to the GORM

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Post{})

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World  postings here👋!")
	})
	app.Listen(":8000")
}
