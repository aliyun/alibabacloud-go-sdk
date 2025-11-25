// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeVpcFirewallZoneRequest
	GetCenId() *string
	SetEnvironment(v string) *DescribeVpcFirewallZoneRequest
	GetEnvironment() *string
	SetLang(v string) *DescribeVpcFirewallZoneRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeVpcFirewallZoneRequest
	GetMemberUid() *string
	SetRegionNo(v string) *DescribeVpcFirewallZoneRequest
	GetRegionNo() *string
	SetSourceIp(v string) *DescribeVpcFirewallZoneRequest
	GetSourceIp() *string
	SetTransitRouterId(v string) *DescribeVpcFirewallZoneRequest
	GetTransitRouterId() *string
}

type DescribeVpcFirewallZoneRequest struct {
	// example:
	//
	// cen-4xbjup276au29r****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// TransitRouter
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 135809047715****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 222.212.86.7XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// tr-m5etmb2q7e0mxcur****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeVpcFirewallZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallZoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallZoneRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeVpcFirewallZoneRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *DescribeVpcFirewallZoneRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallZoneRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallZoneRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallZoneRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeVpcFirewallZoneRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeVpcFirewallZoneRequest) SetCenId(v string) *DescribeVpcFirewallZoneRequest {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallZoneRequest) SetEnvironment(v string) *DescribeVpcFirewallZoneRequest {
	s.Environment = &v
	return s
}

func (s *DescribeVpcFirewallZoneRequest) SetLang(v string) *DescribeVpcFirewallZoneRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallZoneRequest) SetMemberUid(v string) *DescribeVpcFirewallZoneRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallZoneRequest) SetRegionNo(v string) *DescribeVpcFirewallZoneRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallZoneRequest) SetSourceIp(v string) *DescribeVpcFirewallZoneRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVpcFirewallZoneRequest) SetTransitRouterId(v string) *DescribeVpcFirewallZoneRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeVpcFirewallZoneRequest) Validate() error {
	return dara.Validate(s)
}
