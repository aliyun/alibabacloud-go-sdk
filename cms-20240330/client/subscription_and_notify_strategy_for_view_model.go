// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionAndNotifyStrategyForView interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *SubscriptionAndNotifyStrategyForView
	GetCreateTime() *string
	SetDescription(v string) *SubscriptionAndNotifyStrategyForView
	GetDescription() *string
	SetEnabled(v bool) *SubscriptionAndNotifyStrategyForView
	GetEnabled() *bool
	SetMigrationBatchId(v string) *SubscriptionAndNotifyStrategyForView
	GetMigrationBatchId() *string
	SetMigrationMeta(v string) *SubscriptionAndNotifyStrategyForView
	GetMigrationMeta() *string
	SetName(v string) *SubscriptionAndNotifyStrategyForView
	GetName() *string
	SetNotifyStrategy(v *NotifyStrategyForSNSView) *SubscriptionAndNotifyStrategyForView
	GetNotifyStrategy() *NotifyStrategyForSNSView
	SetNotifyStrategyUuid(v string) *SubscriptionAndNotifyStrategyForView
	GetNotifyStrategyUuid() *string
	SetResponsePlan(v *IncidentResponsePlanForSNSView) *SubscriptionAndNotifyStrategyForView
	GetResponsePlan() *IncidentResponsePlanForSNSView
	SetSubscription(v *SubscriptionForSNSView) *SubscriptionAndNotifyStrategyForView
	GetSubscription() *SubscriptionForSNSView
	SetSubscriptionUuid(v string) *SubscriptionAndNotifyStrategyForView
	GetSubscriptionUuid() *string
	SetSubscriptions(v []*SubscriptionForView) *SubscriptionAndNotifyStrategyForView
	GetSubscriptions() []*SubscriptionForView
	SetUpdateTime(v string) *SubscriptionAndNotifyStrategyForView
	GetUpdateTime() *string
	SetUserId(v string) *SubscriptionAndNotifyStrategyForView
	GetUserId() *string
	SetUuid(v string) *SubscriptionAndNotifyStrategyForView
	GetUuid() *string
	SetVersion(v int32) *SubscriptionAndNotifyStrategyForView
	GetVersion() *int32
	SetWorkspace(v string) *SubscriptionAndNotifyStrategyForView
	GetWorkspace() *string
}

type SubscriptionAndNotifyStrategyForView struct {
	CreateTime         *string                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description        *string                         `json:"description,omitempty" xml:"description,omitempty"`
	Enabled            *bool                           `json:"enabled,omitempty" xml:"enabled,omitempty"`
	MigrationBatchId   *string                         `json:"migrationBatchId,omitempty" xml:"migrationBatchId,omitempty"`
	MigrationMeta      *string                         `json:"migrationMeta,omitempty" xml:"migrationMeta,omitempty"`
	Name               *string                         `json:"name,omitempty" xml:"name,omitempty"`
	NotifyStrategy     *NotifyStrategyForSNSView       `json:"notifyStrategy,omitempty" xml:"notifyStrategy,omitempty"`
	NotifyStrategyUuid *string                         `json:"notifyStrategyUuid,omitempty" xml:"notifyStrategyUuid,omitempty"`
	ResponsePlan       *IncidentResponsePlanForSNSView `json:"responsePlan,omitempty" xml:"responsePlan,omitempty"`
	Subscription       *SubscriptionForSNSView         `json:"subscription,omitempty" xml:"subscription,omitempty"`
	SubscriptionUuid   *string                         `json:"subscriptionUuid,omitempty" xml:"subscriptionUuid,omitempty"`
	Subscriptions      []*SubscriptionForView          `json:"subscriptions,omitempty" xml:"subscriptions,omitempty" type:"Repeated"`
	UpdateTime         *string                         `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId             *string                         `json:"userId,omitempty" xml:"userId,omitempty"`
	Uuid               *string                         `json:"uuid,omitempty" xml:"uuid,omitempty"`
	Version            *int32                          `json:"version,omitempty" xml:"version,omitempty"`
	Workspace          *string                         `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SubscriptionAndNotifyStrategyForView) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionAndNotifyStrategyForView) GoString() string {
	return s.String()
}

