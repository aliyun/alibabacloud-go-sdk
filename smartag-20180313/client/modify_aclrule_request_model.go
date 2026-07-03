// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyACLRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *ModifyACLRuleRequest
	GetAclId() *string
	SetAcrId(v string) *ModifyACLRuleRequest
	GetAcrId() *string
	SetDescription(v string) *ModifyACLRuleRequest
	GetDescription() *string
	SetDestCidr(v string) *ModifyACLRuleRequest
	GetDestCidr() *string
	SetDestPortRange(v string) *ModifyACLRuleRequest
	GetDestPortRange() *string
	SetDirection(v string) *ModifyACLRuleRequest
	GetDirection() *string
	SetDpiGroupIds(v []*string) *ModifyACLRuleRequest
	GetDpiGroupIds() []*string
	SetDpiSignatureIds(v []*string) *ModifyACLRuleRequest
	GetDpiSignatureIds() []*string
	SetIpProtocol(v string) *ModifyACLRuleRequest
	GetIpProtocol() *string
	SetName(v string) *ModifyACLRuleRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyACLRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyACLRuleRequest
	GetOwnerId() *int64
	SetPolicy(v string) *ModifyACLRuleRequest
	GetPolicy() *string
	SetPriority(v int32) *ModifyACLRuleRequest
	GetPriority() *int32
	SetRegionId(v string) *ModifyACLRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyACLRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyACLRuleRequest
	GetResourceOwnerId() *int64
	SetSourceCidr(v string) *ModifyACLRuleRequest
	GetSourceCidr() *string
	SetSourcePortRange(v string) *ModifyACLRuleRequest
	GetSourcePortRange() *string
	SetType(v string) *ModifyACLRuleRequest
	GetType() *string
}

type ModifyACLRuleRequest struct {
	// The ID of the ACL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-xhwhyuo43l0n*******
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the access control rule.
	//
	// Call the [DescribeACLAttribute](https://help.aliyun.com/document_detail/114017.html) operation to query the IDs of access control rules in an ACL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// acr-u98qztgtgvhb********
	AcrId *string `json:"AcrId,omitempty" xml:"AcrId,omitempty"`
	// The description of the access control rule.
	//
	// The description must be **1*	- to **512*	- characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// The destination CIDR block must be in CIDR format. For example: 192.168.10.0/24.
	//
	// example:
	//
	// 0.0.0.0/0
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range. Valid values: **-1*	- or **1*	- to **65535**.
	//
	// Examples of the destination port range format:
	//
	// - 1/200: ports 1 to 200.
	//
	// - 80/80: port 80.
	//
	// - -1/-1: all ports.
	//
	// example:
	//
	// 80/80
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The direction in which the access control rule is applied. Valid values:
	//
	// - **in**: inbound. This is the direction of traffic from an external network to the on-premises network where the Smart Access Gateway instance is deployed.
	//
	// - **out**: outbound. This is the direction of traffic from the on-premises network where the Smart Access Gateway instance is deployed to an external network.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// A list of application group IDs that the access control rule matches.
	DpiGroupIds []*string `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Repeated"`
	// A list of application IDs that the access control rule matches.
	DpiSignatureIds []*string `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Repeated"`
	// The protocol used by the access control rule.
	//
	// For the protocols supported by the access control feature, see the information in the console. The protocol is not case-sensitive.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the access control rule.
	//
	// The name must be 2 to 128 characters in length, start with a letter or a Chinese character, and can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// doctest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The authorization policy of the access control rule. Valid values:
	//
	// - **accept**: allows access.
	//
	// - **drop**: denies access.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The priority of the access control rule.
	//
	// A smaller value indicates a higher priority. If rules have the same priority, the one that is first delivered to the Smart Access Gateway device takes precedence.
	//
	// Valid values: 1 to **100**. Default value: **1**.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the access control list (ACL) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source CIDR block.
	//
	// The source CIDR block must be in CIDR format. For example: 192.168.1.0/24.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range. Valid values: **-1*	- or **1*	- to **65535**.
	//
	// Examples of the source port range format:
	//
	// - 1/200: ports 1 to 200.
	//
	// - 80/80: port 80.
	//
	// - -1/-1: all ports.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The type of the access control rule. Valid values:
	//
	// - **LAN**: (Default) private network. This value indicates that the access control rule applies to traffic from private IP addresses.
	//
	// - **WAN**: public network. This value indicates that the access control rule applies to traffic from public IP addresses.
	//
	// example:
	//
	// LAN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyACLRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyACLRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyACLRuleRequest) GetAclId() *string {
	return s.AclId
}

