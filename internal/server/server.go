package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codepnw/rest-api-mysql/internal/database"
)

type Server struct {
	port int
	db   database.MySQLStorage
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   *database.NewMySQLStorage(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

