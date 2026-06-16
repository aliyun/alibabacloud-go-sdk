// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateNetworkZoneRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateNetworkZoneRequest
	GetInstanceId() *string
	SetIpv4Cidrs(v []*string) *UpdateNetworkZoneRequest
	GetIpv4Cidrs() []*string
	SetIpv6Cidrs(v []*string) *UpdateNetworkZoneRequest
	GetIpv6Cidrs() []*string
	SetNetworkZoneId(v string) *UpdateNetworkZoneRequest
	GetNetworkZoneId() *string
	SetNetworkZoneName(v string) *UpdateNetworkZoneRequest
	GetNetworkZoneName() *string
	SetVpcId(v string) *UpdateNetworkZoneRequest
	GetVpcId() *string
}

type UpdateNetworkZoneRequest struct {
	// The idempotence token.
	//
	// example:
	//
	// client-token-examplexxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IPv4 CIDR blocks of the network zone.
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// The IPv6 CIDR blocks of the network zone.
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// The network zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// network_11111
	NetworkZoneId *string `json:"NetworkZoneId,omitempty" xml:"NetworkZoneId,omitempty"`
	// The name of the network zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// IPV4Test
	NetworkZoneName *string `json:"NetworkZoneName,omitempty" xml:"NetworkZoneName,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc_xxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateNetworkZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkZoneRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkZoneRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateNetworkZoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNetworkZoneRequest) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *UpdateNetworkZoneRequest) GetIpv6Cidrs() []*string {
	return s.Ipv6Cidrs
}

func (s *UpdateNetworkZoneRequest) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *UpdateNetworkZoneRequest) GetNetworkZoneName() *string {
	return s.NetworkZoneName
}

func (s *UpdateNetworkZoneRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateNetworkZoneRequest) SetClientToken(v string) *UpdateNetworkZoneRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateNetworkZoneRequest) SetInstanceId(v string) *UpdateNetworkZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNetworkZoneRequest) SetIpv4Cidrs(v []*string) *UpdateNetworkZoneRequest {
	s.Ipv4Cidrs = v
	return s
}

func (s *UpdateNetworkZoneRequest) SetIpv6Cidrs(v []*string) *UpdateNetworkZoneRequest {
	s.Ipv6Cidrs = v
	return s
}

func (s *UpdateNetworkZoneRequest) SetNetworkZoneId(v string) *UpdateNetworkZoneRequest {
	s.NetworkZoneId = &v
	return s
}

func (s *UpdateNetworkZoneRequest) SetNetworkZoneName(v string) *UpdateNetworkZoneRequest {
	s.NetworkZoneName = &v
	return s
}

func (s *UpdateNetworkZoneRequest) SetVpcId(v string) *UpdateNetworkZoneRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateNetworkZoneRequest) Validate() error {
	return dara.Validate(s)
}
