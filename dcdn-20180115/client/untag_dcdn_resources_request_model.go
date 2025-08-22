// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagDcdnResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagDcdnResourcesRequest
	GetAll() *bool
	SetResourceId(v []*string) *UntagDcdnResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagDcdnResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagDcdnResourcesRequest
	GetTagKey() []*string
}

type UntagDcdnResourcesRequest struct {
	// Specifies whether to delete all tags. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The ID of the resource. Valid values of N: **1*	- to **50**.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to **DOMAIN**.
	//
	// This parameter is required.
	//
	// example:
	//
	// DOMAIN
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag. Valid values of N: **1*	- to **20**.
	//
	// example:
	//
	// env
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagDcdnResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagDcdnResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagDcdnResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UntagDcdnResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagDcdnResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagDcdnResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UntagDcdnResourcesRequest) SetAll(v bool) *UntagDcdnResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagDcdnResourcesRequest) SetResourceId(v []*string) *UntagDcdnResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagDcdnResourcesRequest) SetResourceType(v string) *UntagDcdnResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagDcdnResourcesRequest) SetTagKey(v []*string) *UntagDcdnResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagDcdnResourcesRequest) Validate() error {
	return dara.Validate(s)
}
