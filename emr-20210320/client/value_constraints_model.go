// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValueConstraints interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v int32) *ValueConstraints
	GetDefaultValue() *int32
	SetEnd(v int32) *ValueConstraints
	GetEnd() *int32
	SetStart(v int32) *ValueConstraints
	GetStart() *int32
	SetStep(v int32) *ValueConstraints
	GetStep() *int32
	SetType(v string) *ValueConstraints
	GetType() *string
	SetValues(v []*int32) *ValueConstraints
	GetValues() []*int32
}

type ValueConstraints struct {
	// 默认值。
	DefaultValue *int32 `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// 结束值。
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// 起始值。
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// 步长。
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// 值限制类型。
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 枚举值。
	//
	// example:
	//
	// null
	Values []*int32 `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ValueConstraints) String() string {
	return dara.Prettify(s)
}

func (s ValueConstraints) GoString() string {
	return s.String()
}

func (s *ValueConstraints) GetDefaultValue() *int32 {
	return s.DefaultValue
}

func (s *ValueConstraints) GetEnd() *int32 {
	return s.End
}

func (s *ValueConstraints) GetStart() *int32 {
	return s.Start
}

func (s *ValueConstraints) GetStep() *int32 {
	return s.Step
}

func (s *ValueConstraints) GetType() *string {
	return s.Type
}

func (s *ValueConstraints) GetValues() []*int32 {
	return s.Values
}

func (s *ValueConstraints) SetDefaultValue(v int32) *ValueConstraints {
	s.DefaultValue = &v
	return s
}

func (s *ValueConstraints) SetEnd(v int32) *ValueConstraints {
	s.End = &v
	return s
}

func (s *ValueConstraints) SetStart(v int32) *ValueConstraints {
	s.Start = &v
	return s
}

func (s *ValueConstraints) SetStep(v int32) *ValueConstraints {
	s.Step = &v
	return s
}

func (s *ValueConstraints) SetType(v string) *ValueConstraints {
	s.Type = &v
	return s
}

func (s *ValueConstraints) SetValues(v []*int32) *ValueConstraints {
	s.Values = v
	return s
}

func (s *ValueConstraints) Validate() error {
	return dara.Validate(s)
}
