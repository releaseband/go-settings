package settings

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"

	"github.com/spf13/viper"

	"gopkg.in/yaml.v3"
)

var ErrNotImplementedForThisType = errors.New("not implemented for type")

func SaveYamlSettings(cfg interface{}) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("yaml.Marshal: %w", err)
	}

	if err := ReadSettings(data, "yaml"); err != nil {
		return fmt.Errorf("ReadSettings: %w", err)
	}

	return nil
}

func SaveDefaultSettings(cfg interface{}, configType string) error {
	if configType != "yaml" {
		return ErrNotImplementedForThisType
	}

	return SaveYamlSettings(cfg)
}

func UnmarshalSettings(cfg interface{}) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("viper.UnmarshalSettings: %w", err)
	}

	return nil
}

func Unmarshal(data []byte, configType string, cfg interface{}) error {
	if err := ReadSettings(data, configType); err != nil {
		return err
	}

	return UnmarshalSettings(cfg)
}

func UnmarshalFile(fileName, path string, cfg interface{}) error {
	if err := ReadSettingsFromFile(fileName, path); err != nil {
		return err
	}

	return UnmarshalSettings(cfg)
}

func UnmarshalEnv(envPrefix string, cfg interface{}) error {
	return envconfig.Process(envPrefix, cfg)
}

func UnmarshalEmbed(fs embed.FS, embedDirName string, prefix string, cfg interface{}) error {
	secrets, err := GetEmbedConfigs(fs, embedDirName)
	if err != nil {
		return err
	}

	viper.SetEnvPrefix(prefix)

	for k, v := range secrets {
		if err := viper.BindEnv(k); err != nil {
			return fmt.Errorf("bind env %s failed: %w", k, err)
		}

		if err := os.Setenv(strings.ToUpper(prefix+"_"+k), string(v)); err != nil {
			return fmt.Errorf("os.Setenv %s failed: %w", k, err)
		}
	}

	return UnmarshalEnv(prefix, cfg)
}
