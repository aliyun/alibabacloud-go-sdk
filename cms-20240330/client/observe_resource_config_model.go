// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveResourceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEntityDomain(v string) *ObserveResourceConfig
	GetEntityDomain() *string
	SetEntityType(v string) *ObserveResourceConfig
	GetEntityType() *string
	SetNamespace(v string) *ObserveResourceConfig
	GetNamespace() *string
	SetProductCategory(v string) *ObserveResourceConfig
	GetProductCategory() *string
	SetRelationType(v string) *ObserveResourceConfig
	GetRelationType() *string
	SetResources(v []*string) *ObserveResourceConfig
	GetResources() []*string
}

type ObserveResourceConfig struct {
	EntityDomain    *string   `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	EntityType      *string   `json:"entityType,omitempty" xml:"entityType,omitempty"`
	Namespace       *string   `json:"namespace,omitempty" xml:"namespace,omitempty"`
	ProductCategory *string   `json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	RelationType    *string   `json:"relationType,omitempty" xml:"relationType,omitempty"`
	Resources       []*string `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s ObserveResourceConfig) String() string {
	return dara.Prettify(s)
}

func (s ObserveResourceConfig) GoString() string {
	return s.String()
}

func (s *ObserveResourceConfig) GetEntityDomain() *string {
	return s.EntityDomain
}

func (s *ObserveResourceConfig) GetEntityType() *string {
	return s.EntityType
}

func (s *ObserveResourceConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ObserveResourceConfig) GetProductCategory() *string {
	return s.ProductCategory
}

func (s *ObserveResourceConfig) GetRelationType() *string {
	return s.RelationType
}

func (s *ObserveResourceConfig) GetResources() []*string {
	return s.Resources
}

func (s *ObserveResourceConfig) SetEntityDomain(v string) *ObserveResourceConfig {
	s.EntityDomain = &v
	return s
}

func (s *ObserveResourceConfig) SetEntityType(v string) *ObserveResourceConfig {
	s.EntityType = &v
	return s
}

func (s *ObserveResourceConfig) SetNamespace(v string) *ObserveResourceConfig {
	s.Namespace = &v
	return s
}

func (s *ObserveResourceConfig) SetProductCategory(v string) *ObserveResourceConfig {
	s.ProductCategory = &v
	return s
}

func (s *ObserveResourceConfig) SetRelationType(v string) *ObserveResourceConfig {
	s.RelationType = &v
	return s
}

func (s *ObserveResourceConfig) SetResources(v []*string) *ObserveResourceConfig {
	s.Resources = v
	return s
}

func (s *ObserveResourceConfig) Validate() error {
	return dara.Validate(s)
}
