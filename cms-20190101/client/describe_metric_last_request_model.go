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
	// Set the value to a collection of `key:value` pairs. Example: `{"userId":"120886317861****"}` or `{"instanceId":"i-2ze2d6j5uhg20x47****"}`.
	//
	// >  You can query a maximum of 50 instances in each request.
	//
	// example:
	//
	// [{"instanceId":"i-abcdefgh12****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end of the time range to query monitoring data.
	//
	// 	- For second-level data, the start time is obtained by comparing the time that is specified by the StartTime parameter and 20 minutes earlier of the time that is specified by the EndTime parameter. The earlier one of the compared points in time is used as the start time.
	//
	// 	- For minute-level data, the start time is obtained by comparing the time that is specified by the StartTime parameter and 2 hours earlier of the time that is specified by the EndTime parameter. The earlier one of the compared points in time is used as the start time.
	//
	// 	- For hour-level data, the start time is obtained by comparing the time that is specified by the StartTime parameter and two days earlier of the time that is specified by the EndTime parameter. The earlier one of the compared points in time is used as the start time.
	//
	// example:
	//
	// 2019-01-31 10:10:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expression that is used to calculate the query results in real time.
	//
	// example:
	//
	// {"groupby":["userId","instanceId"]}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// The number of entries per page.
	//
	// Default value: 1000. This value indicates that a maximum of 1,000 entries of monitoring data can be returned on each page.
	//
	// >  The maximum value of the Length parameter for each request is 1440.
	//
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The metric that is used to monitor the cloud service.
	//
	// For more information about metric names, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// For more information about the namespaces of cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pagination token.
	//
	// 	- If the number of results exceeds the maximum number of entries allowed on a single page, a pagination token is returned.
	//
	// 	- This token can be used as an input parameter to obtain the next page of results. If all results are obtained, no token is returned.
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
	// >
	//
	// 	- If this parameter is not specified, monitoring data is queried based on the period in which metric values are reported. The statistical period of metrics (`MetricName`) varies for each cloud service. The statistical period of metrics is displayed in the `MinPeriods` column on the **Metrics*	- page for each cloud service. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period   *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query monitoring data.
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
