package settings

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"

	"github.com/releaseband/go-settings/reader"
	"gopkg.in/yaml.v3"
)

var ErrNotImplementedForThisType = errors.New("not implemented for type")

func setYamlSettings(cfg interface{}) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("yaml.Marshal: %w", err)
	}

	if err := reader.ReadSettingFromByte(data, "yaml"); err != nil {
		return fmt.Errorf("ReadSettingFromByte: %w", err)
	}

	return nil
}

func SetInitialData(cfg interface{}, configType string) error {
	if configType != "yaml" {
		return ErrNotImplementedForThisType
	}

	return setYamlSettings(cfg)
}

func unmarshal(cfg interface{}) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("viper.Unmarshal: %w", err)
	}

	return nil
}

func SetSettingsFromFile(fileName, path string, cfg interface{}) error {
	if err := reader.ReadSettingFromFile(fileName, path); err != nil {
		return err
	}

	return unmarshal(cfg)
}

func SetSettings(data []byte, cfg interface{}, configType string) error {
	if err := reader.ReadSettingFromByte(data, configType); err != nil {
		return err
	}

	return unmarshal(cfg)
}
