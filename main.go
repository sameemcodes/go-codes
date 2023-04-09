package main

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"go-codes/config"
	"go-codes/durable"
	"go-codes/controllers"

)
var (
	env        = flag.String("env", "local", "")
	configFile = flag.String("configFile", "internal/config/local.toml", "")
)
func main(){
	flag.Parse()
	if _, err := toml.DecodeFile(*configFile, &config.Config); err != nil {
		panic(err)
	}
	logger := initLogger(env)
	r := mux.NewRouter()
	controllers.RegisterRoutes(r, logger)

}

func initLogger(env *string) *durable.Logger {
	zapConfig := zap.NewDevelopmentConfig()
	if *env == "prod" || *env == "ingestion" {
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.OutputPaths = config.Config.LoggerConfig.OutputPaths
	zapConfig.ErrorOutputPaths = config.Config.LoggerConfig.ErrorOutputPaths

	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	return durable.NewLogger(logger)
}
