package gofig

import (
	"encoding/json"
	"fmt"
	"os"
)

func NewConfig[Config any](path string, jsonData Config) error {
	root := os.Getenv("HOME") + "/.config"

	configDir := root + path

	if mkDirErr := createDirectories(configDir); mkDirErr != nil {
		return mkDirErr
	}

	file, createErr := os.Create(configDir + "/config.json")

	if createErr != nil {
		fmt.Printf("Failed to create config at %s \n", configDir+"/config.json")
		return createErr
	}

	data, jsonErr := json.Marshal(jsonData)

	if jsonErr != nil {
		return jsonErr
	}

	if _, writeErr := file.Write(data); writeErr != nil {
		fmt.Println("Failed to write to config.")
		return writeErr
	}
	return nil
}

func GetConfig[Config any](path string, c *Config) error {

	root := os.Getenv("HOME") + "/.config"

	configDir := root + path + "/config.json"

	data, err := os.ReadFile(configDir)

	if err != nil {
		fmt.Printf("Could not read file at %s \n", configDir)
		return err
	}

	if jsonError := json.Unmarshal(data, c); jsonError != nil {
		fmt.Println("Failed to Parse JSON")
		return jsonError
	}
	return nil
}

func createDirectories(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println("Failed to create directories!")
		return err
	}
	return nil
}
