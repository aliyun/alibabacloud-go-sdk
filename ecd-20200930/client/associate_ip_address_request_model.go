// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipId(v string) *AssociateIpAddressRequest
	GetEipId() *string
	SetNatGatewayId(v string) *AssociateIpAddressRequest
	GetNatGatewayId() *string
	SetNetworkInterfaceId(v string) *AssociateIpAddressRequest
	GetNetworkInterfaceId() *string
	SetOfficeSiteId(v string) *AssociateIpAddressRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *AssociateIpAddressRequest
	GetRegionId() *string
}

type AssociateIpAddressRequest struct {
	// This parameter is required.
	EipId              *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	NatGatewayId       *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OfficeSiteId       *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateIpAddressRequest) GetEipId() *string {
	return s.EipId
}

func (s *AssociateIpAddressRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *AssociateIpAddressRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssociateIpAddressRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *AssociateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateIpAddressRequest) SetEipId(v string) *AssociateIpAddressRequest {
	s.EipId = &v
	return s
}

func (s *AssociateIpAddressRequest) SetNatGatewayId(v string) *AssociateIpAddressRequest {
	s.NatGatewayId = &v
	return s
}

func (s *AssociateIpAddressRequest) SetNetworkInterfaceId(v string) *AssociateIpAddressRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssociateIpAddressRequest) SetOfficeSiteId(v string) *AssociateIpAddressRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *AssociateIpAddressRequest) SetRegionId(v string) *AssociateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
