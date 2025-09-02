// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIAlarmRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRule(v *GetDIAlarmRuleResponseBodyDIAlarmRule) *GetDIAlarmRuleResponseBody
	GetDIAlarmRule() *GetDIAlarmRuleResponseBodyDIAlarmRule
	SetRequestId(v string) *GetDIAlarmRuleResponseBody
	GetRequestId() *string
}

type GetDIAlarmRuleResponseBody struct {
	// The details of the alert rule.
	DIAlarmRule *GetDIAlarmRuleResponseBodyDIAlarmRule `json:"DIAlarmRule,omitempty" xml:"DIAlarmRule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4A807D85-AC9F-55F7-A58F-998D5249CAD9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDIAlarmRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDIAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIAlarmRuleResponseBody) GetDIAlarmRule() *GetDIAlarmRuleResponseBodyDIAlarmRule {
	return s.DIAlarmRule
}

func (s *GetDIAlarmRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDIAlarmRuleResponseBody) SetDIAlarmRule(v *GetDIAlarmRuleResponseBodyDIAlarmRule) *GetDIAlarmRuleResponseBody {
	s.DIAlarmRule = v
	return s
}

func (s *GetDIAlarmRuleResponseBody) SetRequestId(v string) *GetDIAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDIAlarmRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDIAlarmRuleResponseBodyDIAlarmRule struct {
	// The timestamp when the alert rule was created. Unit: seconds.
	//
	// example:
	//
	// 1663573162
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the user who creates the alert rule.
	//
	// example:
	//
	// 10000001
	CreatedUid *string `json:"CreatedUid,omitempty" xml:"CreatedUid,omitempty"`
	// The alert rule ID.
	//
	// example:
	//
	// 34988
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The ID of the task with which the alert rule is associated.
	//
	// example:
	//
	// 11170
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
	NotificationSettings *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	// The conditions that are used to trigger the alert rule.
	TriggerConditions []*GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty" type:"Repeated"`
	// The timestamp when the alert rule was last updated. Unit: seconds.
	//
	// example:
	//
	// 1663573163
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The ID of the user who last updates the alert rule.
	//
	// example:
	//
	// 10000001
	UpdatedUid *string `json:"UpdatedUid,omitempty" xml:"UpdatedUid,omitempty"`
}

func (s GetDIAlarmRuleResponseBodyDIAlarmRule) String() string {
	return dara.Prettify(s)
}

func (s GetDIAlarmRuleResponseBodyDIAlarmRule) GoString() string {
	return s.String()
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetCreatedUid() *string {
	return s.CreatedUid
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetDescription() *string {
	return s.Description
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetMetricType() *string {
	return s.MetricType
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetNotificationSettings() *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings {
	return s.NotificationSettings
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetTriggerConditions() []*GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions {
	return s.TriggerConditions
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) GetUpdatedUid() *string {
	return s.UpdatedUid
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetCreatedTime(v int64) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.CreatedTime = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetCreatedUid(v string) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.CreatedUid = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetDIAlarmRuleId(v int64) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.DIAlarmRuleId = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetDIJobId(v int64) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.DIJobId = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetDescription(v string) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.Description = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetEnabled(v bool) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.Enabled = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetMetricType(v string) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.MetricType = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetNotificationSettings(v *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.NotificationSettings = v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetTriggerConditions(v []*GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.TriggerConditions = v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetUpdatedTime(v int64) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.UpdatedTime = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) SetUpdatedUid(v string) *GetDIAlarmRuleResponseBodyDIAlarmRule {
	s.UpdatedUid = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRule) Validate() error {
	return dara.Validate(s)
}

type GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings struct {
	// The duration of the alert suppression interval. Unit: minutes.
	//
	// example:
	//
	// 5
	InhibitionInterval *int32 `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	// The alert notification methods.
	NotificationChannels []*GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The alert notification recipients.
	NotificationReceivers []*GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) String() string {
	return dara.Prettify(s)
}

func (s GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) GoString() string {
	return s.String()
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) GetInhibitionInterval() *int32 {
	return s.InhibitionInterval
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) GetNotificationChannels() []*GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels {
	return s.NotificationChannels
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) GetNotificationReceivers() []*GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) SetInhibitionInterval(v int32) *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings {
	s.InhibitionInterval = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) SetNotificationChannels(v []*GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels) *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings {
	s.NotificationChannels = v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) SetNotificationReceivers(v []*GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers) *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings {
	s.NotificationReceivers = v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettings) Validate() error {
	return dara.Validate(s)
}

type GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels struct {
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

func (s GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels) GoString() string {
	return s.String()
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels) GetSeverity() *string {
	return s.Severity
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels) SetChannels(v []*string) *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels {
	s.Channels = v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels) SetSeverity(v string) *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels {
	s.Severity = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers struct {
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

func (s GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers) SetReceiverType(v string) *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers) SetReceiverValues(v []*string) *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleNotificationSettingsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions struct {
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
	// 	- If the alert rule is for failovers, specify the number of failovers.
	//
	// 	- If the alert rule is for latency, the threshold is the latency duration, in seconds.
	//
	// example:
	//
	// 5
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) GoString() string {
	return s.String()
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) GetDuration() *int64 {
	return s.Duration
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) GetSeverity() *string {
	return s.Severity
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) GetThreshold() *int64 {
	return s.Threshold
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) SetDuration(v int64) *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions {
	s.Duration = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) SetSeverity(v string) *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions {
	s.Severity = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) SetThreshold(v int64) *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions {
	s.Threshold = &v
	return s
}

func (s *GetDIAlarmRuleResponseBodyDIAlarmRuleTriggerConditions) Validate() error {
	return dara.Validate(s)
}
