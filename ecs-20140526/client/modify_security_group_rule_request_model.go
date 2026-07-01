// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifySecurityGroupRuleRequest
	GetClientToken() *string
	SetDescription(v string) *ModifySecurityGroupRuleRequest
	GetDescription() *string
	SetDestCidrIp(v string) *ModifySecurityGroupRuleRequest
	GetDestCidrIp() *string
	SetIpProtocol(v string) *ModifySecurityGroupRuleRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *ModifySecurityGroupRuleRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *ModifySecurityGroupRuleRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *ModifySecurityGroupRuleRequest
	GetNicType() *string
	SetOwnerAccount(v string) *ModifySecurityGroupRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityGroupRuleRequest
	GetOwnerId() *int64
	SetPolicy(v string) *ModifySecurityGroupRuleRequest
	GetPolicy() *string
	SetPortRange(v string) *ModifySecurityGroupRuleRequest
	GetPortRange() *string
	SetPortRangeListId(v string) *ModifySecurityGroupRuleRequest
	GetPortRangeListId() *string
	SetPriority(v string) *ModifySecurityGroupRuleRequest
	GetPriority() *string
	SetRegionId(v string) *ModifySecurityGroupRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySecurityGroupRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityGroupRuleRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *ModifySecurityGroupRuleRequest
	GetSecurityGroupId() *string
	SetSecurityGroupRuleId(v string) *ModifySecurityGroupRuleRequest
	GetSecurityGroupRuleId() *string
	SetSourceCidrIp(v string) *ModifySecurityGroupRuleRequest
	GetSourceCidrIp() *string
	SetSourceGroupId(v string) *ModifySecurityGroupRuleRequest
	GetSourceGroupId() *string
	SetSourceGroupOwnerAccount(v string) *ModifySecurityGroupRuleRequest
	GetSourceGroupOwnerAccount() *string
	SetSourceGroupOwnerId(v int64) *ModifySecurityGroupRuleRequest
	GetSourceGroupOwnerId() *int64
	SetSourcePortRange(v string) *ModifySecurityGroupRuleRequest
	GetSourcePortRange() *string
	SetSourcePrefixListId(v string) *ModifySecurityGroupRuleRequest
	GetSourcePrefixListId() *string
}

type ModifySecurityGroupRuleRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The ClientToken parameter supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the security group rule. The description must be 1 to 512 characters in length.
	//
	// example:
	//
	// This is a new security group rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 Classless Inter-Domain Routing (CIDR) block. CIDR format and IPv4 format IP address ranges are supported.
	//
	// Default value: null.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The network-layer or transport-layer protocol. Two types of values are supported:
	//
	// 1. Case-insensitive protocol names. Valid values:
	//
	// - ICMP
	//
	// - GRE
	//
	// - TCP
	//
	// - UDP
	//
	// - ALL: all protocols are supported.
	//
	// 2. Protocol numbers that comply with IANA specifications, which are integers from 0 to 255. The following regions currently support this feature:
	//
	// - Philippines
	//
	// - UK (London)
	//
	// - Malaysia
	//
	// - China (Hohhot)
	//
	// - China (Qingdao)
	//
	// - US (Virginia)
	//
	// - Singapore
	//
	// example:
	//
	// ALL
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Settings for the destination IPv6 CIDR block. Classless Inter-Domain Routing (CIDR) format and IPv6 format IP address ranges are supported.
	//
	// >Only VPC-type IP addresses are supported. This parameter and `DestCidrIp` cannot be specified at the same time.
	//
	// Default value: null.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// Settings for the source IPv6 CIDR block for the access permissions. Classless Inter-Domain Routing (CIDR) format and IPv6 format IP address ranges are supported.
	//
	// > Only VPC-type IP addresses are supported. This parameter and `SourceCidrIp` cannot be specified at the same time.
	//
	// Default value: null.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network interface controller (NIC) type.
	//
	// > When you modify a rule by specifying the security group rule ID, this parameter cannot be modified. To make such a change, add a new rule and then delete the current rule.
	//
	// example:
	//
	// intranet
	NicType      *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The access permissions. Valid values:
	//
	//
	//
	// - accept: Accepts access.
	//
	// - drop: Denies access and does not return a deny response.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the transport-layer protocol of the security group. Valid values:
	//
	//
	//
	// - TCP/UDP: valid values are 1 to 65535. Separate the start port and the end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// - GRE: -1/-1.
	//
	// - ALL: -1/-1.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The port address book ID.
	//
	// You can call `DescribePortRangeLists` to query available port address book IDs.
	//
	// This parameter is ignored if you specify the PortRange parameter.
	//
	// For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The priority of the security group rule. Valid values: 1 to 100.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the security group. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The security group rule ID. You can call [DescribeSecurityGroupAttribute](https://help.aliyun.com/document_detail/2679845.html) to query security group rule IDs.
	//
	// example:
	//
	// sgr-bp67acfmxa123b***
	SecurityGroupRuleId *string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty"`
	// Settings for the source IPv4 CIDR block for the access permissions. Classless Inter-Domain Routing (CIDR) format and IPv4 format IP address ranges are supported.
	//
	// Default value: null.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// Settings for the ID of the source security group for the access permissions. Specify at least one of `SourceGroupId` and `SourceCidrIp`.
	//
	// - If `SourceGroupId` is specified but `SourceCidrIp` is not, the `NicType` parameter can only be set to intranet.
	//
	// - If both `SourceGroupId` and `SourceCidrIp` are specified, `SourceCidrIp` takes precedence by default.
	//
	// example:
	//
	// sg-bp67acfmxa123b****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// Settings for the Alibaba Cloud account that owns the source security group when you configure a cross-account security group rule for access permissions.
	//
	//
	//
	// - If neither `SourceGroupOwnerAccount` nor `SourceGroupOwnerId` is configured, the rule is configured for the access permissions of your other security groups.
	//
	// - If the `SourceCidrIp` parameter is specified, the `SourceGroupOwnerAccount` parameter is invalid.
	//
	// example:
	//
	// EcsforCloud@Alibaba.com
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// Settings for the Alibaba Cloud account ID that owns the source security group when you configure a cross-account security group rule for access permissions.
	//
	//
	//
	// - If neither `SourceGroupOwnerId` nor `SourceGroupOwnerAccount` is configured, the rule is configured for the access permissions of your other security groups.
	//
	// - If the `SourceCidrIp` parameter is specified, the `SourceGroupOwnerId` parameter is invalid.
	//
	// example:
	//
	// 12345678910
	SourceGroupOwnerId *int64 `json:"SourceGroupOwnerId,omitempty" xml:"SourceGroupOwnerId,omitempty"`
	// The range of source ports that correspond to the transport-layer protocol of the security group. Valid values:
	//
	//
	//
	// - TCP/UDP: valid values are 1 to 65535. Separate the start port and the end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// - GRE: -1/-1.
	//
	// - ALL: -1/-1.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// Settings for the ID of the source prefix list for the access permissions. You can invoke [DescribePrefixLists](https://help.aliyun.com/document_detail/205046.html) to query available prefix list IDs.
	//
	// This parameter is ignored if you specify one of `SourceCidrIp`, `Ipv6SourceCidrIp`, or `SourceGroupId`.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	SourcePrefixListId *string `json:"SourcePrefixListId,omitempty" xml:"SourcePrefixListId,omitempty"`
}

