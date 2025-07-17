// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *ListAlertEventsResponseBodyPageBean) *ListAlertEventsResponseBody
	GetPageBean() *ListAlertEventsResponseBodyPageBean
	SetRequestId(v string) *ListAlertEventsResponseBody
	GetRequestId() *string
}

type ListAlertEventsResponseBody struct {
	// The returned struct.
	PageBean *ListAlertEventsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2FC13182-B9AF-4E6B-BE51-72669B7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlertEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBody) GetPageBean() *ListAlertEventsResponseBodyPageBean {
	return s.PageBean
}

func (s *ListAlertEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertEventsResponseBody) SetPageBean(v *ListAlertEventsResponseBodyPageBean) *ListAlertEventsResponseBody {
	s.PageBean = v
	return s
}

func (s *ListAlertEventsResponseBody) SetRequestId(v string) *ListAlertEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyPageBean struct {
	// The queried historical alert events.
	Events []*ListAlertEventsResponseBodyPageBeanEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 24
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAlertEventsResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyPageBean) GetEvents() []*ListAlertEventsResponseBodyPageBeanEvents {
	return s.Events
}

func (s *ListAlertEventsResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *ListAlertEventsResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *ListAlertEventsResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *ListAlertEventsResponseBodyPageBean) SetEvents(v []*ListAlertEventsResponseBodyPageBeanEvents) *ListAlertEventsResponseBodyPageBean {
	s.Events = v
	return s
}

func (s *ListAlertEventsResponseBodyPageBean) SetPage(v int64) *ListAlertEventsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBean) SetSize(v int64) *ListAlertEventsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBean) SetTotal(v int64) *ListAlertEventsResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyPageBeanEvents struct {
	// The associated alerts.
	Alarms []*ListAlertEventsResponseBodyPageBeanEventsAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// The name of the alert.
	//
	// example:
	//
	// Test-triggered alert
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The annotations.
	//
	// example:
	//
	// [{\\"Name\\":\\"annotation-a\\",\\"Value\\":\\"annotation a value\\"}]
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The description of the alert event.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2021-12-20 17:42:16
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The URL of the alert event.
	//
	// example:
	//
	// https://xxx.xx/
	GeneratorURL *string `json:"GeneratorURL,omitempty" xml:"GeneratorURL,omitempty"`
	// The user who handled the alert.
	//
	// example:
	//
	// Tom
	HandlerName *string `json:"HandlerName,omitempty" xml:"HandlerName,omitempty"`
	// The name of the alert integration.
	//
	// example:
	//
	// Custom integration
	IntegrationName *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	// The type of the alert integration.
	//
	// example:
	//
	// CUSTOM
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	// The tags.
	//
	// example:
	//
	// [{\\"name\\":\\"severity\\",\\"value\\":\\"error\\"}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The associated notification policies.
	NotificationPolicies []*ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies `json:"NotificationPolicies,omitempty" xml:"NotificationPolicies,omitempty" type:"Repeated"`
	// The time when the alert event was received.
	//
	// example:
	//
	// 2021-12-20 17:42:16
	ReceiveTime *string `json:"ReceiveTime,omitempty" xml:"ReceiveTime,omitempty"`
	// The severity level of the alert. Valid values:
	//
	// 	- critical: P1
	//
	// 	- error: P2
	//
	// 	- warning: P3
	//
	// 	- page: P4
	//
	// 	- default: P6
	//
	// example:
	//
	// critical
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2021-12-20 17:42:16
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the alert event. Valid values:
	//
	// 	- Active
	//
	// 	- Silenced
	//
	// 	- Resolved
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of times the event is triggered.
	//
	// example:
	//
	// 10
	TriggerCount *int64 `json:"TriggerCount,omitempty" xml:"TriggerCount,omitempty"`
}

func (s ListAlertEventsResponseBodyPageBeanEvents) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyPageBeanEvents) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetAlarms() []*ListAlertEventsResponseBodyPageBeanEventsAlarms {
	return s.Alarms
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetAlertName() *string {
	return s.AlertName
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetAnnotations() *string {
	return s.Annotations
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetDescription() *string {
	return s.Description
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetGeneratorURL() *string {
	return s.GeneratorURL
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetHandlerName() *string {
	return s.HandlerName
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetLabels() *string {
	return s.Labels
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetNotificationPolicies() []*ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies {
	return s.NotificationPolicies
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetReceiveTime() *string {
	return s.ReceiveTime
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetSeverity() *string {
	return s.Severity
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetStatus() *string {
	return s.Status
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) GetTriggerCount() *int64 {
	return s.TriggerCount
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetAlarms(v []*ListAlertEventsResponseBodyPageBeanEventsAlarms) *ListAlertEventsResponseBodyPageBeanEvents {
	s.Alarms = v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetAlertName(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.AlertName = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetAnnotations(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.Annotations = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetDescription(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.Description = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetEndTime(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.EndTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetGeneratorURL(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.GeneratorURL = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetHandlerName(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.HandlerName = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetIntegrationName(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.IntegrationName = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetIntegrationType(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.IntegrationType = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetLabels(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.Labels = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetNotificationPolicies(v []*ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies) *ListAlertEventsResponseBodyPageBeanEvents {
	s.NotificationPolicies = v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetReceiveTime(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.ReceiveTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetSeverity(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.Severity = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetStartTime(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.StartTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetStatus(v string) *ListAlertEventsResponseBodyPageBeanEvents {
	s.Status = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) SetTriggerCount(v int64) *ListAlertEventsResponseBodyPageBeanEvents {
	s.TriggerCount = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEvents) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyPageBeanEventsAlarms struct {
	// The ID of the alert.
	//
	// example:
	//
	// 77444
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// Test-triggered alert
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The time when the alert was created.
	//
	// example:
	//
	// 2021-12-20 07:10:18
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The status of the alert. Valid values:
	//
	// 	- 0: The alert is pending.
	//
	// 	- 1: The alert is being handled.
	//
	// 	- 2: The alert is cleared.
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAlertEventsResponseBodyPageBeanEventsAlarms) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyPageBeanEventsAlarms) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) GetAlarmName() *string {
	return s.AlarmName
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) GetState() *int32 {
	return s.State
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) SetAlarmId(v int64) *ListAlertEventsResponseBodyPageBeanEventsAlarms {
	s.AlarmId = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) SetAlarmName(v string) *ListAlertEventsResponseBodyPageBeanEventsAlarms {
	s.AlarmName = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) SetCreateTime(v string) *ListAlertEventsResponseBodyPageBeanEventsAlarms {
	s.CreateTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) SetState(v int32) *ListAlertEventsResponseBodyPageBeanEventsAlarms {
	s.State = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEventsAlarms) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies struct {
	// The ID of the notification policy.
	//
	// example:
	//
	// 646093
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the notification policy.
	//
	// example:
	//
	// P1 alert notification policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies) GetId() *int64 {
	return s.Id
}

func (s *ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies) GetName() *string {
	return s.Name
}

func (s *ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies) SetId(v int64) *ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies {
	s.Id = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies) SetName(v string) *ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies {
	s.Name = &v
	return s
}

func (s *ListAlertEventsResponseBodyPageBeanEventsNotificationPolicies) Validate() error {
	return dara.Validate(s)
}
