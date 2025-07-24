// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateVpcRequest
	GetCidrBlock() *string
	SetClientToken(v string) *CreateVpcRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVpcRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateVpcRequest
	GetDryRun() *bool
	SetEnableDnsHostname(v bool) *CreateVpcRequest
	GetEnableDnsHostname() *bool
	SetEnableIpv6(v bool) *CreateVpcRequest
	GetEnableIpv6() *bool
	SetIpv4CidrMask(v int32) *CreateVpcRequest
	GetIpv4CidrMask() *int32
	SetIpv4IpamPoolId(v string) *CreateVpcRequest
	GetIpv4IpamPoolId() *string
	SetIpv6CidrBlock(v string) *CreateVpcRequest
	GetIpv6CidrBlock() *string
	SetIpv6CidrMask(v int32) *CreateVpcRequest
	GetIpv6CidrMask() *int32
	SetIpv6IpamPoolId(v string) *CreateVpcRequest
	GetIpv6IpamPoolId() *string
	SetIpv6Isp(v string) *CreateVpcRequest
	GetIpv6Isp() *string
	SetOwnerAccount(v string) *CreateVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateVpcRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVpcRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpcRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateVpcRequestTag) *CreateVpcRequest
	GetTag() []*CreateVpcRequestTag
	SetUserCidr(v string) *CreateVpcRequest
	GetUserCidr() *string
	SetVpcName(v string) *CreateVpcRequest
	GetVpcName() *string
}

type CreateVpcRequest struct {
	// The CIDR block of the VPC.
	//
	// 	- We recommend using the private IPv4 address specified in RFC 1918 as the primary IPv4 CIDR block of the VPC with a recommended mask length of 16 to 28 bits. For example, 10.0.0.0/16, 172.16.0.0/16, and 192.168.0.0/16.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the primary IPv4 CIDR block.
	//
	// example:
	//
	// 172.16.0.0/12
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the VPC.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my first Vpc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Whether to enable the DNS hostname feature. Values:
	//
	// - **false*	- (default): Not enabled.
	//
	// - **true**: Enabled.
	//
	// example:
	//
	// false
	EnableDnsHostname *bool `json:"EnableDnsHostname,omitempty" xml:"EnableDnsHostname,omitempty"`
	// Indicates whether IPv6 is enabled. Valid values:
	//
	// 	- **false*	- (default): disabled.
	//
	// 	- **true**: enabled.
	//
	// example:
	//
	// false
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// Allocate VPC from the IPAM address pool by inputting a mask.
	//
	// > When creating a VPC with a specified IPAM address pool, at least one of the parameters CidrBlock or Ipv4CidrMask must be provided.
	//
	// example:
	//
	// 12
	Ipv4CidrMask *int32 `json:"Ipv4CidrMask,omitempty" xml:"Ipv4CidrMask,omitempty"`
	// The ID of the IP Address Manager (IPAM) pool of the IPv4 type.
	//
	// example:
	//
	// ipam-pool-sycmt3p2a9v63i****
	Ipv4IpamPoolId *string `json:"Ipv4IpamPoolId,omitempty" xml:"Ipv4IpamPoolId,omitempty"`
	// The IPv6 CIDR block of the VPC. If you enable IPv6 for a VPC, the system allocates an IPv6 CIDR block. To specify an IPv6 CIDR block, you must call the [AllocateVpcIpv6Cidr](https://help.aliyun.com/document_detail/448916.html) operation to reserve the specified IPv6 CIDR block.
	//
	// example:
	//
	// 2408:XXXX:0:6a::/56
	Ipv6CidrBlock  *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	Ipv6CidrMask   *int32  `json:"Ipv6CidrMask,omitempty" xml:"Ipv6CidrMask,omitempty"`
	Ipv6IpamPoolId *string `json:"Ipv6IpamPoolId,omitempty" xml:"Ipv6IpamPoolId,omitempty"`
	// The type of the IPv6 CIDR block of the VPC. Valid values:
	//
	// 	- **BGP*	- (default)
	//
	// 	- **ChinaMobile**
	//
	// 	- **ChinaUnicom**
	//
	// 	- **ChinaTelecom**
	//
	// >  If you are allowed to use single-ISP bandwidth, you can set the value to **ChinaTelecom**, **ChinaUnicom**, or **ChinaMobile**.
	//
	// example:
	//
	// BGP
	Ipv6Isp      *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the VPC belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a resource group?](https://help.aliyun.com/document_detail/94475.html)
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag of the resource.
	Tag []*CreateVpcRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The user CIDR block. Separate user CIDR blocks with commas (,). You can specify up to three user CIDR blocks.
	//
	// For more information about user CIDR blocks, see the `What is a user CIDR block?` section in [VPC FAQ](https://help.aliyun.com/document_detail/185311.html).
	//
	// example:
	//
	// 192.168.0.0/12
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	// The name of the VPC.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s CreateVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateVpcRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpcRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVpcRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpcRequest) GetEnableDnsHostname() *bool {
	return s.EnableDnsHostname
}