func (s ModifySecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifySecurityGroupRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupRuleRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *ModifySecurityGroupRuleRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifySecurityGroupRuleRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *ModifySecurityGroupRuleRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *ModifySecurityGroupRuleRequest) GetNicType() *string {
	return s.NicType
}

func (s *ModifySecurityGroupRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityGroupRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityGroupRuleRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ModifySecurityGroupRuleRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifySecurityGroupRuleRequest) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *ModifySecurityGroupRuleRequest) GetPriority() *string {
	return s.Priority
}

func (s *ModifySecurityGroupRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityGroupRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityGroupRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityGroupRuleRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifySecurityGroupRuleRequest) GetSecurityGroupRuleId() *string {
	return s.SecurityGroupRuleId
}

func (s *ModifySecurityGroupRuleRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifySecurityGroupRuleRequest) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *ModifySecurityGroupRuleRequest) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *ModifySecurityGroupRuleRequest) GetSourceGroupOwnerId() *int64 {
	return s.SourceGroupOwnerId
}

func (s *ModifySecurityGroupRuleRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifySecurityGroupRuleRequest) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *ModifySecurityGroupRuleRequest) SetClientToken(v string) *ModifySecurityGroupRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetDescription(v string) *ModifySecurityGroupRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetDestCidrIp(v string) *ModifySecurityGroupRuleRequest {
	s.DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetIpProtocol(v string) *ModifySecurityGroupRuleRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetIpv6DestCidrIp(v string) *ModifySecurityGroupRuleRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetIpv6SourceCidrIp(v string) *ModifySecurityGroupRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetNicType(v string) *ModifySecurityGroupRuleRequest {
	s.NicType = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetOwnerAccount(v string) *ModifySecurityGroupRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetOwnerId(v int64) *ModifySecurityGroupRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetPolicy(v string) *ModifySecurityGroupRuleRequest {
	s.Policy = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetPortRange(v string) *ModifySecurityGroupRuleRequest {
	s.PortRange = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetPortRangeListId(v string) *ModifySecurityGroupRuleRequest {
	s.PortRangeListId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetPriority(v string) *ModifySecurityGroupRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetRegionId(v string) *ModifySecurityGroupRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetResourceOwnerAccount(v string) *ModifySecurityGroupRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetResourceOwnerId(v int64) *ModifySecurityGroupRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSecurityGroupId(v string) *ModifySecurityGroupRuleRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSecurityGroupRuleId(v string) *ModifySecurityGroupRuleRequest {
	s.SecurityGroupRuleId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourceCidrIp(v string) *ModifySecurityGroupRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourceGroupId(v string) *ModifySecurityGroupRuleRequest {
	s.SourceGroupId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourceGroupOwnerAccount(v string) *ModifySecurityGroupRuleRequest {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourceGroupOwnerId(v int64) *ModifySecurityGroupRuleRequest {
	s.SourceGroupOwnerId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourcePortRange(v string) *ModifySecurityGroupRuleRequest {
	s.SourcePortRange = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourcePrefixListId(v string) *ModifySecurityGroupRuleRequest {
	s.SourcePrefixListId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
