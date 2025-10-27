// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterResourcePoolPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest
	GetEndTime() *string
	SetKey(v string) *DescribeDBClusterResourcePoolPerformanceRequest
	GetKey() *string
	SetOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterResourcePoolPerformanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterResourcePoolPerformanceRequest
	GetResourceOwnerId() *int64
	SetResourcePools(v string) *DescribeDBClusterResourcePoolPerformanceRequest
	GetResourcePools() *string
	SetStartTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest
	GetStartTime() *string
}

type DescribeDBClusterResourcePoolPerformanceRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to monitor the resource group. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-06-10T07:01Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metrics of the resource group. You can enter multiple metrics at the same time to query the monitoring information. Separate multiple metrics with commas (,). Valid values:
	//
	// 	- **AnalyticDB_RP_CPU**: the average CPU utilization. Unit: %.
	//
	// 	- **AnalyticDB_RP_RT**: the query response time (RT). Unit: milliseconds.
	//
	// 	- **AnalyticDB_RP_QPS**: the queries per second (QPS). The value of this parameter must be a numeric value.
	//
	// 	- **AnalyticDB_RP_WaitTime**: the query waiting time. Unit: milliseconds.
	//
	// 	- **AnalyticDB_RP_OriginalNode**: the number of basic nodes in the resource group.
	//
	// 	- **AnalyticDB_RP_ActualNode**: the number of scheduled nodes that are scaled out in the resource group.
	//
	// 	- **AnalyticDB_RP_PlanNode**: the number of scheduled nodes to be scaled out in the resource group.
	//
	// 	- **AnalyticDB_RP_TotalNode**: the total number of nodes in the resource group. Total number of nodes = Number of basic nodes + Number of scheduled nodes that are scaled out.
	//
	// >
	//
	// 	- If you leave this parameter empty, the monitoring information about all metrics is returned.
	//
	// 	- For more information about scaling plans, see [Create a resource scaling plan](https://help.aliyun.com/document_detail/189507.html).
	//
	// example:
	//
	// AnalyticDB_RP_CPU
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The names of the resource groups that you want to query. You can enter multiple names of resource groups. Separate multiple names with commas (,).
	//
	// >
	//
	// 	- The value of this parameter is case-insensitive. For example, `USER_DEFAULT` and `user_default` specify the same resource group.
	//
	// 	- If you leave this parameter empty, the monitoring information about the `USER_DEFAULT` resource group is returned.
	//
	// example:
	//
	// TEST_POOL
	ResourcePools *string `json:"ResourcePools,omitempty" xml:"ResourcePools,omitempty"`
	// The beginning of the time range to monitor the resource group. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// > You can view only the monitoring information about the resource groups within the last two days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-06-10T07:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBClusterResourcePoolPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterResourcePoolPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetResourcePools() *string {
	return s.ResourcePools
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetEndTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetKey(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetOwnerId(v int64) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetResourcePools(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.ResourcePools = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) SetStartTime(v string) *DescribeDBClusterResourcePoolPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterResourcePoolPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
