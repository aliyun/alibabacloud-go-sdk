// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallSummaryInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVpcFirewallSummaryInfoRequest
	GetLang() *string
	SetUserType(v string) *DescribeVpcFirewallSummaryInfoRequest
	GetUserType() *string
}

type DescribeVpcFirewallSummaryInfoRequest struct {
	// The language type. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The user type. Valid values:
	//
	// **buy**: paid user.
	//
	// **free**: free user.
	//
	// example:
	//
	// buy
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeVpcFirewallSummaryInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallSummaryInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallSummaryInfoRequest) GetUserType() *string {
	return s.UserType
}

func (s *DescribeVpcFirewallSummaryInfoRequest) SetLang(v string) *DescribeVpcFirewallSummaryInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoRequest) SetUserType(v string) *DescribeVpcFirewallSummaryInfoRequest {
	s.UserType = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoRequest) Validate() error {
	return dara.Validate(s)
}
