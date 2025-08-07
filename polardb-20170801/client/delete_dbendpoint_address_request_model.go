// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBEndpointAddressRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *DeleteDBEndpointAddressRequest
	GetDBEndpointId() *string
	SetNetType(v string) *DeleteDBEndpointAddressRequest
	GetNetType() *string
	SetOwnerAccount(v string) *DeleteDBEndpointAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBEndpointAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDBEndpointAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBEndpointAddressRequest
	GetResourceOwnerId() *int64
}

type DeleteDBEndpointAddressRequest struct {
	// The ID of the cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters that belong to your account, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the endpoint.
	//
	// >  You can call the [DescribeDBClusterEndpoints](https://help.aliyun.com/document_detail/98205.html) operation to query the endpoints of a specified PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pe-***************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The network type of the endpoint. Set the value to **Public*	- (public network).
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBEndpointAddressRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBEndpointAddressRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DeleteDBEndpointAddressRequest) GetNetType() *string {
	return s.NetType
}

func (s *DeleteDBEndpointAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBEndpointAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBEndpointAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBEndpointAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBEndpointAddressRequest) SetDBClusterId(v string) *DeleteDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetDBEndpointId(v string) *DeleteDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetNetType(v string) *DeleteDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetOwnerAccount(v string) *DeleteDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetOwnerId(v int64) *DeleteDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *DeleteDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetResourceOwnerId(v int64) *DeleteDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
