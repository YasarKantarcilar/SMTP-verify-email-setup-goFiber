package main

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	T "smtpservice.com/Models"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/verify", func(c *fiber.Ctx) error {
		var payload T.Email

		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		err := sendMail([]string{payload.To})
		if err != nil {
			return c.SendStatus(500)
		}

		return c.Status(fiber.StatusOK).SendString("Email sent")
	})

	app.Get("/verify/:token", func(c *fiber.Ctx) error {
		token := c.Params("token")
		fmt.Println(token)
		return c.SendString("Email verified")
	})

	log.Fatal(app.Listen(":3000"))

	fmt.Println("Hello, world!")

}

func sendMail(sendTo []string) error {
	// Sender data.
	from := "<senderemail>"
	password := "<senderpassword>"

	// Receiver email address.

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := `<html><body>
	<h1>Test email from Go!</h1>
	<a style="display:grid; place-items:center; width: 200px; padding: 20px; background-color:black; color:white;" href="http://localhost:3000/verify/123456789">Verify</a>
	</body></html>`
	message := []byte(subject + mime + body)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, sendTo, []byte(message))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
