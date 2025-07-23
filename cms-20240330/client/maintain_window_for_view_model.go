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
	CreateTime       *string                               `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description      *string                               `json:"description,omitempty" xml:"description,omitempty"`
	EffectTimeRange  *MaintainWindowForViewEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	Effective        *string                               `json:"effective,omitempty" xml:"effective,omitempty"`
	Enable           *bool                                 `json:"enable,omitempty" xml:"enable,omitempty"`
	EndTime          *string                               `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FilterSetting    *FilterSetting                        `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	MaintainWindowId *string                               `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// This parameter is required.
	MaintainWindowName *string `json:"maintainWindowName,omitempty" xml:"maintainWindowName,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UpdateTime         *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace          *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	return dara.Validate(s)
}

type MaintainWindowForViewEffectTimeRange struct {
	DayInWeek         []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	EndTimeInMinute   *int32   `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	StartTimeInMinute *int32   `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
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
