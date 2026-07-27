// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveResourceConfigFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEntityDomain(v *QueryAlertRulesEntityDomainFilter) *ObserveResourceConfigFilter
	GetEntityDomain() *QueryAlertRulesEntityDomainFilter
	SetEntityType(v *QueryAlertRulesEntityTypeFilter) *ObserveResourceConfigFilter
	GetEntityType() *QueryAlertRulesEntityTypeFilter
	SetNamespace(v *QueryAlertRulesNamespaceFilter) *ObserveResourceConfigFilter
	GetNamespace() *QueryAlertRulesNamespaceFilter
	SetProductCategory(v *QueryAlertRulesProductCategoryFilter) *ObserveResourceConfigFilter
	GetProductCategory() *QueryAlertRulesProductCategoryFilter
	SetRelationType(v *QueryAlertRulesRelationTypeFilter) *ObserveResourceConfigFilter
	GetRelationType() *QueryAlertRulesRelationTypeFilter
	SetResources(v *QueryAlertRulesResourcesFilter) *ObserveResourceConfigFilter
	GetResources() *QueryAlertRulesResourcesFilter
}

type ObserveResourceConfigFilter struct {
	EntityDomain    *QueryAlertRulesEntityDomainFilter    `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	EntityType      *QueryAlertRulesEntityTypeFilter      `json:"entityType,omitempty" xml:"entityType,omitempty"`
	Namespace       *QueryAlertRulesNamespaceFilter       `json:"namespace,omitempty" xml:"namespace,omitempty"`
	ProductCategory *QueryAlertRulesProductCategoryFilter `json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	RelationType    *QueryAlertRulesRelationTypeFilter    `json:"relationType,omitempty" xml:"relationType,omitempty"`
	Resources       *QueryAlertRulesResourcesFilter       `json:"resources,omitempty" xml:"resources,omitempty"`
}

func (s ObserveResourceConfigFilter) String() string {
	return dara.Prettify(s)
}

func (s ObserveResourceConfigFilter) GoString() string {
	return s.String()
}

func (s *ObserveResourceConfigFilter) GetEntityDomain() *QueryAlertRulesEntityDomainFilter {
	return s.EntityDomain
}

func (s *ObserveResourceConfigFilter) GetEntityType() *QueryAlertRulesEntityTypeFilter {
	return s.EntityType
}

func (s *ObserveResourceConfigFilter) GetNamespace() *QueryAlertRulesNamespaceFilter {
	return s.Namespace
}

func (s *ObserveResourceConfigFilter) GetProductCategory() *QueryAlertRulesProductCategoryFilter {
	return s.ProductCategory
}

func (s *ObserveResourceConfigFilter) GetRelationType() *QueryAlertRulesRelationTypeFilter {
	return s.RelationType
}

func (s *ObserveResourceConfigFilter) GetResources() *QueryAlertRulesResourcesFilter {
	return s.Resources
}

func (s *ObserveResourceConfigFilter) SetEntityDomain(v *QueryAlertRulesEntityDomainFilter) *ObserveResourceConfigFilter {
	s.EntityDomain = v
	return s
}

func (s *ObserveResourceConfigFilter) SetEntityType(v *QueryAlertRulesEntityTypeFilter) *ObserveResourceConfigFilter {
	s.EntityType = v
	return s
}

func (s *ObserveResourceConfigFilter) SetNamespace(v *QueryAlertRulesNamespaceFilter) *ObserveResourceConfigFilter {
	s.Namespace = v
	return s
}

func (s *ObserveResourceConfigFilter) SetProductCategory(v *QueryAlertRulesProductCategoryFilter) *ObserveResourceConfigFilter {
	s.ProductCategory = v
	return s
}

func (s *ObserveResourceConfigFilter) SetRelationType(v *QueryAlertRulesRelationTypeFilter) *ObserveResourceConfigFilter {
	s.RelationType = v
	return s
}

func (s *ObserveResourceConfigFilter) SetResources(v *QueryAlertRulesResourcesFilter) *ObserveResourceConfigFilter {
	s.Resources = v
	return s
}

func (s *ObserveResourceConfigFilter) Validate() error {
	if s.EntityDomain != nil {
		if err := s.EntityDomain.Validate(); err != nil {
			return err
		}
	}
	if s.EntityType != nil {
		if err := s.EntityType.Validate(); err != nil {
			return err
		}
	}
	if s.Namespace != nil {
		if err := s.Namespace.Validate(); err != nil {
			return err
		}
	}
	if s.ProductCategory != nil {
		if err := s.ProductCategory.Validate(); err != nil {
			return err
		}
	}
	if s.RelationType != nil {
		if err := s.RelationType.Validate(); err != nil {
			return err
		}
	}
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}
