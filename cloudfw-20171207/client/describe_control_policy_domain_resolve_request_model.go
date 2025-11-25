// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControlPolicyDomainResolveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeControlPolicyDomainResolveRequest
	GetDomain() *string
	SetFirewallType(v string) *DescribeControlPolicyDomainResolveRequest
	GetFirewallType() *string
	SetIpVersion(v int32) *DescribeControlPolicyDomainResolveRequest
	GetIpVersion() *int32
	SetLang(v string) *DescribeControlPolicyDomainResolveRequest
	GetLang() *string
	SetRegionNo(v string) *DescribeControlPolicyDomainResolveRequest
	GetRegionNo() *string
}

type DescribeControlPolicyDomainResolveRequest struct {
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// internet
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeControlPolicyDomainResolveRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeControlPolicyDomainResolveRequest) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyDomainResolveRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeControlPolicyDomainResolveRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribeControlPolicyDomainResolveRequest) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *DescribeControlPolicyDomainResolveRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeControlPolicyDomainResolveRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeControlPolicyDomainResolveRequest) SetDomain(v string) *DescribeControlPolicyDomainResolveRequest {
	s.Domain = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveRequest) SetFirewallType(v string) *DescribeControlPolicyDomainResolveRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveRequest) SetIpVersion(v int32) *DescribeControlPolicyDomainResolveRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveRequest) SetLang(v string) *DescribeControlPolicyDomainResolveRequest {
	s.Lang = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveRequest) SetRegionNo(v string) *DescribeControlPolicyDomainResolveRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveRequest) Validate() error {
	return dara.Validate(s)
}
