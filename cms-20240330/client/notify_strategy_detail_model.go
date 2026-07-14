// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyStrategyDetail interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplateEntries(v []*NotifyStrategyDetailCustomTemplateEntries) *NotifyStrategyDetail
	GetCustomTemplateEntries() []*NotifyStrategyDetailCustomTemplateEntries
	SetDescription(v string) *NotifyStrategyDetail
	GetDescription() *string
	SetGroupingSetting(v *NotifyStrategyDetailGroupingSetting) *NotifyStrategyDetail
	GetGroupingSetting() *NotifyStrategyDetailGroupingSetting
	SetIgnoreRestoredNotification(v bool) *NotifyStrategyDetail
	GetIgnoreRestoredNotification() *bool
	SetRoutes(v []*NotifyStrategyDetailRoutes) *NotifyStrategyDetail
	GetRoutes() []*NotifyStrategyDetailRoutes
}

type NotifyStrategyDetail struct {
	// The list of custom notification templates.
	//
	// example:
	//
	// []
	CustomTemplateEntries []*NotifyStrategyDetailCustomTemplateEntries `json:"customTemplateEntries,omitempty" xml:"customTemplateEntries,omitempty" type:"Repeated"`
	// The description of the notification policy.
	//
	// example:
	//
	// 生产环境告警通知策略
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The grouping and merging settings.
	GroupingSetting *NotifyStrategyDetailGroupingSetting `json:"groupingSetting,omitempty" xml:"groupingSetting,omitempty" type:"Struct"`
	// Specifies whether to ignore notifications for recovery events. A value of true indicates that recovery notifications are not sent.
	//
	// example:
	//
	// true
	IgnoreRestoredNotification *bool `json:"ignoreRestoredNotification,omitempty" xml:"ignoreRestoredNotification,omitempty"`
	// The list of notification channel routing settings.
	//
	// example:
	//
	// []
	Routes []*NotifyStrategyDetailRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s NotifyStrategyDetail) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyDetail) GoString() string {
	return s.String()
}

func (s *NotifyStrategyDetail) GetCustomTemplateEntries() []*NotifyStrategyDetailCustomTemplateEntries {
	return s.CustomTemplateEntries
}

func (s *NotifyStrategyDetail) GetDescription() *string {
	return s.Description
}

func (s *NotifyStrategyDetail) GetGroupingSetting() *NotifyStrategyDetailGroupingSetting {
	return s.GroupingSetting
}

func (s *NotifyStrategyDetail) GetIgnoreRestoredNotification() *bool {
	return s.IgnoreRestoredNotification
}

func (s *NotifyStrategyDetail) GetRoutes() []*NotifyStrategyDetailRoutes {
	return s.Routes
}

func (s *NotifyStrategyDetail) SetCustomTemplateEntries(v []*NotifyStrategyDetailCustomTemplateEntries) *NotifyStrategyDetail {
	s.CustomTemplateEntries = v
	return s
}

func (s *NotifyStrategyDetail) SetDescription(v string) *NotifyStrategyDetail {
	s.Description = &v
	return s
}

func (s *NotifyStrategyDetail) SetGroupingSetting(v *NotifyStrategyDetailGroupingSetting) *NotifyStrategyDetail {
	s.GroupingSetting = v
	return s
}

func (s *NotifyStrategyDetail) SetIgnoreRestoredNotification(v bool) *NotifyStrategyDetail {
	s.IgnoreRestoredNotification = &v
	return s
}

func (s *NotifyStrategyDetail) SetRoutes(v []*NotifyStrategyDetailRoutes) *NotifyStrategyDetail {
	s.Routes = v
	return s
}

func (s *NotifyStrategyDetail) Validate() error {
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

type NotifyStrategyDetailCustomTemplateEntries struct {
	// The UUID of the notification template.
	//
	// example:
	//
	// template-uuid-xxx
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyDetailCustomTemplateEntries) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyDetailCustomTemplateEntries) GoString() string {
	return s.String()
}

func (s *NotifyStrategyDetailCustomTemplateEntries) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotifyStrategyDetailCustomTemplateEntries) SetTemplateUuid(v string) *NotifyStrategyDetailCustomTemplateEntries {
	s.TemplateUuid = &v
	return s
}

func (s *NotifyStrategyDetailCustomTemplateEntries) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyDetailGroupingSetting struct {
	// The event fields by which events are grouped for noise reduction. Events in the same group are merged into a single notification. An empty array indicates no grouping.
	//
	// example:
	//
	// ["alertname"]
	GroupingKeys []*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty" type:"Repeated"`
	// The check period in minutes.
	//
	// example:
	//
	// 0
	PeriodMin *int32 `json:"periodMin,omitempty" xml:"periodMin,omitempty"`
	// The silence duration in seconds.
	//
	// example:
	//
	// 0
	SilenceSec *int32 `json:"silenceSec,omitempty" xml:"silenceSec,omitempty"`
	// The number of trigger times.
	//
	// example:
	//
	// 0
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s NotifyStrategyDetailGroupingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyDetailGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyDetailGroupingSetting) GetGroupingKeys() []*string {
	return s.GroupingKeys
}

func (s *NotifyStrategyDetailGroupingSetting) GetPeriodMin() *int32 {
	return s.PeriodMin
}

func (s *NotifyStrategyDetailGroupingSetting) GetSilenceSec() *int32 {
	return s.SilenceSec
}

func (s *NotifyStrategyDetailGroupingSetting) GetTimes() *int32 {
	return s.Times
}

