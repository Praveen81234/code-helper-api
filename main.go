package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CodeRequest struct {
	Code string `json:"code"`
}

type QueryRequest struct {
	Query string `json:"query"`
}

type RunResponse struct {
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

type AutoFixResponse struct {
	FixedCode string `json:"fixed_code"`
}

type HelpResponse struct {
	Tip string `json:"tip"`
}

func main() {
	app := fiber.New()

	// 1. Run Code API - POST /run
	app.Post("/run", func(c *fiber.Ctx) error {
		req := new(CodeRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(RunResponse{
				Error: "invalid request body",
			})
		}

		code := strings.TrimSpace(req.Code)

		if code == "" {
			return c.JSON(RunResponse{
				Error: "no code provided",
			})
		}

		// Simulation
		if strings.Contains(strings.ToLower(code), "error") {
			return c.JSON(RunResponse{
				Error:  "simulated: code has an error",
				Output: "",
			})
		}

		return c.JSON(RunResponse{
			Output: "simulated: code ran successfully",
			Error:  "",
		})
	})

	// 2. Auto-Fix API - POST /autofix
	app.Post("/autofix", func(c *fiber.Ctx) error {
		req := new(CodeRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		code := req.Code
		lines := strings.Split(code, "\n")
		for i, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" {
				continue
			}
			if !strings.HasSuffix(trimmed, ";") &&
				!strings.HasSuffix(trimmed, "{") &&
				!strings.HasSuffix(trimmed, "}") {

				trimmed = trimmed + ";"
			}
			lines[i] = "    " + trimmed
		}
		fixed := strings.Join(lines, "\n")
		fixed = strings.ReplaceAll(fixed, "  ", " ")

		return c.JSON(AutoFixResponse{
			FixedCode: fixed,
		})
	})

	// 3. Help API - POST /help
	app.Post("/help", func(c *fiber.Ctx) error {
		req := new(QueryRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		query := strings.ToLower(req.Query)

		tip := "General tip: Break your code into small parts, print intermediate values, and read error messages carefully."

		if strings.Contains(query, "syntax") {
			tip = "Syntax tip: Check for missing brackets, commas, and semicolons."
		} else if strings.Contains(query, "loop") {
			tip = "Loop tip: Make sure your loop has an exit condition."
		} else if strings.Contains(query, "variable") {
			tip = "Variables tip: Use meaningful names and declare before use."
		} else if strings.Contains(query, "error") {
			tip = "Error tip: Read the error slowly and search exactly that message."
		}

		return c.JSON(HelpResponse{
			Tip: tip,
		})
	})

	log.Println("Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
