// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesByTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagFilter(v *ListResourcesByTagRequestTagFilter) *ListResourcesByTagRequest
	GetTagFilter() *ListResourcesByTagRequestTagFilter
	SetFuzzyType(v string) *ListResourcesByTagRequest
	GetFuzzyType() *string
	SetIncludeAllTags(v bool) *ListResourcesByTagRequest
	GetIncludeAllTags() *bool
	SetMaxResult(v int32) *ListResourcesByTagRequest
	GetMaxResult() *int32
	SetNextToken(v string) *ListResourcesByTagRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListResourcesByTagRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListResourcesByTagRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListResourcesByTagRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListResourcesByTagRequest
	GetResourceOwnerAccount() *string
	SetResourceType(v string) *ListResourcesByTagRequest
	GetResourceType() *string
}

type ListResourcesByTagRequest struct {
	TagFilter *ListResourcesByTagRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
	// The type of the query. Valid values:
	//
	// 	- EQUAL: exact match for resources to which the specified tag is added. This is the default value.
	//
	// 	- NOT: exact match for resources to which the specified tag is not added.
	//
	// example:
	//
	// EQUAL
	FuzzyType *string `json:"FuzzyType,omitempty" xml:"FuzzyType,omitempty"`
	// Specifies whether to return the information of tags added to the resources. Valid values:
	//
	// 	- False: does not return the information of tags added to the resources. This is the default value.
	//
	// 	- True: returns the information of all tags added to the resources.
	//
	// example:
	//
	// False
	IncludeAllTags *bool `json:"IncludeAllTags,omitempty" xml:"IncludeAllTags,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 50. Maximum value: 1000.
	//
	// example:
	//
	// 50
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// For more information about region IDs, see [Endpoints](https://help.aliyun.com/document_detail/2330902.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. This parameter specifies a filter condition for the query.
	//
	// 	- If you set the FuzzyType parameter to EQUAL, you can set this parameter to a value obtained from the response of the [ListSupportResourceTypes](https://help.aliyun.com/document_detail/2330915.html) operation.
	//
	// 	- If you set the FuzzyType parameter to NOT, you can set this parameter to a resource type provided in **Types of resources that support queries based on the NOT operator**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::VPC::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourcesByTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByTagRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagRequest) GetTagFilter() *ListResourcesByTagRequestTagFilter {
	return s.TagFilter
}

func (s *ListResourcesByTagRequest) GetFuzzyType() *string {
	return s.FuzzyType
}

func (s *ListResourcesByTagRequest) GetIncludeAllTags() *bool {
	return s.IncludeAllTags
}

func (s *ListResourcesByTagRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *ListResourcesByTagRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourcesByTagRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListResourcesByTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListResourcesByTagRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourcesByTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListResourcesByTagRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourcesByTagRequest) SetTagFilter(v *ListResourcesByTagRequestTagFilter) *ListResourcesByTagRequest {
	s.TagFilter = v
	return s
}

func (s *ListResourcesByTagRequest) SetFuzzyType(v string) *ListResourcesByTagRequest {
	s.FuzzyType = &v
	return s
}

func (s *ListResourcesByTagRequest) SetIncludeAllTags(v bool) *ListResourcesByTagRequest {
	s.IncludeAllTags = &v
	return s
}

func (s *ListResourcesByTagRequest) SetMaxResult(v int32) *ListResourcesByTagRequest {
	s.MaxResult = &v
	return s
}

func (s *ListResourcesByTagRequest) SetNextToken(v string) *ListResourcesByTagRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourcesByTagRequest) SetOwnerAccount(v string) *ListResourcesByTagRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListResourcesByTagRequest) SetOwnerId(v int64) *ListResourcesByTagRequest {
	s.OwnerId = &v
	return s
}

func (s *ListResourcesByTagRequest) SetRegionId(v string) *ListResourcesByTagRequest {
	s.RegionId = &v
	return s
}

func (s *ListResourcesByTagRequest) SetResourceOwnerAccount(v string) *ListResourcesByTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListResourcesByTagRequest) SetResourceType(v string) *ListResourcesByTagRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesByTagRequest) Validate() error {
	return dara.Validate(s)
}

type ListResourcesByTagRequestTagFilter struct {
	// The tag key. This parameter specifies a filter condition for the query.
	//
	// The tag key can be a maximum of 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. This parameter specifies a filter condition for the query.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourcesByTagRequestTagFilter) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByTagRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagRequestTagFilter) GetKey() *string {
	return s.Key
}

func (s *ListResourcesByTagRequestTagFilter) GetValue() *string {
	return s.Value
}

func (s *ListResourcesByTagRequestTagFilter) SetKey(v string) *ListResourcesByTagRequestTagFilter {
	s.Key = &v
	return s
}

func (s *ListResourcesByTagRequestTagFilter) SetValue(v string) *ListResourcesByTagRequestTagFilter {
	s.Value = &v
	return s
}

func (s *ListResourcesByTagRequestTagFilter) Validate() error {
	return dara.Validate(s)
}
