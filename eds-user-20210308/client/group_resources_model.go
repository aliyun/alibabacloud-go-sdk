// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupResources interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *GroupResources
	GetRegion() *string
	SetResourceId(v string) *GroupResources
	GetResourceId() *string
	SetResourceType(v string) *GroupResources
	GetResourceType() *string
}

type GroupResources struct {
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GroupResources) String() string {
	return dara.Prettify(s)
}

func (s GroupResources) GoString() string {
	return s.String()
}

func (s *GroupResources) GetRegion() *string {
	return s.Region
}

func (s *GroupResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *GroupResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *GroupResources) SetRegion(v string) *GroupResources {
	s.Region = &v
	return s
}

func (s *GroupResources) SetResourceId(v string) *GroupResources {
	s.ResourceId = &v
	return s
}

func (s *GroupResources) SetResourceType(v string) *GroupResources {
	s.ResourceType = &v
	return s
}

func (s *GroupResources) Validate() error {
	return dara.Validate(s)
}
