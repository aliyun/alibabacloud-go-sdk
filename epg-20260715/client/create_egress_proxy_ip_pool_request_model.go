// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyIpPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *CreateEgressProxyIpPoolRequest
	GetAddressType() *string
	SetClientToken(v string) *CreateEgressProxyIpPoolRequest
	GetClientToken() *string
	SetDescription(v string) *CreateEgressProxyIpPoolRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateEgressProxyIpPoolRequest
	GetDryRun() *bool
	SetEgressProxyGatewayId(v string) *CreateEgressProxyIpPoolRequest
	GetEgressProxyGatewayId() *string
	SetEgressProxyIpPoolName(v string) *CreateEgressProxyIpPoolRequest
	GetEgressProxyIpPoolName() *string
	SetHashFactor(v string) *CreateEgressProxyIpPoolRequest
	GetHashFactor() *string
	SetHashHeaderName(v string) *CreateEgressProxyIpPoolRequest
	GetHashHeaderName() *string
	SetScheduler(v string) *CreateEgressProxyIpPoolRequest
	GetScheduler() *string
	SetTag(v []*CreateEgressProxyIpPoolRequestTag) *CreateEgressProxyIpPoolRequest
	GetTag() []*CreateEgressProxyIpPoolRequestTag
	SetZoneMappings(v []*CreateEgressProxyIpPoolRequestZoneMappings) *CreateEgressProxyIpPoolRequest
	GetZoneMappings() []*CreateEgressProxyIpPoolRequestZoneMappings
}

type CreateEgressProxyIpPoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// created by epg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testIpPool
	EgressProxyIpPoolName *string `json:"EgressProxyIpPoolName,omitempty" xml:"EgressProxyIpPoolName,omitempty"`
	// example:
	//
	// SourceIP
	HashFactor *string `json:"HashFactor,omitempty" xml:"HashFactor,omitempty"`
	// example:
	//
	// X-Forwarded-For
	HashHeaderName *string `json:"HashHeaderName,omitempty" xml:"HashHeaderName,omitempty"`
	// example:
	//
	// Random
	Scheduler *string                              `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	Tag       []*CreateEgressProxyIpPoolRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	ZoneMappings []*CreateEgressProxyIpPoolRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyIpPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyIpPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyIpPoolRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *CreateEgressProxyIpPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEgressProxyIpPoolRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyIpPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEgressProxyIpPoolRequest) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *CreateEgressProxyIpPoolRequest) GetEgressProxyIpPoolName() *string {
	return s.EgressProxyIpPoolName
}

func (s *CreateEgressProxyIpPoolRequest) GetHashFactor() *string {
	return s.HashFactor
}

func (s *CreateEgressProxyIpPoolRequest) GetHashHeaderName() *string {
	return s.HashHeaderName
}

func (s *CreateEgressProxyIpPoolRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateEgressProxyIpPoolRequest) GetTag() []*CreateEgressProxyIpPoolRequestTag {
	return s.Tag
}

func (s *CreateEgressProxyIpPoolRequest) GetZoneMappings() []*CreateEgressProxyIpPoolRequestZoneMappings {
	return s.ZoneMappings
}

func (s *CreateEgressProxyIpPoolRequest) SetAddressType(v string) *CreateEgressProxyIpPoolRequest {
	s.AddressType = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetClientToken(v string) *CreateEgressProxyIpPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetDescription(v string) *CreateEgressProxyIpPoolRequest {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetDryRun(v bool) *CreateEgressProxyIpPoolRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetEgressProxyGatewayId(v string) *CreateEgressProxyIpPoolRequest {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetEgressProxyIpPoolName(v string) *CreateEgressProxyIpPoolRequest {
	s.EgressProxyIpPoolName = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetHashFactor(v string) *CreateEgressProxyIpPoolRequest {
	s.HashFactor = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetHashHeaderName(v string) *CreateEgressProxyIpPoolRequest {
	s.HashHeaderName = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetScheduler(v string) *CreateEgressProxyIpPoolRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetTag(v []*CreateEgressProxyIpPoolRequestTag) *CreateEgressProxyIpPoolRequest {
	s.Tag = v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) SetZoneMappings(v []*CreateEgressProxyIpPoolRequestZoneMappings) *CreateEgressProxyIpPoolRequest {
	s.ZoneMappings = v
	return s
}

func (s *CreateEgressProxyIpPoolRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateEgressProxyIpPoolRequestTag struct {
	// example:
	//
	// epg
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEgressProxyIpPoolRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyIpPoolRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyIpPoolRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEgressProxyIpPoolRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEgressProxyIpPoolRequestTag) SetKey(v string) *CreateEgressProxyIpPoolRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequestTag) SetValue(v string) *CreateEgressProxyIpPoolRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyIpPoolRequestZoneMappings struct {
	IpAddresses []*CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	IpCount *int32 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-0jl8x3ridfqbikax5o****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateEgressProxyIpPoolRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyIpPoolRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) GetIpAddresses() []*CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	return s.IpAddresses
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) GetIpCount() *int32 {
	return s.IpCount
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) SetIpAddresses(v []*CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses) *CreateEgressProxyIpPoolRequestZoneMappings {
	s.IpAddresses = v
	return s
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) SetIpCount(v int32) *CreateEgressProxyIpPoolRequestZoneMappings {
	s.IpCount = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) SetVSwitchId(v string) *CreateEgressProxyIpPoolRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) SetZoneId(v string) *CreateEgressProxyIpPoolRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequestZoneMappings) Validate() error {
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

type CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses struct {
	// example:
	//
	// eip-2zep1zbmxnvoyt1vs****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// example:
	//
	// 192.168.10.1
	PrivateIpv4Address *string `json:"PrivateIpv4Address,omitempty" xml:"PrivateIpv4Address,omitempty"`
}

func (s CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses) GetPrivateIpv4Address() *string {
	return s.PrivateIpv4Address
}

func (s *CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses) SetEipAllocationId(v string) *CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	s.EipAllocationId = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses) SetPrivateIpv4Address(v string) *CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses {
	s.PrivateIpv4Address = &v
	return s
}

func (s *CreateEgressProxyIpPoolRequestZoneMappingsIpAddresses) Validate() error {
	return dara.Validate(s)
}
