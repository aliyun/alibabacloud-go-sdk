// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterTableSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSchema(v *Schema) *AlterTableSchemaRequest
	GetSchema() *Schema
}

type AlterTableSchemaRequest struct {
	// The table schema.
	Schema *Schema `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s AlterTableSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterTableSchemaRequest) GoString() string {
	return s.String()
}

func (s *AlterTableSchemaRequest) GetSchema() *Schema {
	return s.Schema
}

func (s *AlterTableSchemaRequest) SetSchema(v *Schema) *AlterTableSchemaRequest {
	s.Schema = v
	return s
}

func (s *AlterTableSchemaRequest) Validate() error {
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	return nil
}
