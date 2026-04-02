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
	SetDisplayName(v string) *AlertRuleV2
	GetDisplayName() *string
	SetEnabled(v bool) *AlertRuleV2
	GetEnabled() *bool
	SetLabels(v map[string]*string) *AlertRuleV2
	GetLabels() map[string]*string
	SetNotifyConfig(v *NotifyConfigUnified) *AlertRuleV2
	GetNotifyConfig() *NotifyConfigUnified
	SetQueryConfig(v *QueryConfigUnified) *AlertRuleV2
	GetQueryConfig() *QueryConfigUnified
	SetScheduleConfig(v *ScheduleConfigUnified) *AlertRuleV2
	GetScheduleConfig() *ScheduleConfigUnified
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
	// 注解
	Annotations           map[string]*string      `json:"annotations,omitempty" xml:"annotations,omitempty"`
	ArmsIntegrationConfig *ArmsIntegrationConfig  `json:"armsIntegrationConfig,omitempty" xml:"armsIntegrationConfig,omitempty"`
	ConditionConfig       *ConditionConfigUnified `json:"conditionConfig,omitempty" xml:"conditionConfig,omitempty"`
	// 内容模板
	ContentTemplate *string `json:"contentTemplate,omitempty" xml:"contentTemplate,omitempty"`
	// 创建时间（只读），ISO 8601
	CreatedAt        *string                  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	DatasourceConfig *DatasourceConfigUnified `json:"datasourceConfig,omitempty" xml:"datasourceConfig,omitempty"`
	// 显示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 是否启用
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 标签
	Labels         map[string]*string     `json:"labels,omitempty" xml:"labels,omitempty"`
	NotifyConfig   *NotifyConfigUnified   `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty"`
	QueryConfig    *QueryConfigUnified    `json:"queryConfig,omitempty" xml:"queryConfig,omitempty"`
	ScheduleConfig *ScheduleConfigUnified `json:"scheduleConfig,omitempty" xml:"scheduleConfig,omitempty"`
	// 告警状态（只读）
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 更新时间（只读），ISO 8601
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// 规则 UUID（系统生成，只读）
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 工作空间
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

func (s *AlertRuleV2) GetQueryConfig() *QueryConfigUnified {
	return s.QueryConfig
}

func (s *AlertRuleV2) GetScheduleConfig() *ScheduleConfigUnified {
	return s.ScheduleConfig
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

func (s *AlertRuleV2) SetQueryConfig(v *QueryConfigUnified) *AlertRuleV2 {
	s.QueryConfig = v
	return s
}

func (s *AlertRuleV2) SetScheduleConfig(v *ScheduleConfigUnified) *AlertRuleV2 {
	s.ScheduleConfig = v
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
