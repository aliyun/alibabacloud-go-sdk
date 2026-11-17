// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *DescribeVpcFirewallControlPolicyRequest
	GetAclAction() *string
	SetAclUuid(v string) *DescribeVpcFirewallControlPolicyRequest
	GetAclUuid() *string
	SetCurrentPage(v string) *DescribeVpcFirewallControlPolicyRequest
	GetCurrentPage() *string
	SetDescription(v string) *DescribeVpcFirewallControlPolicyRequest
	GetDescription() *string
	SetDestination(v string) *DescribeVpcFirewallControlPolicyRequest
	GetDestination() *string
	SetLang(v string) *DescribeVpcFirewallControlPolicyRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeVpcFirewallControlPolicyRequest
	GetMemberUid() *string
	SetPageSize(v string) *DescribeVpcFirewallControlPolicyRequest
	GetPageSize() *string
	SetProto(v string) *DescribeVpcFirewallControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *DescribeVpcFirewallControlPolicyRequest
	GetRelease() *string
	SetRepeatType(v string) *DescribeVpcFirewallControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *DescribeVpcFirewallControlPolicyRequest
	GetSource() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallControlPolicyRequest
	GetVpcFirewallId() *string
}

type DescribeVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic in the access control policy of the virtual private cloud (VPC) firewalls. Valid values:
	//
	// - **accept**: allows the traffic.
	//
	// - **drop**: denies the traffic.
	//
	// - **log**: monitors the traffic.
	//
	// > If you do not specify this parameter, all actions are queried.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique identifier of the access control policy of the virtual private cloud (VPC) firewalls.
	//
	// example:
	//
	// 4037fbf7-3e39-4634-92a4-d0155247****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The page number of the current page for a paged query.
	//
	// > You must specify both CurrentPage and PageSize or omit both. If you specify only one of them, the operation returns the ErrorParametersPageSizeOrNo (400) error.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy of the virtual private cloud (VPC) firewalls. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy of the virtual private cloud (VPC) firewalls. Fuzzy match is supported.
	//
	// > Destination can be a single CIDR block, a domain name, or an address book name.
	//
	// example:
	//
	// 192.0.XX.XX/24
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The language of the request and response.
	//
	// Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member account of the current Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of access control policies of the virtual private cloud (VPC) firewalls to return on each page for a paged query.
	//
	// Maximum value: 50.
	//
	// > You must specify both CurrentPage and PageSize or omit both. If you specify only one of them, the operation returns the ErrorParametersPageSizeOrNo (400) error.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol type of the traffic in the access control policy of the virtual private cloud (VPC) firewalls. Valid values:
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// - **ANY**: all protocol types.
	//
	// > If you do not specify this parameter, all protocol types are queried.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The enabling status of the access control policy. Valid values:
	//
	// - **true**: The access control policy is enabled.
	//
	// - **false**: The access control policy is disabled.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence type of the policy validity period for the access control policy. Valid values:
	//
	// - **Permanent*	- (default): always.
	//
	// - **None**: a specific single time.
	//
	// - **Daily**: every day.
	//
	// - **Weekly**: every week.
	//
	// - **Monthly**: every month.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy of the virtual private cloud (VPC) firewalls. Fuzzy match is supported.
	//
	// > Source can be a single CIDR block or an address book name.
	//
	// example:
	//
	// 192.0.XX.XX/24
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The instance ID of the virtual private cloud (VPC) firewalls. Valid values:
	//
	// - If the virtual private cloud (VPC) firewalls protect traffic between two VPCs connected by CEN, set this parameter to the instance ID of the CEN instance.
	//
	// - If the virtual private cloud (VPC) firewalls protect traffic between two VPCs connected by Express Connect, set this parameter to the instance ID of the virtual private cloud (VPC) firewalls.
	//
	// > You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeVpcFirewallControlPolicyRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetAclAction(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetAclUuid(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetCurrentPage(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetDescription(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetDestination(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetLang(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetMemberUid(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetPageSize(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetProto(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetRelease(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetRepeatType(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetSource(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
