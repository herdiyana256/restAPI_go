package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load() // Load environment variables from .env file (optional)
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	if config.DBHost == "" || config.DBPort == "" || config.DBUser == "" ||
		config.DBPassword == "" || config.DBName == "" {
		return nil, fmt.Errorf("missing required environment variables for database connection")
	}

	return config, nil
}
