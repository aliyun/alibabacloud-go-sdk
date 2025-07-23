// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyStrategyForView interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *NotifyStrategyForView
	GetCreateTime() *string
	SetCustomTemplateEntries(v []*NotifyStrategyForViewCustomTemplateEntries) *NotifyStrategyForView
	GetCustomTemplateEntries() []*NotifyStrategyForViewCustomTemplateEntries
	SetDescription(v string) *NotifyStrategyForView
	GetDescription() *string
	SetEnable(v bool) *NotifyStrategyForView
	GetEnable() *bool
	SetGroupingSetting(v *NotifyStrategyForViewGroupingSetting) *NotifyStrategyForView
	GetGroupingSetting() *NotifyStrategyForViewGroupingSetting
	SetIgnoreRestoredNotification(v bool) *NotifyStrategyForView
	GetIgnoreRestoredNotification() *bool
	SetNotifyStrategyId(v string) *NotifyStrategyForView
	GetNotifyStrategyId() *string
	SetNotifyStrategyName(v string) *NotifyStrategyForView
	GetNotifyStrategyName() *string
	SetRoutes(v []*NotifyStrategyForViewRoutes) *NotifyStrategyForView
	GetRoutes() []*NotifyStrategyForViewRoutes
	SetUpdateTime(v string) *NotifyStrategyForView
	GetUpdateTime() *string
	SetUserId(v string) *NotifyStrategyForView
	GetUserId() *string
	SetWorkspace(v string) *NotifyStrategyForView
	GetWorkspace() *string
}

type NotifyStrategyForView struct {
	CreateTime            *string                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CustomTemplateEntries []*NotifyStrategyForViewCustomTemplateEntries `json:"customTemplateEntries,omitempty" xml:"customTemplateEntries,omitempty" type:"Repeated"`
	Description           *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	Enable                *bool                                         `json:"enable,omitempty" xml:"enable,omitempty"`
	// This parameter is required.
	GroupingSetting            *NotifyStrategyForViewGroupingSetting `json:"groupingSetting,omitempty" xml:"groupingSetting,omitempty" type:"Struct"`
	IgnoreRestoredNotification *bool                                 `json:"ignoreRestoredNotification,omitempty" xml:"ignoreRestoredNotification,omitempty"`
	NotifyStrategyId           *string                               `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// This parameter is required.
	NotifyStrategyName *string `json:"notifyStrategyName,omitempty" xml:"notifyStrategyName,omitempty"`
	// This parameter is required.
	Routes     []*NotifyStrategyForViewRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
	UpdateTime *string                        `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId     *string                        `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace  *string                        `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s NotifyStrategyForView) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForView) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *NotifyStrategyForView) GetCustomTemplateEntries() []*NotifyStrategyForViewCustomTemplateEntries {
	return s.CustomTemplateEntries
}

func (s *NotifyStrategyForView) GetDescription() *string {
	return s.Description
}

func (s *NotifyStrategyForView) GetEnable() *bool {
	return s.Enable
}

func (s *NotifyStrategyForView) GetGroupingSetting() *NotifyStrategyForViewGroupingSetting {
	return s.GroupingSetting
}

func (s *NotifyStrategyForView) GetIgnoreRestoredNotification() *bool {
	return s.IgnoreRestoredNotification
}

func (s *NotifyStrategyForView) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *NotifyStrategyForView) GetNotifyStrategyName() *string {
	return s.NotifyStrategyName
}

func (s *NotifyStrategyForView) GetRoutes() []*NotifyStrategyForViewRoutes {
	return s.Routes
}

func (s *NotifyStrategyForView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *NotifyStrategyForView) GetUserId() *string {
	return s.UserId
}

func (s *NotifyStrategyForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *NotifyStrategyForView) SetCreateTime(v string) *NotifyStrategyForView {
	s.CreateTime = &v
	return s
}

func (s *NotifyStrategyForView) SetCustomTemplateEntries(v []*NotifyStrategyForViewCustomTemplateEntries) *NotifyStrategyForView {
	s.CustomTemplateEntries = v
	return s
}

func (s *NotifyStrategyForView) SetDescription(v string) *NotifyStrategyForView {
	s.Description = &v
	return s
}

func (s *NotifyStrategyForView) SetEnable(v bool) *NotifyStrategyForView {
	s.Enable = &v
	return s
}

func (s *NotifyStrategyForView) SetGroupingSetting(v *NotifyStrategyForViewGroupingSetting) *NotifyStrategyForView {
	s.GroupingSetting = v
	return s
}

func (s *NotifyStrategyForView) SetIgnoreRestoredNotification(v bool) *NotifyStrategyForView {
	s.IgnoreRestoredNotification = &v
	return s
}

func (s *NotifyStrategyForView) SetNotifyStrategyId(v string) *NotifyStrategyForView {
	s.NotifyStrategyId = &v
	return s
}

func (s *NotifyStrategyForView) SetNotifyStrategyName(v string) *NotifyStrategyForView {
	s.NotifyStrategyName = &v
	return s
}

func (s *NotifyStrategyForView) SetRoutes(v []*NotifyStrategyForViewRoutes) *NotifyStrategyForView {
	s.Routes = v
	return s
}

func (s *NotifyStrategyForView) SetUpdateTime(v string) *NotifyStrategyForView {
	s.UpdateTime = &v
	return s
}

func (s *NotifyStrategyForView) SetUserId(v string) *NotifyStrategyForView {
	s.UserId = &v
	return s
}

