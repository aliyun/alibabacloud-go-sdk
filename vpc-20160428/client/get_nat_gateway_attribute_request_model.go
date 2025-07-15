// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatGatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayId(v string) *GetNatGatewayAttributeRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *GetNatGatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetNatGatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetNatGatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetNatGatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetNatGatewayAttributeRequest
	GetResourceOwnerId() *int64
}

type GetNatGatewayAttributeRequest struct {
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1b0lic8uz4r6vf2****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the NAT gateway is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetNatGatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNatGatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetNatGatewayAttributeRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetNatGatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetNatGatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetNatGatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNatGatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetNatGatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetNatGatewayAttributeRequest) SetNatGatewayId(v string) *GetNatGatewayAttributeRequest {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatGatewayAttributeRequest) SetOwnerAccount(v string) *GetNatGatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetNatGatewayAttributeRequest) SetOwnerId(v int64) *GetNatGatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetNatGatewayAttributeRequest) SetRegionId(v string) *GetNatGatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetNatGatewayAttributeRequest) SetResourceOwnerAccount(v string) *GetNatGatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetNatGatewayAttributeRequest) SetResourceOwnerId(v int64) *GetNatGatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetNatGatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
