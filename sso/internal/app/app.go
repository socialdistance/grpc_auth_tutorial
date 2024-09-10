package app

import (
	"grpc_auth_tutorial/sso/internal/services/auth"
	"grpc_auth_tutorial/sso/internal/storage/sqlite"
	"log/slog"
	"time"

	grpcapp "grpc_auth_tutorial/sso/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
