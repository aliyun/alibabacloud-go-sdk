// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConditionalRule interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *EventFilter) *ConditionalRule
	GetFilter() *EventFilter
	SetModifyTime(v string) *ConditionalRule
	GetModifyTime() *string
	SetName(v string) *ConditionalRule
	GetName() *string
	SetOperator(v string) *ConditionalRule
	GetOperator() *string
	SetSampleRate(v float32) *ConditionalRule
	GetSampleRate() *float32
}

type ConditionalRule struct {
	Filter     *EventFilter `json:"Filter,omitempty" xml:"Filter,omitempty"`
	ModifyTime *string      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string      `json:"Name,omitempty" xml:"Name,omitempty"`
	Operator   *string      `json:"Operator,omitempty" xml:"Operator,omitempty"`
	SampleRate *float32     `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s ConditionalRule) String() string {
	return dara.Prettify(s)
}

func (s ConditionalRule) GoString() string {
	return s.String()
}

func (s *ConditionalRule) GetFilter() *EventFilter {
	return s.Filter
}

func (s *ConditionalRule) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ConditionalRule) GetName() *string {
	return s.Name
}

func (s *ConditionalRule) GetOperator() *string {
	return s.Operator
}

func (s *ConditionalRule) GetSampleRate() *float32 {
	return s.SampleRate
}

func (s *ConditionalRule) SetFilter(v *EventFilter) *ConditionalRule {
	s.Filter = v
	return s
}

func (s *ConditionalRule) SetModifyTime(v string) *ConditionalRule {
	s.ModifyTime = &v
	return s
}

func (s *ConditionalRule) SetName(v string) *ConditionalRule {
	s.Name = &v
	return s
}

func (s *ConditionalRule) SetOperator(v string) *ConditionalRule {
	s.Operator = &v
	return s
}

func (s *ConditionalRule) SetSampleRate(v float32) *ConditionalRule {
	s.SampleRate = &v
	return s
}

func (s *ConditionalRule) Validate() error {
	return dara.Validate(s)
}
