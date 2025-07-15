// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpv6AddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6Address(v string) *AllocateIpv6AddressResponseBody
	GetIpv6Address() *string
	SetIpv6AddressId(v string) *AllocateIpv6AddressResponseBody
	GetIpv6AddressId() *string
	SetRequestId(v string) *AllocateIpv6AddressResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *AllocateIpv6AddressResponseBody
	GetResourceGroupId() *string
}

type AllocateIpv6AddressResponseBody struct {
	// The IPv6 address.
	//
	// example:
	//
	// 2408:XXXX:153:3921:851c:c435:7b12:1c5f
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The ID of the IPv6 address.
	//
	// example:
	//
	// ipv6-2zen5j4axcp5l5qyy****
	Ipv6AddressId *string `json:"Ipv6AddressId,omitempty" xml:"Ipv6AddressId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AllocateIpv6AddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpv6AddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateIpv6AddressResponseBody) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *AllocateIpv6AddressResponseBody) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *AllocateIpv6AddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateIpv6AddressResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateIpv6AddressResponseBody) SetIpv6Address(v string) *AllocateIpv6AddressResponseBody {
	s.Ipv6Address = &v
	return s
}

func (s *AllocateIpv6AddressResponseBody) SetIpv6AddressId(v string) *AllocateIpv6AddressResponseBody {
	s.Ipv6AddressId = &v
	return s
}

func (s *AllocateIpv6AddressResponseBody) SetRequestId(v string) *AllocateIpv6AddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateIpv6AddressResponseBody) SetResourceGroupId(v string) *AllocateIpv6AddressResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateIpv6AddressResponseBody) Validate() error {
	return dara.Validate(s)
}
