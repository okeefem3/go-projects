package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func buildPostgresDBSource() (string, error) {
	// In a real project I would have a package devoted to config values, and get these from that. But for now this is good enough
	if _, err := os.Stat("../../.env"); err == nil {
		fmt.Println("Env file exists, attempting to load env file")
		err = godotenv.Load("../../.env")

		if err != nil {
			fmt.Println("Error loading env file")
			return "", err
		}
	} else {
		fmt.Println("Env file does not exist, lets hope the env vars exist")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")

	conn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, db, sslMode)

	fmt.Println(conn)

	return conn, nil
}
