// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyStrategyForModify interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverSeconds(v int32) *NotifyStrategyForModify
	GetAutoRecoverSeconds() *int32
	SetCustomTemplateEntries(v []*NotifyStrategyForModifyCustomTemplateEntries) *NotifyStrategyForModify
	GetCustomTemplateEntries() []*NotifyStrategyForModifyCustomTemplateEntries
	SetDescription(v string) *NotifyStrategyForModify
	GetDescription() *string
	SetEnableIncidentManagement(v bool) *NotifyStrategyForModify
	GetEnableIncidentManagement() *bool
	SetEscalationId(v []*string) *NotifyStrategyForModify
	GetEscalationId() []*string
	SetFilterSetting(v *FilterSetting) *NotifyStrategyForModify
	GetFilterSetting() *FilterSetting
	SetGroupingSetting(v *NotifyStrategyForModifyGroupingSetting) *NotifyStrategyForModify
	GetGroupingSetting() *NotifyStrategyForModifyGroupingSetting
	SetIgnoreRestoredNotification(v bool) *NotifyStrategyForModify
	GetIgnoreRestoredNotification() *bool
	SetNotifyStrategyName(v string) *NotifyStrategyForModify
	GetNotifyStrategyName() *string
	SetPushingSetting(v *NotifyStrategyForModifyPushingSetting) *NotifyStrategyForModify
	GetPushingSetting() *NotifyStrategyForModifyPushingSetting
	SetRepeatNotifySetting(v *NotifyStrategyForModifyRepeatNotifySetting) *NotifyStrategyForModify
	GetRepeatNotifySetting() *NotifyStrategyForModifyRepeatNotifySetting
	SetRoutes(v []*NotifyStrategyForModifyRoutes) *NotifyStrategyForModify
	GetRoutes() []*NotifyStrategyForModifyRoutes
	SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *NotifyStrategyForModify
	GetWorkspaceFilterSetting() *WorkspaceFilterSetting
}

type NotifyStrategyForModify struct {
	AutoRecoverSeconds       *int32                                          `json:"autoRecoverSeconds,omitempty" xml:"autoRecoverSeconds,omitempty"`
	CustomTemplateEntries    []*NotifyStrategyForModifyCustomTemplateEntries `json:"customTemplateEntries,omitempty" xml:"customTemplateEntries,omitempty" type:"Repeated"`
	Description              *string                                         `json:"description,omitempty" xml:"description,omitempty"`
	EnableIncidentManagement *bool                                           `json:"enableIncidentManagement,omitempty" xml:"enableIncidentManagement,omitempty"`
	EscalationId             []*string                                       `json:"escalationId,omitempty" xml:"escalationId,omitempty" type:"Repeated"`
	FilterSetting            *FilterSetting                                  `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// This parameter is required.
	GroupingSetting            *NotifyStrategyForModifyGroupingSetting `json:"groupingSetting,omitempty" xml:"groupingSetting,omitempty" type:"Struct"`
	IgnoreRestoredNotification *bool                                   `json:"ignoreRestoredNotification,omitempty" xml:"ignoreRestoredNotification,omitempty"`
	// This parameter is required.
	NotifyStrategyName  *string                                     `json:"notifyStrategyName,omitempty" xml:"notifyStrategyName,omitempty"`
	PushingSetting      *NotifyStrategyForModifyPushingSetting      `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	RepeatNotifySetting *NotifyStrategyForModifyRepeatNotifySetting `json:"repeatNotifySetting,omitempty" xml:"repeatNotifySetting,omitempty" type:"Struct"`
	// This parameter is required.
	Routes                 []*NotifyStrategyForModifyRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
	WorkspaceFilterSetting *WorkspaceFilterSetting          `json:"workspaceFilterSetting,omitempty" xml:"workspaceFilterSetting,omitempty"`
}

func (s NotifyStrategyForModify) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForModify) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModify) GetAutoRecoverSeconds() *int32 {
	return s.AutoRecoverSeconds
}

