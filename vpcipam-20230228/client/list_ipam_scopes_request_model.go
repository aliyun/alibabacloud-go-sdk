// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamScopesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpamId(v string) *ListIpamScopesRequest
	GetIpamId() *string
	SetIpamScopeIds(v []*string) *ListIpamScopesRequest
	GetIpamScopeIds() []*string
	SetIpamScopeName(v string) *ListIpamScopesRequest
	GetIpamScopeName() *string
	SetIpamScopeType(v string) *ListIpamScopesRequest
	GetIpamScopeType() *string
	SetMaxResults(v int64) *ListIpamScopesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListIpamScopesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListIpamScopesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListIpamScopesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListIpamScopesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListIpamScopesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListIpamScopesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListIpamScopesRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListIpamScopesRequestTags) *ListIpamScopesRequest
	GetTags() []*ListIpamScopesRequestTags
}

type ListIpamScopesRequest struct {
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The IDs of IPAM scopes.
	IpamScopeIds []*string `json:"IpamScopeIds,omitempty" xml:"IpamScopeIds,omitempty" type:"Repeated"`
	// The name of the IPAM scope.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM scope.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tags []*ListIpamScopesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamScopesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamScopesRequest) GoString() string {
	return s.String()
}

func (s *ListIpamScopesRequest) GetIpamId() *string {
	return s.IpamId
}

func (s *ListIpamScopesRequest) GetIpamScopeIds() []*string {
	return s.IpamScopeIds
}

func (s *ListIpamScopesRequest) GetIpamScopeName() *string {
	return s.IpamScopeName
}

func (s *ListIpamScopesRequest) GetIpamScopeType() *string {
	return s.IpamScopeType
}

func (s *ListIpamScopesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIpamScopesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamScopesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListIpamScopesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamScopesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamScopesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpamScopesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListIpamScopesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamScopesRequest) GetTags() []*ListIpamScopesRequestTags {
	return s.Tags
}

func (s *ListIpamScopesRequest) SetIpamId(v string) *ListIpamScopesRequest {
	s.IpamId = &v
	return s
}

func (s *ListIpamScopesRequest) SetIpamScopeIds(v []*string) *ListIpamScopesRequest {
	s.IpamScopeIds = v
	return s
}

func (s *ListIpamScopesRequest) SetIpamScopeName(v string) *ListIpamScopesRequest {
	s.IpamScopeName = &v
	return s
}

func (s *ListIpamScopesRequest) SetIpamScopeType(v string) *ListIpamScopesRequest {
	s.IpamScopeType = &v
	return s
}

func (s *ListIpamScopesRequest) SetMaxResults(v int64) *ListIpamScopesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamScopesRequest) SetNextToken(v string) *ListIpamScopesRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamScopesRequest) SetOwnerAccount(v string) *ListIpamScopesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamScopesRequest) SetOwnerId(v int64) *ListIpamScopesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamScopesRequest) SetRegionId(v string) *ListIpamScopesRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamScopesRequest) SetResourceGroupId(v string) *ListIpamScopesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamScopesRequest) SetResourceOwnerAccount(v string) *ListIpamScopesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamScopesRequest) SetResourceOwnerId(v int64) *ListIpamScopesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamScopesRequest) SetTags(v []*ListIpamScopesRequestTags) *ListIpamScopesRequest {
	s.Tags = v
	return s
}

func (s *ListIpamScopesRequest) Validate() error {
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

type ListIpamScopesRequestTags struct {
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

func (s ListIpamScopesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpamScopesRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpamScopesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListIpamScopesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListIpamScopesRequestTags) SetKey(v string) *ListIpamScopesRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpamScopesRequestTags) SetValue(v string) *ListIpamScopesRequestTags {
	s.Value = &v
	return s
}

func (s *ListIpamScopesRequestTags) Validate() error {
	return dara.Validate(s)
}
