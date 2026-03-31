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
	SetResourceNames(v []*string) *UntagResourcesRequest
	GetResourceNames() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKeys(v []*string) *UntagResourcesRequest
	GetTagKeys() []*string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resources.
	//
	// Enumerated values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The names of the resources. You can specify up to 50 resource names.
	ResourceNames []*string `json:"ResourceNames,omitempty" xml:"ResourceNames,omitempty" type:"Repeated"`
	// The resource type.
	//
	// Enumerated values:
	//
	// 	- role: RAM roles.
	//
	// 	- policy: policies.
	//
	// example:
	//
	// role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags. You can specify up to 20 tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) GetResourceNames() []*string {
	return s.ResourceNames
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

func (s *UntagResourcesRequest) SetResourceNames(v []*string) *UntagResourcesRequest {
	s.ResourceNames = v
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
