// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelatedWorkload interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *RelatedWorkload
	GetResourceId() *string
	SetResourceName(v string) *RelatedWorkload
	GetResourceName() *string
	SetResourceType(v string) *RelatedWorkload
	GetResourceType() *string
}

type RelatedWorkload struct {
	ResourceId   *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s RelatedWorkload) String() string {
	return dara.Prettify(s)
}

func (s RelatedWorkload) GoString() string {
	return s.String()
}

func (s *RelatedWorkload) GetResourceId() *string {
	return s.ResourceId
}

func (s *RelatedWorkload) GetResourceName() *string {
	return s.ResourceName
}

func (s *RelatedWorkload) GetResourceType() *string {
	return s.ResourceType
}

func (s *RelatedWorkload) SetResourceId(v string) *RelatedWorkload {
	s.ResourceId = &v
	return s
}

func (s *RelatedWorkload) SetResourceName(v string) *RelatedWorkload {
	s.ResourceName = &v
	return s
}

func (s *RelatedWorkload) SetResourceType(v string) *RelatedWorkload {
	s.ResourceType = &v
	return s
}

func (s *RelatedWorkload) Validate() error {
	return dara.Validate(s)
}
