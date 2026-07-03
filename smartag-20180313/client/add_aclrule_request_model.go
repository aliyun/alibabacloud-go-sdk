// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddACLRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *AddACLRuleRequest
	GetAclId() *string
	SetDescription(v string) *AddACLRuleRequest
	GetDescription() *string
	SetDestCidr(v string) *AddACLRuleRequest
	GetDestCidr() *string
	SetDestPortRange(v string) *AddACLRuleRequest
	GetDestPortRange() *string
	SetDirection(v string) *AddACLRuleRequest
	GetDirection() *string
	SetDpiGroupIds(v []*string) *AddACLRuleRequest
	GetDpiGroupIds() []*string
	SetDpiSignatureIds(v []*string) *AddACLRuleRequest
	GetDpiSignatureIds() []*string
	SetIpProtocol(v string) *AddACLRuleRequest
	GetIpProtocol() *string
	SetName(v string) *AddACLRuleRequest
	GetName() *string
	SetOwnerAccount(v string) *AddACLRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddACLRuleRequest
	GetOwnerId() *int64
	SetPolicy(v string) *AddACLRuleRequest
	GetPolicy() *string
	SetPriority(v int32) *AddACLRuleRequest
	GetPriority() *int32
	SetRegionId(v string) *AddACLRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddACLRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddACLRuleRequest
	GetResourceOwnerId() *int64
	SetSourceCidr(v string) *AddACLRuleRequest
	GetSourceCidr() *string
	SetSourcePortRange(v string) *AddACLRuleRequest
	GetSourcePortRange() *string
	SetType(v string) *AddACLRuleRequest
	GetType() *string
}

type AddACLRuleRequest struct {
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-xhwhyuo43l0n*****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The description of the ACL rule.
	//
	// The description must be 1 to **512*	- characters in length.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// For example: 192.168.10.0/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.0/24
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range.
	//
	// Valid values: **-1*	- and **1*	- to **65535**.
	//
	// Use the format 1/200 or 80/80. A value of -1/-1 means all ports.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1/200
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The direction of traffic to which the ACL rule applies. Valid values:
	//
	// - **in**: inbound. Traffic from an external network to the local branch where the SAG instance is deployed.
	//
	// - **out**: outbound. Traffic from the local branch where the SAG instance is deployed to an external network.
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// A list of application group IDs. The ACL rule matches traffic of the specified application groups.
	//
	// For more information, see [ListDpiGroups](https://help.aliyun.com/document_detail/196754.html). You can specify up to **10*	- application group IDs.
	//
	// example:
	//
	// 20
	DpiGroupIds []*string `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Repeated"`
	// A list of application IDs. The ACL rule matches traffic of the specified applications.
	//
	// For more information, see [ListDpiSignatures](https://help.aliyun.com/document_detail/196630.html). You can specify up to **10*	- application IDs.
	//
	// example:
	//
	// 1
	DpiSignatureIds []*string `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Repeated"`
	// The protocol to which the ACL rule applies.
	//
	// For a list of supported protocols, see the console. The protocol is not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the ACL rule.
	//
	// The name must be 2 to 100 characters in length, start with a letter, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// doctest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The authorization policy of the ACL rule. Valid values:
	//
	// - **accept**: allows access.
	//
	// - **drop**: denies access.
	//
	// This parameter is required.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The priority of the ACL rule.
	//
	// A smaller value indicates a higher priority. If multiple rules have the same priority, the rule that is first delivered to the Smart Access Gateway device takes precedence.
	//
	// Valid values: 1 to **100**. Default value: **1**.
	//
	// example:
	//
	// 12
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the region where the access control list (ACL) is located.
	//
	// For more information, see [DescribeRegions](https://help.aliyun.com/document_detail/69813.html).
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
	// For example: 192.168.1.0/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.20.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range.
	//
	// Valid values: **-1*	- and **1*	- to **65535**.
	//
	// Use the format 1/200 or 80/80. A value of -1/-1 means all ports.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1/200
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The type of the ACL rule. Valid values:
	//
	// - **LAN**: (Default) private network. The ACL rule controls traffic on private networks.
	//
	// - **WAN**: public network. The ACL rule controls traffic on public networks.
	//
	// example:
	//
	// LAN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddACLRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddACLRuleRequest) GoString() string {
	return s.String()
}

func (s *AddACLRuleRequest) GetAclId() *string {
	return s.AclId
}

func (s *AddACLRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *AddACLRuleRequest) GetDestCidr() *string {
	return s.DestCidr
}

func (s *AddACLRuleRequest) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *AddACLRuleRequest) GetDirection() *string {
	return s.Direction
}

func (s *AddACLRuleRequest) GetDpiGroupIds() []*string {
	return s.DpiGroupIds
}

func (s *AddACLRuleRequest) GetDpiSignatureIds() []*string {
	return s.DpiSignatureIds
}

func (s *AddACLRuleRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AddACLRuleRequest) GetName() *string {
	return s.Name
}

func (s *AddACLRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddACLRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddACLRuleRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AddACLRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *AddACLRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddACLRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddACLRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddACLRuleRequest) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *AddACLRuleRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AddACLRuleRequest) GetType() *string {
	return s.Type
}

func (s *AddACLRuleRequest) SetAclId(v string) *AddACLRuleRequest {
	s.AclId = &v
	return s
}

func (s *AddACLRuleRequest) SetDescription(v string) *AddACLRuleRequest {
	s.Description = &v
	return s
}

func (s *AddACLRuleRequest) SetDestCidr(v string) *AddACLRuleRequest {
	s.DestCidr = &v
	return s
}

func (s *AddACLRuleRequest) SetDestPortRange(v string) *AddACLRuleRequest {
	s.DestPortRange = &v
	return s
}

func (s *AddACLRuleRequest) SetDirection(v string) *AddACLRuleRequest {
	s.Direction = &v
	return s
}

func (s *AddACLRuleRequest) SetDpiGroupIds(v []*string) *AddACLRuleRequest {
	s.DpiGroupIds = v
	return s
}

func (s *AddACLRuleRequest) SetDpiSignatureIds(v []*string) *AddACLRuleRequest {
	s.DpiSignatureIds = v
	return s
}

func (s *AddACLRuleRequest) SetIpProtocol(v string) *AddACLRuleRequest {
	s.IpProtocol = &v
	return s
}

func (s *AddACLRuleRequest) SetName(v string) *AddACLRuleRequest {
	s.Name = &v
	return s
}

func (s *AddACLRuleRequest) SetOwnerAccount(v string) *AddACLRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddACLRuleRequest) SetOwnerId(v int64) *AddACLRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *AddACLRuleRequest) SetPolicy(v string) *AddACLRuleRequest {
	s.Policy = &v
	return s
}

func (s *AddACLRuleRequest) SetPriority(v int32) *AddACLRuleRequest {
	s.Priority = &v
	return s
}

func (s *AddACLRuleRequest) SetRegionId(v string) *AddACLRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AddACLRuleRequest) SetResourceOwnerAccount(v string) *AddACLRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddACLRuleRequest) SetResourceOwnerId(v int64) *AddACLRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddACLRuleRequest) SetSourceCidr(v string) *AddACLRuleRequest {
	s.SourceCidr = &v
	return s
}

func (s *AddACLRuleRequest) SetSourcePortRange(v string) *AddACLRuleRequest {
	s.SourcePortRange = &v
	return s
}

func (s *AddACLRuleRequest) SetType(v string) *AddACLRuleRequest {
	s.Type = &v
	return s
}

func (s *AddACLRuleRequest) Validate() error {
	return dara.Validate(s)
}
