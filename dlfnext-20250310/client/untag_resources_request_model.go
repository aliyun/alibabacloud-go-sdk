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
	SetResourceId(v []*string) *UntagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resources. This parameter and TagKey are mutually exclusive.
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The list of data catalog IDs from which to remove tags. The value is a JSON array. Maximum: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["clg-paimon-0424965be0c240acb4159688c9e2c4b6"]
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The resource type. Valid value: CATALOGRESOURCE.
	//
	// This parameter is required.
	//
	// example:
	//
	// CATALOGRESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The list of tag keys to remove. The value is a JSON array. This parameter and All are mutually exclusive.
	//
	// example:
	//
	// ["team"]
	TagKey []*string `json:"tagKey,omitempty" xml:"tagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
