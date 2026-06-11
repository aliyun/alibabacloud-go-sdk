// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUmodelEntityField interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *UmodelEntityField
	GetField() *string
	SetValue(v string) *UmodelEntityField
	GetValue() *string
}

type UmodelEntityField struct {
	// The name of the entity field.
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The field alias or display value.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UmodelEntityField) String() string {
	return dara.Prettify(s)
}

func (s UmodelEntityField) GoString() string {
	return s.String()
}

func (s *UmodelEntityField) GetField() *string {
	return s.Field
}

func (s *UmodelEntityField) GetValue() *string {
	return s.Value
}

func (s *UmodelEntityField) SetField(v string) *UmodelEntityField {
	s.Field = &v
	return s
}

func (s *UmodelEntityField) SetValue(v string) *UmodelEntityField {
	s.Value = &v
	return s
}

func (s *UmodelEntityField) Validate() error {
	return dara.Validate(s)
}
