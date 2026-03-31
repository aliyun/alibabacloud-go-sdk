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
	SetResourceNamesShrink(v string) *UntagResourcesShrinkRequest
	GetResourceNamesShrink() *string
	SetResourceType(v string) *UntagResourcesShrinkRequest
	GetResourceType() *string
	SetTagKeysShrink(v string) *UntagResourcesShrinkRequest
	GetTagKeysShrink() *string
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to remove all tags from the resources.
	//
	// Enumerated values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The names of the resources. You can specify up to 50 resource names.
	ResourceNamesShrink *string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty"`
	// The resource type.
	//
	// Enumerated values:
	//
	// 	- role: RAM roles.
	//
	// 	- policy: policies.
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags. You can specify up to 20 tag keys.
	TagKeysShrink *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
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

func (s *UntagResourcesShrinkRequest) GetResourceNamesShrink() *string {
	return s.ResourceNamesShrink
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

func (s *UntagResourcesShrinkRequest) SetResourceNamesShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceNamesShrink = &v
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
