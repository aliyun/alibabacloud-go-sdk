// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpAddressesFromEgressProxyIpPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveIpAddressesFromEgressProxyIpPoolRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveIpAddressesFromEgressProxyIpPoolRequest
	GetDryRun() *bool
	SetEgressProxyIpPoolId(v string) *RemoveIpAddressesFromEgressProxyIpPoolRequest
	GetEgressProxyIpPoolId() *string
	SetZoneMappings(v []*RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) *RemoveIpAddressesFromEgressProxyIpPoolRequest
	GetZoneMappings() []*RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings
}

type RemoveIpAddressesFromEgressProxyIpPoolRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eppool-gw80606a59ecc10****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// This parameter is required.
	ZoneMappings []*RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s RemoveIpAddressesFromEgressProxyIpPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpAddressesFromEgressProxyIpPoolRequest) GoString() string {
	return s.String()
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) GetZoneMappings() []*RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings {
	return s.ZoneMappings
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) SetClientToken(v string) *RemoveIpAddressesFromEgressProxyIpPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) SetDryRun(v bool) *RemoveIpAddressesFromEgressProxyIpPoolRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) SetEgressProxyIpPoolId(v string) *RemoveIpAddressesFromEgressProxyIpPoolRequest {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) SetZoneMappings(v []*RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) *RemoveIpAddressesFromEgressProxyIpPoolRequest {
	s.ZoneMappings = v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequest) Validate() error {
	if s.ZoneMappings != nil {
		for _, item := range s.ZoneMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings struct {
	// This parameter is required.
	IpAddresses []*RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) GetIpAddresses() []*RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	return s.IpAddresses
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) SetIpAddresses(v []*RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses) *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings {
	s.IpAddresses = v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) SetZoneId(v string) *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappings) Validate() error {
	if s.IpAddresses != nil {
		for _, item := range s.IpAddresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses struct {
	// example:
	//
	// eip-2zep1zbmxnvoyt1vs****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// example:
	//
	// 192.168.10.1
	PrivateIpv4Address *string `json:"PrivateIpv4Address,omitempty" xml:"PrivateIpv4Address,omitempty"`
}

func (s RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses) GoString() string {
	return s.String()
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses) GetPrivateIpv4Address() *string {
	return s.PrivateIpv4Address
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses) SetEipAllocationId(v string) *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	s.EipAllocationId = &v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses) SetPrivateIpv4Address(v string) *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	s.PrivateIpv4Address = &v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolRequestZoneMappingsIpAddresses) Validate() error {
	return dara.Validate(s)
}
