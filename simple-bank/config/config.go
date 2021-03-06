package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresHost        string        `mapstructure:"POSTGRES_HOST"`
	PostgresUser        string        `mapstructure:"POSTGRES_USER"`
	PostgresPassword    string        `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB          string        `mapstructure:"POSTGRES_DB"`
	PostgresPort        string        `mapstructure:"POSTGRES_PORT"`
	PostgresSSLMode     string        `mapstructure:"POSTGRES_SSL_MODE"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	// This is how you would specify the name of the config file, ex app.env
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// If the env variables exist in the env, they will overwrite the file values
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
