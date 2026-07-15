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
	// > **NodeId*	- is required when specifying a sharded cluster instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2635****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-06-13T11:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data granularity of the performance metrics in seconds. Valid values: 5, 30, 60, 600, 1800, 3600, and 86400.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The performance metrics. For more information, see [Metrics](https://help.aliyun.com/document_detail/216973.html).
	//
	// > To specify multiple metrics, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// CpuUsage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of a mongos or shard node in the sharded cluster instance. This parameter lets you query the performance of a single node.
	//
	// > Available only when **DBInstanceId*	- is set to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp2287****
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The role of a node in a standalone or replica set instance. Valid values:
	//
	// - **Primary**: The primary node.
	//
	// - **Secondary**: A secondary node.
	//
	// > 	- Available only when **DBInstanceId*	- is set to the ID of a standalone or replica set instance.
	//
	// >
	//
	// > 	- If **DBInstanceId*	- is set to the ID of a standalone instance, this parameter only supports the value **Primary**.
	//
	// example:
	//
	// Primary
	ReplicaSetRole       *string `json:"ReplicaSetRole,omitempty" xml:"ReplicaSetRole,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role ID of a node in a standalone or replica set instance. To query the role ID, call the [DescribeReplicaSetRole](https://help.aliyun.com/document_detail/62134.html) operation.
	//
	// > Available only when **DBInstanceId*	- is set to the ID of a standalone or replica set instance.
	//
	// example:
	//
	// 6025****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The Search node ID.
	//
	// > Available only after the Search feature is enabled for the instance.
	//
	// example:
	//
	// dds-2zec12675c9e****-search
	SearchId *string `json:"SearchId,omitempty" xml:"SearchId,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
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
