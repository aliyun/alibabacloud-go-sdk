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
	// The default CIDR mask for allocations from the IPAM address pool.
	//
	// > The IPv4 CIDR mask ranges from **0*	- to **32*	- bits. The IPv6 CIDR mask ranges from **0*	- to **128*	- bits.
	//
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// The maximum CIDR mask for allocations from the IPAM address pool.
	//
	// > The IPv4 CIDR mask ranges from **0*	- to **32*	- bits. The IPv6 CIDR mask ranges from **0*	- to **128*	- bits.
	//
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// The minimum CIDR mask for allocations from the IPAM address pool.
	//
	// > The IPv4 CIDR mask ranges from **0*	- to **32*	- bits. The IPv6 CIDR mask ranges from **0*	- to **128*	- bits.
	//
	// example:
	//
	// 8
	AllocationMinCidrMask *int32 `json:"AllocationMinCidrMask,omitempty" xml:"AllocationMinCidrMask,omitempty"`
	// Specifies whether to enable the auto-import feature for the address pool.
	//
	// example:
	//
	// true
	AutoImport *bool `json:"AutoImport,omitempty" xml:"AutoImport,omitempty"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing required parameters, request format, and instance status. If the request fails the dry run, an error message is returned. If the request passes the dry run, DryRunOperation is returned.
	//
	// - **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is directly performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP address protocol version. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **IPv6**: IPv6.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The description of the IPAM address pool.
	//
	// The description must be 1 to 256 characters in length and must start with an uppercase or lowercase English letter or a Chinese character. It cannot start with `http://` or `https://`. If this parameter is not specified, the description is empty by default.
	//
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// The name of the IPAM address pool.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// The instance ID of the IPAM scope.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The type of IPv6 CIDR block. This parameter takes effect only for public IPv6 address pools. Valid values:
	//
	// - **BGP*	- (default): Alibaba Cloud BGP IPv6.
	//
	// - **ChinaMobile**: China Mobile (single-line).
	//
	// - **ChinaUnicom**: China Unicom (single-line).
	//
	// - **ChinaTelecom**: China Telecom (single-line).
	//
	// > If you are a user who has activated the single-line bandwidth whitelist, this field can be set to **ChinaTelecom*	- (China Telecom), **ChinaUnicom*	- (China Unicom), or **ChinaMobile*	- (China Mobile).
	//
	// example:
	//
	// BGP
	Ipv6Isp      *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the IPAM address pool takes effect.
	//
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// The region ID of the IPAM. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) API to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM address pool.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The instance ID of the source IPAM address pool.
	//
	// > If this parameter is not specified, the created address pool is a parent address pool.
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateIpamPoolRequestTag struct {
	// The tag key of the resource. A maximum of 20 tag keys are supported. Once this value is specified, it cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and must start with a letter or Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// > You must specify at least one of the **ResourceId.N*	- and **Tag.N*	- (**Tag.N.Key*	- and **Tag.N.Value**) parameters.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. A maximum of 20 tag values are supported. Once this value is specified, it can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// > You must specify at least one of the **ResourceId.N*	- and **Tag.N*	- (**Tag.N.Key*	- and **Tag.N.Value**) parameters.
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
