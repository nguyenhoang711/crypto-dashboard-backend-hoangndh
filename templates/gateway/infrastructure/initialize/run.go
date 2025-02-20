package initialize

import (
	"crypto-dashboard/gw-example/infrastructure/global"
	"crypto-dashboard/pkg/appLogger"
	"crypto-dashboard/pkg/common"
	database "crypto-dashboard/pkg/database/postgres"
	"crypto-dashboard/pkg/response"
)

func Run() {
	global.AppConfig = common.LoadConfig[global.Config]()
	var err *response.AppError
	global.AppConfig = common.LoadConfig[global.Config]()
	global.Log, err = appLogger.NewLogger(global.AppConfig.Logger, global.AppConfig.Server)
	if err != nil {
		panic(err)
	}

	global.SQLDB, err = database.NewConnection(global.AppConfig.SQL)
}
