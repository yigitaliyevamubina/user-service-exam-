package storage

import (
	"exam/user-service/pkg/db"
	"exam/user-service/pkg/logger"

	// "exam/user-service/storage/postgres"
	"exam/user-service/storage/postgres"
	"exam/user-service/storage/repo"
)

// Storage
type StorageI interface {
	UserService() repo.UserServiceI
}

type storagePg struct {
	userService repo.UserServiceI
}

func New(db *db.Postgres, log logger.Logger) StorageI {
	return &storagePg{userService: postgres.NewUserRepo(db, log)}
	// return &storagePg{userService: mongodb.NewUserRepo(collection, log)}
}

func (s *storagePg) UserService() repo.UserServiceI {
	return s.userService
}
