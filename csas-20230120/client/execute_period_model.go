// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePeriod interface {
  dara.Model
  String() string
  GoString() string
  SetEffectDay(v *ExecutePeriodEffectDay) *ExecutePeriod
  GetEffectDay() *ExecutePeriodEffectDay 
  SetEffectTime(v *ExecutePeriodEffectTime) *ExecutePeriod
  GetEffectTime() *ExecutePeriodEffectTime 
  SetScheduleEffect(v *ExecutePeriodScheduleEffect) *ExecutePeriod
  GetScheduleEffect() *ExecutePeriodScheduleEffect 
  SetValidType(v string) *ExecutePeriod
  GetValidType() *string 
}

type ExecutePeriod struct {
  EffectDay *ExecutePeriodEffectDay `json:"EffectDay,omitempty" xml:"EffectDay,omitempty" type:"Struct"`
  EffectTime *ExecutePeriodEffectTime `json:"EffectTime,omitempty" xml:"EffectTime,omitempty" type:"Struct"`
  ScheduleEffect *ExecutePeriodScheduleEffect `json:"ScheduleEffect,omitempty" xml:"ScheduleEffect,omitempty" type:"Struct"`
  ValidType *string `json:"ValidType,omitempty" xml:"ValidType,omitempty"`
}

func (s ExecutePeriod) String() string {
  return dara.Prettify(s)
}

func (s ExecutePeriod) GoString() string {
  return s.String()
}

func (s *ExecutePeriod) GetEffectDay() *ExecutePeriodEffectDay  {
  return s.EffectDay
}

func (s *ExecutePeriod) GetEffectTime() *ExecutePeriodEffectTime  {
  return s.EffectTime
}

func (s *ExecutePeriod) GetScheduleEffect() *ExecutePeriodScheduleEffect  {
  return s.ScheduleEffect
}

func (s *ExecutePeriod) GetValidType() *string  {
  return s.ValidType
}

func (s *ExecutePeriod) SetEffectDay(v *ExecutePeriodEffectDay) *ExecutePeriod {
  s.EffectDay = v
  return s
}

func (s *ExecutePeriod) SetEffectTime(v *ExecutePeriodEffectTime) *ExecutePeriod {
  s.EffectTime = v
  return s
}

func (s *ExecutePeriod) SetScheduleEffect(v *ExecutePeriodScheduleEffect) *ExecutePeriod {
  s.ScheduleEffect = v
  return s
}

func (s *ExecutePeriod) SetValidType(v string) *ExecutePeriod {
  s.ValidType = &v
  return s
}

func (s *ExecutePeriod) Validate() error {
  if s.EffectDay != nil {
    if err := s.EffectDay.Validate(); err != nil {
      return err
    }
  }
  if s.EffectTime != nil {
    if err := s.EffectTime.Validate(); err != nil {
      return err
    }
  }
  if s.ScheduleEffect != nil {
    if err := s.ScheduleEffect.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecutePeriodEffectDay struct {
  Friday *bool `json:"Friday,omitempty" xml:"Friday,omitempty"`
  Monday *bool `json:"Monday,omitempty" xml:"Monday,omitempty"`
  Saturday *bool `json:"Saturday,omitempty" xml:"Saturday,omitempty"`
  Sunday *bool `json:"Sunday,omitempty" xml:"Sunday,omitempty"`
  Thursday *bool `json:"Thursday,omitempty" xml:"Thursday,omitempty"`
  Tuesday *bool `json:"Tuesday,omitempty" xml:"Tuesday,omitempty"`
  Wednesday *bool `json:"Wednesday,omitempty" xml:"Wednesday,omitempty"`
}

func (s ExecutePeriodEffectDay) String() string {
  return dara.Prettify(s)
}

func (s ExecutePeriodEffectDay) GoString() string {
  return s.String()
}

func (s *ExecutePeriodEffectDay) GetFriday() *bool  {
  return s.Friday
}

func (s *ExecutePeriodEffectDay) GetMonday() *bool  {
  return s.Monday
}

func (s *ExecutePeriodEffectDay) GetSaturday() *bool  {
  return s.Saturday
}

func (s *ExecutePeriodEffectDay) GetSunday() *bool  {
  return s.Sunday
}

func (s *ExecutePeriodEffectDay) GetThursday() *bool  {
  return s.Thursday
}

func (s *ExecutePeriodEffectDay) GetTuesday() *bool  {
  return s.Tuesday
}

func (s *ExecutePeriodEffectDay) GetWednesday() *bool  {
  return s.Wednesday
}

func (s *ExecutePeriodEffectDay) SetFriday(v bool) *ExecutePeriodEffectDay {
  s.Friday = &v
  return s
}

func (s *ExecutePeriodEffectDay) SetMonday(v bool) *ExecutePeriodEffectDay {
  s.Monday = &v
  return s
}

func (s *ExecutePeriodEffectDay) SetSaturday(v bool) *ExecutePeriodEffectDay {
  s.Saturday = &v
  return s
}

func (s *ExecutePeriodEffectDay) SetSunday(v bool) *ExecutePeriodEffectDay {
  s.Sunday = &v
  return s
}

func (s *ExecutePeriodEffectDay) SetThursday(v bool) *ExecutePeriodEffectDay {
  s.Thursday = &v
  return s
}

func (s *ExecutePeriodEffectDay) SetTuesday(v bool) *ExecutePeriodEffectDay {
  s.Tuesday = &v
  return s
}

func (s *ExecutePeriodEffectDay) SetWednesday(v bool) *ExecutePeriodEffectDay {
  s.Wednesday = &v
  return s
}

func (s *ExecutePeriodEffectDay) Validate() error {
  return dara.Validate(s)
}

type ExecutePeriodEffectTime struct {
  End *string `json:"End,omitempty" xml:"End,omitempty"`
  Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ExecutePeriodEffectTime) String() string {
  return dara.Prettify(s)
}

func (s ExecutePeriodEffectTime) GoString() string {
  return s.String()
}

func (s *ExecutePeriodEffectTime) GetEnd() *string  {
  return s.End
}

func (s *ExecutePeriodEffectTime) GetStart() *string  {
  return s.Start
}

func (s *ExecutePeriodEffectTime) SetEnd(v string) *ExecutePeriodEffectTime {
  s.End = &v
  return s
}

func (s *ExecutePeriodEffectTime) SetStart(v string) *ExecutePeriodEffectTime {
  s.Start = &v
  return s
}

func (s *ExecutePeriodEffectTime) Validate() error {
  return dara.Validate(s)
}

type ExecutePeriodScheduleEffect struct {
  Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
  Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s ExecutePeriodScheduleEffect) String() string {
  return dara.Prettify(s)
}

func (s ExecutePeriodScheduleEffect) GoString() string {
  return s.String()
}

func (s *ExecutePeriodScheduleEffect) GetFrequency() *string  {
  return s.Frequency
}

func (s *ExecutePeriodScheduleEffect) GetInterval() *int64  {
  return s.Interval
}

func (s *ExecutePeriodScheduleEffect) SetFrequency(v string) *ExecutePeriodScheduleEffect {
  s.Frequency = &v
  return s
}

func (s *ExecutePeriodScheduleEffect) SetInterval(v int64) *ExecutePeriodScheduleEffect {
  s.Interval = &v
  return s
}

func (s *ExecutePeriodScheduleEffect) Validate() error {
  return dara.Validate(s)
}

