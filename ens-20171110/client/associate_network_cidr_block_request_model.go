// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateNetworkCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpVersion(v string) *AssociateNetworkCidrBlockRequest
	GetIpVersion() *string
	SetIpv6Isp(v string) *AssociateNetworkCidrBlockRequest
	GetIpv6Isp() *string
	SetNetworkId(v string) *AssociateNetworkCidrBlockRequest
	GetNetworkId() *string
	SetSecondaryCidrBlock(v string) *AssociateNetworkCidrBlockRequest
	GetSecondaryCidrBlock() *string
}

type AssociateNetworkCidrBlockRequest struct {
	// The IP address version. Valid values:
	//
	// - IPV4: IPv4 address.
	//
	// - IPV6: IPv6 address.
	//
	// example:
	//
	// IPV6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ISP of the IPv6 CIDR block. Valid values:
	//
	// - cmcc
	//
	// - unicom
	//
	// - telecom
	//
	// example:
	//
	// telecom
	Ipv6Isp *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	// The network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-257gqcdfvx6n****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The secondary IPv4 CIDR block to be added. The CIDR block must meet the following requirements:
	//
	// - The CIDR block must be within the 10.0.0.0/8, 172.16.0.0/12, or 192.168.0.0/16 range.
	//
	// - The secondary CIDR block must not overlap with the primary CIDR block or existing secondary CIDR blocks of the network.
	//
	// example:
	//
	// 10.XX.XX.0/24
	SecondaryCidrBlock *string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty"`
}

func (s AssociateNetworkCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *AssociateNetworkCidrBlockRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *AssociateNetworkCidrBlockRequest) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *AssociateNetworkCidrBlockRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *AssociateNetworkCidrBlockRequest) GetSecondaryCidrBlock() *string {
	return s.SecondaryCidrBlock
}

func (s *AssociateNetworkCidrBlockRequest) SetIpVersion(v string) *AssociateNetworkCidrBlockRequest {
	s.IpVersion = &v
	return s
}

func (s *AssociateNetworkCidrBlockRequest) SetIpv6Isp(v string) *AssociateNetworkCidrBlockRequest {
	s.Ipv6Isp = &v
	return s
}

func (s *AssociateNetworkCidrBlockRequest) SetNetworkId(v string) *AssociateNetworkCidrBlockRequest {
	s.NetworkId = &v
	return s
}

func (s *AssociateNetworkCidrBlockRequest) SetSecondaryCidrBlock(v string) *AssociateNetworkCidrBlockRequest {
	s.SecondaryCidrBlock = &v
	return s
}

func (s *AssociateNetworkCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