func (s *ModifyACLRuleRequest) GetAcrId() *string {
	return s.AcrId
}

func (s *ModifyACLRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyACLRuleRequest) GetDestCidr() *string {
	return s.DestCidr
}

func (s *ModifyACLRuleRequest) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *ModifyACLRuleRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyACLRuleRequest) GetDpiGroupIds() []*string {
	return s.DpiGroupIds
}

func (s *ModifyACLRuleRequest) GetDpiSignatureIds() []*string {
	return s.DpiSignatureIds
}

func (s *ModifyACLRuleRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyACLRuleRequest) GetName() *string {
	return s.Name
}

func (s *ModifyACLRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyACLRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyACLRuleRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyACLRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyACLRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyACLRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyACLRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyACLRuleRequest) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *ModifyACLRuleRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifyACLRuleRequest) GetType() *string {
	return s.Type
}

func (s *ModifyACLRuleRequest) SetAclId(v string) *ModifyACLRuleRequest {
	s.AclId = &v
	return s
}

func (s *ModifyACLRuleRequest) SetAcrId(v string) *ModifyACLRuleRequest {
	s.AcrId = &v
	return s
}

func (s *ModifyACLRuleRequest) SetDescription(v string) *ModifyACLRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifyACLRuleRequest) SetDestCidr(v string) *ModifyACLRuleRequest {
	s.DestCidr = &v
	return s
}

func (s *ModifyACLRuleRequest) SetDestPortRange(v string) *ModifyACLRuleRequest {
	s.DestPortRange = &v
	return s
}

func (s *ModifyACLRuleRequest) SetDirection(v string) *ModifyACLRuleRequest {
	s.Direction = &v
	return s
}

func (s *ModifyACLRuleRequest) SetDpiGroupIds(v []*string) *ModifyACLRuleRequest {
	s.DpiGroupIds = v
	return s
}

func (s *ModifyACLRuleRequest) SetDpiSignatureIds(v []*string) *ModifyACLRuleRequest {
	s.DpiSignatureIds = v
	return s
}

func (s *ModifyACLRuleRequest) SetIpProtocol(v string) *ModifyACLRuleRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyACLRuleRequest) SetName(v string) *ModifyACLRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyACLRuleRequest) SetOwnerAccount(v string) *ModifyACLRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyACLRuleRequest) SetOwnerId(v int64) *ModifyACLRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyACLRuleRequest) SetPolicy(v string) *ModifyACLRuleRequest {
	s.Policy = &v
	return s
}

func (s *ModifyACLRuleRequest) SetPriority(v int32) *ModifyACLRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifyACLRuleRequest) SetRegionId(v string) *ModifyACLRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyACLRuleRequest) SetResourceOwnerAccount(v string) *ModifyACLRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyACLRuleRequest) SetResourceOwnerId(v int64) *ModifyACLRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyACLRuleRequest) SetSourceCidr(v string) *ModifyACLRuleRequest {
	s.SourceCidr = &v
	return s
}

func (s *ModifyACLRuleRequest) SetSourcePortRange(v string) *ModifyACLRuleRequest {
	s.SourcePortRange = &v
	return s
}

func (s *ModifyACLRuleRequest) SetType(v string) *ModifyACLRuleRequest {
	s.Type = &v
	return s
}

func (s *ModifyACLRuleRequest) Validate() error {
	return dara.Validate(s)
}
