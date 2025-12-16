// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchemaTablesValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SchemaTablesValue
	GetName() *string
	SetPrimaryTable(v bool) *SchemaTablesValue
	GetPrimaryTable() *bool
	SetFields(v map[string]*SchemaTablesValueFieldsValue) *SchemaTablesValue
	GetFields() map[string]*SchemaTablesValueFieldsValue
}

type SchemaTablesValue struct {
	Name         *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	PrimaryTable *bool                                    `json:"primaryTable,omitempty" xml:"primaryTable,omitempty"`
	Fields       map[string]*SchemaTablesValueFieldsValue `json:"fields,omitempty" xml:"fields,omitempty"`
}

func (s SchemaTablesValue) String() string {
	return dara.Prettify(s)
}

func (s SchemaTablesValue) GoString() string {
	return s.String()
}

func (s *SchemaTablesValue) GetName() *string {
	return s.Name
}

func (s *SchemaTablesValue) GetPrimaryTable() *bool {
	return s.PrimaryTable
}

func (s *SchemaTablesValue) GetFields() map[string]*SchemaTablesValueFieldsValue {
	return s.Fields
}

func (s *SchemaTablesValue) SetName(v string) *SchemaTablesValue {
	s.Name = &v
	return s
}

func (s *SchemaTablesValue) SetPrimaryTable(v bool) *SchemaTablesValue {
	s.PrimaryTable = &v
	return s
}

func (s *SchemaTablesValue) SetFields(v map[string]*SchemaTablesValueFieldsValue) *SchemaTablesValue {
	s.Fields = v
	return s
}

func (s *SchemaTablesValue) Validate() error {
	return dara.Validate(s)
}
