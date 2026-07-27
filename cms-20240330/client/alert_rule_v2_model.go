// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleV2 interface {
	dara.Model
	String() string
	GoString() string
	SetActionIntegrationConfig(v *ActionIntegrationConfig) *AlertRuleV2
	GetActionIntegrationConfig() *ActionIntegrationConfig
	SetAnnotations(v map[string]*string) *AlertRuleV2
	GetAnnotations() map[string]*string
	SetArmsIntegrationConfig(v *ArmsIntegrationConfig) *AlertRuleV2
	GetArmsIntegrationConfig() *ArmsIntegrationConfig
	SetBizSource(v string) *AlertRuleV2
	GetBizSource() *string
	SetConditionConfig(v *ConditionConfigUnified) *AlertRuleV2
	GetConditionConfig() *ConditionConfigUnified
	SetContentTemplate(v string) *AlertRuleV2
	GetContentTemplate() *string
	SetCreatedAt(v string) *AlertRuleV2
	GetCreatedAt() *string
	SetDatasourceConfig(v *DatasourceConfigUnified) *AlertRuleV2
	GetDatasourceConfig() *DatasourceConfigUnified
	SetDatasourceType(v string) *AlertRuleV2
	GetDatasourceType() *string
	SetDisplayName(v string) *AlertRuleV2
	GetDisplayName() *string
	SetEnabled(v bool) *AlertRuleV2
	GetEnabled() *bool
	SetLabels(v map[string]*string) *AlertRuleV2
	GetLabels() map[string]*string
	SetNotifyConfig(v *NotifyConfigUnified) *AlertRuleV2
	GetNotifyConfig() *NotifyConfigUnified
	SetNotifyStrategyId(v string) *AlertRuleV2
	GetNotifyStrategyId() *string
	SetObserveResourceConfig(v *ObserveResourceConfig) *AlertRuleV2
	GetObserveResourceConfig() *ObserveResourceConfig
	SetObserveResourceGlobalScope(v bool) *AlertRuleV2
	GetObserveResourceGlobalScope() *bool
	SetObserveResourceList(v []*string) *AlertRuleV2
	GetObserveResourceList() []*string
	SetObserveResourceType(v string) *AlertRuleV2
	GetObserveResourceType() *string
	SetPartitionKey(v string) *AlertRuleV2
	GetPartitionKey() *string
	SetQueryConfig(v *QueryConfigUnified) *AlertRuleV2
	GetQueryConfig() *QueryConfigUnified
	SetRcaConfig(v *AlertRuleRcaConfig) *AlertRuleV2
	GetRcaConfig() *AlertRuleRcaConfig
	SetRegionId(v string) *AlertRuleV2
	GetRegionId() *string
	SetScheduleConfig(v *ScheduleConfigUnified) *AlertRuleV2
	GetScheduleConfig() *ScheduleConfigUnified
	SetSeverityLevels(v string) *AlertRuleV2
	GetSeverityLevels() *string
	SetStatus(v string) *AlertRuleV2
	GetStatus() *string
	SetUpdatedAt(v string) *AlertRuleV2
	GetUpdatedAt() *string
	SetUuid(v string) *AlertRuleV2
	GetUuid() *string
	SetWorkspace(v string) *AlertRuleV2
	GetWorkspace() *string
}

