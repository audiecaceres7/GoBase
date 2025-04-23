package main

import (
	"context"
	repo "go_base/audiecaceres7/internal/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/subosito/gotenv"
)

const (
    PORT = ":4000"
    TEMPLATES_PATH = "./templates/*"
    STATIC_PATH = "./static"
)

type AppConfig struct {
    port string
    Repo repo.Queries
}

func main() {
    err := gotenv.Load()
    if err != nil {
        log.Printf("ERROR: couldn't load env, %v", err) 
    }
    
    ctx := context.Background()
    conn, err := pgx.Connect(ctx, os.Getenv("DB_CONN_STRING"))
    if err != nil {
        log.Printf("ERROR: couldn't connect to database %v", err) 
    }

    appCofig := AppConfig{
        port: PORT,
        Repo: *repo.New(conn),
    }

    router := gin.Default()
    router.Use(func(c *gin.Context) {
        c.Set("config", appCofig) 
        c.Next()
    })

    router.LoadHTMLGlob(TEMPLATES_PATH)
    router.Static("/static", STATIC_PATH)

    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index", struct{ Title string }{ Title: "World" })
    })

    log.Printf("Serving on port %s\n", PORT)
    router.Run(PORT)
}
