// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyStrategyForSNSModify interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplateEntries(v []*NotifyStrategyForSNSModifyCustomTemplateEntries) *NotifyStrategyForSNSModify
	GetCustomTemplateEntries() []*NotifyStrategyForSNSModifyCustomTemplateEntries
	SetDescription(v string) *NotifyStrategyForSNSModify
	GetDescription() *string
	SetEnableIncidentManagement(v bool) *NotifyStrategyForSNSModify
	GetEnableIncidentManagement() *bool
	SetGroupingSetting(v *NotifyStrategyForSNSModifyGroupingSetting) *NotifyStrategyForSNSModify
	GetGroupingSetting() *NotifyStrategyForSNSModifyGroupingSetting
	SetIgnoreRestoredNotification(v bool) *NotifyStrategyForSNSModify
	GetIgnoreRestoredNotification() *bool
	SetRoutes(v []*NotifyStrategyForSNSModifyRoutes) *NotifyStrategyForSNSModify
	GetRoutes() []*NotifyStrategyForSNSModifyRoutes
}

type NotifyStrategyForSNSModify struct {
	CustomTemplateEntries    []*NotifyStrategyForSNSModifyCustomTemplateEntries `json:"customTemplateEntries,omitempty" xml:"customTemplateEntries,omitempty" type:"Repeated"`
	Description              *string                                            `json:"description,omitempty" xml:"description,omitempty"`
	EnableIncidentManagement *bool                                              `json:"enableIncidentManagement,omitempty" xml:"enableIncidentManagement,omitempty"`
	// This parameter is required.
	GroupingSetting            *NotifyStrategyForSNSModifyGroupingSetting `json:"groupingSetting,omitempty" xml:"groupingSetting,omitempty" type:"Struct"`
	IgnoreRestoredNotification *bool                                      `json:"ignoreRestoredNotification,omitempty" xml:"ignoreRestoredNotification,omitempty"`
	// This parameter is required.
	Routes []*NotifyStrategyForSNSModifyRoutes `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForSNSModify) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSModify) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSModify) GetCustomTemplateEntries() []*NotifyStrategyForSNSModifyCustomTemplateEntries {
	return s.CustomTemplateEntries
}

func (s *NotifyStrategyForSNSModify) GetDescription() *string {
	return s.Description
}

func (s *NotifyStrategyForSNSModify) GetEnableIncidentManagement() *bool {
	return s.EnableIncidentManagement
}

func (s *NotifyStrategyForSNSModify) GetGroupingSetting() *NotifyStrategyForSNSModifyGroupingSetting {
	return s.GroupingSetting
}

func (s *NotifyStrategyForSNSModify) GetIgnoreRestoredNotification() *bool {
	return s.IgnoreRestoredNotification
}

func (s *NotifyStrategyForSNSModify) GetRoutes() []*NotifyStrategyForSNSModifyRoutes {
	return s.Routes
}

func (s *NotifyStrategyForSNSModify) SetCustomTemplateEntries(v []*NotifyStrategyForSNSModifyCustomTemplateEntries) *NotifyStrategyForSNSModify {
	s.CustomTemplateEntries = v
	return s
}

func (s *NotifyStrategyForSNSModify) SetDescription(v string) *NotifyStrategyForSNSModify {
	s.Description = &v
	return s
}

func (s *NotifyStrategyForSNSModify) SetEnableIncidentManagement(v bool) *NotifyStrategyForSNSModify {
	s.EnableIncidentManagement = &v
	return s
}

func (s *NotifyStrategyForSNSModify) SetGroupingSetting(v *NotifyStrategyForSNSModifyGroupingSetting) *NotifyStrategyForSNSModify {
	s.GroupingSetting = v
	return s
}

func (s *NotifyStrategyForSNSModify) SetIgnoreRestoredNotification(v bool) *NotifyStrategyForSNSModify {
	s.IgnoreRestoredNotification = &v
	return s
}

func (s *NotifyStrategyForSNSModify) SetRoutes(v []*NotifyStrategyForSNSModifyRoutes) *NotifyStrategyForSNSModify {
	s.Routes = v
	return s
}

func (s *NotifyStrategyForSNSModify) Validate() error {
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

type NotifyStrategyForSNSModifyCustomTemplateEntries struct {
	TargetType   *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s NotifyStrategyForSNSModifyCustomTemplateEntries) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSModifyCustomTemplateEntries) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSModifyCustomTemplateEntries) GetTargetType() *string {
	return s.TargetType
}

func (s *NotifyStrategyForSNSModifyCustomTemplateEntries) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *NotifyStrategyForSNSModifyCustomTemplateEntries) SetTargetType(v string) *NotifyStrategyForSNSModifyCustomTemplateEntries {
	s.TargetType = &v
	return s
}

func (s *NotifyStrategyForSNSModifyCustomTemplateEntries) SetTemplateUuid(v string) *NotifyStrategyForSNSModifyCustomTemplateEntries {
	s.TemplateUuid = &v
	return s
}

func (s *NotifyStrategyForSNSModifyCustomTemplateEntries) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForSNSModifyGroupingSetting struct {
	GroupingKeys []*string `json:"groupingKeys,omitempty" xml:"groupingKeys,omitempty" type:"Repeated"`
	PeriodMin    *int32    `json:"periodMin,omitempty" xml:"periodMin,omitempty"`
	SilenceSec   *int32    `json:"silenceSec,omitempty" xml:"silenceSec,omitempty"`
	Times        *int32    `json:"times,omitempty" xml:"times,omitempty"`
}

func (s NotifyStrategyForSNSModifyGroupingSetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSModifyGroupingSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) GetGroupingKeys() []*string {
	return s.GroupingKeys
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) GetPeriodMin() *int32 {
	return s.PeriodMin
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) GetSilenceSec() *int32 {
	return s.SilenceSec
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) GetTimes() *int32 {
	return s.Times
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) SetGroupingKeys(v []*string) *NotifyStrategyForSNSModifyGroupingSetting {
	s.GroupingKeys = v
	return s
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) SetPeriodMin(v int32) *NotifyStrategyForSNSModifyGroupingSetting {
	s.PeriodMin = &v
	return s
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) SetSilenceSec(v int32) *NotifyStrategyForSNSModifyGroupingSetting {
	s.SilenceSec = &v
	return s
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) SetTimes(v int32) *NotifyStrategyForSNSModifyGroupingSetting {
	s.Times = &v
	return s
}

func (s *NotifyStrategyForSNSModifyGroupingSetting) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForSNSModifyRoutes struct {
	Channels            []*NotifyStrategyForSNSModifyRoutesChannels `json:"channels,omitempty" xml:"channels,omitempty" type:"Repeated"`
	DigitalEmployeeName *string                                     `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The effective period settings for notifications. Defines on which days and during which time range the system sends notifications.
	EffectTimeRange *NotifyStrategyForSNSModifyRoutesEffectTimeRange `json:"effectTimeRange,omitempty" xml:"effectTimeRange,omitempty" type:"Struct"`
	EnableRca       *bool                                            `json:"enableRca,omitempty" xml:"enableRca,omitempty"`
	FilterSetting   *NotifyStrategyForSNSModifyRoutesFilterSetting   `json:"filterSetting,omitempty" xml:"filterSetting,omitempty" type:"Struct"`
	Severities      []*string                                        `json:"severities,omitempty" xml:"severities,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForSNSModifyRoutes) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSModifyRoutes) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSModifyRoutes) GetChannels() []*NotifyStrategyForSNSModifyRoutesChannels {
	return s.Channels
}

func (s *NotifyStrategyForSNSModifyRoutes) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *NotifyStrategyForSNSModifyRoutes) GetEffectTimeRange() *NotifyStrategyForSNSModifyRoutesEffectTimeRange {
	return s.EffectTimeRange
}

func (s *NotifyStrategyForSNSModifyRoutes) GetEnableRca() *bool {
	return s.EnableRca
}

func (s *NotifyStrategyForSNSModifyRoutes) GetFilterSetting() *NotifyStrategyForSNSModifyRoutesFilterSetting {
	return s.FilterSetting
}

func (s *NotifyStrategyForSNSModifyRoutes) GetSeverities() []*string {
	return s.Severities
}

func (s *NotifyStrategyForSNSModifyRoutes) SetChannels(v []*NotifyStrategyForSNSModifyRoutesChannels) *NotifyStrategyForSNSModifyRoutes {
	s.Channels = v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutes) SetDigitalEmployeeName(v string) *NotifyStrategyForSNSModifyRoutes {
	s.DigitalEmployeeName = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutes) SetEffectTimeRange(v *NotifyStrategyForSNSModifyRoutesEffectTimeRange) *NotifyStrategyForSNSModifyRoutes {
	s.EffectTimeRange = v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutes) SetEnableRca(v bool) *NotifyStrategyForSNSModifyRoutes {
	s.EnableRca = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutes) SetFilterSetting(v *NotifyStrategyForSNSModifyRoutesFilterSetting) *NotifyStrategyForSNSModifyRoutes {
	s.FilterSetting = v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutes) SetSeverities(v []*string) *NotifyStrategyForSNSModifyRoutes {
	s.Severities = v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutes) Validate() error {
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

type NotifyStrategyForSNSModifyRoutesChannels struct {
	// The notification channel type. The value must be one of the following uppercase enum values: DING (DingTalk chatbot), WEIXIN (WeCom chatbot), FEISHU (Lark chatbot), SLACK, TEAMS, WEBHOOK (custom webhook), CONTACT (contact, requires enabledSubChannels to specify sub-channels), GROUP (contact group), DUTY (on-call schedule), or DING_COOL_APP (DingTalk Cool App). Note: Lowercase values such as EMAIL or SMS are not supported. To send email, text message, or voice notifications, set channelType to CONTACT and specify EMAIL, SMS, or VOICE in enabledSubChannels.
	//
	// This parameter is required.
	//
	// example:
	//
	// WEBHOOK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// Required only when channelType is CONTACT, GROUP, or DUTY. Valid values: EMAIL (email), SMS (text message), VOICE (voice call), DING (DingTalk work notification), WEIXIN (WeCom message), FEISHU (Lark message), and WEBHOOK. For example, to notify a contact by email and text message, set channelType to CONTACT and enabledSubChannels to ["EMAIL","SMS"]. This field is not required for other channelType values such as WEBHOOK or DING.
	//
	// example:
	//
	// ["EMAIL","SMS"]
	EnabledSubChannels []*string `json:"enabledSubChannels,omitempty" xml:"enabledSubChannels,omitempty" type:"Repeated"`
	// The list of receiver identifiers. For the WEBHOOK type, specify the webhook UUID. For DING, WEIXIN, or FEISHU, specify the chatbot UUID. For CONTACT, specify the contact ID. For GROUP, specify the contact group ID. For DUTY, specify the on-call schedule UUID. At least one element is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["my-webhook-uuid"]
	Receivers []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s NotifyStrategyForSNSModifyRoutesChannels) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSModifyRoutesChannels) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSModifyRoutesChannels) GetChannelType() *string {
	return s.ChannelType
}

func (s *NotifyStrategyForSNSModifyRoutesChannels) GetEnabledSubChannels() []*string {
	return s.EnabledSubChannels
}

func (s *NotifyStrategyForSNSModifyRoutesChannels) GetReceivers() []*string {
	return s.Receivers
}

func (s *NotifyStrategyForSNSModifyRoutesChannels) SetChannelType(v string) *NotifyStrategyForSNSModifyRoutesChannels {
	s.ChannelType = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesChannels) SetEnabledSubChannels(v []*string) *NotifyStrategyForSNSModifyRoutesChannels {
	s.EnabledSubChannels = v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesChannels) SetReceivers(v []*string) *NotifyStrategyForSNSModifyRoutesChannels {
	s.Receivers = v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesChannels) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForSNSModifyRoutesEffectTimeRange struct {
	// The days of the week on which the setting takes effect. Array element values range from 0 to 6 (0 = Sunday, 1 = Monday, 2 = Tuesday, ... 6 = Saturday). Note: The value 7 is not supported. The maximum value is 6. Example for all days: [0,1,2,3,4,5,6]. Example for weekdays only: [1,2,3,4,5].
	//
	// example:
	//
	// [0,1,2,3,4,5,6]
	DayInWeek []*int32 `json:"dayInWeek,omitempty" xml:"dayInWeek,omitempty" type:"Repeated"`
	// The end time of the day, expressed as the number of minutes from 00:00. Valid values: 0 to 1439 (23 × 60 + 59 = 1439, which represents 23:59).
	//
	// example:
	//
	// 1439
	EndTimeInMinute *int32 `json:"endTimeInMinute,omitempty" xml:"endTimeInMinute,omitempty"`
	// The start time of the day, expressed as the number of minutes from 00:00. Valid values: 0 to 1439 (0 represents 00:00).
	//
	// example:
	//
	// 0
	StartTimeInMinute *int32 `json:"startTimeInMinute,omitempty" xml:"startTimeInMinute,omitempty"`
	// The IANA time zone identifier, such as Asia/Shanghai or America/Los_Angeles.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s NotifyStrategyForSNSModifyRoutesEffectTimeRange) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSModifyRoutesEffectTimeRange) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) GetDayInWeek() []*int32 {
	return s.DayInWeek
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) GetEndTimeInMinute() *int32 {
	return s.EndTimeInMinute
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) GetStartTimeInMinute() *int32 {
	return s.StartTimeInMinute
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) GetTimeZone() *string {
	return s.TimeZone
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) SetDayInWeek(v []*int32) *NotifyStrategyForSNSModifyRoutesEffectTimeRange {
	s.DayInWeek = v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) SetEndTimeInMinute(v int32) *NotifyStrategyForSNSModifyRoutesEffectTimeRange {
	s.EndTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) SetStartTimeInMinute(v int32) *NotifyStrategyForSNSModifyRoutesEffectTimeRange {
	s.StartTimeInMinute = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) SetTimeZone(v string) *NotifyStrategyForSNSModifyRoutesEffectTimeRange {
	s.TimeZone = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesEffectTimeRange) Validate() error {
	return dara.Validate(s)
}

type NotifyStrategyForSNSModifyRoutesFilterSetting struct {
	Conditions []*NotifyStrategyForSNSModifyRoutesFilterSettingConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Expression *string                                                    `json:"expression,omitempty" xml:"expression,omitempty"`
	Relation   *string                                                    `json:"relation,omitempty" xml:"relation,omitempty"`
}

func (s NotifyStrategyForSNSModifyRoutesFilterSetting) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSModifyRoutesFilterSetting) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSetting) GetConditions() []*NotifyStrategyForSNSModifyRoutesFilterSettingConditions {
	return s.Conditions
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSetting) GetExpression() *string {
	return s.Expression
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSetting) GetRelation() *string {
	return s.Relation
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSetting) SetConditions(v []*NotifyStrategyForSNSModifyRoutesFilterSettingConditions) *NotifyStrategyForSNSModifyRoutesFilterSetting {
	s.Conditions = v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSetting) SetExpression(v string) *NotifyStrategyForSNSModifyRoutesFilterSetting {
	s.Expression = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSetting) SetRelation(v string) *NotifyStrategyForSNSModifyRoutesFilterSetting {
	s.Relation = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSetting) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NotifyStrategyForSNSModifyRoutesFilterSettingConditions struct {
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	Op    *string `json:"op,omitempty" xml:"op,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s NotifyStrategyForSNSModifyRoutesFilterSettingConditions) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyForSNSModifyRoutesFilterSettingConditions) GoString() string {
	return s.String()
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSettingConditions) GetField() *string {
	return s.Field
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSettingConditions) GetOp() *string {
	return s.Op
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSettingConditions) GetValue() *string {
	return s.Value
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSettingConditions) SetField(v string) *NotifyStrategyForSNSModifyRoutesFilterSettingConditions {
	s.Field = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSettingConditions) SetOp(v string) *NotifyStrategyForSNSModifyRoutesFilterSettingConditions {
	s.Op = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSettingConditions) SetValue(v string) *NotifyStrategyForSNSModifyRoutesFilterSettingConditions {
	s.Value = &v
	return s
}

func (s *NotifyStrategyForSNSModifyRoutesFilterSettingConditions) Validate() error {
	return dara.Validate(s)
}
