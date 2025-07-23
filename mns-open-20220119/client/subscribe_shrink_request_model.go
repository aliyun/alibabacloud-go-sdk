// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDlqPolicyShrink(v string) *SubscribeShrinkRequest
	GetDlqPolicyShrink() *string
	SetDmAttributesShrink(v string) *SubscribeShrinkRequest
	GetDmAttributesShrink() *string
	SetDysmsAttributesShrink(v string) *SubscribeShrinkRequest
	GetDysmsAttributesShrink() *string
	SetEndpoint(v string) *SubscribeShrinkRequest
	GetEndpoint() *string
	SetKafkaAttributesShrink(v string) *SubscribeShrinkRequest
	GetKafkaAttributesShrink() *string
	SetMessageTag(v string) *SubscribeShrinkRequest
	GetMessageTag() *string
	SetNotifyContentFormat(v string) *SubscribeShrinkRequest
	GetNotifyContentFormat() *string
	SetNotifyStrategy(v string) *SubscribeShrinkRequest
	GetNotifyStrategy() *string
	SetPushType(v string) *SubscribeShrinkRequest
	GetPushType() *string
	SetStsRoleArn(v string) *SubscribeShrinkRequest
	GetStsRoleArn() *string
	SetSubscriptionName(v string) *SubscribeShrinkRequest
	GetSubscriptionName() *string
	SetTenantRateLimitPolicyShrink(v string) *SubscribeShrinkRequest
	GetTenantRateLimitPolicyShrink() *string
	SetTopicName(v string) *SubscribeShrinkRequest
	GetTopicName() *string
}

type SubscribeShrinkRequest struct {
	// The dead-letter queue policy.
	DlqPolicyShrink       *string `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty"`
	DmAttributesShrink    *string `json:"DmAttributes,omitempty" xml:"DmAttributes,omitempty"`
	DysmsAttributesShrink *string `json:"DysmsAttributes,omitempty" xml:"DysmsAttributes,omitempty"`
	// The receiver endpoint. The format of the endpoint varies based on the terminal type.
	//
	// 	- If you set PushType to http, set Endpoint to an `HTTP URL that starts with http:// or https://`.
	//
	// 	- If you set PushType to queue, set Endpoint to a `queue name`.
	//
	// 	- If you set PushType to mpush, set Endpoint to an `AppKey`.
	//
	// 	- If you set PushType to alisms, set Endpoint to a `mobile number`.
	//
	// 	- If you set PushType to email, set Endpoint to an `email address`.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com
	Endpoint              *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	KafkaAttributesShrink *string `json:"KafkaAttributes,omitempty" xml:"KafkaAttributes,omitempty"`
	// The tag that is used to filter messages. Only messages that have the same tag can be pushed. Set the value to a string of no more than 16 characters.
	//
	// By default, no tag is specified to filter messages.
	//
	// example:
	//
	// important
	MessageTag *string `json:"MessageTag,omitempty" xml:"MessageTag,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
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
	// The terminal type. Valid values:
	//
	// 	- http: HTTP services
	//
	// 	- queue: queues
	//
	// 	- mpush: mobile devices
	//
	// 	- alisms: Alibaba Cloud Short Message Service (SMS)
	//
	// 	- email: emails
	//
	// This parameter is required.
	//
	// example:
	//
	// queue
	PushType   *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	StsRoleArn *string `json:"StsRoleArn,omitempty" xml:"StsRoleArn,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSubscription
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

func (s SubscribeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscribeShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubscribeShrinkRequest) GetDlqPolicyShrink() *string {
	return s.DlqPolicyShrink
}

func (s *SubscribeShrinkRequest) GetDmAttributesShrink() *string {
	return s.DmAttributesShrink
}

func (s *SubscribeShrinkRequest) GetDysmsAttributesShrink() *string {
	return s.DysmsAttributesShrink
}

func (s *SubscribeShrinkRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SubscribeShrinkRequest) GetKafkaAttributesShrink() *string {
	return s.KafkaAttributesShrink
}

func (s *SubscribeShrinkRequest) GetMessageTag() *string {
	return s.MessageTag
}

func (s *SubscribeShrinkRequest) GetNotifyContentFormat() *string {
	return s.NotifyContentFormat
}

func (s *SubscribeShrinkRequest) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *SubscribeShrinkRequest) GetPushType() *string {
	return s.PushType
}

func (s *SubscribeShrinkRequest) GetStsRoleArn() *string {
	return s.StsRoleArn
}

func (s *SubscribeShrinkRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *SubscribeShrinkRequest) GetTenantRateLimitPolicyShrink() *string {
	return s.TenantRateLimitPolicyShrink
}

func (s *SubscribeShrinkRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *SubscribeShrinkRequest) SetDlqPolicyShrink(v string) *SubscribeShrinkRequest {
	s.DlqPolicyShrink = &v
	return s
}

func (s *SubscribeShrinkRequest) SetDmAttributesShrink(v string) *SubscribeShrinkRequest {
	s.DmAttributesShrink = &v
	return s
}

func (s *SubscribeShrinkRequest) SetDysmsAttributesShrink(v string) *SubscribeShrinkRequest {
	s.DysmsAttributesShrink = &v
	return s
}

func (s *SubscribeShrinkRequest) SetEndpoint(v string) *SubscribeShrinkRequest {
	s.Endpoint = &v
	return s
}

func (s *SubscribeShrinkRequest) SetKafkaAttributesShrink(v string) *SubscribeShrinkRequest {
	s.KafkaAttributesShrink = &v
	return s
}

func (s *SubscribeShrinkRequest) SetMessageTag(v string) *SubscribeShrinkRequest {
	s.MessageTag = &v
	return s
}

func (s *SubscribeShrinkRequest) SetNotifyContentFormat(v string) *SubscribeShrinkRequest {
	s.NotifyContentFormat = &v
	return s
}

func (s *SubscribeShrinkRequest) SetNotifyStrategy(v string) *SubscribeShrinkRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SubscribeShrinkRequest) SetPushType(v string) *SubscribeShrinkRequest {
	s.PushType = &v
	return s
}

func (s *SubscribeShrinkRequest) SetStsRoleArn(v string) *SubscribeShrinkRequest {
	s.StsRoleArn = &v
	return s
}

func (s *SubscribeShrinkRequest) SetSubscriptionName(v string) *SubscribeShrinkRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SubscribeShrinkRequest) SetTenantRateLimitPolicyShrink(v string) *SubscribeShrinkRequest {
	s.TenantRateLimitPolicyShrink = &v
	return s
}

func (s *SubscribeShrinkRequest) SetTopicName(v string) *SubscribeShrinkRequest {
	s.TopicName = &v
	return s
}

func (s *SubscribeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
