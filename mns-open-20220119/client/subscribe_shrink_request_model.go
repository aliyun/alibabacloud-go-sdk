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
	DlqPolicyShrink *string `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty"`
	// The email push attributes. This parameter is required when PushType is set to dm. The value is in JSON format and contains the following fields:
	//
	// - AccountName: The sender address configured in DirectMail (such as notify@example.com).
	//
	// - Subject: The email subject.
	DmAttributesShrink *string `json:"DmAttributes,omitempty" xml:"DmAttributes,omitempty"`
	// The SMS push attributes. This parameter is required when PushType is set to alisms. The value is in JSON format and contains the following fields:
	//
	// - TemplateCode: The SMS template code, which can be obtained from the Short Message Service console.
	//
	// - SignName: The SMS signature name.
	DysmsAttributesShrink *string `json:"DysmsAttributes,omitempty" xml:"DysmsAttributes,omitempty"`
	// ## Endpoint address for receiving messages
	//
	// The format varies depending on the value of `PushType`:
	//
	// - `PushType=http`: An HTTP/HTTPS callback URL, such as `http://example.com/callback` or `https://example.com/callback`.
	//
	// - `PushType=queue`: The ARN of the destination queue, in the format `acs:mns:{RegionId}:{Alibaba Cloud account ID}:queues/{QueueName}`.
	//
	// - `PushType=dm`: The email push endpoint, in the fixed format `smq-ep:dm:{Alibaba Cloud account ID}:__dynamic`. Replace `{Alibaba Cloud account ID}` with your Alibaba Cloud account ID.
	//
	// - `PushType=dysms`: The SMS push endpoint, in the format `smq-ep:dysms:{Alibaba Cloud account ID}:{PhoneNumber}`.
	//
	// - `PushType=kafka`: The Kafka push endpoint. The Kafka push type is deprecated.
	//
	// - `PushType=fc`: The Function Compute endpoint, in the format `acs:fc:{RegionId}:{Alibaba Cloud account ID}:services/{ServiceName}/functions/{FunctionName}`.
	//
	// - `PushType=eventbus`: The EventBridge endpoint, in the format `acs:eventbridge:{RegionId}:{Alibaba Cloud account ID}:eventbus/{EventBusName}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://*****.com/uri1/xxx
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The Kafka push type is deprecated.
	KafkaAttributesShrink *string `json:"KafkaAttributes,omitempty" xml:"KafkaAttributes,omitempty"`
	// The tag used for message filtering in this subscription. Only messages with a matching tag are pushed. The value is a string of up to 16 characters.
	//
	// By default, no message filtering is applied.
	//
	// example:
	//
	// important
	MessageTag *string `json:"MessageTag,omitempty" xml:"MessageTag,omitempty"`
	// ## Format of the pushed message content
	//
	// Valid values:
	//
	// - `XML`: The message body is pushed in XML format. This is the default value.
	//
	// - `JSON`: The message body is pushed in JSON format.
	//
	// - `SIMPLIFIED`: Only the raw message body content is pushed, without SMQ metadata wrapping.
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry strategy when an error occurs while pushing messages to the endpoint. Valid values:
	//
	// - BACKOFF_RETRY: backoff retry.
	//
	// - EXPONENTIAL_DECAY_RETRY: exponential decay retry.
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// ## Push type of the subscription
	//
	// Valid values:
	//
	// - `http`: HTTP/HTTPS push. Pushes messages to a specified HTTP or HTTPS callback URL.
	//
	// - `queue`: Queue push. Pushes messages to a specified SMQ queue.
	//
	// - `dm`: Email push. Sends notifications through DirectMail. You must also set the `DmAttributes` and `StsRoleArn` parameters.
	//
	// - `dysms`: SMS push. Sends notifications through Alibaba Cloud Short Message Service. You must also set the `DysmsAttributes` parameter.
	//
	// - `fc`: Function Compute push. Pushes messages to Alibaba Cloud Function Compute (FC).
	//
	// - `eventbus`: EventBridge push. Pushes messages to an EventBridge event bus.
	//
	// **Note:**
	//
	// The following values are deprecated and are only used for compatibility with legacy subscriptions:
	//
	// - `mpush`: Mobile push.
	//
	// - `alisms`: Legacy SMS.
	//
	// - `email`: Legacy email. Use `dm` instead.
	//
	// - `kafka`: Kafka push type is deprecated.
	//
	// This parameter is required.
	//
	// example:
	//
	// queue
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The ARN of the RAM role assumed by the service. The format is acs:ram::{Alibaba Cloud account ID}:role/{RoleName}. Replace {Alibaba Cloud account ID} with the Alibaba Cloud account ID that calls the API operation.
	//
	// example:
	//
	// acs:ram::1234567890:role/AliyunMNSNotificationRole
	StsRoleArn *string `json:"StsRoleArn,omitempty" xml:"StsRoleArn,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The throttling policy.
	TenantRateLimitPolicyShrink *string `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// topic****1
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
