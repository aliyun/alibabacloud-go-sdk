// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAclGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeVpcFirewallAclGroupListRequest
	GetCurrentPage() *string
	SetFirewallConfigureStatus(v string) *DescribeVpcFirewallAclGroupListRequest
	GetFirewallConfigureStatus() *string
	SetFirewallId(v string) *DescribeVpcFirewallAclGroupListRequest
	GetFirewallId() *string
	SetLang(v string) *DescribeVpcFirewallAclGroupListRequest
	GetLang() *string
	SetPageSize(v string) *DescribeVpcFirewallAclGroupListRequest
	GetPageSize() *string
}

type DescribeVpcFirewallAclGroupListRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The configuration status of the VPC boundary firewall. Valid values:
	//
	// - **configured**: The VPC boundary firewall is configured.
	//
	// - If you omit this parameter, the operation queries the access control policies for all VPC boundary firewalls.
	//
	// example:
	//
	// configured
	FirewallConfigureStatus *string `json:"FirewallConfigureStatus,omitempty" xml:"FirewallConfigureStatus,omitempty"`
	// The instance ID of the VPC boundary firewall.
	//
	// example:
	//
	// vfw-tr-5b202e7f0be64611****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. The maximum value is 50.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallAclGroupListRequest) GetFirewallConfigureStatus() *string {
	return s.FirewallConfigureStatus
}

func (s *DescribeVpcFirewallAclGroupListRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeVpcFirewallAclGroupListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallAclGroupListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetCurrentPage(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetFirewallConfigureStatus(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.FirewallConfigureStatus = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetFirewallId(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetLang(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetPageSize(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) Validate() error {
	return dara.Validate(s)
}
