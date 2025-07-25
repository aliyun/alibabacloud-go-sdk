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
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The IDs of IPAM pools. Valid values of N: 1 to 100. A maximum of 100 IPAM pools can be queried at a time.
	IpamPoolIds []*string `json:"IpamPoolIds,omitempty" xml:"IpamPoolIds,omitempty" type:"Repeated"`
	// The name of the IPAM pool. You can enter at most 20 names.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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
	Ipv6Isp     *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	// Whether it is a shared pool.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// The tag information.
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
	return dara.Validate(s)
}

type ListIpamPoolsRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. It can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It must start with a letter and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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
