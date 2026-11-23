// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v string) *UntagResourcesRequest
	GetAll() *string
	SetResourceId(v []*string) *UntagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all custom tags. Valid values: true, false, and an empty string. If this parameter is not specified or is set to an empty string, the value is treated as false. The value true takes effect only when the tagKey parameter is not specified or is an empty array.
	//
	// example:
	//
	// true
	All *string `json:"all,omitempty" xml:"all,omitempty"`
	// The list of resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["space-a","space-b"]
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to agentspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// agentspace
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The list of tag keys to remove. You can specify up to 20 tag keys. The tag keys must be unique. If this parameter is not empty, it takes precedence over the all parameter.
	//
	// example:
	//
	// ["key1", "key2", "key3"]
	TagKey []*string `json:"tagKey,omitempty" xml:"tagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetAll() *string {
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

func (s *UntagResourcesRequest) SetAll(v string) *UntagResourcesRequest {
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
