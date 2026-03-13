// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesSystemTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesSystemTagsRequest
	GetAll() *bool
	SetRegionId(v string) *UntagResourcesSystemTagsRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UntagResourcesSystemTagsRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagResourcesSystemTagsRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesSystemTagsRequest
	GetTagKey() []*string
	SetTagOwnerUid(v int64) *UntagResourcesSystemTagsRequest
	GetTagOwnerUid() *int64
}

type UntagResourcesSystemTagsRequest struct {
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	// This parameter is required.
	TagOwnerUid *int64 `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
}

func (s UntagResourcesSystemTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesSystemTagsRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesSystemTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesSystemTagsRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagResourcesSystemTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesSystemTagsRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UntagResourcesSystemTagsRequest) GetTagOwnerUid() *int64 {
	return s.TagOwnerUid
}

func (s *UntagResourcesSystemTagsRequest) SetAll(v bool) *UntagResourcesSystemTagsRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetRegionId(v string) *UntagResourcesSystemTagsRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetResourceId(v []*string) *UntagResourcesSystemTagsRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetResourceType(v string) *UntagResourcesSystemTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetTagKey(v []*string) *UntagResourcesSystemTagsRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetTagOwnerUid(v int64) *UntagResourcesSystemTagsRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) Validate() error {
	return dara.Validate(s)
}
