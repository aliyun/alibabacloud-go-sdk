// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricLastRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v string) *DescribeMetricLastRequest
	GetDimensions() *string
	SetEndTime(v string) *DescribeMetricLastRequest
	GetEndTime() *string
	SetExpress(v string) *DescribeMetricLastRequest
	GetExpress() *string
	SetLength(v string) *DescribeMetricLastRequest
	GetLength() *string
	SetMetricName(v string) *DescribeMetricLastRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeMetricLastRequest
	GetNamespace() *string
	SetNextToken(v string) *DescribeMetricLastRequest
	GetNextToken() *string
	SetPeriod(v string) *DescribeMetricLastRequest
	GetPeriod() *string
	SetRegionId(v string) *DescribeMetricLastRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeMetricLastRequest
	GetStartTime() *string
}

type DescribeMetricLastRequest struct {
	// The monitoring dimensions of the specified resource.
	//
	// Format: a collection of `key:value` pairs. Example: `{"userId":"120886317861****"}` or `{"instanceId":"i-2ze2d6j5uhg20x47****"}`.
	//
	// > A maximum of 50 instances can be queried in a single request.
	//
	// example:
	//
	// [{"instanceId":"i-2ze2d6j5uhg20x47****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end time for querying the metric.
	//
	// The system automatically limits the query start time (query window size) based on data granularity:
	//
	// 1. Second-level data: The maximum query span is 20 minutes (start time = Max(startTime, EndTime - 20min)).
	//
	// 2. Minute-level data: The maximum query span is 2 hours (start time = Max(startTime, EndTime - 2h)).
	//
	// 3. Hour-level data: The maximum query span is 2 days (start time = Max(startTime, EndTime - 2d)).
	//
	// > If the specified startTime is earlier than the limit, the system automatically uses the limit as the actual query start time.
	//
	// example:
	//
	// 2019-01-31 10:10:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expression for real-time computation on existing query results.
	//
	// example:
	//
	// {"groupby":["userId","instanceId"]}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// The number of entries per page for monitoring data. This parameter is used for paged query.
	//
	// Default value: 1000, which indicates 1000 entries of monitoring data per page.
	//
	// > The maximum value of Length in a single request is 1440.
	//
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The metric name of the cloud service.
	//
	// For information about metric names, see [Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// For information about namespaces of cloud services, see [Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pagination cursor.
	//
	// - If the query results that match the filter conditions exceed the page size, a pagination cursor is returned.
	//
	// - To retrieve the next page of data, use the returned cursor value as the request parameter until no cursor value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// 15761432850009dd70bb64cff1f0fff6c0b08ffff073be5fb1e785e2b020f7fed9b5e137bd810a6d6cff5ae****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The statistical period of the monitoring data.
	//
	// Valid values: 15, 60, 900, and 3600.
	//
	// Unit: seconds.
	//
	// > - If you do not set the statistical period, monitoring data is queried based on the reporting period that was specified when the metric was registered.
	//
	//     - The statistical period varies by metric (`MetricName`) for each cloud service. For more information, see the `Min Periods` column on the metric list page of the target cloud service in [Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period   *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time.
	//
	// Two formats are supported: 1. A UNIX timestamp string in milliseconds (ms), which represents the number of milliseconds elapsed since 1970-01-01 00:00:00 UTC. 2. A time string in the ISO 8601 format of YYYY-MM-DDThh:mm:ssZ in UTC.
	//
	// > The query range uses a left-exclusive and right-inclusive mode. StartTime cannot be equal to or greater than EndTime, and the interval between them must be less than or equal to 31 days.
	//
	// example:
	//
	// 2019-01-31 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricLastRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricLastRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricLastRequest) GetExpress() *string {
	return s.Express
}

func (s *DescribeMetricLastRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeMetricLastRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricLastRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricLastRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricLastRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricLastRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricLastRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricLastRequest) SetDimensions(v string) *DescribeMetricLastRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricLastRequest) SetEndTime(v string) *DescribeMetricLastRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricLastRequest) SetExpress(v string) *DescribeMetricLastRequest {
	s.Express = &v
	return s
}

func (s *DescribeMetricLastRequest) SetLength(v string) *DescribeMetricLastRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricLastRequest) SetMetricName(v string) *DescribeMetricLastRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricLastRequest) SetNamespace(v string) *DescribeMetricLastRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricLastRequest) SetNextToken(v string) *DescribeMetricLastRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricLastRequest) SetPeriod(v string) *DescribeMetricLastRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricLastRequest) SetRegionId(v string) *DescribeMetricLastRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricLastRequest) SetStartTime(v string) *DescribeMetricLastRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricLastRequest) Validate() error {
	return dara.Validate(s)
}
