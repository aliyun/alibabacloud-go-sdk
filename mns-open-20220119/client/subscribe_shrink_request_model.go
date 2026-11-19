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
	// The attributes for Direct Mail push. This parameter is required when the PushType parameter is set to dm. The value is in JSON format and contains the following fields:
	//
	// - AccountName: the sender address, which is the sender address configured in Direct Mail, such as notify@example.com.
	//
	// - Subject: the email subject.
	DmAttributesShrink *string `json:"DmAttributes,omitempty" xml:"DmAttributes,omitempty"`
	// The attributes for SMS push. This parameter is required when the PushType parameter is set to alisms. The value is in JSON format and contains the following fields:
	//
	// - TemplateCode: the code of the SMS template, which can be obtained from the SMS console.
	//
	// - SignName: the name of the SMS signature.
	DysmsAttributesShrink *string `json:"DysmsAttributes,omitempty" xml:"DysmsAttributes,omitempty"`
	// ## Endpoint for receiving messages
	//
	// The format varies based on the `PushType` parameter:
	//
	// - `PushType=http`: The HTTP or HTTPS callback URL, such as `http://example.com/callback` or `https://example.com/callback`.
	//
	// - `PushType=queue`: The ARN of the destination queue, in the format of `acs:mns:{RegionId}:{Alibaba Cloud account ID}:/queues/{QueueName}`.
	//
	// - `PushType=dm`: The Direct Mail push endpoint, in the fixed format of `smq-ep:dm:{Alibaba Cloud account ID}:__dynamic`. Replace `{Alibaba Cloud account ID}` with the ID of your Alibaba Cloud account.
	//
	// - `PushType=dysms`: The SMS push endpoint, in the format of `smq-ep:dysms:{Alibaba Cloud account ID}:{Phone number}`.
	//
	// - `PushType=kafka`: The Kafka push endpoint. The Kafka push type is deprecated.
	//
	// - `PushType=fc`: The Function Compute endpoint, in the format of `acs:fc:{RegionId}:{Alibaba Cloud account ID}:services/{ServiceName}/functions/{FunctionName}`.
	//
	// - `PushType=eventbus`: The EventBridge endpoint, in the format of `acs:eventbridge:{RegionId}:{Alibaba Cloud account ID}:eventbus/{EventBusName}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://*****.com/uri1/xxx
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The Kafka push type is deprecated.
	KafkaAttributesShrink *string `json:"KafkaAttributes,omitempty" xml:"KafkaAttributes,omitempty"`
	// The tag used for message filtering in the subscription. Only messages with a matching tag are pushed.
	//
	// The value must be a string up to 16 characters in length.
	//
	// By default, message filtering is not performed.
	//
	// example:
	//
	// important
	MessageTag *string `json:"MessageTag,omitempty" xml:"MessageTag,omitempty"`
	// ## Format of the pushed message content
	//
	// Valid values:
	//
	// - `XML`: XML format (default)
	//
	// - `JSON`: JSON format
	//
	// - `SIMPLIFIED`: original message body content without the SMQ metadata wrapper
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry strategy that is used when an error occurs during message pushing to the endpoint. Valid values:
	//
	// - BACKOFF_RETRY: backoff retry (default)
	//
	// - EXPONENTIAL_DECAY_RETRY: exponential decay retry
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// ## Push types for subscriptions
	//
	// Valid values:
	//
	// - `http`: Pushes messages to the specified HTTP or HTTPS callback URL.
	//
	// - `queue`: Pushes messages to the specified Simple Message Queue (SMQ) queue.
	//
	// - `dm`: Sends notifications by using the Direct Mail service. You must also set the `DmAttributes` and `StsRoleArn` parameters.
	//
	// - `dysms`: Sends notifications by using the Alibaba Cloud SMS service. You must also set the `DysmsAttributes` parameter.
	//
	// - `fc`: Pushes messages to Alibaba Cloud Function Compute (FC).
	//
	// - `eventbus`: Pushes messages to an EventBridge event bus.
	//
	// **Note:**
	//
	// The following values are deprecated and are retained only for compatibility with legacy subscriptions:
	//
	// - `mpush`: mobile push
	//
	// - `alisms`: legacy SMS
	//
	// - `email`: legacy email, use `dm` instead
	//
	// - `kafka`: deprecated Kafka push type
	//
	// This parameter is required.
	//
	// example:
	//
	// queue
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The ARN of the RAM role assumed by the service. The format is acs:ram::{Alibaba Cloud account ID}:role/{RoleName}. Replace {Alibaba Cloud account ID} with the ID of the Alibaba Cloud account that calls the API operation.
	//
	// example:
	//
	// acs:ram::1234567890:role/AliyunMNSNotificationRole
	StsRoleArn *string `json:"StsRoleArn,omitempty" xml:"StsRoleArn,omitempty"`
	// The name of the subscription. The name must be 1 to 255 characters in length. It must start with a letter or a digit, and can contain only letters, digits, and hyphens (-). The name cannot start with mns-reserved-.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The throttling policy.
	TenantRateLimitPolicyShrink *string `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty"`
	// The name of the topic. The name must be 1 to 255 characters in length. It must start with a letter or a digit, and can contain only letters, digits, and hyphens (-). The name cannot start with mns-reserved-.
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
