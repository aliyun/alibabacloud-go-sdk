// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilterSetting interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*FilterSettingConditions) *FilterSetting
	GetConditions() []*FilterSettingConditions
	SetExpression(v string) *FilterSetting
	GetExpression() *string
	SetRelation(v string) *FilterSetting
	GetRelation() *string
}

type FilterSetting struct {
	// The subscription conditions.
	Conditions []*FilterSettingConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// Use either expression or relation. If expression is not empty, it takes precedence and relation is ignored. If expression is empty or not specified, relation (AND or OR) is used to perform a simple AND/OR operation on all conditions. Condition numbers correspond to the indexes of the conditions array (starting from 1). Each condition evaluates whether a single event field matches by using field (the event field path, which supports dot-notation nesting such as resource.tags.pod), op (the operator, such as CONTAIN, EQ, or IN), and value (the matching value).
	//
	// example:
	//
	// 1 and 2 or 3
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The logical relationship between conditions. This parameter takes effect when expression is empty.
	//
	// example:
	//
	// AND
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
}

func (s FilterSetting) String() string {
	return dara.Prettify(s)
}

func (s FilterSetting) GoString() string {
	return s.String()
}

func (s *FilterSetting) GetConditions() []*FilterSettingConditions {
	return s.Conditions
}

func (s *FilterSetting) GetExpression() *string {
	return s.Expression
}

func (s *FilterSetting) GetRelation() *string {
	return s.Relation
}

func (s *FilterSetting) SetConditions(v []*FilterSettingConditions) *FilterSetting {
	s.Conditions = v
	return s
}

func (s *FilterSetting) SetExpression(v string) *FilterSetting {
	s.Expression = &v
	return s
}

func (s *FilterSetting) SetRelation(v string) *FilterSetting {
	s.Relation = &v
	return s
}

func (s *FilterSetting) Validate() error {
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

type FilterSettingConditions struct {
	// The JSON path of the event field. Dot-notation nesting is supported.
	//
	// example:
	//
	// labels.alertname
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The comparison operator.
	//
	// example:
	//
	// EQ
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// The matching value. Separate multiple values with commas when using IN or NOT_IN.
	//
	// example:
	//
	// CRITICAL
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FilterSettingConditions) String() string {
	return dara.Prettify(s)
}

func (s FilterSettingConditions) GoString() string {
	return s.String()
}

func (s *FilterSettingConditions) GetField() *string {
	return s.Field
}

func (s *FilterSettingConditions) GetOp() *string {
	return s.Op
}

func (s *FilterSettingConditions) GetValue() *string {
	return s.Value
}

func (s *FilterSettingConditions) SetField(v string) *FilterSettingConditions {
	s.Field = &v
	return s
}

func (s *FilterSettingConditions) SetOp(v string) *FilterSettingConditions {
	s.Op = &v
	return s
}

func (s *FilterSettingConditions) SetValue(v string) *FilterSettingConditions {
	s.Value = &v
	return s
}

func (s *FilterSettingConditions) Validate() error {
	return dara.Validate(s)
}
