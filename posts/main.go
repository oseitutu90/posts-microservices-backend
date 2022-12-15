// Description: Hello World in Go

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql" // need to install this driver
	"gorm.io/gorm"
)

type Post struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {

	// https://github.com/go-sql-driver/mysql
	dsn := "root:Geforce229!@tcp(localhost:3306)/posts_ms?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Pass the connection to the database to the gorm.Open function

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Post{})

	app := fiber.New() // create a new fiber app

	app.Use(cors.New()) // enable cors

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		var posts []Post // create a slice of posts

		db.Find(&posts) // find all posts

		return c.JSON(posts) // return all posts
	})

	app.Post("/api/posts", func(c *fiber.Ctx) error {
		var post Post // create a post

		if err := c.BodyParser(&post); err != nil { // parse the body of the request
			return err
		}

		db.Create(&post) // create a post

		return c.JSON(post) // return the post
	})

	app.Listen(":8000")
}
