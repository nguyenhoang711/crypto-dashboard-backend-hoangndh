package global

import (
	"crypto-dashboard/pkg/appLogger"
	"crypto-dashboard/pkg/database/caching"
	database "crypto-dashboard/pkg/database/postgres"
	"crypto-dashboard/pkg/settings"
)

type Config struct {
	AllowOrigins string                  `mapstructure:"allow_origins"`
	Server       *settings.ServerSetting `mapstructure:"server"`
	SQL          *settings.SQLSetting    `mapstructure:"sql"`
	Logger       *settings.LoggerSetting `mapstructure:"logger"`
	Cache        *settings.CacheSetting  `mapstructure:"redis"`
	JWT          *settings.JWTSetting    `mapstructure:"jwt"`
}

var (
	Log       *appLogger.Logger
	AppConfig *Config
	SQLDB     *database.Connection
	Cache     *caching.CacheClient
)
