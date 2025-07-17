// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateNotificationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectedMode(v bool) *CreateOrUpdateNotificationPolicyRequest
	GetDirectedMode() *bool
	SetEscalationPolicyId(v int64) *CreateOrUpdateNotificationPolicyRequest
	GetEscalationPolicyId() *int64
	SetGroupRule(v string) *CreateOrUpdateNotificationPolicyRequest
	GetGroupRule() *string
	SetId(v int64) *CreateOrUpdateNotificationPolicyRequest
	GetId() *int64
	SetIntegrationId(v int64) *CreateOrUpdateNotificationPolicyRequest
	GetIntegrationId() *int64
	SetMatchingRules(v string) *CreateOrUpdateNotificationPolicyRequest
	GetMatchingRules() *string
	SetName(v string) *CreateOrUpdateNotificationPolicyRequest
	GetName() *string
	SetNotifyRule(v string) *CreateOrUpdateNotificationPolicyRequest
	GetNotifyRule() *string
	SetNotifyTemplate(v string) *CreateOrUpdateNotificationPolicyRequest
	GetNotifyTemplate() *string
	SetRegionId(v string) *CreateOrUpdateNotificationPolicyRequest
	GetRegionId() *string
	SetRepeat(v bool) *CreateOrUpdateNotificationPolicyRequest
	GetRepeat() *bool
	SetRepeatInterval(v int64) *CreateOrUpdateNotificationPolicyRequest
	GetRepeatInterval() *int64
	SetSendRecoverMessage(v bool) *CreateOrUpdateNotificationPolicyRequest
	GetSendRecoverMessage() *bool
	SetState(v string) *CreateOrUpdateNotificationPolicyRequest
	GetState() *string
}

