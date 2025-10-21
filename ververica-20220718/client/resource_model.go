// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResource interface {
	dara.Model
	String() string
	GoString() string
	SetElasticResource(v *ResourceSpec) *Resource
	GetElasticResource() *ResourceSpec
	SetFixedResource(v *ResourceSpec) *Resource
	GetFixedResource() *ResourceSpec
}

type Resource struct {
	ElasticResource *ResourceSpec `json:"elasticResource,omitempty" xml:"elasticResource,omitempty"`
	// This parameter is required.
	FixedResource *ResourceSpec `json:"fixedResource,omitempty" xml:"fixedResource,omitempty"`
}

func (s Resource) String() string {
	return dara.Prettify(s)
}

func (s Resource) GoString() string {
	return s.String()
}

func (s *Resource) GetElasticResource() *ResourceSpec {
	return s.ElasticResource
}

func (s *Resource) GetFixedResource() *ResourceSpec {
	return s.FixedResource
}

func (s *Resource) SetElasticResource(v *ResourceSpec) *Resource {
	s.ElasticResource = v
	return s
}

func (s *Resource) SetFixedResource(v *ResourceSpec) *Resource {
	s.FixedResource = v
	return s
}

func (s *Resource) Validate() error {
	if s.ElasticResource != nil {
		if err := s.ElasticResource.Validate(); err != nil {
			return err
		}
	}
	if s.FixedResource != nil {
		if err := s.FixedResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}
