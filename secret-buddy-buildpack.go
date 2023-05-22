package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type ConsolidatedSecret struct {
	Current  map[string]string `json:"current"`  // Current secret values
	Previous map[string]string `json:"previous"` // Previous secret values, useful for rollbacks
}

func ExportEnvVarsFromMap(env string, rules map[string]string) (map[string]string, error) {
	var consolidatedSecret ConsolidatedSecret
	err := json.Unmarshal([]byte(env), &consolidatedSecret)
	if err != nil {
		return nil, err
	}
	var envVars map[string]string
	envVars = consolidatedSecret.Current
	fmt.Println(envVars)

	for key, value := range rules {

		if strings.Contains(value, "[") {
			var stringValue = ""
			value = strings.Replace(value, "[", "", -1)
			value = strings.Replace(value, "]", "", -1)
			values := strings.Split(value, ",")
			for _, v := range values {
				valueAndKey := strings.Split(v, ".")
				if valueAndKey[0] == "current" {
					stringValue = stringValue + consolidatedSecret.Current[valueAndKey[1]] + ","
				} else if valueAndKey[0] == "previous" {
					stringValue = stringValue + consolidatedSecret.Previous[valueAndKey[1]] + ","
				}
			}

		} else if strings.Contains(value, ",") {
			values := strings.Split(value, ",")
			var stringValue = ""
			for _, v := range values {
				valueAndKey := strings.Split(v, ".")
				if valueAndKey[0] == "current" {
					stringValue = stringValue + consolidatedSecret.Current[valueAndKey[1]] + ","
				} else if valueAndKey[0] == "previous" {
					stringValue = stringValue + consolidatedSecret.Previous[valueAndKey[1]] + ","
				}
			}
			stringValue = strings.TrimSuffix(stringValue, ",")
			envVars[key] = stringValue
		} else {
			valueAndKey := strings.Split(value, ".")
			if valueAndKey[0] == "current" {
				envVars[key] = consolidatedSecret.Current[valueAndKey[1]]
			} else if valueAndKey[0] == "previous" {
				envVars[key] = consolidatedSecret.Previous[valueAndKey[1]]
			}

		}
	}

	return envVars, nil
}

func GetEnvVar(envVarName string) (string, error) {
	envVarValue := os.Getenv(envVarName)
	return envVarValue, nil
}

func ParseRules(ruleString string) (map[string]string, error) {
	var rules map[string]string
	err := json.Unmarshal([]byte(ruleString), &rules)
	if err != nil {
		return nil, err
	}
	return rules, err
}

func main() {

	envVar, err := GetEnvVar("SECRETBUDDY_ENV")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	variableX, err := GetEnvVar("HEROKU_SECRETS_CONFIG")
	if err != nil {
		fmt.Println(err)
	}

	rules, err := ParseRules(variableX)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	consolidatedSecret, err := ExportEnvVarsFromMap(envVar, rules)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for key, value := range consolidatedSecret {
		fmt.Printf("export %s=%v\n", key, value)
	}

}
