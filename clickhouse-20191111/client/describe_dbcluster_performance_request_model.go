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
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time. The interval cannot be more than 32 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-27T16:38Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metrics that you want to query. Separate multiple performance metrics with commas (,). You can query up to five performance metrics at a time. You can query the following performance metrics:
	//
	// >  The **Key*	- parameter is required.
	//
	// 	- **CPU**:
	//
	//     	- **CPU_USAGE**: the CPU utilization
	//
	// 	- **Memory**:
	//
	//     	- **MEM_USAGE**: the memory usage
	//
	//     	- **MEM_USAGE_SIZE**: the used memory. Unit: MB
	//
	// 	- **Disk**:
	//
	//     	- **DISK_USAGE**: the disk usage
	//
	//     	- **DISK_USAGE_SIZE**: the size of the used disks. Unit: MB
	//
	//     	- **IOPS**: the disk Input/Output Operations per Second (IOPS)
	//
	// 	- **Connection**:
	//
	//     	- **CONN_USAGE**: the database connection usage
	//
	//     	- **CONN_USAGE_COUNT**: the number of database connections used
	//
	// 	- **Write**:
	//
	//     	- **TPS**: the number of rows written per second
	//
	//     	- **INSERT_SIZE**: the amount of data written per second. Unit: MB
	//
	// 	- **Query**:
	//
	//     	- **QPS**: the queries per second
	//
	//     	- **AVG_SEEK**: the average number of random seek calls
	//
	// 	- **WAIT**:
	//
	//     	- **ZK_WAIT**: the average ZooKeeper wait time. Unit: ms
	//
	//     	- **IO_WAIT**: the average I/O wait time. Unit: ms
	//
	//     	- **CPU_WAIT**: the average CPU wait time. Unit: ms
	//
	// example:
	//
	// MEM_USAGE
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in Coordinated Universal Time (UTC).
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
