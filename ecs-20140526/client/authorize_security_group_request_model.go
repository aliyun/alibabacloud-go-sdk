// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AuthorizeSecurityGroupRequest
	GetClientToken() *string
	SetDescription(v string) *AuthorizeSecurityGroupRequest
	GetDescription() *string
	SetDestCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetDestCidrIp() *string
	SetIpProtocol(v string) *AuthorizeSecurityGroupRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *AuthorizeSecurityGroupRequest
	GetNicType() *string
	SetOwnerAccount(v string) *AuthorizeSecurityGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AuthorizeSecurityGroupRequest
	GetOwnerId() *int64
	SetPermissions(v []*AuthorizeSecurityGroupRequestPermissions) *AuthorizeSecurityGroupRequest
	GetPermissions() []*AuthorizeSecurityGroupRequestPermissions
	SetPolicy(v string) *AuthorizeSecurityGroupRequest
	GetPolicy() *string
	SetPortRange(v string) *AuthorizeSecurityGroupRequest
	GetPortRange() *string
	SetPriority(v string) *AuthorizeSecurityGroupRequest
	GetPriority() *string
	SetRegionId(v string) *AuthorizeSecurityGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AuthorizeSecurityGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AuthorizeSecurityGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *AuthorizeSecurityGroupRequest
	GetSecurityGroupId() *string
	SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetSourceCidrIp() *string
	SetSourceGroupId(v string) *AuthorizeSecurityGroupRequest
	GetSourceGroupId() *string
	SetSourceGroupOwnerAccount(v string) *AuthorizeSecurityGroupRequest
	GetSourceGroupOwnerAccount() *string
	SetSourceGroupOwnerId(v int64) *AuthorizeSecurityGroupRequest
	GetSourceGroupOwnerId() *int64
	SetSourcePortRange(v string) *AuthorizeSecurityGroupRequest
	GetSourcePortRange() *string
	SetSourcePrefixListId(v string) *AuthorizeSecurityGroupRequest
	GetSourcePrefixListId() *string
}

type AuthorizeSecurityGroupRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Description` to specify the description of the security group rule.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.DestCidrIp` to specify the destination IPv4 CIDR block.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
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
	// Deprecated. Use `Permissions.N.Ipv6DestCidrIp` to specify the destination IPv6 CIDR block.
	//
	// example:
	//
	// null
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Ipv6SourceCidrIp` to specify the source IPv6 Classless Inter-Domain Routing block.
	//
	// example:
	//
	// 2001:250:6000::***
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
	Permissions []*AuthorizeSecurityGroupRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.Policy` to set access permissions.
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
	// Deprecated. Use `Permissions.N.Priority` to specify the security group rule priority.
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
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.SourceCidrIp` to specify the source IPv4 Classless Inter-Domain Routing block.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.SourceGroupId` to specify the source security group ID.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.SourceGroupOwnerAccount` to specify the Alibaba Cloud account that owns the source security group.
	//
	// example:
	//
	// test@aliyun.com
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.SourceGroupOwnerId` to specify the ID of the Alibaba Cloud account that owns the source security group.
	//
	// example:
	//
	// 1234567890
	SourceGroupOwnerId *int64 `json:"SourceGroupOwnerId,omitempty" xml:"SourceGroupOwnerId,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.SourcePortRange` to specify the source port range.
	//
	// example:
	//
	// 22/22
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// Deprecated
	//
	// Deprecated. Use `Permissions.N.SourcePrefixListId` to specify the source prefix list ID.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	SourcePrefixListId *string `json:"SourcePrefixListId,omitempty" xml:"SourcePrefixListId,omitempty"`
}

func (s AuthorizeSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AuthorizeSecurityGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *AuthorizeSecurityGroupRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetNicType() *string {
	return s.NicType
}

func (s *AuthorizeSecurityGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AuthorizeSecurityGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AuthorizeSecurityGroupRequest) GetPermissions() []*AuthorizeSecurityGroupRequestPermissions {
	return s.Permissions
}

func (s *AuthorizeSecurityGroupRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupRequest) GetPriority() *string {
	return s.Priority
}

func (s *AuthorizeSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeSecurityGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AuthorizeSecurityGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AuthorizeSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AuthorizeSecurityGroupRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *AuthorizeSecurityGroupRequest) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *AuthorizeSecurityGroupRequest) GetSourceGroupOwnerId() *int64 {
	return s.SourceGroupOwnerId
}

func (s *AuthorizeSecurityGroupRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupRequest) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *AuthorizeSecurityGroupRequest) SetClientToken(v string) *AuthorizeSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetDescription(v string) *AuthorizeSecurityGroupRequest {
	s.Description = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetDestCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetNicType(v string) *AuthorizeSecurityGroupRequest {
	s.NicType = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetOwnerAccount(v string) *AuthorizeSecurityGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetOwnerId(v int64) *AuthorizeSecurityGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPermissions(v []*AuthorizeSecurityGroupRequestPermissions) *AuthorizeSecurityGroupRequest {
	s.Permissions = v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPolicy(v string) *AuthorizeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPortRange(v string) *AuthorizeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPriority(v string) *AuthorizeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetRegionId(v string) *AuthorizeSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetResourceOwnerAccount(v string) *AuthorizeSecurityGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetResourceOwnerId(v int64) *AuthorizeSecurityGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceGroupId(v string) *AuthorizeSecurityGroupRequest {
	s.SourceGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceGroupOwnerAccount(v string) *AuthorizeSecurityGroupRequest {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceGroupOwnerId(v int64) *AuthorizeSecurityGroupRequest {
	s.SourceGroupOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourcePrefixListId(v string) *AuthorizeSecurityGroupRequest {
	s.SourcePrefixListId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) Validate() error {
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

type AuthorizeSecurityGroupRequestPermissions struct {
	// The description of the security group rule. The description must be 1 to 512 characters in length.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 CIDR block. CIDR blocks and IPv4 address ranges are supported.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The network layer or transport layer protocol. Two types of values are supported:
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
	// - ALL: all protocols.
	//
	// 2. Protocol numbers that comply with IANA specifications, which are integers from 0 to 255. The following regions currently support this feature:
	//
	// - Philippines
	//
	// - UK
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
	// The destination IPv6 CIDR block. CIDR format and IPv6 format address ranges are supported.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// > This parameter is valid only for VPC-connected ECS instances that support IPv6. This parameter and `DestCidrIp` cannot be specified at the same time.
	//
	// example:
	//
	// 2001:250:6000::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block for which you want to set access permissions. Settings for CIDR format and IPv6 format address ranges are supported.
	//
	// > This parameter is valid only for VPC-connected ECS instances that support IPv6. This parameter and `SourceCidrIp` cannot be specified at the same time.
	//
	// example:
	//
	// 2001:250:6000::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network interface controller (NIC) type for a classic network type security group rule. Valid values:
	//
	// - internet: public network interface controller (NIC).
	//
	// - intranet: internal network interface controller (NIC).
	//
	// For VPC security group rules, you do not need to set the network interface controller (NIC) type parameter. The default value is intranet, and only intranet is supported.
	//
	// When you set security groups to access each other (only DestGroupId is specified), only intranet is supported.
	//
	// Default value: internet.
	//
	// example:
	//
	// intranet
	NicType *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	// Settings for access permissions. Valid values:
	//
	// - accept: accepts access.
	//
	// - drop: denies access and does not return a deny message. The request appears to timeout or the connection cannot be established.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the protocol for the security group. Valid values:
	//
	// - TCP/UDP: Valid values are 1 to 65535. Separate the start port and the stop port with a forward slash (/). Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// - GRE: -1/-1.
	//
	// - ALL: -1/-1.
	//
	// For more information about common ports, see [Common scenarios for ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The port address book ID.
	//
	// You can invoke `DescribePortRangeLists` to query available port address book IDs.
	//
	// - If you specify `Permissions.N.PortRange`, this parameter is ignored.
	//
	// - Port address books are not supported for security groups with the classic network type. For more information about security group and port address book limits, see [Security group limits](~~25412#SecurityGroupQuota1~~). Settings for port address books are not available for classic network security groups.
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
	// The source IPv4 CIDR block for which you want to set access permissions. Settings for CIDR format and IPv4 format address ranges are supported.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The ID of the source security group for which you want to set access permissions.
	//
	// - You must specify at least one of the following parameters: `SourceGroupId`, `SourceCidrIp`, `Ipv6SourceCidrIp`, or `SourcePrefixListId`.
	//
	// - If `SourceGroupId` is specified but `SourceCidrIp` or `Ipv6SourceCidrIp` is not specified, the `NicType` parameter can only be set to `intranet`.
	//
	// - If both `SourceGroupId` and `SourceCidrIp` are specified, `SourceCidrIp` takes precedence.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// The Alibaba Cloud account that owns the source security group when you set a cross-account security group rule.
	//
	// - If neither `SourceGroupOwnerAccount` nor `SourceGroupOwnerId` is set, access permissions are configured for another security group within your account.
	//
	// - If the `SourceCidrIp` parameter is set, the `SourceGroupOwnerAccount` parameter is ignored.
	//
	// example:
	//
	// test@aliyun.com
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// The ID of the Alibaba Cloud account that owns the source security group when you set a cross-account security group rule.
	//
	// - If neither `SourceGroupOwnerAccount` nor `SourceGroupOwnerId` is set, access permissions are configured for another security group within your account.
	//
	// - If the `SourceCidrIp` parameter is set, the `SourceGroupOwnerAccount` parameter is ignored.
	//
	// example:
	//
	// 1234567890
	SourceGroupOwnerId *int64 `json:"SourceGroupOwnerId,omitempty" xml:"SourceGroupOwnerId,omitempty"`
	// The range of source ports that correspond to the protocol for the security group. Valid values:
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
	// 7000/8000
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The ID of the source prefix list for which you want to set access permissions. You can call [DescribePrefixLists](https://help.aliyun.com/document_detail/205046.html) to query available prefix list IDs.
	//
	// Notes:
	//
	// If you specify `SourceCidrIp`, `Ipv6SourceCidrIp`, or `SourceGroupId`, this parameter is ignored.
	//
	// For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	SourcePrefixListId *string `json:"SourcePrefixListId,omitempty" xml:"SourcePrefixListId,omitempty"`
}

func (s AuthorizeSecurityGroupRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupRequestPermissions) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetDescription() *string {
	return s.Description
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetNicType() *string {
	return s.NicType
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetPriority() *string {
	return s.Priority
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourceGroupOwnerId() *int64 {
	return s.SourceGroupOwnerId
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetDescription(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Description = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetDestCidrIp(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetIpProtocol(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetNicType(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.NicType = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetPolicy(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetPortRange(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetPortRangeListId(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.PortRangeListId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetPriority(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourceGroupId(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourceGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourceGroupOwnerAccount(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourceGroupOwnerId(v int64) *AuthorizeSecurityGroupRequestPermissions {
	s.SourceGroupOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourcePortRange(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourcePrefixListId(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourcePrefixListId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) Validate() error {
	return dara.Validate(s)
}
