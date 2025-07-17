// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNotificationPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *ListNotificationPoliciesResponseBodyPageBean) *ListNotificationPoliciesResponseBody
	GetPageBean() *ListNotificationPoliciesResponseBodyPageBean
	SetRequestId(v string) *ListNotificationPoliciesResponseBody
	GetRequestId() *string
}

type ListNotificationPoliciesResponseBody struct {
	// The returned pages.
	PageBean *ListNotificationPoliciesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNotificationPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBody) GetPageBean() *ListNotificationPoliciesResponseBodyPageBean {
	return s.PageBean
}

func (s *ListNotificationPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNotificationPoliciesResponseBody) SetPageBean(v *ListNotificationPoliciesResponseBodyPageBean) *ListNotificationPoliciesResponseBody {
	s.PageBean = v
	return s
}

func (s *ListNotificationPoliciesResponseBody) SetRequestId(v string) *ListNotificationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNotificationPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNotificationPoliciesResponseBodyPageBean struct {
	// The queried notification policies.
	NotificationPolicies []*ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies `json:"NotificationPolicies,omitempty" xml:"NotificationPolicies,omitempty" type:"Repeated"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries that are returned on each page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The number of notification policies that are returned.
	//
	// example:
	//
	// 24
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBean) GetNotificationPolicies() []*ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	return s.NotificationPolicies
}

func (s *ListNotificationPoliciesResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *ListNotificationPoliciesResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *ListNotificationPoliciesResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *ListNotificationPoliciesResponseBodyPageBean) SetNotificationPolicies(v []*ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) *ListNotificationPoliciesResponseBodyPageBean {
	s.NotificationPolicies = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBean) SetPage(v int64) *ListNotificationPoliciesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBean) SetSize(v int64) *ListNotificationPoliciesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBean) SetTotal(v int64) *ListNotificationPoliciesResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies struct {
	// Indicates whether simple mode is enabled.
	//
	// example:
	//
	// true
	DirectedMode *bool `json:"DirectedMode,omitempty" xml:"DirectedMode,omitempty"`
	// The ID of the escalation policy.
	//
	// example:
	//
	// 123
	EscalationPolicyId *int64 `json:"EscalationPolicyId,omitempty" xml:"EscalationPolicyId,omitempty"`
	// The grouping rule for alert events.
	GroupRule *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule `json:"GroupRule,omitempty" xml:"GroupRule,omitempty" type:"Struct"`
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
	// The matching rules for alert events.
	MatchingRules []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty" type:"Repeated"`
	// The name of the notification policy.
	//
	// example:
	//
	// notificationpolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The notification rule.
	NotifyRule *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule `json:"NotifyRule,omitempty" xml:"NotifyRule,omitempty" type:"Struct"`
	// The notification template.
	NotifyTemplate *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate `json:"NotifyTemplate,omitempty" xml:"NotifyTemplate,omitempty" type:"Struct"`
	// Indicates whether the system resends notifications for a long-lasting unresolved alert. Valid values:
	//
	// - `true` (default): The system resends notifications for a long-lasting unresolved alert at a specified time interval.
	//
	// - `false`: The system resends notifications for a long-lasting unresolved alert based on an escalation policy.
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
	// Indicates whether the status of an alert automatically changes to Resolved when all events related to the alert change to the Restored state. The system sends a notification to the alert contacts when the alert status changes to Resolved.
	//
	// - `true` (default): The system sends a notification.
	//
	// - `false`: The system does not send a notification.
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

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetDirectedMode() *bool {
	return s.DirectedMode
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetEscalationPolicyId() *int64 {
	return s.EscalationPolicyId
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetGroupRule() *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule {
	return s.GroupRule
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetId() *int64 {
	return s.Id
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetIntegrationId() *int64 {
	return s.IntegrationId
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetMatchingRules() []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules {
	return s.MatchingRules
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetName() *string {
	return s.Name
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetNotifyRule() *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	return s.NotifyRule
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetNotifyTemplate() *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	return s.NotifyTemplate
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetRepeat() *bool {
	return s.Repeat
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetSendRecoverMessage() *bool {
	return s.SendRecoverMessage
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GetState() *string {
	return s.State
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetDirectedMode(v bool) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.DirectedMode = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetEscalationPolicyId(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.EscalationPolicyId = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetGroupRule(v *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.GroupRule = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetId(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.Id = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetIntegrationId(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.IntegrationId = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetMatchingRules(v []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.MatchingRules = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetName(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.Name = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetNotifyRule(v *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.NotifyRule = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetNotifyTemplate(v *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.NotifyTemplate = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetRepeat(v bool) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.Repeat = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetRepeatInterval(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.RepeatInterval = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetSendRecoverMessage(v bool) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.SendRecoverMessage = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetState(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.State = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) Validate() error {
	return dara.Validate(s)
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule struct {
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

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) GetGroupInterval() *int64 {
	return s.GroupInterval
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) GetGroupWait() *int64 {
	return s.GroupWait
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) GetGroupingFields() []*string {
	return s.GroupingFields
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) SetGroupInterval(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule {
	s.GroupInterval = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) SetGroupWait(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule {
	s.GroupWait = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) SetGroupingFields(v []*string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule {
	s.GroupingFields = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) Validate() error {
	return dara.Validate(s)
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules struct {
	// The matching conditions.
	MatchingConditions []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty" type:"Repeated"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) GetMatchingConditions() []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions {
	return s.MatchingConditions
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) SetMatchingConditions(v []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules {
	s.MatchingConditions = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) Validate() error {
	return dara.Validate(s)
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions struct {
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

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) GetKey() *string {
	return s.Key
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) GetOperator() *string {
	return s.Operator
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) GetValue() *string {
	return s.Value
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) SetKey(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions {
	s.Key = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) SetOperator(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions {
	s.Operator = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) SetValue(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions {
	s.Value = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) Validate() error {
	return dara.Validate(s)
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule struct {
	// The notification method.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The end time of the notification window.
	//
	// example:
	//
	// 23:59
	NotifyEndTime *string `json:"NotifyEndTime,omitempty" xml:"NotifyEndTime,omitempty"`
	// The notification objects.
	NotifyObjects []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
	// The start time of the notification window.
	//
	// example:
	//
	// 00:00
	NotifyStartTime *string `json:"NotifyStartTime,omitempty" xml:"NotifyStartTime,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) GetNotifyEndTime() *string {
	return s.NotifyEndTime
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) GetNotifyObjects() []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects {
	return s.NotifyObjects
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) GetNotifyStartTime() *string {
	return s.NotifyStartTime
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) SetNotifyChannels(v []*string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	s.NotifyChannels = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) SetNotifyEndTime(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	s.NotifyEndTime = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) SetNotifyObjects(v []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	s.NotifyObjects = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) SetNotifyStartTime(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	s.NotifyStartTime = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) Validate() error {
	return dara.Validate(s)
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects struct {
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
	// - CONTACT: an individual contact
	//
	// - CONTACT_GROUP: a contact group
	//
	// - DING_ROBOT: an instant messaging (IM) chatbot
	//
	// - CONTACT_SCHEDULE: a person on duty based on an established schedule
	//
	// example:
	//
	// CONTACT
	NotifyObjectType *string `json:"NotifyObjectType,omitempty" xml:"NotifyObjectType,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) GetNotifyObjectId() *int64 {
	return s.NotifyObjectId
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) GetNotifyObjectName() *string {
	return s.NotifyObjectName
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) GetNotifyObjectType() *string {
	return s.NotifyObjectType
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) SetNotifyChannels(v []*string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects {
	s.NotifyChannels = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) SetNotifyObjectId(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects {
	s.NotifyObjectId = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) SetNotifyObjectName(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects {
	s.NotifyObjectName = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) SetNotifyObjectType(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects {
	s.NotifyObjectType = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) Validate() error {
	return dara.Validate(s)
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate struct {
	// The content of the alert notification sent by email.
	//
	// example:
	//
	// Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert time: {{ .startTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{if .generatorURL }} \\<a href="{{.generatorURL}}" >Link\\</a> {{end}} {{end}}
	EmailContent *string `json:"EmailContent,omitempty" xml:"EmailContent,omitempty"`
	// The content of the alert resolution notification sent by email.
	//
	// example:
	//
	// Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert resolution time: {{ .endTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{if .generatorURL }} \\<a href="{{.generatorURL}}" >Link\\</a> {{end}} {{end}}
	EmailRecoverContent *string `json:"EmailRecoverContent,omitempty" xml:"EmailRecoverContent,omitempty"`
	// The title of the alert resolution notification sent by email.
	//
	// example:
	//
	// {{ .commonLabels.alertname }}
	EmailRecoverTitle *string `json:"EmailRecoverTitle,omitempty" xml:"EmailRecoverTitle,omitempty"`
	// The title of the alert notification sent by email.
	//
	// example:
	//
	// {{ .commonLabels.alertname }}
	EmailTitle *string `json:"EmailTitle,omitempty" xml:"EmailTitle,omitempty"`
	// The content of the alert notification sent by an IM chatbot.
	//
	// example:
	//
	// {{if .commonLabels.clustername }} > Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} > Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}{{ for .alerts }}> {{.annotations.message}} {{if .generatorURL }} [Link]\\({{.generatorURL}}) {{ end }} {{if eq "true" .labels._aliyun_arms_is_denoise_filtered }} (Suspected noise) {{end}} {{end}}
	RobotContent *string `json:"RobotContent,omitempty" xml:"RobotContent,omitempty"`
	// The content of the alert notification sent by text message.
	//
	// example:
	//
	// \\<SmsContent>Notification on the occurrence of a {{ .level }} alert. Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert time: {{ .startTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{ end }}\\</SmsContent>
	SmsContent *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	// The content of the alert resolution notification sent by text message.
	//
	// example:
	//
	// \\<SmsRecoverContent>Alert resolution notification. Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert resolution time: {{ .endTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{ end }}\\</SmsRecoverContent>
	SmsRecoverContent *string `json:"SmsRecoverContent,omitempty" xml:"SmsRecoverContent,omitempty"`
	// The content of the alert notification sent by phone.
	//
	// example:
	//
	// \\<TtsContent>Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert time: {{ .startTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{ end }}\\</TtsContent>
	TtsContent *string `json:"TtsContent,omitempty" xml:"TtsContent,omitempty"`
	// The content of the alert resolution notification sent by phone.
	//
	// example:
	//
	// \\<TtsRecoverContent>Alert name: {{ .commonLabels.alertname }}{{if .commonLabels.clustername }} Cluster name: {{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }} Application name: {{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }} Notification policy: {{ .dispatchRuleName }} Alert resolution time: {{ .endTime }} Alert content: {{ for .alerts }} {{.annotations.message}} {{ end }}\\</TtsRecoverContent>
	TtsRecoverContent *string `json:"TtsRecoverContent,omitempty" xml:"TtsRecoverContent,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetEmailContent() *string {
	return s.EmailContent
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetEmailRecoverContent() *string {
	return s.EmailRecoverContent
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetEmailRecoverTitle() *string {
	return s.EmailRecoverTitle
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetEmailTitle() *string {
	return s.EmailTitle
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetRobotContent() *string {
	return s.RobotContent
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetSmsContent() *string {
	return s.SmsContent
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetSmsRecoverContent() *string {
	return s.SmsRecoverContent
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetTtsContent() *string {
	return s.TtsContent
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GetTtsRecoverContent() *string {
	return s.TtsRecoverContent
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetEmailContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.EmailContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetEmailRecoverContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.EmailRecoverContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetEmailRecoverTitle(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.EmailRecoverTitle = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetEmailTitle(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.EmailTitle = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetRobotContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.RobotContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetSmsContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.SmsContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetSmsRecoverContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.SmsRecoverContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetTtsContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.TtsContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetTtsRecoverContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.TtsRecoverContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) Validate() error {
	return dara.Validate(s)
}
