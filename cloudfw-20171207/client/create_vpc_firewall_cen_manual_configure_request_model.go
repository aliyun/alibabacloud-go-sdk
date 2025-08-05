// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallCenManualConfigureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateVpcFirewallCenManualConfigureRequest
	GetCenId() *string
	SetLang(v string) *CreateVpcFirewallCenManualConfigureRequest
	GetLang() *string
	SetMemberUid(v string) *CreateVpcFirewallCenManualConfigureRequest
	GetMemberUid() *string
	SetVSwitchId(v string) *CreateVpcFirewallCenManualConfigureRequest
	GetVSwitchId() *string
	SetVpcFirewallName(v string) *CreateVpcFirewallCenManualConfigureRequest
	GetVpcFirewallName() *string
	SetVpcId(v string) *CreateVpcFirewallCenManualConfigureRequest
	GetVpcId() *string
}

type CreateVpcFirewallCenManualConfigureRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cen-37nddhri7jf0d2****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-qzeaol304m***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateVpcFirewallCenManualConfigureRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallCenManualConfigureRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenManualConfigureRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateVpcFirewallCenManualConfigureRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateVpcFirewallCenManualConfigureRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *CreateVpcFirewallCenManualConfigureRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateVpcFirewallCenManualConfigureRequest) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *CreateVpcFirewallCenManualConfigureRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcFirewallCenManualConfigureRequest) SetCenId(v string) *CreateVpcFirewallCenManualConfigureRequest {
	s.CenId = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureRequest) SetLang(v string) *CreateVpcFirewallCenManualConfigureRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureRequest) SetMemberUid(v string) *CreateVpcFirewallCenManualConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureRequest) SetVSwitchId(v string) *CreateVpcFirewallCenManualConfigureRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureRequest) SetVpcFirewallName(v string) *CreateVpcFirewallCenManualConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureRequest) SetVpcId(v string) *CreateVpcFirewallCenManualConfigureRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureRequest) Validate() error {
	return dara.Validate(s)
}
