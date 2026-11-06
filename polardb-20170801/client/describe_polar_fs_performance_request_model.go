// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePolarFsPerformanceRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribePolarFsPerformanceRequest
	GetEndTime() *string
	SetInterval(v string) *DescribePolarFsPerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribePolarFsPerformanceRequest
	GetKey() *string
	SetPolarFsInstanceId(v string) *DescribePolarFsPerformanceRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *DescribePolarFsPerformanceRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribePolarFsPerformanceRequest
	GetStartTime() *string
}

type DescribePolarFsPerformanceRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all Data Warehouse Edition clusters in the specified region.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time must be in UTC and must not be earlier than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-07T03:04:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sampling granularity of performance data, in seconds. The value must be a multiple of 60 and greater than 0. Default value: 60.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The monitoring metric group. Valid values: Capacity (capacity), Inode (file nodes), Iops (read and write operations per second), Latency (read and write latency), QpsMeta (metadata operation requests), and Throughput (read and write throughput).
	//
	// This parameter is required.
	//
	// example:
	//
	// Capacity
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-07T02:04:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePolarFsPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarFsPerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarFsPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePolarFsPerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribePolarFsPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribePolarFsPerformanceRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsPerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarFsPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePolarFsPerformanceRequest) SetDBClusterId(v string) *DescribePolarFsPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarFsPerformanceRequest) SetEndTime(v string) *DescribePolarFsPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePolarFsPerformanceRequest) SetInterval(v string) *DescribePolarFsPerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribePolarFsPerformanceRequest) SetKey(v string) *DescribePolarFsPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribePolarFsPerformanceRequest) SetPolarFsInstanceId(v string) *DescribePolarFsPerformanceRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsPerformanceRequest) SetRegionId(v string) *DescribePolarFsPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolarFsPerformanceRequest) SetStartTime(v string) *DescribePolarFsPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePolarFsPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
