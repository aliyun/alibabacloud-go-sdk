// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataRedundancyType(v string) *DescribeResourcesShrinkRequest
	GetDataRedundancyType() *string
	SetFailedRuleTemplate(v string) *DescribeResourcesShrinkRequest
	GetFailedRuleTemplate() *string
	SetMaxResults(v int32) *DescribeResourcesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeResourcesShrinkRequest
	GetNextToken() *string
	SetResourceArn(v string) *DescribeResourcesShrinkRequest
	GetResourceArn() *string
	SetResourceCategoryId(v string) *DescribeResourcesShrinkRequest
	GetResourceCategoryId() *string
	SetResourceId(v string) *DescribeResourcesShrinkRequest
	GetResourceId() *string
	SetResourceOwnerIdsShrink(v string) *DescribeResourcesShrinkRequest
	GetResourceOwnerIdsShrink() *string
	SetResourceRegionId(v string) *DescribeResourcesShrinkRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *DescribeResourcesShrinkRequest
	GetResourceType() *string
	SetSortBy(v string) *DescribeResourcesShrinkRequest
	GetSortBy() *string
	SetSortOrder(v string) *DescribeResourcesShrinkRequest
	GetSortOrder() *string
	SetStorageClass(v string) *DescribeResourcesShrinkRequest
	GetStorageClass() *string
}

type DescribeResourcesShrinkRequest struct {
	// The data redundancy type.
	//
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	// A filter for rules that failed the scoring.
	//
	// example:
	//
	// rule-000c***yc9
	FailedRuleTemplate *string `json:"FailedRuleTemplate,omitempty" xml:"FailedRuleTemplate,omitempty"`
	// The page size. Default: 10. Maximum: 100. Values less than 10 are set to 10, and values greater than 100 are set to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. The service returns a token if the response is truncated. To retrieve the next page of results, include this token in your next request. If no token is returned, all results have been retrieved.
	//
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The Resource ARN.
	//
	// example:
	//
	// acs:ecs:cn-hangzhou:123***7890:instance/i-123***7890
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The ID of the resource category.
	//
	// example:
	//
	// rc-000***123
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	// The resource ID. For example, for an instance, this is the instance ID.
	//
	// example:
	//
	// i-0003***110
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// A list of resource owner IDs. Use this parameter for cross-account scenarios. If you omit this parameter, the service returns data for the current account by default.
	//
	// example:
	//
	// [123***7890]
	ResourceOwnerIdsShrink *string `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty"`
	// The resource region ID.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The sort key.
	//
	// example:
	//
	// protectionScore
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order.
	//
	// example:
	//
	// ASC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The storage class.
	//
	// example:
	//
	// ARCHIVE
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s DescribeResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcesShrinkRequest) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *DescribeResourcesShrinkRequest) GetFailedRuleTemplate() *string {
	return s.FailedRuleTemplate
}

func (s *DescribeResourcesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeResourcesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourcesShrinkRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *DescribeResourcesShrinkRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeResourcesShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourcesShrinkRequest) GetResourceOwnerIdsShrink() *string {
	return s.ResourceOwnerIdsShrink
}

func (s *DescribeResourcesShrinkRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourcesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeResourcesShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeResourcesShrinkRequest) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeResourcesShrinkRequest) SetDataRedundancyType(v string) *DescribeResourcesShrinkRequest {
	s.DataRedundancyType = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetFailedRuleTemplate(v string) *DescribeResourcesShrinkRequest {
	s.FailedRuleTemplate = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetMaxResults(v int32) *DescribeResourcesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetNextToken(v string) *DescribeResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetResourceArn(v string) *DescribeResourcesShrinkRequest {
	s.ResourceArn = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetResourceCategoryId(v string) *DescribeResourcesShrinkRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetResourceId(v string) *DescribeResourcesShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetResourceOwnerIdsShrink(v string) *DescribeResourcesShrinkRequest {
	s.ResourceOwnerIdsShrink = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetResourceRegionId(v string) *DescribeResourcesShrinkRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetResourceType(v string) *DescribeResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetSortBy(v string) *DescribeResourcesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetSortOrder(v string) *DescribeResourcesShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) SetStorageClass(v string) *DescribeResourcesShrinkRequest {
	s.StorageClass = &v
	return s
}

func (s *DescribeResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
