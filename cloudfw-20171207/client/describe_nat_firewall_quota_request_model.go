// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNatFirewallQuotaRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeNatFirewallQuotaRequest
	GetSourceIp() *string
}

type DescribeNatFirewallQuotaRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 113.132.26.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeNatFirewallQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallQuotaRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNatFirewallQuotaRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeNatFirewallQuotaRequest) SetLang(v string) *DescribeNatFirewallQuotaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallQuotaRequest) SetSourceIp(v string) *DescribeNatFirewallQuotaRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNatFirewallQuotaRequest) Validate() error {
	return dara.Validate(s)
}
