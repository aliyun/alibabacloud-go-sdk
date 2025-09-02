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
	// The pagination information.
	DIAlarmRulePaging *ListDIAlarmRulesResponseBodyDIAlarmRulePaging `json:"DIAlarmRulePaging,omitempty" xml:"DIAlarmRulePaging,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 74C2FECD-5B3A-554A-BCF5-351A36DE9815
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
	return dara.Validate(s)
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePaging struct {
	// The alert rules.
	DIJobAlarmRules []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules `json:"DIJobAlarmRules,omitempty" xml:"DIJobAlarmRules,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
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
	return dara.Validate(s)
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRules struct {
	// The alert rule ID.
	//
	// example:
	//
	// 41998
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The ID of the task with which the alert rule is associated.
	//
	// example:
	//
	// 11260
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the alert rule.
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
	// The metric type in the alert rule. Valid values:
	//
	// 	- Heartbeat
	//
	// 	- FailoverCount
	//
	// 	- Delay
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The alert notification settings.
	NotificationSettings *ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	// The conditions that are used to trigger the alert rule.
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
	return dara.Validate(s)
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettings struct {
	// The duration of the alert suppression interval. Unit: minutes.
	//
	// example:
	//
	// 5
	InhibitionInterval *int32 `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	// The alert notification methods.
	NotificationChannels []*ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The settings of alert notification recipients.
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
	return dara.Validate(s)
}

type ListDIAlarmRulesResponseBodyDIAlarmRulePagingDIJobAlarmRulesNotificationSettingsNotificationChannels struct {
	// The alert notification methods.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The severity level. Valid values:
	//
	// 	- Warning
	//
	// 	- Critical
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
	// The recipient type. Valid values: AliyunUid and DingToken.
	//
	// 	- If the alert notification method is Mail, Phone, or Sms, the value of this parameter is **AliyunUid**, which indicates the Alibaba Cloud account ID.
	//
	// 	- If the alert notification method is Ding, the value of this parameter is **DingToken**, which indicates the DingTalk chatbot token.
	//
	// example:
	//
	// DingToken
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The recipients.
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
	// The time interval for alert calculation. Unit: minutes.
	//
	// example:
	//
	// 15
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The severity level. Valid values:
	//
	// 	- Warning
	//
	// 	- Critical
	//
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The alert threshold.
	//
	// 	- If the alert rule is for task status, no threshold is used.
	//
	// 	- If the alert rule is for failovers, the threshold is the number of failovers.
	//
	// 	- If the alert rule is for latency, the threshold is the latency duration, in seconds.
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
