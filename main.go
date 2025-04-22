package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

const (
    PORT = ":4000"
    TEMPLATES_PATH = "./templates/*"
)

type AppConfig struct {
    port string
}

func main() {
    err := gotenv.Load; if err != nil {
        log.Printf("ERROR: couldn't load env, %v", err) 
    }

    router := gin.Default()
    router.LoadHTMLGlob(TEMPLATES_PATH)

    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index", nil)
    })

    log.Printf("Serving on port %s\n", PORT)
    router.Run(PORT)
}
