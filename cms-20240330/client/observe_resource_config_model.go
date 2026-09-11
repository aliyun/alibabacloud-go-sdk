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
	// The UModel entity domain.
	//
	// example:
	//
	// cloud_monitor
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// The UModel entity type.
	//
	// example:
	//
	// ACS::ECS::Instance
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// The CloudMonitor namespace.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The CloudMonitor product category.
	//
	// example:
	//
	// ecs
	ProductCategory *string `json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	// The relation type. TAG is supported only for alert rules where datasourceConfig.type is set to APM and queryConfig.type is set to APM_MULTI_QUERY. UMODEL_ENTITY does not support writes and is used only for reading existing data.
	//
	// example:
	//
	// ALL
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
	// The list of resources. If relationType is set to ALL, this parameter can be left empty, which indicates all resources. If relationType is set to TAG, this parameter is a list of labels in key=value format (such as ["env=prod", "app=foo"]). This is supported only for APM data sources with APM_MULTI_QUERY.
	Resources []*string `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
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
