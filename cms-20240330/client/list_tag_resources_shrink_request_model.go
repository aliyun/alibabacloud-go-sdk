// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagResourcesShrinkRequest
	GetMaxResults() *int32
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxxxxxxx
	NextToken        *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Service
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	TagShrink    *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
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

func (s *ListTagResourcesShrinkRequest) SetMaxResults(v int32) *ListTagResourcesShrinkRequest {
	s.MaxResults = &v
	return s
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
