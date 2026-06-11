// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUmodelEntityFilter interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *UmodelEntityFilter
	GetField() *string
	SetOperator(v string) *UmodelEntityFilter
	GetOperator() *string
	SetValue(v string) *UmodelEntityFilter
	GetValue() *string
}

type UmodelEntityFilter struct {
	// The field name to filter on.
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The comparison operator to use. Supported operators include `=`, `>`, `<`, `!=`, `IN`, and `NOT IN`.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The value to compare the field against.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UmodelEntityFilter) String() string {
	return dara.Prettify(s)
}

func (s UmodelEntityFilter) GoString() string {
	return s.String()
}

func (s *UmodelEntityFilter) GetField() *string {
	return s.Field
}

func (s *UmodelEntityFilter) GetOperator() *string {
	return s.Operator
}

func (s *UmodelEntityFilter) GetValue() *string {
	return s.Value
}

func (s *UmodelEntityFilter) SetField(v string) *UmodelEntityFilter {
	s.Field = &v
	return s
}

func (s *UmodelEntityFilter) SetOperator(v string) *UmodelEntityFilter {
	s.Operator = &v
	return s
}

func (s *UmodelEntityFilter) SetValue(v string) *UmodelEntityFilter {
	s.Value = &v
	return s
}

func (s *UmodelEntityFilter) Validate() error {
	return dara.Validate(s)
}
