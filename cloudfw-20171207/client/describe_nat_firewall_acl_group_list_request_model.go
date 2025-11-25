// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallAclGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNatFirewallAclGroupListRequest
	GetLang() *string
}

type DescribeNatFirewallAclGroupListRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeNatFirewallAclGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallAclGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallAclGroupListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNatFirewallAclGroupListRequest) SetLang(v string) *DescribeNatFirewallAclGroupListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallAclGroupListRequest) Validate() error {
	return dara.Validate(s)
}
