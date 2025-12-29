// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstancePerformanceRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeDBInstancePerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribeDBInstancePerformanceRequest
	GetKey() *string
	SetNodeId(v string) *DescribeDBInstancePerformanceRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeDBInstancePerformanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancePerformanceRequest
	GetOwnerId() *int64
	SetReplicaSetRole(v string) *DescribeDBInstancePerformanceRequest
	GetReplicaSetRole() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancePerformanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancePerformanceRequest
	GetResourceOwnerId() *int64
	SetRoleId(v string) *DescribeDBInstancePerformanceRequest
	GetRoleId() *string
	SetSearchId(v string) *DescribeDBInstancePerformanceRequest
	GetSearchId() *string
	SetStartTime(v string) *DescribeDBInstancePerformanceRequest
	GetStartTime() *string
}

type DescribeDBInstancePerformanceRequest struct {
	// The instance ID.
	//
	// >  If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2635****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-06-13T11:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The interval at which performance data is collected. Valid values: 5, 30, 60, 600, 1800, 3600, 86400.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The performance metric. For more information about valid values, see [Monitoring items and metrics](https://help.aliyun.com/document_detail/216973.html).
	//
	// >  If you need to specify multiple metrics, separate the metrics with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// CpuUsage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the mongos or shard node in a sharded cluster instance. You can specify this parameter to view the performance data of a single node.
	//
	// >  This parameter is valid when you set the **DBInstanceId*	- parameter to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp2287****
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The role of the node in the standalone or replica set instance. Valid values:
	//
	// 	- **Primary**
	//
	// 	- **Secondary**
	//
	// >  	- This parameter is valid only when you specify the **DBInstanceId*	- parameter to the ID of a standalone instance or a replica set instance.
	//
	// > 	- This parameter can be set only to **Primary*	- when you specify the **DBInstanceId*	- parameter to the ID of a standalone instance.
	//
	// example:
	//
	// Primary
	ReplicaSetRole       *string `json:"ReplicaSetRole,omitempty" xml:"ReplicaSetRole,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role ID of the node in a standalone or replica set instance. You can call the [DescribeReplicaSetRole](https://help.aliyun.com/document_detail/62134.html) operation to query the role ID of the node.
	//
	// >  This parameter is available when you set the **DBInstanceId*	- parameter to the ID of a standalone instance or a replica set instance.
	//
	// example:
	//
	// 6025****
	RoleId   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	SearchId *string `json:"SearchId,omitempty" xml:"SearchId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-06-13T10:58Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstancePerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeDBInstancePerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancePerformanceRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstancePerformanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancePerformanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancePerformanceRequest) GetReplicaSetRole() *string {
	return s.ReplicaSetRole
}

func (s *DescribeDBInstancePerformanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancePerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancePerformanceRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *DescribeDBInstancePerformanceRequest) GetSearchId() *string {
	return s.SearchId
}

func (s *DescribeDBInstancePerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstancePerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetInterval(v string) *DescribeDBInstancePerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetNodeId(v string) *DescribeDBInstancePerformanceRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetOwnerAccount(v string) *DescribeDBInstancePerformanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetOwnerId(v int64) *DescribeDBInstancePerformanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetReplicaSetRole(v string) *DescribeDBInstancePerformanceRequest {
	s.ReplicaSetRole = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancePerformanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBInstancePerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetRoleId(v string) *DescribeDBInstancePerformanceRequest {
	s.RoleId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetSearchId(v string) *DescribeDBInstancePerformanceRequest {
	s.SearchId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
