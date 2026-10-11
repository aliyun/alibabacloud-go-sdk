// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyIpPoolAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetAddressType() *string
	SetCreateTime(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetDescription() *string
	SetEgressProxyGatewayId(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetEgressProxyGatewayId() *string
	SetEgressProxyIpPoolId(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetEgressProxyIpPoolId() *string
	SetEgressProxyIpPoolName(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetEgressProxyIpPoolName() *string
	SetHashFactor(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetHashFactor() *string
	SetHashHeaderName(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetHashHeaderName() *string
	SetRequestId(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetRequestId() *string
	SetScheduler(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetScheduler() *string
	SetStatus(v string) *GetEgressProxyIpPoolAttributeResponseBody
	GetStatus() *string
	SetTags(v []*GetEgressProxyIpPoolAttributeResponseBodyTags) *GetEgressProxyIpPoolAttributeResponseBody
	GetTags() []*GetEgressProxyIpPoolAttributeResponseBodyTags
	SetZoneMappings(v []*GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) *GetEgressProxyIpPoolAttributeResponseBody
	GetZoneMappings() []*GetEgressProxyIpPoolAttributeResponseBodyZoneMappings
}

type GetEgressProxyIpPoolAttributeResponseBody struct {
	// example:
	//
	// Intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// 2022-07-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// example:
	//
	// testEpg
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
	// 01A05AB8-D4FC-5B68-B645-5BA513A94F36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Random
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// example:
	//
	// Active
	Status       *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags         []*GetEgressProxyIpPoolAttributeResponseBodyTags         `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneMappings []*GetEgressProxyIpPoolAttributeResponseBodyZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetEgressProxyIpPoolAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyIpPoolAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetAddressType() *string {
	return s.AddressType
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetEgressProxyIpPoolName() *string {
	return s.EgressProxyIpPoolName
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetHashFactor() *string {
	return s.HashFactor
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetHashHeaderName() *string {
	return s.HashHeaderName
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetTags() []*GetEgressProxyIpPoolAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) GetZoneMappings() []*GetEgressProxyIpPoolAttributeResponseBodyZoneMappings {
	return s.ZoneMappings
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetAddressType(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetCreateTime(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetDescription(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetEgressProxyGatewayId(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetEgressProxyIpPoolId(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetEgressProxyIpPoolName(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.EgressProxyIpPoolName = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetHashFactor(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.HashFactor = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetHashHeaderName(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.HashHeaderName = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetRequestId(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetScheduler(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetStatus(v string) *GetEgressProxyIpPoolAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetTags(v []*GetEgressProxyIpPoolAttributeResponseBodyTags) *GetEgressProxyIpPoolAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) SetZoneMappings(v []*GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) *GetEgressProxyIpPoolAttributeResponseBody {
	s.ZoneMappings = v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
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

type GetEgressProxyIpPoolAttributeResponseBodyTags struct {
	// example:
	//
	// epg
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEgressProxyIpPoolAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyIpPoolAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyTags) SetKey(v string) *GetEgressProxyIpPoolAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyTags) SetValue(v string) *GetEgressProxyIpPoolAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyIpPoolAttributeResponseBodyZoneMappings struct {
	IpAddresses []*GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// vsw-uf64h3tjn2xsd94mb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-guangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) GoString() string {
	return s.String()
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) GetIpAddresses() []*GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses {
	return s.IpAddresses
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) SetIpAddresses(v []*GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings {
	s.IpAddresses = v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) SetVSwitchId(v string) *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) SetZoneId(v string) *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappings) Validate() error {
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

type GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses struct {
	// example:
	//
	// eip-2zep1zbmxnvoyt1vsj****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// example:
	//
	// 192.168.10.1
	PrivateIpv4Address *string `json:"PrivateIpv4Address,omitempty" xml:"PrivateIpv4Address,omitempty"`
	// example:
	//
	// 48.101.2.1
	PublicIpv4Address *string `json:"PublicIpv4Address,omitempty" xml:"PublicIpv4Address,omitempty"`
}

func (s GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) GoString() string {
	return s.String()
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) GetPrivateIpv4Address() *string {
	return s.PrivateIpv4Address
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) GetPublicIpv4Address() *string {
	return s.PublicIpv4Address
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) SetEipAllocationId(v string) *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses {
	s.EipAllocationId = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) SetPrivateIpv4Address(v string) *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses {
	s.PrivateIpv4Address = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) SetPublicIpv4Address(v string) *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses {
	s.PublicIpv4Address = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponseBodyZoneMappingsIpAddresses) Validate() error {
	return dara.Validate(s)
}
