// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSubscriptionAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDlqPolicy(v *SetSubscriptionAttributesRequestDlqPolicy) *SetSubscriptionAttributesRequest
	GetDlqPolicy() *SetSubscriptionAttributesRequestDlqPolicy
	SetNotifyStrategy(v string) *SetSubscriptionAttributesRequest
	GetNotifyStrategy() *string
	SetSubscriptionName(v string) *SetSubscriptionAttributesRequest
	GetSubscriptionName() *string
	SetTenantRateLimitPolicy(v *SetSubscriptionAttributesRequestTenantRateLimitPolicy) *SetSubscriptionAttributesRequest
	GetTenantRateLimitPolicy() *SetSubscriptionAttributesRequestTenantRateLimitPolicy
	SetTopicName(v string) *SetSubscriptionAttributesRequest
	GetTopicName() *string
}

type SetSubscriptionAttributesRequest struct {
	// The dead-letter queue policy.
	DlqPolicy *SetSubscriptionAttributesRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
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
	SubscriptionName      *string                                                `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	TenantRateLimitPolicy *SetSubscriptionAttributesRequestTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SetSubscriptionAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSubscriptionAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesRequest) GetDlqPolicy() *SetSubscriptionAttributesRequestDlqPolicy {
	return s.DlqPolicy
}

func (s *SetSubscriptionAttributesRequest) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *SetSubscriptionAttributesRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *SetSubscriptionAttributesRequest) GetTenantRateLimitPolicy() *SetSubscriptionAttributesRequestTenantRateLimitPolicy {
	return s.TenantRateLimitPolicy
}

func (s *SetSubscriptionAttributesRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *SetSubscriptionAttributesRequest) SetDlqPolicy(v *SetSubscriptionAttributesRequestDlqPolicy) *SetSubscriptionAttributesRequest {
	s.DlqPolicy = v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetNotifyStrategy(v string) *SetSubscriptionAttributesRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetSubscriptionName(v string) *SetSubscriptionAttributesRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetTenantRateLimitPolicy(v *SetSubscriptionAttributesRequestTenantRateLimitPolicy) *SetSubscriptionAttributesRequest {
	s.TenantRateLimitPolicy = v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetTopicName(v string) *SetSubscriptionAttributesRequest {
	s.TopicName = &v
	return s
}

func (s *SetSubscriptionAttributesRequest) Validate() error {
	return dara.Validate(s)
}

type SetSubscriptionAttributesRequestDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable the dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s SetSubscriptionAttributesRequestDlqPolicy) String() string {
	return dara.Prettify(s)
}

func (s SetSubscriptionAttributesRequestDlqPolicy) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesRequestDlqPolicy) GetDeadLetterTargetQueue() *string {
	return s.DeadLetterTargetQueue
}

func (s *SetSubscriptionAttributesRequestDlqPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *SetSubscriptionAttributesRequestDlqPolicy) SetDeadLetterTargetQueue(v string) *SetSubscriptionAttributesRequestDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *SetSubscriptionAttributesRequestDlqPolicy) SetEnabled(v bool) *SetSubscriptionAttributesRequestDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *SetSubscriptionAttributesRequestDlqPolicy) Validate() error {
	return dara.Validate(s)
}

type SetSubscriptionAttributesRequestTenantRateLimitPolicy struct {
	Enabled              *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	MaxReceivesPerSecond *int32 `json:"MaxReceivesPerSecond,omitempty" xml:"MaxReceivesPerSecond,omitempty"`
}

func (s SetSubscriptionAttributesRequestTenantRateLimitPolicy) String() string {
	return dara.Prettify(s)
}

func (s SetSubscriptionAttributesRequestTenantRateLimitPolicy) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesRequestTenantRateLimitPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *SetSubscriptionAttributesRequestTenantRateLimitPolicy) GetMaxReceivesPerSecond() *int32 {
	return s.MaxReceivesPerSecond
}

func (s *SetSubscriptionAttributesRequestTenantRateLimitPolicy) SetEnabled(v bool) *SetSubscriptionAttributesRequestTenantRateLimitPolicy {
	s.Enabled = &v
	return s
}

func (s *SetSubscriptionAttributesRequestTenantRateLimitPolicy) SetMaxReceivesPerSecond(v int32) *SetSubscriptionAttributesRequestTenantRateLimitPolicy {
	s.MaxReceivesPerSecond = &v
	return s
}

func (s *SetSubscriptionAttributesRequestTenantRateLimitPolicy) Validate() error {
	return dara.Validate(s)
}
