package settings

import (
	"embed"
	"fmt"
	"strings"
)

func isHiddenConfig(fileName string) bool {
	return strings.HasPrefix(fileName, ".")
}

func skip(fileName string) bool {
	return isHiddenConfig(fileName)
}

func GetEmbedConfigs(fs embed.FS, dirName string) (map[string][]byte, error) {
	dirEntries, err := fs.ReadDir(dirName)
	if err != nil {
		return nil, fmt.Errorf("ReadDir: %w", err)
	}

	configs := make(map[string][]byte, len(dirEntries))

	for _, ent := range dirEntries {
		fileName := ent.Name()

		if !skip(fileName) {
			val, err := fs.ReadFile(dirName + "/" + fileName)
			if err != nil {
				return nil, fmt.Errorf("read file %s failed: %w", fileName, err)
			}

			configs[fileName] = val
		}
	}

	return configs, nil
}
