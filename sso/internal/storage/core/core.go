package core

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	"grpc_auth_tutorial/sso/internal/domain/models"
	"grpc_auth_tutorial/sso/internal/storage"
)

// TODO change mutex to atomic
type CoreStore struct {
	mu    sync.RWMutex
	store map[uuid.UUID]models.StoreEvent
}

func NewCoreStore() *CoreStore {
	return &CoreStore{
		store: make(map[uuid.UUID]models.StoreEvent, 0),
	}
}

func (c *CoreStore) Get(key uuid.UUID) (models.StoreEvent, error) {
	const op = "storage.core.Get"

	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.store[key]

	if !ok {
		return models.StoreEvent{}, fmt.Errorf("%s: %w", op, storage.ErrorNoSuchKey)
	}

	return value, nil
}

func (c *CoreStore) Put(key uuid.UUID, value models.StoreEvent) error {
	// const op = "storage.core.Put"

	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = value

	return nil
}

func (c *CoreStore) Delete(key uuid.UUID) error {
	// const op = "storage.core.Delete"

	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.store, key)

	return nil
}
