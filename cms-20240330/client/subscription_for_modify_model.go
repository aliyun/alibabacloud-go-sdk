// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionForModify interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubscriptionForModify
	GetDescription() *string
	SetFilterSetting(v *FilterSetting) *SubscriptionForModify
	GetFilterSetting() *FilterSetting
	SetNotifyStrategyId(v string) *SubscriptionForModify
	GetNotifyStrategyId() *string
	SetPushingSetting(v *SubscriptionForModifyPushingSetting) *SubscriptionForModify
	GetPushingSetting() *SubscriptionForModifyPushingSetting
	SetSubscriptionName(v string) *SubscriptionForModify
	GetSubscriptionName() *string
	SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionForModify
	GetWorkspaceFilterSetting() *WorkspaceFilterSetting
}

type SubscriptionForModify struct {
	// Description.
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Filtering settings.
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// Notification policy UUID.
	//
	// example:
	//
	// 123123
	NotifyStrategyId *string `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// Push settings.
	PushingSetting *SubscriptionForModifyPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	// Name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test subscription.
	SubscriptionName       *string                 `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
	WorkspaceFilterSetting *WorkspaceFilterSetting `json:"workspaceFilterSetting,omitempty" xml:"workspaceFilterSetting,omitempty"`
}

func (s SubscriptionForModify) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForModify) GoString() string {
	return s.String()
}

func (s *SubscriptionForModify) GetDescription() *string {
	return s.Description
}

func (s *SubscriptionForModify) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *SubscriptionForModify) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *SubscriptionForModify) GetPushingSetting() *SubscriptionForModifyPushingSetting {
	return s.PushingSetting
}

func (s *SubscriptionForModify) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *SubscriptionForModify) GetWorkspaceFilterSetting() *WorkspaceFilterSetting {
	return s.WorkspaceFilterSetting
}

func (s *SubscriptionForModify) SetDescription(v string) *SubscriptionForModify {
	s.Description = &v
	return s
}

func (s *SubscriptionForModify) SetFilterSetting(v *FilterSetting) *SubscriptionForModify {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionForModify) SetNotifyStrategyId(v string) *SubscriptionForModify {
	s.NotifyStrategyId = &v
	return s
}

func (s *SubscriptionForModify) SetPushingSetting(v *SubscriptionForModifyPushingSetting) *SubscriptionForModify {
	s.PushingSetting = v
	return s
}

func (s *SubscriptionForModify) SetSubscriptionName(v string) *SubscriptionForModify {
	s.SubscriptionName = &v
	return s
}

func (s *SubscriptionForModify) SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionForModify {
	s.WorkspaceFilterSetting = v
	return s
}

func (s *SubscriptionForModify) Validate() error {
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

type SubscriptionForModifyPushingSetting struct {
	// Alert push action plan ID list.
	AlertActionIds []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	// Action plan ID.
	//
	// example:
	//
	// 123123123
	ResponsePlanId *string `json:"responsePlanId,omitempty" xml:"responsePlanId,omitempty"`
	// Action integration plan ID list.
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	// Template UUID.
	//
	// example:
	//
	// 123123123
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
}

func (s SubscriptionForModifyPushingSetting) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForModifyPushingSetting) GoString() string {
	return s.String()
}

func (s *SubscriptionForModifyPushingSetting) GetAlertActionIds() []*string {
	return s.AlertActionIds
}

func (s *SubscriptionForModifyPushingSetting) GetResponsePlanId() *string {
	return s.ResponsePlanId
}

func (s *SubscriptionForModifyPushingSetting) GetRestoreActionIds() []*string {
	return s.RestoreActionIds
}

func (s *SubscriptionForModifyPushingSetting) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *SubscriptionForModifyPushingSetting) SetAlertActionIds(v []*string) *SubscriptionForModifyPushingSetting {
	s.AlertActionIds = v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetResponsePlanId(v string) *SubscriptionForModifyPushingSetting {
	s.ResponsePlanId = &v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetRestoreActionIds(v []*string) *SubscriptionForModifyPushingSetting {
	s.RestoreActionIds = v
	return s
}

func (s *SubscriptionForModifyPushingSetting) SetTemplateUuid(v string) *SubscriptionForModifyPushingSetting {
	s.TemplateUuid = &v
	return s
}

func (s *SubscriptionForModifyPushingSetting) Validate() error {
	return dara.Validate(s)
}
