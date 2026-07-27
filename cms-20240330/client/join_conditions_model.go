// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinConditions interface {
	dara.Model
	String() string
	GoString() string
	SetLhsField(v string) *JoinConditions
	GetLhsField() *string
	SetOperator(v string) *JoinConditions
	GetOperator() *string
	SetRhsField(v string) *JoinConditions
	GetRhsField() *string
}

type JoinConditions struct {
	LhsField *string `json:"lhsField,omitempty" xml:"lhsField,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	RhsField *string `json:"rhsField,omitempty" xml:"rhsField,omitempty"`
}

func (s JoinConditions) String() string {
	return dara.Prettify(s)
}

func (s JoinConditions) GoString() string {
	return s.String()
}

func (s *JoinConditions) GetLhsField() *string {
	return s.LhsField
}

func (s *JoinConditions) GetOperator() *string {
	return s.Operator
}

func (s *JoinConditions) GetRhsField() *string {
	return s.RhsField
}

func (s *JoinConditions) SetLhsField(v string) *JoinConditions {
	s.LhsField = &v
	return s
}

func (s *JoinConditions) SetOperator(v string) *JoinConditions {
	s.Operator = &v
	return s
}

func (s *JoinConditions) SetRhsField(v string) *JoinConditions {
	s.RhsField = &v
	return s
}

func (s *JoinConditions) Validate() error {
	return dara.Validate(s)
}
