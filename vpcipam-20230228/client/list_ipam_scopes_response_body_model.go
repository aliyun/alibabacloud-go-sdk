// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamScopesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListIpamScopesResponseBody
	GetCount() *int64
	SetIpamScopes(v []*ListIpamScopesResponseBodyIpamScopes) *ListIpamScopesResponseBody
	GetIpamScopes() []*ListIpamScopesResponseBodyIpamScopes
	SetMaxResults(v int64) *ListIpamScopesResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListIpamScopesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamScopesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamScopesResponseBody
	GetTotalCount() *int64
}

type ListIpamScopesResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IPAM scopes.
	IpamScopes []*ListIpamScopesResponseBodyIpamScopes `json:"IpamScopes,omitempty" xml:"IpamScopes,omitempty" type:"Repeated"`
	// The number of entries per page.
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
	// 8859C501-97E7-53D4-B94B-2A9E16003B22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamScopesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamScopesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamScopesResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListIpamScopesResponseBody) GetIpamScopes() []*ListIpamScopesResponseBodyIpamScopes {
	return s.IpamScopes
}

func (s *ListIpamScopesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIpamScopesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamScopesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamScopesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamScopesResponseBody) SetCount(v int64) *ListIpamScopesResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamScopesResponseBody) SetIpamScopes(v []*ListIpamScopesResponseBodyIpamScopes) *ListIpamScopesResponseBody {
	s.IpamScopes = v
	return s
}

func (s *ListIpamScopesResponseBody) SetMaxResults(v int64) *ListIpamScopesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamScopesResponseBody) SetNextToken(v string) *ListIpamScopesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamScopesResponseBody) SetRequestId(v string) *ListIpamScopesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamScopesResponseBody) SetTotalCount(v int64) *ListIpamScopesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamScopesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIpamScopesResponseBodyIpamScopes struct {
	// The time when the IPAM scope was created.
	//
	// example:
	//
	// 2022-04-18T03:12:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The description of the IPAM scope.
	//
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// The ID of the IPAM scope.
	//
	// example:
	//
	// ipam-scope-glfmcyldpm8lsy****
	IpamScopeId *string `json:"IpamScopeId,omitempty" xml:"IpamScopeId,omitempty"`
	// The name of the IPAM scope.
	//
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
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
	// Indicates whether the scope is the default scope. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The Alibaba Cloud account that owns the IPAM scope.
	//
	// example:
	//
	// 1210123456******
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of pools in the IPAM scope.
	//
	// example:
	//
	// 2
	PoolCount *int32 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	// The region ID of the IPAM.
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
	// The status of the IPAM scope. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
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
	Tags []*ListIpamScopesResponseBodyIpamScopesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamScopesResponseBodyIpamScopes) String() string {
	return dara.Prettify(s)
}

func (s ListIpamScopesResponseBodyIpamScopes) GoString() string {
	return s.String()
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetIpamId() *string {
	return s.IpamId
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetIpamScopeDescription() *string {
	return s.IpamScopeDescription
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetIpamScopeId() *string {
	return s.IpamScopeId
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetIpamScopeName() *string {
	return s.IpamScopeName
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetIpamScopeType() *string {
	return s.IpamScopeType
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetPoolCount() *int32 {
	return s.PoolCount
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetStatus() *string {
	return s.Status
}

func (s *ListIpamScopesResponseBodyIpamScopes) GetTags() []*ListIpamScopesResponseBodyIpamScopesTags {
	return s.Tags
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetCreateTime(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.CreateTime = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamId(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamScopeDescription(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamScopeDescription = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamScopeId(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamScopeId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamScopeName(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamScopeName = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIpamScopeType(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.IpamScopeType = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetIsDefault(v bool) *ListIpamScopesResponseBodyIpamScopes {
	s.IsDefault = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetOwnerId(v int64) *ListIpamScopesResponseBodyIpamScopes {
	s.OwnerId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetPoolCount(v int32) *ListIpamScopesResponseBodyIpamScopes {
	s.PoolCount = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetRegionId(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.RegionId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetResourceGroupId(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetStatus(v string) *ListIpamScopesResponseBodyIpamScopes {
	s.Status = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) SetTags(v []*ListIpamScopesResponseBodyIpamScopesTags) *ListIpamScopesResponseBodyIpamScopes {
	s.Tags = v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopes) Validate() error {
	return dara.Validate(s)
}

type ListIpamScopesResponseBodyIpamScopesTags struct {
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
	// FinanceDept
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamScopesResponseBodyIpamScopesTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpamScopesResponseBodyIpamScopesTags) GoString() string {
	return s.String()
}

func (s *ListIpamScopesResponseBodyIpamScopesTags) GetKey() *string {
	return s.Key
}

func (s *ListIpamScopesResponseBodyIpamScopesTags) GetValue() *string {
	return s.Value
}

func (s *ListIpamScopesResponseBodyIpamScopesTags) SetKey(v string) *ListIpamScopesResponseBodyIpamScopesTags {
	s.Key = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopesTags) SetValue(v string) *ListIpamScopesResponseBodyIpamScopesTags {
	s.Value = &v
	return s
}

func (s *ListIpamScopesResponseBodyIpamScopesTags) Validate() error {
	return dara.Validate(s)
}