func (s *SubscriptionAndNotifyStrategyForView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SubscriptionAndNotifyStrategyForView) GetDescription() *string {
	return s.Description
}

func (s *SubscriptionAndNotifyStrategyForView) GetEnabled() *bool {
	return s.Enabled
}

func (s *SubscriptionAndNotifyStrategyForView) GetMigrationBatchId() *string {
	return s.MigrationBatchId
}

func (s *SubscriptionAndNotifyStrategyForView) GetMigrationMeta() *string {
	return s.MigrationMeta
}

func (s *SubscriptionAndNotifyStrategyForView) GetName() *string {
	return s.Name
}

func (s *SubscriptionAndNotifyStrategyForView) GetNotifyStrategy() *NotifyStrategyForSNSView {
	return s.NotifyStrategy
}

func (s *SubscriptionAndNotifyStrategyForView) GetNotifyStrategyUuid() *string {
	return s.NotifyStrategyUuid
}

func (s *SubscriptionAndNotifyStrategyForView) GetResponsePlan() *IncidentResponsePlanForSNSView {
	return s.ResponsePlan
}

func (s *SubscriptionAndNotifyStrategyForView) GetSubscription() *SubscriptionForSNSView {
	return s.Subscription
}

func (s *SubscriptionAndNotifyStrategyForView) GetSubscriptionUuid() *string {
	return s.SubscriptionUuid
}

func (s *SubscriptionAndNotifyStrategyForView) GetSubscriptions() []*SubscriptionForView {
	return s.Subscriptions
}

func (s *SubscriptionAndNotifyStrategyForView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SubscriptionAndNotifyStrategyForView) GetUserId() *string {
	return s.UserId
}

func (s *SubscriptionAndNotifyStrategyForView) GetUuid() *string {
	return s.Uuid
}

func (s *SubscriptionAndNotifyStrategyForView) GetVersion() *int32 {
	return s.Version
}

func (s *SubscriptionAndNotifyStrategyForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *SubscriptionAndNotifyStrategyForView) SetCreateTime(v string) *SubscriptionAndNotifyStrategyForView {
	s.CreateTime = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetDescription(v string) *SubscriptionAndNotifyStrategyForView {
	s.Description = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetEnabled(v bool) *SubscriptionAndNotifyStrategyForView {
	s.Enabled = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetMigrationBatchId(v string) *SubscriptionAndNotifyStrategyForView {
	s.MigrationBatchId = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetMigrationMeta(v string) *SubscriptionAndNotifyStrategyForView {
	s.MigrationMeta = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetName(v string) *SubscriptionAndNotifyStrategyForView {
	s.Name = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetNotifyStrategy(v *NotifyStrategyForSNSView) *SubscriptionAndNotifyStrategyForView {
	s.NotifyStrategy = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetNotifyStrategyUuid(v string) *SubscriptionAndNotifyStrategyForView {
	s.NotifyStrategyUuid = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetResponsePlan(v *IncidentResponsePlanForSNSView) *SubscriptionAndNotifyStrategyForView {
	s.ResponsePlan = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetSubscription(v *SubscriptionForSNSView) *SubscriptionAndNotifyStrategyForView {
	s.Subscription = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetSubscriptionUuid(v string) *SubscriptionAndNotifyStrategyForView {
	s.SubscriptionUuid = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetSubscriptions(v []*SubscriptionForView) *SubscriptionAndNotifyStrategyForView {
	s.Subscriptions = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetUpdateTime(v string) *SubscriptionAndNotifyStrategyForView {
	s.UpdateTime = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetUserId(v string) *SubscriptionAndNotifyStrategyForView {
	s.UserId = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetUuid(v string) *SubscriptionAndNotifyStrategyForView {
	s.Uuid = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetVersion(v int32) *SubscriptionAndNotifyStrategyForView {
	s.Version = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) SetWorkspace(v string) *SubscriptionAndNotifyStrategyForView {
	s.Workspace = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForView) Validate() error {
	if s.NotifyStrategy != nil {
		if err := s.NotifyStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.ResponsePlan != nil {
		if err := s.ResponsePlan.Validate(); err != nil {
			return err
		}
	}
	if s.Subscription != nil {
		if err := s.Subscription.Validate(); err != nil {
			return err
		}
	}
	if s.Subscriptions != nil {
		for _, item := range s.Subscriptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
