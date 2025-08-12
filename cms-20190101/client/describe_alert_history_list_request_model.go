// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertHistoryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscending(v bool) *DescribeAlertHistoryListRequest
	GetAscending() *bool
	SetEndTime(v string) *DescribeAlertHistoryListRequest
	GetEndTime() *string
	SetGroupId(v string) *DescribeAlertHistoryListRequest
	GetGroupId() *string
	SetMetricName(v string) *DescribeAlertHistoryListRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeAlertHistoryListRequest
	GetNamespace() *string
	SetPage(v int32) *DescribeAlertHistoryListRequest
	GetPage() *int32
	SetPageSize(v int32) *DescribeAlertHistoryListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAlertHistoryListRequest
	GetRegionId() *string
	SetRuleId(v string) *DescribeAlertHistoryListRequest
	GetRuleId() *string
	SetRuleName(v string) *DescribeAlertHistoryListRequest
	GetRuleName() *string
	SetStartTime(v string) *DescribeAlertHistoryListRequest
	GetStartTime() *string
	SetState(v string) *DescribeAlertHistoryListRequest
	GetState() *string
	SetStatus(v string) *DescribeAlertHistoryListRequest
	GetStatus() *string
}

type DescribeAlertHistoryListRequest struct {
	// The order of alerts. Valid values:
	//
	// 	- true (default value): reverse chronological order
	//
	// 	- false: chronological order
	//
	// example:
	//
	// true
	Ascending *bool `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	// The end timestamp of the historical alerts that you want to query.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1640608200000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the application group.
	//
	// For information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// example:
	//
	// 7671****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The metric that is used to monitor the cloud service.
	//
	// For information about how to query the name of a metric, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// For information about how to query the namespace of a cloud service, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule.
	//
	// For information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// applyTemplate61dc81b5-d357-4cf6-a9b7-9f83c1d5****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// For information about how to query the name of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// ECS_Rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The start timestamp of the historical alerts that you want to query.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1640237400000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the alert. Valid values:
	//
	// 	- ALARM (default value): Alerts are triggered.
	//
	// 	- OK: No alerts are triggered.
	//
	// example:
	//
	// ALARM
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Specifies whether alerts are muted. Valid values:
	//
	// 	- 2 (default value): Alerts are muted and are not triggered within the mute period, even if the condition specified in the alert rule is met.
	//
	// 	- 0: Alerts are triggered or cleared.
	//
	// 	- 1: The alert rule is ineffective.
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAlertHistoryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListRequest) GetAscending() *bool {
	return s.Ascending
}

func (s *DescribeAlertHistoryListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAlertHistoryListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertHistoryListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertHistoryListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertHistoryListRequest) GetPage() *int32 {
	return s.Page
}

func (s *DescribeAlertHistoryListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertHistoryListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertHistoryListRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertHistoryListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAlertHistoryListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAlertHistoryListRequest) GetState() *string {
	return s.State
}

func (s *DescribeAlertHistoryListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAlertHistoryListRequest) SetAscending(v bool) *DescribeAlertHistoryListRequest {
	s.Ascending = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetEndTime(v string) *DescribeAlertHistoryListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetGroupId(v string) *DescribeAlertHistoryListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetMetricName(v string) *DescribeAlertHistoryListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetNamespace(v string) *DescribeAlertHistoryListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetPage(v int32) *DescribeAlertHistoryListRequest {
	s.Page = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetPageSize(v int32) *DescribeAlertHistoryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetRegionId(v string) *DescribeAlertHistoryListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetRuleId(v string) *DescribeAlertHistoryListRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetRuleName(v string) *DescribeAlertHistoryListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetStartTime(v string) *DescribeAlertHistoryListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetState(v string) *DescribeAlertHistoryListRequest {
	s.State = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetStatus(v string) *DescribeAlertHistoryListRequest {
	s.Status = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) Validate() error {
	return dara.Validate(s)
}
