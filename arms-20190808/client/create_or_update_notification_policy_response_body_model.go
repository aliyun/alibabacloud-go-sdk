// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateNotificationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationPolicy(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) *CreateOrUpdateNotificationPolicyResponseBody
	GetNotificationPolicy() *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy
	SetRequestId(v string) *CreateOrUpdateNotificationPolicyResponseBody
	GetRequestId() *string
}

type CreateOrUpdateNotificationPolicyResponseBody struct {
	// An array of notification policy objects.
	NotificationPolicy *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy `json:"NotificationPolicy,omitempty" xml:"NotificationPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBody) GetNotificationPolicy() *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	return s.NotificationPolicy
}

func (s *CreateOrUpdateNotificationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateNotificationPolicyResponseBody) SetNotificationPolicy(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) *CreateOrUpdateNotificationPolicyResponseBody {
	s.NotificationPolicy = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBody) SetRequestId(v string) *CreateOrUpdateNotificationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy struct {
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
	GroupRule *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule `json:"GroupRule,omitempty" xml:"GroupRule,omitempty" type:"Struct"`
	// The ID of the notification policy.
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
	// The matching rules.
	MatchingRules []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty" type:"Repeated"`
	// The name of the notification policy.
	//
	// example:
	//
	// notificationpolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array of notification rule objects.
	NotifyRule *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule `json:"NotifyRule,omitempty" xml:"NotifyRule,omitempty" type:"Struct"`
	// The notification template.
	NotifyTemplate *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate `json:"NotifyTemplate,omitempty" xml:"NotifyTemplate,omitempty" type:"Struct"`
	// Indicates whether a notification is resent for a long-lasting unresolved alert. Default value: true. Valid values:
	//
	// 	- `true`: The system resends a notification for a long-lasting unresolved alert at a specified time interval.
	//
	// 	- `false`: The system sends a notification for a long-lasting unresolved alert based on an escalation policy.
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
	// Indicates whether the status of an alert automatically changes to Resolved when all events related to the alert change to the Restored state. ARMS notifies contacts when the alert status changes to Resolved.
	//
	// 	- `true`: The system sends a notification.
	//
	// 	- `false`: The system does not send a notification.
	//
	// example:
	//
	// true
	SendRecoverMessage *bool `json:"SendRecoverMessage,omitempty" xml:"SendRecoverMessage,omitempty"`
	// Indicates whether the notification policy is enabled. Valid values: enable and disable.
	//
	// example:
	//
	// enable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetDirectedMode() *bool {
	return s.DirectedMode
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetEscalationPolicyId() *int64 {
	return s.EscalationPolicyId
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetGroupRule() *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule {
	return s.GroupRule
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetIntegrationId() *int64 {
	return s.IntegrationId
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetMatchingRules() []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules {
	return s.MatchingRules
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetNotifyRule() *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	return s.NotifyRule
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetNotifyTemplate() *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	return s.NotifyTemplate
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetRepeat() *bool {
	return s.Repeat
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetSendRecoverMessage() *bool {
	return s.SendRecoverMessage
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GetState() *string {
	return s.State
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetDirectedMode(v bool) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.DirectedMode = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetEscalationPolicyId(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.EscalationPolicyId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetGroupRule(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.GroupRule = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetId(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetIntegrationId(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.IntegrationId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetMatchingRules(v []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.MatchingRules = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetName(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetNotifyRule(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.NotifyRule = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetNotifyTemplate(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.NotifyTemplate = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetRepeat(v bool) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.Repeat = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetRepeatInterval(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.RepeatInterval = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetSendRecoverMessage(v bool) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.SendRecoverMessage = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetState(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.State = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule struct {
	// The time interval of grouping. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 30
	GroupInterval *int64 `json:"GroupInterval,omitempty" xml:"GroupInterval,omitempty"`
	// The waiting time for grouping. Unit: seconds. Default value: 5.
	//
	// example:
	//
	// 5
	GroupWait *int64 `json:"GroupWait,omitempty" xml:"GroupWait,omitempty"`
	// An array of alert event group objects.
	//
	// 	- If you do not specify the groupingFields field, all alerts will be sent to contacts based on `alertname`.
	//
	// 	- If you specify the groupingFields field, alerts with the same field will be sent to contacts in one notification.
	GroupingFields []*string `json:"GroupingFields,omitempty" xml:"GroupingFields,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) GetGroupInterval() *int64 {
	return s.GroupInterval
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) GetGroupWait() *int64 {
	return s.GroupWait
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) GetGroupingFields() []*string {
	return s.GroupingFields
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) SetGroupInterval(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule {
	s.GroupInterval = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) SetGroupWait(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule {
	s.GroupWait = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) SetGroupingFields(v []*string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule {
	s.GroupingFields = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules struct {
	// The matching conditions.
	MatchingConditions []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) GetMatchingConditions() []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions {
	return s.MatchingConditions
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) SetMatchingConditions(v []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules {
	s.MatchingConditions = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions struct {
	// The key of the matching condition.
	//
	// example:
	//
	// altertname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The logical operator of the matching condition. Valid values:
	//
	// 	- `eq`: equal to
	//
	// 	- `neq`: not equal to
	//
	// 	- `in`: contains
	//
	// 	- `nin`: does not contain
	//
	// 	- `re`: regular expression match
	//
	// 	- `nre`: regular expression mismatch
	//
	// example:
	//
	// eq
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the matching condition.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) GetKey() *string {
	return s.Key
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) GetOperator() *string {
	return s.Operator
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) SetKey(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) SetOperator(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) SetValue(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule struct {
	// The notification method.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The end time of the notification window.
	//
	// example:
	//
	// 23:59
	NotifyEndTime *string `json:"NotifyEndTime,omitempty" xml:"NotifyEndTime,omitempty"`
	// An array of notification objects.
	NotifyObjects []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
	// The start time of the notification window.
	//
	// example:
	//
	// 00:00
	NotifyStartTime *string `json:"NotifyStartTime,omitempty" xml:"NotifyStartTime,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) GetNotifyEndTime() *string {
	return s.NotifyEndTime
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) GetNotifyObjects() []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects {
	return s.NotifyObjects
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) GetNotifyStartTime() *string {
	return s.NotifyStartTime
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) SetNotifyChannels(v []*string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	s.NotifyChannels = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) SetNotifyEndTime(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	s.NotifyEndTime = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) SetNotifyObjects(v []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	s.NotifyObjects = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) SetNotifyStartTime(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	s.NotifyStartTime = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects struct {
	// The notification methods specified for a contact.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The ID of the notification object.
	//
	// example:
	//
	// 123
	NotifyObjectId *int64 `json:"NotifyObjectId,omitempty" xml:"NotifyObjectId,omitempty"`
	// The name of the notification object.
	//
	// example:
	//
	// test
	NotifyObjectName *string `json:"NotifyObjectName,omitempty" xml:"NotifyObjectName,omitempty"`
	// The type of the notification object. Valid values:
	//
	// 	- CONTACT: contact
	//
	// 	- CONTACT_GROUP: contact group
	//
	// 	- ARMS_CONTACT: ARMS contact
	//
	// 	- ARMS_CONTACT_GROUP: ARMS contact group
	//
	// 	- DING_ROBOT_GROUP: DingTalk, Lark, WeCom, or IM robot
	//
	// 	- CONTACT_SCHEDULE: user on duty defined by a schedule
	//
	// example:
	//
	// CONTACT
	NotifyObjectType *string `json:"NotifyObjectType,omitempty" xml:"NotifyObjectType,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) GetNotifyObjectId() *int64 {
	return s.NotifyObjectId
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) GetNotifyObjectName() *string {
	return s.NotifyObjectName
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) GetNotifyObjectType() *string {
	return s.NotifyObjectType
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) SetNotifyChannels(v []*string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects {
	s.NotifyChannels = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) SetNotifyObjectId(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects {
	s.NotifyObjectId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) SetNotifyObjectName(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects {
	s.NotifyObjectName = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) SetNotifyObjectType(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects {
	s.NotifyObjectType = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate struct {
	// The content of the alert notification sent through email.
	//
	// example:
	//
	// Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert time: {{ .startTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{if .generatorURL }} \\<a href="{{.generatorURL}}" >Link\\</a> {{end}} {{end}}
	EmailContent *string `json:"EmailContent,omitempty" xml:"EmailContent,omitempty"`
	// The content of the alert resolution notification sent through email.
	//
	// example:
	//
	// Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert resolution time: {{ .endTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{if .generatorURL }} \\<a href="{{.generatorURL}}" >Link\\</a> {{end}} {{end}}
	EmailRecoverContent *string `json:"EmailRecoverContent,omitempty" xml:"EmailRecoverContent,omitempty"`
	// The title of the alert resolution notification sent through email.
	//
	// example:
	//
	// {{ .commonLabels.alertname }}
	EmailRecoverTitle *string `json:"EmailRecoverTitle,omitempty" xml:"EmailRecoverTitle,omitempty"`
	// The title of the alert notification sent through email.
	//
	// example:
	//
	// {{ .commonLabels.alertname }}
	EmailTitle *string `json:"EmailTitle,omitempty" xml:"EmailTitle,omitempty"`
	// The content of the alert notification sent by the IM robot.
	//
	// example:
	//
	// {{if .commonLabels.clustername }} > Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} > Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}{{ for .alerts }}> {{.annotations.message}} {{if .generatorURL }} [Link]\\({{.generatorURL}}) {{ end }} {{if eq "true" .labels._aliyun_arms_is_denoise_filtered }} (Suspected noise) {{end}} {{end}}
	RobotContent *string `json:"RobotContent,omitempty" xml:"RobotContent,omitempty"`
	// The content of the alert notification sent through text message.
	//
	// example:
	//
	// \\<SmsContent>Notification on the occurrence of a {{ .level }} alert. Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert time: {{ .startTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{ end }}\\</SmsContent>
	SmsContent *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	// The content of the alert resolution notification sent through text message.
	//
	// example:
	//
	// \\<SmsRecoverContent>Alert resolution notification. Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert resolution time: {{ .endTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{ end }}\\</SmsRecoverContent>
	SmsRecoverContent *string `json:"SmsRecoverContent,omitempty" xml:"SmsRecoverContent,omitempty"`
	// The content of the alert notification by phone.
	//
	// example:
	//
	// \\<TtsContent>Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert time: {{ .startTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{ end }}\\</TtsContent>
	TtsContent *string `json:"TtsContent,omitempty" xml:"TtsContent,omitempty"`
	// The content of the alert resolution notification by phone.
	//
	// example:
	//
	// \\<TtsRecoverContent>Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert resolution time: {{ .endTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{ end }}\\</TtsRecoverContent>
	TtsRecoverContent *string `json:"TtsRecoverContent,omitempty" xml:"TtsRecoverContent,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetEmailContent() *string {
	return s.EmailContent
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetEmailRecoverContent() *string {
	return s.EmailRecoverContent
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetEmailRecoverTitle() *string {
	return s.EmailRecoverTitle
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetEmailTitle() *string {
	return s.EmailTitle
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetRobotContent() *string {
	return s.RobotContent
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetSmsContent() *string {
	return s.SmsContent
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetSmsRecoverContent() *string {
	return s.SmsRecoverContent
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetTtsContent() *string {
	return s.TtsContent
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GetTtsRecoverContent() *string {
	return s.TtsRecoverContent
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetEmailContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.EmailContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetEmailRecoverContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.EmailRecoverContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetEmailRecoverTitle(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.EmailRecoverTitle = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetEmailTitle(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.EmailTitle = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetRobotContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.RobotContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetSmsContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.SmsContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetSmsRecoverContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.SmsRecoverContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetTtsContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.TtsContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetTtsRecoverContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.TtsRecoverContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) Validate() error {
	return dara.Validate(s)
}
