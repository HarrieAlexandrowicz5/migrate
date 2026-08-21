package pgx5

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func WithInstance(db *pgxpool.Pool, config *Config) (database.Driver, error) {
	// ... existing implementation ...
	d := &Pgx5{
		db:     db,
		config: config,
	}

	if err := d.ensureVersion(); err != nil {
		d.Close() // Clean up on failure
		return nil, err
	}

	return d, nil
}

func (p *Pgx5) Close() error {
	// ... existing close implementation ...
	return nil
}

func (p *Pgx5) ensureVersion() error {
	// ... existing logic ...
	return nil
}