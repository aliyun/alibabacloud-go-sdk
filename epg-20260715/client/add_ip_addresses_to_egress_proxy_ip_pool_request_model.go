// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpAddressesToEgressProxyIpPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddIpAddressesToEgressProxyIpPoolRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddIpAddressesToEgressProxyIpPoolRequest
	GetDryRun() *bool
	SetEgressProxyIpPoolId(v string) *AddIpAddressesToEgressProxyIpPoolRequest
	GetEgressProxyIpPoolId() *string
	SetZoneMappings(v []*AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) *AddIpAddressesToEgressProxyIpPoolRequest
	GetZoneMappings() []*AddIpAddressesToEgressProxyIpPoolRequestZoneMappings
}

type AddIpAddressesToEgressProxyIpPoolRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42661417****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eppool-gw80606a59ecc10****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// This parameter is required.
	ZoneMappings []*AddIpAddressesToEgressProxyIpPoolRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s AddIpAddressesToEgressProxyIpPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIpAddressesToEgressProxyIpPoolRequest) GoString() string {
	return s.String()
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) GetZoneMappings() []*AddIpAddressesToEgressProxyIpPoolRequestZoneMappings {
	return s.ZoneMappings
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) SetClientToken(v string) *AddIpAddressesToEgressProxyIpPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) SetDryRun(v bool) *AddIpAddressesToEgressProxyIpPoolRequest {
	s.DryRun = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) SetEgressProxyIpPoolId(v string) *AddIpAddressesToEgressProxyIpPoolRequest {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) SetZoneMappings(v []*AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) *AddIpAddressesToEgressProxyIpPoolRequest {
	s.ZoneMappings = v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequest) Validate() error {
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

type AddIpAddressesToEgressProxyIpPoolRequestZoneMappings struct {
	IpAddresses []*AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	IpCount *int32 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) GetIpAddresses() []*AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	return s.IpAddresses
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) GetIpCount() *int32 {
	return s.IpCount
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) SetIpAddresses(v []*AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses) *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings {
	s.IpAddresses = v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) SetIpCount(v int32) *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings {
	s.IpCount = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) SetZoneId(v string) *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappings) Validate() error {
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

type AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses struct {
	// example:
	//
	// eip-2zep1zbmxnvoyt1vsj****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// example:
	//
	// 172.16.1.1
	PrivateIpv4Address *string `json:"PrivateIpv4Address,omitempty" xml:"PrivateIpv4Address,omitempty"`
}

func (s AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses) GoString() string {
	return s.String()
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses) GetPrivateIpv4Address() *string {
	return s.PrivateIpv4Address
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses) SetEipAllocationId(v string) *AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	s.EipAllocationId = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses) SetPrivateIpv4Address(v string) *AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	s.PrivateIpv4Address = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolRequestZoneMappingsIpAddresses) Validate() error {
	return dara.Validate(s)
}
