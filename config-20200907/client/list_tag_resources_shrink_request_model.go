// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagResourcesShrinkRequest
	GetRegionId() *string
	SetResourceId(v []*string) *ListTagResourcesShrinkRequest
	GetResourceId() []*string
	SetResourceType(v string) *ListTagResourcesShrinkRequest
	GetResourceType() *string
	SetTagShrink(v string) *ListTagResourcesShrinkRequest
	GetTagShrink() *string
}

type ListTagResourcesShrinkRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// You can add tags to up to 50 resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- `ACS::Config::Rule`
	//
	// 	- `ACS::Config::AggregateConfigRule`
	//
	// 	- `ACS::Config::Aggregator`
	//
	// 	- `ACS::Config::CompliancePack`
	//
	// 	- `ACS::Config::AggregateCompliancePack`
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::Config:rule
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesShrinkRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListTagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetRegionId(v string) *ListTagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceId(v []*string) *ListTagResourcesShrinkRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
