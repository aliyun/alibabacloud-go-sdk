// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionForView interface {
	dara.Model
	String() string
	GoString() string
	SetAgentConfig(v *SubscriptionForViewAgentConfig) *SubscriptionForView
	GetAgentConfig() *SubscriptionForViewAgentConfig
	SetCreateTime(v string) *SubscriptionForView
	GetCreateTime() *string
	SetDescription(v string) *SubscriptionForView
	GetDescription() *string
	SetEnable(v bool) *SubscriptionForView
	GetEnable() *bool
	SetFilterSetting(v *FilterSetting) *SubscriptionForView
	GetFilterSetting() *FilterSetting
	SetNotifyStrategyId(v string) *SubscriptionForView
	GetNotifyStrategyId() *string
	SetPushingSetting(v *SubscriptionForViewPushingSetting) *SubscriptionForView
	GetPushingSetting() *SubscriptionForViewPushingSetting
	SetSubscriptionId(v string) *SubscriptionForView
	GetSubscriptionId() *string
	SetSubscriptionName(v string) *SubscriptionForView
	GetSubscriptionName() *string
	SetSubscriptionType(v string) *SubscriptionForView
	GetSubscriptionType() *string
	SetSyncFromType(v string) *SubscriptionForView
	GetSyncFromType() *string
	SetUpdateTime(v string) *SubscriptionForView
	GetUpdateTime() *string
	SetUserId(v string) *SubscriptionForView
	GetUserId() *string
	SetWorkspace(v string) *SubscriptionForView
	GetWorkspace() *string
	SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionForView
	GetWorkspaceFilterSetting() *WorkspaceFilterSetting
}

type SubscriptionForView struct {
	AgentConfig *SubscriptionForViewAgentConfig `json:"agentConfig,omitempty" xml:"agentConfig,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description.
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the subscription is enabled.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The filter settings.
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// The UUID of the notification policy.
	//
	// example:
	//
	// 23123123
	NotifyStrategyId *string `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// The push settings.
	PushingSetting *SubscriptionForViewPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	// The UUID.
	//
	// example:
	//
	// 123123123123
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 订阅测试
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
	// example:
	//
	// NORMAL
	SubscriptionType *string `json:"subscriptionType,omitempty" xml:"subscriptionType,omitempty"`
	// The source type of the synchronization policy.
	//
	// example:
	//
	// "ARMS"
	SyncFromType *string `json:"syncFromType,omitempty" xml:"syncFromType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-05-23T02:29:02Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123123123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// Specifies the workspace.
	//
	// example:
	//
	// workspace-test
	Workspace              *string                 `json:"workspace,omitempty" xml:"workspace,omitempty"`
	WorkspaceFilterSetting *WorkspaceFilterSetting `json:"workspaceFilterSetting,omitempty" xml:"workspaceFilterSetting,omitempty"`
}

func (s SubscriptionForView) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForView) GoString() string {
	return s.String()
}

func (s *SubscriptionForView) GetAgentConfig() *SubscriptionForViewAgentConfig {
	return s.AgentConfig
}

func (s *SubscriptionForView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SubscriptionForView) GetDescription() *string {
	return s.Description
}

func (s *SubscriptionForView) GetEnable() *bool {
	return s.Enable
}

func (s *SubscriptionForView) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *SubscriptionForView) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *SubscriptionForView) GetPushingSetting() *SubscriptionForViewPushingSetting {
	return s.PushingSetting
}

func (s *SubscriptionForView) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *SubscriptionForView) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *SubscriptionForView) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *SubscriptionForView) GetSyncFromType() *string {
	return s.SyncFromType
}

func (s *SubscriptionForView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SubscriptionForView) GetUserId() *string {
	return s.UserId
}

func (s *SubscriptionForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *SubscriptionForView) GetWorkspaceFilterSetting() *WorkspaceFilterSetting {
	return s.WorkspaceFilterSetting
}

func (s *SubscriptionForView) SetAgentConfig(v *SubscriptionForViewAgentConfig) *SubscriptionForView {
	s.AgentConfig = v
	return s
}

func (s *SubscriptionForView) SetCreateTime(v string) *SubscriptionForView {
	s.CreateTime = &v
	return s
}

func (s *SubscriptionForView) SetDescription(v string) *SubscriptionForView {
	s.Description = &v
	return s
}

func (s *SubscriptionForView) SetEnable(v bool) *SubscriptionForView {
	s.Enable = &v
	return s
}

func (s *SubscriptionForView) SetFilterSetting(v *FilterSetting) *SubscriptionForView {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionForView) SetNotifyStrategyId(v string) *SubscriptionForView {
	s.NotifyStrategyId = &v
	return s
}

func (s *SubscriptionForView) SetPushingSetting(v *SubscriptionForViewPushingSetting) *SubscriptionForView {
	s.PushingSetting = v
	return s
}

func (s *SubscriptionForView) SetSubscriptionId(v string) *SubscriptionForView {
	s.SubscriptionId = &v
	return s
}

func (s *SubscriptionForView) SetSubscriptionName(v string) *SubscriptionForView {
	s.SubscriptionName = &v
	return s
}

func (s *SubscriptionForView) SetSubscriptionType(v string) *SubscriptionForView {
	s.SubscriptionType = &v
	return s
}

func (s *SubscriptionForView) SetSyncFromType(v string) *SubscriptionForView {
	s.SyncFromType = &v
	return s
}

func (s *SubscriptionForView) SetUpdateTime(v string) *SubscriptionForView {
	s.UpdateTime = &v
	return s
}

func (s *SubscriptionForView) SetUserId(v string) *SubscriptionForView {
	s.UserId = &v
	return s
}

func (s *SubscriptionForView) SetWorkspace(v string) *SubscriptionForView {
	s.Workspace = &v
	return s
}

func (s *SubscriptionForView) SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionForView {
	s.WorkspaceFilterSetting = v
	return s
}

func (s *SubscriptionForView) Validate() error {
	if s.AgentConfig != nil {
		if err := s.AgentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
			return err
		}
	}
	if s.PushingSetting != nil {
		if err := s.PushingSetting.Validate(); err != nil {
			return err
		}
	}
	if s.WorkspaceFilterSetting != nil {
		if err := s.WorkspaceFilterSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubscriptionForViewAgentConfig struct {
	AgentUuid *string                       `json:"agentUuid,omitempty" xml:"agentUuid,omitempty"`
	Routes    []*NotifyRouteForSubscription `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s SubscriptionForViewAgentConfig) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForViewAgentConfig) GoString() string {
	return s.String()
}

