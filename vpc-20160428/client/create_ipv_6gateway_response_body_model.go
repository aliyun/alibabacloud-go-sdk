// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv6GatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6GatewayId(v string) *CreateIpv6GatewayResponseBody
	GetIpv6GatewayId() *string
	SetRequestId(v string) *CreateIpv6GatewayResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateIpv6GatewayResponseBody
	GetResourceGroupId() *string
}

type CreateIpv6GatewayResponseBody struct {
	// The ID of the IPv6 gateway.
	//
	// example:
	//
	// ipv6gw-hp3y0l3ln89j8cdvf****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the IPv6 gateway belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateIpv6GatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv6GatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpv6GatewayResponseBody) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *CreateIpv6GatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpv6GatewayResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIpv6GatewayResponseBody) SetIpv6GatewayId(v string) *CreateIpv6GatewayResponseBody {
	s.Ipv6GatewayId = &v
	return s
}

func (s *CreateIpv6GatewayResponseBody) SetRequestId(v string) *CreateIpv6GatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpv6GatewayResponseBody) SetResourceGroupId(v string) *CreateIpv6GatewayResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpv6GatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
