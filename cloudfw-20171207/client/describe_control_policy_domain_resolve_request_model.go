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
	// The domain name. This parameter is required and must be specified when you call this operation.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The type of the cloud firewall.
	//
	// example:
	//
	// internet
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The IP address version. Valid values: **4*	- (IPv4) and **6*	- (IPv6).
	//
	// > This parameter is unconditionally required and has no dependency on RegionNo. If this parameter is not specified, the error MissingParameter.IpVersion is returned (-200157). If the value is invalid, the error ErrorParameterIpVersion is returned (-200135).
	//
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language type.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID. This parameter is required. If this parameter is not specified, the error MissingParameter.RegionNo is returned (-200155, The required parameter \\"RegionNo\\" is not provided.).
	//
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
