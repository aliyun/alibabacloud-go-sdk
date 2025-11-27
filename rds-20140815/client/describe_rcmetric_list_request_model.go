// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCMetricListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v string) *DescribeRCMetricListRequest
	GetDimensions() *string
	SetEndTime(v string) *DescribeRCMetricListRequest
	GetEndTime() *string
	SetExpress(v string) *DescribeRCMetricListRequest
	GetExpress() *string
	SetInstanceId(v string) *DescribeRCMetricListRequest
	GetInstanceId() *string
	SetLength(v string) *DescribeRCMetricListRequest
	GetLength() *string
	SetMetricName(v string) *DescribeRCMetricListRequest
	GetMetricName() *string
	SetNextToken(v string) *DescribeRCMetricListRequest
	GetNextToken() *string
	SetPeriod(v string) *DescribeRCMetricListRequest
	GetPeriod() *string
	SetRegionId(v string) *DescribeRCMetricListRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeRCMetricListRequest
	GetStartTime() *string
}

type DescribeRCMetricListRequest struct {
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Example: `2024-08-06 10:15:00`.
	//
	// example:
	//
	// 2024-08-06 10:15:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The metric that you want to use. For more information, see [CloudMonitor metrics](https://cms.console.aliyun.com/metric-meta/acs_ecs_dashboard/ecs).
	//
	// This parameter is required.
	//
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 6178f1825f9fb76ce0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The statistical period of the monitoring data.
	//
	// Set the value to 60 or an integer multiple of 60.
	//
	// Unit: seconds.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Example: `2024-08-06 10:05:00`.
	//
	// example:
	//
	// 2024-08-06 10:05:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRCMetricListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCMetricListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCMetricListRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeRCMetricListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRCMetricListRequest) GetExpress() *string {
	return s.Express
}

func (s *DescribeRCMetricListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCMetricListRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeRCMetricListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeRCMetricListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCMetricListRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeRCMetricListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCMetricListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRCMetricListRequest) SetDimensions(v string) *DescribeRCMetricListRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetEndTime(v string) *DescribeRCMetricListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetExpress(v string) *DescribeRCMetricListRequest {
	s.Express = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetInstanceId(v string) *DescribeRCMetricListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetLength(v string) *DescribeRCMetricListRequest {
	s.Length = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetMetricName(v string) *DescribeRCMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetNextToken(v string) *DescribeRCMetricListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetPeriod(v string) *DescribeRCMetricListRequest {
	s.Period = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetRegionId(v string) *DescribeRCMetricListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCMetricListRequest) SetStartTime(v string) *DescribeRCMetricListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRCMetricListRequest) Validate() error {
	return dara.Validate(s)
}
