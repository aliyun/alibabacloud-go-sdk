// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceHARequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *SwitchDBInstanceHARequest
	GetDBInstanceId() *string
	SetNodeId(v string) *SwitchDBInstanceHARequest
	GetNodeId() *string
	SetOwnerAccount(v string) *SwitchDBInstanceHARequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchDBInstanceHARequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SwitchDBInstanceHARequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchDBInstanceHARequest
	GetResourceOwnerId() *int64
	SetRoleIds(v string) *SwitchDBInstanceHARequest
	GetRoleIds() *string
	SetSwitchMode(v int32) *SwitchDBInstanceHARequest
	GetSwitchMode() *int32
}

type SwitchDBInstanceHARequest struct {
	// The ID of the instance
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the shard node in the sharded cluster instance.
	//
	// > You must specify this parameter if you set the **DBInstanceId*	- parameter to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bpxxxxxxxx
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the roles who switch the primary and secondary nodes for the instance. You can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/123802.html) operation to view the IDs and information of roles of nodes.
	//
	// >
	//
	// 	- Separate role IDs with commas (,). If this parameter is not specified, the primary and secondary nodes are switched.
	//
	// 	- If you set the **DBInstanceId*	- parameter to the ID of a sharded cluster instance, the roles who switch the primary and secondary nodes for the instance must belong to one shard node.
	//
	// example:
	//
	// 972xxxx,972xxxx
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The time when the primary and secondary nodes are switched. Valid values:
	//
	// 	- 0: The primary and secondary nodes are immediately switched.
	//
	// 	- 1: The primary and secondary nodes are switched during the O\\&M time period.
	//
	// example:
	//
	// 0
	SwitchMode *int32 `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s SwitchDBInstanceHARequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceHARequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHARequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SwitchDBInstanceHARequest) GetNodeId() *string {
	return s.NodeId
}

func (s *SwitchDBInstanceHARequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchDBInstanceHARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchDBInstanceHARequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchDBInstanceHARequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchDBInstanceHARequest) GetRoleIds() *string {
	return s.RoleIds
}

func (s *SwitchDBInstanceHARequest) GetSwitchMode() *int32 {
	return s.SwitchMode
}

func (s *SwitchDBInstanceHARequest) SetDBInstanceId(v string) *SwitchDBInstanceHARequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetNodeId(v string) *SwitchDBInstanceHARequest {
	s.NodeId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetOwnerAccount(v string) *SwitchDBInstanceHARequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetOwnerId(v int64) *SwitchDBInstanceHARequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetResourceOwnerAccount(v string) *SwitchDBInstanceHARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetResourceOwnerId(v int64) *SwitchDBInstanceHARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetRoleIds(v string) *SwitchDBInstanceHARequest {
	s.RoleIds = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetSwitchMode(v int32) *SwitchDBInstanceHARequest {
	s.SwitchMode = &v
	return s
}

func (s *SwitchDBInstanceHARequest) Validate() error {
	return dara.Validate(s)
}
