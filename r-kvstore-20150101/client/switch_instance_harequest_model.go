// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceHARequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SwitchInstanceHARequest
	GetInstanceId() *string
	SetNodeId(v string) *SwitchInstanceHARequest
	GetNodeId() *string
	SetOwnerAccount(v string) *SwitchInstanceHARequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchInstanceHARequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SwitchInstanceHARequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchInstanceHARequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *SwitchInstanceHARequest
	GetSecurityToken() *string
	SetSourceNodeId(v string) *SwitchInstanceHARequest
	GetSourceNodeId() *string
	SetSwitchMode(v int32) *SwitchInstanceHARequest
	GetSwitchMode() *int32
	SetSwitchType(v string) *SwitchInstanceHARequest
	GetSwitchType() *string
	SetTargetNodeId(v string) *SwitchInstanceHARequest
	GetTargetNodeId() *string
	SetTargetShardName(v string) *SwitchInstanceHARequest
	GetTargetShardName() *string
}

type SwitchInstanceHARequest struct {
	// The instance ID. You can call [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the data shard node. You can call [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/473782.html) to obtain the CustinsId parameter. Separate multiple data shard node IDs with commas (,). To specify all nodes, enter `all`.
	//
	// > This parameter is available and required only when the instance uses the cluster or read/write splitting architecture.
	//
	// example:
	//
	// 56****19,56****20
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The node ID of the original MASTER node in the shard.
	//
	// example:
	//
	// 52717408
	SourceNodeId *string `json:"SourceNodeId,omitempty" xml:"SourceNodeId,omitempty"`
	// The execution time. Valid values:
	//
	// 	- **0**: immediately. This is the default value.
	//
	// 	- **1**: during the maintenance window.
	//
	// > You can call [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) to modify the maintenance window of the instance.
	//
	// example:
	//
	// 0
	SwitchMode *int32 `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
	// The switchover mode. Valid values:
	//
	// 	- **ReliabilityPriority (default)**: Reliability is prioritized. The primary/secondary switchover is performed only when primary/secondary synchronization has no latency, which prevents data loss. In scenarios with heavy write workloads and persistent synchronization latency, this mode may cause the primary/secondary switchover to fail.
	//
	// 	- **AvailablePriority**: Availability is prioritized. The primary/secondary switchover is performed immediately regardless of primary/secondary latency, which may cause minor data loss.
	//
	// > Evaluate your business requirements for data integrity and service availability before selecting a switchover mode.
	//
	// example:
	//
	// ReliabilityPriority
	SwitchType *string `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
	// The node ID of the target MASTER node after the switchover.
	//
	// example:
	//
	// 52717403
	TargetNodeId *string `json:"TargetNodeId,omitempty" xml:"TargetNodeId,omitempty"`
	// The shard name of the instance.
	//
	// example:
	//
	// r-2zegk3jyxxxwixfo6c-db-1
	TargetShardName *string `json:"TargetShardName,omitempty" xml:"TargetShardName,omitempty"`
}

func (s SwitchInstanceHARequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceHARequest) GoString() string {
	return s.String()
}

func (s *SwitchInstanceHARequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchInstanceHARequest) GetNodeId() *string {
	return s.NodeId
}

func (s *SwitchInstanceHARequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchInstanceHARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchInstanceHARequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchInstanceHARequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchInstanceHARequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SwitchInstanceHARequest) GetSourceNodeId() *string {
	return s.SourceNodeId
}

func (s *SwitchInstanceHARequest) GetSwitchMode() *int32 {
	return s.SwitchMode
}

func (s *SwitchInstanceHARequest) GetSwitchType() *string {
	return s.SwitchType
}

func (s *SwitchInstanceHARequest) GetTargetNodeId() *string {
	return s.TargetNodeId
}

func (s *SwitchInstanceHARequest) GetTargetShardName() *string {
	return s.TargetShardName
}

func (s *SwitchInstanceHARequest) SetInstanceId(v string) *SwitchInstanceHARequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetNodeId(v string) *SwitchInstanceHARequest {
	s.NodeId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetOwnerAccount(v string) *SwitchInstanceHARequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchInstanceHARequest) SetOwnerId(v int64) *SwitchInstanceHARequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetResourceOwnerAccount(v string) *SwitchInstanceHARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchInstanceHARequest) SetResourceOwnerId(v int64) *SwitchInstanceHARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSecurityToken(v string) *SwitchInstanceHARequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSourceNodeId(v string) *SwitchInstanceHARequest {
	s.SourceNodeId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSwitchMode(v int32) *SwitchInstanceHARequest {
	s.SwitchMode = &v
	return s
}

func (s *SwitchInstanceHARequest) SetSwitchType(v string) *SwitchInstanceHARequest {
	s.SwitchType = &v
	return s
}

func (s *SwitchInstanceHARequest) SetTargetNodeId(v string) *SwitchInstanceHARequest {
	s.TargetNodeId = &v
	return s
}

func (s *SwitchInstanceHARequest) SetTargetShardName(v string) *SwitchInstanceHARequest {
	s.TargetShardName = &v
	return s
}

func (s *SwitchInstanceHARequest) Validate() error {
	return dara.Validate(s)
}
