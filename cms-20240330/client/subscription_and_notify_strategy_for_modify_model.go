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
	// The description of the alert policy.
	//
	// example:
	//
	// Used to monitor the CPU utilization of ECS instances
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Read-only. This parameter does not take effect even if specified. The backend forcibly sets this parameter to true during creation and retains the current value during updates. To enable or disable the policy, call the EnableAlertPolicy or DisableAlertPolicy operation.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Policy Name of the alert policy. If this parameter is not specified, the backend derives Policy Name from notifyStrategy.
	//
	// example:
	//
	// my-alert-policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The notification configuration that defines noise reduction rules, notification channel routing, and templates. This parameter is required for Create operations.
	NotifyStrategy *NotifyStrategyForSNSModify `json:"notifyStrategy,omitempty" xml:"notifyStrategy,omitempty"`
	// The event management configuration that defines recovery notifications, repeat notifications, automatic recovery, and escalation policies.
	ResponsePlan *IncidentResponsePlanForSNSModify `json:"responsePlan,omitempty" xml:"responsePlan,omitempty"`
	// The single primary subscription configuration that defines event filter conditions. This parameter is mutually exclusive with subscriptions. Do not specify both parameters at the same time.
	Subscription *SubscriptionForSNSModify `json:"subscription,omitempty" xml:"subscription,omitempty"`
	// Dedicated to Update operations. Performs batch create, update, or remove adjustments on member subscriptions.
	Subscriptions []*SubscriptionOp `json:"subscriptions,omitempty" xml:"subscriptions,omitempty" type:"Repeated"`
	// The unique identifier of the alert policy. This parameter is required for Update operations. Do not specify this parameter for Create operations because the backend automatically generates the value.
	//
	// example:
	//
	// 7076c75c-c804-461e-975f-c6f9ed5af745
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// The optimistic lock version number. This parameter is required for Update operations and must match the current value on the backend. Otherwise, a 409 VersionConflict error is returned. The version number increments by 1 after each successful update.
	//
	// example:
	//
	// 1
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
