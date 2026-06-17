// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v string) *DescribeMetricListRequest
	GetDimensions() *string
	SetEndTime(v string) *DescribeMetricListRequest
	GetEndTime() *string
	SetExpress(v string) *DescribeMetricListRequest
	GetExpress() *string
	SetLength(v string) *DescribeMetricListRequest
	GetLength() *string
	SetMetricName(v string) *DescribeMetricListRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeMetricListRequest
	GetNamespace() *string
	SetNextToken(v string) *DescribeMetricListRequest
	GetNextToken() *string
	SetPeriod(v string) *DescribeMetricListRequest
	GetPeriod() *string
	SetRegionId(v string) *DescribeMetricListRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeMetricListRequest
	GetStartTime() *string
}

type DescribeMetricListRequest struct {
	// The dimensions that specify the resources to be monitored.
	//
	// Format: a collection of key-value pairs, such as `{"userId":"120886317861****"}` and `{"instanceId":"i-2ze2d6j5uhg20x47****"}`.
	//
	// > A single request can be used to query a maximum of 50 instances.
	//
	// example:
	//
	// [{"instanceId":"i-2ze2d6j5uhg20x47****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end of the time range to query. The following formats are supported:
	//
	// - UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// - Format: YYYY-MM-DD hh:mm:ss.
	//
	// > The interval between \\`StartTime\\` and \\`EndTime\\` must be less than or equal to 31 days.
	//
	// example:
	//
	// 2019-01-30 00:10:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expression that is used for real-time computing based on the query results.
	//
	// > Only the groupby expression is supported. This expression is similar to the GROUP BY statement in databases.
	//
	// example:
	//
	// {"groupby":["userId","instanceId"]}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// The number of entries to return on each page for a paged query.
	//
	// > The maximum value of \\`Length\\` in a single request is 1440.
	//
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The name of the metric.
	//
	// For more information, see [Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_idle
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// For more information, see [Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pagination cursor.
	//
	// > If you do not set this parameter, the first page of data is returned. If a value is returned for this parameter, it indicates that more data is available. To retrieve the next page, use the returned value as the \\`NextToken\\` in your next request. A null value indicates that all data has been retrieved.
	//
	// example:
	//
	// 15761485350009dd70bb64cff1f0fff750b08ffff073be5fb1e785e2b020f1a949d5ea14aea7fed82f01dd8****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The statistical period of the monitoring data.
	//
	// Valid values: 15, 60, 900, and 3600.
	//
	// Unit: seconds.
	//
	// > - If you do not set this parameter, the reporting period that was specified when the metric was registered is used.
	//
	// - The statistical period of each metric (`MetricName`) of a cloud service is different. For more information, see [Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period   *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The following formats are supported:
	//
	// - UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// - Format: YYYY-MM-DD hh:mm:ss.
	//
	// > 	- The time range is a left-open and right-closed interval. The value of \\`StartTime\\` must be earlier than the value of \\`EndTime\\`.
	//
	// - The interval between \\`StartTime\\` and \\`EndTime\\` must be less than or equal to 31 days.
	//
	// example:
	//
	// 2019-01-30 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricListRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricListRequest) GetExpress() *string {
	return s.Express
}

func (s *DescribeMetricListRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeMetricListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricListRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricListRequest) SetDimensions(v string) *DescribeMetricListRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricListRequest) SetEndTime(v string) *DescribeMetricListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricListRequest) SetExpress(v string) *DescribeMetricListRequest {
	s.Express = &v
	return s
}

func (s *DescribeMetricListRequest) SetLength(v string) *DescribeMetricListRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricListRequest) SetMetricName(v string) *DescribeMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricListRequest) SetNamespace(v string) *DescribeMetricListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricListRequest) SetNextToken(v string) *DescribeMetricListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricListRequest) SetPeriod(v string) *DescribeMetricListRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricListRequest) SetRegionId(v string) *DescribeMetricListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricListRequest) SetStartTime(v string) *DescribeMetricListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricListRequest) Validate() error {
	return dara.Validate(s)
}
