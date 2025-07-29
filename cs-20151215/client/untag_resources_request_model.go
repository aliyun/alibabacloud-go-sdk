// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesRequest
	GetAll() *bool
	SetRegionId(v string) *UntagResourcesRequest
	GetRegionId() *string
	SetResourceIds(v []*string) *UntagResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKeys(v []*string) *UntagResourcesRequest
	GetTagKeys() []*string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all custom labels. This parameter takes effect only when `tag_keys` is left empty. Valid values:
	//
	// 	- `true`: Remove all custom labels.
	//
	// 	- `false`: Do not remove all custom labels.
	//
	// example:
	//
	// true
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The region ID of the resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The list of resource IDs.
	//
	// This parameter is required.
	ResourceIds []*string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty" type:"Repeated"`
	// The type of resource. Set the value to `CLUSTER`.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLUSTER
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The list of keys of the labels that you want to remove.
	//
	// This parameter is required.
	TagKeys []*string `json:"tag_keys,omitempty" xml:"tag_keys,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
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

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
