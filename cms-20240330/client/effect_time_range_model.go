// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectTimeRange interface {
  dara.Model
  String() string
  GoString() string
  SetDayInWeek(v []*int32) *EffectTimeRange
  GetDayInWeek() []*int32 
  SetEndTimeInMinute(v int32) *EffectTimeRange
  GetEndTimeInMinute() *int32 
  SetStartTimeInMinute(v int32) *EffectTimeRange
  GetStartTimeInMinute() *int32 
  SetTimeZone(v string) *EffectTimeRange
  GetTimeZone() *string 
}

type EffectTimeRange struct {
  DayInWeek []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
  EndTimeInMinute *int32 `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
  StartTimeInMinute *int32 `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
  TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s EffectTimeRange) String() string {
  return dara.Prettify(s)
}

func (s EffectTimeRange) GoString() string {
  return s.String()
}

func (s *EffectTimeRange) GetDayInWeek() []*int32  {
  return s.DayInWeek
}

func (s *EffectTimeRange) GetEndTimeInMinute() *int32  {
  return s.EndTimeInMinute
}

func (s *EffectTimeRange) GetStartTimeInMinute() *int32  {
  return s.StartTimeInMinute
}

func (s *EffectTimeRange) GetTimeZone() *string  {
  return s.TimeZone
}

func (s *EffectTimeRange) SetDayInWeek(v []*int32) *EffectTimeRange {
  s.DayInWeek = v
  return s
}

func (s *EffectTimeRange) SetEndTimeInMinute(v int32) *EffectTimeRange {
  s.EndTimeInMinute = &v
  return s
}

func (s *EffectTimeRange) SetStartTimeInMinute(v int32) *EffectTimeRange {
  s.StartTimeInMinute = &v
  return s
}

func (s *EffectTimeRange) SetTimeZone(v string) *EffectTimeRange {
  s.TimeZone = &v
  return s
}

func (s *EffectTimeRange) Validate() error {
  return dara.Validate(s)
}

