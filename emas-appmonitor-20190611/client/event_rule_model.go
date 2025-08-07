// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventRule interface {
  dara.Model
  String() string
  GoString() string
  SetConditional(v []*ConditionalRule) *EventRule
  GetConditional() []*ConditionalRule 
  SetEnable(v bool) *EventRule
  GetEnable() *bool 
  SetEventId(v string) *EventRule
  GetEventId() *string 
  SetModifyTime(v string) *EventRule
  GetModifyTime() *string 
  SetOperator(v string) *EventRule
  GetOperator() *string 
  SetSampleRate(v float32) *EventRule
  GetSampleRate() *float32 
}

type EventRule struct {
  Conditional []*ConditionalRule `json:"Conditional,omitempty" xml:"Conditional,omitempty" type:"Repeated"`
  Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
  EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
  ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
  Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
  SampleRate *float32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s EventRule) String() string {
  return dara.Prettify(s)
}

func (s EventRule) GoString() string {
  return s.String()
}

func (s *EventRule) GetConditional() []*ConditionalRule  {
  return s.Conditional
}

func (s *EventRule) GetEnable() *bool  {
  return s.Enable
}

func (s *EventRule) GetEventId() *string  {
  return s.EventId
}

func (s *EventRule) GetModifyTime() *string  {
  return s.ModifyTime
}

func (s *EventRule) GetOperator() *string  {
  return s.Operator
}

func (s *EventRule) GetSampleRate() *float32  {
  return s.SampleRate
}

func (s *EventRule) SetConditional(v []*ConditionalRule) *EventRule {
  s.Conditional = v
  return s
}

func (s *EventRule) SetEnable(v bool) *EventRule {
  s.Enable = &v
  return s
}

func (s *EventRule) SetEventId(v string) *EventRule {
  s.EventId = &v
  return s
}

func (s *EventRule) SetModifyTime(v string) *EventRule {
  s.ModifyTime = &v
  return s
}

func (s *EventRule) SetOperator(v string) *EventRule {
  s.Operator = &v
  return s
}

func (s *EventRule) SetSampleRate(v float32) *EventRule {
  s.SampleRate = &v
  return s
}

func (s *EventRule) Validate() error {
  return dara.Validate(s)
}

