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
	SetResourcePrincipalName(v []*string) *UntagResourcesRequest
	GetResourcePrincipalName() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. Valid values:
	//
	// 	- true: remove all tags from the resources.
	//
	// 	- false (default): does not remove all tags from the resources.
	//
	// > This parameter takes effect only when TagKey.N is not set in the request.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The IDs of resources.
	//
	// Valid values of N: 1 to 50. If the ResourceType parameter is set to user, the resource ID is the ID of the RAM user.
	//
	// > You must specify only one of the following parameters: ResourceId and ResourcePrincipalName.
	//
	// example:
	//
	// UntagResources
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The names of resources.
	//
	// Valid values of N: 1 to 50. If the ResourceType parameter is set to user, the resource name is the name of the RAM user.
	//
	// > You must specify only one of the following parameters: ResourceId and ResourcePrincipalName.
	ResourcePrincipalName []*string `json:"ResourcePrincipalName,omitempty" xml:"ResourcePrincipalName,omitempty" type:"Repeated"`
	// The type of the resource. Valid value:
	//
	// 	- user: a RAM user
	//
	// example:
	//
	// user
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys of resources.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) GetResourcePrincipalName() []*string {
	return s.ResourcePrincipalName
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

func (s *UntagResourcesRequest) SetResourcePrincipalName(v []*string) *UntagResourcesRequest {
	s.ResourcePrincipalName = v
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
