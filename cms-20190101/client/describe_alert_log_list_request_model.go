// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroup(v string) *DescribeAlertLogListRequest
	GetContactGroup() *string
	SetEndTime(v int64) *DescribeAlertLogListRequest
	GetEndTime() *int64
	SetEventType(v string) *DescribeAlertLogListRequest
	GetEventType() *string
	SetGroupBy(v string) *DescribeAlertLogListRequest
	GetGroupBy() *string
	SetGroupId(v string) *DescribeAlertLogListRequest
	GetGroupId() *string
	SetLastMin(v string) *DescribeAlertLogListRequest
	GetLastMin() *string
	SetLevel(v string) *DescribeAlertLogListRequest
	GetLevel() *string
	SetMetricName(v string) *DescribeAlertLogListRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeAlertLogListRequest
	GetNamespace() *string
	SetPageNumber(v int32) *DescribeAlertLogListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAlertLogListRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAlertLogListRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeAlertLogListRequest
	GetRegionId() *string
	SetRuleId(v string) *DescribeAlertLogListRequest
	GetRuleId() *string
	SetRuleName(v string) *DescribeAlertLogListRequest
	GetRuleName() *string
	SetSearchKey(v string) *DescribeAlertLogListRequest
	GetSearchKey() *string
	SetSendStatus(v string) *DescribeAlertLogListRequest
	GetSendStatus() *string
	SetSourceType(v string) *DescribeAlertLogListRequest
	GetSourceType() *string
	SetStartTime(v int64) *DescribeAlertLogListRequest
	GetStartTime() *int64
}

type DescribeAlertLogListRequest struct {
	// The alert contact group.
	//
	// example:
	//
	// ECS_Group
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// The end of the time range to query the alert history.
	//
	// Unit: milliseconds.
	//
	// You can query only the alert history within the last year. If the query time range exceeds one year, the return value of the `AlertLogList` parameter is empty.
	//
	// > The interval between the start time (`StartTime`) and end time (`EndTime`) must be less than or equal to 15 days. Both parameters must be specified or unspecified at the same time. If they are not specified, the alert history within the last 15 minutes is queried by default.
	//
	// example:
	//
	// 1610074409694
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The alert type. Valid values:
	//
	// - TRIGGERED: The alert is triggered.
	//
	// - RESOLVED: The alert is cleared.
	//
	// example:
	//
	// TRIGGERED
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The spatial dimension by which the data is aggregated, which is equivalent to Group By in SQL. Valid values:
	//
	// - `product`: aggregates data by cloud service.
	//
	// - `level`: aggregates data by alert level.
	//
	// - `groupId`: aggregates data by application group.
	//
	// - `contactGroup`: aggregates data by alert contact group.
	//
	// - `product,metricName`: aggregates data by cloud service and metric.
	//
	// example:
	//
	// product
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 7301****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The interval at which logs are obtained. Unit: minutes.
	//
	// example:
	//
	// 360
	LastMin *string `json:"LastMin,omitempty" xml:"LastMin,omitempty"`
	// The alert level and notification methods. Valid values:
	//
	// <props="china">- P2: phone calls, text messages, emails, and DingTalk chatbots.
	//
	// <props="china">- P3: text messages, emails, and DingTalk chatbots.
	//
	// <props="china">- P4: emails and DingTalk chatbots.
	//
	// <props="china">- OK: no alerts.
	//
	// <props="intl">- P4: emails and DingTalk chatbots.
	//
	// <props="intl">- OK: no alerts.
	//
	// <props="partner">- P4: emails and DingTalk chatbots.
	//
	// <props="partner">- OK: no alerts.
	//
	// example:
	//
	// P4
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the metric.
	//
	// >For more information about the metrics of cloud services, see [Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// IntranetInRate
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// > For more information about the namespaces of cloud services, see [Metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// The abbreviation of the cloud service name.
	//
	// For more information about how to obtain the abbreviation of a cloud service name, see [DescribeProductsOfActiveMetricRule](https://help.aliyun.com/document_detail/114930.html).
	//
	// example:
	//
	// ECS
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule.
	//
	// For more information about how to query the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// bc369e8_30f87e517ed2fc****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// test123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The keyword used to query the alert history.
	//
	// example:
	//
	// alert
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The alert status. Valid values:
	//
	// - 0: An alert is triggered or cleared.
	//
	// - 1: The current time is not within the effective period of the alert.
	//
	// - 2: The current time is within the channel silence period.
	//
	// - 3: The host is being restarted.
	//
	// - 4: No alerts are sent.
	//
	// <props="china">When the alert status is 0, an alert is triggered if Level is set to P2, P3, or P4; the alert is cleared if Level is set to OK.
	//
	// <props="intl">When the alert status is 0, an alert is triggered if Level is set to P4; the alert is cleared if Level is set to OK.
	//
	// <props="partner">When the alert status is 0, an alert is triggered if Level is set to P4; the alert is cleared if Level is set to OK.
	//
	// example:
	//
	// 0
	SendStatus *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// The type of the alert rule. Valid value: METRIC, which indicates a time series metric alert rule.
	//
	// example:
	//
	// METRIC
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The beginning of the time range to query the alert history.
	//
	// Unit: milliseconds.
	//
	// You can query only the alert history within the last year. If the query time range exceeds one year, the return value of the `AlertLogList` parameter is empty.
	//
	// > The interval between the start time (`StartTime`) and end time (`EndTime`) must be less than or equal to 15 days. Both parameters must be specified or unspecified at the same time. If they are not specified, the alert history within the last 15 minutes is queried by default.
	//
	// example:
	//
	// 1609988009694
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertLogListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListRequest) GetContactGroup() *string {
	return s.ContactGroup
}

func (s *DescribeAlertLogListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlertLogListRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeAlertLogListRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeAlertLogListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertLogListRequest) GetLastMin() *string {
	return s.LastMin
}

func (s *DescribeAlertLogListRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeAlertLogListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlertLogListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertLogListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAlertLogListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertLogListRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAlertLogListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertLogListRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertLogListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAlertLogListRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeAlertLogListRequest) GetSendStatus() *string {
	return s.SendStatus
}

func (s *DescribeAlertLogListRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeAlertLogListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlertLogListRequest) SetContactGroup(v string) *DescribeAlertLogListRequest {
	s.ContactGroup = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetEndTime(v int64) *DescribeAlertLogListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetEventType(v string) *DescribeAlertLogListRequest {
	s.EventType = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetGroupBy(v string) *DescribeAlertLogListRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetGroupId(v string) *DescribeAlertLogListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetLastMin(v string) *DescribeAlertLogListRequest {
	s.LastMin = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetLevel(v string) *DescribeAlertLogListRequest {
	s.Level = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetMetricName(v string) *DescribeAlertLogListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetNamespace(v string) *DescribeAlertLogListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetPageNumber(v int32) *DescribeAlertLogListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetPageSize(v int32) *DescribeAlertLogListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetProduct(v string) *DescribeAlertLogListRequest {
	s.Product = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetRegionId(v string) *DescribeAlertLogListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetRuleId(v string) *DescribeAlertLogListRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetRuleName(v string) *DescribeAlertLogListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetSearchKey(v string) *DescribeAlertLogListRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetSendStatus(v string) *DescribeAlertLogListRequest {
	s.SendStatus = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetSourceType(v string) *DescribeAlertLogListRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetStartTime(v int64) *DescribeAlertLogListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertLogListRequest) Validate() error {
	return dara.Validate(s)
}
