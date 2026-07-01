// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupEgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AuthorizeSecurityGroupEgressRequest
	GetClientToken() *string
	SetDescription(v string) *AuthorizeSecurityGroupEgressRequest
	GetDescription() *string
	SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestCidrIp() *string
	SetDestGroupId(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestGroupId() *string
	SetDestGroupOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestGroupOwnerAccount() *string
	SetDestGroupOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest
	GetDestGroupOwnerId() *int64
	SetDestPrefixListId(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestPrefixListId() *string
	SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *AuthorizeSecurityGroupEgressRequest
	GetNicType() *string
	SetOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest
	GetOwnerId() *int64
	SetPermissions(v []*AuthorizeSecurityGroupEgressRequestPermissions) *AuthorizeSecurityGroupEgressRequest
	GetPermissions() []*AuthorizeSecurityGroupEgressRequestPermissions
	SetPolicy(v string) *AuthorizeSecurityGroupEgressRequest
	GetPolicy() *string
	SetPortRange(v string) *AuthorizeSecurityGroupEgressRequest
	GetPortRange() *string
	SetPriority(v string) *AuthorizeSecurityGroupEgressRequest
	GetPriority() *string
	SetRegionId(v string) *AuthorizeSecurityGroupEgressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressRequest
	GetSecurityGroupId() *string
	SetSourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequest
	GetSourcePortRange() *string
}

type AuthorizeSecurityGroupEgressRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
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
	// sg-bp67acfmxazb4p****
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
	// ALL
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
	// The security group rules. Array length: 1 to 100.
	Permissions []*AuthorizeSecurityGroupEgressRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Policy` to configure the access permission settings.
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
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Priority` to specify the rule priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the source security group. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
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
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeSecurityGroupEgressRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDescription() *string {
	return s.Description
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestGroupId() *string {
	return s.DestGroupId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestGroupOwnerAccount() *string {
	return s.DestGroupOwnerAccount
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestGroupOwnerId() *int64 {
	return s.DestGroupOwnerId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestPrefixListId() *string {
	return s.DestPrefixListId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupEgressRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetNicType() *string {
	return s.NicType
}

func (s *AuthorizeSecurityGroupEgressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AuthorizeSecurityGroupEgressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPermissions() []*AuthorizeSecurityGroupEgressRequestPermissions {
	return s.Permissions
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPriority() *string {
	return s.Priority
}

func (s *AuthorizeSecurityGroupEgressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AuthorizeSecurityGroupEgressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupEgressRequest) SetClientToken(v string) *AuthorizeSecurityGroupEgressRequest {
	s.ClientToken = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDescription(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Description = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestGroupId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestGroupOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestGroupOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestGroupOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest {
	s.DestGroupOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestPrefixListId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestPrefixListId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetNicType(v string) *AuthorizeSecurityGroupEgressRequest {
	s.NicType = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest {
	s.OwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPermissions(v []*AuthorizeSecurityGroupEgressRequestPermissions) *AuthorizeSecurityGroupEgressRequest {
	s.Permissions = v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPolicy(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPriority(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetRegionId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetResourceOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetResourceOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) Validate() error {
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

type AuthorizeSecurityGroupEgressRequestPermissions struct {
	// The description of the security group rule. The description must be 1 to 512 characters in length.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 Classless Inter-Domain Routing (CIDR) block for which you want to configure access permission settings. Both CIDR format and IPv4 format address ranges are supported.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The ID of the destination security group for which you want to set access permissions.
	//
	// - You must specify at least one of the following parameters: `DestGroupId`, `DestCidrIp`, `Ipv6DestCidrIp`, or `DestPrefixListId`.
	//
	// - If `DestGroupId` is specified but `DestCidrIp` is not, the `NicType` parameter can only be set to intranet.
	//
	// - If both `DestGroupId` and `DestCidrIp` are specified, `DestCidrIp` takes precedence.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	DestGroupId *string `json:"DestGroupId,omitempty" xml:"DestGroupId,omitempty"`
	// The Alibaba Cloud account that manages the destination security group when you configure a cross-account security group rule settings.
	//
	//
	//
	// - If neither `DestGroupOwnerAccount` nor `DestGroupOwnerId` is specified, the access permissions are configured for another security group within your account.
	//
	// - If the `DestCidrIp` parameter is specified, the `DestGroupOwnerAccount` parameter is ignored.
	//
	// example:
	//
	// Test@aliyun.com
	DestGroupOwnerAccount *string `json:"DestGroupOwnerAccount,omitempty" xml:"DestGroupOwnerAccount,omitempty"`
	// The ID of the Alibaba Cloud account that manages the destination security group when you configure a cross-account security group rule settings.
	//
	//
	//
	// - If neither `DestGroupOwnerId` nor `DestGroupOwnerAccount` is specified, the access permissions are configured for another security group within your account.
	//
	// - If the `DestCidrIp` parameter is specified, the `DestGroupOwnerId` parameter is ignored.
	//
	// example:
	//
	// 12345678910
	DestGroupOwnerId *int64 `json:"DestGroupOwnerId,omitempty" xml:"DestGroupOwnerId,omitempty"`
	// The ID of the destination prefix list for which you want to set access permissions. You can call [DescribePrefixLists](https://help.aliyun.com/document_detail/205046.html) to query available prefix list IDs.
	//
	// Notes:
	//
	// If you specify one of the `DestCidrIp`, `Ipv6DestCidrIp`, or `DestGroupId` parameters, this parameter is ignored.
	//
	// For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	DestPrefixListId *string `json:"DestPrefixListId,omitempty" xml:"DestPrefixListId,omitempty"`
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
	// - ALL: All protocols are supported.
	//
	// 2. Protocol numbers that comply with IANA specifications, which are integers from 0 to 255. The following regions currently support this feature:
	//
	// - Philippines
	//
	// - UK
	//
	// - Malaysia
	//
	// - Hohhot
	//
	// - Qingdao
	//
	// - US West
	//
	// - Singapore
	//
	// example:
	//
	// ALL
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 Classless Inter-Domain Routing (CIDR) block for which you want to configure access permission settings. Both CIDR format and IPv6 format address ranges are supported.
	//
	// > This parameter is valid only for VPC-type ECS instances that support IPv6. This parameter and `DestCidrIp` cannot be specified at the same time.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block. Both CIDR format and IPv6 format address ranges are supported.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// > This parameter is valid only for VPC-type ECS instances that support IPv6. This parameter and `DestCidrIp` cannot be specified at the same time.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network interface controller (NIC) type settings for a classic network security group rule. Valid values:
	//
	//
	//
	// - internet: public network interface controller (NIC).
	//
	// - intranet: internal network interface controller (NIC).
	//
	//     - For VPC-type security group rules, you do not need to configure the network interface controller (NIC) type settings. The default value is intranet.
	//
	//     - When you configure security groups to access each other, meaning only the DestGroupId parameter is specified, the value can only be intranet.
	//
	// Default value: internet.
	//
	// example:
	//
	// intranet
	NicType *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	// The access permission settings. Valid values:
	//
	//
	//
	// - accept: Accepts access.
	//
	// - drop: Denies access and does not return a deny message. The request times out or a timeout error similar to a connection failure is returned.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the protocol for the security group. Valid values:
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
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The port list ID.
	//
	// You can invoke `DescribePortRangeLists` to query available port list IDs.
	//
	// - If you specify `Permissions.N.PortRange`, this parameter is ignored.
	//
	// - Port lists are not supported for classic network security group settings. For more information about security group and port list limits, see [Security group limits](~~25412#SecurityGroupQuota1~~).
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
	// The source IPv4 CIDR block. Both CIDR format and IPv4 format address ranges are supported.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of source ports that correspond to the protocol for the security group. Valid values:
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
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeSecurityGroupEgressRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressRequestPermissions) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDescription() *string {
	return s.Description
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestGroupId() *string {
	return s.DestGroupId
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestGroupOwnerAccount() *string {
	return s.DestGroupOwnerAccount
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestGroupOwnerId() *int64 {
	return s.DestGroupOwnerId
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestPrefixListId() *string {
	return s.DestPrefixListId
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetNicType() *string {
	return s.NicType
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetPriority() *string {
	return s.Priority
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDescription(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Description = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestGroupId(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestGroupOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestGroupOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestGroupOwnerId(v int64) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestGroupOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestPrefixListId(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestPrefixListId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetNicType(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.NicType = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetPolicy(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetPortRange(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetPortRangeListId(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.PortRangeListId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetPriority(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetSourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) Validate() error {
	return dara.Validate(s)
}
