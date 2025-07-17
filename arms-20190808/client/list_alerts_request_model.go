// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *ListAlertsRequest
	GetAlertName() *string
	SetDispatchRuleId(v int64) *ListAlertsRequest
	GetDispatchRuleId() *int64
	SetEndTime(v string) *ListAlertsRequest
	GetEndTime() *string
	SetIntegrationType(v string) *ListAlertsRequest
	GetIntegrationType() *string
	SetOwner(v string) *ListAlertsRequest
	GetOwner() *string
	SetPage(v int64) *ListAlertsRequest
	GetPage() *int64
	SetRegionId(v string) *ListAlertsRequest
	GetRegionId() *string
	SetSeverity(v string) *ListAlertsRequest
	GetSeverity() *string
	SetShowActivities(v bool) *ListAlertsRequest
	GetShowActivities() *bool
	SetShowEvents(v bool) *ListAlertsRequest
	GetShowEvents() *bool
	SetSize(v int64) *ListAlertsRequest
	GetSize() *int64
	SetStartTime(v string) *ListAlertsRequest
	GetStartTime() *string
	SetState(v int64) *ListAlertsRequest
	GetState() *int64
}

type ListAlertsRequest struct {
	// The name of the alert rule.
	//
	// example:
	//
	// Test alert
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the notification policy.
	//
	// example:
	//
	// 12345
	DispatchRuleId *int64 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	// The end time of the alert sending history that you want to query. Specify the time in the `YYYY-MM-DD HH:mm:ss` format.
	//
	// example:
	//
	// 2021-12-22 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The integration type.
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
	// The notification object configured in the notification policy, responsible for handling alerts.
	//
	// example:
	//
	// Alice
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The severity level of the alert. Valid values: P6, P5, P4, P3, P2, and P1. The preceding values are listed in ascending order of severity.
	//
	// example:
	//
	// P6
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// Specifies whether to query the activities that correspond to alerts. Valid values:
	//
	// 	- `false` (default value): The activities are not queried.
	//
	// 	- `true`: The activities in the last three days are queried.
	//
	// example:
	//
	// true
	ShowActivities *bool `json:"ShowActivities,omitempty" xml:"ShowActivities,omitempty"`
	// Specifies whether to query the events that correspond to alerts. Valid values:
	//
	// 	- `false` (default value): The events are not queried.
	//
	// 	- `true`: The events are queried.
	//
	// example:
	//
	// true
	ShowEvents *bool `json:"ShowEvents,omitempty" xml:"ShowEvents,omitempty"`
	// The number of alerts to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time of the alert sending history that you want to query. Specify the time in the `YYYY-MM-DD HH:mm:ss` format.
	//
	// example:
	//
	// 2021-12-10 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// 2
	State *int64 `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertsRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *ListAlertsRequest) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *ListAlertsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAlertsRequest) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *ListAlertsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListAlertsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAlertsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertsRequest) GetSeverity() *string {
	return s.Severity
}

func (s *ListAlertsRequest) GetShowActivities() *bool {
	return s.ShowActivities
}

func (s *ListAlertsRequest) GetShowEvents() *bool {
	return s.ShowEvents
}

func (s *ListAlertsRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListAlertsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAlertsRequest) GetState() *int64 {
	return s.State
}

func (s *ListAlertsRequest) SetAlertName(v string) *ListAlertsRequest {
	s.AlertName = &v
	return s
}

func (s *ListAlertsRequest) SetDispatchRuleId(v int64) *ListAlertsRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *ListAlertsRequest) SetEndTime(v string) *ListAlertsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlertsRequest) SetIntegrationType(v string) *ListAlertsRequest {
	s.IntegrationType = &v
	return s
}

func (s *ListAlertsRequest) SetOwner(v string) *ListAlertsRequest {
	s.Owner = &v
	return s
}

func (s *ListAlertsRequest) SetPage(v int64) *ListAlertsRequest {
	s.Page = &v
	return s
}

func (s *ListAlertsRequest) SetRegionId(v string) *ListAlertsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAlertsRequest) SetSeverity(v string) *ListAlertsRequest {
	s.Severity = &v
	return s
}

func (s *ListAlertsRequest) SetShowActivities(v bool) *ListAlertsRequest {
	s.ShowActivities = &v
	return s
}

func (s *ListAlertsRequest) SetShowEvents(v bool) *ListAlertsRequest {
	s.ShowEvents = &v
	return s
}

func (s *ListAlertsRequest) SetSize(v int64) *ListAlertsRequest {
	s.Size = &v
	return s
}

func (s *ListAlertsRequest) SetStartTime(v string) *ListAlertsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAlertsRequest) SetState(v int64) *ListAlertsRequest {
	s.State = &v
	return s
}

func (s *ListAlertsRequest) Validate() error {
	return dara.Validate(s)
}
