// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckResourceCountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *PrecheckResourceCountShrinkRequest
	GetResourceType() *string
	SetTagResourceMatchersShrink(v string) *PrecheckResourceCountShrinkRequest
	GetTagResourceMatchersShrink() *string
}

type PrecheckResourceCountShrinkRequest struct {
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	TagResourceMatchersShrink *string `json:"TagResourceMatchers,omitempty" xml:"TagResourceMatchers,omitempty"`
}

func (s PrecheckResourceCountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PrecheckResourceCountShrinkRequest) GoString() string {
	return s.String()
}

func (s *PrecheckResourceCountShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *PrecheckResourceCountShrinkRequest) GetTagResourceMatchersShrink() *string {
	return s.TagResourceMatchersShrink
}

func (s *PrecheckResourceCountShrinkRequest) SetResourceType(v string) *PrecheckResourceCountShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *PrecheckResourceCountShrinkRequest) SetTagResourceMatchersShrink(v string) *PrecheckResourceCountShrinkRequest {
	s.TagResourceMatchersShrink = &v
	return s
}

func (s *PrecheckResourceCountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
