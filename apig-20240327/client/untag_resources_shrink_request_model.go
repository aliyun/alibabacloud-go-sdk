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
	// Specifies whether to delete all tags. This parameter is valid only when tags are empty. Valid values:
	//
	// - true
	//
	// - false
	//
	// - True
	//
	// - False
	//
	// Default value: false.
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The resource IDs. You can specify up to 50 subkeys.
	//
	// example:
	//
	// ["gw-xxx","gw-yyy"]
	ResourceIdShrink *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// gateway
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tag keys. You can specify up to 20 subkeys.
	//
	// example:
	//
	// ["env","team"]
	TagKeyShrink *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
