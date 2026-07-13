// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataRedundancyType(v string) *DescribeResourcesRequest
	GetDataRedundancyType() *string
	SetFailedRuleTemplate(v string) *DescribeResourcesRequest
	GetFailedRuleTemplate() *string
	SetMaxResults(v int32) *DescribeResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeResourcesRequest
	GetNextToken() *string
	SetResourceArn(v string) *DescribeResourcesRequest
	GetResourceArn() *string
	SetResourceCategoryId(v string) *DescribeResourcesRequest
	GetResourceCategoryId() *string
	SetResourceId(v string) *DescribeResourcesRequest
	GetResourceId() *string
	SetResourceOwnerIds(v []*int64) *DescribeResourcesRequest
	GetResourceOwnerIds() []*int64
	SetResourceRegionId(v string) *DescribeResourcesRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *DescribeResourcesRequest
	GetResourceType() *string
	SetSortBy(v string) *DescribeResourcesRequest
	GetSortBy() *string
	SetSortOrder(v string) *DescribeResourcesRequest
	GetSortOrder() *string
	SetStorageClass(v string) *DescribeResourcesRequest
	GetStorageClass() *string
}

type DescribeResourcesRequest struct {
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
	ResourceOwnerIds []*int64 `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty" type:"Repeated"`
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

func (s DescribeResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcesRequest) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *DescribeResourcesRequest) GetFailedRuleTemplate() *string {
	return s.FailedRuleTemplate
}

func (s *DescribeResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourcesRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *DescribeResourcesRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourcesRequest) GetResourceOwnerIds() []*int64 {
	return s.ResourceOwnerIds
}

func (s *DescribeResourcesRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourcesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeResourcesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeResourcesRequest) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeResourcesRequest) SetDataRedundancyType(v string) *DescribeResourcesRequest {
	s.DataRedundancyType = &v
	return s
}

func (s *DescribeResourcesRequest) SetFailedRuleTemplate(v string) *DescribeResourcesRequest {
	s.FailedRuleTemplate = &v
	return s
}

func (s *DescribeResourcesRequest) SetMaxResults(v int32) *DescribeResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourcesRequest) SetNextToken(v string) *DescribeResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeResourcesRequest) SetResourceArn(v string) *DescribeResourcesRequest {
	s.ResourceArn = &v
	return s
}

func (s *DescribeResourcesRequest) SetResourceCategoryId(v string) *DescribeResourcesRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeResourcesRequest) SetResourceId(v string) *DescribeResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourcesRequest) SetResourceOwnerIds(v []*int64) *DescribeResourcesRequest {
	s.ResourceOwnerIds = v
	return s
}

func (s *DescribeResourcesRequest) SetResourceRegionId(v string) *DescribeResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeResourcesRequest) SetResourceType(v string) *DescribeResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourcesRequest) SetSortBy(v string) *DescribeResourcesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeResourcesRequest) SetSortOrder(v string) *DescribeResourcesRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeResourcesRequest) SetStorageClass(v string) *DescribeResourcesRequest {
	s.StorageClass = &v
	return s
}

func (s *DescribeResourcesRequest) Validate() error {
	return dara.Validate(s)
}
