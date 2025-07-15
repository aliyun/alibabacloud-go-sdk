// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrefixListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListPrefixListsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListPrefixListsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListPrefixListsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListPrefixListsRequest
	GetOwnerId() *int64
	SetPrefixListIds(v []*string) *ListPrefixListsRequest
	GetPrefixListIds() []*string
	SetPrefixListName(v string) *ListPrefixListsRequest
	GetPrefixListName() *string
	SetRegionId(v string) *ListPrefixListsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListPrefixListsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListPrefixListsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListPrefixListsRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListPrefixListsRequestTags) *ListPrefixListsRequest
	GetTags() []*ListPrefixListsRequestTags
}

type ListPrefixListsRequest struct {
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
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
	// The IDs of prefix lists to be queried. Valid values of **N*	- are **1*	- to **100**, which specifies that you can query up to 100 prefix lists at a time.
	//
	// example:
	//
	// pl-m5estsqsdqwg88hjf****
	PrefixListIds []*string `json:"PrefixListIds,omitempty" xml:"PrefixListIds,omitempty" type:"Repeated"`
	// The name of the prefix list to query.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// name
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The ID of the region where you want to query prefix lists.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the prefix list belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tags []*ListPrefixListsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPrefixListsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrefixListsRequest) GoString() string {
	return s.String()
}

func (s *ListPrefixListsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListPrefixListsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrefixListsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListPrefixListsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPrefixListsRequest) GetPrefixListIds() []*string {
	return s.PrefixListIds
}

func (s *ListPrefixListsRequest) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *ListPrefixListsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrefixListsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrefixListsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPrefixListsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListPrefixListsRequest) GetTags() []*ListPrefixListsRequestTags {
	return s.Tags
}

func (s *ListPrefixListsRequest) SetMaxResults(v int64) *ListPrefixListsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPrefixListsRequest) SetNextToken(v string) *ListPrefixListsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPrefixListsRequest) SetOwnerAccount(v string) *ListPrefixListsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPrefixListsRequest) SetOwnerId(v int64) *ListPrefixListsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPrefixListsRequest) SetPrefixListIds(v []*string) *ListPrefixListsRequest {
	s.PrefixListIds = v
	return s
}

func (s *ListPrefixListsRequest) SetPrefixListName(v string) *ListPrefixListsRequest {
	s.PrefixListName = &v
	return s
}

func (s *ListPrefixListsRequest) SetRegionId(v string) *ListPrefixListsRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrefixListsRequest) SetResourceGroupId(v string) *ListPrefixListsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrefixListsRequest) SetResourceOwnerAccount(v string) *ListPrefixListsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPrefixListsRequest) SetResourceOwnerId(v int64) *ListPrefixListsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListPrefixListsRequest) SetTags(v []*ListPrefixListsRequestTags) *ListPrefixListsRequest {
	s.Tags = v
	return s
}

func (s *ListPrefixListsRequest) Validate() error {
	return dara.Validate(s)
}

type ListPrefixListsRequestTags struct {
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The key cannot exceed 64 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value cannot exceed 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrefixListsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListPrefixListsRequestTags) GoString() string {
	return s.String()
}

func (s *ListPrefixListsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListPrefixListsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListPrefixListsRequestTags) SetKey(v string) *ListPrefixListsRequestTags {
	s.Key = &v
	return s
}

func (s *ListPrefixListsRequestTags) SetValue(v string) *ListPrefixListsRequestTags {
	s.Value = &v
	return s
}

func (s *ListPrefixListsRequestTags) Validate() error {
	return dara.Validate(s)
}
