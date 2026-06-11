// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyStrategyForSNSView interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *NotifyStrategyForSNSView
	GetCreateTime() *string
	SetCustomTemplateEntries(v []*NotifyStrategyForSNSViewCustomTemplateEntries) *NotifyStrategyForSNSView
	GetCustomTemplateEntries() []*NotifyStrategyForSNSViewCustomTemplateEntries
	SetDescription(v string) *NotifyStrategyForSNSView
	GetDescription() *string
	SetEnable(v bool) *NotifyStrategyForSNSView
	GetEnable() *bool
	SetEnableIncidentManagement(v bool) *NotifyStrategyForSNSView
	GetEnableIncidentManagement() *bool
	SetGroupingSetting(v *NotifyStrategyForSNSViewGroupingSetting) *NotifyStrategyForSNSView
	GetGroupingSetting() *NotifyStrategyForSNSViewGroupingSetting
	SetIgnoreRestoredNotification(v bool) *NotifyStrategyForSNSView
	GetIgnoreRestoredNotification() *bool
	SetIncidentResponsePlanId(v string) *NotifyStrategyForSNSView
	GetIncidentResponsePlanId() *string
	SetMode(v string) *NotifyStrategyForSNSView
	GetMode() *string
	SetNotifyStrategyId(v string) *NotifyStrategyForSNSView
	GetNotifyStrategyId() *string
	SetNotifyStrategyName(v string) *NotifyStrategyForSNSView
	GetNotifyStrategyName() *string
	SetRoutes(v []*NotifyStrategyForSNSViewRoutes) *NotifyStrategyForSNSView
	GetRoutes() []*NotifyStrategyForSNSViewRoutes
	SetSyncFromType(v string) *NotifyStrategyForSNSView
	GetSyncFromType() *string
	SetUpdateTime(v string) *NotifyStrategyForSNSView
	GetUpdateTime() *string
	SetUserId(v string) *NotifyStrategyForSNSView
	GetUserId() *string
	SetWorkspace(v string) *NotifyStrategyForSNSView
	GetWorkspace() *string
}