type AlertRuleV2 struct {
	ActionIntegrationConfig *ActionIntegrationConfig `json:"actionIntegrationConfig,omitempty" xml:"actionIntegrationConfig,omitempty"`
	Annotations             map[string]*string       `json:"annotations,omitempty" xml:"annotations,omitempty"`
	ArmsIntegrationConfig   *ArmsIntegrationConfig   `json:"armsIntegrationConfig,omitempty" xml:"armsIntegrationConfig,omitempty"`
	BizSource               *string                  `json:"bizSource,omitempty" xml:"bizSource,omitempty"`
	ConditionConfig         *ConditionConfigUnified  `json:"conditionConfig,omitempty" xml:"conditionConfig,omitempty"`
	ContentTemplate         *string                  `json:"contentTemplate,omitempty" xml:"contentTemplate,omitempty"`
	CreatedAt               *string                  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DatasourceConfig        *DatasourceConfigUnified `json:"datasourceConfig,omitempty" xml:"datasourceConfig,omitempty"`
	DatasourceType          *string                  `json:"datasourceType,omitempty" xml:"datasourceType,omitempty"`
	DisplayName             *string                  `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Enabled                 *bool                    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Labels                  map[string]*string       `json:"labels,omitempty" xml:"labels,omitempty"`
	NotifyConfig            *NotifyConfigUnified     `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty"`
	NotifyStrategyId        *string                  `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	ObserveResourceConfig   *ObserveResourceConfig   `json:"observeResourceConfig,omitempty" xml:"observeResourceConfig,omitempty"`
	// Deprecated
	ObserveResourceGlobalScope *bool     `json:"observeResourceGlobalScope,omitempty" xml:"observeResourceGlobalScope,omitempty"`
	ObserveResourceList        []*string `json:"observeResourceList,omitempty" xml:"observeResourceList,omitempty" type:"Repeated"`
	// Deprecated
	ObserveResourceType *string                `json:"observeResourceType,omitempty" xml:"observeResourceType,omitempty"`
	PartitionKey        *string                `json:"partitionKey,omitempty" xml:"partitionKey,omitempty"`
	QueryConfig         *QueryConfigUnified    `json:"queryConfig,omitempty" xml:"queryConfig,omitempty"`
	RcaConfig           *AlertRuleRcaConfig    `json:"rcaConfig,omitempty" xml:"rcaConfig,omitempty"`
	RegionId            *string                `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ScheduleConfig      *ScheduleConfigUnified `json:"scheduleConfig,omitempty" xml:"scheduleConfig,omitempty"`
	SeverityLevels      *string                `json:"severityLevels,omitempty" xml:"severityLevels,omitempty"`
	Status              *string                `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt           *string                `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Uuid                *string                `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Workspace           *string                `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s AlertRuleV2) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleV2) GoString() string {
	return s.String()
}

func (s *AlertRuleV2) GetActionIntegrationConfig() *ActionIntegrationConfig {
	return s.ActionIntegrationConfig
}

func (s *AlertRuleV2) GetAnnotations() map[string]*string {
	return s.Annotations
}

func (s *AlertRuleV2) GetArmsIntegrationConfig() *ArmsIntegrationConfig {
	return s.ArmsIntegrationConfig
}

func (s *AlertRuleV2) GetBizSource() *string {
	return s.BizSource
}

func (s *AlertRuleV2) GetConditionConfig() *ConditionConfigUnified {
	return s.ConditionConfig
}

func (s *AlertRuleV2) GetContentTemplate() *string {
	return s.ContentTemplate
}

func (s *AlertRuleV2) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *AlertRuleV2) GetDatasourceConfig() *DatasourceConfigUnified {
	return s.DatasourceConfig
}

func (s *AlertRuleV2) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *AlertRuleV2) GetDisplayName() *string {
	return s.DisplayName
}

func (s *AlertRuleV2) GetEnabled() *bool {
	return s.Enabled
}

func (s *AlertRuleV2) GetLabels() map[string]*string {
	return s.Labels
}

func (s *AlertRuleV2) GetNotifyConfig() *NotifyConfigUnified {
	return s.NotifyConfig
}

func (s *AlertRuleV2) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *AlertRuleV2) GetObserveResourceConfig() *ObserveResourceConfig {
	return s.ObserveResourceConfig
}

func (s *AlertRuleV2) GetObserveResourceGlobalScope() *bool {
	return s.ObserveResourceGlobalScope
}

func (s *AlertRuleV2) GetObserveResourceList() []*string {
	return s.ObserveResourceList
}

func (s *AlertRuleV2) GetObserveResourceType() *string {
	return s.ObserveResourceType
}

func (s *AlertRuleV2) GetPartitionKey() *string {
	return s.PartitionKey
}

func (s *AlertRuleV2) GetQueryConfig() *QueryConfigUnified {
	return s.QueryConfig
}

func (s *AlertRuleV2) GetRcaConfig() *AlertRuleRcaConfig {
	return s.RcaConfig
}

func (s *AlertRuleV2) GetRegionId() *string {
	return s.RegionId
}

func (s *AlertRuleV2) GetScheduleConfig() *ScheduleConfigUnified {
	return s.ScheduleConfig
}

func (s *AlertRuleV2) GetSeverityLevels() *string {
	return s.SeverityLevels
}

func (s *AlertRuleV2) GetStatus() *string {
	return s.Status
}

func (s *AlertRuleV2) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *AlertRuleV2) GetUuid() *string {
	return s.Uuid
}

func (s *AlertRuleV2) GetWorkspace() *string {
	return s.Workspace
}

func (s *AlertRuleV2) SetActionIntegrationConfig(v *ActionIntegrationConfig) *AlertRuleV2 {
	s.ActionIntegrationConfig = v
	return s
}

func (s *AlertRuleV2) SetAnnotations(v map[string]*string) *AlertRuleV2 {
	s.Annotations = v
	return s
}

