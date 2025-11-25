// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAssetRegionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberUid(v string) *DescribeVpcFirewallAssetRegionListRequest
	GetMemberUid() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallAssetRegionListRequest
	GetVpcFirewallId() *string
}

type DescribeVpcFirewallAssetRegionListRequest struct {
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallAssetRegionListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAssetRegionListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAssetRegionListRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallAssetRegionListRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallAssetRegionListRequest) SetMemberUid(v string) *DescribeVpcFirewallAssetRegionListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallAssetRegionListRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallAssetRegionListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallAssetRegionListRequest) Validate() error {
	return dara.Validate(s)
}
