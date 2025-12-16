// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchemaTablesValueFieldsValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SchemaTablesValueFieldsValue
	GetName() *string
	SetPrimaryKey(v bool) *SchemaTablesValueFieldsValue
	GetPrimaryKey() *bool
	SetType(v string) *SchemaTablesValueFieldsValue
	GetType() *string
	SetJoinWith(v []*string) *SchemaTablesValueFieldsValue
	GetJoinWith() []*string
	SetLabel(v string) *SchemaTablesValueFieldsValue
	GetLabel() *string
}

type SchemaTablesValueFieldsValue struct {
	Name       *string   `json:"name,omitempty" xml:"name,omitempty"`
	PrimaryKey *bool     `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	Type       *string   `json:"type,omitempty" xml:"type,omitempty"`
	JoinWith   []*string `json:"joinWith,omitempty" xml:"joinWith,omitempty" type:"Repeated"`
	Label      *string   `json:"label,omitempty" xml:"label,omitempty"`
}

func (s SchemaTablesValueFieldsValue) String() string {
	return dara.Prettify(s)
}

func (s SchemaTablesValueFieldsValue) GoString() string {
	return s.String()
}

func (s *SchemaTablesValueFieldsValue) GetName() *string {
	return s.Name
}

func (s *SchemaTablesValueFieldsValue) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *SchemaTablesValueFieldsValue) GetType() *string {
	return s.Type
}

func (s *SchemaTablesValueFieldsValue) GetJoinWith() []*string {
	return s.JoinWith
}

func (s *SchemaTablesValueFieldsValue) GetLabel() *string {
	return s.Label
}

func (s *SchemaTablesValueFieldsValue) SetName(v string) *SchemaTablesValueFieldsValue {
	s.Name = &v
	return s
}

func (s *SchemaTablesValueFieldsValue) SetPrimaryKey(v bool) *SchemaTablesValueFieldsValue {
	s.PrimaryKey = &v
	return s
}

func (s *SchemaTablesValueFieldsValue) SetType(v string) *SchemaTablesValueFieldsValue {
	s.Type = &v
	return s
}

func (s *SchemaTablesValueFieldsValue) SetJoinWith(v []*string) *SchemaTablesValueFieldsValue {
	s.JoinWith = v
	return s
}

func (s *SchemaTablesValueFieldsValue) SetLabel(v string) *SchemaTablesValueFieldsValue {
	s.Label = &v
	return s
}

func (s *SchemaTablesValueFieldsValue) Validate() error {
	return dara.Validate(s)
}
