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
	SetOwnerAccount(v string) *DescribeDBClusterPerformanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterPerformanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterPerformanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterPerformanceRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDBClusterPerformanceRequest
	GetStartTime() *string
}

type DescribeDBClusterPerformanceRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp125e3uu94wo****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. Format: yyyy-MM-ddTHH:mmZ (UTC time).
	//
	// >The end time must be later than the start time, and the time interval between the start time and the end time cannot exceed 32 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-27T16:38Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The names of the performance metrics to query. Separate multiple values with commas (,). You can query up to 5 performance metrics at a time. The following performance metrics are supported:
	//
	// >**Key*	- is required.
	//
	// - **CPU**:
	//
	//     - **CPU_USAGE**: CPU utilization.
	//
	// - **Memory**:
	//
	//     - **MEM_USAGE**: Memory utilization.
	//
	//     - **MEM_USAGE_SIZE**: Memory usage (unit: MB).
	//
	// - **Disk**:
	//
	//     - **DISK_USAGE**: Disk utilization.
	//
	//     - **DISK_USAGE_SIZE**: Disk usage (unit: MB).
	//
	//     - **IOPS**: Disk IOPS.
	//
	// - **Connections**:
	//
	//     - **CONN_USAGE**: Database connection utilization.
	//
	//     - **CONN_USAGE_COUNT**: Number of database connections in use.
	//
	// - **Writes**:
	//
	//     - **TPS**: Number of rows written per second.
	//
	//     - **INSERT_SIZE**: Size of data written per second (unit: MB).
	//
	// - **Queries**:
	//
	//     - **QPS**: Queries per second (QPS).
	//
	//     - **AVG_SEEK**: Number of random SEEK calls.
	//
	// - **WAIT**:
	//
	//     - **ZK_WAIT**: Average wait time for ZooKeeper (unit: ms).
	//
	//     - **IO_WAIT**: Average wait time for IO_WAIT (unit: ms).
	//
	//     - **CPU_WAIT**: Average wait time for CPU_WAIT (unit: ms).
	//
	// example:
	//
	// MEM_USAGE
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time of the query. Format: yyyy-MM-ddTHH:mmZ (UTC time).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-27T16:37Z
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

func (s *DescribeDBClusterPerformanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterPerformanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterPerformanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterPerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *DescribeDBClusterPerformanceRequest) SetOwnerAccount(v string) *DescribeDBClusterPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetOwnerId(v int64) *DescribeDBClusterPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBClusterPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
