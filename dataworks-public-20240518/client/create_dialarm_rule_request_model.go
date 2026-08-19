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
	// The idempotency parameter.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The task ID associated with the alert rule.
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
	// Alert description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the alert rule. By default, the alert rule is disabled.
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
	// - Delay: task latency alert.
	//
	// - DdlReport: DDL notification.
	//
	// - ResourceUtilization: resource group utilization.
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
	// The list of alert trigger conditions. Multiple conditions are supported.
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

type CreateDIAlarmRuleRequestNotificationSettings struct {
	// Deprecated
	//
	// **[Deprecated]*	- Use the MuteInterval parameter instead.
	//
	// example:
	//
	// 5
	InhibitionInterval *int32 `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	// The alert mute interval. Unit: minutes. Default value: 5.
	//
	// example:
	//
	// 5
	MuteInterval *int32 `json:"MuteInterval,omitempty" xml:"MuteInterval,omitempty"`
	// The alert notification channels.
	NotificationChannels []*CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The alert notification receivers.
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

type CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels struct {
	// The notification channel. Valid values:
	//
	// - Mail: email.
	//
	// - Phone: phone call.
	//
	// - Sms: text message.
	//
	// - Ding: DingTalk.
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
	// The receiver type. Valid values: AliyunUid, DingToken, FeishuToken, and WebHookUrl.
	//
	// example:
	//
	// DingToken
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The receiver values.
	//
	// - If the receiver type is AliyunUid, the value is the Alibaba Cloud account ID.
	//
	// - If the receiver type is DingToken, the value is the DingTalk token.
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
	// **[Deprecated]*	- Use the DdlTypes parameter instead.
	DdlReportTags []*string `json:"DdlReportTags,omitempty" xml:"DdlReportTags,omitempty" type:"Repeated"`
	// The list of DDL types that take effect. This parameter takes effect only when the metric type is DDL notification.
	DdlTypes []*string `json:"DdlTypes,omitempty" xml:"DdlTypes,omitempty" type:"Repeated"`
	// The time window for alert calculation. Unit: minutes.
	//
	// example:
	//
	// 10
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
	// - Task status alert: no threshold is required.
	//
	// - Failover count alert: the threshold is the number of failovers.
	//
	// - Task latency alert: the threshold is the latency duration. Unit: seconds.
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
