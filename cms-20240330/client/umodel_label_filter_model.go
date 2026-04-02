// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUmodelLabelFilter interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UmodelLabelFilter
	GetName() *string
	SetOperator(v string) *UmodelLabelFilter
	GetOperator() *string
	SetValue(v string) *UmodelLabelFilter
	GetValue() *string
}

type UmodelLabelFilter struct {
	// 标签名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签过滤操作符，仅支持 = 或 !=
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 标签值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UmodelLabelFilter) String() string {
	return dara.Prettify(s)
}

func (s UmodelLabelFilter) GoString() string {
	return s.String()
}

func (s *UmodelLabelFilter) GetName() *string {
	return s.Name
}

func (s *UmodelLabelFilter) GetOperator() *string {
	return s.Operator
}

func (s *UmodelLabelFilter) GetValue() *string {
	return s.Value
}

func (s *UmodelLabelFilter) SetName(v string) *UmodelLabelFilter {
	s.Name = &v
	return s
}

func (s *UmodelLabelFilter) SetOperator(v string) *UmodelLabelFilter {
	s.Operator = &v
	return s
}

func (s *UmodelLabelFilter) SetValue(v string) *UmodelLabelFilter {
	s.Value = &v
	return s
}

func (s *UmodelLabelFilter) Validate() error {
	return dara.Validate(s)
}
