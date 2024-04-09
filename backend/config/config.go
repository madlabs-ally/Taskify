package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/spf13/viper"
)

var (
    AppConfig *Config
    DB        *gorm.DB
)

type Config struct {
    Port     string
    DBConfig DBConfig
}

type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Name     string
}

func LoadConfig() error {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        return err
    }

    AppConfig = &Config{}
    err = viper.Unmarshal(AppConfig)
    if err != nil {
        return err
    }

    return nil
}

func InitDB() error {
    dbURI := "host=" + AppConfig.DBConfig.Host +
        " port=" + AppConfig.DBConfig.Port +
        " user=" + AppConfig.DBConfig.User +
        " dbname=" + AppConfig.DBConfig.Name +
        " password=" + AppConfig.DBConfig.Password +
        " sslmode=disable"

    db, err := gorm.Open("postgres", dbURI)
    if err != nil {
        return err
    }

    DB = db
    return nil
}
