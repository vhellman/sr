package sr

// Schemas holds the JSON strings for parent and child schemas.
type Schemas struct {
	ParentSchema   string
	ChildSchemas   map[string]string
	ExpectedSchema string
}

/**
*
*				TEST SCHEMA
*
 */
var Employee = Schemas{
	ParentSchema: `{
		"type": "record",
		"name": "Employee",
		"namespace": "com.company.hr",
		"fields": [
			{"name": "personalInfo", "type": "PersonalInfo"},
			{"name": "department", "type": "string"},
			{"name": "role", "type": "string"},
			{"name": "yearsOfExperience", "type": "int"}
		]
	}`,

	ChildSchemas: map[string]string{
		"PersonalInfo": `{
			"type": "record",
			"name": "PersonalInfo",
			"namespace": "com.company.hr",
			"fields": [
				{"name": "name", "type": "string"},
				{"name": "age", "type": "int"},
				{"name": "address", "type": "string"}
			]
		}`,
	},
	ExpectedSchema: `{
		"type": "record",
		"name": "Employee",
		"namespace": "com.company.hr",
		"fields": [
			{
				"name": "personalInfo", 
				"type": {
					"type": "record",
					"name": "PersonalInfo",
					"namespace": "com.company.hr",
					"fields": [
						{"name": "name", "type": "string"},
						{"name": "age", "type": "int"},
						{"name": "address", "type": "string"}
					]
				}
			},
			{"name": "department", "type": "string"},
			{"name": "role", "type": "string"},
			{"name": "yearsOfExperience", "type": "int"}
		]
	}
	`,
}

/**
*
*				ORIGINAL TEST SCHEMA
*
 */
var User = Schemas{
	ParentSchema: `{
		"type": "record",
		"name": "User",
		"namespace": "com.example",
		"fields": [
			{"name": "headers", "type": "Headers"},
			{"name": "things", "type": "Things"},
			{"name": "name", "type": "string"},
			{"name": "age",  "type": "int"}
		]
	}`,

	ChildSchemas: map[string]string{
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
	},

	ExpectedSchema: `{
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
	}`,
}
