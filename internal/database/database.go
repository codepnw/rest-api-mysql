package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type MySQLStorage struct {
	db *sql.DB
}

var (
	driver     = "mysql"
	dbname     = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *MySQLStorage
)

func NewMySQLStorage() *MySQLStorage {
	if dbInstance != nil {
		return dbInstance
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	log.Println("connected to MySQL.")

	dbInstance = &MySQLStorage{db: db}

	return dbInstance
}
