package router

import (
	"Health-Check/controller/health"
	"Health-Check/db"
	"Health-Check/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	connString := os.Getenv("POSTGRES_CONN_STR")
	connString = connString + "?sslmode=disable"
	postgresRepo, err := db.NewPostgreSQL(connString)
	if err != nil {
		log.Println("could not connect to database with err : %v", err)
	}

	healthService := service.NewHealthService(postgresRepo)

	healthController := health.NewHealthController(healthService)
	router.GET("/healthz", healthController.GetHealth)

	router.Use(func(context *gin.Context) {
		if context.Request.URL.Path == "/healthz" && context.Request.Method != http.MethodGet {
			context.JSON(http.StatusMethodNotAllowed, nil)
			context.Abort()
		}
	})

	return router
}
