package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID        uuid.UUID
	Title     string
	Started   time.Time
	Ended     time.Time
	ServiceID int
}
