// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleDataSource interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *AlertRuleDataSource
	GetAppType() *string
	SetDsList(v []*AlertRuleDataSourceDsList) *AlertRuleDataSource
	GetDsList() []*AlertRuleDataSourceDsList
	SetInstanceId(v string) *AlertRuleDataSource
	GetInstanceId() *string
	SetNamespace(v string) *AlertRuleDataSource
	GetNamespace() *string
	SetRegionId(v string) *AlertRuleDataSource
	GetRegionId() *string
	SetTenantId(v string) *AlertRuleDataSource
	GetTenantId() *string
	SetType(v string) *AlertRuleDataSource
	GetType() *string
}

type AlertRuleDataSource struct {
	// Applicable data source type: APM_DS.
	//
	// Application type:
	//
	// - apm.
	//
	// example:
	//
	// apm
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// Applicable data source type: SLS_MULTI_DS.
	//
	// List of sub-data sources.
	DsList []*AlertRuleDataSourceDsList `json:"dsList,omitempty" xml:"dsList,omitempty" type:"Repeated"`
	// Applicable data source type: PROMETHEUS_DS.
	//
	// Prometheus instance ID.
	//
	// example:
	//
	// rw-bbe8961b4a59be0
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Applicable data source type: ENTERPRISE_DS.
	//
	// Name of the enterprise cloud monitoring metric repository.
	//
	// example:
	//
	// aliyun-default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Applicable data source types: APM_DS, PROMETHEUS_DS.
	//
	// The regionId to which the data source belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// Data source type.
	//
	// Valid values:
	//
	// - PROMETHEUS_DS: Prometheus data source.
	//
	// - SLS_MULTI_DS: SLS data source.
	//
	// - APM_DS: Application monitoring data source.
	//
	// - CMS_BASIC_DS: Basic cloud monitoring data source.
	//
	// - ENTERPRISE_DS: Enterprise cloud monitoring data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROMETHEUS_DS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleDataSource) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleDataSource) GoString() string {
	return s.String()
}

func (s *AlertRuleDataSource) GetAppType() *string {
	return s.AppType
}

func (s *AlertRuleDataSource) GetDsList() []*AlertRuleDataSourceDsList {
	return s.DsList
}

func (s *AlertRuleDataSource) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AlertRuleDataSource) GetNamespace() *string {
	return s.Namespace
}

func (s *AlertRuleDataSource) GetRegionId() *string {
	return s.RegionId
}

func (s *AlertRuleDataSource) GetTenantId() *string {
	return s.TenantId
}

func (s *AlertRuleDataSource) GetType() *string {
	return s.Type
}

func (s *AlertRuleDataSource) SetAppType(v string) *AlertRuleDataSource {
	s.AppType = &v
	return s
}

func (s *AlertRuleDataSource) SetDsList(v []*AlertRuleDataSourceDsList) *AlertRuleDataSource {
	s.DsList = v
	return s
}

func (s *AlertRuleDataSource) SetInstanceId(v string) *AlertRuleDataSource {
	s.InstanceId = &v
	return s
}

func (s *AlertRuleDataSource) SetNamespace(v string) *AlertRuleDataSource {
	s.Namespace = &v
	return s
}

func (s *AlertRuleDataSource) SetRegionId(v string) *AlertRuleDataSource {
	s.RegionId = &v
	return s
}

func (s *AlertRuleDataSource) SetTenantId(v string) *AlertRuleDataSource {
	s.TenantId = &v
	return s
}

func (s *AlertRuleDataSource) SetType(v string) *AlertRuleDataSource {
	s.Type = &v
	return s
}

func (s *AlertRuleDataSource) Validate() error {
	if s.DsList != nil {
		for _, item := range s.DsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleDataSourceDsList struct {
	// SLS project
	//
	// example:
	//
	// mySlsProject
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The region where the SLS project is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// LogStore/MetricStore name.
	//
	// example:
	//
	// mySlsLogStore
	Store *string `json:"store,omitempty" xml:"store,omitempty"`
	// Type of SLS data sub-source:
	//
	// - SLS_LOG_DS: LogStore data source.
	//
	// - SLS_METRIC_DS: MetricStore data source.
	//
	// example:
	//
	// SLS_LOG_DS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertRuleDataSourceDsList) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleDataSourceDsList) GoString() string {
	return s.String()
}

func (s *AlertRuleDataSourceDsList) GetProject() *string {
	return s.Project
}

func (s *AlertRuleDataSourceDsList) GetRegionId() *string {
	return s.RegionId
}

func (s *AlertRuleDataSourceDsList) GetStore() *string {
	return s.Store
}

func (s *AlertRuleDataSourceDsList) GetType() *string {
	return s.Type
}

func (s *AlertRuleDataSourceDsList) SetProject(v string) *AlertRuleDataSourceDsList {
	s.Project = &v
	return s
}

func (s *AlertRuleDataSourceDsList) SetRegionId(v string) *AlertRuleDataSourceDsList {
	s.RegionId = &v
	return s
}

func (s *AlertRuleDataSourceDsList) SetStore(v string) *AlertRuleDataSourceDsList {
	s.Store = &v
	return s
}

func (s *AlertRuleDataSourceDsList) SetType(v string) *AlertRuleDataSourceDsList {
	s.Type = &v
	return s
}

func (s *AlertRuleDataSourceDsList) Validate() error {
	return dara.Validate(s)
}
