// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewayAvailableZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeVpnGatewayAvailableZonesRequest
	GetAcceptLanguage() *string
	SetOwnerAccount(v string) *DescribeVpnGatewayAvailableZonesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnGatewayAvailableZonesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVpnGatewayAvailableZonesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnGatewayAvailableZonesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnGatewayAvailableZonesRequest
	GetResourceOwnerId() *int64
	SetSpec(v string) *DescribeVpnGatewayAvailableZonesRequest
	GetSpec() *string
}

type DescribeVpnGatewayAvailableZonesRequest struct {
	// The language in which the returned results are displayed. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The bandwidth specification.
	//
	// 	- If an IPsec-VPN connection can be associated with the VPN gateway, this parameter specifies the bandwidth specification of the VPN gateway.
	//
	// 	- In scenarios where an IPsec-VPN connection can be associated with a transit router. This parameter specifies the bandwidth specification supported by an IPsec-VPN connection.
	//
	// Different bandwidth specifications may affect returned zone information. Valid values:
	//
	// 	- **5M**
	//
	// 	- **10M**
	//
	// 	- **20M**
	//
	// 	- **50M**
	//
	// 	- **100M**
	//
	// 	- **200M**
	//
	// 	- **500M**
	//
	// 	- **1000M**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s DescribeVpnGatewayAvailableZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayAvailableZonesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeVpnGatewayAvailableZonesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnGatewayAvailableZonesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnGatewayAvailableZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnGatewayAvailableZonesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnGatewayAvailableZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnGatewayAvailableZonesRequest) GetSpec() *string {
	return s.Spec
}

func (s *DescribeVpnGatewayAvailableZonesRequest) SetAcceptLanguage(v string) *DescribeVpnGatewayAvailableZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesRequest) SetOwnerAccount(v string) *DescribeVpnGatewayAvailableZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesRequest) SetOwnerId(v int64) *DescribeVpnGatewayAvailableZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesRequest) SetRegionId(v string) *DescribeVpnGatewayAvailableZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesRequest) SetResourceOwnerAccount(v string) *DescribeVpnGatewayAvailableZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesRequest) SetResourceOwnerId(v int64) *DescribeVpnGatewayAvailableZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesRequest) SetSpec(v string) *DescribeVpnGatewayAvailableZonesRequest {
	s.Spec = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesRequest) Validate() error {
	return dara.Validate(s)
}
