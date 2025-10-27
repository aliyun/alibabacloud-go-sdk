// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateDnsEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceName(v string) *CreatePrivateDnsEndpointRequest
	GetAccessInstanceName() *string
	SetFirewallType(v []*string) *CreatePrivateDnsEndpointRequest
	GetFirewallType() []*string
	SetIpProtocol(v string) *CreatePrivateDnsEndpointRequest
	GetIpProtocol() *string
	SetMemberUid(v int64) *CreatePrivateDnsEndpointRequest
	GetMemberUid() *int64
	SetPort(v int32) *CreatePrivateDnsEndpointRequest
	GetPort() *int32
	SetPrimaryDns(v string) *CreatePrivateDnsEndpointRequest
	GetPrimaryDns() *string
	SetPrimaryVSwitchId(v string) *CreatePrivateDnsEndpointRequest
	GetPrimaryVSwitchId() *string
	SetPrimaryVSwitchIp(v string) *CreatePrivateDnsEndpointRequest
	GetPrimaryVSwitchIp() *string
	SetPrivateDnsType(v string) *CreatePrivateDnsEndpointRequest
	GetPrivateDnsType() *string
	SetRegionNo(v string) *CreatePrivateDnsEndpointRequest
	GetRegionNo() *string
	SetStandbyDns(v string) *CreatePrivateDnsEndpointRequest
	GetStandbyDns() *string
	SetStandbyVSwitchId(v string) *CreatePrivateDnsEndpointRequest
	GetStandbyVSwitchId() *string
	SetStandbyVSwitchIp(v string) *CreatePrivateDnsEndpointRequest
	GetStandbyVSwitchIp() *string
	SetVpcId(v string) *CreatePrivateDnsEndpointRequest
	GetVpcId() *string
}

type CreatePrivateDnsEndpointRequest struct {
	// This parameter is required.
	AccessInstanceName *string `json:"AccessInstanceName,omitempty" xml:"AccessInstanceName,omitempty"`
	// This parameter is required.
	FirewallType []*string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty" type:"Repeated"`
	// example:
	//
	// UDP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 1.1.1.1
	PrimaryDns *string `json:"PrimaryDns,omitempty" xml:"PrimaryDns,omitempty"`
	// example:
	//
	// vsw-uf6b0dkyryer8******
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// example:
	//
	// 10.1.1.1
	PrimaryVSwitchIp *string `json:"PrimaryVSwitchIp,omitempty" xml:"PrimaryVSwitchIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Custom
	PrivateDnsType *string `json:"PrivateDnsType,omitempty" xml:"PrivateDnsType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 1.1.1.2
	StandbyDns *string `json:"StandbyDns,omitempty" xml:"StandbyDns,omitempty"`
	// example:
	//
	// vsw-8vb6jk75wfcwn******
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// example:
	//
	// 10.2.2.2
	StandbyVSwitchIp *string `json:"StandbyVSwitchIp,omitempty" xml:"StandbyVSwitchIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6b5lyul0x******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreatePrivateDnsEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateDnsEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateDnsEndpointRequest) GetAccessInstanceName() *string {
	return s.AccessInstanceName
}

func (s *CreatePrivateDnsEndpointRequest) GetFirewallType() []*string {
	return s.FirewallType
}

func (s *CreatePrivateDnsEndpointRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreatePrivateDnsEndpointRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *CreatePrivateDnsEndpointRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreatePrivateDnsEndpointRequest) GetPrimaryDns() *string {
	return s.PrimaryDns
}

func (s *CreatePrivateDnsEndpointRequest) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *CreatePrivateDnsEndpointRequest) GetPrimaryVSwitchIp() *string {
	return s.PrimaryVSwitchIp
}

func (s *CreatePrivateDnsEndpointRequest) GetPrivateDnsType() *string {
	return s.PrivateDnsType
}

func (s *CreatePrivateDnsEndpointRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *CreatePrivateDnsEndpointRequest) GetStandbyDns() *string {
	return s.StandbyDns
}

func (s *CreatePrivateDnsEndpointRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *CreatePrivateDnsEndpointRequest) GetStandbyVSwitchIp() *string {
	return s.StandbyVSwitchIp
}

func (s *CreatePrivateDnsEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreatePrivateDnsEndpointRequest) SetAccessInstanceName(v string) *CreatePrivateDnsEndpointRequest {
	s.AccessInstanceName = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetFirewallType(v []*string) *CreatePrivateDnsEndpointRequest {
	s.FirewallType = v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetIpProtocol(v string) *CreatePrivateDnsEndpointRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetMemberUid(v int64) *CreatePrivateDnsEndpointRequest {
	s.MemberUid = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetPort(v int32) *CreatePrivateDnsEndpointRequest {
	s.Port = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetPrimaryDns(v string) *CreatePrivateDnsEndpointRequest {
	s.PrimaryDns = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetPrimaryVSwitchId(v string) *CreatePrivateDnsEndpointRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetPrimaryVSwitchIp(v string) *CreatePrivateDnsEndpointRequest {
	s.PrimaryVSwitchIp = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetPrivateDnsType(v string) *CreatePrivateDnsEndpointRequest {
	s.PrivateDnsType = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetRegionNo(v string) *CreatePrivateDnsEndpointRequest {
	s.RegionNo = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetStandbyDns(v string) *CreatePrivateDnsEndpointRequest {
	s.StandbyDns = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetStandbyVSwitchId(v string) *CreatePrivateDnsEndpointRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetStandbyVSwitchIp(v string) *CreatePrivateDnsEndpointRequest {
	s.StandbyVSwitchIp = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) SetVpcId(v string) *CreatePrivateDnsEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreatePrivateDnsEndpointRequest) Validate() error {
	return dara.Validate(s)
}
