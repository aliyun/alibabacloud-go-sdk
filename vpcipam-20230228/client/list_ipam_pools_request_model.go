// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpVersion(v string) *ListIpamPoolsRequest
	GetIpVersion() *string
	SetIpamPoolIds(v []*string) *ListIpamPoolsRequest
	GetIpamPoolIds() []*string
	SetIpamPoolName(v string) *ListIpamPoolsRequest
	GetIpamPoolName() *string
	SetIpamScopeId(v string) *ListIpamPoolsRequest
	GetIpamScopeId() *string
	SetIpv6Isp(v string) *ListIpamPoolsRequest
	GetIpv6Isp() *string
	SetIsShared(v bool) *ListIpamPoolsRequest
	GetIsShared() *bool
	SetMaxResults(v int32) *ListIpamPoolsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamPoolsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListIpamPoolsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListIpamPoolsRequest
	GetOwnerId() *int64
	SetPoolRegionId(v string) *ListIpamPoolsRequest
	GetPoolRegionId() *string
	SetRegionId(v string) *ListIpamPoolsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListIpamPoolsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListIpamPoolsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListIpamPoolsRequest
	GetResourceOwnerId() *int64
	SetSourceIpamPoolId(v string) *ListIpamPoolsRequest
	GetSourceIpamPoolId() *string
	SetTags(v []*ListIpamPoolsRequestTags) *ListIpamPoolsRequest
	GetTags() []*ListIpamPoolsRequestTags
}

type ListIpamPoolsRequest struct {
	// The IP version. Valid values:
	//
	// - **IPv4**
	//
	// - **IPv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// A list of IPAM pool IDs. You can specify up to 100 IDs.
	IpamPoolIds []*string `json:"IpamPoolIds,omitempty" xml:"IpamPoolIds,omitempty" type:"Repeated"`
	// The name of the IPAM pool.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The line type of the IPv6 CIDR block. This parameter is valid only for public IPv6 address pools. Valid values:
	//
	// - **BGP*	- (default): Alibaba Cloud BGP IPv6.
	//
	// - **ChinaMobile**
	//
	// - **ChinaUnicom**
	//
	// - **ChinaTelecom**
	//
	// > If your account is whitelisted for single-line bandwidth, you can set this parameter to **ChinaTelecom**, **ChinaUnicom**, or **ChinaMobile**.
	//
	// example:
	//
	// BGP
	Ipv6Isp *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	// Specifies whether the address pool is a shared pool.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Valid values:
	//
	// - You do not need to specify this parameter for the first call.
	//
	// - Set this parameter to the value of NextToken that was returned in the previous call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the IPAM pool is available.
	//
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// The ID of the managed region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPAM pool belongs.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the source IPAM pool.
	//
	// example:
	//
	// ipam-pool-lfnwi4jok1ss0g****
	SourceIpamPoolId *string `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
	// The tags.
	Tags []*ListIpamPoolsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamPoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListIpamPoolsRequest) GetIpamPoolIds() []*string {
	return s.IpamPoolIds
}

func (s *ListIpamPoolsRequest) GetIpamPoolName() *string {
	return s.IpamPoolName
}

func (s *ListIpamPoolsRequest) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *ListIpamPoolsRequest) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *ListIpamPoolsRequest) GetIsShared() *bool {
	return s.IsShared
}

func (s *ListIpamPoolsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamPoolsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamPoolsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListIpamPoolsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamPoolsRequest) GetPoolRegionId() *string {
	return s.PoolRegionId
}

func (s *ListIpamPoolsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamPoolsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpamPoolsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListIpamPoolsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamPoolsRequest) GetSourceIpamPoolId() *string {
	return s.SourceIpamPoolId
}

func (s *ListIpamPoolsRequest) GetTags() []*ListIpamPoolsRequestTags {
	return s.Tags
}

func (s *ListIpamPoolsRequest) SetIpVersion(v string) *ListIpamPoolsRequest {
	s.IpVersion = &v
	return s
}

func (s *ListIpamPoolsRequest) SetIpamPoolIds(v []*string) *ListIpamPoolsRequest {
	s.IpamPoolIds = v
	return s
}

func (s *ListIpamPoolsRequest) SetIpamPoolName(v string) *ListIpamPoolsRequest {
	s.IpamPoolName = &v
	return s
}

func (s *ListIpamPoolsRequest) SetIpamScopeId(v string) *ListIpamPoolsRequest {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetIpv6Isp(v string) *ListIpamPoolsRequest {
	s.Ipv6Isp = &v
	return s
}

func (s *ListIpamPoolsRequest) SetIsShared(v bool) *ListIpamPoolsRequest {
	s.IsShared = &v
	return s
}

func (s *ListIpamPoolsRequest) SetMaxResults(v int32) *ListIpamPoolsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolsRequest) SetNextToken(v string) *ListIpamPoolsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolsRequest) SetOwnerAccount(v string) *ListIpamPoolsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamPoolsRequest) SetOwnerId(v int64) *ListIpamPoolsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetPoolRegionId(v string) *ListIpamPoolsRequest {
	s.PoolRegionId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetRegionId(v string) *ListIpamPoolsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetResourceGroupId(v string) *ListIpamPoolsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetResourceOwnerAccount(v string) *ListIpamPoolsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamPoolsRequest) SetResourceOwnerId(v int64) *ListIpamPoolsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetSourceIpamPoolId(v string) *ListIpamPoolsRequest {
	s.SourceIpamPoolId = &v
	return s
}

func (s *ListIpamPoolsRequest) SetTags(v []*ListIpamPoolsRequestTags) *ListIpamPoolsRequest {
	s.Tags = v
	return s
}

func (s *ListIpamPoolsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIpamPoolsRequestTags struct {
	// The tag key. You can specify up to 20 tag keys. The key cannot be an empty string.
	//
	// The key can be up to 64 characters in length. It must start with a letter and can contain digits, periods (.), underscores (_), and hyphens (-). The key cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The value can be an empty string.
	//
	// The value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamPoolsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolsRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListIpamPoolsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListIpamPoolsRequestTags) SetKey(v string) *ListIpamPoolsRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpamPoolsRequestTags) SetValue(v string) *ListIpamPoolsRequestTags {
	s.Value = &v
	return s
}

func (s *ListIpamPoolsRequestTags) Validate() error {
	return dara.Validate(s)
}
