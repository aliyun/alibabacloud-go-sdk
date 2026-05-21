// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UntagResourcesRequest
	GetRegionId() *string
	SetCategory(v string) *UntagResourcesRequest
	GetCategory() *string
	SetResourceIds(v []*string) *UntagResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKeys(v []*string) *UntagResourcesRequest
	GetTagKeys() []*string
	SetTagOwnerUid(v int64) *UntagResourcesRequest
	GetTagOwnerUid() *int64
}

type UntagResourcesRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Custom
	Category    *string   `json:"category,omitempty" xml:"category,omitempty"`
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// INSTANCe
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	TagKeys      []*string `json:"tagKeys,omitempty" xml:"tagKeys,omitempty" type:"Repeated"`
	// example:
	//
	// 1062017779051424
	TagOwnerUid *int64 `json:"tagOwnerUid,omitempty" xml:"tagOwnerUid,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesRequest) GetCategory() *string {
	return s.Category
}

func (s *UntagResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *UntagResourcesRequest) GetTagOwnerUid() *int64 {
	return s.TagOwnerUid
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetCategory(v string) *UntagResourcesRequest {
	s.Category = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v []*string) *UntagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v []*string) *UntagResourcesRequest {
	s.TagKeys = v
	return s
}

func (s *UntagResourcesRequest) SetTagOwnerUid(v int64) *UntagResourcesRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
