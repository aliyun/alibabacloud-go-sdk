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
	// The action that Cloud Firewall performs on the traffic in the access control policy of the DNS border firewall. Valid values:
	//
	// - **accept**: Allow
	//
	// - **drop**: Deny
	//
	// - **log**: Monitor
	//
	// > If you do not specify this parameter, all actions are queried.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the firewall rule.
	//
	// example:
	//
	// b6c8f905-2eb6-442a-ba35-9416e****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The page number of the current page for a paging query. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the DNS border firewall policy.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy of the DNS border firewall. Fuzzy match is supported.
	//
	// > Destination can be a single CIDR block, a domain name, or an address book name.
	//
	// example:
	//
	// 10.2.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The IP address version. Valid values:
	//
	// - **4**: IPv4
	//
	// - **6**: IPv6
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return in the current query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the access control policy. By default, the policy is enabled after it is created. Valid values:
	//
	// - **true**: Enabled
	//
	// - **false**: Disabled
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy of the DNS border firewall. Fuzzy match is supported.
	//
	// > Source can be a single CIDR block or an address book name.
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source IP address of the requester.
	//
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
