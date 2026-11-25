// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregate(v bool) *DescribeAgenticDBPerformanceRequest
	GetAggregate() *bool
	SetBranchId(v string) *DescribeAgenticDBPerformanceRequest
	GetBranchId() *string
	SetDBClusterId(v string) *DescribeAgenticDBPerformanceRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeAgenticDBPerformanceRequest
	GetEndTime() *string
	SetEndpointId(v string) *DescribeAgenticDBPerformanceRequest
	GetEndpointId() *string
	SetInterval(v string) *DescribeAgenticDBPerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribeAgenticDBPerformanceRequest
	GetKey() *string
	SetStartTime(v string) *DescribeAgenticDBPerformanceRequest
	GetStartTime() *string
}

type DescribeAgenticDBPerformanceRequest struct {
	// The alert aggregation method. Valid values vary by metric:
	//
	// - Value: The value of the HTTP status code.
	//
	// - Average: The average HTTP wait time.
	//
	// - Value: The value of the Telnet status code.
	//
	// - TelnetLatency: The average Telnet wait time.
	//
	// - Average: The average ping packet loss rate.
	//
	// example:
	//
	// Average
	Aggregate *bool `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
	// The anchor branch ID. If you do not specify this parameter, the main branch of the project is used by default.
	//
	// example:
	//
	// br-7g8h9i0j1k2l
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-01-15T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-3m4n5o6p7q8r
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The granularity of performance data. Valid values:
	//
	// - 5
	//
	// - 30
	//
	// - 60
	//
	// - 600
	//
	// - 1800
	//
	// - 3600
	//
	// - 86400
	//
	// example:
	//
	// 5
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The performance metrics to query. Separate multiple values with commas (,).
	//
	// Note
	//
	// You can specify up to 5 performance metrics.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDBCPUForPCU,PolarDBPCU,PolarDBMemoryForPCU,PolarDBQPSTPS,PolarDBConnections
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The start time of the query. Specify the time in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-01T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAgenticDBPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBPerformanceRequest) GetAggregate() *bool {
	return s.Aggregate
}

func (s *DescribeAgenticDBPerformanceRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBPerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAgenticDBPerformanceRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeAgenticDBPerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeAgenticDBPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeAgenticDBPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAgenticDBPerformanceRequest) SetAggregate(v bool) *DescribeAgenticDBPerformanceRequest {
	s.Aggregate = &v
	return s
}

func (s *DescribeAgenticDBPerformanceRequest) SetBranchId(v string) *DescribeAgenticDBPerformanceRequest {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBPerformanceRequest) SetDBClusterId(v string) *DescribeAgenticDBPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBPerformanceRequest) SetEndTime(v string) *DescribeAgenticDBPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAgenticDBPerformanceRequest) SetEndpointId(v string) *DescribeAgenticDBPerformanceRequest {
	s.EndpointId = &v
	return s
}

func (s *DescribeAgenticDBPerformanceRequest) SetInterval(v string) *DescribeAgenticDBPerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribeAgenticDBPerformanceRequest) SetKey(v string) *DescribeAgenticDBPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeAgenticDBPerformanceRequest) SetStartTime(v string) *DescribeAgenticDBPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAgenticDBPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
