// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpv6InternetBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInternetBandwidthId(v string) *AllocateIpv6InternetBandwidthResponseBody
	GetInternetBandwidthId() *string
	SetIpv6AddressId(v string) *AllocateIpv6InternetBandwidthResponseBody
	GetIpv6AddressId() *string
	SetRequestId(v string) *AllocateIpv6InternetBandwidthResponseBody
	GetRequestId() *string
}

type AllocateIpv6InternetBandwidthResponseBody struct {
	// The ID of the Internet bandwidth that you purchased for the IPv6 gateway.
	//
	// example:
	//
	// ipv6bw-uf6hcyzu65v98v3du****
	InternetBandwidthId *string `json:"InternetBandwidthId,omitempty" xml:"InternetBandwidthId,omitempty"`
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
	// 6972A26E-99B1-4367-9890-FBDEBB0F5E7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateIpv6InternetBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpv6InternetBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateIpv6InternetBandwidthResponseBody) GetInternetBandwidthId() *string {
	return s.InternetBandwidthId
}

func (s *AllocateIpv6InternetBandwidthResponseBody) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *AllocateIpv6InternetBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateIpv6InternetBandwidthResponseBody) SetInternetBandwidthId(v string) *AllocateIpv6InternetBandwidthResponseBody {
	s.InternetBandwidthId = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthResponseBody) SetIpv6AddressId(v string) *AllocateIpv6InternetBandwidthResponseBody {
	s.Ipv6AddressId = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthResponseBody) SetRequestId(v string) *AllocateIpv6InternetBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateIpv6InternetBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
