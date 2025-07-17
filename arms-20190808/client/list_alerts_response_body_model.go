// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListAlertsResponseBody
	GetMessage() *string
	SetPageBean(v *ListAlertsResponseBodyPageBean) *ListAlertsResponseBody
	GetPageBean() *ListAlertsResponseBodyPageBean
	SetRequestId(v string) *ListAlertsResponseBody
	GetRequestId() *string
}

type ListAlertsResponseBody struct {
	// The returned error message.
	//
	// example:
	//
	// alert.manager.error.code.signature.invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The struct returned.
	PageBean *ListAlertsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2FC13182-B9AF-4E6B-BE51-72669B7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertsResponseBody) GetPageBean() *ListAlertsResponseBodyPageBean {
	return s.PageBean
}

func (s *ListAlertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertsResponseBody) SetMessage(v string) *ListAlertsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertsResponseBody) SetPageBean(v *ListAlertsResponseBodyPageBean) *ListAlertsResponseBody {
	s.PageBean = v
	return s
}

func (s *ListAlertsResponseBody) SetRequestId(v string) *ListAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlertsResponseBodyPageBean struct {
	// The queried alert notification history records.
	ListAlerts []*ListAlertsResponseBodyPageBeanListAlerts `json:"ListAlerts,omitempty" xml:"ListAlerts,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of alerts returned per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of queried alerts.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAlertsResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyPageBean) GetListAlerts() []*ListAlertsResponseBodyPageBeanListAlerts {
	return s.ListAlerts
}

func (s *ListAlertsResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *ListAlertsResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *ListAlertsResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *ListAlertsResponseBodyPageBean) SetListAlerts(v []*ListAlertsResponseBodyPageBeanListAlerts) *ListAlertsResponseBodyPageBean {
	s.ListAlerts = v
	return s
}

func (s *ListAlertsResponseBodyPageBean) SetPage(v int64) *ListAlertsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListAlertsResponseBodyPageBean) SetSize(v int64) *ListAlertsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListAlertsResponseBodyPageBean) SetTotal(v int64) *ListAlertsResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *ListAlertsResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type ListAlertsResponseBodyPageBeanListAlerts struct {
	// Time to claim the alarm.
	//
	// example:
	//
	// -1
	AcknowledgeTime *int64 `json:"AcknowledgeTime,omitempty" xml:"AcknowledgeTime,omitempty"`
	// The list of activities.
	Activities []*ListAlertsResponseBodyPageBeanListAlertsActivities `json:"Activities,omitempty" xml:"Activities,omitempty" type:"Repeated"`
	// The list of events.
	AlertEvents []*ListAlertsResponseBodyPageBeanListAlertsAlertEvents `json:"AlertEvents,omitempty" xml:"AlertEvents,omitempty" type:"Repeated"`
	// The alert ID.
	//
	// example:
	//
	// 2279
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// Test alert
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The time when the alert was created.
	//
	// example:
	//
	// 2022-01-18 00:21:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of a event execution status.
	//
	// example:
	//
	// [Notification Strategy: ARMS Front-end Alarm]\\nPage Indicator Page Name: Home JS Error Number Average in the last 5 minutes &gt;= 1.0 times, current value 1.0000 times\\n
	Describe *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// The ID of the notification policy.
	//
	// example:
	//
	// 12345
	DispatchRuleId *float32 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	// The name of the notification policy.
	//
	// example:
	//
	// DingTalk Notification
	DispatchRuleName *string `json:"DispatchRuleName,omitempty" xml:"DispatchRuleName,omitempty"`
	// Alarm handler.
	//
	// example:
	//
	// Alice
	Handler *string `json:"Handler,omitempty" xml:"Handler,omitempty"`
	// The contact card of an instant messaging app.
	//
	// example:
	//
	// Test Robot
	NotifyRobots *string `json:"NotifyRobots,omitempty" xml:"NotifyRobots,omitempty"`
	// The notification object configured in the notification policy, responsible for handling alerts.
	//
	// example:
	//
	// Alice
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Alarm recovery time.
	//
	// example:
	//
	// -1
	RecoverTime *int64 `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
	// The severity level of the alert. Valid values: P6, P5, P4, P3, P2, and P1. The preceding values are listed in ascending order of severity.
	//
	// example:
	//
	// P6
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The Alert solution.
	//
	// example:
	//
	// --
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The status of the alert. Valid values:
	//
	// 	- 0: The alert is pending.
	//
	// 	- 1: The alert is being handled.
	//
	// 	- 2: The alert is handled.
	//
	// example:
	//
	// 0
	State *int64 `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAlertsResponseBodyPageBeanListAlerts) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponseBodyPageBeanListAlerts) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetAcknowledgeTime() *int64 {
	return s.AcknowledgeTime
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetActivities() []*ListAlertsResponseBodyPageBeanListAlertsActivities {
	return s.Activities
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetAlertEvents() []*ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	return s.AlertEvents
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetAlertId() *int64 {
	return s.AlertId
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetAlertName() *string {
	return s.AlertName
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetDescribe() *string {
	return s.Describe
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetDispatchRuleId() *float32 {
	return s.DispatchRuleId
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetDispatchRuleName() *string {
	return s.DispatchRuleName
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetHandler() *string {
	return s.Handler
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetNotifyRobots() *string {
	return s.NotifyRobots
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetOwner() *string {
	return s.Owner
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetRecoverTime() *int64 {
	return s.RecoverTime
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetSeverity() *string {
	return s.Severity
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetSolution() *string {
	return s.Solution
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) GetState() *int64 {
	return s.State
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetAcknowledgeTime(v int64) *ListAlertsResponseBodyPageBeanListAlerts {
	s.AcknowledgeTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetActivities(v []*ListAlertsResponseBodyPageBeanListAlertsActivities) *ListAlertsResponseBodyPageBeanListAlerts {
	s.Activities = v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetAlertEvents(v []*ListAlertsResponseBodyPageBeanListAlertsAlertEvents) *ListAlertsResponseBodyPageBeanListAlerts {
	s.AlertEvents = v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetAlertId(v int64) *ListAlertsResponseBodyPageBeanListAlerts {
	s.AlertId = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetAlertName(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.AlertName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetCreateTime(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.CreateTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetDescribe(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.Describe = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetDispatchRuleId(v float32) *ListAlertsResponseBodyPageBeanListAlerts {
	s.DispatchRuleId = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetDispatchRuleName(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.DispatchRuleName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetHandler(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.Handler = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetNotifyRobots(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.NotifyRobots = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetOwner(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.Owner = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetRecoverTime(v int64) *ListAlertsResponseBodyPageBeanListAlerts {
	s.RecoverTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetSeverity(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.Severity = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetSolution(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.Solution = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetState(v int64) *ListAlertsResponseBodyPageBeanListAlerts {
	s.State = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) Validate() error {
	return dara.Validate(s)
}

type ListAlertsResponseBodyPageBeanListAlertsActivities struct {
	// The content of the alert notification.
	//
	// example:
	//
	// [Notification policy: Send notifications to DingTalk groups] Host monitoring. Host IP address: 10.76.XX.XX. Average memory usage of the host in the last 1 minute ≥ 1.0%. Current value: 84.7454%.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the activity.
	//
	// example:
	//
	// [Alert Claimed] The alert is claimed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the handler.
	//
	// example:
	//
	// O\\&M Engineer A
	HandlerName *string `json:"HandlerName,omitempty" xml:"HandlerName,omitempty"`
	// The operation time of the activity.
	//
	// example:
	//
	// 2021-12-20 19:08:57
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The type of the activity. Valid values:
	//
	// 	- 1: The alert is claimed.
	//
	// 	- 2: The alert is disclaimed.
	//
	// 	- 3: A comment is added for the alert.
	//
	// 	- 4: The alert is disabled.
	//
	// 	- 5: An alert notification is sent.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertsResponseBodyPageBeanListAlertsActivities) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponseBodyPageBeanListAlertsActivities) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) GetContent() *string {
	return s.Content
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) GetDescription() *string {
	return s.Description
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) GetHandlerName() *string {
	return s.HandlerName
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) GetTime() *string {
	return s.Time
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) GetType() *int64 {
	return s.Type
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetContent(v string) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.Content = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetDescription(v string) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.Description = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetHandlerName(v string) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.HandlerName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetTime(v string) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.Time = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetType(v int64) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.Type = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) Validate() error {
	return dara.Validate(s)
}

type ListAlertsResponseBodyPageBeanListAlertsAlertEvents struct {
	// The name of the event.
	//
	// example:
	//
	// Test-triggered alert
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The annotations.
	//
	// example:
	//
	// { "_aliyun_arms_alert_value":"4.0" "_aliyun_arms_alert_traceId":"ac10c43116421327442277073d5461-105075299"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The description of the event.
	//
	// example:
	//
	// Host monitoring. Host IP address: 10.76.XX.XX. Average memory usage of the host in the last 1 minute ≥ 1.0%. Current value: 84.7454%.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the event ended.
	//
	// example:
	//
	// 2022-01-19 17:10:31
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The URL of the event.
	//
	// example:
	//
	// http://arms.console.aliyun.com
	GeneratorURL *string `json:"GeneratorURL,omitempty" xml:"GeneratorURL,omitempty"`
	// The name of the integration that corresponds to the alert event.
	//
	// example:
	//
	// ARMS
	IntegrationName *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	// The type of the integration that corresponds to the alert event. Valid values:
	//
	// 	- ARMS
	//
	// 	- CLOUD_MONITOR
	//
	// 	- MSE
	//
	// 	- ARMS_CLOUD_DIALTEST
	//
	// 	- PROMETHEUS
	//
	// 	- LOG_SERVICE
	//
	// 	- CUSTOM
	//
	// 	- ARMS_PROMETHEUS
	//
	// 	- ARMS_APP_MON
	//
	// 	- ARMS_FRONT_MON
	//
	// 	- ARMS_CUSTOM
	//
	// 	- XTRACE
	//
	// 	- GRAFANA
	//
	// 	- ZABBIX
	//
	// 	- SKYWALKING
	//
	// 	- EVENT_BRIDGE
	//
	// 	- NAGIOS
	//
	// 	- OPENFALCON
	//
	// 	- ARMS_INSIGHTS
	//
	// example:
	//
	// ARMS_APP_MON
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	// The tags.
	//
	// example:
	//
	// { "severity":"warning" "_aliyun_arms_alert_level":"ERROR" "_aliyun_arms_entropy":"0.30170457417889235"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The time when the event was created.
	//
	// example:
	//
	// 2022-01-19 17:05:42
	ReceiveTime *string `json:"ReceiveTime,omitempty" xml:"ReceiveTime,omitempty"`
	// The severity level of the event. Valid values:
	//
	// 	- critical
	//
	// 	- error
	//
	// 	- warning
	//
	// 	- info
	//
	// example:
	//
	// warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The time when the event started.
	//
	// example:
	//
	// 2022-01-18 00:14:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- Active: The event is not cleared.
	//
	// 	- Silenced: The event is silenced.
	//
	// 	- Resolved: The event is cleared.
	//
	// example:
	//
	// Active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAlertsResponseBodyPageBeanListAlertsAlertEvents) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetAlertName() *string {
	return s.AlertName
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetAnnotations() *string {
	return s.Annotations
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetDescription() *string {
	return s.Description
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetGeneratorURL() *string {
	return s.GeneratorURL
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetLabels() *string {
	return s.Labels
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetReceiveTime() *string {
	return s.ReceiveTime
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetSeverity() *string {
	return s.Severity
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GetState() *string {
	return s.State
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetAlertName(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.AlertName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetAnnotations(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.Annotations = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetDescription(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.Description = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetEndTime(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.EndTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetGeneratorURL(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.GeneratorURL = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetIntegrationName(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.IntegrationName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetIntegrationType(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.IntegrationType = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetLabels(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.Labels = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetReceiveTime(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.ReceiveTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetSeverity(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.Severity = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetStartTime(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.StartTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetState(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.State = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) Validate() error {
	return dara.Validate(s)
}
