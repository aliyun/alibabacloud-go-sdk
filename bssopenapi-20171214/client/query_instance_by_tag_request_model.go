// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceByTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v []*string) *QueryInstanceByTagRequest
	GetResourceId() []*string
	SetResourceType(v string) *QueryInstanceByTagRequest
	GetResourceType() *string
	SetTag(v []*QueryInstanceByTagRequestTag) *QueryInstanceByTagRequest
	GetTag() []*QueryInstanceByTagRequestTag
}

type QueryInstanceByTagRequest struct {
	// The IDs of resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Specify the savings plan instance as the type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*QueryInstanceByTagRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryInstanceByTagRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceByTagRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *QueryInstanceByTagRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryInstanceByTagRequest) GetTag() []*QueryInstanceByTagRequestTag {
	return s.Tag
}

func (s *QueryInstanceByTagRequest) SetResourceId(v []*string) *QueryInstanceByTagRequest {
	s.ResourceId = v
	return s
}

func (s *QueryInstanceByTagRequest) SetResourceType(v string) *QueryInstanceByTagRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryInstanceByTagRequest) SetTag(v []*QueryInstanceByTagRequestTag) *QueryInstanceByTagRequest {
	s.Tag = v
	return s
}

func (s *QueryInstanceByTagRequest) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceByTagRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// ecs
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryInstanceByTagRequestTag) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceByTagRequestTag) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagRequestTag) GetKey() *string {
	return s.Key
}

func (s *QueryInstanceByTagRequestTag) GetValue() *string {
	return s.Value
}

func (s *QueryInstanceByTagRequestTag) SetKey(v string) *QueryInstanceByTagRequestTag {
	s.Key = &v
	return s
}

func (s *QueryInstanceByTagRequestTag) SetValue(v string) *QueryInstanceByTagRequestTag {
	s.Value = &v
	return s
}

func (s *QueryInstanceByTagRequestTag) Validate() error {
	return dara.Validate(s)
}
