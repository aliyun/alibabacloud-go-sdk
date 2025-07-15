// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpnConnectionRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteVpnConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVpnConnectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVpnConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpnConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpnConnectionRequest
	GetResourceOwnerId() *int64
	SetVpnConnectionId(v string) *DeleteVpnConnectionRequest
	GetVpnConnectionId() *string
}

type DeleteVpnConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPsec-VPN connection is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-bp1bbi27hojx80nck****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DeleteVpnConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpnConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpnConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpnConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVpnConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpnConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpnConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpnConnectionRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DeleteVpnConnectionRequest) SetClientToken(v string) *DeleteVpnConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpnConnectionRequest) SetOwnerAccount(v string) *DeleteVpnConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpnConnectionRequest) SetOwnerId(v int64) *DeleteVpnConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVpnConnectionRequest) SetRegionId(v string) *DeleteVpnConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpnConnectionRequest) SetResourceOwnerAccount(v string) *DeleteVpnConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpnConnectionRequest) SetResourceOwnerId(v int64) *DeleteVpnConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpnConnectionRequest) SetVpnConnectionId(v string) *DeleteVpnConnectionRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DeleteVpnConnectionRequest) Validate() error {
	return dara.Validate(s)
}
