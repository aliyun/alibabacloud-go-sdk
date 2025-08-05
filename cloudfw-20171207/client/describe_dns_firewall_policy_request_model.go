// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsFirewallPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *DescribeDnsFirewallPolicyRequest
	GetAclAction() *string
	SetAclUuid(v string) *DescribeDnsFirewallPolicyRequest
	GetAclUuid() *string
	SetCurrentPage(v string) *DescribeDnsFirewallPolicyRequest
	GetCurrentPage() *string
	SetDescription(v string) *DescribeDnsFirewallPolicyRequest
	GetDescription() *string
	SetDestination(v string) *DescribeDnsFirewallPolicyRequest
	GetDestination() *string
	SetIpVersion(v string) *DescribeDnsFirewallPolicyRequest
	GetIpVersion() *string
	SetLang(v string) *DescribeDnsFirewallPolicyRequest
	GetLang() *string
	SetPageSize(v string) *DescribeDnsFirewallPolicyRequest
	GetPageSize() *string
	SetRelease(v string) *DescribeDnsFirewallPolicyRequest
	GetRelease() *string
	SetSource(v string) *DescribeDnsFirewallPolicyRequest
	GetSource() *string
	SetSourceIp(v string) *DescribeDnsFirewallPolicyRequest
	GetSourceIp() *string
}

type DescribeDnsFirewallPolicyRequest struct {
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// example:
	//
	// b6c8f905-2eb6-442a-ba35-9416e****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10.2.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 140.205.118.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDnsFirewallPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsFirewallPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsFirewallPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeDnsFirewallPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeDnsFirewallPolicyRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeDnsFirewallPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeDnsFirewallPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *DescribeDnsFirewallPolicyRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeDnsFirewallPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsFirewallPolicyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDnsFirewallPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *DescribeDnsFirewallPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeDnsFirewallPolicyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDnsFirewallPolicyRequest) SetAclAction(v string) *DescribeDnsFirewallPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetAclUuid(v string) *DescribeDnsFirewallPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetCurrentPage(v string) *DescribeDnsFirewallPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetDescription(v string) *DescribeDnsFirewallPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetDestination(v string) *DescribeDnsFirewallPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetIpVersion(v string) *DescribeDnsFirewallPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetLang(v string) *DescribeDnsFirewallPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetPageSize(v string) *DescribeDnsFirewallPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetRelease(v string) *DescribeDnsFirewallPolicyRequest {
	s.Release = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetSource(v string) *DescribeDnsFirewallPolicyRequest {
	s.Source = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) SetSourceIp(v string) *DescribeDnsFirewallPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDnsFirewallPolicyRequest) Validate() error {
	return dara.Validate(s)
}
