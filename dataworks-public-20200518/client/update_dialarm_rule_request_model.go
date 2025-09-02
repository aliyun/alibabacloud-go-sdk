// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDIAlarmRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRuleId(v int64) *UpdateDIAlarmRuleRequest
	GetDIAlarmRuleId() *int64
	SetDescription(v string) *UpdateDIAlarmRuleRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateDIAlarmRuleRequest
	GetEnabled() *bool
	SetMetricType(v string) *UpdateDIAlarmRuleRequest
	GetMetricType() *string
	SetNotificationSettings(v *UpdateDIAlarmRuleRequestNotificationSettings) *UpdateDIAlarmRuleRequest
	GetNotificationSettings() *UpdateDIAlarmRuleRequestNotificationSettings
	SetTriggerConditions(v []*UpdateDIAlarmRuleRequestTriggerConditions) *UpdateDIAlarmRuleRequest
	GetTriggerConditions() []*UpdateDIAlarmRuleRequestTriggerConditions
}

type UpdateDIAlarmRuleRequest struct {
	// The alert rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34982
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The description of the alert rule.
	//
	// example:
	//
	// mysql synchronizes to hologres heartbeat alert
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the alert rule. By default, the alert rule is disabled.
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
	// This parameter is required.
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The alert notification settings.
	//
	// This parameter is required.
	NotificationSettings *UpdateDIAlarmRuleRequestNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	// The conditions that are used to trigger the alert rule.
	//
	// This parameter is required.
	TriggerConditions []*UpdateDIAlarmRuleRequestTriggerConditions `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty" type:"Repeated"`
}

func (s UpdateDIAlarmRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequest) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *UpdateDIAlarmRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDIAlarmRuleRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateDIAlarmRuleRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *UpdateDIAlarmRuleRequest) GetNotificationSettings() *UpdateDIAlarmRuleRequestNotificationSettings {
	return s.NotificationSettings
}

func (s *UpdateDIAlarmRuleRequest) GetTriggerConditions() []*UpdateDIAlarmRuleRequestTriggerConditions {
	return s.TriggerConditions
}

func (s *UpdateDIAlarmRuleRequest) SetDIAlarmRuleId(v int64) *UpdateDIAlarmRuleRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetDescription(v string) *UpdateDIAlarmRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetEnabled(v bool) *UpdateDIAlarmRuleRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetMetricType(v string) *UpdateDIAlarmRuleRequest {
	s.MetricType = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetNotificationSettings(v *UpdateDIAlarmRuleRequestNotificationSettings) *UpdateDIAlarmRuleRequest {
	s.NotificationSettings = v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetTriggerConditions(v []*UpdateDIAlarmRuleRequestTriggerConditions) *UpdateDIAlarmRuleRequest {
	s.TriggerConditions = v
	return s
}

func (s *UpdateDIAlarmRuleRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateDIAlarmRuleRequestNotificationSettings struct {
	// The duration of the alert suppression interval. Default value: 5. Unit: minutes.
	//
	// example:
	//
	// 5
	InhibitionInterval *int32 `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	// The alert notification methods.
	//
	// This parameter is required.
	NotificationChannels []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The settings of alert notification recipients.
	//
	// This parameter is required.
	NotificationReceivers []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s UpdateDIAlarmRuleRequestNotificationSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIAlarmRuleRequestNotificationSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) GetInhibitionInterval() *int32 {
	return s.InhibitionInterval
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) GetNotificationChannels() []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	return s.NotificationChannels
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) GetNotificationReceivers() []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) SetInhibitionInterval(v int32) *UpdateDIAlarmRuleRequestNotificationSettings {
	s.InhibitionInterval = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) SetNotificationChannels(v []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) *UpdateDIAlarmRuleRequestNotificationSettings {
	s.NotificationChannels = v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) SetNotificationReceivers(v []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) *UpdateDIAlarmRuleRequestNotificationSettings {
	s.NotificationReceivers = v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels struct {
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

func (s UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) SetChannels(v []*string) *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	s.Channels = v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) SetSeverity(v string) *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	s.Severity = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers struct {
	// The recipient type.
	//
	// 	- If the alert notification method is Mail, Phone, or Sms, the recipient type is the Alibaba Cloud account ID.
	//
	// 	- If the alert notification method is Ding, the recipient type is the DingTalk chatbot token.
	//
	// example:
	//
	// DingToken
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) SetReceiverType(v string) *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) SetReceiverValues(v []*string) *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type UpdateDIAlarmRuleRequestTriggerConditions struct {
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
	// 	- If the alert rule is for task status, you do not need to specify a threshold.
	//
	// 	- If the alert rule is for failovers, you must specify the number of failovers.
	//
	// 	- If the alert rule is for latency, you must specify the latency duration, in seconds.
	//
	// example:
	//
	// 5
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateDIAlarmRuleRequestTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateDIAlarmRuleRequestTriggerConditions) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) GetDuration() *int64 {
	return s.Duration
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) GetThreshold() *int64 {
	return s.Threshold
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) SetDuration(v int64) *UpdateDIAlarmRuleRequestTriggerConditions {
	s.Duration = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) SetSeverity(v string) *UpdateDIAlarmRuleRequestTriggerConditions {
	s.Severity = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) SetThreshold(v int64) *UpdateDIAlarmRuleRequestTriggerConditions {
	s.Threshold = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) Validate() error {
	return dara.Validate(s)
}
