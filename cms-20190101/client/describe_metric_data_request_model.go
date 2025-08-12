// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v string) *DescribeMetricDataRequest
	GetDimensions() *string
	SetEndTime(v string) *DescribeMetricDataRequest
	GetEndTime() *string
	SetExpress(v string) *DescribeMetricDataRequest
	GetExpress() *string
	SetLength(v string) *DescribeMetricDataRequest
	GetLength() *string
	SetMetricName(v string) *DescribeMetricDataRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeMetricDataRequest
	GetNamespace() *string
	SetPeriod(v string) *DescribeMetricDataRequest
	GetPeriod() *string
	SetRegionId(v string) *DescribeMetricDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeMetricDataRequest
	GetStartTime() *string
}

type DescribeMetricDataRequest struct {
	// The dimensions based on which the resources are queried.
	//
	// Set the value to a collection of key-value pairs. A typical key-value pair is `instanceId:i-2ze2d6j5uhg20x47****`.
	//
	// >  You can query a maximum of 50 instances in a single request.
	//
	// example:
	//
	// [{"instanceId": "i-abcdefgh12****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end of the time range to query monitoring data.
	//
	// 	- If the `StartTime` and `EndTime` parameters are not specified, the monitoring data of the last statistical period is queried.``
	//
	// 	- If the `StartTime` and `EndTime` parameters are specified, the monitoring data of the last statistical period in the specified time range is queried.```` The following examples demonstrate how to determine the period in which monitoring data is queried:
	//
	//     	- If you set the `Period` parameter to 15, the specified time range must be less than or equal to 20 minutes. For example, if you set the StartTime parameter to 2021-05-08 08:10:00 and the EndTime parameter to 2021-05-08 08:30:00, the monitoring data of the last 15 seconds in the time range is queried.
	//
	//     	- If you set the `Period` to 60 or 900, the specified time range must be less than or equal to 2 hours. For example, if you set the Period parameter to 60, the StartTime parameter to 2021-05-08 08:00:00, and the EndTime parameter to 2021-05-08 10:00:00, the monitoring data of the last 60 seconds in the time range is queried.
	//
	//     	- If you set the `Period` parameter to 3600, the specified time range must be less than or equal to 2 days. For example, if you set the StartTime parameter to 2021-05-08 08:00:00 and the EndTime parameter to 2021-05-10 08:00:00, the monitoring data of the last 3,600 seconds in the time range is queried.
	//
	// The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 UTC on Thursday, January 1, 1970.
	//
	// 	- UTC time: the UTC time that follows the YYYY-MM-DDThh:mm:ssZ format.
	//
	// >  We recommend that you use UNIX timestamps to prevent time zone-related issues.
	//
	// example:
	//
	// 1618368960000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expression that is used to compute the query results in real time.
	//
	// >  Only the `groupby` expression is supported. This expression is similar to the `GROUP BY` statement that is used in databases.
	//
	// example:
	//
	// {"groupby":["userId","instanceId"]}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// The number of entries per page.
	//
	// Default value: 1000.
	//
	// >  The maximum value of the Length parameter in a request is 1440.
	//
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The metric that is used to monitor the cloud service.
	//
	// For more information about the metrics of cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_idle
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
	// The statistical period of the metric.
	//
	// Valid values: 15, 60, 900, and 3600.
	//
	// Unit: seconds.
	//
	// >
	//
	// 	- If this parameter is not specified, monitoring data is queried based on the period in which metric values are reported.
	//
	// 	- For more information about the statistical period of a metric that is specified by the `MetricName` parameter, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period   *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query monitoring data.
	//
	// 	- If the `StartTime` and `EndTime` parameters are not specified, the monitoring data of the last statistical period is queried.``
	//
	// 	- If the `StartTime` and `EndTime` parameters are specified, the monitoring data of the last statistical period in the specified time range is queried.```` The following examples demonstrate how to determine the period in which monitoring data is queried:
	//
	//     	- If you set the `Period` parameter to 15, the specified time range must be less than or equal to 20 minutes. For example, if you set the StartTime parameter to 2021-05-08 08:10:00 and the EndTime parameter to 2021-05-08 08:30:00, the monitoring data of the last 15 seconds in the time range is queried.
	//
	//     	- If you set the `Period` to 60 or 900, the specified time range must be less than or equal to 2 hours. For example, if you set the Period parameter to 60, the StartTime parameter to 2021-05-08 08:00:00, and the EndTime parameter to 2021-05-08 10:00:00, the monitoring data of the last 60 seconds in the time range is queried.
	//
	//     	- If you set the `Period` parameter to 3600, the specified time range must be less than or equal to 2 days. For example, if you set the StartTime parameter to 2021-05-08 08:00:00 and the EndTime parameter to 2021-05-10 08:00:00, the monitoring data of the last 3,600 seconds in the time range is queried.
	//
	// The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 UTC on Thursday, January 1, 1970.
	//
	// 	- UTC time: the UTC time that follows the YYYY-MM-DDThh:mm:ssZ format.
	//
	// >
	//
	// 	- You must set the `StartTime` parameter to a point in time that is later than 00:00:00 UTC on Thursday, January 1, 1970. Otherwise, this parameter is invalid.
	//
	// 	- We recommend that you use UNIX timestamps to prevent time zone-related issues.
	//
	// example:
	//
	// 1618368900000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricDataRequest) GetExpress() *string {
	return s.Express
}

func (s *DescribeMetricDataRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeMetricDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricDataRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricDataRequest) SetDimensions(v string) *DescribeMetricDataRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricDataRequest) SetEndTime(v string) *DescribeMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricDataRequest) SetExpress(v string) *DescribeMetricDataRequest {
	s.Express = &v
	return s
}

func (s *DescribeMetricDataRequest) SetLength(v string) *DescribeMetricDataRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricDataRequest) SetMetricName(v string) *DescribeMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricDataRequest) SetNamespace(v string) *DescribeMetricDataRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricDataRequest) SetPeriod(v string) *DescribeMetricDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricDataRequest) SetRegionId(v string) *DescribeMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricDataRequest) SetStartTime(v string) *DescribeMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
