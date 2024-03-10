package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

type Database struct {
	db *sql.DB
}

func NewDatabase(connectionString string) (*Database, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return &Database{db: db}, nil
}

func (db *Database) Close() error {
	return db.db.Close()
}
