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
	SetProductCategory(v string) *DatasourceConfigUnified
	GetProductCategory() *string
	SetProject(v string) *DatasourceConfigUnified
	GetProject() *string
	SetRegionId(v string) *DatasourceConfigUnified
	GetRegionId() *string
	SetStores(v []*Stores) *DatasourceConfigUnified
	GetStores() []*Stores
	SetType(v string) *DatasourceConfigUnified
	GetType() *string
}

type DatasourceConfigUnified struct {
	// The Prometheus instance ID (required when type=PROMETHEUS; ignored for other types).
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The original V1 datasource JSON string returned as a fallback when type=UNKNOWN and the read path fails to parse the datasource. If the frontend detects that this field is not empty, display it as read-only.
	LegacyRaw *string `json:"legacyRaw,omitempty" xml:"legacyRaw,omitempty"`
	// Returned when type=UNKNOWN, indicating that this rule cannot be edited through the new API. Submit a ticket to contact the CloudMonitor team.
	LegacyType *string `json:"legacyType,omitempty" xml:"legacyType,omitempty"`
	// The Alibaba Cloud service category (optional when type=CLOUD_MONITORING). If the source does not contain this information, the value unknown is returned.
	ProductCategory *string `json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	// The Simple Log Service project name (required when type=SLS; all stores share the same project).
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The region ID (optional for PROMETHEUS / UMODEL / APM / SLS types; defaults to the same region as the rule or gateway. CLOUD_MONITORING does not use this field; use AlertRuleV2.regionId instead).
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The list of Simple Log Service stores (used when type=SLS; at least one store is required). Each store contains store and storeType fields. The project and regionId fields have been moved to the top level. The deprecated fields with the same names that remain in stores cause a 400 error if used in write paths.
	Stores []*Stores `json:"stores,omitempty" xml:"stores,omitempty" type:"Repeated"`
	// The datasource type. Valid values: PROMETHEUS (instanceId is required; regionId is optional). UMODEL (regionId is optional; other settings are carried in queryConfig/conditionConfig). APM (regionId is optional). CLOUD_MONITORING (regionId and productCategory are optional). UNKNOWN (read-only fallback; do not use in write paths). Do not use non-enumerated values (such as CMS_BASIC_DS or SLS_DS). The backend returns an Invalidtype 400 error.
	//
	// This parameter is required.
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
