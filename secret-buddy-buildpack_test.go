package main

import (
	"os"
	"testing"
)

// # This will set the KEY1 env var to a comma-delimited list of current and previous values:
// heroku config:set HEROKU_SECRETS_CONFIG='{"KEY1": "current.KEY1,previous.KEY1"}'
func TestExportEnvVarsFromMap(t *testing.T) {
	// Define a test case
	testCase := struct {
		inputEnv       string
		inputRules     map[string]string
		expectedOutput map[string]string
	}{
		inputEnv: `{"current": {"DB_PASSWORD": "current_password"}, "previous": {"DB_PASSWORD": "previous_password"}}`,
		inputRules: map[string]string{
			"DB_PASSWORD": "current.DB_PASSWORD,previous.DB_PASSWORD",
		},
		expectedOutput: map[string]string{
			"DB_PASSWORD": "current_password,previous_password",
		},
	}

	// Call the ExportEnvVarsFromMap function with the test case input
	outputValue, err := ExportEnvVarsFromMap(testCase.inputEnv, testCase.inputRules)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

// # This will use the previous version of the KEY1 secret for the KEY1 env var:
// heroku config:set HEROKU_SECRETS_CONFIG='{"KEY1": "previous.KEY1"}'
func TestExportPreviousEnvVarsFromMap(t *testing.T) {
	// Define a test case
	testCase := struct {
		inputEnv       string
		inputRules     map[string]string
		expectedOutput map[string]string
	}{
		inputEnv: `{"current": {"DB_PASSWORD": "current_password"}, "previous": {"DB_PASSWORD": "previous_password"}}`,
		inputRules: map[string]string{
			"DB_PASSWORD": "previous.DB_PASSWORD",
		},
		expectedOutput: map[string]string{
			"DB_PASSWORD": "previous_password",
		},
	}

	// Call the ExportEnvVarsFromMap function with the test case input
	outputValue, err := ExportEnvVarsFromMap(testCase.inputEnv, testCase.inputRules)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

// # This will set KEY1 to the current value, and another secret to the previous value.
// # (The default still applies to KEY1, although you could be explicit if you want):
func TestExportAlternatePreviousEnvVarsFromMap(t *testing.T) {
	// Define a test case
	testCase := struct {
		inputEnv       string
		inputRules     map[string]string
		expectedOutput map[string]string
	}{
		inputEnv: `{"current": {"DB_PASSWORD": "current_password"}, "previous": {"DB_PASSWORD": "previous_password"}}`,
		inputRules: map[string]string{
			"OTHER_DB_PASSWORD": "previous.DB_PASSWORD",
		},
		expectedOutput: map[string]string{
			"OTHER_DB_PASSWORD": "previous_password",
		},
	}

	// Call the ExportEnvVarsFromMap function with the test case input
	outputValue, err := ExportEnvVarsFromMap(testCase.inputEnv, testCase.inputRules)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

// # This will set KEY1 to a JSON-encoded list including current and previous values:
// heroku config:set HEROKU_SECRETS_CONFIG='{"KEY1": "[\"current.KEY1\",\"previous.KEY1\"]"}
func TestExportJSONEnvVarsFromMap(t *testing.T) {
	// Define a test case
	testCase := struct {
		inputEnv       string
		inputRules     map[string]string
		expectedOutput map[string]string
	}{
		inputEnv: `{"current": {"DB_PASSWORD": "current_password"}, "previous": {"DB_PASSWORD": "previous_password"}}`,
		inputRules: map[string]string{
			"DB_PASSWORD": "[\"current.DB_PASSWORD\", \"previous.DB_PASSWORD\"]",
		},
		expectedOutput: map[string]string{
			"DB_PASSWORD": "current_password",
		},
	}

	// Call the ExportEnvVarsFromMap function with the test case input
	outputValue, err := ExportEnvVarsFromMap(testCase.inputEnv, testCase.inputRules)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestExportEnvVarsFromMapEmptyRules(t *testing.T) {
	// Define a test case
	testCase := struct {
		inputEnv       string
		inputRules     map[string]string
		expectedOutput map[string]string
	}{
		inputEnv: `{"current": {"DB_PASSWORD": "current_password"}, "previous": {"DB_PASSWORD": "previous_password"}}`,
		inputRules: map[string]string{
			"": "",
		},
		expectedOutput: map[string]string{
			"DB_PASSWORD": "current_password",
		},
	}

	// Call the ExportEnvVarsFromMap function with the test case input
	outputValue, err := ExportEnvVarsFromMap(testCase.inputEnv, testCase.inputRules)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestExportEnvVarsFromMapEmptyRulesWithTickChar(t *testing.T) {
	// Define a test case
	testCase := struct {
		inputEnv       string
		inputRules     map[string]string
		expectedOutput map[string]string
	}{
		inputEnv: `{"current": {"DB_PASSWORD": "current'password'"}, "previous": {"DB_PASSWORD": "previous_password"}}`,
		inputRules: map[string]string{
			"": "",
		},
		expectedOutput: map[string]string{
			"DB_PASSWORD": "current'password'",
		},
	}

	// Call the ExportEnvVarsFromMap function with the test case input
	outputValue, err := ExportEnvVarsFromMap(testCase.inputEnv, testCase.inputRules)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestParseRules(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput map[string]string
	}{
		input: `{"DB_PASSWORD": "current.DB_PASSWORD"}`,
		expectedOutput: map[string]string{
			"DB_PASSWORD": "current.DB_PASSWORD",
		},
	}

	// Call the ParseRules function with the test case input
	outputValue, err := ParseRules(testCase.input)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestParseEmptyRules(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput map[string]string
	}{
		input: "{}",
		expectedOutput: map[string]string{
			"": "",
		},
	}

	// Call the ParseRules function with the test case input
	outputValue, err := ParseRules(testCase.input)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestGetEnvVar(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "SECRETBUDDY_ENV",
		expectedOutput: "test_value",
	}

	// Set the environment variable for the test case
	err := os.Setenv(testCase.input, testCase.expectedOutput)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Call the GetEnvVar function with the test case input
	outputValue, exists := os.LookupEnv(testCase.input)

	// Check if there was an error
	if exists {
		// Check if the output matches the expected output
		if outputValue != testCase.expectedOutput {
			t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
		}
	} else {
		t.Errorf("Expected that exists is true but got %v", exists)
	}
}

func TestNotExistsGetEnvVar(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "SECRETBUDDY_ENV",
		expectedOutput: "test_value",
	}

	// Call the GetEnvVar function with the test case input
	_, exists := os.LookupEnv(testCase.input)

	// Check if there was an error
	if exists {
		t.Errorf("EnvVar should not exist")
	}
}

func TestGetEmptyEnvVar(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "SECRETBUDDY_ENV",
		expectedOutput: "",
	}

	// Set the environment variable for the test case
	err := os.Setenv(testCase.input, testCase.expectedOutput)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Call the GetEnvVar function with the test case input
	outputValue, exists := os.LookupEnv(testCase.input)

	// Check if there was an error
	if exists {
		// Check if the output matches the expected output
		if outputValue != testCase.expectedOutput {
			t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
		}
	} else {
		t.Errorf("Expected that exists is true but got %v", exists)
	}
}

func TestEscapedStringsSecret(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "SECRETBUDDY_SECRET",
		expectedOutput: "SECRETBUDDY_SECRET",
	}

	outputValue := EscapeSingleQuote(testCase.input)
	// Check if the output matches the expected output
	if outputValue != testCase.expectedOutput {
		t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
	}
}

func TestEscapedStringsQuoteSecret(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "SECRETBUDDY'SECRET",
		expectedOutput: "SECRETBUDDY'\\''SECRET",
	}

	outputValue := EscapeSingleQuote(testCase.input)
	// Check if the output matches the expected output
	if outputValue != testCase.expectedOutput {
		t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
	}
}

func TestEscapedStringsTWOQuoteSecret(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "SECRET'BUDDY'SECRET",
		expectedOutput: "SECRET'\\''BUDDY'\\''SECRET",
	}

	outputValue := EscapeSingleQuote(testCase.input)
	// Check if the output matches the expected output
	if outputValue != testCase.expectedOutput {
		t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
	}
}

func TestEscapedStringsEmptySecret(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "",
		expectedOutput: "",
	}

	outputValue := EscapeSingleQuote(testCase.input)
	// Check if the output matches the expected output
	if outputValue != testCase.expectedOutput {
		t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
	}
}

func TestEscapedStringsSingleTick(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "'",
		expectedOutput: "'\\''",
	}

	outputValue := EscapeSingleQuote(testCase.input)
	// Check if the output matches the expected output
	if outputValue != testCase.expectedOutput {
		t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
	}
}

func TestEscapedStringsSecretDoubleQuote(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "SECRET\"BUDDY\"_SECRET",
		expectedOutput: "SECRET\"BUDDY\"_SECRET",
	}

	outputValue := EscapeSingleQuote(testCase.input)
	// Check if the output matches the expected output
	if outputValue != testCase.expectedOutput {
		t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
	}
}
