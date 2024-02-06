package bin

import (
	"fmt"
	"log"
	"starter-go-fiber/libs"
	"starter-go-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type App struct {
	Router *fiber.App
}

func (a *App) Initialize() {
	engine := html.New("./views", ".html")

	a.Router = fiber.New(fiber.Config{
		Views: engine,
	})
	a.setRoutes()
}

func (a *App) setRoutes() {
	routes.SetUp(a.Router)
}

func (a *App) Run() {
	router := a.Router
	port := fmt.Sprint(":", libs.EnvVariable("APP_PORT"))
	if libs.EnvVariable("APP_PORT") == "" {
		port = ":3000"
	}

	err := router.Listen(port)

	if err != nil {
		log.Panic("Failed to start the server:", err)
	}
}

func NewApp() *App {
	app := App{}
	app.Initialize()

	return &app
}
