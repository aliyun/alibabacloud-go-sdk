// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaintainWindowForModify interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *MaintainWindowForModify
	GetDescription() *string
	SetEffectTimeRange(v *MaintainWindowForModifyEffectTimeRange) *MaintainWindowForModify
	GetEffectTimeRange() *MaintainWindowForModifyEffectTimeRange
	SetEffective(v string) *MaintainWindowForModify
	GetEffective() *string
	SetEndTime(v string) *MaintainWindowForModify
	GetEndTime() *string
	SetFilterSetting(v *FilterSetting) *MaintainWindowForModify
	GetFilterSetting() *FilterSetting
	SetMaintainWindowName(v string) *MaintainWindowForModify
	GetMaintainWindowName() *string
	SetStartTime(v string) *MaintainWindowForModify
	GetStartTime() *string
}

type MaintainWindowForModify struct {
	// 描述
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 生效时间范围
	EffectTimeRange *MaintainWindowForModifyEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	// Crontab 表达式
	//
	// example:
	//
	// 	- 	- 14-18 ? 	- *
	Effective *string `json:"effective,omitempty" xml:"effective,omitempty"`
	// 生效结束时间
	//
	// example:
	//
	// 2024-09-05 09:30:40
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 筛选条件
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// 名称
	//
	// This parameter is required.
	//
	// example:
	//
	// test strategy
	MaintainWindowName *string `json:"maintainWindowName,omitempty" xml:"maintainWindowName,omitempty"`
	// 生效开始时间
	//
	// example:
	//
	// 2025-04-11 07:55:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s MaintainWindowForModify) String() string {
	return dara.Prettify(s)
}

func (s MaintainWindowForModify) GoString() string {
	return s.String()
}

func (s *MaintainWindowForModify) GetDescription() *string {
	return s.Description
}

func (s *MaintainWindowForModify) GetEffectTimeRange() *MaintainWindowForModifyEffectTimeRange {
	return s.EffectTimeRange
}

func (s *MaintainWindowForModify) GetEffective() *string {
	return s.Effective
}

func (s *MaintainWindowForModify) GetEndTime() *string {
	return s.EndTime
}

func (s *MaintainWindowForModify) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *MaintainWindowForModify) GetMaintainWindowName() *string {
	return s.MaintainWindowName
}

func (s *MaintainWindowForModify) GetStartTime() *string {
	return s.StartTime
}

func (s *MaintainWindowForModify) SetDescription(v string) *MaintainWindowForModify {
	s.Description = &v
	return s
}

func (s *MaintainWindowForModify) SetEffectTimeRange(v *MaintainWindowForModifyEffectTimeRange) *MaintainWindowForModify {
	s.EffectTimeRange = v
	return s
}

func (s *MaintainWindowForModify) SetEffective(v string) *MaintainWindowForModify {
	s.Effective = &v
	return s
}

func (s *MaintainWindowForModify) SetEndTime(v string) *MaintainWindowForModify {
	s.EndTime = &v
	return s
}

func (s *MaintainWindowForModify) SetFilterSetting(v *FilterSetting) *MaintainWindowForModify {
	s.FilterSetting = v
	return s
}

func (s *MaintainWindowForModify) SetMaintainWindowName(v string) *MaintainWindowForModify {
	s.MaintainWindowName = &v
	return s
}

func (s *MaintainWindowForModify) SetStartTime(v string) *MaintainWindowForModify {
	s.StartTime = &v
	return s
}

func (s *MaintainWindowForModify) Validate() error {
	if s.EffectTimeRange != nil {
		if err := s.EffectTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MaintainWindowForModifyEffectTimeRange struct {
	// 生效天(周一到周日)
	DayInWeek []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	// 结束时间(分钟数)
	//
	// example:
	//
	// 60
	EndTimeInMinute *int32 `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	// 开始时间(分钟数)
	//
	// example:
	//
	// 60
	StartTimeInMinute *int32 `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	// 时区
	//
	// example:
	//
	// +08:00
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s MaintainWindowForModifyEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s MaintainWindowForModifyEffectTimeRange) GoString() string {
	return s.String()
}

func (s *MaintainWindowForModifyEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *MaintainWindowForModifyEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *MaintainWindowForModifyEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *MaintainWindowForModifyEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *MaintainWindowForModifyEffectTimeRange) SetDayInWeek(v []*int32) *MaintainWindowForModifyEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *MaintainWindowForModifyEffectTimeRange) SetEndTimeInMinute(v int32) *MaintainWindowForModifyEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *MaintainWindowForModifyEffectTimeRange) SetStartTimeInMinute(v int32) *MaintainWindowForModifyEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *MaintainWindowForModifyEffectTimeRange) SetTimeZone(v string) *MaintainWindowForModifyEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *MaintainWindowForModifyEffectTimeRange) Validate() error {
	return dara.Validate(s)
}