func (s *SubscriptionForViewAgentConfig) GetAgentUuid() *string {
	return s.AgentUuid
}

func (s *SubscriptionForViewAgentConfig) GetRoutes() []*NotifyRouteForSubscription {
	return s.Routes
}

func (s *SubscriptionForViewAgentConfig) SetAgentUuid(v string) *SubscriptionForViewAgentConfig {
	s.AgentUuid = &v
	return s
}

func (s *SubscriptionForViewAgentConfig) SetRoutes(v []*NotifyRouteForSubscription) *SubscriptionForViewAgentConfig {
	s.Routes = v
	return s
}

func (s *SubscriptionForViewAgentConfig) Validate() error {
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubscriptionForViewPushingSetting struct {
	// A list of action integration IDs for alert pushes.
	AlertActionIds []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	// The action plan ID.
	//
	// example:
	//
	// 123123123
	ResponsePlanId *string `json:"responsePlanId,omitempty" xml:"responsePlanId,omitempty"`
	// A list of action integration IDs for recovery pushes.
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	// The template UUID.
	//
	// example:
	//
	// 123123123
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s SubscriptionForViewPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForViewPushingSetting) GoString() string {
	return s.String()
}

func (s *SubscriptionForViewPushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *SubscriptionForViewPushingSetting) GetResponsePlanId() *string {
	return s.ResponsePlanId
}

func (s *SubscriptionForViewPushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *SubscriptionForViewPushingSetting) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *SubscriptionForViewPushingSetting) SetAlertActionIds(v []*string) *SubscriptionForViewPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *SubscriptionForViewPushingSetting) SetResponsePlanId(v string) *SubscriptionForViewPushingSetting {
	s.ResponsePlanId = &v
	return s
}

func (s *SubscriptionForViewPushingSetting) SetRestoreActionIds(v []*string) *SubscriptionForViewPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *SubscriptionForViewPushingSetting) SetTemplateUuid(v string) *SubscriptionForViewPushingSetting {
	s.TemplateUuid = &v
	return s
}

func (s *SubscriptionForViewPushingSetting) Validate() error {
	return dara.Validate(s)
}
