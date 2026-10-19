// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateNetworkCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *AssociateNetworkCidrBlockResponseBody
	GetCidrBlock() *string
	SetIpVersion(v string) *AssociateNetworkCidrBlockResponseBody
	GetIpVersion() *string
	SetRequestId(v string) *AssociateNetworkCidrBlockResponseBody
	GetRequestId() *string
}

type AssociateNetworkCidrBlockResponseBody struct {
	// The allocated secondary CIDR block.
	//
	// example:
	//
	// 2408:XXXX:0:a600::/110
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The IP address type of the secondary CIDR block. Valid values:
	//
	// - IPV4: IPv4 address.
	//
	// - IPV6: IPv6 address.
	//
	// example:
	//
	// IPV6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateNetworkCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateNetworkCidrBlockResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *AssociateNetworkCidrBlockResponseBody) GetIpVersion() *string {
	return s.IpVersion
}

func (s *AssociateNetworkCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateNetworkCidrBlockResponseBody) SetCidrBlock(v string) *AssociateNetworkCidrBlockResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *AssociateNetworkCidrBlockResponseBody) SetIpVersion(v string) *AssociateNetworkCidrBlockResponseBody {
	s.IpVersion = &v
	return s
}

func (s *AssociateNetworkCidrBlockResponseBody) SetRequestId(v string) *AssociateNetworkCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateNetworkCidrBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
