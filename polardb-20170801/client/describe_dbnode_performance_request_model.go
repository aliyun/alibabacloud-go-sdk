// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBNodePerformanceRequest
	GetDBClusterId() *string
	SetDBNodeId(v string) *DescribeDBNodePerformanceRequest
	GetDBNodeId() *string
	SetEndTime(v string) *DescribeDBNodePerformanceRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDBNodePerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribeDBNodePerformanceRequest
	GetKey() *string
	SetStartTime(v string) *DescribeDBNodePerformanceRequest
	GetStartTime() *string
	SetType(v string) *DescribeDBNodePerformanceRequest
	GetType() *string
}

type DescribeDBNodePerformanceRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the node in the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The end of the time range to query. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:01Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The granularity of the performance data. Valid values:
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
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The performance metrics to query. Separate multiple metrics with commas (,). For more information, see [Performance metrics](https://help.aliyun.com/document_detail/141787.html).
	//
	// > - You can query a maximum of five performance metrics.
	//
	// >
	//
	// > - If your cluster has Serverless enabled for fixed specifications, querying PolarDBCPU or PolarDBMemory alone ignores the Interval parameter and returns performance metrics per second. To get data at your specified Interval, query multiple metrics.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDBDiskUsage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The beginning of the time range to query. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in Coordinated Universal Time (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A special metric. Currently, only orca is supported.
	//
	// example:
	//
	// orca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDBNodePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBNodePerformanceRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBNodePerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBNodePerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDBNodePerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBNodePerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBNodePerformanceRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDBNodePerformanceRequest) SetDBClusterId(v string) *DescribeDBNodePerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBNodeId(v string) *DescribeDBNodePerformanceRequest {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetEndTime(v string) *DescribeDBNodePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetInterval(v string) *DescribeDBNodePerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetKey(v string) *DescribeDBNodePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetStartTime(v string) *DescribeDBNodePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetType(v string) *DescribeDBNodePerformanceRequest {
	s.Type = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
