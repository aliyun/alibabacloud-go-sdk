// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagVodResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UnTagVodResourcesRequest
	GetAll() *bool
	SetOwnerId(v int64) *UnTagVodResourcesRequest
	GetOwnerId() *int64
	SetResourceId(v []*string) *UnTagVodResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UnTagVodResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UnTagVodResourcesRequest
	GetTagKey() []*string
}

type UnTagVodResourcesRequest struct {
	All     *bool  `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagVodResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagVodResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagVodResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UnTagVodResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnTagVodResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UnTagVodResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnTagVodResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UnTagVodResourcesRequest) SetAll(v bool) *UnTagVodResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagVodResourcesRequest) SetOwnerId(v int64) *UnTagVodResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UnTagVodResourcesRequest) SetResourceId(v []*string) *UnTagVodResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagVodResourcesRequest) SetResourceType(v string) *UnTagVodResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagVodResourcesRequest) SetTagKey(v []*string) *UnTagVodResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UnTagVodResourcesRequest) Validate() error {
	return dara.Validate(s)
}
