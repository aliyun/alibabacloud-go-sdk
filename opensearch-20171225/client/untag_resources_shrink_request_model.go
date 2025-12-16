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
	// Specifies whether to remove all tags from the specified one or more resources. This parameter takes effect only if the tagKey parameter is not specified. Valid values: true and false. Default value: false.
	//
	// example:
	//
	// true
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 IDs.
	//
	// This parameter is required.
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProductVersion
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The keys of tags. You can specify a maximum number of 20 keys.
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
