// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroup(v string) *DescribeAlertLogCountRequest
	GetContactGroup() *string
	SetEndTime(v int64) *DescribeAlertLogCountRequest
	GetEndTime() *int64
	SetEventType(v string) *DescribeAlertLogCountRequest
	GetEventType() *string
	SetGroupBy(v string) *DescribeAlertLogCountRequest
	GetGroupBy() *string
	SetGroupId(v string) *DescribeAlertLogCountRequest
	GetGroupId() *string
	SetLastMin(v string) *DescribeAlertLogCountRequest
	GetLastMin() *string
	SetLevel(v string) *DescribeAlertLogCountRequest
	GetLevel() *string
	SetMetricName(v string) *DescribeAlertLogCountRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeAlertLogCountRequest
	GetNamespace() *string
	SetPageNumber(v int32) *DescribeAlertLogCountRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAlertLogCountRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAlertLogCountRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeAlertLogCountRequest
	GetRegionId() *string
	SetRuleId(v string) *DescribeAlertLogCountRequest
	GetRuleId() *string
	SetRuleName(v string) *DescribeAlertLogCountRequest
	GetRuleName() *string
	SetSearchKey(v string) *DescribeAlertLogCountRequest
	GetSearchKey() *string
	SetSendStatus(v string) *DescribeAlertLogCountRequest
	GetSendStatus() *string
	SetSourceType(v string) *DescribeAlertLogCountRequest
	GetSourceType() *string
	SetStartTime(v int64) *DescribeAlertLogCountRequest
	GetStartTime() *int64
}

type DescribeAlertLogCountRequest struct {
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
	// You can query only the alert history within the last year. If the query time range exceeds one year, the `AlertLogCount` response parameter is empty.
	//
	// > The interval between the start time (StartTime) and end time (EndTime) must be less than or equal to 15 days. Both parameters must be specified or left empty at the same time. If neither is specified, the alert history of the last 15 minutes is queried by default.
	//
	// example:
	//
	// 1610074409694
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The alert type. Valid values:
	//
	// - TRIGGERED: The alert is triggered.
	//
	// - RESOLVED: The alert recovers.
	//
	// example:
	//
	// TRIGGERED
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Aggregates data by a spatial dimension, which is equivalent to the Group By clause in SQL. Valid values:
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
	// This parameter is required.
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
	// <props="china">- P2: Phone call + text message + email + DingTalk chatbot.
	//
	// <props="china">- P3: Text message + email + DingTalk chatbot.
	//
	// <props="china">- P4: Email + DingTalk chatbot.
	//
	// <props="china">- OK: No alert.
	//
	// <props="intl">- P4: Email + DingTalk chatbot.
	//
	// <props="intl">- OK: No alert.
	//
	// <props="partner">- P4: Email + DingTalk chatbot.
	//
	// <props="partner">- OK: No alert.
	//
	// example:
	//
	// P4
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The metric name.
	//
	// > For information about the metrics of Alibaba Cloud services, see [Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the Alibaba Cloud service.
	//
	// > For information about the namespaces of Alibaba Cloud services, see [Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
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
	// The ID of the alert rule.
	//
	//
	// For information about how to obtain the alert rule ID, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// bfae2ca5b4e07d2c7278772e***********
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// test123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The search keyword for querying alert history statistics.
	//
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The alert status. Valid values:
	//
	// - 0: An alert is triggered or the alert recovers to Normal.
	//
	// - 1: The alert rule is not effective.
	//
	// - 2: The mute period.
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
	// 无
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start timestamp for querying alert history.
	//
	// Unit: milliseconds.
	//
	// You can query only the alert history within the last year. If the query time range exceeds one year, the `AlertLogCount` response parameter is empty.
	//
	// > The interval between the start time (StartTime) and end time (EndTime) must be less than or equal to 15 days. Both parameters must be specified or left empty at the same time. If neither is specified, the alert history of the last 15 minutes is queried by default.
	//
	// example:
	//
	// 1609988009694
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertLogCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountRequest) GetContactGroup() *string {
	return s.ContactGroup
}

func (s *DescribeAlertLogCountRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlertLogCountRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeAlertLogCountRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeAlertLogCountRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertLogCountRequest) GetLastMin() *string {
	return s.LastMin
}

func (s *DescribeAlertLogCountRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeAlertLogCountRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertLogCountRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertLogCountRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAlertLogCountRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertLogCountRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAlertLogCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertLogCountRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertLogCountRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAlertLogCountRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeAlertLogCountRequest) GetSendStatus() *string {
	return s.SendStatus
}

func (s *DescribeAlertLogCountRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAlertLogCountRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlertLogCountRequest) SetContactGroup(v string) *DescribeAlertLogCountRequest {
	s.ContactGroup = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetEndTime(v int64) *DescribeAlertLogCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetEventType(v string) *DescribeAlertLogCountRequest {
	s.EventType = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetGroupBy(v string) *DescribeAlertLogCountRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetGroupId(v string) *DescribeAlertLogCountRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetLastMin(v string) *DescribeAlertLogCountRequest {
	s.LastMin = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetLevel(v string) *DescribeAlertLogCountRequest {
	s.Level = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetMetricName(v string) *DescribeAlertLogCountRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetNamespace(v string) *DescribeAlertLogCountRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetPageNumber(v int32) *DescribeAlertLogCountRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetPageSize(v int32) *DescribeAlertLogCountRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetProduct(v string) *DescribeAlertLogCountRequest {
	s.Product = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetRegionId(v string) *DescribeAlertLogCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetRuleId(v string) *DescribeAlertLogCountRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetRuleName(v string) *DescribeAlertLogCountRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetSearchKey(v string) *DescribeAlertLogCountRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetSendStatus(v string) *DescribeAlertLogCountRequest {
	s.SendStatus = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetSourceType(v string) *DescribeAlertLogCountRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetStartTime(v int64) *DescribeAlertLogCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertLogCountRequest) Validate() error {
	return dara.Validate(s)
}
