// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteAll(v bool) *UntagResourcesRequest
	GetDeleteAll() *bool
	SetResourceIds(v string) *UntagResourcesRequest
	GetResourceIds() *string
	SetResourceRegionId(v string) *UntagResourcesRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKeys(v string) *UntagResourcesRequest
	GetTagKeys() *string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all existing tags from the specified resources. Default value: false. Valid values:
	//
	// 	- **true**: removes all existing tags from the specified resources.
	//
	// 	- **false**: does not remove all existing tags from the specified resources.
	//
	// > All existing tags of a resource are removed only if the **tagKeys*	- parameter is left empty and the **DeleteAll*	- parameter is set to true.
	//
	// example:
	//
	// true
	DeleteAll *bool `json:"DeleteAll,omitempty" xml:"DeleteAll,omitempty"`
	// The IDs of the resources from which you want to remove tags. You can specify up to 20 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["f5ad6ff7-xxxx-xxxx-xxxx-2axxxx82xxxx"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The region in which the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **application**: Enterprise Distributed Application Service (EDAS) application
	//
	// 	- **cluster**: EDAS cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to remove. You can specify up to 20 tags. Set this parameter to a JSON array.
	//
	// example:
	//
	// ["tagKey1","tagKey2"]
	TagKeys *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetDeleteAll() *bool {
	return s.DeleteAll
}

func (s *UntagResourcesRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *UntagResourcesRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKeys() *string {
	return s.TagKeys
}

func (s *UntagResourcesRequest) SetDeleteAll(v bool) *UntagResourcesRequest {
	s.DeleteAll = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v string) *UntagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceRegionId(v string) *UntagResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v string) *UntagResourcesRequest {
	s.TagKeys = &v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
