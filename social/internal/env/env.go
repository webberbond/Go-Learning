package env

import (
	"os"
	"strconv"
	"time"
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

func ParseDuration(key string, fallback time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	duration, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}
	return duration
}
