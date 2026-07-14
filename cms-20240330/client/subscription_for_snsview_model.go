// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionForSNSView interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *SubscriptionForSNSView
	GetCreateTime() *string
	SetEnable(v bool) *SubscriptionForSNSView
	GetEnable() *bool
	SetFilterSetting(v *FilterSetting) *SubscriptionForSNSView
	GetFilterSetting() *FilterSetting
	SetMode(v string) *SubscriptionForSNSView
	GetMode() *string
	SetName(v string) *SubscriptionForSNSView
	GetName() *string
	SetNotifyStrategyUuid(v string) *SubscriptionForSNSView
	GetNotifyStrategyUuid() *string
	SetRegionId(v string) *SubscriptionForSNSView
	GetRegionId() *string
	SetSubscribeLegacyEvent(v bool) *SubscriptionForSNSView
	GetSubscribeLegacyEvent() *bool
	SetSubscriptionType(v string) *SubscriptionForSNSView
	GetSubscriptionType() *string
	SetSyncFromType(v string) *SubscriptionForSNSView
	GetSyncFromType() *string
	SetUpdateTime(v string) *SubscriptionForSNSView
	GetUpdateTime() *string
	SetUserId(v string) *SubscriptionForSNSView
	GetUserId() *string
	SetUuid(v string) *SubscriptionForSNSView
	GetUuid() *string
	SetWorkspace(v string) *SubscriptionForSNSView
	GetWorkspace() *string
	SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionForSNSView
	GetWorkspaceFilterSetting() *WorkspaceFilterSetting
}

type SubscriptionForSNSView struct {
	CreateTime         *string        `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Enable             *bool          `json:"enable,omitempty" xml:"enable,omitempty"`
	FilterSetting      *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	Mode               *string        `json:"mode,omitempty" xml:"mode,omitempty"`
	Name               *string        `json:"name,omitempty" xml:"name,omitempty"`
	NotifyStrategyUuid *string        `json:"notifyStrategyUuid,omitempty" xml:"notifyStrategyUuid,omitempty"`
	RegionId           *string        `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Specifies whether to subscribe to legacy product events (CMS 1.0, ARMS, or SLS events where workspace=null). Valid values: true: subscribed. false or null: not subscribed.
	SubscribeLegacyEvent   *bool                   `json:"subscribeLegacyEvent,omitempty" xml:"subscribeLegacyEvent,omitempty"`
	SubscriptionType       *string                 `json:"subscriptionType,omitempty" xml:"subscriptionType,omitempty"`
	SyncFromType           *string                 `json:"syncFromType,omitempty" xml:"syncFromType,omitempty"`
	UpdateTime             *string                 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId                 *string                 `json:"userId,omitempty" xml:"userId,omitempty"`
	Uuid                   *string                 `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Workspace              *string                 `json:"workspace,omitempty" xml:"workspace,omitempty"`
	WorkspaceFilterSetting *WorkspaceFilterSetting `json:"workspaceFilterSetting,omitempty" xml:"workspaceFilterSetting,omitempty"`
}

func (s SubscriptionForSNSView) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForSNSView) GoString() string {
	return s.String()
}

func (s *SubscriptionForSNSView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SubscriptionForSNSView) GetEnable() *bool {
	return s.Enable
}

func (s *SubscriptionForSNSView) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *SubscriptionForSNSView) GetMode() *string {
	return s.Mode
}

func (s *SubscriptionForSNSView) GetName() *string {
	return s.Name
}

func (s *SubscriptionForSNSView) GetNotifyStrategyUuid() *string {
	return s.NotifyStrategyUuid
}

func (s *SubscriptionForSNSView) GetRegionId() *string {
	return s.RegionId
}

func (s *SubscriptionForSNSView) GetSubscribeLegacyEvent() *bool {
	return s.SubscribeLegacyEvent
}

func (s *SubscriptionForSNSView) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *SubscriptionForSNSView) GetSyncFromType() *string {
	return s.SyncFromType
}

func (s *SubscriptionForSNSView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SubscriptionForSNSView) GetUserId() *string {
	return s.UserId
}

func (s *SubscriptionForSNSView) GetUuid() *string {
	return s.Uuid
}

func (s *SubscriptionForSNSView) GetWorkspace() *string {
	return s.Workspace
}

func (s *SubscriptionForSNSView) GetWorkspaceFilterSetting() *WorkspaceFilterSetting {
	return s.WorkspaceFilterSetting
}

func (s *SubscriptionForSNSView) SetCreateTime(v string) *SubscriptionForSNSView {
	s.CreateTime = &v
	return s
}

func (s *SubscriptionForSNSView) SetEnable(v bool) *SubscriptionForSNSView {
	s.Enable = &v
	return s
}

func (s *SubscriptionForSNSView) SetFilterSetting(v *FilterSetting) *SubscriptionForSNSView {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionForSNSView) SetMode(v string) *SubscriptionForSNSView {
	s.Mode = &v
	return s
}

func (s *SubscriptionForSNSView) SetName(v string) *SubscriptionForSNSView {
	s.Name = &v
	return s
}

func (s *SubscriptionForSNSView) SetNotifyStrategyUuid(v string) *SubscriptionForSNSView {
	s.NotifyStrategyUuid = &v
	return s
}

func (s *SubscriptionForSNSView) SetRegionId(v string) *SubscriptionForSNSView {
	s.RegionId = &v
	return s
}

func (s *SubscriptionForSNSView) SetSubscribeLegacyEvent(v bool) *SubscriptionForSNSView {
	s.SubscribeLegacyEvent = &v
	return s
}

func (s *SubscriptionForSNSView) SetSubscriptionType(v string) *SubscriptionForSNSView {
	s.SubscriptionType = &v
	return s
}

func (s *SubscriptionForSNSView) SetSyncFromType(v string) *SubscriptionForSNSView {
	s.SyncFromType = &v
	return s
}

func (s *SubscriptionForSNSView) SetUpdateTime(v string) *SubscriptionForSNSView {
	s.UpdateTime = &v
	return s
}

func (s *SubscriptionForSNSView) SetUserId(v string) *SubscriptionForSNSView {
	s.UserId = &v
	return s
}

func (s *SubscriptionForSNSView) SetUuid(v string) *SubscriptionForSNSView {
	s.Uuid = &v
	return s
}

func (s *SubscriptionForSNSView) SetWorkspace(v string) *SubscriptionForSNSView {
	s.Workspace = &v
	return s
}

func (s *SubscriptionForSNSView) SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionForSNSView {
	s.WorkspaceFilterSetting = v
	return s
}

func (s *SubscriptionForSNSView) Validate() error {
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
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
