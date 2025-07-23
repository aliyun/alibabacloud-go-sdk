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
	Conditions []*FilterSettingConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Expression *string                    `json:"expression,omitempty" xml:"expression,omitempty"`
	Relation   *string                    `json:"relation,omitempty" xml:"relation,omitempty"`
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
	return dara.Validate(s)
}

type FilterSettingConditions struct {
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	Op    *string `json:"op,omitempty" xml:"op,omitempty"`
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
