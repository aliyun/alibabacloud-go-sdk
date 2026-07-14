// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *NotifyPolicy
	GetCreateTime() *string
	SetDescription(v string) *NotifyPolicy
	GetDescription() *string
	SetEnabled(v bool) *NotifyPolicy
	GetEnabled() *bool
	SetName(v string) *NotifyPolicy
	GetName() *string
	SetNotifyStrategy(v *NotifyStrategyDetail) *NotifyPolicy
	GetNotifyStrategy() *NotifyStrategyDetail
	SetResponsePlan(v *ResponsePlanDetail) *NotifyPolicy
	GetResponsePlan() *ResponsePlanDetail
	SetSubscription(v *SubscriptionDetail) *NotifyPolicy
	GetSubscription() *SubscriptionDetail
	SetUpdateTime(v string) *NotifyPolicy
	GetUpdateTime() *string
	SetUserId(v string) *NotifyPolicy
	GetUserId() *string
	SetUuid(v string) *NotifyPolicy
	GetUuid() *string
	SetVersion(v int32) *NotifyPolicy
	GetVersion() *int32
	SetWorkspace(v string) *NotifyPolicy
	GetWorkspace() *string
}

type NotifyPolicy struct {
	// The creation time. The value is a UNIX timestamp string in milliseconds.
	//
	// example:
	//
	// 1710000000000
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The policy description.
	//
	// example:
	//
	// 生产环境告警通知策略
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the policy is enabled. This is a read-only field controlled by the Enable or Disable operation.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The policy name.
	//
	// example:
	//
	// 生产环境告警通知策略
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The notification policy sub-entity details.
	//
	// example:
	//
	// {}
	NotifyStrategy *NotifyStrategyDetail `json:"notifyStrategy,omitempty" xml:"notifyStrategy,omitempty"`
	// The response plan sub-entity details.
	//
	// example:
	//
	// {}
	ResponsePlan *ResponsePlanDetail `json:"responsePlan,omitempty" xml:"responsePlan,omitempty"`
	// The subscription sub-entity details.
	//
	// example:
	//
	// {}
	Subscription *SubscriptionDetail `json:"subscription,omitempty" xml:"subscription,omitempty"`
	// The update time. The value is a UNIX timestamp string in milliseconds.
	//
	// example:
	//
	// 1710000000000
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The Alibaba Cloud account UID.
	//
	// example:
	//
	// 1234567890123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The unique identifier of the policy.
	//
	// example:
	//
	// 04779a183add4f2ca06ab440f16cc580
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// The optimistic locking version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
	// The workspace identifier.
	//
	// example:
	//
	// default-cms-xxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s NotifyPolicy) String() string {
	return dara.Prettify(s)
}

func (s NotifyPolicy) GoString() string {
	return s.String()
}

func (s *NotifyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *NotifyPolicy) GetDescription() *string {
	return s.Description
}

func (s *NotifyPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *NotifyPolicy) GetName() *string {
	return s.Name
}

func (s *NotifyPolicy) GetNotifyStrategy() *NotifyStrategyDetail {
	return s.NotifyStrategy
}

func (s *NotifyPolicy) GetResponsePlan() *ResponsePlanDetail {
	return s.ResponsePlan
}

func (s *NotifyPolicy) GetSubscription() *SubscriptionDetail {
	return s.Subscription
}

func (s *NotifyPolicy) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *NotifyPolicy) GetUserId() *string {
	return s.UserId
}

func (s *NotifyPolicy) GetUuid() *string {
	return s.Uuid
}

func (s *NotifyPolicy) GetVersion() *int32 {
	return s.Version
}

func (s *NotifyPolicy) GetWorkspace() *string {
	return s.Workspace
}

func (s *NotifyPolicy) SetCreateTime(v string) *NotifyPolicy {
	s.CreateTime = &v
	return s
}

func (s *NotifyPolicy) SetDescription(v string) *NotifyPolicy {
	s.Description = &v
	return s
}

func (s *NotifyPolicy) SetEnabled(v bool) *NotifyPolicy {
	s.Enabled = &v
	return s
}

func (s *NotifyPolicy) SetName(v string) *NotifyPolicy {
	s.Name = &v
	return s
}

func (s *NotifyPolicy) SetNotifyStrategy(v *NotifyStrategyDetail) *NotifyPolicy {
	s.NotifyStrategy = v
	return s
}

func (s *NotifyPolicy) SetResponsePlan(v *ResponsePlanDetail) *NotifyPolicy {
	s.ResponsePlan = v
	return s
}

func (s *NotifyPolicy) SetSubscription(v *SubscriptionDetail) *NotifyPolicy {
	s.Subscription = v
	return s
}

func (s *NotifyPolicy) SetUpdateTime(v string) *NotifyPolicy {
	s.UpdateTime = &v
	return s
}

func (s *NotifyPolicy) SetUserId(v string) *NotifyPolicy {
	s.UserId = &v
	return s
}

func (s *NotifyPolicy) SetUuid(v string) *NotifyPolicy {
	s.Uuid = &v
	return s
}

func (s *NotifyPolicy) SetVersion(v int32) *NotifyPolicy {
	s.Version = &v
	return s
}

func (s *NotifyPolicy) SetWorkspace(v string) *NotifyPolicy {
	s.Workspace = &v
	return s
}

func (s *NotifyPolicy) Validate() error {
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
	return nil
}
