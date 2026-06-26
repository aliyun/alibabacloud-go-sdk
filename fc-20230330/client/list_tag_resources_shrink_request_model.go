// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListTagResourcesShrinkRequest
	GetLimit() *int32
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
	// The number of resources to return.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of resource IDs.
	ResourceIdShrink *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::FC::FUNCTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags.
	//
	// You can specify up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) GetLimit() *int32 {
	return s.Limit
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

func (s *ListTagResourcesShrinkRequest) SetLimit(v int32) *ListTagResourcesShrinkRequest {
	s.Limit = &v
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
