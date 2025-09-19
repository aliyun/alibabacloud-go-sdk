// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListIpamPoolsResponseBody
	GetCount() *int64
	SetIpamPools(v []*ListIpamPoolsResponseBodyIpamPools) *ListIpamPoolsResponseBody
	GetIpamPools() []*ListIpamPoolsResponseBodyIpamPools
	SetMaxResults(v int64) *ListIpamPoolsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListIpamPoolsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamPoolsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamPoolsResponseBody
	GetTotalCount() *int64
}

type ListIpamPoolsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IPAM pools.
	IpamPools []*ListIpamPoolsResponseBodyIpamPools `json:"IpamPools,omitempty" xml:"IpamPools,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B54867DE-83DC-56B4-A57E-69A03119D0B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListIpamPoolsResponseBody) GetIpamPools() []*ListIpamPoolsResponseBodyIpamPools {
	return s.IpamPools
}

func (s *ListIpamPoolsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIpamPoolsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamPoolsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamPoolsResponseBody) SetCount(v int64) *ListIpamPoolsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamPoolsResponseBody) SetIpamPools(v []*ListIpamPoolsResponseBodyIpamPools) *ListIpamPoolsResponseBody {
	s.IpamPools = v
	return s
}

func (s *ListIpamPoolsResponseBody) SetMaxResults(v int64) *ListIpamPoolsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolsResponseBody) SetNextToken(v string) *ListIpamPoolsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolsResponseBody) SetRequestId(v string) *ListIpamPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamPoolsResponseBody) SetTotalCount(v int64) *ListIpamPoolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamPoolsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIpamPoolsResponseBodyIpamPools struct {
	// The default network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 28
	AllocationDefaultCidrMask *int32 `json:"AllocationDefaultCidrMask,omitempty" xml:"AllocationDefaultCidrMask,omitempty"`
	// The maximum network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
	//
	// example:
	//
	// 32
	AllocationMaxCidrMask *int32 `json:"AllocationMaxCidrMask,omitempty" xml:"AllocationMaxCidrMask,omitempty"`
	// The minimum network mask assigned to the IPAM pool.
	//
	// An IPv4 mask must be **0 to 32*	- bits in length.
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
	AutoImport *bool     `json:"AutoImport,omitempty" xml:"AutoImport,omitempty"`
	Cidrs      []*string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Repeated"`
	// The time when the IPAM pool was created.
	//
	// example:
	//
	// 2023-04-19T16:49:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the pool is a subpool. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasSubPool *bool `json:"HasSubPool,omitempty" xml:"HasSubPool,omitempty"`
	// The IP version. Only **IPv4*	- may be returned.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-b5mtlx3q7xcnyr****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The description of the IPAM pool.
	//
	// example:
	//
	// test description
	IpamPoolDescription *string `json:"IpamPoolDescription,omitempty" xml:"IpamPoolDescription,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The name of the IPAM pool.
	//
	// example:
	//
	// test
	IpamPoolName *string `json:"IpamPoolName,omitempty" xml:"IpamPoolName,omitempty"`
	// The ID of the region where the IPAM to which the IPAM pool belongs is hosted.
	//
	// example:
	//
	// cn-hangzhou
	IpamRegionId *string `json:"IpamRegionId,omitempty" xml:"IpamRegionId,omitempty"`
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The type of the IPAM scope. Valid values:
	//
	// 	- **public**
	//
	// 	- **private**
	//
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	Ipv6Isp       *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	// Whether it is a shared pool.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The Alibaba Cloud account of the owner for the IPAM pool.
	//
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The depth of the IPAM pool. Valid values: **0 to 10**.
	//
	// example:
	//
	// 2
	PoolDepth *int32 `json:"PoolDepth,omitempty" xml:"PoolDepth,omitempty"`
	// The effective region of the IPAM pool. The ID of the effective region for the IPAM pool.
	//
	// example:
	//
	// cn-hangzhou
	PoolRegionId *string `json:"PoolRegionId,omitempty" xml:"PoolRegionId,omitempty"`
	// The ID of the region where the operation is called.
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
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the source IPAM pool.
	//
	// example:
	//
	// ipam-pool-lfnwi4jok1ss0g****
	SourceIpamPoolId *string `json:"SourceIpamPoolId,omitempty" xml:"SourceIpamPoolId,omitempty"`
	// The status of the IPAM pool. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**: indicates that the creation is complete.
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListIpamPoolsResponseBodyIpamPoolsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamPoolsResponseBodyIpamPools) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolsResponseBodyIpamPools) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetAllocationDefaultCidrMask() *int32 {
	return s.AllocationDefaultCidrMask
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetAllocationMaxCidrMask() *int32 {
	return s.AllocationMaxCidrMask
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetAllocationMinCidrMask() *int32 {
	return s.AllocationMinCidrMask
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetAutoImport() *bool {
	return s.AutoImport
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetCidrs() []*string {
	return s.Cidrs
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetHasSubPool() *bool {
	return s.HasSubPool
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpamId() *string {
	return s.IpamId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpamPoolDescription() *string {
	return s.IpamPoolDescription
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpamPoolName() *string {
	return s.IpamPoolName
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpamRegionId() *string {
	return s.IpamRegionId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpamScopeType() *string {
	return s.IpamScopeType
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetIsShared() *bool {
	return s.IsShared
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetPoolDepth() *int32 {
	return s.PoolDepth
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetPoolRegionId() *string {
	return s.PoolRegionId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetSourceIpamPoolId() *string {
	return s.SourceIpamPoolId
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetStatus() *string {
	return s.Status
}

func (s *ListIpamPoolsResponseBodyIpamPools) GetTags() []*ListIpamPoolsResponseBodyIpamPoolsTags {
	return s.Tags
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetAllocationDefaultCidrMask(v int32) *ListIpamPoolsResponseBodyIpamPools {
	s.AllocationDefaultCidrMask = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetAllocationMaxCidrMask(v int32) *ListIpamPoolsResponseBodyIpamPools {
	s.AllocationMaxCidrMask = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetAllocationMinCidrMask(v int32) *ListIpamPoolsResponseBodyIpamPools {
	s.AllocationMinCidrMask = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetAutoImport(v bool) *ListIpamPoolsResponseBodyIpamPools {
	s.AutoImport = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetCidrs(v []*string) *ListIpamPoolsResponseBodyIpamPools {
	s.Cidrs = v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetCreateTime(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.CreateTime = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetHasSubPool(v bool) *ListIpamPoolsResponseBodyIpamPools {
	s.HasSubPool = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpVersion(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpVersion = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamPoolDescription(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamPoolDescription = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamPoolId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamPoolName(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamPoolName = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamRegionId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamRegionId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamScopeId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpamScopeType(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.IpamScopeType = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIpv6Isp(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.Ipv6Isp = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetIsShared(v bool) *ListIpamPoolsResponseBodyIpamPools {
	s.IsShared = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetOwnerId(v int64) *ListIpamPoolsResponseBodyIpamPools {
	s.OwnerId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetPoolDepth(v int32) *ListIpamPoolsResponseBodyIpamPools {
	s.PoolDepth = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetPoolRegionId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.PoolRegionId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetRegionId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.RegionId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetResourceGroupId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetSourceIpamPoolId(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.SourceIpamPoolId = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetStatus(v string) *ListIpamPoolsResponseBodyIpamPools {
	s.Status = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) SetTags(v []*ListIpamPoolsResponseBodyIpamPoolsTags) *ListIpamPoolsResponseBodyIpamPools {
	s.Tags = v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPools) Validate() error {
	return dara.Validate(s)
}

type ListIpamPoolsResponseBodyIpamPoolsTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamPoolsResponseBodyIpamPoolsTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolsResponseBodyIpamPoolsTags) GoString() string {
	return s.String()
}

func (s *ListIpamPoolsResponseBodyIpamPoolsTags) GetKey() *string {
	return s.Key
}

func (s *ListIpamPoolsResponseBodyIpamPoolsTags) GetValue() *string {
	return s.Value
}

func (s *ListIpamPoolsResponseBodyIpamPoolsTags) SetKey(v string) *ListIpamPoolsResponseBodyIpamPoolsTags {
	s.Key = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPoolsTags) SetValue(v string) *ListIpamPoolsResponseBodyIpamPoolsTags {
	s.Value = &v
	return s
}

func (s *ListIpamPoolsResponseBodyIpamPoolsTags) Validate() error {
	return dara.Validate(s)
}
