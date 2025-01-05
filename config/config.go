package config

import (
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type Server struct {
	Port int
}
type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	TimeZone string
}
type Config struct {
	Server *Server
	Db     *Db
}

var once sync.Once
var configInstance *Config

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})
	return configInstance
}
