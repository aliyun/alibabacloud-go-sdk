// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSlsJoinCondition interface {
	dara.Model
	String() string
	GoString() string
	SetLhsField(v string) *SlsJoinCondition
	GetLhsField() *string
	SetOperator(v string) *SlsJoinCondition
	GetOperator() *string
	SetRhsField(v string) *SlsJoinCondition
	GetRhsField() *string
}

type SlsJoinCondition struct {
	LhsField *string `json:"lhsField,omitempty" xml:"lhsField,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	RhsField *string `json:"rhsField,omitempty" xml:"rhsField,omitempty"`
}

func (s SlsJoinCondition) String() string {
	return dara.Prettify(s)
}

func (s SlsJoinCondition) GoString() string {
	return s.String()
}

func (s *SlsJoinCondition) GetLhsField() *string {
	return s.LhsField
}

func (s *SlsJoinCondition) GetOperator() *string {
	return s.Operator
}

func (s *SlsJoinCondition) GetRhsField() *string {
	return s.RhsField
}

func (s *SlsJoinCondition) SetLhsField(v string) *SlsJoinCondition {
	s.LhsField = &v
	return s
}

func (s *SlsJoinCondition) SetOperator(v string) *SlsJoinCondition {
	s.Operator = &v
	return s
}

func (s *SlsJoinCondition) SetRhsField(v string) *SlsJoinCondition {
	s.RhsField = &v
	return s
}

func (s *SlsJoinCondition) Validate() error {
	return dara.Validate(s)
}
