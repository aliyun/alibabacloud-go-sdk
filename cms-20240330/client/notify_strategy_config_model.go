// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyStrategyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplateEntries(v []*NotifyStrategyConfigCustomTemplateEntries) *NotifyStrategyConfig
	GetCustomTemplateEntries() []*NotifyStrategyConfigCustomTemplateEntries
	SetDescription(v string) *NotifyStrategyConfig
	GetDescription() *string
	SetGroupingSetting(v *NotifyStrategyConfigGroupingSetting) *NotifyStrategyConfig
	GetGroupingSetting() *NotifyStrategyConfigGroupingSetting
	SetIgnoreRestoredNotification(v bool) *NotifyStrategyConfig
	GetIgnoreRestoredNotification() *bool
	SetRoutes(v []*NotifyStrategyConfigRoutes) *NotifyStrategyConfig
	GetRoutes() []*NotifyStrategyConfigRoutes
}

type NotifyStrategyConfig struct {
	// The list of custom notification templates.
	//
	// example:
	//
	// []
	CustomTemplateEntries []*NotifyStrategyConfigCustomTemplateEntries `json:"customTemplateEntries,omitempty" xml:"customTemplateEntries,omitempty" type:"Repeated"`
	// The description of the notification policy.
	//
	// example:
	//
	// 生产环境告警通知策略
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The noise reduction settings.
	//
	// This parameter is required.
	GroupingSetting *NotifyStrategyConfigGroupingSetting `json:"groupingSetting,omitempty" xml:"groupingSetting,omitempty" type:"Struct"`
	// Specifies whether to ignore notifications for recovery events. A value of true indicates that recovery notifications are not sent.
	//
	// example:
	//
	// true
	IgnoreRestoredNotification *bool `json:"ignoreRestoredNotification,omitempty" xml:"ignoreRestoredNotification,omitempty"`
	// The list of notification channel routing settings.
	//
	// This parameter is required.
	//
	// example:
	//
	// []
	Routes []*NotifyStrategyConfigRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s NotifyStrategyConfig) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyConfig) GoString() string {
	return s.String()
}

func (s *NotifyStrategyConfig) GetCustomTemplateEntries() []*NotifyStrategyConfigCustomTemplateEntries {
	return s.CustomTemplateEntries
}

func (s *NotifyStrategyConfig) GetDescription() *string {
	return s.Description
}

func (s *NotifyStrategyConfig) GetGroupingSetting() *NotifyStrategyConfigGroupingSetting {
	return s.GroupingSetting
}

func (s *NotifyStrategyConfig) GetIgnoreRestoredNotification() *bool {
	return s.IgnoreRestoredNotification
}

func (s *NotifyStrategyConfig) GetRoutes() []*NotifyStrategyConfigRoutes {
	return s.Routes
}

func (s *NotifyStrategyConfig) SetCustomTemplateEntries(v []*NotifyStrategyConfigCustomTemplateEntries) *NotifyStrategyConfig {
	s.CustomTemplateEntries = v
	return s
}

func (s *NotifyStrategyConfig) SetDescription(v string) *NotifyStrategyConfig {
	s.Description = &v
	return s
}

func (s *NotifyStrategyConfig) SetGroupingSetting(v *NotifyStrategyConfigGroupingSetting) *NotifyStrategyConfig {
	s.GroupingSetting = v
	return s
}

func (s *NotifyStrategyConfig) SetIgnoreRestoredNotification(v bool) *NotifyStrategyConfig {
	s.IgnoreRestoredNotification = &v
	return s
}

func (s *NotifyStrategyConfig) SetRoutes(v []*NotifyStrategyConfigRoutes) *NotifyStrategyConfig {
	s.Routes = v
	return s
}

func (s *NotifyStrategyConfig) Validate() error {
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

type NotifyStrategyConfigCustomTemplateEntries struct {
	// The UUID of the notification template.
	//
	// example:
	//
	// template-uuid-xxx
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyConfigCustomTemplateEntries) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyConfigCustomTemplateEntries) GoString() string {
	return s.String()
}

func (s *NotifyStrategyConfigCustomTemplateEntries) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotifyStrategyConfigCustomTemplateEntries) SetTemplateUuid(v string) *NotifyStrategyConfigCustomTemplateEntries {
	s.TemplateUuid = &v
	return s
}

func (s *NotifyStrategyConfigCustomTemplateEntries) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyConfigGroupingSetting struct {
	// The event fields by which events are grouped. Events in the same group are merged into a single notification. An empty array indicates no grouping.
	//
	// example:
	//
	// ["alertname"]
	GroupingKeys []*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty" type:"Repeated"`
	// This parameter does not take effect for this operation. You do not need to set this parameter.
	//
	// example:
	//
	// 0
	PeriodMin *int32 `json:"periodMin,omitempty" xml:"periodMin,omitempty"`
	// This parameter does not take effect for this operation. You do not need to set this parameter.
	//
	// example:
	//
	// 0
	SilenceSec *int32 `json:"silenceSec,omitempty" xml:"silenceSec,omitempty"`
	// This parameter does not take effect for this operation. You do not need to set this parameter.
	//
	// example:
	//
	// 0
	Times *int32 `json:"times,omitempty" xml:"times,omitempty"`
}

func (s NotifyStrategyConfigGroupingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyConfigGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyConfigGroupingSetting) GetGroupingKeys() []*string {
	return s.GroupingKeys
}

