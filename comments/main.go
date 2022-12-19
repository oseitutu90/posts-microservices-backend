package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql" // need to install this
	"gorm.io/gorm"
)

type Comment struct { // Comment model
	Id     uint   `json:"id"`
	PostId uint   `json: "postId"`
	Text   string `json: "text"`
}

func main() {
	dsn := "root:Geforce229!@tcp(localhost:3306)/comments_ms?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Pass the connection to the database to the GORM

	if err != nil { // Check for errors
		panic(err)
	}

	db.AutoMigrate(&Comment{}) // Migrate the schema

	app := fiber.New()

	app.Use(cors.New()) // allow cors

	app.Post("/api/posts/:id/comments", func(c *fiber.Ctx) error {
		var comments []Comment

		db.Find(&comments, "post_id = ?", c.Params("id")) // Get all comments

		return c.JSON(comments)
	})

	app.Post("/api/comments", func(c *fiber.Ctx) error {
		var comment Comment                            // Create a comment
		if err := c.BodyParser(&comment); err != nil { // Parse the body of the request
			return err
		} // Create comment

		db.Create(&comment)

		return c.JSON(comment) // Return the comment
	}) // Get all comments

	app.Listen(":8001")
}
