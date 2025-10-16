// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallPrecheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateVpcFirewallPrecheckRequest
	GetCenId() *string
	SetLang(v string) *CreateVpcFirewallPrecheckRequest
	GetLang() *string
	SetMemberUid(v string) *CreateVpcFirewallPrecheckRequest
	GetMemberUid() *string
	SetNetworkInstanceType(v string) *CreateVpcFirewallPrecheckRequest
	GetNetworkInstanceType() *string
	SetRegion(v string) *CreateVpcFirewallPrecheckRequest
	GetRegion() *string
	SetTransitRouterId(v string) *CreateVpcFirewallPrecheckRequest
	GetTransitRouterId() *string
	SetVpcId(v string) *CreateVpcFirewallPrecheckRequest
	GetVpcId() *string
}

type CreateVpcFirewallPrecheckRequest struct {
	// example:
	//
	// cen-iv8m2lj2fqg1xt****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 128599825273****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// cen_tr_firewall
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// example:
	//
	// cn-chengdu
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// tr-t4n16htuv1jalj9cq****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// example:
	//
	// vpc-bp132e2wpu9o6qth****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateVpcFirewallPrecheckRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallPrecheckRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallPrecheckRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateVpcFirewallPrecheckRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateVpcFirewallPrecheckRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *CreateVpcFirewallPrecheckRequest) GetNetworkInstanceType() *string {
	return s.NetworkInstanceType
}

func (s *CreateVpcFirewallPrecheckRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateVpcFirewallPrecheckRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateVpcFirewallPrecheckRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcFirewallPrecheckRequest) SetCenId(v string) *CreateVpcFirewallPrecheckRequest {
	s.CenId = &v
	return s
}

func (s *CreateVpcFirewallPrecheckRequest) SetLang(v string) *CreateVpcFirewallPrecheckRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallPrecheckRequest) SetMemberUid(v string) *CreateVpcFirewallPrecheckRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallPrecheckRequest) SetNetworkInstanceType(v string) *CreateVpcFirewallPrecheckRequest {
	s.NetworkInstanceType = &v
	return s
}

func (s *CreateVpcFirewallPrecheckRequest) SetRegion(v string) *CreateVpcFirewallPrecheckRequest {
	s.Region = &v
	return s
}

func (s *CreateVpcFirewallPrecheckRequest) SetTransitRouterId(v string) *CreateVpcFirewallPrecheckRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateVpcFirewallPrecheckRequest) SetVpcId(v string) *CreateVpcFirewallPrecheckRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVpcFirewallPrecheckRequest) Validate() error {
	return dara.Validate(s)
}
