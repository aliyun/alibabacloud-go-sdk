// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeRulesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRulesShrinkRequest
	GetNextToken() *string
	SetResourceCategoryId(v string) *DescribeRulesShrinkRequest
	GetResourceCategoryId() *string
	SetResourceOwnerIdsShrink(v string) *DescribeRulesShrinkRequest
	GetResourceOwnerIdsShrink() *string
	SetResourceRegionId(v string) *DescribeRulesShrinkRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *DescribeRulesShrinkRequest
	GetResourceType() *string
}

type DescribeRulesShrinkRequest struct {
	// The number of entries to return on each page. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. You can obtain this token from the `NextToken` parameter in the previous response.
	//
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource category ID.
	//
	// example:
	//
	// rc-000***123
	ResourceCategoryId     *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	ResourceOwnerIdsShrink *string `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty"`
	// The ID of the region where the resource resides.
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
}

func (s DescribeRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRulesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRulesShrinkRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeRulesShrinkRequest) GetResourceOwnerIdsShrink() *string {
	return s.ResourceOwnerIdsShrink
}

func (s *DescribeRulesShrinkRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeRulesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRulesShrinkRequest) SetMaxResults(v int32) *DescribeRulesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRulesShrinkRequest) SetNextToken(v string) *DescribeRulesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRulesShrinkRequest) SetResourceCategoryId(v string) *DescribeRulesShrinkRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeRulesShrinkRequest) SetResourceOwnerIdsShrink(v string) *DescribeRulesShrinkRequest {
	s.ResourceOwnerIdsShrink = &v
	return s
}

func (s *DescribeRulesShrinkRequest) SetResourceRegionId(v string) *DescribeRulesShrinkRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeRulesShrinkRequest) SetResourceType(v string) *DescribeRulesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