func (s *NotifyStrategyForModify) GetCustomTemplateEntries() []*NotifyStrategyForModifyCustomTemplateEntries {
	return s.CustomTemplateEntries
}

func (s *NotifyStrategyForModify) GetDescription() *string {
	return s.Description
}

func (s *NotifyStrategyForModify) GetEnableIncidentManagement() *bool {
	return s.EnableIncidentManagement
}

func (s *NotifyStrategyForModify) GetEscalationId() []*string {
	return s.EscalationId
}

func (s *NotifyStrategyForModify) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *NotifyStrategyForModify) GetGroupingSetting() *NotifyStrategyForModifyGroupingSetting {
	return s.GroupingSetting
}

func (s *NotifyStrategyForModify) GetIgnoreRestoredNotification() *bool {
	return s.IgnoreRestoredNotification
}

func (s *NotifyStrategyForModify) GetNotifyStrategyName() *string {
	return s.NotifyStrategyName
}

func (s *NotifyStrategyForModify) GetPushingSetting() *NotifyStrategyForModifyPushingSetting {
	return s.PushingSetting
}

func (s *NotifyStrategyForModify) GetRepeatNotifySetting() *NotifyStrategyForModifyRepeatNotifySetting {
	return s.RepeatNotifySetting
}

func (s *NotifyStrategyForModify) GetRoutes() []*NotifyStrategyForModifyRoutes {
	return s.Routes
}

func (s *NotifyStrategyForModify) GetWorkspaceFilterSetting() *WorkspaceFilterSetting {
	return s.WorkspaceFilterSetting
}

func (s *NotifyStrategyForModify) SetAutoRecoverSeconds(v int32) *NotifyStrategyForModify {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *NotifyStrategyForModify) SetCustomTemplateEntries(v []*NotifyStrategyForModifyCustomTemplateEntries) *NotifyStrategyForModify {
	s.CustomTemplateEntries = v
	return s
}

func (s *NotifyStrategyForModify) SetDescription(v string) *NotifyStrategyForModify {
	s.Description = &v
	return s
}

func (s *NotifyStrategyForModify) SetEnableIncidentManagement(v bool) *NotifyStrategyForModify {
	s.EnableIncidentManagement = &v
	return s
}

func (s *NotifyStrategyForModify) SetEscalationId(v []*string) *NotifyStrategyForModify {
	s.EscalationId = v
	return s
}

func (s *NotifyStrategyForModify) SetFilterSetting(v *FilterSetting) *NotifyStrategyForModify {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyForModify) SetGroupingSetting(v *NotifyStrategyForModifyGroupingSetting) *NotifyStrategyForModify {
	s.GroupingSetting = v
	return s
}

func (s *NotifyStrategyForModify) SetIgnoreRestoredNotification(v bool) *NotifyStrategyForModify {
	s.IgnoreRestoredNotification = &v
	return s
}

func (s *NotifyStrategyForModify) SetNotifyStrategyName(v string) *NotifyStrategyForModify {
	s.NotifyStrategyName = &v
	return s
}

func (s *NotifyStrategyForModify) SetPushingSetting(v *NotifyStrategyForModifyPushingSetting) *NotifyStrategyForModify {
	s.PushingSetting = v
	return s
}

func (s *NotifyStrategyForModify) SetRepeatNotifySetting(v *NotifyStrategyForModifyRepeatNotifySetting) *NotifyStrategyForModify {
	s.RepeatNotifySetting = v
	return s
}

func (s *NotifyStrategyForModify) SetRoutes(v []*NotifyStrategyForModifyRoutes) *NotifyStrategyForModify {
	s.Routes = v
	return s
}

func (s *NotifyStrategyForModify) SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *NotifyStrategyForModify {
	s.WorkspaceFilterSetting = v
	return s
}

