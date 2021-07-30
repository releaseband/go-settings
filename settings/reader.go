package settings

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var ErrFileNotFound = errors.New("file not found")

func ReadSettings(data []byte, configType string) error {
	viper.SetConfigType(configType)

	if err := viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return fmt.Errorf("viper.ReadConfig: %w", err)
	}

	return nil
}

func ReadSettingsFromFile(fileName, path string) error {
	if len(strings.TrimSpace(fileName)) != 0 {
		viper.SetConfigFile(fileName)
	}

	viper.SetConfigName(fileName)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		var errFileNotFound viper.ConfigFileNotFoundError

		if errors.As(err, &errFileNotFound) {
			return ErrFileNotFound
		}

		return fmt.Errorf("viper.ReadConfig(`%s/%s`): %w", path, fileName, err)
	}

	return nil
}

func ReadFile(filePath string) ([]byte, error) {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}

		return nil, fmt.Errorf("os.Stat: %w", err)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile: %w", err)
	}

	return data, nil
}
