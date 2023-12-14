package sr

import (
	"encoding/json"
	"strings"
	"testing"
)

// Test case structure
type testCase struct {
	name                   string
	parentSchema           string
	childSchemas           map[string]string
	expectedCombinedSchema string
}

// Test cases
func Test_CombineSchemas(t *testing.T) {
	// Define your test cases
	tests := []testCase{
		{
			name:                   "Test Employee Schema",
			parentSchema:           Employee.ParentSchema,
			childSchemas:           Employee.ChildSchemas,
			expectedCombinedSchema: Employee.ExpectedSchema,
		},
		{
			name:                   "Test User Schema",
			parentSchema:           User.ParentSchema,
			childSchemas:           User.ChildSchemas,
			expectedCombinedSchema: User.ExpectedSchema,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			combinedSchema, err := MergeSchemas(tc.parentSchema, tc.childSchemas)
			if err != nil {
				t.Errorf("CombineSchemas() returned an error: %v", err)
			}

			// Normalize both strings by removing whitespaces and line breaks
			normalizedExpected := normalizeString(tc.expectedCombinedSchema)
			normalizedActual := normalizeString(combinedSchema)

			// Compare the normalized strings
			if normalizedExpected != normalizedActual {
				t.Errorf("CombineSchemas() returned incorrect combined schema for %s.\nExpected: %s\nGot: %s", tc.name, normalizedExpected, normalizedActual)
			}
		})
	}
}

// normalizeString removes all whitespaces, line breaks, and indentation from a string.
func normalizeString(str string) string {
	return strings.Join(strings.Fields(str), "")
}

// TestCombineSchemas compares JSON objects instead of string representations.
func TestCombineSchemas(t *testing.T) {
	// Original test data
	parentSchema := `{
		"type": "record",
		"name": "User",
		"namespace": "com.example",
		"fields": [
			{"name": "headers", "type": "Headers"},
			{"name": "things", "type": "Things"},
			{"name": "name", "type": "string"},
			{"name": "age",  "type": "int"}
		]
	}`
	childSchemas := map[string]string{
		"Headers": `{
			"type": "record",
			"name": "Headers",
			"namespace": "com.example",
			"fields": [
				{"name": "headerName", "type": "string"},
				{"name": "headerAge",  "type": "int"}
			]
		}`,
		"Things": `{
			"type": "record",
			"name": "Things",
			"namespace": "com.example",
			"fields": [
				{"name": "stuff", "type": "string"}
			]
		}`,
	}
	expectedCombinedSchema := `{
		"type": "record",
		"name": "User",
		"namespace": "com.example",
		"fields": [
			{"name": "headers", 
			"type": {
				"type": "record",
				"name": "Headers",
				"namespace": "com.example",
				"fields": [
					{"name": "headerName", "type": "string"},
					{"name": "headerAge",  "type": "int"}
				]
			}},
			{"name": "things", "type": {
				"type": "record",
				"name": "Things",
				"namespace": "com.example",
				"fields": [
					{"name": "stuff", "type": "string"}
				]
			}},
			{"name": "name", "type": "string"},
			{"name": "age",  "type": "int"}
		]
	}`

	combinedSchema, err := MergeSchemas(parentSchema, childSchemas)
	if err != nil {
		t.Errorf("CombineSchemas() returned an error: %v", err)
	}

	// Unmarshal expected and actual schemas into interface{} for comparison
	var expected, actual interface{}
	if err := json.Unmarshal([]byte(expectedCombinedSchema), &expected); err != nil {
		t.Errorf("Failed to unmarshal expected schema: %v", err)
	}
	if err := json.Unmarshal([]byte(combinedSchema), &actual); err != nil {
		t.Errorf("Failed to unmarshal actual schema: %v", err)
	}

	// Compare the unmarshalled JSON objects
	if !jsonEqual(expected, actual) {
		t.Errorf("CombineSchemas() returned incorrect combined schema.\nExpected: %s\nGot: %s", expectedCombinedSchema, combinedSchema)
	}
}

// jsonEqual compares two JSON objects for structural and content equality.
func jsonEqual(a, b interface{}) bool {
	ajson, _ := json.Marshal(a)
	bjson, _ := json.Marshal(b)
	return string(ajson) == string(bjson)
}
