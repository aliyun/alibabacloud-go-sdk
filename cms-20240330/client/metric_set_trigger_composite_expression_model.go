// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricSetTriggerCompositeExpression interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*MetricSetTriggerSimpleExpression) *MetricSetTriggerCompositeExpression
	GetConditions() []*MetricSetTriggerSimpleExpression
	SetExpressionType(v string) *MetricSetTriggerCompositeExpression
	GetExpressionType() *string
	SetLogicOperator(v string) *MetricSetTriggerCompositeExpression
	GetLogicOperator() *string
}

type MetricSetTriggerCompositeExpression struct {
	// The list of sub-conditions, where each item is a simple comparison expression.
	//
	// example:
	//
	// [{"expressionType":"SIMPLE","queryName":"cpuQuery","operator":"GT","threshold":80}]
	Conditions []*MetricSetTriggerSimpleExpression `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The expression type, which is fixed to COMPOSITE.
	//
	// example:
	//
	// COMPOSITE
	ExpressionType *string `json:"expressionType,omitempty" xml:"expressionType,omitempty"`
	// The logical operator. Valid values: AND (all conditions are met), OR (any condition is met), and UNLESS (the first condition is met and none of the others are met).
	//
	// example:
	//
	// AND
	LogicOperator *string `json:"logicOperator,omitempty" xml:"logicOperator,omitempty"`
}

func (s MetricSetTriggerCompositeExpression) String() string {
	return dara.Prettify(s)
}

func (s MetricSetTriggerCompositeExpression) GoString() string {
	return s.String()
}

func (s *MetricSetTriggerCompositeExpression) GetConditions() []*MetricSetTriggerSimpleExpression {
	return s.Conditions
}

func (s *MetricSetTriggerCompositeExpression) GetExpressionType() *string {
	return s.ExpressionType
}

func (s *MetricSetTriggerCompositeExpression) GetLogicOperator() *string {
	return s.LogicOperator
}

func (s *MetricSetTriggerCompositeExpression) SetConditions(v []*MetricSetTriggerSimpleExpression) *MetricSetTriggerCompositeExpression {
	s.Conditions = v
	return s
}

func (s *MetricSetTriggerCompositeExpression) SetExpressionType(v string) *MetricSetTriggerCompositeExpression {
	s.ExpressionType = &v
	return s
}

func (s *MetricSetTriggerCompositeExpression) SetLogicOperator(v string) *MetricSetTriggerCompositeExpression {
	s.LogicOperator = &v
	return s
}

func (s *MetricSetTriggerCompositeExpression) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
