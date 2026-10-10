// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResource interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *TagResource
	GetResourceId() *string
	SetResourceType(v string) *TagResource
	GetResourceType() *string
	SetTagKey(v string) *TagResource
	GetTagKey() *string
	SetTagValue(v string) *TagResource
	GetTagValue() *string
}

type TagResource struct {
	// The resource ID, which is the data catalog ID.
	//
	// example:
	//
	// clg-paimon-0424965be0c240acb4159688c9e2c4b6
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type, which is fixed to CATALOGRESOURCE.
	//
	// example:
	//
	// CATALOGRESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// team
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// recommendation
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s TagResource) String() string {
	return dara.Prettify(s)
}

func (s TagResource) GoString() string {
	return s.String()
}

func (s *TagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *TagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *TagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *TagResource) SetResourceId(v string) *TagResource {
	s.ResourceId = &v
	return s
}

func (s *TagResource) SetResourceType(v string) *TagResource {
	s.ResourceType = &v
	return s
}

func (s *TagResource) SetTagKey(v string) *TagResource {
	s.TagKey = &v
	return s
}

func (s *TagResource) SetTagValue(v string) *TagResource {
	s.TagValue = &v
	return s
}

func (s *TagResource) Validate() error {
	return dara.Validate(s)
}
