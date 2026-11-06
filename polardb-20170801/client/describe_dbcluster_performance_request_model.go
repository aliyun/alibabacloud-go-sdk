// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeDBClusterPerformanceRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDBClusterPerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribeDBClusterPerformanceRequest
	GetKey() *string
	SetStartTime(v string) *DescribeDBClusterPerformanceRequest
	GetStartTime() *string
	SetSubGroupName(v string) *DescribeDBClusterPerformanceRequest
	GetSubGroupName() *string
	SetType(v string) *DescribeDBClusterPerformanceRequest
	GetType() *string
}

type DescribeDBClusterPerformanceRequest struct {
	// The ID of the PolarDB cluster. Before you query performance data, the cluster must be in the Running state. You can call the DescribeDBClusters operation and check the DBClusterStatus parameter to confirm the cluster status.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data granularity of the performance data. Valid values:
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
	// The performance metrics to query. Separate multiple values with commas (,). For more information, see [Performance metrics](https://help.aliyun.com/document_detail/141787.html).
	//
	// > You can specify up to 5 performance metrics.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDBDiskUsage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The start time of the query. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the subgroup for the performance metrics. Specify this parameter only when the queried metrics contain a subgroup.
	//
	// example:
	//
	// example-name
	SubGroupName *string `json:"SubGroupName,omitempty" xml:"SubGroupName,omitempty"`
	// The query type.
	//
	// example:
	//
	// orca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterPerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDBClusterPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterPerformanceRequest) GetSubGroupName() *string {
	return s.SubGroupName
}

func (s *DescribeDBClusterPerformanceRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetInterval(v string) *DescribeDBClusterPerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetSubGroupName(v string) *DescribeDBClusterPerformanceRequest {
	s.SubGroupName = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetType(v string) *DescribeDBClusterPerformanceRequest {
	s.Type = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
