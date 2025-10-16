// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainGroupUseDns(v bool) *DescribeAclWhitelistResponseBody
	GetDomainGroupUseDns() *bool
	SetNatDomainGroupUseDns(v bool) *DescribeAclWhitelistResponseBody
	GetNatDomainGroupUseDns() *bool
	SetRequestId(v string) *DescribeAclWhitelistResponseBody
	GetRequestId() *string
	SetSupportMessageType(v bool) *DescribeAclWhitelistResponseBody
	GetSupportMessageType() *bool
	SetVpcDomainGroupUseDns(v bool) *DescribeAclWhitelistResponseBody
	GetVpcDomainGroupUseDns() *bool
}

type DescribeAclWhitelistResponseBody struct {
	// example:
	//
	// true
	DomainGroupUseDns *bool `json:"DomainGroupUseDns,omitempty" xml:"DomainGroupUseDns,omitempty"`
	// example:
	//
	// false
	NatDomainGroupUseDns *bool `json:"NatDomainGroupUseDns,omitempty" xml:"NatDomainGroupUseDns,omitempty"`
	// example:
	//
	// 7D45133B-DBC0-506B-9DF9-AB0735D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SupportMessageType *bool `json:"SupportMessageType,omitempty" xml:"SupportMessageType,omitempty"`
	// example:
	//
	// false
	VpcDomainGroupUseDns *bool `json:"VpcDomainGroupUseDns,omitempty" xml:"VpcDomainGroupUseDns,omitempty"`
}

func (s DescribeAclWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclWhitelistResponseBody) GetDomainGroupUseDns() *bool {
	return s.DomainGroupUseDns
}

func (s *DescribeAclWhitelistResponseBody) GetNatDomainGroupUseDns() *bool {
	return s.NatDomainGroupUseDns
}

func (s *DescribeAclWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclWhitelistResponseBody) GetSupportMessageType() *bool {
	return s.SupportMessageType
}

func (s *DescribeAclWhitelistResponseBody) GetVpcDomainGroupUseDns() *bool {
	return s.VpcDomainGroupUseDns
}

func (s *DescribeAclWhitelistResponseBody) SetDomainGroupUseDns(v bool) *DescribeAclWhitelistResponseBody {
	s.DomainGroupUseDns = &v
	return s
}

func (s *DescribeAclWhitelistResponseBody) SetNatDomainGroupUseDns(v bool) *DescribeAclWhitelistResponseBody {
	s.NatDomainGroupUseDns = &v
	return s
}

func (s *DescribeAclWhitelistResponseBody) SetRequestId(v string) *DescribeAclWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclWhitelistResponseBody) SetSupportMessageType(v bool) *DescribeAclWhitelistResponseBody {
	s.SupportMessageType = &v
	return s
}

func (s *DescribeAclWhitelistResponseBody) SetVpcDomainGroupUseDns(v bool) *DescribeAclWhitelistResponseBody {
	s.VpcDomainGroupUseDns = &v
	return s
}

func (s *DescribeAclWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
