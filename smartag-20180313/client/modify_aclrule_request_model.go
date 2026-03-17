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
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-xhwhyuo43l0n*******
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the ACL rule.
	//
	// You can call the [DescribeACLAttribute](https://help.aliyun.com/document_detail/114017.html) operation to query the ID of the ACL rule that is added to the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// acr-u98qztgtgvhb********
	AcrId *string `json:"AcrId,omitempty" xml:"AcrId,omitempty"`
	// The description of the ACL rule.
	//
	// The description must be **1*	- to **512*	- characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// Specify the value of this parameter in CIDR notation. Example: 192.168.10.0/24.
	//
	// example:
	//
	// 0.0.0.0/0
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range. Valid values: **1*	- to **65535*	- and **-1**.
	//
	// Examples:
	//
	// 	- 1/200: port 1 to port 200.
	//
	// 	- 80/80: port 80.
	//
	// 	- \\-1/-1: all ports.
	//
	// example:
	//
	// 80/80
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The direction of traffic in which the ACL rule is applied. Valid values:
	//
	// 	- **in**: The ACL rule controls inbound network traffic of the on-premises network that is associated with the SAG instance.
	//
	// 	- **out**: The ACL rule controls outbound network traffic of the on-premises network that is associated with the SAG instance.
	//
	// example:
	//
	// in
	Direction       *string   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DpiGroupIds     []*string `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Repeated"`
	DpiSignatureIds []*string `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Repeated"`
	// The protocol used by the ACL rule.
	//
	// The supported protocols provided in this topic are for reference only. The actual protocols in the SAG console shall prevail. The value of the parameter is not case-sensitive.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the ACL rule.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// doctest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The action of the ACL rule. Valid values:
	//
	// 	- **accept**: allows network traffic.
	//
	// 	- **drop**: blocks network traffic.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The priority of the ACL rule.
	//
	// A smaller value indicates a higher priority. If multiple rules have the same priority, the rule that is applied earlier takes effect.
	//
	// Valid values: **1 to 100**. Default value: **1**.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the region where the ACL is deployed.
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
	// Specify the value of this parameter in CIDR notation. Example: 192.168.1.0/24.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range. Valid values: **1*	- to **65535*	- and **-1**.
	//
	// Examples:
	//
	// 	- 1/200: port 1 to port 200.
	//
	// 	- 80/80: port 80.
	//
	// 	- \\-1/-1: all ports.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The type of the ACL rule: Valid values:
	//
	// 	- **LAN**: The ACL rule controls traffic of private IP addresses.
	//
	// 	- **WAN**: The ACL rule controls traffic of public IP addresses.
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
