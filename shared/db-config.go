package shared

import (
	"context"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
)

type (
	DBConfig struct {
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		Name     string `env:"DB_NAME"`
	}
)

func ConnectDB() (*pgx.Conn, error) {
	var cfg DBConfig

	// Load configuration from environment variables
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading environment variables: %v\n", err)
		return nil, err
	}

	// Construct the connection string
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}