func (s *NotifyStrategyForView) SetWorkspace(v string) *NotifyStrategyForView {
	s.Workspace = &v
	return s
}

func (s *NotifyStrategyForView) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForViewCustomTemplateEntries struct {
	// This parameter is required.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyForViewCustomTemplateEntries) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForViewCustomTemplateEntries) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewCustomTemplateEntries) GetTargetType() *string {
	return s.TargetType
}

func (s *NotifyStrategyForViewCustomTemplateEntries) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotifyStrategyForViewCustomTemplateEntries) SetTargetType(v string) *NotifyStrategyForViewCustomTemplateEntries {
	s.TargetType = &v
	return s
}

func (s *NotifyStrategyForViewCustomTemplateEntries) SetTemplateUuid(v string) *NotifyStrategyForViewCustomTemplateEntries {
	s.TemplateUuid = &v
	return s
}

func (s *NotifyStrategyForViewCustomTemplateEntries) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForViewGroupingSetting struct {
	GroupingKeys []*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty" type:"Repeated"`
	PeriodMin    *int32    `json:"periodMin,omitempty" xml:"periodMin,omitempty"`
	SilenceSec   *int32    `json:"silenceSec,omitempty" xml:"silenceSec,omitempty"`
	Times        *int32    `json:"times,omitempty" xml:"times,omitempty"`
}

func (s NotifyStrategyForViewGroupingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForViewGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewGroupingSetting) GetGroupingKeys() []*string {
	return s.GroupingKeys
}

func (s *NotifyStrategyForViewGroupingSetting) GetPeriodMin() *int32 {
	return s.PeriodMin
}

func (s *NotifyStrategyForViewGroupingSetting) GetSilenceSec() *int32 {
	return s.SilenceSec
}

func (s *NotifyStrategyForViewGroupingSetting) GetTimes() *int32 {
	return s.Times
}

func (s *NotifyStrategyForViewGroupingSetting) SetGroupingKeys(v []*string) *NotifyStrategyForViewGroupingSetting {
	s.GroupingKeys = v
	return s
}

func (s *NotifyStrategyForViewGroupingSetting) SetPeriodMin(v int32) *NotifyStrategyForViewGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotifyStrategyForViewGroupingSetting) SetSilenceSec(v int32) *NotifyStrategyForViewGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotifyStrategyForViewGroupingSetting) SetTimes(v int32) *NotifyStrategyForViewGroupingSetting {
	s.Times = &v
	return s
}

func (s *NotifyStrategyForViewGroupingSetting) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForViewRoutes struct {
	Channels        []*NotifyStrategyForViewRoutesChannels      `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	EffectTimeRange *NotifyStrategyForViewRoutesEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	FilterSetting   *FilterSetting                              `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	Severities      []*string                                   `json:"severities,omitempty" xml:"severities,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForViewRoutes) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForViewRoutes) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewRoutes) GetChannels() []*NotifyStrategyForViewRoutesChannels {
	return s.Channels
}

func (s *NotifyStrategyForViewRoutes) GetEffectTimeRange() *NotifyStrategyForViewRoutesEffectTimeRange {
	return s.EffectTimeRange
}

func (s *NotifyStrategyForViewRoutes) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *NotifyStrategyForViewRoutes) GetSeverities() []*string {
	return s.Severities
}

func (s *NotifyStrategyForViewRoutes) SetChannels(v []*NotifyStrategyForViewRoutesChannels) *NotifyStrategyForViewRoutes {
	s.Channels = v
	return s
}

func (s *NotifyStrategyForViewRoutes) SetEffectTimeRange(v *NotifyStrategyForViewRoutesEffectTimeRange) *NotifyStrategyForViewRoutes {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyStrategyForViewRoutes) SetFilterSetting(v *FilterSetting) *NotifyStrategyForViewRoutes {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyForViewRoutes) SetSeverities(v []*string) *NotifyStrategyForViewRoutes {
	s.Severities = v
	return s
}

func (s *NotifyStrategyForViewRoutes) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForViewRoutesChannels struct {
	// This parameter is required.
	ChannelType        *string   `json:"channelType,omitempty" xml:"channelType,omitempty"`
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	// This parameter is required.
	Receivers []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForViewRoutesChannels) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForViewRoutesChannels) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewRoutesChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotifyStrategyForViewRoutesChannels) GetEnabledSubChannels() []*string {
	return s.EnabledSubChannels
}

func (s *NotifyStrategyForViewRoutesChannels) GetReceivers() []*string {
	return s.Receivers
}

func (s *NotifyStrategyForViewRoutesChannels) SetChannelType(v string) *NotifyStrategyForViewRoutesChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyStrategyForViewRoutesChannels) SetEnabledSubChannels(v []*string) *NotifyStrategyForViewRoutesChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyStrategyForViewRoutesChannels) SetReceivers(v []*string) *NotifyStrategyForViewRoutesChannels {
	s.Receivers = v
	return s
}

func (s *NotifyStrategyForViewRoutesChannels) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForViewRoutesEffectTimeRange struct {
	DayInWeek         []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	EndTimeInMinute   *int32   `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	StartTimeInMinute *int32   `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s NotifyStrategyForViewRoutesEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForViewRoutesEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) SetDayInWeek(v []*int32) *NotifyStrategyForViewRoutesEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyStrategyForViewRoutesEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyStrategyForViewRoutesEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) SetTimeZone(v string) *NotifyStrategyForViewRoutesEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *NotifyStrategyForViewRoutesEffectTimeRange) Validate() error {
	return dara.Validate(s)
}
