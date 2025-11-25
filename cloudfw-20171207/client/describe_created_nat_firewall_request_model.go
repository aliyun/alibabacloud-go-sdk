// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreatedNatFirewallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCreatedNatFirewallRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeCreatedNatFirewallRequest
	GetSourceIp() *string
}

type DescribeCreatedNatFirewallRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 121.225.255.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCreatedNatFirewallRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreatedNatFirewallRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreatedNatFirewallRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCreatedNatFirewallRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCreatedNatFirewallRequest) SetLang(v string) *DescribeCreatedNatFirewallRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCreatedNatFirewallRequest) SetSourceIp(v string) *DescribeCreatedNatFirewallRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCreatedNatFirewallRequest) Validate() error {
	return dara.Validate(s)
}
