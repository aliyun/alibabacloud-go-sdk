// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilterCondition interface {
	dara.Model
	String() string
	GoString() string
	SetField(v string) *FilterCondition
	GetField() *string
	SetOp(v string) *FilterCondition
	GetOp() *string
	SetValue(v string) *FilterCondition
	GetValue() *string
}

type FilterCondition struct {
	// The name of the field to filter on.
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The comparison operator, such as `equals` or `startsWith`.
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// The value to match for the specified field and operator.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FilterCondition) String() string {
	return dara.Prettify(s)
}

func (s FilterCondition) GoString() string {
	return s.String()
}

func (s *FilterCondition) GetField() *string {
	return s.Field
}

func (s *FilterCondition) GetOp() *string {
	return s.Op
}

func (s *FilterCondition) GetValue() *string {
	return s.Value
}

func (s *FilterCondition) SetField(v string) *FilterCondition {
	s.Field = &v
	return s
}

func (s *FilterCondition) SetOp(v string) *FilterCondition {
	s.Op = &v
	return s
}

func (s *FilterCondition) SetValue(v string) *FilterCondition {
	s.Value = &v
	return s
}

func (s *FilterCondition) Validate() error {
	return dara.Validate(s)
}
