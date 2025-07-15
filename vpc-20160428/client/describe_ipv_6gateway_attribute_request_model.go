// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6GatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6GatewayId(v string) *DescribeIpv6GatewayAttributeRequest
	GetIpv6GatewayId() *string
	SetOwnerAccount(v string) *DescribeIpv6GatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIpv6GatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeIpv6GatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeIpv6GatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIpv6GatewayAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeIpv6GatewayAttributeRequest struct {
	// The ID of the IPv6 gateway that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6gw-hp3y0l3ln89j8cdvf****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPv6 gateway is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeIpv6GatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewayAttributeRequest) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *DescribeIpv6GatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIpv6GatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIpv6GatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpv6GatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIpv6GatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIpv6GatewayAttributeRequest) SetIpv6GatewayId(v string) *DescribeIpv6GatewayAttributeRequest {
	s.Ipv6GatewayId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeRequest) SetOwnerAccount(v string) *DescribeIpv6GatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeRequest) SetOwnerId(v int64) *DescribeIpv6GatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeRequest) SetRegionId(v string) *DescribeIpv6GatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeRequest) SetResourceOwnerAccount(v string) *DescribeIpv6GatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeRequest) SetResourceOwnerId(v int64) *DescribeIpv6GatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
