// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseNodePrivateNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionType(v string) *ReleaseNodePrivateNetworkAddressRequest
	GetConnectionType() *string
	SetDBInstanceId(v string) *ReleaseNodePrivateNetworkAddressRequest
	GetDBInstanceId() *string
	SetNetworkType(v string) *ReleaseNodePrivateNetworkAddressRequest
	GetNetworkType() *string
	SetNodeId(v string) *ReleaseNodePrivateNetworkAddressRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *ReleaseNodePrivateNetworkAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseNodePrivateNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleaseNodePrivateNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseNodePrivateNetworkAddressRequest
	GetResourceOwnerId() *int64
}

type ReleaseNodePrivateNetworkAddressRequest struct {
	// The public endpoint type. Valid values:
	//
	// 	- **SRV**
	//
	// 	- **Normal**
	//
	// >  This parameter is valid only when you want to release an SRV endpoint.
	//
	// example:
	//
	// SRV
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The ID of the sharded cluster instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1a7009eb24****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type of the internal endpoint. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC).
	//
	// 	- **Classic**: classic network.
	//
	// >  You can call the [DescribeShardingNetworkAddress](https://help.aliyun.com/document_detail/62135.html) operation to query the network type of the internal endpoint.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the shard or Configserver node.
	//
	// >  You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the ID of the shard or Configserver node.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp128a003436****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseNodePrivateNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseNodePrivateNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseNodePrivateNetworkAddressRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *ReleaseNodePrivateNetworkAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ReleaseNodePrivateNetworkAddressRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ReleaseNodePrivateNetworkAddressRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ReleaseNodePrivateNetworkAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseNodePrivateNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseNodePrivateNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseNodePrivateNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetConnectionType(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.ConnectionType = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetDBInstanceId(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetNetworkType(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.NetworkType = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetNodeId(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.NodeId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetOwnerAccount(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetOwnerId(v int64) *ReleaseNodePrivateNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetResourceOwnerAccount(v string) *ReleaseNodePrivateNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) SetResourceOwnerId(v int64) *ReleaseNodePrivateNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
