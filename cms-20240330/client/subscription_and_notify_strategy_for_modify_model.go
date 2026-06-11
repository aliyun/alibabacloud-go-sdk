// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionAndNotifyStrategyForModify interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubscriptionAndNotifyStrategyForModify
	GetDescription() *string
	SetEnabled(v bool) *SubscriptionAndNotifyStrategyForModify
	GetEnabled() *bool
	SetName(v string) *SubscriptionAndNotifyStrategyForModify
	GetName() *string
	SetNotifyStrategy(v *NotifyStrategyForSNSModify) *SubscriptionAndNotifyStrategyForModify
	GetNotifyStrategy() *NotifyStrategyForSNSModify
	SetResponsePlan(v *IncidentResponsePlanForSNSModify) *SubscriptionAndNotifyStrategyForModify
	GetResponsePlan() *IncidentResponsePlanForSNSModify
	SetSubscription(v *SubscriptionForSNSModify) *SubscriptionAndNotifyStrategyForModify
	GetSubscription() *SubscriptionForSNSModify
	SetSubscriptions(v []*SubscriptionOp) *SubscriptionAndNotifyStrategyForModify
	GetSubscriptions() []*SubscriptionOp
	SetUuid(v string) *SubscriptionAndNotifyStrategyForModify
	GetUuid() *string
	SetVersion(v int32) *SubscriptionAndNotifyStrategyForModify
	GetVersion() *int32
}

type SubscriptionAndNotifyStrategyForModify struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Optional. If omitted, the backend derives the name from `notifyStrategy`.
	Name           *string                           `json:"name,omitempty" xml:"name,omitempty"`
	NotifyStrategy *NotifyStrategyForSNSModify       `json:"notifyStrategy,omitempty" xml:"notifyStrategy,omitempty"`
	ResponsePlan   *IncidentResponsePlanForSNSModify `json:"responsePlan,omitempty" xml:"responsePlan,omitempty"`
	Subscription   *SubscriptionForSNSModify         `json:"subscription,omitempty" xml:"subscription,omitempty"`
	// For update operations only. Use this parameter to batch create, update, and remove member subscriptions.
	Subscriptions []*SubscriptionOp `json:"subscriptions,omitempty" xml:"subscriptions,omitempty" type:"Repeated"`
	// Required for update operations but optional for create operations. If omitted during creation, the backend automatically generates a UUID.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// Required for update operations. The value must match the current version of the record. If the versions do not match, the request fails with an `OPTIMISTIC_LOCK_FAILED` error.
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SubscriptionAndNotifyStrategyForModify) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionAndNotifyStrategyForModify) GoString() string {
	return s.String()
}

func (s *SubscriptionAndNotifyStrategyForModify) GetDescription() *string {
	return s.Description
}

func (s *SubscriptionAndNotifyStrategyForModify) GetEnabled() *bool {
	return s.Enabled
}

func (s *SubscriptionAndNotifyStrategyForModify) GetName() *string {
	return s.Name
}

func (s *SubscriptionAndNotifyStrategyForModify) GetNotifyStrategy() *NotifyStrategyForSNSModify {
	return s.NotifyStrategy
}

func (s *SubscriptionAndNotifyStrategyForModify) GetResponsePlan() *IncidentResponsePlanForSNSModify {
	return s.ResponsePlan
}

func (s *SubscriptionAndNotifyStrategyForModify) GetSubscription() *SubscriptionForSNSModify {
	return s.Subscription
}

func (s *SubscriptionAndNotifyStrategyForModify) GetSubscriptions() []*SubscriptionOp {
	return s.Subscriptions
}

func (s *SubscriptionAndNotifyStrategyForModify) GetUuid() *string {
	return s.Uuid
}

func (s *SubscriptionAndNotifyStrategyForModify) GetVersion() *int32 {
	return s.Version
}

func (s *SubscriptionAndNotifyStrategyForModify) SetDescription(v string) *SubscriptionAndNotifyStrategyForModify {
	s.Description = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) SetEnabled(v bool) *SubscriptionAndNotifyStrategyForModify {
	s.Enabled = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) SetName(v string) *SubscriptionAndNotifyStrategyForModify {
	s.Name = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) SetNotifyStrategy(v *NotifyStrategyForSNSModify) *SubscriptionAndNotifyStrategyForModify {
	s.NotifyStrategy = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) SetResponsePlan(v *IncidentResponsePlanForSNSModify) *SubscriptionAndNotifyStrategyForModify {
	s.ResponsePlan = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) SetSubscription(v *SubscriptionForSNSModify) *SubscriptionAndNotifyStrategyForModify {
	s.Subscription = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) SetSubscriptions(v []*SubscriptionOp) *SubscriptionAndNotifyStrategyForModify {
	s.Subscriptions = v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) SetUuid(v string) *SubscriptionAndNotifyStrategyForModify {
	s.Uuid = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) SetVersion(v int32) *SubscriptionAndNotifyStrategyForModify {
	s.Version = &v
	return s
}

func (s *SubscriptionAndNotifyStrategyForModify) Validate() error {
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
