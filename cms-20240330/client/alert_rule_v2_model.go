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
	// Configuration for action integrations, such as webhooks, that execute when an alert is triggered.
	ActionIntegrationConfig *ActionIntegrationConfig `json:"actionIntegrationConfig,omitempty" xml:"actionIntegrationConfig,omitempty"`
	// A set of key-value pairs that serve as annotations, providing additional, non-identifying information, such as a description or a runbook link.
	Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// The configuration for integrating the alert rule with Application Real-Time Monitoring Service (ARMS).
	ArmsIntegrationConfig *ArmsIntegrationConfig `json:"armsIntegrationConfig,omitempty" xml:"armsIntegrationConfig,omitempty"`
	// The configuration for the conditions that trigger an alert.
	ConditionConfig *ConditionConfigUnified `json:"conditionConfig,omitempty" xml:"conditionConfig,omitempty"`
	// The template for the alert notification content.
	ContentTemplate *string `json:"contentTemplate,omitempty" xml:"contentTemplate,omitempty"`
	// The time the alert rule was created.
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The configuration for the data source to be evaluated.
	DatasourceConfig *DatasourceConfigUnified `json:"datasourceConfig,omitempty" xml:"datasourceConfig,omitempty"`
	// The data source type. Examples: `sls`, `prometheus`.
	DatasourceType *string `json:"datasourceType,omitempty" xml:"datasourceType,omitempty"`
	// The user-defined display name for the alert rule.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Indicates whether the alert rule is active. Set to `true` to enable the rule, or `false` to disable it.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// A set of key-value pairs that serve as labels to filter and group alert rules.
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The configuration for sending notifications when an alert is triggered.
	NotifyConfig *NotifyConfigUnified `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty"`
	// The ID of the notification strategy to use for this alert rule.
	NotifyStrategyId *string `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// Indicates whether the alert rule monitors all resources of the specified type. If `true`, the rule applies globally within the workspace.
	ObserveResourceGlobalScope *bool `json:"observeResourceGlobalScope,omitempty" xml:"observeResourceGlobalScope,omitempty"`
	// A list of specific resource IDs to monitor, used only when `observeResourceGlobalScope` is `false`.
	ObserveResourceList []*string `json:"observeResourceList,omitempty" xml:"observeResourceList,omitempty" type:"Repeated"`
	// The type of resource that the alert rule monitors.
	ObserveResourceType *string `json:"observeResourceType,omitempty" xml:"observeResourceType,omitempty"`
	// The partition key used to group alerts. Alerts with the same partition key are treated as a single incident.
	PartitionKey *string `json:"partitionKey,omitempty" xml:"partitionKey,omitempty"`
	// The configuration for querying and processing data from the data source.
	QueryConfig *QueryConfigUnified `json:"queryConfig,omitempty" xml:"queryConfig,omitempty"`
	// The configuration for how often the alert rule is evaluated.
	ScheduleConfig *ScheduleConfigUnified `json:"scheduleConfig,omitempty" xml:"scheduleConfig,omitempty"`
	// The severity level of the alert. Examples: `critical`, `warning`.
	SeverityLevels *string `json:"severityLevels,omitempty" xml:"severityLevels,omitempty"`
	// The current status of the alert rule. Examples: `RUNNING`, `STOPPED`.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time the alert rule was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The unique identifier for the alert rule.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// The ID of the workspace that contains the alert rule.
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
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
	if s.QueryConfig != nil {
		if err := s.QueryConfig.Validate(); err != nil {
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
