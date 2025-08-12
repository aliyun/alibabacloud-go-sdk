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
	// The end timestamp of the alert logs to be queried.
	//
	// Unit: milliseconds.
	//
	// You can query only the alert logs within the last year. If the query time is longer than one year, the return value of the `AlertLogList` parameter is empty.
	//
	// >  The time period between the start time specified by `StartTime` and end time specified by `EndTime` must be less than or equal to 15 days. You must specify StartTime and EndTime at the same time, or leave StartTime and EndTime empty at the same time. If you do not specify this parameter, the alert logs within the last 15 minutes are queried by default.
	//
	// example:
	//
	// 1610074409694
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the alert event. Valid values:
	//
	// 	- TRIGGERED: The alert is triggered.
	//
	// 	- RESOLVED: The alert is resolved.
	//
	// example:
	//
	// TRIGGERED
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The dimensions based on which data is aggregated. This parameter is equivalent to the GROUP BY clause in SQL. Valid values:
	//
	// 	- `product`: aggregates data by cloud service.
	//
	// 	- `level`: aggregates data by alert level.
	//
	// 	- `groupId`: aggregates data by application group.
	//
	// 	- `contactGroup`: aggregates data by alert contact group.
	//
	// 	- `product,metricName`: aggregates data both by cloud service and by metric.
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
	// The statistical period of alert logs. Unit: minutes.
	//
	// example:
	//
	// 360
	LastMin *string `json:"LastMin,omitempty" xml:"LastMin,omitempty"`
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
	// > For more information about the metrics of different cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// IntranetInRate
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// >  For information about how to query the namespace of a cloud service, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// The abbreviation of the service name.
	//
	// For information about how to obtain the abbreviation of a cloud service name, see [DescribeProductsOfActiveMetricRule](https://help.aliyun.com/document_detail/114930.html).
	//
	// example:
	//
	// ECS
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule.
	//
	// For information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
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
	// The search keyword that is used to query alert logs.
	//
	// example:
	//
	// alert
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The status of the alert. Valid values:
	//
	// 	- 0: The alert is triggered or cleared.
	//
	// 	- 1: The alert is ineffective.
	//
	// 	- 2: The alert is muted.
	//
	// 	- 3: The host is restarting.
	//
	// 	- 4: No alert notification is sent.
	//
	// If the value of the SendStatus parameter is 0, the value P4 of the Level parameter indicates a triggered alert and the value OK indicates a cleared alert.
	//
	// example:
	//
	// 0
	SendStatus *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// The type of the alert rule. Valid value: METRIC. This value indicates an alert rule for time series metrics.
	//
	// example:
	//
	// METRIC
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start timestamp of the alert logs to be queried.
	//
	// Unit: milliseconds.
	//
	// You can query only the alert logs within the last year. If the query time is longer than one year, the return value of the `AlertLogList` parameter is empty.
	//
	// >  The time period between the start time specified by `StartTime` and the end time specified by `EndTime` must be less than or equal to 15 days. You must specify StartTime and EndTime at the same time, or leave StartTime and EndTime empty at the same time. If you do not specify this parameter, the alert logs within the last 15 minutes are queried by default.
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
