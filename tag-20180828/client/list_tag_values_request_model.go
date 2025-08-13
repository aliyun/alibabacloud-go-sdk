// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagFilter(v *ListTagValuesRequestTagFilter) *ListTagValuesRequest
	GetTagFilter() *ListTagValuesRequestTagFilter
	SetFuzzyType(v string) *ListTagValuesRequest
	GetFuzzyType() *string
	SetKey(v string) *ListTagValuesRequest
	GetKey() *string
	SetNextToken(v string) *ListTagValuesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTagValuesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTagValuesRequest
	GetOwnerId() *int64
	SetPageSize(v int32) *ListTagValuesRequest
	GetPageSize() *int32
	SetQueryType(v string) *ListTagValuesRequest
	GetQueryType() *string
	SetRegionId(v string) *ListTagValuesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTagValuesRequest
	GetResourceOwnerAccount() *string
	SetResourceType(v string) *ListTagValuesRequest
	GetResourceType() *string
}

type ListTagValuesRequest struct {
	TagFilter *ListTagValuesRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
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
	// The tag key. This parameter specifies a filter condition for the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of tag values to return on each page.
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

func (s ListTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) GetTagFilter() *ListTagValuesRequestTagFilter {
	return s.TagFilter
}

func (s *ListTagValuesRequest) GetFuzzyType() *string {
	return s.FuzzyType
}

func (s *ListTagValuesRequest) GetKey() *string {
	return s.Key
}

func (s *ListTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTagValuesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTagValuesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagValuesRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListTagValuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagValuesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTagValuesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagValuesRequest) SetTagFilter(v *ListTagValuesRequestTagFilter) *ListTagValuesRequest {
	s.TagFilter = v
	return s
}

func (s *ListTagValuesRequest) SetFuzzyType(v string) *ListTagValuesRequest {
	s.FuzzyType = &v
	return s
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerAccount(v string) *ListTagValuesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerId(v int64) *ListTagValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagValuesRequest) SetPageSize(v int32) *ListTagValuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesRequest) SetQueryType(v string) *ListTagValuesRequest {
	s.QueryType = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceOwnerAccount(v string) *ListTagValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) Validate() error {
	return dara.Validate(s)
}

type ListTagValuesRequestTagFilter struct {
	// The tag value for a fuzzy query.
	//
	// This parameter is used together with the `FuzzyType` parameter.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagValuesRequestTagFilter) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequestTagFilter) GetValue() *string {
	return s.Value
}

func (s *ListTagValuesRequestTagFilter) SetValue(v string) *ListTagValuesRequestTagFilter {
	s.Value = &v
	return s
}

func (s *ListTagValuesRequestTagFilter) Validate() error {
	return dara.Validate(s)
}