func (s *AlertRuleV2) SetArmsIntegrationConfig(v *ArmsIntegrationConfig) *AlertRuleV2 {
	s.ArmsIntegrationConfig = v
	return s
}

func (s *AlertRuleV2) SetBizSource(v string) *AlertRuleV2 {
	s.BizSource = &v
	return s
}

func (s *AlertRuleV2) SetConditionConfig(v *ConditionConfigUnified) *AlertRuleV2 {
	s.ConditionConfig = v
	return s
}

func (s *AlertRuleV2) SetContentTemplate(v string) *AlertRuleV2 {
	s.ContentTemplate = &v
	return s
}

func (s *AlertRuleV2) SetCreatedAt(v string) *AlertRuleV2 {
	s.CreatedAt = &v
	return s
}

func (s *AlertRuleV2) SetDatasourceConfig(v *DatasourceConfigUnified) *AlertRuleV2 {
	s.DatasourceConfig = v
	return s
}

func (s *AlertRuleV2) SetDatasourceType(v string) *AlertRuleV2 {
	s.DatasourceType = &v
	return s
}

func (s *AlertRuleV2) SetDisplayName(v string) *AlertRuleV2 {
	s.DisplayName = &v
	return s
}

func (s *AlertRuleV2) SetEnabled(v bool) *AlertRuleV2 {
	s.Enabled = &v
	return s
}

func (s *AlertRuleV2) SetLabels(v map[string]*string) *AlertRuleV2 {
	s.Labels = v
	return s
}

func (s *AlertRuleV2) SetNotifyConfig(v *NotifyConfigUnified) *AlertRuleV2 {
	s.NotifyConfig = v
	return s
}

func (s *AlertRuleV2) SetNotifyStrategyId(v string) *AlertRuleV2 {
	s.NotifyStrategyId = &v
	return s
}

func (s *AlertRuleV2) SetObserveResourceConfig(v *ObserveResourceConfig) *AlertRuleV2 {
	s.ObserveResourceConfig = v
	return s
}

func (s *AlertRuleV2) SetObserveResourceGlobalScope(v bool) *AlertRuleV2 {
	s.ObserveResourceGlobalScope = &v
	return s
}

func (s *AlertRuleV2) SetObserveResourceList(v []*string) *AlertRuleV2 {
	s.ObserveResourceList = v
	return s
}

func (s *AlertRuleV2) SetObserveResourceType(v string) *AlertRuleV2 {
	s.ObserveResourceType = &v
	return s
}

func (s *AlertRuleV2) SetPartitionKey(v string) *AlertRuleV2 {
	s.PartitionKey = &v
	return s
}

func (s *AlertRuleV2) SetQueryConfig(v *QueryConfigUnified) *AlertRuleV2 {
	s.QueryConfig = v
	return s
}

func (s *AlertRuleV2) SetRcaConfig(v *AlertRuleRcaConfig) *AlertRuleV2 {
	s.RcaConfig = v
	return s
}

func (s *AlertRuleV2) SetRegionId(v string) *AlertRuleV2 {
	s.RegionId = &v
	return s
}

func (s *AlertRuleV2) SetScheduleConfig(v *ScheduleConfigUnified) *AlertRuleV2 {
	s.ScheduleConfig = v
	return s
}

func (s *AlertRuleV2) SetSeverityLevels(v string) *AlertRuleV2 {
	s.SeverityLevels = &v
	return s
}

func (s *AlertRuleV2) SetStatus(v string) *AlertRuleV2 {
	s.Status = &v
	return s
}

func (s *AlertRuleV2) SetUpdatedAt(v string) *AlertRuleV2 {
	s.UpdatedAt = &v
	return s
}

func (s *AlertRuleV2) SetUuid(v string) *AlertRuleV2 {
	s.Uuid = &v
	return s
}

func (s *AlertRuleV2) SetWorkspace(v string) *AlertRuleV2 {
	s.Workspace = &v
	return s
}

func (s *AlertRuleV2) Validate() error {
	if s.ActionIntegrationConfig != nil {
		if err := s.ActionIntegrationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ArmsIntegrationConfig != nil {
		if err := s.ArmsIntegrationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ConditionConfig != nil {
		if err := s.ConditionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DatasourceConfig != nil {
		if err := s.DatasourceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NotifyConfig != nil {
		if err := s.NotifyConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ObserveResourceConfig != nil {
		if err := s.ObserveResourceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.QueryConfig != nil {
		if err := s.QueryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RcaConfig != nil {
		if err := s.RcaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
