// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEventSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v string) *SetEventSubscriptionRequest
	GetActive() *string
	SetChannelType(v string) *SetEventSubscriptionRequest
	GetChannelType() *string
	SetContactGroupName(v string) *SetEventSubscriptionRequest
	GetContactGroupName() *string
	SetContactName(v string) *SetEventSubscriptionRequest
	GetContactName() *string
	SetDispatchRule(v string) *SetEventSubscriptionRequest
	GetDispatchRule() *string
	SetEventContext(v string) *SetEventSubscriptionRequest
	GetEventContext() *string
	SetInstanceId(v string) *SetEventSubscriptionRequest
	GetInstanceId() *string
	SetLang(v string) *SetEventSubscriptionRequest
	GetLang() *string
	SetLevel(v string) *SetEventSubscriptionRequest
	GetLevel() *string
	SetMinInterval(v string) *SetEventSubscriptionRequest
	GetMinInterval() *string
	SetSeverity(v string) *SetEventSubscriptionRequest
	GetSeverity() *string
}

type SetEventSubscriptionRequest struct {
	// Specifies whether to enable the event subscription feature. Valid values:
	//
	// 	- **0**: disables the event subscription feature.
	//
	// 	- **1**: enables the event subscription feature.
	//
	// example:
	//
	// 1
	Active *string `json:"Active,omitempty" xml:"Active,omitempty"`
	// The notification method. Valid values:
	//
	// 	- **hdm_alarm_sms**: text message.
	//
	// 	- **dingtalk**: DingTalk chatbot.
	//
	// 	- **hdm_alarm_sms_and_email**: text message and email.
	//
	// 	- **hdm_alarm_sms,dingtalk**: text message and DingTalk chatbot.
	//
	// example:
	//
	// hdm_alarm_sms,dingtalk
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The name of the contact group that receives alert notifications. Separate multiple names with commas (,).
	//
	// example:
	//
	// Default contact group
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The name of the contact who receives alert notifications. Separate multiple names with commas (,).
	//
	// example:
	//
	// Default contact
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The notification rules based on the event type. If you leave this parameter empty, the values of **MinInterval*	- and **ChannelType*	- prevail.
	//
	// Specify this parameter in the following format: `{"silenced": {"Event type 1":Specifies whether to enable adaptive silence, "Event type 2":Specify whether to enable adaptive silence},"min_interval": {"Event type 1":Minimum interval between event notifications, "Event type 2":Minimum interval between event notifications},"alert_type": {"Event type 1":"Notification method", "Event type 2":"Notification method"}}`.
	//
	// 	- **silenced**: specifies whether to enable adaptive silence. After you enable adaptive silence, the interval between consecutive alert notifications for an event is the greater one of the minimum interval specified by **min_interval*	- and one third of the event duration. Valid values:
	//
	//     	- 1: enables adaptive silence.
	//
	//     	- 2: disables adaptive silence.
	//
	// 	- **min_interval**: the minimum interval between event notifications. Unit: seconds.
	//
	// 	- **alert_type**: the notification method. Valid values:
	//
	//     	- **hdm_alarm_sms**: text message.
	//
	//     	- **dingtalk**: DingTalk chatbot.
	//
	//     	- **hdm_alarm_sms_and_email**: text message and email.
	//
	//     	- **hdm_alarm_sms,dingtalk**: text message and DingTalk chatbot.
	//
	// example:
	//
	// {"silenced": {"AutoScale":1, "SQLThrottle":0, "TimeSeriesAbnormal": 1}, "min_interval": {"AutoScale":300, "SQLThrottle":360, "TimeSeriesAbnormal": 120}, "alert_type": {"AutoScale":"hdm_alarm_sms", "SQLThrottle":"hdm_alarm_sms_and_email", "TimeSeriesAbnormal": "hdm_alarm_sms,dingtalk"}}
	DispatchRule *string `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty"`
	// The supported event scenarios. You can set the value to **AllContext**, which indicates that all scenarios are supported.
	//
	// example:
	//
	// AllContext
	EventContext *string `json:"EventContext,omitempty" xml:"EventContext,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of event notifications. You can set the value to **zh-CN**, which indicates that event notifications are sent in Chinese.
	//
	// example:
	//
	// zh-CN
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The risk level of the events. Valid values:
	//
	// 	- **Notice**: events that trigger notifications, including events at the **Notice**, **Optimization**, **Warn**, and **Critical*	- levels.
	//
	// 	- **Optimization**: events that trigger optimizations, including events at the **Optimization**, **Warn**, and **Critical*	- levels.
	//
	// 	- **Warn**: events that trigger warnings, including events at the **Warn*	- and **Critical*	- levels.
	//
	// 	- **Critical**: events that trigger critical warnings.
	//
	// The following content describes the events at each level in detail:
	//
	// 	- Notice: events that are related to database exceptions for which no suggestions are generated.
	//
	// 	- Optimization: events for which optimization suggestions are generated based on the status of the database.
	//
	// 	- Warn: events that may affect the running of the database.
	//
	// 	- Critical: events that affect the running of the database.
	//
	// example:
	//
	// Optimization
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The minimum interval between consecutive event notifications. Unit: seconds.
	//
	// example:
	//
	// 60
	MinInterval *string `json:"MinInterval,omitempty" xml:"MinInterval,omitempty"`
	// The alert severity based on the event type.
	//
	// Specify this parameter in the following format: `{"Event type 1":"Alert severity", "Event type 2":"Alert severity"}`.
	//
	// Valid values of event types:
	//
	// 	- **AutoScale**: auto scaling event.
	//
	// 	- **SQLThrottle**: throttling event.
	//
	// 	- **TimeSeriesAbnormal**: event for detecting time series anomalies.
	//
	// 	- **SQLOptimize**: SQL optimization event.
	//
	// 	- **ResourceOptimize**: storage optimization event.
	//
	// Valid values of alert severities:
	//
	// 	- **info**
	//
	// 	- **noticed**
	//
	// 	- **warning**
	//
	// 	- **critical**
	//
	// example:
	//
	// {"AutoScale":"critical","SQLThrottle":"info","TimeSeriesAbnormal":"warning"}
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s SetEventSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetEventSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *SetEventSubscriptionRequest) GetActive() *string {
	return s.Active
}

func (s *SetEventSubscriptionRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *SetEventSubscriptionRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *SetEventSubscriptionRequest) GetContactName() *string {
	return s.ContactName
}

func (s *SetEventSubscriptionRequest) GetDispatchRule() *string {
	return s.DispatchRule
}

func (s *SetEventSubscriptionRequest) GetEventContext() *string {
	return s.EventContext
}

func (s *SetEventSubscriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetEventSubscriptionRequest) GetLang() *string {
	return s.Lang
}

func (s *SetEventSubscriptionRequest) GetLevel() *string {
	return s.Level
}

func (s *SetEventSubscriptionRequest) GetMinInterval() *string {
	return s.MinInterval
}

func (s *SetEventSubscriptionRequest) GetSeverity() *string {
	return s.Severity
}

func (s *SetEventSubscriptionRequest) SetActive(v string) *SetEventSubscriptionRequest {
	s.Active = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetChannelType(v string) *SetEventSubscriptionRequest {
	s.ChannelType = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetContactGroupName(v string) *SetEventSubscriptionRequest {
	s.ContactGroupName = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetContactName(v string) *SetEventSubscriptionRequest {
	s.ContactName = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetDispatchRule(v string) *SetEventSubscriptionRequest {
	s.DispatchRule = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetEventContext(v string) *SetEventSubscriptionRequest {
	s.EventContext = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetInstanceId(v string) *SetEventSubscriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetLang(v string) *SetEventSubscriptionRequest {
	s.Lang = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetLevel(v string) *SetEventSubscriptionRequest {
	s.Level = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetMinInterval(v string) *SetEventSubscriptionRequest {
	s.MinInterval = &v
	return s
}

func (s *SetEventSubscriptionRequest) SetSeverity(v string) *SetEventSubscriptionRequest {
	s.Severity = &v
	return s
}

func (s *SetEventSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
