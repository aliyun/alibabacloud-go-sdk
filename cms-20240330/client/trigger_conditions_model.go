// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerConditions interface {
	dara.Model
	String() string
	GoString() string
	SetExpressionType(v string) *TriggerConditions
	GetExpressionType() *string
	SetMax(v float64) *TriggerConditions
	GetMax() *float64
	SetMin(v float64) *TriggerConditions
	GetMin() *float64
	SetOperator(v string) *TriggerConditions
	GetOperator() *string
	SetQueryName(v string) *TriggerConditions
	GetQueryName() *string
	SetThreshold(v float64) *TriggerConditions
	GetThreshold() *float64
}

type TriggerConditions struct {
	ExpressionType *string  `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	Max            *float64 `json:"max,omitempty" xml:"max,omitempty"`
	Min            *float64 `json:"min,omitempty" xml:"min,omitempty"`
	Operator       *string  `json:"operator,omitempty" xml:"operator,omitempty"`
	QueryName      *string  `json:"queryName,omitempty" xml:"queryName,omitempty"`
	Threshold      *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s TriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s TriggerConditions) GoString() string {
	return s.String()
}

func (s *TriggerConditions) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *TriggerConditions) GetMax() *float64 {
	return s.Max
}

func (s *TriggerConditions) GetMin() *float64 {
	return s.Min
}

func (s *TriggerConditions) GetOperator() *string {
	return s.Operator
}

func (s *TriggerConditions) GetQueryName() *string {
	return s.QueryName
}

func (s *TriggerConditions) GetThreshold() *float64 {
	return s.Threshold
}

func (s *TriggerConditions) SetExpressionType(v string) *TriggerConditions {
	s.ExpressionType = &v
	return s
}

func (s *TriggerConditions) SetMax(v float64) *TriggerConditions {
	s.Max = &v
	return s
}

func (s *TriggerConditions) SetMin(v float64) *TriggerConditions {
	s.Min = &v
	return s
}

func (s *TriggerConditions) SetOperator(v string) *TriggerConditions {
	s.Operator = &v
	return s
}

func (s *TriggerConditions) SetQueryName(v string) *TriggerConditions {
	s.QueryName = &v
	return s
}

func (s *TriggerConditions) SetThreshold(v float64) *TriggerConditions {
	s.Threshold = &v
	return s
}

func (s *TriggerConditions) Validate() error {
	return dara.Validate(s)
}
