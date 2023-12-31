package list

import (
	"time"
)

type Entity struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	ID        string    `db:"id"`
	Title     *string   `db:"title"`
	ActiveAt  *string   `db:"active_at"`
	Status    *bool     `db:"status"`
}
