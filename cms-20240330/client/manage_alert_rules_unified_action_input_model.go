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
	SetObserveResourceInstanceId(v string) *ManageAlertRulesUnifiedActionInput
	GetObserveResourceInstanceId() *string
	SetObserveResourceType(v string) *ManageAlertRulesUnifiedActionInput
	GetObserveResourceType() *string
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
	// This parameter is required.
	Action                    *string                  `json:"action,omitempty" xml:"action,omitempty"`
	ActionIntegrationConfig   *ActionIntegrationConfig `json:"actionIntegrationConfig,omitempty" xml:"actionIntegrationConfig,omitempty"`
	Annotations               map[string]*string       `json:"annotations,omitempty" xml:"annotations,omitempty"`
	ArmsIntegrationConfig     *ArmsIntegrationConfig   `json:"armsIntegrationConfig,omitempty" xml:"armsIntegrationConfig,omitempty"`
	ConditionConfig           *ConditionConfigUnified  `json:"conditionConfig,omitempty" xml:"conditionConfig,omitempty"`
	ContentTemplate           *string                  `json:"contentTemplate,omitempty" xml:"contentTemplate,omitempty"`
	DatasourceConfig          *DatasourceConfigUnified `json:"datasourceConfig,omitempty" xml:"datasourceConfig,omitempty"`
	DisplayName               *string                  `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Enabled                   *bool                    `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Labels                    map[string]*string       `json:"labels,omitempty" xml:"labels,omitempty"`
	NotifyConfig              *NotifyConfigUnified     `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty"`
	ObserveResourceInstanceId *string                  `json:"observeResourceInstanceId,omitempty" xml:"observeResourceInstanceId,omitempty"`
	ObserveResourceType       *string                  `json:"observeResourceType,omitempty" xml:"observeResourceType,omitempty"`
	QueryConfig               *QueryConfigUnified      `json:"queryConfig,omitempty" xml:"queryConfig,omitempty"`
	ScheduleConfig            *ScheduleConfigUnified   `json:"scheduleConfig,omitempty" xml:"scheduleConfig,omitempty"`
	Uuid                      *string                  `json:"uuid,omitempty" xml:"uuid,omitempty"`
	UuidList                  []*string                `json:"uuidList,omitempty" xml:"uuidList,omitempty" type:"Repeated"`
	Workspace                 *string                  `json:"workspace,omitempty" xml:"workspace,omitempty"`
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

func (s *ManageAlertRulesUnifiedActionInput) GetObserveResourceInstanceId() *string {
	return s.ObserveResourceInstanceId
}

func (s *ManageAlertRulesUnifiedActionInput) GetObserveResourceType() *string {
	return s.ObserveResourceType
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

func (s *ManageAlertRulesUnifiedActionInput) SetObserveResourceInstanceId(v string) *ManageAlertRulesUnifiedActionInput {
	s.ObserveResourceInstanceId = &v
	return s
}

func (s *ManageAlertRulesUnifiedActionInput) SetObserveResourceType(v string) *ManageAlertRulesUnifiedActionInput {
	s.ObserveResourceType = &v
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
