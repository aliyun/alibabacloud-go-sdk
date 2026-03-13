// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListTagResourcesShrinkRequest
	GetCategory() *string
	SetNextToken(v string) *ListTagResourcesShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagResourcesShrinkRequest
	GetRegionId() *string
	SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest
	GetResourceIdShrink() *string
	SetResourceType(v string) *ListTagResourcesShrinkRequest
	GetResourceType() *string
	SetScope(v string) *ListTagResourcesShrinkRequest
	GetScope() *string
	SetTagShrink(v string) *ListTagResourcesShrinkRequest
	GetTagShrink() *string
	SetTagOwnerUid(v int64) *ListTagResourcesShrinkRequest
	GetTagOwnerUid() *int64
}

type ListTagResourcesShrinkRequest struct {
	Category  *string `json:"Category,omitempty" xml:"Category,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIdShrink *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope        *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagShrink    *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// This parameter is required.
	TagOwnerUid *int64 `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ListTagResourcesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesShrinkRequest) GetResourceIdShrink() *string {
	return s.ResourceIdShrink
}

func (s *ListTagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *ListTagResourcesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListTagResourcesShrinkRequest) GetTagOwnerUid() *int64 {
	return s.TagOwnerUid
}

func (s *ListTagResourcesShrinkRequest) SetCategory(v string) *ListTagResourcesShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetRegionId(v string) *ListTagResourcesShrinkRequest {
	s.RegionId = &v
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

func (s *ListTagResourcesShrinkRequest) SetScope(v string) *ListTagResourcesShrinkRequest {
	s.Scope = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagOwnerUid(v int64) *ListTagResourcesShrinkRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