type CreateOrUpdateNotificationPolicyRequest struct {
	// Specifies whether to enable simple mode.
	//
	// example:
	//
	// false
	DirectedMode *bool `json:"DirectedMode,omitempty" xml:"DirectedMode,omitempty"`
	// The ID of the escalation policy.
	//
	// example:
	//
	// 123
	EscalationPolicyId *int64 `json:"EscalationPolicyId,omitempty" xml:"EscalationPolicyId,omitempty"`
	// An array of alert event group objects.
	//
	// 	- If you do not specify the groupingFields field, all alerts will be sent to contacts based on `alertname`.
	//
	// 	- If you specify the groupingFields field, alerts with the same field will be sent to contacts in one notification.
	//
	// Sample statement:
	//
	//     {
	//
	//     "groupWait":5,    // The waiting time for grouping.
	//
	//     "groupInterval":30,     // The time interval of grouping.
	//
	//     "groupingFields":["alertname"]       // The field that is used to group alert events.
	//
	//     }
	//
	// example:
	//
	// { 	"groupWait":5, 	"groupInterval":30, 	"groupingFields":["alertname"] }
	GroupRule *string `json:"GroupRule,omitempty" xml:"GroupRule,omitempty"`
	// The ID of the notification policy.
	//
	// 	- If you do not specify this parameter, a new notification policy is created.
	//
	// 	- If you specify this parameter, the specified notification policy is modified.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The integration ID of the ticket system to which alerts are pushed.
	//
	// example:
	//
	// 34
	IntegrationId *int64 `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	// The matching rules. Format:
	//
	//     [
	//
	//      {
	//
	//      "matchingConditions": [
	//
	//      {
	//
	//      "value": "test",    // The value of the matching condition.
	//
	//      "key": "alertname",     // The key of the matching condition.
	//
	//      "operator": "eq"   // The logical operator of the matching condition, including eq (equal to), neq (not equal to), in (contains), nin (does not contain), re (regular expression match), and nre (regular expression mismatch).
	//
	//      }
	//
	//      ]
	//
	//      }
	//
	//      ]
	//
	// example:
	//
	// [ 		 { 		 "matchingConditions": [          { 		 "value": "test", 		 "key": "alertname", 		 "operator": "eq"         }       ]     }   ]
	MatchingRules *string `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty"`
	// The name of the notification policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// notificationpolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array of notification rule objects. Format:
	//
	//     {
	//
	//      "notifyStartTime":"00:00",      // The start time of the notification window.
	//
	//      "notifyEndTime":"23:59",       // The end time of the notification window.
	//
	//      "notifyChannels":["dingTalk", "email", "sms", "tts", "webhook"],       // The notification methods. Valid values: dingTalk, email, sms, tts, and webhook.
	//
	//      "notifyObjects":[{       // An array of notification objects.
	//
	//      "notifyObjectType":"CONTACT",       // The type of the notification object. Valid values: CONTACT (contact), CONTACT_GROUP (contact group), ARMS_CONTACT (ARMS contact), ARMS_CONTACT_GROUP (ARMS contact group), DING_ROBOT_GROUP (DingTalk, Lark, WeCom, or IM robot), and CONTACT_SCHEDULE (user on duty defined by a schedule).
	//
	//      "notifyObjectId":123,       // The ID of the notification object.
	//
	//      "notifyObjectName":"test"       // The name of the notification object.
	//
	//      }]
	//
	// This parameter is required.
	//
	// example:
	//
	// {     "notifyStartTime":"00:00",     "notifyEndTime":"23:59",     "notifyChannels":[         "dingTalk",         "email",         "sms",         "tts",         "webhook"     ],     "notifyObjects":[         {             "notifyObjectType":"CONTACT",             "notifyObjectId":123,             "notifyObjectName":"test"         }     ] }
	NotifyRule *string `json:"NotifyRule,omitempty" xml:"NotifyRule,omitempty"`
	// The notification template. The default notification template is provided below the table.
	//
	// example:
	//
	// "robotContent":"{{if .commonLabels.clustername }} > Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} > Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}{{ for .alerts }} > {{.annotations.message}} {{if .generatorURL }} [Link]\\({{.generatorURL}}) {{ end }} {{if eq "true" .labels._aliyun_arms_is_denoise_filtered }} (Suspected noise) {{end}} {{end}}"
	NotifyTemplate *string `json:"NotifyTemplate,omitempty" xml:"NotifyTemplate,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to resend a notification for a long-lasting unresolved alert. Default value: true. Valid values:
	//
	// 	- `true`: If you set this parameter to `true`, you must set **RepeatInterval**.
	//
	// 	- `false`: If you set this parameter to `false`, you must set **EscalationPolicyId**.
	//
	// example:
	//
	// true
	Repeat *bool `json:"Repeat,omitempty" xml:"Repeat,omitempty"`
	// The time interval at which a notification is resent for a long-lasting unresolved alert. Unit: seconds.
	//
	// example:
	//
	// 600
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// Specifies whether the status of an alert automatically changes to Resolved when all events related to the alert change to the Restored state. ARMS notifies contacts when the alert status changes to Resolved.
	//
	// 	- `true`: The system sends a notification.
	//
	// 	- `false`: The system does not send a notification.
	//
	// example:
	//
	// true
	SendRecoverMessage *bool `json:"SendRecoverMessage,omitempty" xml:"SendRecoverMessage,omitempty"`
	// Specifies whether to enable the notification policy. Valid values: enable and disable.
	//
	// example:
	//
	// enable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetDirectedMode() *bool {
	return s.DirectedMode
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetEscalationPolicyId() *int64 {
	return s.EscalationPolicyId
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetGroupRule() *string {
	return s.GroupRule
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetIntegrationId() *int64 {
	return s.IntegrationId
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetMatchingRules() *string {
	return s.MatchingRules
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetNotifyRule() *string {
	return s.NotifyRule
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetNotifyTemplate() *string {
	return s.NotifyTemplate
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetRepeat() *bool {
	return s.Repeat
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetSendRecoverMessage() *bool {
	return s.SendRecoverMessage
}

func (s *CreateOrUpdateNotificationPolicyRequest) GetState() *string {
	return s.State
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetDirectedMode(v bool) *CreateOrUpdateNotificationPolicyRequest {
	s.DirectedMode = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetEscalationPolicyId(v int64) *CreateOrUpdateNotificationPolicyRequest {
	s.EscalationPolicyId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetGroupRule(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.GroupRule = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetId(v int64) *CreateOrUpdateNotificationPolicyRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetIntegrationId(v int64) *CreateOrUpdateNotificationPolicyRequest {
	s.IntegrationId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetMatchingRules(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.MatchingRules = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetName(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetNotifyRule(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.NotifyRule = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetNotifyTemplate(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.NotifyTemplate = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetRegionId(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetRepeat(v bool) *CreateOrUpdateNotificationPolicyRequest {
	s.Repeat = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetRepeatInterval(v int64) *CreateOrUpdateNotificationPolicyRequest {
	s.RepeatInterval = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetSendRecoverMessage(v bool) *CreateOrUpdateNotificationPolicyRequest {
	s.SendRecoverMessage = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetState(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.State = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
