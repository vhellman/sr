package sr

import (
	"encoding/json"
	"log/slog"

	"github.com/riferrei/srclient"
)

type sr struct {
	client *srclient.SchemaRegistryClient
}

type AvroSchema struct {
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Namespace string  `json:"namespace"`
	Fields    []Field `json:"fields"`
}

type Field struct {
	Name string      `json:"name"`
	Type interface{} `json:"type"`
}

// NewSchemaRegistryClient creates a new schema registry client
func New(url string) (*sr, error) {
	client := srclient.CreateSchemaRegistryClient(url)
	return &sr{client: client}, nil
}

func (s *sr) GetSchema(subject string) (*srclient.Schema, error) {
	schema, err := s.client.GetLatestSchema(subject)
	if err != nil {
		slog.Error("Error fetching schema", "subject", subject, "error", err)
		return nil, err
	}

	return schema, nil
}

func (s *sr) GetSchemaByID(ID int) (*srclient.Schema, error) {
	schema, err := s.client.GetSchema(ID)
	if err != nil {
		slog.Error("Error fetching schema", "id", ID, "error", err)
		return nil, err
	}

	return schema, nil
}

func (s *sr) GetSchemaBySubjectAndVersion(subject string, version int) (*srclient.Schema, error) {
	schema, err := s.client.GetSchemaByVersion(subject, version)
	if err != nil {
		slog.Error("Error fetching schema", "subject", subject, "version", version, "error", err)
		return nil, err
	}

	return schema, nil
}

// MergeSchemas takes a parent schema and a map of child schemas and returns a combined schema
func MergeSchemas(parentSchema string, childSchemas map[string]string) (string, error) {
	var parent AvroSchema
	err := json.Unmarshal([]byte(parentSchema), &parent)
	if err != nil {
		return "", err
	}

	// Iterate over the parent fields and replace references with actual child schemas
	for i, field := range parent.Fields {
		if childSchema, ok := childSchemas[field.Type.(string)]; ok {
			var child AvroSchema
			err := json.Unmarshal([]byte(childSchema), &child)
			if err != nil {
				return "", err
			}
			parent.Fields[i].Type = child
		}
	}

	result, err := json.MarshalIndent(parent, "", "  ")
	if err != nil {
		return "", err
	}

	return string(result), nil
}
