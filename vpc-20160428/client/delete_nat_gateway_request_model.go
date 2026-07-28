// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteNatGatewayRequest
	GetForce() *bool
	SetNatGatewayId(v string) *DeleteNatGatewayRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *DeleteNatGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNatGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNatGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteNatGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNatGatewayRequest
	GetResourceOwnerId() *int64
}

type DeleteNatGatewayRequest struct {
	// Specifies whether to forcefully delete the NAT gateway. Valid values:
	//
	// - **true**: forcefully deletes the NAT gateway. If this parameter is set to **true**:
	//
	//     - If the NAT gateway has SNAT rules, the system force deletes the SNAT rules.
	//
	//     - If the NAT gateway has DNAT rules, the system force deletes the DNAT rules.
	//
	//     - If the NAT gateway has associated Elastic IP Addresses (EIPs), the system automatically disassociates the EIPs.
	//
	//     - If the NAT gateway has NAT service plans that are not deleted, the system force deletes the NAT service plans.
	//
	// - **false*	- (default): does not forcefully delete the NAT gateway. If this parameter is set to **false**:
	//
	//     - If the NAT gateway has NAT service plans that are not deleted, delete the NAT service plans first.
	//
	//     - If the NAT gateway has SNAT rules, delete the SNAT rules first.
	//
	//     - If the NAT gateway has DNAT rules, delete the DNAT rules first.
	//
	//     - If the NAT gateway has associated EIPs, disassociate the EIPs first.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The instance ID of the NAT gateway that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatGatewayRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteNatGatewayRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNatGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNatGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNatGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNatGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNatGatewayRequest) SetForce(v bool) *DeleteNatGatewayRequest {
	s.Force = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetNatGatewayId(v string) *DeleteNatGatewayRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetOwnerAccount(v string) *DeleteNatGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetOwnerId(v int64) *DeleteNatGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetRegionId(v string) *DeleteNatGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetResourceOwnerAccount(v string) *DeleteNatGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetResourceOwnerId(v int64) *DeleteNatGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNatGatewayRequest) Validate() error {
	return dara.Validate(s)
}
