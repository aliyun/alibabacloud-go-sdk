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
	// The dimensions that specify the resources whose monitoring data you want to query.
	//
	// Set the value to a collection of key-value pairs. A typical key-value pair is `instanceId:i-2ze2d6j5uhg20x47****`.
	//
	// >  You can query a maximum of 50 instances in a single request.
	//
	// example:
	//
	// [{"instanceId": "i-abcdefgh12****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end of the time range to query. The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 Thursday, January 1, 1970
	//
	// 	- UTC time: the UTC time that follows the YYYY-MM-DDThh:mm:ssZ format
	//
	// example:
	//
	// 2019-01-30 00:10:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expression that is used to compute the query results in real time.
	//
	// >  Only the groupby expression is supported. This expression is similar to the GROUP BY statement that is used in databases.
	//
	// example:
	//
	// {"groupby":["userId","instanceId"]}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// The number of entries to return on each page.
	//
	// >  The maximum value of the Length parameter in a request is 1440.
	//
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The name of the metric.
	//
	// For more information about metric names, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_idle
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service. Format: acs_service name.
	//
	// For more information about the namespaces of cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The paging token.
	//
	// >  If this parameter is not specified, the data on the first page is returned. A return value other than Null of this parameter indicates that not all entries have been returned. You can use this value as an input parameter to obtain entries on the next page. The value Null indicates that all query results have been returned.
	//
	// example:
	//
	// 15761485350009dd70bb64cff1f0fff750b08ffff073be5fb1e785e2b020f1a949d5ea14aea7fed82f01dd8****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The interval at which the monitoring data is queried.
	//
	// Valid values: 60, 300, and 900.
	//
	// Unit: seconds.
	//
	// >  Configure this parameter based on your business scenario.
	//
	// example:
	//
	// 60
	Period   *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 Thursday, January 1, 1970
	//
	// 	- UTC time: the UTC time that follows the YYYY-MM-DDThh:mm:ssZ format
	//
	// >  The specified period includes the end time and excludes the start time. The start time must be earlier than the end time.
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
