// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionAndNotifyStrategyForListView interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *SubscriptionAndNotifyStrategyForListView
	GetCreateTime() *string
	SetDescription(v string) *SubscriptionAndNotifyStrategyForListView
	GetDescription() *string
	SetEnabled(v bool) *SubscriptionAndNotifyStrategyForListView
	GetEnabled() *bool
	SetMigrationBatchId(v string) *SubscriptionAndNotifyStrategyForListView
	GetMigrationBatchId() *string
	SetMigrationMeta(v string) *SubscriptionAndNotifyStrategyForListView
	GetMigrationMeta() *string
	SetName(v string) *SubscriptionAndNotifyStrategyForListView
	GetName() *string
	SetNotifyStrategy(v *NotifyStrategyForSNSView) *SubscriptionAndNotifyStrategyForListView
	GetNotifyStrategy() *NotifyStrategyForSNSView
	SetNotifyStrategyUuid(v string) *SubscriptionAndNotifyStrategyForListView
	GetNotifyStrategyUuid() *string
	SetSubscriptionUuid(v string) *SubscriptionAndNotifyStrategyForListView
	GetSubscriptionUuid() *string
	SetUpdateTime(v string) *SubscriptionAndNotifyStrategyForListView
	GetUpdateTime() *string
	SetUserId(v string) *SubscriptionAndNotifyStrategyForListView
	GetUserId() *string
	SetUuid(v string) *SubscriptionAndNotifyStrategyForListView
	GetUuid() *string
	SetVersion(v int32) *SubscriptionAndNotifyStrategyForListView
	GetVersion() *int32
	SetWorkspace(v string) *SubscriptionAndNotifyStrategyForListView
	GetWorkspace() *string
}

type SubscriptionAndNotifyStrategyForListView struct {
	CreateTime         *string                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description        *string                   `json:"description,omitempty" xml:"description,omitempty"`
	Enabled            *bool                     `json:"enabled,omitempty" xml:"enabled,omitempty"`
	MigrationBatchId   *string                   `json:"migrationBatchId,omitempty" xml:"migrationBatchId,omitempty"`
	MigrationMeta      *string                   `json:"migrationMeta,omitempty" xml:"migrationMeta,omitempty"`
	Name               *string                   `json:"name,omitempty" xml:"name,omitempty"`
	NotifyStrategy     *NotifyStrategyForSNSView `json:"notifyStrategy,omitempty" xml:"notifyStrategy,omitempty"`
	NotifyStrategyUuid *string                   `json:"notifyStrategyUuid,omitempty" xml:"notifyStrategyUuid,omitempty"`
	SubscriptionUuid   *string                   `json:"subscriptionUuid,omitempty" xml:"subscriptionUuid,omitempty"`
	UpdateTime         *string                   `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId             *string                   `json:"userId,omitempty" xml:"userId,omitempty"`
	Uuid               *string                   `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version            *int32                    `json:"version,omitempty" xml:"version,omitempty"`
	Workspace          *string                   `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SubscriptionAndNotifyStrategyForListView) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionAndNotifyStrategyForListView) GoString() string {
	return s.String()
}

func (s *SubscriptionAndNotifyStrategyForListView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SubscriptionAndNotifyStrategyForListView) GetDescription() *string {
	return s.Description
}

func (s *SubscriptionAndNotifyStrategyForListView) GetEnabled() *bool {
	return s.Enabled
}

func (s *SubscriptionAndNotifyStrategyForListView) GetMigrationBatchId() *string {
	return s.MigrationBatchId
}

func (s *SubscriptionAndNotifyStrategyForListView) GetMigrationMeta() *string {
	return s.MigrationMeta
}

func (s *SubscriptionAndNotifyStrategyForListView) GetName() *string {
	return s.Name
}

func (s *SubscriptionAndNotifyStrategyForListView) GetNotifyStrategy() *NotifyStrategyForSNSView {
	return s.NotifyStrategy
}

func (s *SubscriptionAndNotifyStrategyForListView) GetNotifyStrategyUuid() *string {
	return s.NotifyStrategyUuid
}

func (s *SubscriptionAndNotifyStrategyForListView) GetSubscriptionUuid() *string {
	return s.SubscriptionUuid
}

func (s *SubscriptionAndNotifyStrategyForListView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SubscriptionAndNotifyStrategyForListView) GetUserId() *string {
	return s.UserId
}

func (s *SubscriptionAndNotifyStrategyForListView) GetUuid() *string {
	return s.Uuid
}

func (s *SubscriptionAndNotifyStrategyForListView) GetVersion() *int32 {
	return s.Version
}

func (s *SubscriptionAndNotifyStrategyForListView) GetWorkspace() *string {
	return s.Workspace
}

func (s *SubscriptionAndNotifyStrategyForListView) SetCreateTime(v string) *SubscriptionAndNotifyStrategyForListView {
	s.CreateTime = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetDescription(v string) *SubscriptionAndNotifyStrategyForListView {
	s.Description = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetEnabled(v bool) *SubscriptionAndNotifyStrategyForListView {
	s.Enabled = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetMigrationBatchId(v string) *SubscriptionAndNotifyStrategyForListView {
	s.MigrationBatchId = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetMigrationMeta(v string) *SubscriptionAndNotifyStrategyForListView {
	s.MigrationMeta = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetName(v string) *SubscriptionAndNotifyStrategyForListView {
	s.Name = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetNotifyStrategy(v *NotifyStrategyForSNSView) *SubscriptionAndNotifyStrategyForListView {
	s.NotifyStrategy = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetNotifyStrategyUuid(v string) *SubscriptionAndNotifyStrategyForListView {
	s.NotifyStrategyUuid = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetSubscriptionUuid(v string) *SubscriptionAndNotifyStrategyForListView {
	s.SubscriptionUuid = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetUpdateTime(v string) *SubscriptionAndNotifyStrategyForListView {
	s.UpdateTime = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetUserId(v string) *SubscriptionAndNotifyStrategyForListView {
	s.UserId = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetUuid(v string) *SubscriptionAndNotifyStrategyForListView {
	s.Uuid = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetVersion(v int32) *SubscriptionAndNotifyStrategyForListView {
	s.Version = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) SetWorkspace(v string) *SubscriptionAndNotifyStrategyForListView {
	s.Workspace = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForListView) Validate() error {
	if s.NotifyStrategy != nil {
		if err := s.NotifyStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
