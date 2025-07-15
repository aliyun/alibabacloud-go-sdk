// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateVpcIpv6CidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6CidrBlock(v string) *AllocateVpcIpv6CidrResponseBody
	GetIpv6CidrBlock() *string
	SetRequestId(v string) *AllocateVpcIpv6CidrResponseBody
	GetRequestId() *string
}

type AllocateVpcIpv6CidrResponseBody struct {
	// The IPv6 CIDR block that is reserved.
	//
	// example:
	//
	// 2408:XXXX:0:a600::/56
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D20C13EA-2584-53BC-8393-69DE6D98EFF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateVpcIpv6CidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateVpcIpv6CidrResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateVpcIpv6CidrResponseBody) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *AllocateVpcIpv6CidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateVpcIpv6CidrResponseBody) SetIpv6CidrBlock(v string) *AllocateVpcIpv6CidrResponseBody {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *AllocateVpcIpv6CidrResponseBody) SetRequestId(v string) *AllocateVpcIpv6CidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateVpcIpv6CidrResponseBody) Validate() error {
	return dara.Validate(s)
}
