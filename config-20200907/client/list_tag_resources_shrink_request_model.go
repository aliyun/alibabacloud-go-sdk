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
	// The token to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region to which the tags belong.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// You can query tags for a maximum of 50 resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type in CloudConfig. Valid values:
	//
	// - `ACS::Config::Rule`: a rule for a single account.
	//
	// - `ACS::Config::AggregateConfigRule`: a rule for multiple accounts.
	//
	// - `ACS::Config::Aggregator`: an account group.
	//
	// - `ACS::Config::CompliancePack`: a compliance package for a single account.
	//
	// - `ACS::Config::AggregateCompliancePack`: a compliance package for multiple accounts.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::Config::Rule
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	//
	// You can filter resources by a maximum of 20 tags.
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
