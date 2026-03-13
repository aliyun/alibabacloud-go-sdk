// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesSystemTagsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesSystemTagsShrinkRequest
	GetAll() *bool
	SetRegionId(v string) *UntagResourcesSystemTagsShrinkRequest
	GetRegionId() *string
	SetResourceIdShrink(v string) *UntagResourcesSystemTagsShrinkRequest
	GetResourceIdShrink() *string
	SetResourceType(v string) *UntagResourcesSystemTagsShrinkRequest
	GetResourceType() *string
	SetTagKeyShrink(v string) *UntagResourcesSystemTagsShrinkRequest
	GetTagKeyShrink() *string
	SetTagOwnerUid(v int64) *UntagResourcesSystemTagsShrinkRequest
	GetTagOwnerUid() *int64
}

type UntagResourcesSystemTagsShrinkRequest struct {
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceIdShrink *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKeyShrink *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// This parameter is required.
	TagOwnerUid *int64 `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
}

func (s UntagResourcesSystemTagsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesSystemTagsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsShrinkRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesSystemTagsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesSystemTagsShrinkRequest) GetResourceIdShrink() *string {
	return s.ResourceIdShrink
}

func (s *UntagResourcesSystemTagsShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesSystemTagsShrinkRequest) GetTagKeyShrink() *string {
	return s.TagKeyShrink
}

func (s *UntagResourcesSystemTagsShrinkRequest) GetTagOwnerUid() *int64 {
	return s.TagOwnerUid
}

func (s *UntagResourcesSystemTagsShrinkRequest) SetAll(v bool) *UntagResourcesSystemTagsShrinkRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesSystemTagsShrinkRequest) SetRegionId(v string) *UntagResourcesSystemTagsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesSystemTagsShrinkRequest) SetResourceIdShrink(v string) *UntagResourcesSystemTagsShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *UntagResourcesSystemTagsShrinkRequest) SetResourceType(v string) *UntagResourcesSystemTagsShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesSystemTagsShrinkRequest) SetTagKeyShrink(v string) *UntagResourcesSystemTagsShrinkRequest {
	s.TagKeyShrink = &v
	return s
}

func (s *UntagResourcesSystemTagsShrinkRequest) SetTagOwnerUid(v int64) *UntagResourcesSystemTagsShrinkRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *UntagResourcesSystemTagsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
