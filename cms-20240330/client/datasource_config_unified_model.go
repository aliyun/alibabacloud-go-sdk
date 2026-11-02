// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasourceConfigUnified interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DatasourceConfigUnified
	GetInstanceId() *string
	SetLegacyRaw(v string) *DatasourceConfigUnified
	GetLegacyRaw() *string
	SetLegacyType(v string) *DatasourceConfigUnified
	GetLegacyType() *string
	SetNamespace(v string) *DatasourceConfigUnified
	GetNamespace() *string
	SetProductCategory(v string) *DatasourceConfigUnified
	GetProductCategory() *string
	SetProject(v string) *DatasourceConfigUnified
	GetProject() *string
	SetRegionId(v string) *DatasourceConfigUnified
	GetRegionId() *string
	SetStores(v []*Stores) *DatasourceConfigUnified
	GetStores() []*Stores
	SetTenantId(v string) *DatasourceConfigUnified
	GetTenantId() *string
	SetType(v string) *DatasourceConfigUnified
	GetType() *string
}

type DatasourceConfigUnified struct {
	// The Prometheus instance ID. This field is used when type is PROMETHEUS.
	//
	// example:
	//
	// rw-xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The original V1 datasource JSON string returned as a fallback when type is UNKNOWN and read path parsing fails. The frontend displays this field as read-only when it is not empty.
	//
	// example:
	//
	// {"type":"SLS"}
	LegacyRaw *string `json:"legacyRaw,omitempty" xml:"legacyRaw,omitempty"`
	// The legacy type, which is returned when type is UNKNOWN to indicate that this rule cannot be edited by using the new API. Submit a ticket to contact the CloudMonitor team.
	//
	// example:
	//
	// SLS
	LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
	// The namespace. This field is optional when type is VIRTUAL_PROMETHEUS and identifies the namespace to which the virtual Prometheus instance belongs.
	//
	// example:
	//
	// selectdb
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The cloud service category. This field is optional when type is CLOUD_MONITORING. Outputs unknown when this information is missing from the source.
	//
	// example:
	//
	// ecs
	ProductCategory *string `json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	// The Simple Log Service project name. This field is required when type is SLS, and all stores share the same project.
	//
	// example:
	//
	// my-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The region ID. This field is optional for all types and defaults to the rule or gateway region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The Simple Log Service store list. This field is used when type is SLS and requires at least one store. Each store includes store and storeType. The project and regionId fields are moved to the top level. The fields with the same names remaining in Stores are deprecated and return 400 in the write path.
	//
	// example:
	//
	// [{"store":"cms-alert-log","storeType":"log"}]
	Stores []*Stores `json:"stores,omitempty" xml:"stores,omitempty" type:"Repeated"`
	// The tenant ID. This field is optional when type is VIRTUAL_PROMETHEUS and identifies the tenant to which the virtual Prometheus instance belongs.
	//
	// example:
	//
	// t-xxxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The data source type.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROMETHEUS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DatasourceConfigUnified) String() string {
	return dara.Prettify(s)
}

func (s DatasourceConfigUnified) GoString() string {
	return s.String()
}

func (s *DatasourceConfigUnified) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DatasourceConfigUnified) GetLegacyRaw() *string {
	return s.LegacyRaw
}

func (s *DatasourceConfigUnified) GetLegacyType() *string {
	return s.LegacyType
}

func (s *DatasourceConfigUnified) GetNamespace() *string {
	return s.Namespace
}

func (s *DatasourceConfigUnified) GetProductCategory() *string {
	return s.ProductCategory
}

func (s *DatasourceConfigUnified) GetProject() *string {
	return s.Project
}

func (s *DatasourceConfigUnified) GetRegionId() *string {
	return s.RegionId
}

func (s *DatasourceConfigUnified) GetStores() []*Stores {
	return s.Stores
}

func (s *DatasourceConfigUnified) GetTenantId() *string {
	return s.TenantId
}

func (s *DatasourceConfigUnified) GetType() *string {
	return s.Type
}

func (s *DatasourceConfigUnified) SetInstanceId(v string) *DatasourceConfigUnified {
	s.InstanceId = &v
	return s
}

func (s *DatasourceConfigUnified) SetLegacyRaw(v string) *DatasourceConfigUnified {
	s.LegacyRaw = &v
	return s
}

func (s *DatasourceConfigUnified) SetLegacyType(v string) *DatasourceConfigUnified {
	s.LegacyType = &v
	return s
}

func (s *DatasourceConfigUnified) SetNamespace(v string) *DatasourceConfigUnified {
	s.Namespace = &v
	return s
}

func (s *DatasourceConfigUnified) SetProductCategory(v string) *DatasourceConfigUnified {
	s.ProductCategory = &v
	return s
}

func (s *DatasourceConfigUnified) SetProject(v string) *DatasourceConfigUnified {
	s.Project = &v
	return s
}

func (s *DatasourceConfigUnified) SetRegionId(v string) *DatasourceConfigUnified {
	s.RegionId = &v
	return s
}

func (s *DatasourceConfigUnified) SetStores(v []*Stores) *DatasourceConfigUnified {
	s.Stores = v
	return s
}

func (s *DatasourceConfigUnified) SetTenantId(v string) *DatasourceConfigUnified {
	s.TenantId = &v
	return s
}

func (s *DatasourceConfigUnified) SetType(v string) *DatasourceConfigUnified {
	s.Type = &v
	return s
}

func (s *DatasourceConfigUnified) Validate() error {
	if s.Stores != nil {
		for _, item := range s.Stores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
