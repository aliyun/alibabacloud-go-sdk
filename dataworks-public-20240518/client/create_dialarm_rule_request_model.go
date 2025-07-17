// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIAlarmRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDIAlarmRuleRequest
	GetClientToken() *string
	SetDIJobId(v int64) *CreateDIAlarmRuleRequest
	GetDIJobId() *int64
	SetDescription(v string) *CreateDIAlarmRuleRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateDIAlarmRuleRequest
	GetEnabled() *bool
	SetMetricType(v string) *CreateDIAlarmRuleRequest
	GetMetricType() *string
	SetName(v string) *CreateDIAlarmRuleRequest
	GetName() *string
	SetNotificationSettings(v *CreateDIAlarmRuleRequestNotificationSettings) *CreateDIAlarmRuleRequest
	GetNotificationSettings() *CreateDIAlarmRuleRequestNotificationSettings
	SetTriggerConditions(v []*CreateDIAlarmRuleRequestTriggerConditions) *CreateDIAlarmRuleRequest
	GetTriggerConditions() []*CreateDIAlarmRuleRequestTriggerConditions
}

type CreateDIAlarmRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the synchronization task with which the alert rule is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the alert rule.
	//
	// example:
	//
	// The description of the alert rule.
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
	// 	- DdlReport
	//
	// 	- ResourceUtilization
	//
	// This parameter is required.
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The name of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// alartRule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert notification settings.
	//
	// This parameter is required.
	NotificationSettings *CreateDIAlarmRuleRequestNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	// The conditions that can trigger the alert rule.
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

func (s *CreateDIAlarmRuleRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *CreateDIAlarmRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateDIAlarmRuleRequest) GetNotificationSettings() *CreateDIAlarmRuleRequestNotificationSettings {
	return s.NotificationSettings
}

func (s *CreateDIAlarmRuleRequest) GetTriggerConditions() []*CreateDIAlarmRuleRequestTriggerConditions {
	return s.TriggerConditions
}

func (s *CreateDIAlarmRuleRequest) SetClientToken(v string) *CreateDIAlarmRuleRequest {
	s.ClientToken = &v
	return s
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

func (s *CreateDIAlarmRuleRequest) SetName(v string) *CreateDIAlarmRuleRequest {
	s.Name = &v
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
	// Deprecated
	//
	// This parameter is deprecated and replaced by the MuteInterval parameter.
	//
	// example:
	//
	// 5
	InhibitionInterval *int32 `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	// The duration of the alert suppression interval. Default value: 5. Unit: minutes.
	//
	// example:
	//
	// 5
	MuteInterval *int32 `json:"MuteInterval,omitempty" xml:"MuteInterval,omitempty"`
	// The alert notification methods.
	NotificationChannels []*CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The settings of alert notification recipients.
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

func (s *CreateDIAlarmRuleRequestNotificationSettings) GetMuteInterval() *int32 {
	return s.MuteInterval
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

func (s *CreateDIAlarmRuleRequestNotificationSettings) SetMuteInterval(v int32) *CreateDIAlarmRuleRequestNotificationSettings {
	s.MuteInterval = &v
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
	// The alert notification method. Valid values:
	//
	// 	- Mail
	//
	// 	- Phone
	//
	// 	- Sms
	//
	// 	- Ding
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
	// The recipient type. Valid values: AliyunUid, DingToken, FeishuToken, and WebHookUrl.
	//
	// example:
	//
	// DingToken
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The recipient.
	//
	// 	- If the ReceiverType parameter is set to AliyunUid, set this parameter to the Alibaba Cloud account ID of a user.
	//
	// 	- If the ReceiverType parameter is set to DingToken, set this parameter to the token of a DingTalk chatbot.
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
	// Deprecated
	//
	// This parameter is deprecated and replaced by the DdlTypes parameter.
	DdlReportTags []*string `json:"DdlReportTags,omitempty" xml:"DdlReportTags,omitempty" type:"Repeated"`
	// The types of DDL operations for which the alert rule takes effect.
	DdlTypes []*string `json:"DdlTypes,omitempty" xml:"DdlTypes,omitempty" type:"Repeated"`
	// The time interval for alert calculation. Unit: minutes.
	//
	// example:
	//
	// 10
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
	// 10
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateDIAlarmRuleRequestTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleRequestTriggerConditions) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) GetDdlReportTags() []*string {
	return s.DdlReportTags
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) GetDdlTypes() []*string {
	return s.DdlTypes
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

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetDdlReportTags(v []*string) *CreateDIAlarmRuleRequestTriggerConditions {
	s.DdlReportTags = v
	return s
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetDdlTypes(v []*string) *CreateDIAlarmRuleRequestTriggerConditions {
	s.DdlTypes = v
	return s
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