type NotifyStrategyForSNSView struct {
	// The creation time of the notification strategy.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The list of custom templates.
	CustomTemplateEntries []*NotifyStrategyForSNSViewCustomTemplateEntries `json:"customTemplateEntries,omitempty" xml:"customTemplateEntries,omitempty" type:"Repeated"`
	// The description of the notification strategy.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable the notification strategy. Valid values: true, false.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Specifies whether to enable incident management. Valid values: true, false.
	EnableIncidentManagement *bool `json:"enableIncidentManagement,omitempty" xml:"enableIncidentManagement,omitempty"`
	// The settings for alert grouping.
	GroupingSetting *NotifyStrategyForSNSViewGroupingSetting `json:"groupingSetting,omitempty" xml:"groupingSetting,omitempty" type:"Struct"`
	// Specifies whether to ignore notifications for restored alerts. Valid values: true, false.
	IgnoreRestoredNotification *bool `json:"ignoreRestoredNotification,omitempty" xml:"ignoreRestoredNotification,omitempty"`
	// The ID of the incident response plan.
	IncidentResponsePlanId *string `json:"incidentResponsePlanId,omitempty" xml:"incidentResponsePlanId,omitempty"`
	// The mode of the notification strategy.
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The ID of the notification strategy.
	NotifyStrategyId *string `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// The name of the notification strategy.
	NotifyStrategyName *string `json:"notifyStrategyName,omitempty" xml:"notifyStrategyName,omitempty"`
	// The list of notification routes.
	Routes []*NotifyStrategyForSNSViewRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
	// The source from which the strategy is synchronized.
	SyncFromType *string `json:"syncFromType,omitempty" xml:"syncFromType,omitempty"`
	// The last update time of the notification strategy.
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The user ID.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The workspace to which the notification strategy belongs.
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s NotifyStrategyForSNSView) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSView) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *NotifyStrategyForSNSView) GetCustomTemplateEntries() []*NotifyStrategyForSNSViewCustomTemplateEntries {
	return s.CustomTemplateEntries
}

func (s *NotifyStrategyForSNSView) GetDescription() *string {
	return s.Description
}

func (s *NotifyStrategyForSNSView) GetEnable() *bool {
	return s.Enable
}

func (s *NotifyStrategyForSNSView) GetEnableIncidentManagement() *bool {
	return s.EnableIncidentManagement
}

func (s *NotifyStrategyForSNSView) GetGroupingSetting() *NotifyStrategyForSNSViewGroupingSetting {
	return s.GroupingSetting
}

func (s *NotifyStrategyForSNSView) GetIgnoreRestoredNotification() *bool {
	return s.IgnoreRestoredNotification
}

func (s *NotifyStrategyForSNSView) GetIncidentResponsePlanId() *string {
	return s.IncidentResponsePlanId
}

func (s *NotifyStrategyForSNSView) GetMode() *string {
	return s.Mode
}

func (s *NotifyStrategyForSNSView) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *NotifyStrategyForSNSView) GetNotifyStrategyName() *string {
	return s.NotifyStrategyName
}

func (s *NotifyStrategyForSNSView) GetRoutes() []*NotifyStrategyForSNSViewRoutes {
	return s.Routes
}

func (s *NotifyStrategyForSNSView) GetSyncFromType() *string {
	return s.SyncFromType
}

func (s *NotifyStrategyForSNSView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *NotifyStrategyForSNSView) GetUserId() *string {
	return s.UserId
}

func (s *NotifyStrategyForSNSView) GetWorkspace() *string {
	return s.Workspace
}

func (s *NotifyStrategyForSNSView) SetCreateTime(v string) *NotifyStrategyForSNSView {
	s.CreateTime = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetCustomTemplateEntries(v []*NotifyStrategyForSNSViewCustomTemplateEntries) *NotifyStrategyForSNSView {
	s.CustomTemplateEntries = v
	return s
}

func (s *NotifyStrategyForSNSView) SetDescription(v string) *NotifyStrategyForSNSView {
	s.Description = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetEnable(v bool) *NotifyStrategyForSNSView {
	s.Enable = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetEnableIncidentManagement(v bool) *NotifyStrategyForSNSView {
	s.EnableIncidentManagement = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetGroupingSetting(v *NotifyStrategyForSNSViewGroupingSetting) *NotifyStrategyForSNSView {
	s.GroupingSetting = v
	return s
}

func (s *NotifyStrategyForSNSView) SetIgnoreRestoredNotification(v bool) *NotifyStrategyForSNSView {
	s.IgnoreRestoredNotification = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetIncidentResponsePlanId(v string) *NotifyStrategyForSNSView {
	s.IncidentResponsePlanId = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetMode(v string) *NotifyStrategyForSNSView {
	s.Mode = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetNotifyStrategyId(v string) *NotifyStrategyForSNSView {
	s.NotifyStrategyId = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetNotifyStrategyName(v string) *NotifyStrategyForSNSView {
	s.NotifyStrategyName = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetRoutes(v []*NotifyStrategyForSNSViewRoutes) *NotifyStrategyForSNSView {
	s.Routes = v
	return s
}

func (s *NotifyStrategyForSNSView) SetSyncFromType(v string) *NotifyStrategyForSNSView {
	s.SyncFromType = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetUpdateTime(v string) *NotifyStrategyForSNSView {
	s.UpdateTime = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetUserId(v string) *NotifyStrategyForSNSView {
	s.UserId = &v
	return s
}

func (s *NotifyStrategyForSNSView) SetWorkspace(v string) *NotifyStrategyForSNSView {
	s.Workspace = &v
	return s
}

func (s *NotifyStrategyForSNSView) Validate() error {
	if s.CustomTemplateEntries != nil {
		for _, item := range s.CustomTemplateEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GroupingSetting != nil {
		if err := s.GroupingSetting.Validate(); err != nil {
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
	return nil
}

type NotifyStrategyForSNSViewCustomTemplateEntries struct {
	// The target type for the custom template.
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The unique identifier (UUID) of the template.
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyForSNSViewCustomTemplateEntries) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSViewCustomTemplateEntries) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSViewCustomTemplateEntries) GetTargetType() *string {
	return s.TargetType
}

func (s *NotifyStrategyForSNSViewCustomTemplateEntries) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotifyStrategyForSNSViewCustomTemplateEntries) SetTargetType(v string) *NotifyStrategyForSNSViewCustomTemplateEntries {
	s.TargetType = &v
	return s
}

func (s *NotifyStrategyForSNSViewCustomTemplateEntries) SetTemplateUuid(v string) *NotifyStrategyForSNSViewCustomTemplateEntries {
	s.TemplateUuid = &v
	return s
}

func (s *NotifyStrategyForSNSViewCustomTemplateEntries) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForSNSViewGroupingSetting struct {
	// The keys for grouping alerts.
	GroupingKeys []*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty" type:"Repeated"`
	// The time window in minutes for grouping alerts.
	PeriodMin *int32 `json:"periodMin,omitempty" xml:"periodMin,omitempty"`
	// The silence period in seconds after a notification is sent for a group.
	SilenceSec *int32 `json:"silenceSec,omitempty" xml:"silenceSec,omitempty"`
	// The number of times to send notifications for a group.
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s NotifyStrategyForSNSViewGroupingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSViewGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSViewGroupingSetting) GetGroupingKeys() []*string {
	return s.GroupingKeys
}

func (s *NotifyStrategyForSNSViewGroupingSetting) GetPeriodMin() *int32 {
	return s.PeriodMin
}

func (s *NotifyStrategyForSNSViewGroupingSetting) GetSilenceSec() *int32 {
	return s.SilenceSec
}

func (s *NotifyStrategyForSNSViewGroupingSetting) GetTimes() *int32 {
	return s.Times
}

func (s *NotifyStrategyForSNSViewGroupingSetting) SetGroupingKeys(v []*string) *NotifyStrategyForSNSViewGroupingSetting {
	s.GroupingKeys = v
	return s
}

