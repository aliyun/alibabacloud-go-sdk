// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallPreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateNatFirewallPreCheckRequest
	GetLang() *string
	SetNatGatewayId(v string) *CreateNatFirewallPreCheckRequest
	GetNatGatewayId() *string
	SetRegionNo(v string) *CreateNatFirewallPreCheckRequest
	GetRegionNo() *string
	SetVpcId(v string) *CreateNatFirewallPreCheckRequest
	GetVpcId() *string
}

type CreateNatFirewallPreCheckRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ngw-uf69hlxv5c817iqrk****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-2zeiljdml8pble168****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNatFirewallPreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallPreCheckRequest) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallPreCheckRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateNatFirewallPreCheckRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateNatFirewallPreCheckRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *CreateNatFirewallPreCheckRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNatFirewallPreCheckRequest) SetLang(v string) *CreateNatFirewallPreCheckRequest {
	s.Lang = &v
	return s
}

func (s *CreateNatFirewallPreCheckRequest) SetNatGatewayId(v string) *CreateNatFirewallPreCheckRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateNatFirewallPreCheckRequest) SetRegionNo(v string) *CreateNatFirewallPreCheckRequest {
	s.RegionNo = &v
	return s
}

func (s *CreateNatFirewallPreCheckRequest) SetVpcId(v string) *CreateNatFirewallPreCheckRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNatFirewallPreCheckRequest) Validate() error {
	return dara.Validate(s)
}
