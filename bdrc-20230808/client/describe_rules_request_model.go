// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRulesRequest
	GetNextToken() *string
	SetResourceCategoryId(v string) *DescribeRulesRequest
	GetResourceCategoryId() *string
	SetResourceOwnerIds(v []*int64) *DescribeRulesRequest
	GetResourceOwnerIds() []*int64
	SetResourceRegionId(v string) *DescribeRulesRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *DescribeRulesRequest
	GetResourceType() *string
}

type DescribeRulesRequest struct {
	// The maximum number of entries per page for a paged query. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether a next query token exists. Valid values: If NextToken is empty, no more results exist. If NextToken is returned, the value indicates the token for the next query.
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
	ResourceCategoryId *string  `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
	ResourceOwnerIds   []*int64 `json:"ResourceOwnerIds,omitempty" xml:"ResourceOwnerIds,omitempty" type:"Repeated"`
	// The region ID of the resource.
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

func (s DescribeRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRulesRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *DescribeRulesRequest) GetResourceOwnerIds() []*int64 {
	return s.ResourceOwnerIds
}

func (s *DescribeRulesRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeRulesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRulesRequest) SetMaxResults(v int32) *DescribeRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRulesRequest) SetNextToken(v string) *DescribeRulesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceCategoryId(v string) *DescribeRulesRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceOwnerIds(v []*int64) *DescribeRulesRequest {
	s.ResourceOwnerIds = v
	return s
}

func (s *DescribeRulesRequest) SetResourceRegionId(v string) *DescribeRulesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceType(v string) *DescribeRulesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRulesRequest) Validate() error {
	return dara.Validate(s)
}