func (s *NotifyStrategyForModify) Validate() error {
	if s.CustomTemplateEntries != nil {
		for _, item := range s.CustomTemplateEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
			return err
		}
	}
	if s.GroupingSetting != nil {
		if err := s.GroupingSetting.Validate(); err != nil {
			return err
		}
	}
	if s.PushingSetting != nil {
		if err := s.PushingSetting.Validate(); err != nil {
			return err
		}
	}
	if s.RepeatNotifySetting != nil {
		if err := s.RepeatNotifySetting.Validate(); err != nil {
			return err
		}
	}
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WorkspaceFilterSetting != nil {
		if err := s.WorkspaceFilterSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NotifyStrategyForModifyCustomTemplateEntries struct {
	// This parameter is required.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyForModifyCustomTemplateEntries) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForModifyCustomTemplateEntries) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyCustomTemplateEntries) GetTargetType() *string {
	return s.TargetType
}

func (s *NotifyStrategyForModifyCustomTemplateEntries) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotifyStrategyForModifyCustomTemplateEntries) SetTargetType(v string) *NotifyStrategyForModifyCustomTemplateEntries {
	s.TargetType = &v
	return s
}

func (s *NotifyStrategyForModifyCustomTemplateEntries) SetTemplateUuid(v string) *NotifyStrategyForModifyCustomTemplateEntries {
	s.TemplateUuid = &v
	return s
}

func (s *NotifyStrategyForModifyCustomTemplateEntries) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForModifyGroupingSetting struct {
	GroupingKeys []*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty" type:"Repeated"`
	PeriodMin    *int32    `json:"periodMin,omitempty" xml:"periodMin,omitempty"`
	SilenceSec   *int32    `json:"silenceSec,omitempty" xml:"silenceSec,omitempty"`
	Times        *int32    `json:"times,omitempty" xml:"times,omitempty"`
}

func (s NotifyStrategyForModifyGroupingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForModifyGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyGroupingSetting) GetGroupingKeys() []*string {
	return s.GroupingKeys
}

func (s *NotifyStrategyForModifyGroupingSetting) GetPeriodMin() *int32 {
	return s.PeriodMin
}

func (s *NotifyStrategyForModifyGroupingSetting) GetSilenceSec() *int32 {
	return s.SilenceSec
}

func (s *NotifyStrategyForModifyGroupingSetting) GetTimes() *int32 {
	return s.Times
}

func (s *NotifyStrategyForModifyGroupingSetting) SetGroupingKeys(v []*string) *NotifyStrategyForModifyGroupingSetting {
	s.GroupingKeys = v
	return s
}

func (s *NotifyStrategyForModifyGroupingSetting) SetPeriodMin(v int32) *NotifyStrategyForModifyGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotifyStrategyForModifyGroupingSetting) SetSilenceSec(v int32) *NotifyStrategyForModifyGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotifyStrategyForModifyGroupingSetting) SetTimes(v int32) *NotifyStrategyForModifyGroupingSetting {
	s.Times = &v
	return s
}

func (s *NotifyStrategyForModifyGroupingSetting) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForModifyPushingSetting struct {
	AlertActionIds   []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	TemplateUuid     *string   `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyForModifyPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForModifyPushingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyPushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *NotifyStrategyForModifyPushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *NotifyStrategyForModifyPushingSetting) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotifyStrategyForModifyPushingSetting) SetAlertActionIds(v []*string) *NotifyStrategyForModifyPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *NotifyStrategyForModifyPushingSetting) SetRestoreActionIds(v []*string) *NotifyStrategyForModifyPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *NotifyStrategyForModifyPushingSetting) SetTemplateUuid(v string) *NotifyStrategyForModifyPushingSetting {
	s.TemplateUuid = &v
	return s
}

func (s *NotifyStrategyForModifyPushingSetting) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForModifyRepeatNotifySetting struct {
	EndIncidentState *string `json:"endIncidentState,omitempty" xml:"endIncidentState,omitempty"`
	RepeatInterval   *int32  `json:"repeatInterval,omitempty" xml:"repeatInterval,omitempty"`
}

func (s NotifyStrategyForModifyRepeatNotifySetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForModifyRepeatNotifySetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyRepeatNotifySetting) GetEndIncidentState() *string {
	return s.EndIncidentState
}

func (s *NotifyStrategyForModifyRepeatNotifySetting) GetRepeatInterval() *int32 {
	return s.RepeatInterval
}

func (s *NotifyStrategyForModifyRepeatNotifySetting) SetEndIncidentState(v string) *NotifyStrategyForModifyRepeatNotifySetting {
	s.EndIncidentState = &v
	return s
}

func (s *NotifyStrategyForModifyRepeatNotifySetting) SetRepeatInterval(v int32) *NotifyStrategyForModifyRepeatNotifySetting {
	s.RepeatInterval = &v
	return s
}

func (s *NotifyStrategyForModifyRepeatNotifySetting) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForModifyRoutes struct {
	Channels        []*NotifyStrategyForModifyRoutesChannels      `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	EffectTimeRange *NotifyStrategyForModifyRoutesEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	FilterSetting   *FilterSetting                                `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	Severities      []*string                                     `json:"severities,omitempty" xml:"severities,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForModifyRoutes) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForModifyRoutes) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyRoutes) GetChannels() []*NotifyStrategyForModifyRoutesChannels {
	return s.Channels
}

