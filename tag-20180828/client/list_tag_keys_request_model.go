// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagFilter(v *ListTagKeysRequestTagFilter) *ListTagKeysRequest
	GetTagFilter() *ListTagKeysRequestTagFilter
	SetCategory(v string) *ListTagKeysRequest
	GetCategory() *string
	SetFuzzyType(v string) *ListTagKeysRequest
	GetFuzzyType() *string
	SetNextToken(v string) *ListTagKeysRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTagKeysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTagKeysRequest
	GetOwnerId() *int64
	SetPageSize(v int32) *ListTagKeysRequest
	GetPageSize() *int32
	SetQueryType(v string) *ListTagKeysRequest
	GetQueryType() *string
	SetRegionId(v string) *ListTagKeysRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTagKeysRequest
	GetResourceOwnerAccount() *string
	SetResourceType(v string) *ListTagKeysRequest
	GetResourceType() *string
}

type ListTagKeysRequest struct {
	TagFilter *ListTagKeysRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
	// The type of the resource tags. This parameter specifies a filter condition for the query. Valid values:
	//
	// 	- all (default value)
	//
	// 	- custom
	//
	// 	- system
	//
	// >  The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// all
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The type of the query. Valid values:
	//
	// 	- EQUAL (default): exact match
	//
	// 	- PREFIX: prefix-based fuzzy match
	//
	// example:
	//
	// EQUAL
	FuzzyType *string `json:"FuzzyType,omitempty" xml:"FuzzyType,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of tag keys to return on each page.
	//
	// Maximum value: 1000. Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The category of the tags. This parameter specifies a filter condition for the query. Valid values:
	//
	// 	- ResourceTag: resource tags, including custom and system tags. This is the default value.
	//
	// 	- MetaTag: preset tags.
	//
	// >  The value of this parameter is not case-sensitive.
	//
	// example:
	//
	// ResourceTag
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The region ID.
	//
	// For more information about region IDs, see [Endpoints](https://help.aliyun.com/document_detail/2330902.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. This parameter specifies a filter condition for the query.
	//
	// Format: `ALIYUN::${ProductCode}::${ResourceType}`. All letters in the value of this parameter must be in uppercase.
	//
	// 	- `ProductCode`: the service code. You can set this field to a value obtained from the response of the [ListSupportResourceTypes](https://help.aliyun.com/document_detail/2330915.html) operation.
	//
	// 	- `ResourceType`: the resource type. You can set this field to a value obtained from the response of the [ListSupportResourceTypes](https://help.aliyun.com/document_detail/2330915.html) operation.
	//
	// example:
	//
	// ALIYUN::ECS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) GetTagFilter() *ListTagKeysRequestTagFilter {
	return s.TagFilter
}

func (s *ListTagKeysRequest) GetCategory() *string {
	return s.Category
}

func (s *ListTagKeysRequest) GetFuzzyType() *string {
	return s.FuzzyType
}

func (s *ListTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTagKeysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTagKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagKeysRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListTagKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagKeysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagKeysRequest) SetTagFilter(v *ListTagKeysRequestTagFilter) *ListTagKeysRequest {
	s.TagFilter = v
	return s
}

func (s *ListTagKeysRequest) SetCategory(v string) *ListTagKeysRequest {
	s.Category = &v
	return s
}

func (s *ListTagKeysRequest) SetFuzzyType(v string) *ListTagKeysRequest {
	s.FuzzyType = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetOwnerAccount(v string) *ListTagKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetOwnerId(v int64) *ListTagKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetQueryType(v string) *ListTagKeysRequest {
	s.QueryType = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceOwnerAccount(v string) *ListTagKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) Validate() error {
	return dara.Validate(s)
}

type ListTagKeysRequestTagFilter struct {
	// The tag key for a fuzzy query.
	//
	// This parameter is used together with the `FuzzyType` parameter.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysRequestTagFilter) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequestTagFilter) GetKey() *string {
	return s.Key
}

func (s *ListTagKeysRequestTagFilter) SetKey(v string) *ListTagKeysRequestTagFilter {
	s.Key = &v
	return s
}

func (s *ListTagKeysRequestTagFilter) Validate() error {
	return dara.Validate(s)
}
