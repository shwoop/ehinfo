package main

import (
	"bytes"
	"os"
)

// Config is the general configuation struct
type Config struct {
	key, configDir, configFile, lockDir, lockFile string
	action                                        ActionFunction
}

// buildConfig validates the user input.
// The relevant filesystem paths are also generated at this time.
func buildConfig() (Config, error) {
	var config Config
	args := os.Args[1:]
	if len(args) == 0 {
		help(0)
	}
	var objType string
	var objUuid string
	var err error
	config = Config{}
	argLength := len(args)
	if argLength >= 3 {
		config.action = putJson
		objType, err = checkType(args[1])
		// checkError(err)
		if err != nil {
			return Config{}, err
		}
		objUuid = args[2]
	}
	if argLength == 3 {
		switch args[0] {
		case "json":
			config.action = putJson
		case "put":
			config.action = putText
		case "info":
			config.action = getAll
		default:
			help(1)
		}
	} else if argLength == 4 {
		if args[0] != "get" {
			help(1)
		}
		config.action = getSingleValue
		config.key = args[3]
	}

	// stateDir := os.Getenv("STATEDIR")
	// lockDir := os.Getenv("LOCKDIR")
	// testing
	stateDir := "/tmp/eh/state"
	lockDir := "/tmp/eh/lock"

	var filename bytes.Buffer
	filename.WriteString(stateDir)
	filename.WriteRune('/')
	filename.WriteString(objType)
	filename.WriteRune('/')
	filename.WriteString(objUuid)
	config.configDir = filename.String()

	filename.WriteString("/config.json")
	config.configFile = filename.String()

	filename.Reset()

	filename.WriteString(lockDir)
	filename.WriteRune('/')
	filename.WriteString(objType)
	filename.WriteRune('/')
	config.lockDir = filename.String()

	filename.WriteString(objUuid)
	config.lockFile = filename.String()
	return config, nil
}
