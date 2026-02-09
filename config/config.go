package config

import (
	"github.com/nico0302/go-print-server/internal/auth"
	"github.com/nico0302/go-print-server/internal/printer"
	"github.com/spf13/viper"
)

type PrinterConfig struct {
	printer.PrinterConfig     `mapstructure:",squash"`
	printer.CupsPrinterConfig `mapstructure:",squash"`
}

type (
	Config struct {
		Printers map[string]PrinterConfig      `mapstructure:"printers"`
		Cups     map[string]printer.CupsConfig `mapstructure:"cups"`
		Presets  map[string]Preset             `mapstructure:"presets"`
		HTTP     HttpConfig                    `mapstructure:"http"`
		Log      LogConfig                     `mapstructure:"logger"`
	}

	Preset struct {
		Printer       string
		TypstTemplate TypstTemplateConfig   `mapstructure:"typst"`
		JobAttributes printer.JobAttributes `mapstructure:"attributes"`
	}

	TypstTemplateConfig struct {
		WorkingDirectory string `mapstructure:"working-directory"`
		Content          string `mapstructure:"content"`
	}

	HttpConfig struct {
		Port  string               `mapstructure:"port"`
		Users map[string]auth.User `mapstructure:"users"`
	}

	LogConfig struct {
		Level string
	}
)

func NewConfig() (*Config, error) {
	var cfg Config

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AddConfigPath("../../config")
	v.AddConfigPath("./config")
	v.AddConfigPath("etc/cloudprint")
	v.ReadInConfig()

	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
