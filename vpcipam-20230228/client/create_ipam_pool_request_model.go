// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationDefaultCidrMask(v int32) *CreateIpamPoolRequest
	GetAllocationDefaultCidrMask() *int32
	SetAllocationMaxCidrMask(v int32) *CreateIpamPoolRequest
	GetAllocationMaxCidrMask() *int32
	SetAllocationMinCidrMask(v int32) *CreateIpamPoolRequest
	GetAllocationMinCidrMask() *int32
	SetAutoImport(v bool) *CreateIpamPoolRequest
	GetAutoImport() *bool
	SetClientToken(v string) *CreateIpamPoolRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateIpamPoolRequest
	GetDryRun() *bool
	SetIpVersion(v string) *CreateIpamPoolRequest
	GetIpVersion() *string
	SetIpamPoolDescription(v string) *CreateIpamPoolRequest
	GetIpamPoolDescription() *string
	SetIpamPoolName(v string) *CreateIpamPoolRequest
	GetIpamPoolName() *string
	SetIpamScopeId(v string) *CreateIpamPoolRequest
	GetIpamScopeId() *string
	SetIpv6Isp(v string) *CreateIpamPoolRequest
	GetIpv6Isp() *string
	SetOwnerAccount(v string) *CreateIpamPoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIpamPoolRequest
	GetOwnerId() *int64
	SetPoolRegionId(v string) *CreateIpamPoolRequest
	GetPoolRegionId() *string
	SetRegionId(v string) *CreateIpamPoolRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateIpamPoolRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateIpamPoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIpamPoolRequest
	GetResourceOwnerId() *int64
	SetSourceIpamPoolId(v string) *CreateIpamPoolRequest
	GetSourceIpamPoolId() *string
	SetTag(v []*CreateIpamPoolRequestTag) *CreateIpamPoolRequest
	GetTag() []*CreateIpamPoolRequestTag
}

type CreateIpamPoolRequest struct {
	// The default network mask assigned by the IPAM address pool.
	//
	// > The IPv4 network mask value range is 0 to 32 bits, and the IPv6 network mask value range is 0 to 128 bits.
	//
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// The maximum network mask assigned by the IPAM address pool.
	//
	// > The IPv4 network mask value range is **0 to 32*	- bits, and the IPv6 network mask value range is **0 to 128*	- bits.
	//
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// The minimum network mask assigned by the IPAM address pool.
	//
	// > The IPv4 network mask value range is **0 to 32*	- bits, and the IPv6 network mask value range is **0 to 128*	- bits.
	//
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// Whether the pool has the auto-import feature enabled.
	//
	// example:
	//
	// true
	AutoImport *bool `json:"AutoImport,omitempty" xml:"AutoImport,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// IP address protocol version. Values:
	//
	// - **IPv4**: IPv4 protocol.
	//
	// - **IPv6**: IPv6 protocol.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Description of the IPAM address pool.
	//
	// The length should be between 1 to 256 characters, and it must start with an uppercase or lowercase English letter or a Chinese character, but cannot begin with `http://` or `https://`. If left blank, the default value is empty.
	//
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// The name of the IPAM pool.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// The ID of the IPAM scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
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
	// The effective region of the IPAM pool.
	//
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the source IPAM pool.
	//
	// >  If you do not specify this parameter, the pool is a parent pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	SourceIpamPoolId *string `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
	// The tag list.
	Tag []*CreateIpamPoolRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateIpamPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolRequest) GetAllocationDefaultCidrMask() *int32 {
	return s.AllocationDefaultCidrMask
}

func (s *CreateIpamPoolRequest) GetAllocationMaxCidrMask() *int32 {
	return s.AllocationMaxCidrMask
}

func (s *CreateIpamPoolRequest) GetAllocationMinCidrMask() *int32 {
	return s.AllocationMinCidrMask
}

func (s *CreateIpamPoolRequest) GetAutoImport() *bool {
	return s.AutoImport
}

func (s *CreateIpamPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpamPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateIpamPoolRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateIpamPoolRequest) GetIpamPoolDescription() *string {
	return s.IpamPoolDescription
}

func (s *CreateIpamPoolRequest) GetIpamPoolName() *string {
	return s.IpamPoolName
}

func (s *CreateIpamPoolRequest) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *CreateIpamPoolRequest) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *CreateIpamPoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIpamPoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIpamPoolRequest) GetPoolRegionId() *string {
	return s.PoolRegionId
}

func (s *CreateIpamPoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpamPoolRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIpamPoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIpamPoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIpamPoolRequest) GetSourceIpamPoolId() *string {
	return s.SourceIpamPoolId
}

func (s *CreateIpamPoolRequest) GetTag() []*CreateIpamPoolRequestTag {
	return s.Tag
}

func (s *CreateIpamPoolRequest) SetAllocationDefaultCidrMask(v int32) *CreateIpamPoolRequest {
	s.AllocationDefaultCidrMask = &v
	return s
}

func (s *CreateIpamPoolRequest) SetAllocationMaxCidrMask(v int32) *CreateIpamPoolRequest {
	s.AllocationMaxCidrMask = &v
	return s
}

func (s *CreateIpamPoolRequest) SetAllocationMinCidrMask(v int32) *CreateIpamPoolRequest {
	s.AllocationMinCidrMask = &v
	return s
}

func (s *CreateIpamPoolRequest) SetAutoImport(v bool) *CreateIpamPoolRequest {
	s.AutoImport = &v
	return s
}

func (s *CreateIpamPoolRequest) SetClientToken(v string) *CreateIpamPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamPoolRequest) SetDryRun(v bool) *CreateIpamPoolRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpVersion(v string) *CreateIpamPoolRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpamPoolDescription(v string) *CreateIpamPoolRequest {
	s.IpamPoolDescription = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpamPoolName(v string) *CreateIpamPoolRequest {
	s.IpamPoolName = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpamScopeId(v string) *CreateIpamPoolRequest {
	s.IpamScopeId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetIpv6Isp(v string) *CreateIpamPoolRequest {
	s.Ipv6Isp = &v
	return s
}

func (s *CreateIpamPoolRequest) SetOwnerAccount(v string) *CreateIpamPoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpamPoolRequest) SetOwnerId(v int64) *CreateIpamPoolRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetPoolRegionId(v string) *CreateIpamPoolRequest {
	s.PoolRegionId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetRegionId(v string) *CreateIpamPoolRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetResourceGroupId(v string) *CreateIpamPoolRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetResourceOwnerAccount(v string) *CreateIpamPoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamPoolRequest) SetResourceOwnerId(v int64) *CreateIpamPoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetSourceIpamPoolId(v string) *CreateIpamPoolRequest {
	s.SourceIpamPoolId = &v
	return s
}

func (s *CreateIpamPoolRequest) SetTag(v []*CreateIpamPoolRequestTag) *CreateIpamPoolRequest {
	s.Tag = v
	return s
}

func (s *CreateIpamPoolRequest) Validate() error {
	return dara.Validate(s)
}

type CreateIpamPoolRequestTag struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with a `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpamPoolRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamPoolRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateIpamPoolRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateIpamPoolRequestTag) SetKey(v string) *CreateIpamPoolRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpamPoolRequestTag) SetValue(v string) *CreateIpamPoolRequestTag {
	s.Value = &v
	return s
}

func (s *CreateIpamPoolRequestTag) Validate() error {
	return dara.Validate(s)
}
