// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagDcdnResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v []*string) *TagDcdnResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *TagDcdnResourcesRequest
	GetResourceType() *string
	SetTag(v []*TagDcdnResourcesRequestTag) *TagDcdnResourcesRequest
	GetTag() []*TagDcdnResourcesRequestTag
}

type TagDcdnResourcesRequest struct {
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
	// The tags.
	//
	// This parameter is required.
	Tag []*TagDcdnResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagDcdnResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagDcdnResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagDcdnResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagDcdnResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagDcdnResourcesRequest) GetTag() []*TagDcdnResourcesRequestTag {
	return s.Tag
}

func (s *TagDcdnResourcesRequest) SetResourceId(v []*string) *TagDcdnResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagDcdnResourcesRequest) SetResourceType(v string) *TagDcdnResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagDcdnResourcesRequest) SetTag(v []*TagDcdnResourcesRequestTag) *TagDcdnResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagDcdnResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type TagDcdnResourcesRequestTag struct {
	// The key of the tag. Valid values of N: **1*	- to **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a tag. Valid values of N: **1*	- to **20**.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagDcdnResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagDcdnResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagDcdnResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagDcdnResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagDcdnResourcesRequestTag) SetKey(v string) *TagDcdnResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagDcdnResourcesRequestTag) SetValue(v string) *TagDcdnResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *TagDcdnResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
