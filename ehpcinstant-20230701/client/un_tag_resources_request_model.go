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
	SetResourceId(v []*string) *UnTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UnTagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UnTagResourcesRequest
	GetTagKey() []*string
}

type UnTagResourcesRequest struct {
	// example:
	//
	// False
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Job
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UnTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UnTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnTagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UnTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
