// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallManualVSwitchListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVpcFirewallManualVSwitchListRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeVpcFirewallManualVSwitchListRequest
	GetMemberUid() *string
	SetOwnerId(v int64) *DescribeVpcFirewallManualVSwitchListRequest
	GetOwnerId() *int64
	SetRegionNo(v string) *DescribeVpcFirewallManualVSwitchListRequest
	GetRegionNo() *string
	SetVpcId(v string) *DescribeVpcFirewallManualVSwitchListRequest
	GetVpcId() *string
}

type DescribeVpcFirewallManualVSwitchListRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 18820897691****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// This parameter is required.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-m5ewlqkuf7orclr1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallManualVSwitchListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallManualVSwitchListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) SetLang(v string) *DescribeVpcFirewallManualVSwitchListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) SetMemberUid(v string) *DescribeVpcFirewallManualVSwitchListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) SetOwnerId(v int64) *DescribeVpcFirewallManualVSwitchListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) SetRegionNo(v string) *DescribeVpcFirewallManualVSwitchListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) SetVpcId(v string) *DescribeVpcFirewallManualVSwitchListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListRequest) Validate() error {
	return dara.Validate(s)
}