func (s *NotifyStrategyDetailGroupingSetting) SetGroupingKeys(v []*string) *NotifyStrategyDetailGroupingSetting {
	s.GroupingKeys = v
	return s
}

func (s *NotifyStrategyDetailGroupingSetting) SetPeriodMin(v int32) *NotifyStrategyDetailGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotifyStrategyDetailGroupingSetting) SetSilenceSec(v int32) *NotifyStrategyDetailGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotifyStrategyDetailGroupingSetting) SetTimes(v int32) *NotifyStrategyDetailGroupingSetting {
	s.Times = &v
	return s
}

func (s *NotifyStrategyDetailGroupingSetting) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyDetailRoutes struct {
	// The list of notification channels.
	//
	// example:
	//
	// []
	Channels []*NotifyStrategyDetailRoutesChannels `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	// The digital employee name. This parameter is required when enableRca is set to true.
	//
	// example:
	//
	// 数字员工名
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The effective time range.
	EffectTimeRange *NotifyStrategyDetailRoutesEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	// Specifies whether to enable root cause analysis (RCA).
	//
	// example:
	//
	// false
	EnableRca *bool `json:"enableRca,omitempty" xml:"enableRca,omitempty"`
	// The route-level event filter conditions.
	//
	// example:
	//
	// {}
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
}

func (s NotifyStrategyDetailRoutes) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyDetailRoutes) GoString() string {
	return s.String()
}

func (s *NotifyStrategyDetailRoutes) GetChannels() []*NotifyStrategyDetailRoutesChannels {
	return s.Channels
}

func (s *NotifyStrategyDetailRoutes) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *NotifyStrategyDetailRoutes) GetEffectTimeRange() *NotifyStrategyDetailRoutesEffectTimeRange {
	return s.EffectTimeRange
}

func (s *NotifyStrategyDetailRoutes) GetEnableRca() *bool {
	return s.EnableRca
}

func (s *NotifyStrategyDetailRoutes) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *NotifyStrategyDetailRoutes) SetChannels(v []*NotifyStrategyDetailRoutesChannels) *NotifyStrategyDetailRoutes {
	s.Channels = v
	return s
}

func (s *NotifyStrategyDetailRoutes) SetDigitalEmployeeName(v string) *NotifyStrategyDetailRoutes {
	s.DigitalEmployeeName = &v
	return s
}

func (s *NotifyStrategyDetailRoutes) SetEffectTimeRange(v *NotifyStrategyDetailRoutesEffectTimeRange) *NotifyStrategyDetailRoutes {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyStrategyDetailRoutes) SetEnableRca(v bool) *NotifyStrategyDetailRoutes {
	s.EnableRca = &v
	return s
}

func (s *NotifyStrategyDetailRoutes) SetFilterSetting(v *FilterSetting) *NotifyStrategyDetailRoutes {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyDetailRoutes) Validate() error {
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

type NotifyStrategyDetailRoutesChannels struct {
	// The channel type. Valid values: DING, WEIXIN, FEISHU, SLACK, TEAMS, WEBHOOK, CONTACT, GROUP, DUTY, and DING_COOL_APP.
	//
	// example:
	//
	// WEBHOOK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The enabled sub-channels. Valid values: EMAIL, SMS, VOICE, DING, WEIXIN, FEISHU, and WEBHOOK.
	//
	// example:
	//
	// ["EMAIL","SMS"]
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	// The list of receiver identifiers.
	//
	// example:
	//
	// ["webhook-uuid-xxx"]
	Receivers []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyStrategyDetailRoutesChannels) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyDetailRoutesChannels) GoString() string {
	return s.String()
}

func (s *NotifyStrategyDetailRoutesChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotifyStrategyDetailRoutesChannels) GetEnabledSubChannels() []*string {
	return s.EnabledSubChannels
}

func (s *NotifyStrategyDetailRoutesChannels) GetReceivers() []*string {
	return s.Receivers
}

func (s *NotifyStrategyDetailRoutesChannels) SetChannelType(v string) *NotifyStrategyDetailRoutesChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyStrategyDetailRoutesChannels) SetEnabledSubChannels(v []*string) *NotifyStrategyDetailRoutesChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyStrategyDetailRoutesChannels) SetReceivers(v []*string) *NotifyStrategyDetailRoutesChannels {
	s.Receivers = v
	return s
}

func (s *NotifyStrategyDetailRoutesChannels) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyDetailRoutesEffectTimeRange struct {
	// The effective days of the week. Valid values: 0 to 6 (0 = Sunday, 6 = Saturday). The value 7 is not supported.
	//
	// example:
	//
	// [0,1,2,3,4,5,6]
	DayInWeek []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	// The end time of the day in minutes. Valid values: 0 to 1439.
	//
	// example:
	//
	// 1439
	EndTimeInMinute *int32 `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	// The start time of the day in minutes. Valid values: 0 to 1438.
	//
	// example:
	//
	// 0
	StartTimeInMinute *int32 `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	// The IANA time zone identifier.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s NotifyStrategyDetailRoutesEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyDetailRoutesEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) SetDayInWeek(v []*int32) *NotifyStrategyDetailRoutesEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyStrategyDetailRoutesEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyStrategyDetailRoutesEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) SetTimeZone(v string) *NotifyStrategyDetailRoutesEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *NotifyStrategyDetailRoutesEffectTimeRange) Validate() error {
	return dara.Validate(s)
}
