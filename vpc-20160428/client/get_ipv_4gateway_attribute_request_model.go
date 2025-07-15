// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpv4GatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpv4GatewayId(v string) *GetIpv4GatewayAttributeRequest
	GetIpv4GatewayId() *string
	SetOwnerAccount(v string) *GetIpv4GatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetIpv4GatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetIpv4GatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetIpv4GatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetIpv4GatewayAttributeRequest
	GetResourceOwnerId() *int64
}

type GetIpv4GatewayAttributeRequest struct {
	// The ID of the IPv4 gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	Ipv4GatewayId *string `json:"Ipv4GatewayId,omitempty" xml:"Ipv4GatewayId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv4 gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetIpv4GatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIpv4GatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetIpv4GatewayAttributeRequest) GetIpv4GatewayId() *string {
	return s.Ipv4GatewayId
}

func (s *GetIpv4GatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetIpv4GatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetIpv4GatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIpv4GatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetIpv4GatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetIpv4GatewayAttributeRequest) SetIpv4GatewayId(v string) *GetIpv4GatewayAttributeRequest {
	s.Ipv4GatewayId = &v
	return s
}

func (s *GetIpv4GatewayAttributeRequest) SetOwnerAccount(v string) *GetIpv4GatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetIpv4GatewayAttributeRequest) SetOwnerId(v int64) *GetIpv4GatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetIpv4GatewayAttributeRequest) SetRegionId(v string) *GetIpv4GatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetIpv4GatewayAttributeRequest) SetResourceOwnerAccount(v string) *GetIpv4GatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetIpv4GatewayAttributeRequest) SetResourceOwnerId(v int64) *GetIpv4GatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetIpv4GatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
