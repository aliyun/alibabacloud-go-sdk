// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShardingNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeShardingNetworkAddressRequest
	GetDBInstanceId() *string
	SetNetworkType(v string) *DescribeShardingNetworkAddressRequest
	GetNetworkType() *string
	SetNodeId(v string) *DescribeShardingNetworkAddressRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeShardingNetworkAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeShardingNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeShardingNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeShardingNetworkAddressRequest
	GetResourceOwnerId() *int64
}

type DescribeShardingNetworkAddressRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the mongos, shard, or Configserver node in the sharded cluster instance.
	//
	// >  You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to view the ID of the mongos, shard, or Configserver node.
	//
	// example:
	//
	// d-bpxxxxxxxx
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeShardingNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardingNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeShardingNetworkAddressRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeShardingNetworkAddressRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeShardingNetworkAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeShardingNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeShardingNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeShardingNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeShardingNetworkAddressRequest) SetDBInstanceId(v string) *DescribeShardingNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetNetworkType(v string) *DescribeShardingNetworkAddressRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetNodeId(v string) *DescribeShardingNetworkAddressRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetOwnerAccount(v string) *DescribeShardingNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetOwnerId(v int64) *DescribeShardingNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetResourceOwnerAccount(v string) *DescribeShardingNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) SetResourceOwnerId(v int64) *DescribeShardingNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeShardingNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
