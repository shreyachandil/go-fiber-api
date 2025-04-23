package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Analytics() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		userAgent := c.Get("User-Agent")
		endpoint := c.OriginalURL()
		method := c.Method()
		timestamp := time.Now().Format(time.RFC3339)

		fmt.Printf("[%s] %s %s | IP: %s | UA: %s\n", timestamp, method, endpoint, ip, userAgent)

		return c.Next()
	}
}
