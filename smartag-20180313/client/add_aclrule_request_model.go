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
	// The description must be **1 to 512*	- characters in length.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The range of the destination IP addresses.
	//
	// Specify the value of this parameter in CIDR notation. Example: 192.168.10.0/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.0/24
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range.
	//
	// Valid values: **1*	- to **65535*	- and **-1**.
	//
	// Set the destination port range in one of the following formats: 1/200 or 80/80. A value of -1/-1 indicates all ports.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1/200
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The direction of traffic in which the ACL rule is applied. Valid values:
	//
	// 	- **in**: The ACL rule controls inbound network traffic of the on-premises network that is associated with the Smart Access Gateway (SAG) instance.
	//
	// 	- **out**: The ACL rule controls outbound network traffic of the on-premises network that is associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 20
	DpiGroupIds []*string `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DpiSignatureIds []*string `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Repeated"`
	// The protocol used by the ACL rule.
	//
	// The protocols that are provided in this topic are for reference only. The protocols available in the SAG console may vary. The value of the parameter is not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the ACL rule.
	//
	// The name must be 2 to 100 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// doctest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The action policy of the ACL rule. Valid values:
	//
	// 	- **accept**: allows network traffic.
	//
	// 	- **drop**: blocks the network traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The priority of the ACL rule.
	//
	// A smaller value indicates a higher priority. If rules have the same priority, whichever applied to the SAG devices earlier takes effect.
	//
	// Valid values: **1 to 100**. Default value: **1**.
	//
	// example:
	//
	// 12
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the region where the ACL is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The range of the source IP addresses.
	//
	// Specify the value of this parameter in CIDR notation. Example: 192.168.1.0/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.20.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range.
	//
	// Valid values: **1*	- to **65535*	- and **-1**.
	//
	// Set the source port range in one of the following formats: 1/200 or 80/80. A value of -1/-1 indicates all ports.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1/200
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The type of the ACL rule: Valid values:
	//
	// 	- **LAN**: The ACL rule controls network traffic transmitted through private IP addresses.
	//
	// 	- **WAN**: The ACL rule controls network traffic transmitted through public IP addresses.
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
