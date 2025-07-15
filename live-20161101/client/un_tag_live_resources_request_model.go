// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagLiveResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UnTagLiveResourcesRequest
	GetAll() *bool
	SetOwnerId(v int64) *UnTagLiveResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnTagLiveResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UnTagLiveResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UnTagLiveResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UnTagLiveResourcesRequest
	GetTagKey() []*string
}

type UnTagLiveResourcesRequest struct {
	// example:
	//
	// false
	All      *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// DOMAIN
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// env
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagLiveResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagLiveResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagLiveResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UnTagLiveResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnTagLiveResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnTagLiveResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UnTagLiveResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnTagLiveResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UnTagLiveResourcesRequest) SetAll(v bool) *UnTagLiveResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagLiveResourcesRequest) SetOwnerId(v int64) *UnTagLiveResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UnTagLiveResourcesRequest) SetRegionId(v string) *UnTagLiveResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagLiveResourcesRequest) SetResourceId(v []*string) *UnTagLiveResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagLiveResourcesRequest) SetResourceType(v string) *UnTagLiveResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagLiveResourcesRequest) SetTagKey(v []*string) *UnTagLiveResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UnTagLiveResourcesRequest) Validate() error {
	return dara.Validate(s)
}