func (s *NotifyStrategyForSNSViewGroupingSetting) SetPeriodMin(v int32) *NotifyStrategyForSNSViewGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotifyStrategyForSNSViewGroupingSetting) SetSilenceSec(v int32) *NotifyStrategyForSNSViewGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotifyStrategyForSNSViewGroupingSetting) SetTimes(v int32) *NotifyStrategyForSNSViewGroupingSetting {
	s.Times = &v
	return s
}

func (s *NotifyStrategyForSNSViewGroupingSetting) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForSNSViewRoutes struct {
	// The notification channels for the route.
	Channels []*NotifyStrategyForSNSViewRoutesChannels `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	// The name of the digital employee assigned to this route.
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The time range during which the notification route is active.
	EffectTimeRange *NotifyStrategyForSNSViewRoutesEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	// Specifies whether to enable root cause analysis (RCA) for alerts that match this route. Valid values: true, false.
	EnableRca *bool `json:"enableRca,omitempty" xml:"enableRca,omitempty"`
	// The filter settings for the route.
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// The alert severities that trigger this route.
	Severities []*string `json:"severities,omitempty" xml:"severities,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForSNSViewRoutes) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSViewRoutes) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSViewRoutes) GetChannels() []*NotifyStrategyForSNSViewRoutesChannels {
	return s.Channels
}

func (s *NotifyStrategyForSNSViewRoutes) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *NotifyStrategyForSNSViewRoutes) GetEffectTimeRange() *NotifyStrategyForSNSViewRoutesEffectTimeRange {
	return s.EffectTimeRange
}

func (s *NotifyStrategyForSNSViewRoutes) GetEnableRca() *bool {
	return s.EnableRca
}

func (s *NotifyStrategyForSNSViewRoutes) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *NotifyStrategyForSNSViewRoutes) GetSeverities() []*string {
	return s.Severities
}

func (s *NotifyStrategyForSNSViewRoutes) SetChannels(v []*NotifyStrategyForSNSViewRoutesChannels) *NotifyStrategyForSNSViewRoutes {
	s.Channels = v
	return s
}

func (s *NotifyStrategyForSNSViewRoutes) SetDigitalEmployeeName(v string) *NotifyStrategyForSNSViewRoutes {
	s.DigitalEmployeeName = &v
	return s
}

func (s *NotifyStrategyForSNSViewRoutes) SetEffectTimeRange(v *NotifyStrategyForSNSViewRoutesEffectTimeRange) *NotifyStrategyForSNSViewRoutes {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyStrategyForSNSViewRoutes) SetEnableRca(v bool) *NotifyStrategyForSNSViewRoutes {
	s.EnableRca = &v
	return s
}

func (s *NotifyStrategyForSNSViewRoutes) SetFilterSetting(v *FilterSetting) *NotifyStrategyForSNSViewRoutes {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyForSNSViewRoutes) SetSeverities(v []*string) *NotifyStrategyForSNSViewRoutes {
	s.Severities = v
	return s
}

func (s *NotifyStrategyForSNSViewRoutes) Validate() error {
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

type NotifyStrategyForSNSViewRoutesChannels struct {
	// The type of the notification channel, such as \\"sms\\" or \\"email\\".
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The enabled sub-channels.
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	// The list of receivers for the channel.
	Receivers []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForSNSViewRoutesChannels) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSViewRoutesChannels) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSViewRoutesChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotifyStrategyForSNSViewRoutesChannels) GetEnabledSubChannels() []*string {
	return s.EnabledSubChannels
}

func (s *NotifyStrategyForSNSViewRoutesChannels) GetReceivers() []*string {
	return s.Receivers
}

func (s *NotifyStrategyForSNSViewRoutesChannels) SetChannelType(v string) *NotifyStrategyForSNSViewRoutesChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyStrategyForSNSViewRoutesChannels) SetEnabledSubChannels(v []*string) *NotifyStrategyForSNSViewRoutesChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyStrategyForSNSViewRoutesChannels) SetReceivers(v []*string) *NotifyStrategyForSNSViewRoutesChannels {
	s.Receivers = v
	return s
}

func (s *NotifyStrategyForSNSViewRoutesChannels) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForSNSViewRoutesEffectTimeRange struct {
	// The days of the week when the route is active.
	DayInWeek []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	// The end time of the active period, specified in minutes from 00:00.
	EndTimeInMinute *int32 `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	// The start time of the active period, specified in minutes from 00:00.
	StartTimeInMinute *int32 `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	// The time zone for the active period. For example, \\"Asia/Shanghai\\".
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s NotifyStrategyForSNSViewRoutesEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSViewRoutesEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) SetDayInWeek(v []*int32) *NotifyStrategyForSNSViewRoutesEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyStrategyForSNSViewRoutesEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyStrategyForSNSViewRoutesEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) SetTimeZone(v string) *NotifyStrategyForSNSViewRoutesEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *NotifyStrategyForSNSViewRoutesEffectTimeRange) Validate() error {
	return dara.Validate(s)
}
