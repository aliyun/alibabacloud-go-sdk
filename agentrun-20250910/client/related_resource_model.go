// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelatedResource interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *RelatedResource
	GetResourceId() *string
	SetResourceName(v string) *RelatedResource
	GetResourceName() *string
	SetResourceType(v string) *RelatedResource
	GetResourceType() *string
}

type RelatedResource struct {
	ResourceId   *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s RelatedResource) String() string {
	return dara.Prettify(s)
}

func (s RelatedResource) GoString() string {
	return s.String()
}

func (s *RelatedResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *RelatedResource) GetResourceName() *string {
	return s.ResourceName
}

func (s *RelatedResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *RelatedResource) SetResourceId(v string) *RelatedResource {
	s.ResourceId = &v
	return s
}

func (s *RelatedResource) SetResourceName(v string) *RelatedResource {
	s.ResourceName = &v
	return s
}

func (s *RelatedResource) SetResourceType(v string) *RelatedResource {
	s.ResourceType = &v
	return s
}

func (s *RelatedResource) Validate() error {
	return dara.Validate(s)
}
