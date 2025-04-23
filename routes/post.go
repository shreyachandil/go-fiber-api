package routes

import (
	"go-fiber-api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var posts = []models.Post{
	{ID: 1, Title: "First Post", Content: "Hello, API world!"},
}

func SetupPostRoutes(app *fiber.App) {
	api := app.Group("/posts")

	api.Get("/", getPosts)
	api.Get("/:id", getPost)
	api.Post("/", createPost)
	api.Put("/:id", updatePost)
	api.Delete("/:id", deletePost)
}

func getPosts(c *fiber.Ctx) error {
	return c.JSON(posts)
}

func getPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for _, p := range posts {
		if p.ID == id {
			return c.JSON(p)
		}
	}
	return c.Status(404).SendString("Post not found")
}

func createPost(c *fiber.Ctx) error {
	var newPost models.Post
	if err := c.BodyParser(&newPost); err != nil {
		return err
	}
	newPost.ID = len(posts) + 1
	posts = append(posts, newPost)
	return c.Status(201).JSON(newPost)
}

func updatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updated models.Post
	if err := c.BodyParser(&updated); err != nil {
		return err
	}
	for i, p := range posts {
		if p.ID == id {
			updated.ID = id
			posts[i] = updated
			return c.JSON(updated)
		}
	}
	return c.Status(404).SendString("Post not found")
}

func deletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			return c.SendStatus(204)
		}
	}
	return c.Status(404).SendString("Post not found")
}
