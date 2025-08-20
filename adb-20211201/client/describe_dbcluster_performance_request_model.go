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
	SetKey(v string) *DescribeDBClusterPerformanceRequest
	GetKey() *string
	SetRegionId(v string) *DescribeDBClusterPerformanceRequest
	GetRegionId() *string
	SetResourcePools(v string) *DescribeDBClusterPerformanceRequest
	GetResourcePools() *string
	SetStartTime(v string) *DescribeDBClusterPerformanceRequest
	GetStartTime() *string
}

type DescribeDBClusterPerformanceRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](~~~612397~~~) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1hx5n1o8f61****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// > The end time must be later than the start time. The maximum time range that can be specified is two days.
	//
	// example:
	//
	// 2022-03-11T15:01Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The key of the performance metric that you want to query. Separate multiple keys with commas (,). For more information about the performance metrics, see [Metric overview](https://help.aliyun.com/document_detail/2863211.html).
	//
	// example:
	//
	// AnalyticDB_CPU_Usage_Percentage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// user_default
	ResourcePools *string `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2022-03-10T23:56Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *DescribeDBClusterPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterPerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterPerformanceRequest) GetResourcePools() *string {
	return s.ResourcePools
}

func (s *DescribeDBClusterPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetRegionId(v string) *DescribeDBClusterPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourcePools(v string) *DescribeDBClusterPerformanceRequest {
	s.ResourcePools = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
