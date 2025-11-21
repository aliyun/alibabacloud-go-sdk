// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UnTagResourcesShrinkRequest
	GetAll() *bool
	SetRegionId(v string) *UnTagResourcesShrinkRequest
	GetRegionId() *string
	SetResourceIdShrink(v string) *UnTagResourcesShrinkRequest
	GetResourceIdShrink() *string
	SetResourceType(v string) *UnTagResourcesShrinkRequest
	GetResourceType() *string
	SetTagKeyShrink(v string) *UnTagResourcesShrinkRequest
	GetTagKeyShrink() *string
}

type UnTagResourcesShrinkRequest struct {
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceIdShrink *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKeyShrink *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s UnTagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesShrinkRequest) GetAll() *bool {
	return s.All
}

func (s *UnTagResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnTagResourcesShrinkRequest) GetResourceIdShrink() *string {
	return s.ResourceIdShrink
}

func (s *UnTagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnTagResourcesShrinkRequest) GetTagKeyShrink() *string {
	return s.TagKeyShrink
}

func (s *UnTagResourcesShrinkRequest) SetAll(v bool) *UnTagResourcesShrinkRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesShrinkRequest) SetRegionId(v string) *UnTagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesShrinkRequest) SetResourceIdShrink(v string) *UnTagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *UnTagResourcesShrinkRequest) SetResourceType(v string) *UnTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesShrinkRequest) SetTagKeyShrink(v string) *UnTagResourcesShrinkRequest {
	s.TagKeyShrink = &v
	return s
}

func (s *UnTagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
