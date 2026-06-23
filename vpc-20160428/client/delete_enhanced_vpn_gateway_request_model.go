// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnhancedVpnGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEnhancedVpnGatewayRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteEnhancedVpnGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteEnhancedVpnGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteEnhancedVpnGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteEnhancedVpnGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteEnhancedVpnGatewayRequest
	GetResourceOwnerId() *int64
	SetVpnGatewayId(v string) *DeleteEnhancedVpnGatewayRequest
	GetVpnGatewayId() *string
}

type DeleteEnhancedVpnGatewayRequest struct {
	// The client token used to ensure the idempotence of the request.<br> You can generate this token by using your client, but you must ensure that it is unique for each request. The token can contain only ASCII characters and must not exceed 64 characters in length.<br>
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the enhanced VPN gateway is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the enhanced VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DeleteEnhancedVpnGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnhancedVpnGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnhancedVpnGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEnhancedVpnGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteEnhancedVpnGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteEnhancedVpnGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEnhancedVpnGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteEnhancedVpnGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteEnhancedVpnGatewayRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DeleteEnhancedVpnGatewayRequest) SetClientToken(v string) *DeleteEnhancedVpnGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayRequest) SetOwnerAccount(v string) *DeleteEnhancedVpnGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayRequest) SetOwnerId(v int64) *DeleteEnhancedVpnGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayRequest) SetRegionId(v string) *DeleteEnhancedVpnGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayRequest) SetResourceOwnerAccount(v string) *DeleteEnhancedVpnGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayRequest) SetResourceOwnerId(v int64) *DeleteEnhancedVpnGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayRequest) SetVpnGatewayId(v string) *DeleteEnhancedVpnGatewayRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DeleteEnhancedVpnGatewayRequest) Validate() error {
	return dara.Validate(s)
}
