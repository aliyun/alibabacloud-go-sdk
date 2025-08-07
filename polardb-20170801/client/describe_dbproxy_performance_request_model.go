// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBProxyPerformanceRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *DescribeDBProxyPerformanceRequest
	GetDBEndpointId() *string
	SetDBNodeId(v string) *DescribeDBProxyPerformanceRequest
	GetDBNodeId() *string
	SetEndTime(v string) *DescribeDBProxyPerformanceRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDBProxyPerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribeDBProxyPerformanceRequest
	GetKey() *string
	SetStartTime(v string) *DescribeDBProxyPerformanceRequest
	GetStartTime() *string
	SetType(v string) *DescribeDBProxyPerformanceRequest
	GetType() *string
}

type DescribeDBProxyPerformanceRequest struct {
	// The ID of cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// pe-****************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The ID of the node in the cluster. This parameter can be used to query the performance metrics of PolarProxy on different nodes. The following metrics are supported: PolarProxy_DBConns, PolarProxy_DBQps, and PolarProxy_DBActionOps.
	//
	// example:
	//
	// pi-******************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The end of the time range to query. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-24T02:08Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The interval at which performance data is collected. Valid values: 5, 30, 60, 600, 1800, 3600, 86400, in seconds.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The performance metrics that you want to query. Separate multiple indicators with commas (,). For more information, see [Performance parameters](https://help.aliyun.com/document_detail/141787.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarProxy_CpuUsage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The beginning of the time range to query. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Special metric. Set the value to tair, which indicates the PolarTair architecture.
	//
	// example:
	//
	// tair
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDBProxyPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBProxyPerformanceRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribeDBProxyPerformanceRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBProxyPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBProxyPerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDBProxyPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBProxyPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBProxyPerformanceRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDBProxyPerformanceRequest) SetDBClusterId(v string) *DescribeDBProxyPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetDBEndpointId(v string) *DescribeDBProxyPerformanceRequest {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetDBNodeId(v string) *DescribeDBProxyPerformanceRequest {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetEndTime(v string) *DescribeDBProxyPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetInterval(v string) *DescribeDBProxyPerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetKey(v string) *DescribeDBProxyPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetStartTime(v string) *DescribeDBProxyPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) SetType(v string) *DescribeDBProxyPerformanceRequest {
	s.Type = &v
	return s
}

func (s *DescribeDBProxyPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
