// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIAlarmRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRulePaging(v *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) *ListDIAlarmRulesResponseBody
	GetDIAlarmRulePaging() *ListDIAlarmRulesResponseBodyDIAlarmRulePaging
	SetRequestId(v string) *ListDIAlarmRulesResponseBody
	GetRequestId() *string
}

type ListDIAlarmRulesResponseBody struct {
	// The paginated result of alert rules.
	DIAlarmRulePaging *ListDIAlarmRulesResponseBodyDIAlarmRulePaging `json:"DIAlarmRulePaging,omitempty" xml:"DIAlarmRulePaging,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 74C2FECD-5B3A-554A-BCF5-35****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIAlarmRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBody) GetDIAlarmRulePaging() *ListDIAlarmRulesResponseBodyDIAlarmRulePaging {
	return s.DIAlarmRulePaging
}

func (s *ListDIAlarmRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDIAlarmRulesResponseBody) SetDIAlarmRulePaging(v *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) *ListDIAlarmRulesResponseBody {
	s.DIAlarmRulePaging = v
	return s
}

func (s *ListDIAlarmRulesResponseBody) SetRequestId(v string) *ListDIAlarmRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDIAlarmRulesResponseBody) Validate() error {
	if s.DIAlarmRulePaging != nil {
		if err := s.DIAlarmRulePaging.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePaging struct {
	// The list of alert rules.
	DIJobAlarmRules []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules `json:"DIJobAlarmRules,omitempty" xml:"DIJobAlarmRules,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePaging) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePaging) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) GetDIJobAlarmRules() []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules {
	return s.DIJobAlarmRules
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) SetDIJobAlarmRules(v []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) *ListDIAlarmRulesResponseBodyDIAlarmRulePaging {
	s.DIJobAlarmRules = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) SetPageNumber(v int64) *ListDIAlarmRulesResponseBodyDIAlarmRulePaging {
	s.PageNumber = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) SetPageSize(v int64) *ListDIAlarmRulesResponseBodyDIAlarmRulePaging {
	s.PageSize = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) SetTotalCount(v int64) *ListDIAlarmRulesResponseBodyDIAlarmRulePaging {
	s.TotalCount = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePaging) Validate() error {
	if s.DIJobAlarmRules != nil {
		for _, item := range s.DIJobAlarmRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules struct {
	// The alert rule ID.
	//
	// example:
	//
	// 41998
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The task ID. This is the task ID associated with the alert rule.
	//
	// example:
	//
	// 11260
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description.
	//
	// example:
	//
	// mysql synchronizes to hologres heartbeat alert
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the alert rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The alert metric type. Valid values:
	//
	// - Heartbeat: task status alert.
	//
	// - FailoverCount: failover count alert.
	//
	// - Delay: task delay alert.
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The alert notification settings.
	NotificationSettings *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	// The list of alert trigger conditions. Multiple conditions are supported.
	TriggerConditions []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) GetDescription() *string {
	return s.Description
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) GetMetricType() *string {
	return s.MetricType
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) GetNotificationSettings() *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings {
	return s.NotificationSettings
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) GetTriggerConditions() []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions {
	return s.TriggerConditions
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) SetDIAlarmRuleId(v int64) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules {
	s.DIAlarmRuleId = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) SetDIJobId(v int64) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules {
	s.DIJobId = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) SetDescription(v string) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules {
	s.Description = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) SetEnabled(v bool) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules {
	s.Enabled = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) SetMetricType(v string) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules {
	s.MetricType = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) SetNotificationSettings(v *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules {
	s.NotificationSettings = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) SetTriggerConditions(v []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules {
	s.TriggerConditions = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules) Validate() error {
	if s.NotificationSettings != nil {
		if err := s.NotificationSettings.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerConditions != nil {
		for _, item := range s.TriggerConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings struct {
	// The alert suppression interval. Unit: minutes.
	//
	// example:
	//
	// 5
	InhibitionInterval *int32 `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	// The alert notification channels. Multiple values are supported.
	NotificationChannels []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The alert notification receivers. Multiple values are supported.
	NotificationReceivers []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) GetInhibitionInterval() *int32 {
	return s.InhibitionInterval
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) GetNotificationChannels() []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels {
	return s.NotificationChannels
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) GetNotificationReceivers() []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) SetInhibitionInterval(v int32) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings {
	s.InhibitionInterval = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) SetNotificationChannels(v []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings {
	s.NotificationChannels = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) SetNotificationReceivers(v []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings {
	s.NotificationReceivers = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings) Validate() error {
	if s.NotificationChannels != nil {
		for _, item := range s.NotificationChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NotificationReceivers != nil {
		for _, item := range s.NotificationReceivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels struct {
	// The list of channels.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The severity level. Valid values:
	//
	// - Warning
	//
	// - Critical
	//
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels) GetSeverity() *string {
	return s.Severity
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels) SetChannels(v []*string) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels {
	s.Channels = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels) SetSeverity(v string) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels {
	s.Severity = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers struct {
	// The receiver type.
	//
	// - If the alert notification channel is email, phone call, or text message, the receiver type is Alibaba Cloud user ID (**AliyunUid**).
	//
	// - If the alert notification channel is DingTalk, the receiver type is DingTalk token (**DingToken**).
	//
	// example:
	//
	// DingToken
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The list of receiver values.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers) SetReceiverType(v string) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers) SetReceiverValues(v []*string) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions struct {
	// The time window for alert calculation. Unit: minutes.
	//
	// example:
	//
	// 15
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The severity level. Valid values:
	//
	// - Warning
	//
	// - Critical
	//
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The alert threshold.
	//
	// - Task status alert: no threshold.
	//
	// - Failover count alert: the threshold is the number of failovers.
	//
	// - Task delay alert: the threshold is the delay duration. Unit: seconds.
	//
	// example:
	//
	// 5
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) GetDuration() *int64 {
	return s.Duration
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) GetSeverity() *string {
	return s.Severity
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) GetThreshold() *int64 {
	return s.Threshold
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) SetDuration(v int64) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions {
	s.Duration = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) SetSeverity(v string) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions {
	s.Severity = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) SetThreshold(v int64) *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions {
	s.Threshold = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesTriggerConditions) Validate() error {
	return dara.Validate(s)
}
