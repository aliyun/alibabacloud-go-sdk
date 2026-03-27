// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaintainWindowForView interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *MaintainWindowForView
	GetCreateTime() *string
	SetDescription(v string) *MaintainWindowForView
	GetDescription() *string
	SetEffectTimeRange(v *MaintainWindowForViewEffectTimeRange) *MaintainWindowForView
	GetEffectTimeRange() *MaintainWindowForViewEffectTimeRange
	SetEffective(v string) *MaintainWindowForView
	GetEffective() *string
	SetEnable(v bool) *MaintainWindowForView
	GetEnable() *bool
	SetEndTime(v string) *MaintainWindowForView
	GetEndTime() *string
	SetFilterSetting(v *FilterSetting) *MaintainWindowForView
	GetFilterSetting() *FilterSetting
	SetMaintainWindowId(v string) *MaintainWindowForView
	GetMaintainWindowId() *string
	SetMaintainWindowName(v string) *MaintainWindowForView
	GetMaintainWindowName() *string
	SetStartTime(v string) *MaintainWindowForView
	GetStartTime() *string
	SetUpdateTime(v string) *MaintainWindowForView
	GetUpdateTime() *string
	SetUserId(v string) *MaintainWindowForView
	GetUserId() *string
	SetWorkspace(v string) *MaintainWindowForView
	GetWorkspace() *string
}

type MaintainWindowForView struct {
	// Creation time.
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Effective time range.
	EffectTimeRange *MaintainWindowForViewEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	// Crontab expression.
	//
	// example:
	//
	// 	- 	- 14-18 ? 	- *
	Effective *string `json:"effective,omitempty" xml:"effective,omitempty"`
	// Whether enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Effective end time.
	//
	// example:
	//
	// 2024-11-26 12:02:01
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Filtering conditions.
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// UUID
	//
	// example:
	//
	// 123-12-312-31-23123
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// Name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test policy.
	MaintainWindowName *string `json:"maintainWindowName,omitempty" xml:"maintainWindowName,omitempty"`
	// Effective start time.
	//
	// example:
	//
	// 2024-11-26 10:02:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Update time.
	//
	// example:
	//
	// 2025-05-07T02:19:05Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// User ID.
	//
	// example:
	//
	// 12312312***
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// workspace
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s MaintainWindowForView) String() string {
	return dara.Prettify(s)
}

func (s MaintainWindowForView) GoString() string {
	return s.String()
}

func (s *MaintainWindowForView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MaintainWindowForView) GetDescription() *string {
	return s.Description
}

func (s *MaintainWindowForView) GetEffectTimeRange() *MaintainWindowForViewEffectTimeRange {
	return s.EffectTimeRange
}

func (s *MaintainWindowForView) GetEffective() *string {
	return s.Effective
}

func (s *MaintainWindowForView) GetEnable() *bool {
	return s.Enable
}

func (s *MaintainWindowForView) GetEndTime() *string {
	return s.EndTime
}

func (s *MaintainWindowForView) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *MaintainWindowForView) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *MaintainWindowForView) GetMaintainWindowName() *string {
	return s.MaintainWindowName
}

func (s *MaintainWindowForView) GetStartTime() *string {
	return s.StartTime
}

func (s *MaintainWindowForView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *MaintainWindowForView) GetUserId() *string {
	return s.UserId
}

func (s *MaintainWindowForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *MaintainWindowForView) SetCreateTime(v string) *MaintainWindowForView {
	s.CreateTime = &v
	return s
}

func (s *MaintainWindowForView) SetDescription(v string) *MaintainWindowForView {
	s.Description = &v
	return s
}

func (s *MaintainWindowForView) SetEffectTimeRange(v *MaintainWindowForViewEffectTimeRange) *MaintainWindowForView {
	s.EffectTimeRange = v
	return s
}

func (s *MaintainWindowForView) SetEffective(v string) *MaintainWindowForView {
	s.Effective = &v
	return s
}

func (s *MaintainWindowForView) SetEnable(v bool) *MaintainWindowForView {
	s.Enable = &v
	return s
}

func (s *MaintainWindowForView) SetEndTime(v string) *MaintainWindowForView {
	s.EndTime = &v
	return s
}

func (s *MaintainWindowForView) SetFilterSetting(v *FilterSetting) *MaintainWindowForView {
	s.FilterSetting = v
	return s
}

func (s *MaintainWindowForView) SetMaintainWindowId(v string) *MaintainWindowForView {
	s.MaintainWindowId = &v
	return s
}

func (s *MaintainWindowForView) SetMaintainWindowName(v string) *MaintainWindowForView {
	s.MaintainWindowName = &v
	return s
}

func (s *MaintainWindowForView) SetStartTime(v string) *MaintainWindowForView {
	s.StartTime = &v
	return s
}

func (s *MaintainWindowForView) SetUpdateTime(v string) *MaintainWindowForView {
	s.UpdateTime = &v
	return s
}

func (s *MaintainWindowForView) SetUserId(v string) *MaintainWindowForView {
	s.UserId = &v
	return s
}

func (s *MaintainWindowForView) SetWorkspace(v string) *MaintainWindowForView {
	s.Workspace = &v
	return s
}

func (s *MaintainWindowForView) Validate() error {
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

type MaintainWindowForViewEffectTimeRange struct {
	// Effective days (Monday to Sunday).
	DayInWeek []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	// End time (in minutes).
	//
	// example:
	//
	// 360
	EndTimeInMinute *int32 `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	// Start time (in minutes).
	//
	// example:
	//
	// 60
	StartTimeInMinute *int32 `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	// Time Zone.
	//
	// example:
	//
	// +08:00
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s MaintainWindowForViewEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s MaintainWindowForViewEffectTimeRange) GoString() string {
	return s.String()
}

func (s *MaintainWindowForViewEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *MaintainWindowForViewEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *MaintainWindowForViewEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *MaintainWindowForViewEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *MaintainWindowForViewEffectTimeRange) SetDayInWeek(v []*int32) *MaintainWindowForViewEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *MaintainWindowForViewEffectTimeRange) SetEndTimeInMinute(v int32) *MaintainWindowForViewEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *MaintainWindowForViewEffectTimeRange) SetStartTimeInMinute(v int32) *MaintainWindowForViewEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *MaintainWindowForViewEffectTimeRange) SetTimeZone(v string) *MaintainWindowForViewEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *MaintainWindowForViewEffectTimeRange) Validate() error {
	return dara.Validate(s)
}