func (s *NotifyStrategyConfigGroupingSetting) GetPeriodMin() *int32 {
	return s.PeriodMin
}

func (s *NotifyStrategyConfigGroupingSetting) GetSilenceSec() *int32 {
	return s.SilenceSec
}

func (s *NotifyStrategyConfigGroupingSetting) GetTimes() *int32 {
	return s.Times
}

func (s *NotifyStrategyConfigGroupingSetting) SetGroupingKeys(v []*string) *NotifyStrategyConfigGroupingSetting {
	s.GroupingKeys = v
	return s
}

func (s *NotifyStrategyConfigGroupingSetting) SetPeriodMin(v int32) *NotifyStrategyConfigGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotifyStrategyConfigGroupingSetting) SetSilenceSec(v int32) *NotifyStrategyConfigGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotifyStrategyConfigGroupingSetting) SetTimes(v int32) *NotifyStrategyConfigGroupingSetting {
	s.Times = &v
	return s
}

func (s *NotifyStrategyConfigGroupingSetting) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyConfigRoutes struct {
	// The list of notification channels.
	//
	// example:
	//
	// []
	Channels []*NotifyStrategyConfigRoutesChannels `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	// The digital employee name. Required when enableRca is set to true.
	//
	// example:
	//
	// 数字员工名
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The effective time range.
	EffectTimeRange *NotifyStrategyConfigRoutesEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	// Specifies whether to enable Root Cause Analysis (RCA).
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

func (s NotifyStrategyConfigRoutes) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyConfigRoutes) GoString() string {
	return s.String()
}

func (s *NotifyStrategyConfigRoutes) GetChannels() []*NotifyStrategyConfigRoutesChannels {
	return s.Channels
}

func (s *NotifyStrategyConfigRoutes) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *NotifyStrategyConfigRoutes) GetEffectTimeRange() *NotifyStrategyConfigRoutesEffectTimeRange {
	return s.EffectTimeRange
}

func (s *NotifyStrategyConfigRoutes) GetEnableRca() *bool {
	return s.EnableRca
}

func (s *NotifyStrategyConfigRoutes) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *NotifyStrategyConfigRoutes) SetChannels(v []*NotifyStrategyConfigRoutesChannels) *NotifyStrategyConfigRoutes {
	s.Channels = v
	return s
}

func (s *NotifyStrategyConfigRoutes) SetDigitalEmployeeName(v string) *NotifyStrategyConfigRoutes {
	s.DigitalEmployeeName = &v
	return s
}

func (s *NotifyStrategyConfigRoutes) SetEffectTimeRange(v *NotifyStrategyConfigRoutesEffectTimeRange) *NotifyStrategyConfigRoutes {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyStrategyConfigRoutes) SetEnableRca(v bool) *NotifyStrategyConfigRoutes {
	s.EnableRca = &v
	return s
}

func (s *NotifyStrategyConfigRoutes) SetFilterSetting(v *FilterSetting) *NotifyStrategyConfigRoutes {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyConfigRoutes) Validate() error {
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

type NotifyStrategyConfigRoutesChannels struct {
	// The channel type. Valid values: DING, WEIXIN, FEISHU, SLACK, TEAMS, WEBHOOK, CONTACT, GROUP, DUTY, and DING_COOL_APP. Lowercase values are not supported. For email, text message, or phone call notifications, use CONTACT with enabledSubChannels.
	//
	// This parameter is required.
	//
	// example:
	//
	// WEBHOOK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// Required only for CONTACT, GROUP, or DUTY. The sub-channel types in uppercase. Valid values: EMAIL, SMS, VOICE, DING, WEIXIN, FEISHU, and WEBHOOK.
	//
	// example:
	//
	// ["EMAIL","SMS"]
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	// The list of receiver identifiers. At least one receiver is required. Specify a webhook UUID for WEBHOOK, a robot UUID for chatbots, or a contact ID for CONTACT.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["webhook-uuid-xxx"]
	Receivers []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyStrategyConfigRoutesChannels) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyConfigRoutesChannels) GoString() string {
	return s.String()
}

func (s *NotifyStrategyConfigRoutesChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotifyStrategyConfigRoutesChannels) GetEnabledSubChannels() []*string {
	return s.EnabledSubChannels
}

func (s *NotifyStrategyConfigRoutesChannels) GetReceivers() []*string {
	return s.Receivers
}

func (s *NotifyStrategyConfigRoutesChannels) SetChannelType(v string) *NotifyStrategyConfigRoutesChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyStrategyConfigRoutesChannels) SetEnabledSubChannels(v []*string) *NotifyStrategyConfigRoutesChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyStrategyConfigRoutesChannels) SetReceivers(v []*string) *NotifyStrategyConfigRoutesChannels {
	s.Receivers = v
	return s
}

func (s *NotifyStrategyConfigRoutesChannels) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyConfigRoutesEffectTimeRange struct {
	// The effective days. Valid values: 0 to 6 (0 = Sunday, 6 = Saturday). The value 7 is not supported.
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

func (s NotifyStrategyConfigRoutesEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyConfigRoutesEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) SetDayInWeek(v []*int32) *NotifyStrategyConfigRoutesEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyStrategyConfigRoutesEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyStrategyConfigRoutesEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) SetTimeZone(v string) *NotifyStrategyConfigRoutesEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *NotifyStrategyConfigRoutesEffectTimeRange) Validate() error {
	return dara.Validate(s)
}
