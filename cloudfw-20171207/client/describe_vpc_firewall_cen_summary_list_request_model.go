// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenSummaryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeVpcFirewallCenSummaryListRequest
	GetCurrentPage() *string
	SetLang(v string) *DescribeVpcFirewallCenSummaryListRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeVpcFirewallCenSummaryListRequest
	GetMemberUid() *string
	SetPageSize(v string) *DescribeVpcFirewallCenSummaryListRequest
	GetPageSize() *string
	SetTransitRouterType(v string) *DescribeVpcFirewallCenSummaryListRequest
	GetTransitRouterType() *string
}

type DescribeVpcFirewallCenSummaryListRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
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
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Basic
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
}

func (s DescribeVpcFirewallCenSummaryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenSummaryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenSummaryListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallCenSummaryListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallCenSummaryListRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallCenSummaryListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallCenSummaryListRequest) GetTransitRouterType() *string {
	return s.TransitRouterType
}

func (s *DescribeVpcFirewallCenSummaryListRequest) SetCurrentPage(v string) *DescribeVpcFirewallCenSummaryListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListRequest) SetLang(v string) *DescribeVpcFirewallCenSummaryListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListRequest) SetMemberUid(v string) *DescribeVpcFirewallCenSummaryListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListRequest) SetPageSize(v string) *DescribeVpcFirewallCenSummaryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListRequest) SetTransitRouterType(v string) *DescribeVpcFirewallCenSummaryListRequest {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListRequest) Validate() error {
	return dara.Validate(s)
}
