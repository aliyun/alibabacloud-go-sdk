// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSubscriptionAttributesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDlqPolicyShrink(v string) *SetSubscriptionAttributesShrinkRequest
	GetDlqPolicyShrink() *string
	SetNotifyStrategy(v string) *SetSubscriptionAttributesShrinkRequest
	GetNotifyStrategy() *string
	SetSubscriptionName(v string) *SetSubscriptionAttributesShrinkRequest
	GetSubscriptionName() *string
	SetTenantRateLimitPolicyShrink(v string) *SetSubscriptionAttributesShrinkRequest
	GetTenantRateLimitPolicyShrink() *string
	SetTopicName(v string) *SetSubscriptionAttributesShrinkRequest
	GetTopicName() *string
}

type SetSubscriptionAttributesShrinkRequest struct {
	// The dead-letter queue policy.
	DlqPolicyShrink *string `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName            *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	TenantRateLimitPolicyShrink *string `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SetSubscriptionAttributesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSubscriptionAttributesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesShrinkRequest) GetDlqPolicyShrink() *string {
	return s.DlqPolicyShrink
}

func (s *SetSubscriptionAttributesShrinkRequest) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *SetSubscriptionAttributesShrinkRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *SetSubscriptionAttributesShrinkRequest) GetTenantRateLimitPolicyShrink() *string {
	return s.TenantRateLimitPolicyShrink
}

func (s *SetSubscriptionAttributesShrinkRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *SetSubscriptionAttributesShrinkRequest) SetDlqPolicyShrink(v string) *SetSubscriptionAttributesShrinkRequest {
	s.DlqPolicyShrink = &v
	return s
}

func (s *SetSubscriptionAttributesShrinkRequest) SetNotifyStrategy(v string) *SetSubscriptionAttributesShrinkRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SetSubscriptionAttributesShrinkRequest) SetSubscriptionName(v string) *SetSubscriptionAttributesShrinkRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SetSubscriptionAttributesShrinkRequest) SetTenantRateLimitPolicyShrink(v string) *SetSubscriptionAttributesShrinkRequest {
	s.TenantRateLimitPolicyShrink = &v
	return s
}

func (s *SetSubscriptionAttributesShrinkRequest) SetTopicName(v string) *SetSubscriptionAttributesShrinkRequest {
	s.TopicName = &v
	return s
}

func (s *SetSubscriptionAttributesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
