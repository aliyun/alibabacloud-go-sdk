// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertHistoryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmHistoryList(v *DescribeAlertHistoryListResponseBodyAlarmHistoryList) *DescribeAlertHistoryListResponseBody
	GetAlarmHistoryList() *DescribeAlertHistoryListResponseBodyAlarmHistoryList
	SetCode(v string) *DescribeAlertHistoryListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeAlertHistoryListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertHistoryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertHistoryListResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeAlertHistoryListResponseBody
	GetTotal() *string
}

type DescribeAlertHistoryListResponseBody struct {
	// The details of historical alerts.
	AlarmHistoryList *DescribeAlertHistoryListResponseBodyAlarmHistoryList `json:"AlarmHistoryList,omitempty" xml:"AlarmHistoryList,omitempty" type:"Struct"`
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C3C69FBE-2262-541F-A409-C52F380BAE63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAlertHistoryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBody) GetAlarmHistoryList() *DescribeAlertHistoryListResponseBodyAlarmHistoryList {
	return s.AlarmHistoryList
}

func (s *DescribeAlertHistoryListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAlertHistoryListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertHistoryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertHistoryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertHistoryListResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeAlertHistoryListResponseBody) SetAlarmHistoryList(v *DescribeAlertHistoryListResponseBodyAlarmHistoryList) *DescribeAlertHistoryListResponseBody {
	s.AlarmHistoryList = v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetCode(v string) *DescribeAlertHistoryListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetMessage(v string) *DescribeAlertHistoryListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetRequestId(v string) *DescribeAlertHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetSuccess(v bool) *DescribeAlertHistoryListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetTotal(v string) *DescribeAlertHistoryListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryList struct {
	AlarmHistory []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory `json:"AlarmHistory,omitempty" xml:"AlarmHistory,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryList) GetAlarmHistory() []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	return s.AlarmHistory
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryList) SetAlarmHistory(v []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) *DescribeAlertHistoryListResponseBodyAlarmHistoryList {
	s.AlarmHistory = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryList) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory struct {
	// The timestamp when the alert was triggered. Unit: milliseconds.
	//
	// example:
	//
	// 1640586600000
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// The TradeManager IDs of the alert contacts.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	ContactALIIMs *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs `json:"ContactALIIMs,omitempty" xml:"ContactALIIMs,omitempty" type:"Struct"`
	// The alert contact groups.
	ContactGroups *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	// The email addresses of the alert contacts.
	ContactMails *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails `json:"ContactMails,omitempty" xml:"ContactMails,omitempty" type:"Struct"`
	// The mobile numbers of the alert contacts.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	ContactSmses *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses `json:"ContactSmses,omitempty" xml:"ContactSmses,omitempty" type:"Struct"`
	// The alert contacts that receive alert notifications.
	Contacts *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	// The resources that are monitored.
	//
	// example:
	//
	// {\\"instanceId\\":\\"i-bp1cqhiw1za2****\\"}
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before an alert is triggered.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The expression that is used to trigger alerts.
	//
	// example:
	//
	// $Average>=10
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 7671****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// i-bp1cqhiw1za2****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The duration of the alert. Unit: milliseconds.
	//
	// example:
	//
	// 360133
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The severity level and notification methods of the alert. Valid values:
	//
	// 	- P4: Alert notifications are sent by using emails and DingTalk chatbots.
	//
	// 	- OK: No alert is generated.
	//
	// example:
	//
	// P4
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The metric name.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// applyTemplate61dc81b5-d357-4cf6-a9b7-9f83c1d5****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// ECS_Rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The alert status. Valid values:
	//
	// 	- ALARM: Alerts are triggered.
	//
	// 	- OK: No alerts are triggered.
	//
	// example:
	//
	// ALARM
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates whether alerts are muted. Valid values:
	//
	// 	- 2 (default): Alerts are muted and are not triggered within the mute period, even if the condition specified in the alert rule is met.
	//
	// 	- 0: Alerts are triggered or cleared.
	//
	// 	- 1: The alert rule is ineffective.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The threshold of the metric value to trigger or clear an alert.
	//
	// example:
	//
	// 10.58
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhooks *string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetAlertTime() *int64 {
	return s.AlertTime
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContactALIIMs() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs {
	return s.ContactALIIMs
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContactGroups() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups {
	return s.ContactGroups
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContactMails() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails {
	return s.ContactMails
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContactSmses() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses {
	return s.ContactSmses
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetContacts() *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts {
	return s.Contacts
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetExpression() *string {
	return s.Expression
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetLevel() *string {
	return s.Level
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetState() *string {
	return s.State
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetValue() *string {
	return s.Value
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GetWebhooks() *string {
	return s.Webhooks
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetAlertTime(v int64) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.AlertTime = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactALIIMs(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactALIIMs = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactGroups(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactGroups = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactMails(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactMails = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactSmses(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactSmses = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContacts(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Contacts = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetDimensions(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetEvaluationCount(v int32) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetExpression(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Expression = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetGroupId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetInstanceName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetLastTime(v int64) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.LastTime = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetLevel(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Level = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetMetricName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetNamespace(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetRuleId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetRuleName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetState(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.State = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetStatus(v int32) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Status = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetValue(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Value = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetWebhooks(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Webhooks = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs struct {
	ContactALIIM []*string `json:"ContactALIIM,omitempty" xml:"ContactALIIM,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) GetContactALIIM() []*string {
	return s.ContactALIIM
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) SetContactALIIM(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs {
	s.ContactALIIM = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) GetContactGroup() []*string {
	return s.ContactGroup
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) SetContactGroup(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups {
	s.ContactGroup = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails struct {
	ContactMail []*string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) GetContactMail() []*string {
	return s.ContactMail
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) SetContactMail(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails {
	s.ContactMail = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses struct {
	ContactSms []*string `json:"ContactSms,omitempty" xml:"ContactSms,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) GetContactSms() []*string {
	return s.ContactSms
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) SetContactSms(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses {
	s.ContactSms = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts struct {
	Contact []*string `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) GetContact() []*string {
	return s.Contact
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) SetContact(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts {
	s.Contact = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) Validate() error {
	return dara.Validate(s)
}
