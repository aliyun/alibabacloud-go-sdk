// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupEgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RevokeSecurityGroupEgressRequest
	GetClientToken() *string
	SetDescription(v string) *RevokeSecurityGroupEgressRequest
	GetDescription() *string
	SetDestCidrIp(v string) *RevokeSecurityGroupEgressRequest
	GetDestCidrIp() *string
	SetDestGroupId(v string) *RevokeSecurityGroupEgressRequest
	GetDestGroupId() *string
	SetDestGroupOwnerAccount(v string) *RevokeSecurityGroupEgressRequest
	GetDestGroupOwnerAccount() *string
	SetDestGroupOwnerId(v int64) *RevokeSecurityGroupEgressRequest
	GetDestGroupOwnerId() *int64
	SetDestPrefixListId(v string) *RevokeSecurityGroupEgressRequest
	GetDestPrefixListId() *string
	SetIpProtocol(v string) *RevokeSecurityGroupEgressRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *RevokeSecurityGroupEgressRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *RevokeSecurityGroupEgressRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *RevokeSecurityGroupEgressRequest
	GetNicType() *string
	SetOwnerAccount(v string) *RevokeSecurityGroupEgressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeSecurityGroupEgressRequest
	GetOwnerId() *int64
	SetPermissions(v []*RevokeSecurityGroupEgressRequestPermissions) *RevokeSecurityGroupEgressRequest
	GetPermissions() []*RevokeSecurityGroupEgressRequestPermissions
	SetPolicy(v string) *RevokeSecurityGroupEgressRequest
	GetPolicy() *string
	SetPortRange(v string) *RevokeSecurityGroupEgressRequest
	GetPortRange() *string
	SetPriority(v string) *RevokeSecurityGroupEgressRequest
	GetPriority() *string
	SetRegionId(v string) *RevokeSecurityGroupEgressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RevokeSecurityGroupEgressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeSecurityGroupEgressRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *RevokeSecurityGroupEgressRequest
	GetSecurityGroupId() *string
	SetSecurityGroupRuleId(v []*string) *RevokeSecurityGroupEgressRequest
	GetSecurityGroupRuleId() []*string
	SetSourceCidrIp(v string) *RevokeSecurityGroupEgressRequest
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *RevokeSecurityGroupEgressRequest
	GetSourcePortRange() *string
}

type RevokeSecurityGroupEgressRequest struct {
	// A client token used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Description` to specify the rule description.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.DestCidrIp` to specify the destination IPv4 Classless Inter-Domain Routing (CIDR) block.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.DestGroupId` to specify the destination security group ID.
	//
	// example:
	//
	// sg-bp67acfmxa123b****
	DestGroupId *string `json:"DestGroupId,omitempty" xml:"DestGroupId,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.DestGroupOwnerAccount` to specify the Alibaba Cloud account that manages the destination security group.
	//
	// example:
	//
	// Test@aliyun.com
	DestGroupOwnerAccount *string `json:"DestGroupOwnerAccount,omitempty" xml:"DestGroupOwnerAccount,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.DestGroupOwnerId` to specify the ID of the Alibaba Cloud account that manages the destination security group.
	//
	// example:
	//
	// 12345678910
	DestGroupOwnerId *int64 `json:"DestGroupOwnerId,omitempty" xml:"DestGroupOwnerId,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.DestPrefixListId` to specify the source prefix list ID.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	DestPrefixListId *string `json:"DestPrefixListId,omitempty" xml:"DestPrefixListId,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.IpProtocol` to specify the protocol type.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Ipv6DestCidrIp` to specify the destination IPv6 Classless Inter-Domain Routing (CIDR) block.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Ipv6SourceCidrIp` to specify the source IPv6 CIDR block.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.NicType` to specify the NIC type.
	//
	// example:
	//
	// intranet
	NicType      *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The security group rules. Array length: 0 to 100.
	Permissions []*RevokeSecurityGroupEgressRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Policy` to configure the Settings for access permissions.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.PortRange` to specify the port range.
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Priority` to specify the rule priority.
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
	// The IDs of security group rules. Array length: 0 to 100.
	SecurityGroupRuleId []*string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty" type:"Repeated"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.SourceCidrIp` to specify the source IPv4 CIDR block.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.SourcePortRange` to specify the source port range.
	//
	// example:
	//
	// 22/22
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s RevokeSecurityGroupEgressRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RevokeSecurityGroupEgressRequest) GetDescription() *string {
	return s.Description
}

func (s *RevokeSecurityGroupEgressRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *RevokeSecurityGroupEgressRequest) GetDestGroupId() *string {
	return s.DestGroupId
}

func (s *RevokeSecurityGroupEgressRequest) GetDestGroupOwnerAccount() *string {
	return s.DestGroupOwnerAccount
}

func (s *RevokeSecurityGroupEgressRequest) GetDestGroupOwnerId() *int64 {
	return s.DestGroupOwnerId
}

