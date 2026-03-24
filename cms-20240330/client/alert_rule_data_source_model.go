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
	SetType(v string) *AlertRuleDataSource
	GetType() *string
}

type AlertRuleDataSource struct {
	AppType *string                      `json:"appType,omitempty" xml:"appType,omitempty"`
	DsList  []*AlertRuleDataSourceDsList `json:"dsList,omitempty" xml:"dsList,omitempty" type:"Repeated"`
	// 实例id，当type=PROMETHEUS_DS/ENTERPRISE_DS时必填，为prometheus实例的clusterId或指标仓库名称
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Namespace  *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 数据源类型
	//
	// This parameter is required.
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
	Project  *string `json:"project,omitempty" xml:"project,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Store    *string `json:"store,omitempty" xml:"store,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
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
