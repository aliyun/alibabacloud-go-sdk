// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogHistogramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroup(v string) *DescribeAlertLogHistogramRequest
	GetContactGroup() *string
	SetEndTime(v int64) *DescribeAlertLogHistogramRequest
	GetEndTime() *int64
	SetEventType(v string) *DescribeAlertLogHistogramRequest
	GetEventType() *string
	SetGroupBy(v string) *DescribeAlertLogHistogramRequest
	GetGroupBy() *string
	SetGroupId(v string) *DescribeAlertLogHistogramRequest
	GetGroupId() *string
	SetLastMin(v string) *DescribeAlertLogHistogramRequest
	GetLastMin() *string
	SetLevel(v string) *DescribeAlertLogHistogramRequest
	GetLevel() *string
	SetMetricName(v string) *DescribeAlertLogHistogramRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeAlertLogHistogramRequest
	GetNamespace() *string
	SetPageNumber(v int32) *DescribeAlertLogHistogramRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAlertLogHistogramRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAlertLogHistogramRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeAlertLogHistogramRequest
	GetRegionId() *string
	SetRuleId(v string) *DescribeAlertLogHistogramRequest
	GetRuleId() *string
	SetRuleName(v string) *DescribeAlertLogHistogramRequest
	GetRuleName() *string
	SetSearchKey(v string) *DescribeAlertLogHistogramRequest
	GetSearchKey() *string
	SetSendStatus(v string) *DescribeAlertLogHistogramRequest
	GetSendStatus() *string
	SetSourceType(v string) *DescribeAlertLogHistogramRequest
	GetSourceType() *string
	SetStartTime(v int64) *DescribeAlertLogHistogramRequest
	GetStartTime() *int64
}

type DescribeAlertLogHistogramRequest struct {
	// The alert contact group.
	//
	// example:
	//
	// ECS_Group
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// The end timestamp for querying alert history.
	//
	// Unit: milliseconds.
	//
	// > - You can query only the alert history within the last year.
	//
	// > - The interval between the start time (StartTime) and end time (EndTime) must be less than or equal to 15 days.
	//
	// example:
	//
	// 1609989009694
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The alert type. Valid values:
	//
	// - TRIGGERED: The alert is triggered.
	//
	// - RESOLVED: The alert recovers.
	//
	// example:
	//
	// RESOLVED
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Aggregates data by spatial dimension, equivalent to Group By in SQL. Valid values:
	//
	// - `product`: aggregates by Alibaba Cloud service.
	//
	// - `level`: aggregates by alert level.
	//
	// - `groupId`: aggregates by application group.
	//
	// - `contactGroup`: aggregates by alert contact group.
	//
	// - `product,metricName`: aggregates by Alibaba Cloud service and metric.
	//
	// example:
	//
	// product
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The application group ID.
	//
	// example:
	//
	// 7301****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The period for obtaining logs. Unit: minutes.
	//
	// example:
	//
	// 360
	LastMin *string `json:"LastMin,omitempty" xml:"LastMin,omitempty"`
	// The alert level and notification method. Valid values:
	//
	// <props="china">- P2: phone call + text message + email + DingTalk chatbot.
	//
	// <props="china">- P3: text message + email + DingTalk chatbot.
	//
	// <props="china">- P4: email + DingTalk chatbot.
	//
	// <props="china">- OK: no alert.
	//
	// <props="intl">- P4: email + DingTalk chatbot.
	//
	// <props="intl">- OK: no alert.
	//
	// <props="partner">- P4: email + DingTalk chatbot.
	//
	// <props="partner">- OK: no alert.
	//
	// example:
	//
	// P4
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The metric name.
	//
	// > For information about the metrics of Alibaba Cloud services, see [Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the Alibaba Cloud service.
	//
	// > For information about the namespaces of Alibaba Cloud services, see [Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ECS
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The alert rule ID.
	//
	//
	// For information about how to obtain the alert rule ID, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// ae06917_75a8c43178ab66****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The alert rule name.
	//
	// example:
	//
	// test123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The keyword for searching alert history.
	//
	// example:
	//
	// alert
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The alert status. Valid values:
	//
	// - 0: An alert is triggered or the alert recovers to Normal.
	//
	// - 1: The alert is not in the effective period.
	//
	// - 2: The mute for epoch.
	//
	// - 3: The host is restarting.
	//
	// - 4: Alerts are not send.
	//
	// <props="china">When the alert status is 0, if the value of Level is P2, P3, or P4, an alert is triggered. If the value of Level is OK, the alert recovers to Normal.
	//
	// <props="intl">When the alert status is 0, if the value of Level is P4, an alert is triggered. If the value of Level is OK, the alert recovers to Normal.
	//
	// <props="partner">When the alert status is 0, if the value of Level is P4, an alert is triggered. If the value of Level is OK, the alert recovers to Normal.
	//
	// example:
	//
	// 0
	SendStatus *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This parameter is deprecated and does not need to be specified.
	//
	// example:
	//
	// None
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start timestamp for querying alert history.
	//
	// Unit: milliseconds.
	//
	// > - You can query only the alert history within the last year.
	//
	// > - The interval between the start time (StartTime) and end time (EndTime) must be less than or equal to 15 days.
	//
	// example:
	//
	// 1609988009694
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertLogHistogramRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogHistogramRequest) GetContactGroup() *string {
	return s.ContactGroup
}

func (s *DescribeAlertLogHistogramRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlertLogHistogramRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeAlertLogHistogramRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeAlertLogHistogramRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertLogHistogramRequest) GetLastMin() *string {
	return s.LastMin
}

func (s *DescribeAlertLogHistogramRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeAlertLogHistogramRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertLogHistogramRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertLogHistogramRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAlertLogHistogramRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertLogHistogramRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAlertLogHistogramRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertLogHistogramRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertLogHistogramRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAlertLogHistogramRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeAlertLogHistogramRequest) GetSendStatus() *string {
	return s.SendStatus
}

func (s *DescribeAlertLogHistogramRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAlertLogHistogramRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlertLogHistogramRequest) SetContactGroup(v string) *DescribeAlertLogHistogramRequest {
	s.ContactGroup = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetEndTime(v int64) *DescribeAlertLogHistogramRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetEventType(v string) *DescribeAlertLogHistogramRequest {
	s.EventType = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetGroupBy(v string) *DescribeAlertLogHistogramRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetGroupId(v string) *DescribeAlertLogHistogramRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetLastMin(v string) *DescribeAlertLogHistogramRequest {
	s.LastMin = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetLevel(v string) *DescribeAlertLogHistogramRequest {
	s.Level = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetMetricName(v string) *DescribeAlertLogHistogramRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetNamespace(v string) *DescribeAlertLogHistogramRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetPageNumber(v int32) *DescribeAlertLogHistogramRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetPageSize(v int32) *DescribeAlertLogHistogramRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetProduct(v string) *DescribeAlertLogHistogramRequest {
	s.Product = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetRegionId(v string) *DescribeAlertLogHistogramRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetRuleId(v string) *DescribeAlertLogHistogramRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetRuleName(v string) *DescribeAlertLogHistogramRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetSearchKey(v string) *DescribeAlertLogHistogramRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetSendStatus(v string) *DescribeAlertLogHistogramRequest {
	s.SendStatus = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetSourceType(v string) *DescribeAlertLogHistogramRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetStartTime(v int64) *DescribeAlertLogHistogramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) Validate() error {
	return dara.Validate(s)
}
