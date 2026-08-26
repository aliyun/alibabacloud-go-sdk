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
	// The Prometheus instance ID. Required when type is PROMETHEUS or VIRTUAL_PROMETHEUS. Ignored for other types.
	//
	// example:
	//
	// prom-xxxxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The raw V1 datasource JSON string returned as a fallback when type is UNKNOWN and read-path parsing fails. When the frontend detects that this field is not empty, display it as read-only.
	LegacyRaw *string `json:"legacyRaw,omitempty" xml:"legacyRaw,omitempty"`
	// Returned when type is UNKNOWN. Indicates that this rule cannot be edited through the new API. Submit a ticket to contact the CloudMonitor team.
	LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
	// The namespace. Optional when type is VIRTUAL_PROMETHEUS. Identifies the namespace to which the virtual Prometheus instance belongs.
	//
	// example:
	//
	// selectdb
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The Alibaba Cloud service category. Optional when type is CLOUD_MONITORING. Returns unknown when the source lacks this information.
	ProductCategory *string `json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	// The Simple Log Service (SLS) project name. Required when type is SLS. All stores share the same project.
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The region ID. Optional for PROMETHEUS, VIRTUAL_PROMETHEUS, UMODEL, APM, XTRACE, EBPF, RUM, and SLS types. Defaults to the region of the rule or gateway. Not used for CLOUD_MONITORING. Use AlertRuleV2.regionId instead for CLOUD_MONITORING.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The list of SLS stores. Used when type is SLS. At least one store is required. Each store contains store and storeType fields. The project and regionId fields have been moved to the top level. The deprecated fields with the same names that remain in stores return a 400 error if used in write paths.
	Stores []*Stores `json:"stores,omitempty" xml:"stores,omitempty" type:"Repeated"`
	// The tenant ID. Optional when type is VIRTUAL_PROMETHEUS. Identifies the tenant to which the virtual Prometheus instance belongs.
	//
	// example:
	//
	// t-xxxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The data source type. Valid values and associated fields: PROMETHEUS (instanceId required; regionId optional). VIRTUAL_PROMETHEUS (instanceId required; regionId, namespace, and tenantId optional). UMODEL (regionId optional; other fields are carried in queryConfig/conditionConfig). APM (regionId optional). XTRACE (regionId optional). EBPF (regionId optional). RUM (regionId optional). CLOUD_MONITORING (regionId and productCategory optional). SLS (project and stores required). UNKNOWN (read-only fallback; do not use in write paths). Non-enumerated values (such as CMS_BASIC_DS/SLS_DS) are prohibited and the backend returns an Invalidtype 400 error.
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
