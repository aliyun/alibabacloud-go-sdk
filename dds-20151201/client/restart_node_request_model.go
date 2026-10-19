// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RestartNodeRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *RestartNodeRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *RestartNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartNodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestartNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartNodeRequest
	GetResourceOwnerId() *int64
	SetRoleId(v string) *RestartNodeRequest
	GetRoleId() *string
	SetSwitchMode(v string) *RestartNodeRequest
	GetSwitchMode() *string
}

type RestartNodeRequest struct {
	// The instance ID.
	//
	// > If the instance is a sharded cluster instance, you must also specify `NodeId`.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Instance ID of the Mongos node, Shard node, or ConfigServer node in the sharded cluster instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query instance IDs of Mongos nodes, Shard nodes, or ConfigServer nodes.
	//
	// > This parameter is required when **DBInstanceId*	- is set to instance ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp128a003436****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role ID of the node.
	//
	// 1.  You can call the [DescribeReplicaSetRole](https://help.aliyun.com/document_detail/468469.html) operation to query the node IDs of a replica set instance.
	//
	// 2.  You can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/468472.html) operation to query the node IDs of a sharded cluster instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6025****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The execution time. Valid values:
	//
	// 	- **0**: immediately. This is the default value.
	//
	// 	- **1**: during the maintenance window.
	//
	// > You can call the [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) operation to modify the maintenance window of the instance.
	//
	// example:
	//
	// 1
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s RestartNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartNodeRequest) GoString() string {
	return s.String()
}

func (s *RestartNodeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestartNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *RestartNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartNodeRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *RestartNodeRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *RestartNodeRequest) SetDBInstanceId(v string) *RestartNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartNodeRequest) SetNodeId(v string) *RestartNodeRequest {
	s.NodeId = &v
	return s
}

func (s *RestartNodeRequest) SetOwnerAccount(v string) *RestartNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartNodeRequest) SetOwnerId(v int64) *RestartNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartNodeRequest) SetResourceOwnerAccount(v string) *RestartNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartNodeRequest) SetResourceOwnerId(v int64) *RestartNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartNodeRequest) SetRoleId(v string) *RestartNodeRequest {
	s.RoleId = &v
	return s
}

func (s *RestartNodeRequest) SetSwitchMode(v string) *RestartNodeRequest {
	s.SwitchMode = &v
	return s
}

func (s *RestartNodeRequest) Validate() error {
	return dara.Validate(s)
}
