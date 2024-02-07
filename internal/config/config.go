package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

//const (
//	envPrefix   = ""
//	envFilePath = "configs/.env"
//)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	DB         `yaml:"db"`
	HTTPServer `yaml:"http_server"`
	Api        `yaml:"api"`
}

type Api struct {
	AgeUrl         string `yaml:"age_url"`
	GenderUrl      string `yaml:"gender_url"`
	NationalityUrl string `yaml:"nationality_url"`
}

type DB struct {
	Username           string `yaml:"username" env-default:"mobile"`
	Host               string `yaml:"host" env-default:"localhost"`
	Port               string `yaml:"port" env-default:"5040"`
	DBName             string `yaml:"db_name" env-default:"profilesdb"`
	Password           string `yaml:"password" env-default:"password"`
	SSLMode            string `yaml:"sslmode" env-default:"disable"`
	DatabaseURL        string `yaml:"db_url" required:"true"`
	MaxOpenConnections int    `yaml:"db_max_open_connections" default:"10"`
}

type HTTPServer struct {
	Address     string        `yaml:"HTTP_SERVER_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"HTTP_SERVER_TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"HTTP_SERVER_CTX_TIMEOUT" env-default:"60s"`
	CtxTimeout  time.Duration `yaml:"HTTP_SERVER_IDLE_TIMEOUT" default:"60s"`
}

//var CONFIG_PATH = "./config/local.yaml"

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
