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
	SetResourceIds(v string) *UntagResourcesRequest
	GetResourceIds() *string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetTagKeys(v string) *UntagResourcesRequest
	GetTagKeys() *string
	SetBody(v string) *UntagResourcesRequest
	GetBody() *string
}

type UntagResourcesRequest struct {
	// Specifies whether to delete all parts. Default value: **false*	- . This parameter is valid only when **TagKeys*	- is not specified.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The resource list that you want to delete.
	//
	// example:
	//
	// ["es-cn-09k1rocex0006****","es-cn-oew1rgiev0009****"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource. Fixed to **INSTANCE*	- .
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags that you want to delete. The list can contain up to 20 subitems.
	//
	// example:
	//
	// ["tagKey1","tagKey2"]
	TagKeys *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
	Body    *string `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UntagResourcesRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetTagKeys() *string {
	return s.TagKeys
}

func (s *UntagResourcesRequest) GetBody() *string {
	return s.Body
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v string) *UntagResourcesRequest {
	s.ResourceIds = &v
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

func (s *UntagResourcesRequest) SetBody(v string) *UntagResourcesRequest {
	s.Body = &v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
