// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAlertRulesUnifiedActionInput interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ManageAlertRulesUnifiedActionInput
	GetAction() *string
	SetActionIntegrationConfig(v *ActionIntegrationConfig) *ManageAlertRulesUnifiedActionInput
	GetActionIntegrationConfig() *ActionIntegrationConfig
	SetAnnotations(v map[string]*string) *ManageAlertRulesUnifiedActionInput
	GetAnnotations() map[string]*string
	SetArmsIntegrationConfig(v *ArmsIntegrationConfig) *ManageAlertRulesUnifiedActionInput
	GetArmsIntegrationConfig() *ArmsIntegrationConfig
	SetConditionConfig(v *ConditionConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetConditionConfig() *ConditionConfigUnified
	SetContentTemplate(v string) *ManageAlertRulesUnifiedActionInput
	GetContentTemplate() *string
	SetDatasourceConfig(v *DatasourceConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetDatasourceConfig() *DatasourceConfigUnified
	SetDisplayName(v string) *ManageAlertRulesUnifiedActionInput
	GetDisplayName() *string
	SetEnabled(v bool) *ManageAlertRulesUnifiedActionInput
	GetEnabled() *bool
	SetLabels(v map[string]*string) *ManageAlertRulesUnifiedActionInput
	GetLabels() map[string]*string
	SetNotifyConfig(v *NotifyConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetNotifyConfig() *NotifyConfigUnified
	SetQueryConfig(v *QueryConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetQueryConfig() *QueryConfigUnified
	SetScheduleConfig(v *ScheduleConfigUnified) *ManageAlertRulesUnifiedActionInput
	GetScheduleConfig() *ScheduleConfigUnified
	SetUuid(v string) *ManageAlertRulesUnifiedActionInput
	GetUuid() *string
	SetUuidList(v []*string) *ManageAlertRulesUnifiedActionInput
	GetUuidList() []*string
	SetWorkspace(v string) *ManageAlertRulesUnifiedActionInput
	GetWorkspace() *string
}

type ManageAlertRulesUnifiedActionInput struct {
	// The action to perform on the alert rule. For example, `create` or `update`.
	//
	// This parameter is required.
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The action integration configuration for triggering automated workflows or actions in external systems.
	ActionIntegrationConfig *ActionIntegrationConfig `json:"actionIntegrationConfig,omitempty" xml:"actionIntegrationConfig,omitempty"`
	// A collection of key-value pairs attached to the alert as annotations. Use annotations to provide additional, non-identifying information, such as descriptions or runbook links.
	Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// Configuration for integrating with Application Real-Time Monitoring Service (ARMS).
	ArmsIntegrationConfig *ArmsIntegrationConfig `json:"armsIntegrationConfig,omitempty" xml:"armsIntegrationConfig,omitempty"`
	// The condition configuration that specifies the trigger criteria for the alert.
	ConditionConfig *ConditionConfigUnified `json:"conditionConfig,omitempty" xml:"conditionConfig,omitempty"`
	// The content template for the alert notification. You can use variables to customize the message.
	ContentTemplate *string `json:"contentTemplate,omitempty" xml:"contentTemplate,omitempty"`
	// The data source configuration for the alert rule.
	DatasourceConfig *DatasourceConfigUnified `json:"datasourceConfig,omitempty" xml:"datasourceConfig,omitempty"`
	// The display name of the alert rule, as shown in the console.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether the alert rule is enabled. A value of `true` indicates the rule is active, and `false` indicates it is inactive. Default: `true`.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// A collection of key-value pairs attached to the alert rule as labels. Use labels for categorization and filtering.
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The notification configuration that specifies how and where to send alert notifications.
	NotifyConfig *NotifyConfigUnified `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty"`
	// The query configuration that defines the data for rule evaluation.
	QueryConfig *QueryConfigUnified `json:"queryConfig,omitempty" xml:"queryConfig,omitempty"`
	// The schedule configuration that determines how often the system evaluates the rule.
	ScheduleConfig *ScheduleConfigUnified `json:"scheduleConfig,omitempty" xml:"scheduleConfig,omitempty"`
	// The unique identifier (UUID) of the alert rule. This parameter is required when you update or delete an existing rule.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// A list of UUIDs. Use this parameter to perform bulk actions on multiple rules, such as batch deletion.
	UuidList []*string `json:"uuidList,omitempty" xml:"uuidList,omitempty" type:"Repeated"`
	// The ID of the workspace that contains the alert rule.
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ManageAlertRulesUnifiedActionInput) String() string {
	return dara.Prettify(s)
}

func (s ManageAlertRulesUnifiedActionInput) GoString() string {
	return s.String()
}

func (s *ManageAlertRulesUnifiedActionInput) GetAction() *string {
	return s.Action
}

func (s *ManageAlertRulesUnifiedActionInput) GetActionIntegrationConfig() *ActionIntegrationConfig {
	return s.ActionIntegrationConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetAnnotations() map[string]*string {
	return s.Annotations
}

func (s *ManageAlertRulesUnifiedActionInput) GetArmsIntegrationConfig() *ArmsIntegrationConfig {
	return s.ArmsIntegrationConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetConditionConfig() *ConditionConfigUnified {
	return s.ConditionConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetContentTemplate() *string {
	return s.ContentTemplate
}

func (s *ManageAlertRulesUnifiedActionInput) GetDatasourceConfig() *DatasourceConfigUnified {
	return s.DatasourceConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ManageAlertRulesUnifiedActionInput) GetEnabled() *bool {
	return s.Enabled
}

func (s *ManageAlertRulesUnifiedActionInput) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ManageAlertRulesUnifiedActionInput) GetNotifyConfig() *NotifyConfigUnified {
	return s.NotifyConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetQueryConfig() *QueryConfigUnified {
	return s.QueryConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetScheduleConfig() *ScheduleConfigUnified {
	return s.ScheduleConfig
}

func (s *ManageAlertRulesUnifiedActionInput) GetUuid() *string {
	return s.Uuid
}

func (s *ManageAlertRulesUnifiedActionInput) GetUuidList() []*string {
	return s.UuidList
}

func (s *ManageAlertRulesUnifiedActionInput) GetWorkspace() *string {
	return s.Workspace
}

func (s *ManageAlertRulesUnifiedActionInput) SetAction(v string) *ManageAlertRulesUnifiedActionInput {
	s.Action = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetActionIntegrationConfig(v *ActionIntegrationConfig) *ManageAlertRulesUnifiedActionInput {
	s.ActionIntegrationConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetAnnotations(v map[string]*string) *ManageAlertRulesUnifiedActionInput {
	s.Annotations = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetArmsIntegrationConfig(v *ArmsIntegrationConfig) *ManageAlertRulesUnifiedActionInput {
	s.ArmsIntegrationConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetConditionConfig(v *ConditionConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.ConditionConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetContentTemplate(v string) *ManageAlertRulesUnifiedActionInput {
	s.ContentTemplate = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetDatasourceConfig(v *DatasourceConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.DatasourceConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetDisplayName(v string) *ManageAlertRulesUnifiedActionInput {
	s.DisplayName = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetEnabled(v bool) *ManageAlertRulesUnifiedActionInput {
	s.Enabled = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetLabels(v map[string]*string) *ManageAlertRulesUnifiedActionInput {
	s.Labels = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetNotifyConfig(v *NotifyConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.NotifyConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetQueryConfig(v *QueryConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.QueryConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetScheduleConfig(v *ScheduleConfigUnified) *ManageAlertRulesUnifiedActionInput {
	s.ScheduleConfig = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetUuid(v string) *ManageAlertRulesUnifiedActionInput {
	s.Uuid = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetUuidList(v []*string) *ManageAlertRulesUnifiedActionInput {
	s.UuidList = v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetWorkspace(v string) *ManageAlertRulesUnifiedActionInput {
	s.Workspace = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) Validate() error {
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
