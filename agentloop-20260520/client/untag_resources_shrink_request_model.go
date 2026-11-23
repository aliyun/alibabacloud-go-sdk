// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v string) *UntagResourcesShrinkRequest
	GetAll() *string
	SetResourceIdShrink(v string) *UntagResourcesShrinkRequest
	GetResourceIdShrink() *string
	SetResourceType(v string) *UntagResourcesShrinkRequest
	GetResourceType() *string
	SetTagKeyShrink(v string) *UntagResourcesShrinkRequest
	GetTagKeyShrink() *string
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to remove all custom tags. Valid values: true, false, and an empty string. If this parameter is not specified or is set to an empty string, the value is treated as false. The value true takes effect only when the tagKey parameter is not specified or is an empty array.
	//
	// example:
	//
	// true
	All *string `json:"all,omitempty" xml:"all,omitempty"`
	// The list of resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["space-a","space-b"]
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type. Set the value to agentspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// agentspace
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The list of tag keys to remove. You can specify up to 20 tag keys. The tag keys must be unique. If this parameter is not empty, it takes precedence over the all parameter.
	//
	// example:
	//
	// ["key1", "key2", "key3"]
	TagKeyShrink *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
}

func (s UntagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesShrinkRequest) GetAll() *string {
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

func (s *UntagResourcesShrinkRequest) SetAll(v string) *UntagResourcesShrinkRequest {
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
