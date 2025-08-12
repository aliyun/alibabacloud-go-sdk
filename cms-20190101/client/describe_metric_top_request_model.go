// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricTopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v string) *DescribeMetricTopRequest
	GetDimensions() *string
	SetEndTime(v string) *DescribeMetricTopRequest
	GetEndTime() *string
	SetExpress(v string) *DescribeMetricTopRequest
	GetExpress() *string
	SetLength(v string) *DescribeMetricTopRequest
	GetLength() *string
	SetMetricName(v string) *DescribeMetricTopRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeMetricTopRequest
	GetNamespace() *string
	SetOrderDesc(v string) *DescribeMetricTopRequest
	GetOrderDesc() *string
	SetOrderby(v string) *DescribeMetricTopRequest
	GetOrderby() *string
	SetPeriod(v string) *DescribeMetricTopRequest
	GetPeriod() *string
	SetRegionId(v string) *DescribeMetricTopRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeMetricTopRequest
	GetStartTime() *string
}

type DescribeMetricTopRequest struct {
	// The monitoring dimensions of the specified resource.
	//
	// Set the value to a collection of `key:value` pairs. Example: `{"userId":"120886317861****"}` or `{"instanceId":"i-2ze2d6j5uhg20x47****"}`.
	//
	// >  You can query a maximum of 50 instances in each request.
	//
	// example:
	//
	// [{"instanceId": "i-abcdefgh12****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end of the time range to query monitoring data.
	//
	// 	- If the `StartTime` and `EndTime` parameters are not specified, the monitoring data of the last statistical period is queried.``
	//
	// 	- If the `StartTime` and `EndTime` parameters are specified, the monitoring data of the last statistical period in the specified time range is queried.````
	//
	//     	- If you set the `Period` parameter to 15, the specified time range must be less than or equal to 20 minutes. For example, if you set the StartTime parameter to 2021-05-08 08:10:00 and the EndTime parameter to 2021-05-08 08:30:00, the monitoring data of the last 15 seconds in the time range is queried.
	//
	//     	- If you set the `Period` parameter to 60 or 900, the specified time range must be less than or equal to 2 hours. For example, if you set the Period parameter to 60, the StartTime parameter to 2021-05-08 08:00:00, and the EndTime parameter to 2021-05-08 10:00:00, the monitoring data of the last 60 seconds in the time range is queried.
	//
	//     	- If you set the `Period` parameter to 3600, the specified time range must be less than or equal to two days. For example, if you set the StartTime parameter to 2021-05-08 08:00:00 and the EndTime parameter to 2021-05-10 08:00:00, the monitoring data of the last 3,600 seconds in the time range is queried.
	//
	// The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 Thursday, January 1, 1970
	//
	// 	- Time format: YYYY-MM-DDThh:mm:ssZ
	//
	// >  We recommend that you use UNIX timestamps to prevent time zone-related issues.
	//
	// example:
	//
	// 2021-05-08 10:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expression that is used to compute the query results in real time.
	//
	// >  Only the `groupby` expression is supported. This expression is similar to the GROUP BY statement used in databases.
	//
	// example:
	//
	// {"groupby":["userId","instanceId"]}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// >  The maximum value of the Length parameter in a request is 1440.
	//
	// example:
	//
	// 10
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The metric that is used to monitor the cloud service.
	//
	// For more information about metric names, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// The order in which data is sorted. Valid values:
	//
	// 	- True: sorts data in ascending order.
	//
	// 	- False (default): sorts data in descending order.
	//
	// example:
	//
	// False
	OrderDesc *string `json:"OrderDesc,omitempty" xml:"OrderDesc,omitempty"`
	// The field based on which data is sorted. Valid values:
	//
	// 	- Average
	//
	// 	- Minimum
	//
	// 	- Maximum
	//
	// This parameter is required.
	//
	// example:
	//
	// Average
	Orderby *string `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
	// The statistical period of the monitoring data.
	//
	// Valid values: 15, 60, 900, and 3600.
	//
	// Unit: seconds.
	//
	// >
	//
	// 	- If this parameter is not specified, monitoring data is queried based on the period in which metric values are reported.
	//
	// 	- Statistical periods vary based on the metrics that are specified by `MetricName`. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// 	- If the `StartTime` and `EndTime` parameters are specified, the monitoring data of the last statistical period in the specified time range is queried.````
	//
	//     	- If you set the `Period` parameter to 15, the specified time range must be less than or equal to 20 minutes. For example, if you set the StartTime parameter to 2021-05-08 08:10:00 and the EndTime parameter to 2021-05-08 08:30:00, the monitoring data of the last 15 seconds in the time range is queried.
	//
	//     	- If you set the `Period` parameter to 60 or 900, the specified time range must be less than or equal to 2 hours. For example, if you set the Period parameter to 60, the StartTime parameter to 2021-05-08 08:00:00, and the EndTime parameter to 2021-05-08 10:00:00, the monitoring data of the last 60 seconds in the time range is queried.
	//
	//     	- If you set the `Period` parameter to 3600, the specified time range must be less than or equal to two days. For example, if you set the StartTime parameter to 2021-05-08 08:00:00 and the EndTime parameter to 2021-05-10 08:00:00, the monitoring data of the last 3,600 seconds in the time range is queried.
	//
	// The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 Thursday, January 1, 1970
	//
	// 	- Time format: YYYY-MM-DDThh:mm:ssZ
	//
	// >
	//
	// 	- You must set the `StartTime` parameter to a point in time that is later than 00:00:00 Thursday, January 1, 1970. Otherwise, this parameter is invalid.
	//
	// 	- We recommend that you use UNIX timestamps to prevent time zone-related issues.
	//
	// example:
	//
	// 2021-05-08 08:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricTopRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricTopRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricTopRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricTopRequest) GetExpress() *string {
	return s.Express
}

func (s *DescribeMetricTopRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeMetricTopRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricTopRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricTopRequest) GetOrderDesc() *string {
	return s.OrderDesc
}

func (s *DescribeMetricTopRequest) GetOrderby() *string {
	return s.Orderby
}

func (s *DescribeMetricTopRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricTopRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricTopRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricTopRequest) SetDimensions(v string) *DescribeMetricTopRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricTopRequest) SetEndTime(v string) *DescribeMetricTopRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricTopRequest) SetExpress(v string) *DescribeMetricTopRequest {
	s.Express = &v
	return s
}

func (s *DescribeMetricTopRequest) SetLength(v string) *DescribeMetricTopRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricTopRequest) SetMetricName(v string) *DescribeMetricTopRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricTopRequest) SetNamespace(v string) *DescribeMetricTopRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricTopRequest) SetOrderDesc(v string) *DescribeMetricTopRequest {
	s.OrderDesc = &v
	return s
}

func (s *DescribeMetricTopRequest) SetOrderby(v string) *DescribeMetricTopRequest {
	s.Orderby = &v
	return s
}

func (s *DescribeMetricTopRequest) SetPeriod(v string) *DescribeMetricTopRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricTopRequest) SetRegionId(v string) *DescribeMetricTopRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricTopRequest) SetStartTime(v string) *DescribeMetricTopRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricTopRequest) Validate() error {
	return dara.Validate(s)
}
