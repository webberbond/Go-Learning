package env

import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	strVal, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return strVal
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return intVal
}
