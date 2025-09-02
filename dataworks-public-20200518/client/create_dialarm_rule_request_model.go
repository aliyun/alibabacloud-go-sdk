// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIAlarmRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *CreateDIAlarmRuleRequest
	GetDIJobId() *int64
	SetDescription(v string) *CreateDIAlarmRuleRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateDIAlarmRuleRequest
	GetEnabled() *bool
	SetMetricType(v string) *CreateDIAlarmRuleRequest
	GetMetricType() *string
	SetNotificationSettings(v *CreateDIAlarmRuleRequestNotificationSettings) *CreateDIAlarmRuleRequest
	GetNotificationSettings() *CreateDIAlarmRuleRequestNotificationSettings
	SetTriggerConditions(v []*CreateDIAlarmRuleRequestTriggerConditions) *CreateDIAlarmRuleRequest
	GetTriggerConditions() []*CreateDIAlarmRuleRequestTriggerConditions
}

type CreateDIAlarmRuleRequest struct {
	// The ID of the task with which the alert rule is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11265
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the task.
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
	NotificationSettings *CreateDIAlarmRuleRequestNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	// The conditions that are used to trigger the alert rule.
	//
	// This parameter is required.
	TriggerConditions []*CreateDIAlarmRuleRequestTriggerConditions `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty" type:"Repeated"`
}

func (s CreateDIAlarmRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *CreateDIAlarmRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDIAlarmRuleRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateDIAlarmRuleRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *CreateDIAlarmRuleRequest) GetNotificationSettings() *CreateDIAlarmRuleRequestNotificationSettings {
	return s.NotificationSettings
}

func (s *CreateDIAlarmRuleRequest) GetTriggerConditions() []*CreateDIAlarmRuleRequestTriggerConditions {
	return s.TriggerConditions
}

func (s *CreateDIAlarmRuleRequest) SetDIJobId(v int64) *CreateDIAlarmRuleRequest {
	s.DIJobId = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetDescription(v string) *CreateDIAlarmRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetEnabled(v bool) *CreateDIAlarmRuleRequest {
	s.Enabled = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetMetricType(v string) *CreateDIAlarmRuleRequest {
	s.MetricType = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetNotificationSettings(v *CreateDIAlarmRuleRequestNotificationSettings) *CreateDIAlarmRuleRequest {
	s.NotificationSettings = v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetTriggerConditions(v []*CreateDIAlarmRuleRequestTriggerConditions) *CreateDIAlarmRuleRequest {
	s.TriggerConditions = v
	return s
}

func (s *CreateDIAlarmRuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDIAlarmRuleRequestNotificationSettings struct {
	// The duration of the alert suppression interval. Default value: 5. Unit: minutes.
	//
	// example:
	//
	// 5
	InhibitionInterval *int32 `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	// The alert notification methods.
	//
	// This parameter is required.
	NotificationChannels []*CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The settings of alert notification recipients.
	//
	// This parameter is required.
	NotificationReceivers []*CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s CreateDIAlarmRuleRequestNotificationSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleRequestNotificationSettings) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) GetInhibitionInterval() *int32 {
	return s.InhibitionInterval
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) GetNotificationChannels() []*CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	return s.NotificationChannels
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) GetNotificationReceivers() []*CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) SetInhibitionInterval(v int32) *CreateDIAlarmRuleRequestNotificationSettings {
	s.InhibitionInterval = &v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) SetNotificationChannels(v []*CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) *CreateDIAlarmRuleRequestNotificationSettings {
	s.NotificationChannels = v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) SetNotificationReceivers(v []*CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) *CreateDIAlarmRuleRequestNotificationSettings {
	s.NotificationReceivers = v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) Validate() error {
	return dara.Validate(s)
}

type CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels struct {
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

func (s CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) GetSeverity() *string {
	return s.Severity
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) SetChannels(v []*string) *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	s.Channels = v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) SetSeverity(v string) *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	s.Severity = &v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers struct {
	// The recipient type. Valid values: AliyunUid and DingToken.
	//
	// 	- If the alert notification method is Mail, Phone, or Sms, set this parameter to **AliyunUid**, which specifies the Alibaba Cloud account ID.
	//
	// 	- If the alert notification method is Ding, set this parameter to **DingToken**, which indicates the DingTalk chatbot token.
	//
	// example:
	//
	// DingToken
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) SetReceiverType(v string) *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) SetReceiverValues(v []*string) *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type CreateDIAlarmRuleRequestTriggerConditions struct {
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
	// 	- If the alert rule is for failovers, specify the number of failovers.
	//
	// 	- If the alert rule is for latency, the threshold is the latency duration, in seconds.
	//
	// example:
	//
	// 5
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateDIAlarmRuleRequestTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleRequestTriggerConditions) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) GetSeverity() *string {
	return s.Severity
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) GetThreshold() *int64 {
	return s.Threshold
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetDuration(v int64) *CreateDIAlarmRuleRequestTriggerConditions {
	s.Duration = &v
	return s
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetSeverity(v string) *CreateDIAlarmRuleRequestTriggerConditions {
	s.Severity = &v
	return s
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetThreshold(v int64) *CreateDIAlarmRuleRequestTriggerConditions {
	s.Threshold = &v
	return s
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) Validate() error {
	return dara.Validate(s)
}
