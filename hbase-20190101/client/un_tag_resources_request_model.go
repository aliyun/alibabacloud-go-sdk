// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UnTagResourcesRequest
	GetAll() *bool
	SetRegionId(v string) *UnTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UnTagResourcesRequest
	GetResourceId() []*string
	SetTagKey(v []*string) *UnTagResourcesRequest
	GetTagKey() []*string
}

type UnTagResourcesRequest struct {
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bds-bp15e022622fk0w1
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// example:
	//
	// key1
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UnTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UnTagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UnTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
