// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionType(v string) *ReleasePublicNetworkAddressRequest
	GetConnectionType() *string
	SetDBInstanceId(v string) *ReleasePublicNetworkAddressRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *ReleasePublicNetworkAddressRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *ReleasePublicNetworkAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleasePublicNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleasePublicNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleasePublicNetworkAddressRequest
	GetResourceOwnerId() *int64
}

type ReleasePublicNetworkAddressRequest struct {
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
	// The instance ID.
	//
	// >  If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2235****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the mongos, shard, or Configserver node in the sharded cluster instance.
	//
	// > 	- This parameter is valid only if you set the **DBInstanceId*	- parameter to the ID of a sharded cluster instance.
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to view the ID of the mongos, shard, or Configserver node.
	//
	// example:
	//
	// s-bp2235****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleasePublicNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *ReleasePublicNetworkAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ReleasePublicNetworkAddressRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ReleasePublicNetworkAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleasePublicNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleasePublicNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleasePublicNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleasePublicNetworkAddressRequest) SetConnectionType(v string) *ReleasePublicNetworkAddressRequest {
	s.ConnectionType = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetDBInstanceId(v string) *ReleasePublicNetworkAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetNodeId(v string) *ReleasePublicNetworkAddressRequest {
	s.NodeId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetOwnerAccount(v string) *ReleasePublicNetworkAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetOwnerId(v int64) *ReleasePublicNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetResourceOwnerAccount(v string) *ReleasePublicNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetResourceOwnerId(v int64) *ReleasePublicNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
