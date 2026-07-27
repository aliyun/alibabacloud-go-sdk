// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabelFilters interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *LabelFilters
	GetName() *string
	SetOperator(v string) *LabelFilters
	GetOperator() *string
	SetValue(v string) *LabelFilters
	GetValue() *string
}

type LabelFilters struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s LabelFilters) String() string {
	return dara.Prettify(s)
}

func (s LabelFilters) GoString() string {
	return s.String()
}

func (s *LabelFilters) GetName() *string {
	return s.Name
}

func (s *LabelFilters) GetOperator() *string {
	return s.Operator
}

func (s *LabelFilters) GetValue() *string {
	return s.Value
}

func (s *LabelFilters) SetName(v string) *LabelFilters {
	s.Name = &v
	return s
}

func (s *LabelFilters) SetOperator(v string) *LabelFilters {
	s.Operator = &v
	return s
}

func (s *LabelFilters) SetValue(v string) *LabelFilters {
	s.Value = &v
	return s
}

func (s *LabelFilters) Validate() error {
	return dara.Validate(s)
}