func (s *RevokeSecurityGroupEgressRequest) GetDestPrefixListId() *string {
	return s.DestPrefixListId
}

func (s *RevokeSecurityGroupEgressRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *RevokeSecurityGroupEgressRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *RevokeSecurityGroupEgressRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *RevokeSecurityGroupEgressRequest) GetNicType() *string {
	return s.NicType
}

func (s *RevokeSecurityGroupEgressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeSecurityGroupEgressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeSecurityGroupEgressRequest) GetPermissions() []*RevokeSecurityGroupEgressRequestPermissions {
	return s.Permissions
}

func (s *RevokeSecurityGroupEgressRequest) GetPolicy() *string {
	return s.Policy
}

func (s *RevokeSecurityGroupEgressRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *RevokeSecurityGroupEgressRequest) GetPriority() *string {
	return s.Priority
}

func (s *RevokeSecurityGroupEgressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeSecurityGroupEgressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeSecurityGroupEgressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeSecurityGroupEgressRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RevokeSecurityGroupEgressRequest) GetSecurityGroupRuleId() []*string {
	return s.SecurityGroupRuleId
}

func (s *RevokeSecurityGroupEgressRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *RevokeSecurityGroupEgressRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *RevokeSecurityGroupEgressRequest) SetClientToken(v string) *RevokeSecurityGroupEgressRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetDescription(v string) *RevokeSecurityGroupEgressRequest {
	s.Description = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetDestCidrIp(v string) *RevokeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetDestGroupId(v string) *RevokeSecurityGroupEgressRequest {
	s.DestGroupId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetDestGroupOwnerAccount(v string) *RevokeSecurityGroupEgressRequest {
	s.DestGroupOwnerAccount = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetDestGroupOwnerId(v int64) *RevokeSecurityGroupEgressRequest {
	s.DestGroupOwnerId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetDestPrefixListId(v string) *RevokeSecurityGroupEgressRequest {
	s.DestPrefixListId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetIpProtocol(v string) *RevokeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetIpv6DestCidrIp(v string) *RevokeSecurityGroupEgressRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetIpv6SourceCidrIp(v string) *RevokeSecurityGroupEgressRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetNicType(v string) *RevokeSecurityGroupEgressRequest {
	s.NicType = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetOwnerAccount(v string) *RevokeSecurityGroupEgressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetOwnerId(v int64) *RevokeSecurityGroupEgressRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPermissions(v []*RevokeSecurityGroupEgressRequestPermissions) *RevokeSecurityGroupEgressRequest {
	s.Permissions = v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPolicy(v string) *RevokeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPortRange(v string) *RevokeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetPriority(v string) *RevokeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetRegionId(v string) *RevokeSecurityGroupEgressRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetResourceOwnerAccount(v string) *RevokeSecurityGroupEgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetResourceOwnerId(v int64) *RevokeSecurityGroupEgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *RevokeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSecurityGroupRuleId(v []*string) *RevokeSecurityGroupEgressRequest {
	s.SecurityGroupRuleId = v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSourceCidrIp(v string) *RevokeSecurityGroupEgressRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) SetSourcePortRange(v string) *RevokeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequest) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RevokeSecurityGroupEgressRequestPermissions struct {
	// The description of the security group rule. The description must be 1 to 512 characters in length.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 Classless Inter-Domain Routing (CIDR) block for which you want to revoke access permissions. The format supports both CIDR blocks and IPv4 address ranges.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The ID of the destination security group for which you want to revoke access permissions.
	//
	// - Specify at least one of `DestGroupId`, `DestCidrIp`, `Ipv6DestCidrIp`, or `DestPrefixListId`.
	//
	// - If `DestGroupId` is specified but `DestCidrIp` is not, the `NicType` parameter can only be set to intranet.
	//
	// - If both `DestGroupId` and `DestCidrIp` are specified, `DestCidrIp` takes precedence.
	//
	// Note:
	//
	// - Advanced security groups do not support authorization for other security groups.
	//
	// - A maximum of 20 security groups can be authorized for a basic security group.
	//
	// example:
	//
	// sg-bp67acfmxa123b****
	DestGroupId *string `json:"DestGroupId,omitempty" xml:"DestGroupId,omitempty"`
	// The Alibaba Cloud account that manages the destination security group when you revoke a cross-account authorization security group rule.
	//
	// - If neither `DestGroupOwnerAccount` nor `DestGroupOwnerId` is specified, the access permissions for another security group within your account are revoked. Settings for cross-account scenarios require this parameter.
	//
	// - If `DestCidrIp` is specified, this parameter is ignored.
	//
	// example:
	//
	// Test@aliyun.com
	DestGroupOwnerAccount *string `json:"DestGroupOwnerAccount,omitempty" xml:"DestGroupOwnerAccount,omitempty"`
	// The ID of the Alibaba Cloud account that manages the destination security group when you revoke a cross-account authorization security group rule.
	//
	// - If neither `DestGroupOwnerId` nor `DestGroupOwnerAccount` is specified, the access permissions for another security group within your account are revoked. Settings for cross-account scenarios require this parameter.
	//
	// - If `DestCidrIp` is specified, this parameter is ignored.
	//
	// example:
	//
	// 12345678910
	DestGroupOwnerId *string `json:"DestGroupOwnerId,omitempty" xml:"DestGroupOwnerId,omitempty"`
	// The ID of the destination prefix list for which you want to revoke access permissions. You can call [DescribePrefixLists](https://help.aliyun.com/document_detail/205046.html) to query available prefix list IDs.
	//
	// Note:
	//
	// If one of `DestCidrIp`, `Ipv6DestCidrIp`, or `DestGroupId` is specified, this parameter is ignored.
	//
	// For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	DestPrefixListId *string `json:"DestPrefixListId,omitempty" xml:"DestPrefixListId,omitempty"`
	// The protocol type. The value is case-insensitive. Valid values:
	//
	//
	//
	// - TCP.
	//
	// - UDP.
	//
	// - ICMP.
	//
	// - ICMPv6.
	//
	// - GRE.
	//
	// - ALL: all protocols.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 Classless Inter-Domain Routing (CIDR) block for which you want to revoke access permissions. The format supports both CIDR blocks and IPv6 address ranges.
	//
	// > This parameter is valid only for VPC-connected ECS instances that support IPv6 and cannot be specified together with the `DestCidrIp` parameter. Settings for IPv6 destinations use this parameter exclusively.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block. CIDR blocks and IPv6 address ranges are supported.
	//
	// This parameter is used for quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// > This parameter is valid only for VPC-connected ECS instances that support IPv6 and cannot be specified together with the `DestCidrIp` parameter. Settings for IPv6 sources use this parameter exclusively.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network interface controller (NIC) type of the security group rule. For VPC-type security groups, you do not need to configure the network interface controller (NIC) type. The default value is intranet and only intranet is supported. Settings for VPC-type security groups ignore this parameter.
	//
	// > The classic network feature has been taken offline. For details, see [Discontinuation notice](https://help.aliyun.com/document_detail/2833134.html). The network interface controller (NIC) type for classic network-type security group rules. Valid values:
	//
	// > - internet: public network interface controller (NIC).
	//
	// > - intranet: internal network interface controller (NIC).
	//
	// example:
	//
	// intranet
	NicType *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	// The access permissions. Valid values:
	//
	//
	//
	// - accept: Accepts access.
	//
	// - drop: Denies access and returns no deny information. The request timeout or the connection cannot be established.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the transport layer protocol. Valid values:
	//
	//
	//
	// - TCP/UDP: Valid values are 1 to 65535. Separate the start port and the end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// - GRE: -1/-1.
	//
	// - ALL: -1/-1.
	//
	// example:
	//
	// 22/22
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The port list ID.
	//
	// You can call `DescribePortRangeLists` to query available port list IDs.
	//
	// If `Permissions.N.PortRange` is specified, this parameter is ignored.
	//
	// For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The priority of the security group rule. A smaller value indicates a higher priority. Valid values: 1 to 100.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The source IPv4 CIDR block. CIDR blocks and IPv4 address ranges are supported.
	//
	// This parameter is used for quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of source ports that correspond to the transport layer protocol. Valid values:
	//
	//
	//
	// - TCP/UDP: Valid values are 1 to 65535. Separate the start port and the end port with a forward slash (/). Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// - GRE: -1/-1.
	//
	// - ALL: -1/-1.
	//
	// This parameter is used for quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 22/22
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s RevokeSecurityGroupEgressRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupEgressRequestPermissions) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetDescription() *string {
	return s.Description
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetDestGroupId() *string {
	return s.DestGroupId
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetDestGroupOwnerAccount() *string {
	return s.DestGroupOwnerAccount
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetDestGroupOwnerId() *string {
	return s.DestGroupOwnerId
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetDestPrefixListId() *string {
	return s.DestPrefixListId
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetNicType() *string {
	return s.NicType
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetPriority() *string {
	return s.Priority
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *RevokeSecurityGroupEgressRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetDescription(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.Description = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetDestCidrIp(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetDestGroupId(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.DestGroupId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetDestGroupOwnerAccount(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.DestGroupOwnerAccount = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetDestGroupOwnerId(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.DestGroupOwnerId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetDestPrefixListId(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.DestPrefixListId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetIpProtocol(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetIpv6DestCidrIp(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetIpv6SourceCidrIp(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetNicType(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.NicType = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetPolicy(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetPortRange(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetPortRangeListId(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.PortRangeListId = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetPriority(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetSourceCidrIp(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) SetSourcePortRange(v string) *RevokeSecurityGroupEgressRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *RevokeSecurityGroupEgressRequestPermissions) Validate() error {
	return dara.Validate(s)
}
