// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceNamesShrink(v string) *TagResourcesShrinkRequest
	GetResourceNamesShrink() *string
	SetResourceType(v string) *TagResourcesShrinkRequest
	GetResourceType() *string
	SetTagShrink(v string) *TagResourcesShrinkRequest
	GetTagShrink() *string
}

type TagResourcesShrinkRequest struct {
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
	// The tags. You can specify up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s TagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesShrinkRequest) GetResourceNamesShrink() *string {
	return s.ResourceNamesShrink
}

func (s *TagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *TagResourcesShrinkRequest) SetResourceNamesShrink(v string) *TagResourcesShrinkRequest {
	s.ResourceNamesShrink = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceType(v string) *TagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetTagShrink(v string) *TagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *TagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
