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
	SetResourceIdShrink(v string) *UntagResourcesShrinkRequest
	GetResourceIdShrink() *string
	SetResourceType(v string) *UntagResourcesShrinkRequest
	GetResourceType() *string
	SetTagKeyShrink(v string) *UntagResourcesShrinkRequest
	GetTagKeyShrink() *string
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to remove all tags from the resources. This parameter and TagKey are mutually exclusive.
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The list of data catalog IDs from which to remove tags. The value is a JSON array. Maximum: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["clg-paimon-0424965be0c240acb4159688c9e2c4b6"]
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type. Valid value: CATALOGRESOURCE.
	//
	// This parameter is required.
	//
	// example:
	//
	// CATALOGRESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The list of tag keys to remove. The value is a JSON array. This parameter and All are mutually exclusive.
	//
	// example:
	//
	// ["team"]
	TagKeyShrink *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
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

func (s *UntagResourcesShrinkRequest) GetResourceIdShrink() *string {
	return s.ResourceIdShrink
}

func (s *UntagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesShrinkRequest) GetTagKeyShrink() *string {
	return s.TagKeyShrink
}

func (s *UntagResourcesShrinkRequest) SetAll(v bool) *UntagResourcesShrinkRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceIdShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceType(v string) *UntagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetTagKeyShrink(v string) *UntagResourcesShrinkRequest {
	s.TagKeyShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
