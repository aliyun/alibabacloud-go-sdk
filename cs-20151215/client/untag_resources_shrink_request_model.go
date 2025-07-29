// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesShrinkRequest
	GetAll() *bool
	SetRegionId(v string) *UntagResourcesShrinkRequest
	GetRegionId() *string
	SetResourceIdsShrink(v string) *UntagResourcesShrinkRequest
	GetResourceIdsShrink() *string
	SetResourceType(v string) *UntagResourcesShrinkRequest
	GetResourceType() *string
	SetTagKeysShrink(v string) *UntagResourcesShrinkRequest
	GetTagKeysShrink() *string
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to remove all custom labels. This parameter takes effect only when `tag_keys` is left empty. Valid values:
	//
	// 	- `true`: Remove all custom labels.
	//
	// 	- `false`: Do not remove all custom labels.
	//
	// example:
	//
	// true
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The region ID of the resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The list of resource IDs.
	//
	// This parameter is required.
	ResourceIdsShrink *string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty"`
	// The type of resource. Set the value to `CLUSTER`.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLUSTER
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The list of keys of the labels that you want to remove.
	//
	// This parameter is required.
	TagKeysShrink *string `json:"tag_keys,omitempty" xml:"tag_keys,omitempty"`
}

func (s UntagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesShrinkRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *UntagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesShrinkRequest) GetTagKeysShrink() *string {
	return s.TagKeysShrink
}

func (s *UntagResourcesShrinkRequest) SetAll(v bool) *UntagResourcesShrinkRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetRegionId(v string) *UntagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceIdsShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceType(v string) *UntagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetTagKeysShrink(v string) *UntagResourcesShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
