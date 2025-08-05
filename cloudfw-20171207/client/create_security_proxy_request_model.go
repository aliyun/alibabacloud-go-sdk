// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallSwitch(v string) *CreateSecurityProxyRequest
	GetFirewallSwitch() *string
	SetLang(v string) *CreateSecurityProxyRequest
	GetLang() *string
	SetNatGatewayId(v string) *CreateSecurityProxyRequest
	GetNatGatewayId() *string
	SetNatRouteEntryList(v []*CreateSecurityProxyRequestNatRouteEntryList) *CreateSecurityProxyRequest
	GetNatRouteEntryList() []*CreateSecurityProxyRequestNatRouteEntryList
	SetProxyName(v string) *CreateSecurityProxyRequest
	GetProxyName() *string
	SetRegionNo(v string) *CreateSecurityProxyRequest
	GetRegionNo() *string
	SetStrictMode(v int32) *CreateSecurityProxyRequest
	GetStrictMode() *int32
	SetVpcId(v string) *CreateSecurityProxyRequest
	GetVpcId() *string
	SetVswitchAuto(v string) *CreateSecurityProxyRequest
	GetVswitchAuto() *string
	SetVswitchCidr(v string) *CreateSecurityProxyRequest
	GetVswitchCidr() *string
	SetVswitchId(v string) *CreateSecurityProxyRequest
	GetVswitchId() *string
}

type CreateSecurityProxyRequest struct {
	// The status of the NAT firewall. Valid values:
	//
	// 	- **open**: enabled
	//
	// 	- **close**: disabled
	//
	// example:
	//
	// close
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1okz6k7s4n4mnk5f1g3
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The routes to be switched to the NAT gateway.
	//
	// This parameter is required.
	NatRouteEntryList []*CreateSecurityProxyRequestNatRouteEntryList `json:"NatRouteEntryList,omitempty" xml:"NatRouteEntryList,omitempty" type:"Repeated"`
	// The name of the NAT firewall. The name must be 4 to 50 characters in length, and can contain letters, digits, and underscores (_). However, it cannot start with an underscore.
	//
	// This parameter is required.
	//
	// example:
	//
	// nat-idmp-fir
	ProxyName *string `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	// The region ID of the virtual private cloud (VPC).
	//
	// >  For more information about Cloud Firewall supported regions, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// Specifies whether to enable the strict mode. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// 0
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6b5lyul0xfgv74i01ph
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mode of the vSwitch that you want to use. Valid values:
	//
	// 	- **true**: automatic
	//
	// 	- **false**: manual
	//
	// example:
	//
	// true
	VswitchAuto *string `json:"VswitchAuto,omitempty" xml:"VswitchAuto,omitempty"`
	// The CIDR block of the vSwitch.
	//
	// example:
	//
	// 0.0.0.0/0
	VswitchCidr *string `json:"VswitchCidr,omitempty" xml:"VswitchCidr,omitempty"`
	// The ID of the vSwitch. This parameter is required if you set the VswitchAuto parameter to true.
	//
	// example:
	//
	// vsw-bp1sqg9wms9w9y1uxcs1x
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateSecurityProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityProxyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyRequest) GetFirewallSwitch() *string {
	return s.FirewallSwitch
}

func (s *CreateSecurityProxyRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSecurityProxyRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateSecurityProxyRequest) GetNatRouteEntryList() []*CreateSecurityProxyRequestNatRouteEntryList {
	return s.NatRouteEntryList
}

func (s *CreateSecurityProxyRequest) GetProxyName() *string {
	return s.ProxyName
}

func (s *CreateSecurityProxyRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *CreateSecurityProxyRequest) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *CreateSecurityProxyRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateSecurityProxyRequest) GetVswitchAuto() *string {
	return s.VswitchAuto
}

func (s *CreateSecurityProxyRequest) GetVswitchCidr() *string {
	return s.VswitchCidr
}

func (s *CreateSecurityProxyRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateSecurityProxyRequest) SetFirewallSwitch(v string) *CreateSecurityProxyRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetLang(v string) *CreateSecurityProxyRequest {
	s.Lang = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetNatGatewayId(v string) *CreateSecurityProxyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetNatRouteEntryList(v []*CreateSecurityProxyRequestNatRouteEntryList) *CreateSecurityProxyRequest {
	s.NatRouteEntryList = v
	return s
}

func (s *CreateSecurityProxyRequest) SetProxyName(v string) *CreateSecurityProxyRequest {
	s.ProxyName = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetRegionNo(v string) *CreateSecurityProxyRequest {
	s.RegionNo = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetStrictMode(v int32) *CreateSecurityProxyRequest {
	s.StrictMode = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVpcId(v string) *CreateSecurityProxyRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchAuto(v string) *CreateSecurityProxyRequest {
	s.VswitchAuto = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchCidr(v string) *CreateSecurityProxyRequest {
	s.VswitchCidr = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchId(v string) *CreateSecurityProxyRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateSecurityProxyRequest) Validate() error {
	return dara.Validate(s)
}

type CreateSecurityProxyRequestNatRouteEntryList struct {
	// The destination CIDR block of the default route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The next hop of the original NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1okz6k7s4n4mnk5f1g3
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The network type of the next hop. Set the value to NatGateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// NatGateway
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The route table to which the default route of the NAT gateway belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-2ze13wrgz7wsu9yiqeffg
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateSecurityProxyRequestNatRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityProxyRequestNatRouteEntryList) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) GetNextHopId() *string {
	return s.NextHopId
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) GetNextHopType() *string {
	return s.NextHopType
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetDestinationCidr(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetNextHopId(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.NextHopId = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetNextHopType(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.NextHopType = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetRouteTableId(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.RouteTableId = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) Validate() error {
	return dara.Validate(s)
}
