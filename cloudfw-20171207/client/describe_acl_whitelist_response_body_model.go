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
	// Indicates whether the whitelist is enabled. Valid values:
	//
	// - **true**: enabled
	//
	// - **false**: disabled
	//
	// example:
	//
	// true
	DomainGroupUseDns *bool `json:"DomainGroupUseDns,omitempty" xml:"DomainGroupUseDns,omitempty"`
	// Indicates whether DNS is supported for domain names in NAT scenarios. Valid values:
	//
	// - **true**: supported
	//
	// - **false**: not supported
	//
	// example:
	//
	// false
	NatDomainGroupUseDns *bool `json:"NatDomainGroupUseDns,omitempty" xml:"NatDomainGroupUseDns,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7D45133B-DBC0-506B-9DF9-AB0735D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the message type is supported. Valid values:
	//
	// - **true**: supported
	//
	// - **false**: not supported
	//
	// example:
	//
	// true
	SupportMessageType *bool `json:"SupportMessageType,omitempty" xml:"SupportMessageType,omitempty"`
	// Indicates whether DNS is supported for domain names in VPC scenarios. Valid values:
	//
	// - **true**: supported
	//
	// - **false**: not supported
	//
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
