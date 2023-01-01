/*
//
//  ______    _   _           _  __      _   _     ____   ___
// |  ____|  | | | |         | |/ _|    | | | |   |___ \ / _ \
// | |__ __ _| |_| |__   __ _| | |_ __ _| |_| |__   __) | | | |
// |  __/ _` | __| '_ \ / _` | |  _/ _` | __| '_ \ |__ <| | | |
// | | | (_| | |_| | | | (_| | | || (_| | |_| | | |___) | |_| |
// |_|  \__,_|\__|_| |_|\__,_|_|_| \__,_|\__|_| |_|____/ \___/
//
// Written by Fathalfath30.
// Email : fathalfath30@gmail.com
// Follow me on:
//  Github : https://github.com/fathalfath30
//  Gitlab : https://gitlab.com/Fathalfath30
//
*/

package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"scum-helper/services/helper"
	"strings"
)

type App struct {
	host string
	port string

	logger *zap.Logger
}

// GetHost will return assigned host that has been set on environment variable
// as default if it not set any value it will return "127.0.0.1" as default
// host
func (app App) GetHost() string {
	return app.host
}

// GetPort will return assigned port that has been set on environment variable
// if port is not set it will throw panic
func (app App) GetPort() string {
	return app.port
}

func Logger() *zap.Logger {
	lc := zap.NewDevelopmentConfig()
	lc.DisableStacktrace = false
	lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := lc.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

// NewApp will return application configuration and it may throw panic
func NewApp() App {
	logger := Logger()
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			Logger().Error("error while loading .env file", zap.Error(err))
		}
	} else {
		Logger().Warn("running service without configuration from .env")
	}

	host := os.Getenv(helper.AppHostEnvName)
	if host == "" {
		host = helper.ApplicationDefaultHost
	}

	port := os.Getenv(helper.AppPortEnvName)
	if port == "" {
		logger.Error("error while loading port: port is required")
		panic("port is required")
	}

	return App{
		host: strings.TrimSpace(host),
		port: strings.TrimSpace(port),

		logger: logger,
	}
}
