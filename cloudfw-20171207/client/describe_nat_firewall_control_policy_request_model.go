// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *DescribeNatFirewallControlPolicyRequest
	GetAclAction() *string
	SetAclUuid(v string) *DescribeNatFirewallControlPolicyRequest
	GetAclUuid() *string
	SetCurrentPage(v string) *DescribeNatFirewallControlPolicyRequest
	GetCurrentPage() *string
	SetDescription(v string) *DescribeNatFirewallControlPolicyRequest
	GetDescription() *string
	SetDestination(v string) *DescribeNatFirewallControlPolicyRequest
	GetDestination() *string
	SetDirection(v string) *DescribeNatFirewallControlPolicyRequest
	GetDirection() *string
	SetLang(v string) *DescribeNatFirewallControlPolicyRequest
	GetLang() *string
	SetNatGatewayId(v string) *DescribeNatFirewallControlPolicyRequest
	GetNatGatewayId() *string
	SetPageSize(v string) *DescribeNatFirewallControlPolicyRequest
	GetPageSize() *string
	SetProto(v string) *DescribeNatFirewallControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *DescribeNatFirewallControlPolicyRequest
	GetRelease() *string
	SetRepeatType(v string) *DescribeNatFirewallControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *DescribeNatFirewallControlPolicyRequest
	GetSource() *string
}

type DescribeNatFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic.
	//
	// Valid values:
	//
	// - **accept**: Allow
	//
	// - **drop**: Deny
	//
	// - **log**: Monitor
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// example:
	//
	// 323f0697-2a21-4e43-b142-*****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The page number of the current page for a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy. Fuzzy queries are supported.
	//
	// > If you do not set this parameter, the descriptions of all policies are queried.
	//
	// example:
	//
	// test-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy queries are supported. The value of this parameter varies based on the value of the DestinationType parameter.
	//
	// - If DestinationType is set to `net`, the value of this parameter is a CIDR block. Example: 10.0.3.0/24.
	//
	// - If DestinationType is set to `domain`, the value of this parameter is a domain name. Example: aliyun.
	//
	// - If DestinationType is set to `group`, the value of this parameter is the name of an address book. Example: db_group.
	//
	// - If DestinationType is set to `location`, the value of this parameter is a region name. For more information, see [AddControlPolicy](https://help.aliyun.com/document_detail/474128.html). Example: ["BJ11", "ZB"].
	//
	// > If you do not set this parameter, all types of destination addresses are queried.
	//
	// example:
	//
	// x.x.x.x/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The traffic direction of the access control policy. Valid values:
	//
	// - **out**: outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-xxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The maximum number of entries to return on each page for a paged query. The default value is 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol type of the traffic in the access control policy. Valid values:
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// - **ANY*	- (all protocol types)
	//
	// > If you do not set this parameter, all protocol types are queried.
	//
	// example:
	//
	// ANY
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// - **true**: enabled
	//
	// - **false**: disabled
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence type for the policy validity period. Valid values:
	//
	// - **Permanent*	- (default): always
	//
	// - **None**: one-time
	//
	// - **Daily**: daily
	//
	// - **Weekly**: weekly
	//
	// - **Monthly**: monthly
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Fuzzy queries are supported. The value of this parameter varies based on the value of the SourceType parameter.
	//
	// - If SourceType is set to `net`, the value of this parameter is a CIDR block. Example: 192.0.XX.XX/24.
	//
	// - If SourceType is set to `group`, the value of this parameter is the name of an address book. Example: db_group. If you leave this parameter empty, all source addresses are queried.
	//
	// - If SourceType is set to `location`, the value of this parameter is a source region. Example: Beijing or beijing. You can use either Chinese or English to specify the region.
	//
	// > If you do not set this parameter, all types of source addresses are queried.
	//
	// example:
	//
	// 1.1.1.1/32
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeNatFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeNatFirewallControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeNatFirewallControlPolicyRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeNatFirewallControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeNatFirewallControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *DescribeNatFirewallControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeNatFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNatFirewallControlPolicyRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallControlPolicyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeNatFirewallControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *DescribeNatFirewallControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *DescribeNatFirewallControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *DescribeNatFirewallControlPolicyRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeNatFirewallControlPolicyRequest) SetAclAction(v string) *DescribeNatFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetAclUuid(v string) *DescribeNatFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetCurrentPage(v string) *DescribeNatFirewallControlPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetDescription(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetDestination(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetDirection(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetLang(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetNatGatewayId(v string) *DescribeNatFirewallControlPolicyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetPageSize(v string) *DescribeNatFirewallControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetProto(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetRelease(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetRepeatType(v string) *DescribeNatFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetSource(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
