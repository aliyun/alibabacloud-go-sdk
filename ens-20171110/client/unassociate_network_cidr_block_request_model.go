// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateNetworkCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6CidrBlock(v string) *UnassociateNetworkCidrBlockRequest
	GetIpv6CidrBlock() *string
	SetNetworkId(v string) *UnassociateNetworkCidrBlockRequest
	GetNetworkId() *string
	SetSecondaryCidrBlock(v string) *UnassociateNetworkCidrBlockRequest
	GetSecondaryCidrBlock() *string
}

type UnassociateNetworkCidrBlockRequest struct {
	// The secondary IPv6 CIDR block to delete.
	//
	// example:
	//
	// 2408:XXXX:0:a600::/110
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-257gqcdfvx6n****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The secondary IPv4 CIDR block to delete.
	//
	// example:
	//
	// 10.XX.XX.0/24
	SecondaryCidrBlock *string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty"`
}

func (s UnassociateNetworkCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkCidrBlockRequest) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *UnassociateNetworkCidrBlockRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *UnassociateNetworkCidrBlockRequest) GetSecondaryCidrBlock() *string {
	return s.SecondaryCidrBlock
}

func (s *UnassociateNetworkCidrBlockRequest) SetIpv6CidrBlock(v string) *UnassociateNetworkCidrBlockRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *UnassociateNetworkCidrBlockRequest) SetNetworkId(v string) *UnassociateNetworkCidrBlockRequest {
	s.NetworkId = &v
	return s
}

func (s *UnassociateNetworkCidrBlockRequest) SetSecondaryCidrBlock(v string) *UnassociateNetworkCidrBlockRequest {
	s.SecondaryCidrBlock = &v
	return s
}

func (s *UnassociateNetworkCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
