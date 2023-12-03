package remoteTest

import (
	"log"
)

func LogInfo(message string) {
	log.Printf("Information: %v \n", message)
}

func LogError(message string) {
	log.Printf("Error: %v\n", message)
}

func LogWarning(message string) {
	log.Printf("Warning: %v\n", message)
}

func LogSuccess(message string) {
	log.Printf("Success: %s\n ", message)
}
