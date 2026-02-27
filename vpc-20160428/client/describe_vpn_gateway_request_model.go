// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeReservationData(v bool) *DescribeVpnGatewayRequest
	GetIncludeReservationData() *bool
	SetOwnerAccount(v string) *DescribeVpnGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVpnGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnGatewayRequest
	GetResourceOwnerId() *int64
	SetVpnGatewayId(v string) *DescribeVpnGatewayRequest
	GetVpnGatewayId() *string
}

type DescribeVpnGatewayRequest struct {
	// Specifies whether to include the data about pending orders. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// true
	IncludeReservationData *bool   `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// vpn-bp1r3v1xqkl0w519g****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeVpnGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
}

func (s *DescribeVpnGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnGatewayRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnGatewayRequest) SetIncludeReservationData(v bool) *DescribeVpnGatewayRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeVpnGatewayRequest) SetOwnerAccount(v string) *DescribeVpnGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnGatewayRequest) SetOwnerId(v int64) *DescribeVpnGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnGatewayRequest) SetRegionId(v string) *DescribeVpnGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnGatewayRequest) SetResourceOwnerAccount(v string) *DescribeVpnGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnGatewayRequest) SetResourceOwnerId(v int64) *DescribeVpnGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnGatewayRequest) SetVpnGatewayId(v string) *DescribeVpnGatewayRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnGatewayRequest) Validate() error {
	return dara.Validate(s)
}
