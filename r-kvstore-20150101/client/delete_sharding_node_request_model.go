// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteShardingNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveTime(v string) *DeleteShardingNodeRequest
	GetEffectiveTime() *string
	SetForceTrans(v bool) *DeleteShardingNodeRequest
	GetForceTrans() *bool
	SetInstanceId(v string) *DeleteShardingNodeRequest
	GetInstanceId() *string
	SetNodeId(v string) *DeleteShardingNodeRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DeleteShardingNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteShardingNodeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteShardingNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteShardingNodeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteShardingNodeRequest
	GetSecurityToken() *string
	SetShardCount(v int32) *DeleteShardingNodeRequest
	GetShardCount() *int32
}

type DeleteShardingNodeRequest struct {
	// The time when the proxy deletion takes effect if proxy nodes are deleted when you delete shards from a cluster instance in proxy mode. Valid values:
	//
	// 	- **0 or Immediately**: Takes effect immediately. This is the default value.
	//
	// 	- **1 or MaintainTime**: Takes effect during the maintenance window.
	//
	// > You can call the [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) operation to modify the maintenance window of the instance.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// Specifies whether to enable forced migration. Valid values:
	//
	// - **false*	- (default): Before the upgrade/downgrade, the system checks the current minor engine version of the instance. If the kernel version is too low, an error is reported. You must upgrade the minor engine version and retry.
	//
	// - **true**: Skips the checks and directly executes the upgrade/downgrade operation.
	//
	// example:
	//
	// false
	ForceTrans *bool `json:"ForceTrans,omitempty" xml:"ForceTrans,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// The ID of the data shard node to be deleted. To specify multiple shard node IDs, separate them with commas (,).
	//
	// > If both the NodeId and ShardCount parameters are specified, the NodeId parameter takes precedence.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0,r-bp1zxszhcgatnx****-db-1
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards to be deleted. The system deletes the specified number of shard nodes starting from the end.
	//
	// > For example, if an instance has five shard nodes: db-0, db-1, db-2, db-3, and db-4, and you set this parameter to 2, the db-3 and db-4 nodes are deleted.
	//
	// example:
	//
	// 1
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
}

func (s DeleteShardingNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteShardingNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteShardingNodeRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *DeleteShardingNodeRequest) GetForceTrans() *bool {
	return s.ForceTrans
}

func (s *DeleteShardingNodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteShardingNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteShardingNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteShardingNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteShardingNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteShardingNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteShardingNodeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteShardingNodeRequest) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *DeleteShardingNodeRequest) SetEffectiveTime(v string) *DeleteShardingNodeRequest {
	s.EffectiveTime = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetForceTrans(v bool) *DeleteShardingNodeRequest {
	s.ForceTrans = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetInstanceId(v string) *DeleteShardingNodeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetNodeId(v string) *DeleteShardingNodeRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetOwnerAccount(v string) *DeleteShardingNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetOwnerId(v int64) *DeleteShardingNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetResourceOwnerAccount(v string) *DeleteShardingNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetResourceOwnerId(v int64) *DeleteShardingNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetSecurityToken(v string) *DeleteShardingNodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteShardingNodeRequest) SetShardCount(v int32) *DeleteShardingNodeRequest {
	s.ShardCount = &v
	return s
}

func (s *DeleteShardingNodeRequest) Validate() error {
	return dara.Validate(s)
}
