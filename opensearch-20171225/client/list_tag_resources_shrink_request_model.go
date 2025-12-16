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
	SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest
	GetResourceIdShrink() *string
	SetResourceType(v string) *ListTagResourcesShrinkRequest
	GetResourceType() *string
	SetTagShrink(v string) *ListTagResourcesShrinkRequest
	GetTagShrink() *string
}

type ListTagResourcesShrinkRequest struct {
	// The token that is used to retrieve the next page.
	//
	// example:
	//
	// 60
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 resource IDs.
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// BIGDATA
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags. You can specify a maximum number of 20 tags.
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
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

func (s *ListTagResourcesShrinkRequest) GetResourceIdShrink() *string {
	return s.ResourceIdShrink
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

func (s *ListTagResourcesShrinkRequest) SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
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
