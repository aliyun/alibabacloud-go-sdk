// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConditionExpression interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ConditionExpression
	GetKey() *string
	SetOperator(v string) *ConditionExpression
	GetOperator() *string
	SetValues(v []*string) *ConditionExpression
	GetValues() []*string
}

type ConditionExpression struct {
	// This parameter is required.
	//
	// example:
	//
	// SupportedMachineTypes
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// in
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ConditionExpression) String() string {
	return dara.Prettify(s)
}

func (s ConditionExpression) GoString() string {
	return s.String()
}

func (s *ConditionExpression) GetKey() *string {
	return s.Key
}

func (s *ConditionExpression) GetOperator() *string {
	return s.Operator
}

func (s *ConditionExpression) GetValues() []*string {
	return s.Values
}

func (s *ConditionExpression) SetKey(v string) *ConditionExpression {
	s.Key = &v
	return s
}

func (s *ConditionExpression) SetOperator(v string) *ConditionExpression {
	s.Operator = &v
	return s
}

func (s *ConditionExpression) SetValues(v []*string) *ConditionExpression {
	s.Values = v
	return s
}

func (s *ConditionExpression) Validate() error {
	return dara.Validate(s)
}
