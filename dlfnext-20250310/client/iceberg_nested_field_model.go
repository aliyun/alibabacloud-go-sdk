// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIcebergNestedField interface {
	dara.Model
	String() string
	GoString() string
	SetDoc(v string) *IcebergNestedField
	GetDoc() *string
	SetId(v int64) *IcebergNestedField
	GetId() *int64
	SetName(v string) *IcebergNestedField
	GetName() *string
	SetOptional(v bool) *IcebergNestedField
	GetOptional() *bool
	SetType(v string) *IcebergNestedField
	GetType() *string
}

type IcebergNestedField struct {
	Doc      *string `json:"doc,omitempty" xml:"doc,omitempty"`
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Optional *bool   `json:"optional,omitempty" xml:"optional,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IcebergNestedField) String() string {
	return dara.Prettify(s)
}

func (s IcebergNestedField) GoString() string {
	return s.String()
}

func (s *IcebergNestedField) GetDoc() *string {
	return s.Doc
}

func (s *IcebergNestedField) GetId() *int64 {
	return s.Id
}

func (s *IcebergNestedField) GetName() *string {
	return s.Name
}

func (s *IcebergNestedField) GetOptional() *bool {
	return s.Optional
}

func (s *IcebergNestedField) GetType() *string {
	return s.Type
}

func (s *IcebergNestedField) SetDoc(v string) *IcebergNestedField {
	s.Doc = &v
	return s
}

func (s *IcebergNestedField) SetId(v int64) *IcebergNestedField {
	s.Id = &v
	return s
}

func (s *IcebergNestedField) SetName(v string) *IcebergNestedField {
	s.Name = &v
	return s
}

func (s *IcebergNestedField) SetOptional(v bool) *IcebergNestedField {
	s.Optional = &v
	return s
}

func (s *IcebergNestedField) SetType(v string) *IcebergNestedField {
	s.Type = &v
	return s
}

func (s *IcebergNestedField) Validate() error {
	return dara.Validate(s)
}
