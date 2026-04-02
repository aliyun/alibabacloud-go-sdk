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
	// 过滤字段名
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// 过滤操作符，仅支持 = 或 !=
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 过滤值
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
