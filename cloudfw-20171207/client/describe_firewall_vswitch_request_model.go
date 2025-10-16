// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallId(v string) *DescribeFirewallVSwitchRequest
	GetFirewallId() *string
	SetLang(v string) *DescribeFirewallVSwitchRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeFirewallVSwitchRequest
	GetMemberUid() *string
	SetPageNo(v string) *DescribeFirewallVSwitchRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeFirewallVSwitchRequest
	GetPageSize() *string
	SetRegionNo(v string) *DescribeFirewallVSwitchRequest
	GetRegionNo() *string
	SetVpcId(v string) *DescribeFirewallVSwitchRequest
	GetVpcId() *string
	SetVswitchId(v string) *DescribeFirewallVSwitchRequest
	GetVswitchId() *string
}

type DescribeFirewallVSwitchRequest struct {
	// example:
	//
	// vfw-tr-5b202e7f0be64611****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 102910763545****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vpc-uf6b5lyul0x******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-bp1sqg9w******
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeFirewallVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVSwitchRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeFirewallVSwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFirewallVSwitchRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeFirewallVSwitchRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeFirewallVSwitchRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeFirewallVSwitchRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeFirewallVSwitchRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeFirewallVSwitchRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeFirewallVSwitchRequest) SetFirewallId(v string) *DescribeFirewallVSwitchRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeFirewallVSwitchRequest) SetLang(v string) *DescribeFirewallVSwitchRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFirewallVSwitchRequest) SetMemberUid(v string) *DescribeFirewallVSwitchRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeFirewallVSwitchRequest) SetPageNo(v string) *DescribeFirewallVSwitchRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeFirewallVSwitchRequest) SetPageSize(v string) *DescribeFirewallVSwitchRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallVSwitchRequest) SetRegionNo(v string) *DescribeFirewallVSwitchRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeFirewallVSwitchRequest) SetVpcId(v string) *DescribeFirewallVSwitchRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeFirewallVSwitchRequest) SetVswitchId(v string) *DescribeFirewallVSwitchRequest {
	s.VswitchId = &v
	return s
}

func (s *DescribeFirewallVSwitchRequest) Validate() error {
	return dara.Validate(s)
}
