// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *AllocatePublicNetworkAddressRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *AllocatePublicNetworkAddressRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *AllocatePublicNetworkAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocatePublicNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AllocatePublicNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocatePublicNetworkAddressRequest
	GetResourceOwnerId() *int64
}

type AllocatePublicNetworkAddressRequest struct {
	// The ID of the instance
	//
	// > If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp11483712c1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the mongos, shard, or Configserver node in the sharded cluster instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to view the ID of the mongos, shard, or Configserver node.
	//
	// > This parameter is required only when you specify the **DBInstanceId*	- parameter to the ID of a sharded cluster instance.
	//
	// example:
	//
	// s-bp18e6d84ae3****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocatePublicNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocatePublicNetworkAddressRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AllocatePublicNetworkAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocatePublicNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocatePublicNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocatePublicNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocatePublicNetworkAddressRequest) SetDBInstanceId(v string) *AllocatePublicNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetNodeId(v string) *AllocatePublicNetworkAddressRequest {
	s.NodeId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetOwnerAccount(v string) *AllocatePublicNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetOwnerId(v int64) *AllocatePublicNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetResourceOwnerAccount(v string) *AllocatePublicNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetResourceOwnerId(v int64) *AllocatePublicNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