func (s *NotifyStrategyForModifyRoutes) GetEffectTimeRange() *NotifyStrategyForModifyRoutesEffectTimeRange {
	return s.EffectTimeRange
}

func (s *NotifyStrategyForModifyRoutes) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *NotifyStrategyForModifyRoutes) GetSeverities() []*string {
	return s.Severities
}

func (s *NotifyStrategyForModifyRoutes) SetChannels(v []*NotifyStrategyForModifyRoutesChannels) *NotifyStrategyForModifyRoutes {
	s.Channels = v
	return s
}

func (s *NotifyStrategyForModifyRoutes) SetEffectTimeRange(v *NotifyStrategyForModifyRoutesEffectTimeRange) *NotifyStrategyForModifyRoutes {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyStrategyForModifyRoutes) SetFilterSetting(v *FilterSetting) *NotifyStrategyForModifyRoutes {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyForModifyRoutes) SetSeverities(v []*string) *NotifyStrategyForModifyRoutes {
	s.Severities = v
	return s
}

func (s *NotifyStrategyForModifyRoutes) Validate() error {
	if s.Channels != nil {
		for _, item := range s.Channels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type NotifyStrategyForModifyRoutesChannels struct {
	// This parameter is required.
	ChannelType        *string   `json:"channelType,omitempty" xml:"channelType,omitempty"`
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	// This parameter is required.
	Receivers []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForModifyRoutesChannels) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForModifyRoutesChannels) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyRoutesChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotifyStrategyForModifyRoutesChannels) GetEnabledSubChannels() []*string {
	return s.EnabledSubChannels
}

func (s *NotifyStrategyForModifyRoutesChannels) GetReceivers() []*string {
	return s.Receivers
}

func (s *NotifyStrategyForModifyRoutesChannels) SetChannelType(v string) *NotifyStrategyForModifyRoutesChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyStrategyForModifyRoutesChannels) SetEnabledSubChannels(v []*string) *NotifyStrategyForModifyRoutesChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyStrategyForModifyRoutesChannels) SetReceivers(v []*string) *NotifyStrategyForModifyRoutesChannels {
	s.Receivers = v
	return s
}

func (s *NotifyStrategyForModifyRoutesChannels) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForModifyRoutesEffectTimeRange struct {
	DayInWeek         []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	EndTimeInMinute   *int32   `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	StartTimeInMinute *int32   `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	TimeZone          *string  `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s NotifyStrategyForModifyRoutesEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForModifyRoutesEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) SetDayInWeek(v []*int32) *NotifyStrategyForModifyRoutesEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyStrategyForModifyRoutesEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyStrategyForModifyRoutesEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) SetTimeZone(v string) *NotifyStrategyForModifyRoutesEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *NotifyStrategyForModifyRoutesEffectTimeRange) Validate() error {
	return dara.Validate(s)
}
