// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionForView interface {
	dara.Model
	String() string
	GoString() string
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
	SetSyncFromType(v string) *SubscriptionForView
	GetSyncFromType() *string
	SetUpdateTime(v string) *SubscriptionForView
	GetUpdateTime() *string
	SetUserId(v string) *SubscriptionForView
	GetUserId() *string
	SetWorkspace(v string) *SubscriptionForView
	GetWorkspace() *string
}

type SubscriptionForView struct {
	CreateTime       *string                            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description      *string                            `json:"description,omitempty" xml:"description,omitempty"`
	Enable           *bool                              `json:"enable,omitempty" xml:"enable,omitempty"`
	FilterSetting    *FilterSetting                     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	NotifyStrategyId *string                            `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	PushingSetting   *SubscriptionForViewPushingSetting `json:"pushingSetting,omitempty" xml:"pushingSetting,omitempty" type:"Struct"`
	SubscriptionId   *string                            `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	// This parameter is required.
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
	SyncFromType     *string `json:"syncFromType,omitempty" xml:"syncFromType,omitempty"`
	UpdateTime       *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace        *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SubscriptionForView) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForView) GoString() string {
	return s.String()
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

func (s *SubscriptionForView) Validate() error {
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
	return nil
}

type SubscriptionForViewPushingSetting struct {
	AlertActionIds   []*string `json:"alertActionIds,omitempty" xml:"alertActionIds,omitempty" type:"Repeated"`
	ResponsePlanId   *string   `json:"responsePlanId,omitempty" xml:"responsePlanId,omitempty"`
	RestoreActionIds []*string `json:"restoreActionIds,omitempty" xml:"restoreActionIds,omitempty" type:"Repeated"`
	TemplateUuid     *string   `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
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
