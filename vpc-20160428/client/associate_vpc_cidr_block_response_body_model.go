// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpcCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *AssociateVpcCidrBlockResponseBody
	GetCidrBlock() *string
	SetIpVersion(v string) *AssociateVpcCidrBlockResponseBody
	GetIpVersion() *string
	SetRequestId(v string) *AssociateVpcCidrBlockResponseBody
	GetRequestId() *string
}

type AssociateVpcCidrBlockResponseBody struct {
	// The IPv4 CIDR block to be added.
	//
	// example:
	//
	// 192.168.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The version of the IP address. Valid values:
	//
	// 	- **IPV4**: the IPv4 address.
	//
	// 	- **IPV6**: the IPv6 address.
	//
	// example:
	//
	// IPV4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C1221A1F-2ACD-4592-8F27-474E02883159
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateVpcCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpcCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateVpcCidrBlockResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *AssociateVpcCidrBlockResponseBody) GetIpVersion() *string {
	return s.IpVersion
}

func (s *AssociateVpcCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateVpcCidrBlockResponseBody) SetCidrBlock(v string) *AssociateVpcCidrBlockResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *AssociateVpcCidrBlockResponseBody) SetIpVersion(v string) *AssociateVpcCidrBlockResponseBody {
	s.IpVersion = &v
	return s
}

func (s *AssociateVpcCidrBlockResponseBody) SetRequestId(v string) *AssociateVpcCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateVpcCidrBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
