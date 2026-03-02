// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRuleCondition interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *RuleCondition
	GetName() *string
	SetPosition(v string) *RuleCondition
	GetPosition() *string
	SetRelation(v string) *RuleCondition
	GetRelation() *string
	SetValue(v string) *RuleCondition
	GetValue() *string
	SetValueType(v string) *RuleCondition
	GetValueType() *string
}

type RuleCondition struct {
	// example:
	//
	// userId
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// header
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// ==
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// example:
	//
	// 10
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s RuleCondition) String() string {
	return dara.Prettify(s)
}

func (s RuleCondition) GoString() string {
	return s.String()
}

func (s *RuleCondition) GetName() *string {
	return s.Name
}

func (s *RuleCondition) GetPosition() *string {
	return s.Position
}

func (s *RuleCondition) GetRelation() *string {
	return s.Relation
}

func (s *RuleCondition) GetValue() *string {
	return s.Value
}

func (s *RuleCondition) GetValueType() *string {
	return s.ValueType
}

func (s *RuleCondition) SetName(v string) *RuleCondition {
	s.Name = &v
	return s
}

func (s *RuleCondition) SetPosition(v string) *RuleCondition {
	s.Position = &v
	return s
}

func (s *RuleCondition) SetRelation(v string) *RuleCondition {
	s.Relation = &v
	return s
}

func (s *RuleCondition) SetValue(v string) *RuleCondition {
	s.Value = &v
	return s
}

func (s *RuleCondition) SetValueType(v string) *RuleCondition {
	s.ValueType = &v
	return s
}

func (s *RuleCondition) Validate() error {
	return dara.Validate(s)
}
