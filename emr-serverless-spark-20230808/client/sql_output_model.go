// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlOutput interface {
	dara.Model
	String() string
	GoString() string
	SetRows(v []*SqlOutputRows) *SqlOutput
	GetRows() []*SqlOutputRows
	SetSchema(v *SqlOutputSchema) *SqlOutput
	GetSchema() *SqlOutputSchema
}

type SqlOutput struct {
	Rows   []*SqlOutputRows `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
	Schema *SqlOutputSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
}

func (s SqlOutput) String() string {
	return dara.Prettify(s)
}

func (s SqlOutput) GoString() string {
	return s.String()
}

func (s *SqlOutput) GetRows() []*SqlOutputRows {
	return s.Rows
}

func (s *SqlOutput) GetSchema() *SqlOutputSchema {
	return s.Schema
}

func (s *SqlOutput) SetRows(v []*SqlOutputRows) *SqlOutput {
	s.Rows = v
	return s
}

func (s *SqlOutput) SetSchema(v *SqlOutputSchema) *SqlOutput {
	s.Schema = v
	return s
}

func (s *SqlOutput) Validate() error {
	return dara.Validate(s)
}

type SqlOutputRows struct {
	// example:
	//
	// null
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s SqlOutputRows) String() string {
	return dara.Prettify(s)
}

func (s SqlOutputRows) GoString() string {
	return s.String()
}

func (s *SqlOutputRows) GetValues() []*string {
	return s.Values
}

func (s *SqlOutputRows) SetValues(v []*string) *SqlOutputRows {
	s.Values = v
	return s
}

func (s *SqlOutputRows) Validate() error {
	return dara.Validate(s)
}

type SqlOutputSchema struct {
	Fields []*SqlOutputSchemaFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s SqlOutputSchema) String() string {
	return dara.Prettify(s)
}

func (s SqlOutputSchema) GoString() string {
	return s.String()
}

func (s *SqlOutputSchema) GetFields() []*SqlOutputSchemaFields {
	return s.Fields
}

func (s *SqlOutputSchema) SetFields(v []*SqlOutputSchemaFields) *SqlOutputSchema {
	s.Fields = v
	return s
}

func (s *SqlOutputSchema) Validate() error {
	return dara.Validate(s)
}

type SqlOutputSchemaFields struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Nullable *bool   `json:"nullable,omitempty" xml:"nullable,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SqlOutputSchemaFields) String() string {
	return dara.Prettify(s)
}

func (s SqlOutputSchemaFields) GoString() string {
	return s.String()
}

func (s *SqlOutputSchemaFields) GetName() *string {
	return s.Name
}

func (s *SqlOutputSchemaFields) GetNullable() *bool {
	return s.Nullable
}

func (s *SqlOutputSchemaFields) GetType() *string {
	return s.Type
}

func (s *SqlOutputSchemaFields) SetName(v string) *SqlOutputSchemaFields {
	s.Name = &v
	return s
}

func (s *SqlOutputSchemaFields) SetNullable(v bool) *SqlOutputSchemaFields {
	s.Nullable = &v
	return s
}

func (s *SqlOutputSchemaFields) SetType(v string) *SqlOutputSchemaFields {
	s.Type = &v
	return s
}

func (s *SqlOutputSchemaFields) Validate() error {
	return dara.Validate(s)
}
