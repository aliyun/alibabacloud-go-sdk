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
	// Applies to the APM_DS data source type.
	//
	// The type of the application. Valid value:
	//
	// - apm
	//
	// example:
	//
	// apm
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// Applies to the SLS_MULTI_DS data source type.
	//
	// A list of sub-data sources.
	DsList []*AlertRuleDataSourceDsList `json:"dsList,omitempty" xml:"dsList,omitempty" type:"Repeated"`
	// Applies to the PROMETHEUS_DS data source type.
	//
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// rw-bbe8961b4a59be0
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Applies to the ENTERPRISE_DS data source type.
	//
	// The name of the Hybrid Cloud Monitoring metric repository.
	//
	// example:
	//
	// aliyun-default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Applies to the APM_DS and PROMETHEUS_DS data source types.
	//
	// The ID of the region where the data source is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The data source type.
	//
	// Valid values:
	//
	// - PROMETHEUS_DS: A Prometheus data source.
	//
	// - SLS_MULTI_DS: An SLS data source.
	//
	// - APM_DS: An Application Monitoring data source.
	//
	// - CMS_BASIC_DS: A basic Cloud Monitor data source.
	//
	// - ENTERPRISE_DS: A Hybrid Cloud Monitoring data source.
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
	// The SLS project.
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
	// The name of the LogStore or MetricStore.
	//
	// example:
	//
	// mySlsLogStore
	Store *string `json:"store,omitempty" xml:"store,omitempty"`
	// The type of the SLS sub-data source. Valid values:
	//
	// - SLS_LOG_DS: A LogStore data source.
	//
	// - SLS_METRIC_DS: A MetricStore data source.
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
