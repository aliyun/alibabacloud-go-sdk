// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneToVpcEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddZoneToVpcEndpointRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddZoneToVpcEndpointRequest
	GetDryRun() *bool
	SetEndpointId(v string) *AddZoneToVpcEndpointRequest
	GetEndpointId() *string
	SetIpv6Address(v string) *AddZoneToVpcEndpointRequest
	GetIpv6Address() *string
	SetRegionId(v string) *AddZoneToVpcEndpointRequest
	GetRegionId() *string
	SetVSwitchId(v string) *AddZoneToVpcEndpointRequest
	GetVSwitchId() *string
	SetZoneId(v string) *AddZoneToVpcEndpointRequest
	GetZoneId() *string
	SetIp(v string) *AddZoneToVpcEndpointRequest
	GetIp() *string
}

type AddZoneToVpcEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the endpoint to which you want to add the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The IPv6 address of the endpoint ENI in the zone that you want to add.
	//
	// example:
	//
	// 2408:4005:3b6:****:6955:c3cb:34c:****
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the vSwitch in the zone that you want to add. The system automatically creates an endpoint ENI in the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-hjkshjvdkdvd****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone that you want to add.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The IP address of the endpoint elastic network interface (ENI) in the zone that you want to add.
	//
	// example:
	//
	// 192.XX.XX.32
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s AddZoneToVpcEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s AddZoneToVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *AddZoneToVpcEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddZoneToVpcEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddZoneToVpcEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *AddZoneToVpcEndpointRequest) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *AddZoneToVpcEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddZoneToVpcEndpointRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AddZoneToVpcEndpointRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddZoneToVpcEndpointRequest) GetIp() *string {
	return s.Ip
}

func (s *AddZoneToVpcEndpointRequest) SetClientToken(v string) *AddZoneToVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetDryRun(v bool) *AddZoneToVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetEndpointId(v string) *AddZoneToVpcEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetIpv6Address(v string) *AddZoneToVpcEndpointRequest {
	s.Ipv6Address = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetRegionId(v string) *AddZoneToVpcEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetVSwitchId(v string) *AddZoneToVpcEndpointRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetZoneId(v string) *AddZoneToVpcEndpointRequest {
	s.ZoneId = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) SetIp(v string) *AddZoneToVpcEndpointRequest {
	s.Ip = &v
	return s
}

func (s *AddZoneToVpcEndpointRequest) Validate() error {
	return dara.Validate(s)
}
