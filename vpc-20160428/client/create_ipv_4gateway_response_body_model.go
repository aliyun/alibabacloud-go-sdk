// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv4GatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv4GatewayId(v string) *CreateIpv4GatewayResponseBody
	GetIpv4GatewayId() *string
	SetRequestId(v string) *CreateIpv4GatewayResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateIpv4GatewayResponseBody
	GetResourceGroupId() *string
}

type CreateIpv4GatewayResponseBody struct {
	// The ID of the IPv4 gateway.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	Ipv4GatewayId *string `json:"Ipv4GatewayId,omitempty" xml:"Ipv4GatewayId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F282742B-1BBB-5F63-A3AF-E92EC575A1A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateIpv4GatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv4GatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpv4GatewayResponseBody) GetIpv4GatewayId() *string {
	return s.Ipv4GatewayId
}

func (s *CreateIpv4GatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpv4GatewayResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIpv4GatewayResponseBody) SetIpv4GatewayId(v string) *CreateIpv4GatewayResponseBody {
	s.Ipv4GatewayId = &v
	return s
}

func (s *CreateIpv4GatewayResponseBody) SetRequestId(v string) *CreateIpv4GatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpv4GatewayResponseBody) SetResourceGroupId(v string) *CreateIpv4GatewayResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpv4GatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