func (s *CreateVpcRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *CreateVpcRequest) GetIpv4CidrMask() *int32 {
	return s.Ipv4CidrMask
}

func (s *CreateVpcRequest) GetIpv4IpamPoolId() *string {
	return s.Ipv4IpamPoolId
}

func (s *CreateVpcRequest) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *CreateVpcRequest) GetIpv6CidrMask() *int32 {
	return s.Ipv6CidrMask
}

func (s *CreateVpcRequest) GetIpv6IpamPoolId() *string {
	return s.Ipv6IpamPoolId
}

func (s *CreateVpcRequest) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *CreateVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpcRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpcRequest) GetTag() []*CreateVpcRequestTag {
	return s.Tag
}

func (s *CreateVpcRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *CreateVpcRequest) GetVpcName() *string {
	return s.VpcName
}

func (s *CreateVpcRequest) SetCidrBlock(v string) *CreateVpcRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateVpcRequest) SetClientToken(v string) *CreateVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcRequest) SetDescription(v string) *CreateVpcRequest {
	s.Description = &v
	return s
}

func (s *CreateVpcRequest) SetDryRun(v bool) *CreateVpcRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcRequest) SetEnableDnsHostname(v bool) *CreateVpcRequest {
	s.EnableDnsHostname = &v
	return s
}

func (s *CreateVpcRequest) SetEnableIpv6(v bool) *CreateVpcRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateVpcRequest) SetIpv4CidrMask(v int32) *CreateVpcRequest {
	s.Ipv4CidrMask = &v
	return s
}

func (s *CreateVpcRequest) SetIpv4IpamPoolId(v string) *CreateVpcRequest {
	s.Ipv4IpamPoolId = &v
	return s
}

func (s *CreateVpcRequest) SetIpv6CidrBlock(v string) *CreateVpcRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *CreateVpcRequest) SetIpv6CidrMask(v int32) *CreateVpcRequest {
	s.Ipv6CidrMask = &v
	return s
}

func (s *CreateVpcRequest) SetIpv6IpamPoolId(v string) *CreateVpcRequest {
	s.Ipv6IpamPoolId = &v
	return s
}

func (s *CreateVpcRequest) SetIpv6Isp(v string) *CreateVpcRequest {
	s.Ipv6Isp = &v
	return s
}

func (s *CreateVpcRequest) SetOwnerAccount(v string) *CreateVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpcRequest) SetOwnerId(v int64) *CreateVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVpcRequest) SetRegionId(v string) *CreateVpcRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcRequest) SetResourceGroupId(v string) *CreateVpcRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcRequest) SetResourceOwnerAccount(v string) *CreateVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpcRequest) SetResourceOwnerId(v int64) *CreateVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpcRequest) SetTag(v []*CreateVpcRequestTag) *CreateVpcRequest {
	s.Tag = v
	return s
}

func (s *CreateVpcRequest) SetUserCidr(v string) *CreateVpcRequest {
	s.UserCidr = &v
	return s
}

func (s *CreateVpcRequest) SetVpcName(v string) *CreateVpcRequest {
	s.VpcName = &v
	return s
}

func (s *CreateVpcRequest) Validate() error {
	return dara.Validate(s)
}

type CreateVpcRequestTag struct {
	// The key of tag N to add to the resource. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be at most 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, but cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpcRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpcRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVpcRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVpcRequestTag) SetKey(v string) *CreateVpcRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpcRequestTag) SetValue(v string) *CreateVpcRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVpcRequestTag) Validate() error {
	return dara.Validate(s)
}
