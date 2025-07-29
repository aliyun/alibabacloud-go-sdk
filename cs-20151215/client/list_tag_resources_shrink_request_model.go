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
	SetResourceIdsShrink(v string) *ListTagResourcesShrinkRequest
	GetResourceIdsShrink() *string
	SetResourceType(v string) *ListTagResourcesShrinkRequest
	GetResourceType() *string
	SetTagsShrink(v string) *ListTagResourcesShrinkRequest
	GetTagsShrink() *string
}

type ListTagResourcesShrinkRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ***
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The list of cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["xxxxx","xxxxxx"]
	ResourceIdsShrink *string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty"`
	// The resource type. Set the value to `CLUSTER`.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLUSTER
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The list of labels that you want to query. You can specify up to 20 labels.
	//
	// example:
	//
	// [{\\"key\\":\\"env\\",\\"value\\",\\"dev\\"},{\\"key\\":\\"dev\\", \\"value\\":\\"IT\\"}]
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
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

func (s *ListTagResourcesShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *ListTagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetRegionId(v string) *ListTagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdsShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagsShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
