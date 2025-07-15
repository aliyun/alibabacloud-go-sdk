// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpnGatewayRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteVpnGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVpnGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVpnGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpnGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpnGatewayRequest
	GetResourceOwnerId() *int64
	SetVpnGatewayId(v string) *DeleteVpnGatewayRequest
	GetVpnGatewayId() *string
}

type DeleteVpnGatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPN gateway.
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
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DeleteVpnGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpnGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpnGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpnGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVpnGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpnGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpnGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpnGatewayRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DeleteVpnGatewayRequest) SetClientToken(v string) *DeleteVpnGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpnGatewayRequest) SetOwnerAccount(v string) *DeleteVpnGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpnGatewayRequest) SetOwnerId(v int64) *DeleteVpnGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVpnGatewayRequest) SetRegionId(v string) *DeleteVpnGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpnGatewayRequest) SetResourceOwnerAccount(v string) *DeleteVpnGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpnGatewayRequest) SetResourceOwnerId(v int64) *DeleteVpnGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpnGatewayRequest) SetVpnGatewayId(v string) *DeleteVpnGatewayRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DeleteVpnGatewayRequest) Validate() error {
	return dara.Validate(s)
}
