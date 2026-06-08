// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchemaColumn interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SchemaColumn
	GetName() *string
	SetType(v string) *SchemaColumn
	GetType() *string
}

type SchemaColumn struct {
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// VARCHAR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SchemaColumn) String() string {
	return dara.Prettify(s)
}

func (s SchemaColumn) GoString() string {
	return s.String()
}

func (s *SchemaColumn) GetName() *string {
	return s.Name
}

func (s *SchemaColumn) GetType() *string {
	return s.Type
}

func (s *SchemaColumn) SetName(v string) *SchemaColumn {
	s.Name = &v
	return s
}

func (s *SchemaColumn) SetType(v string) *SchemaColumn {
	s.Type = &v
	return s
}

func (s *SchemaColumn) Validate() error {
	return dara.Validate(s)
}
