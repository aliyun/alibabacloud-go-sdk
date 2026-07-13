// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDlqPolicy(v *SubscribeRequestDlqPolicy) *SubscribeRequest
	GetDlqPolicy() *SubscribeRequestDlqPolicy
	SetDmAttributes(v *SubscribeRequestDmAttributes) *SubscribeRequest
	GetDmAttributes() *SubscribeRequestDmAttributes
	SetDysmsAttributes(v *SubscribeRequestDysmsAttributes) *SubscribeRequest
	GetDysmsAttributes() *SubscribeRequestDysmsAttributes
	SetEndpoint(v string) *SubscribeRequest
	GetEndpoint() *string
	SetKafkaAttributes(v *SubscribeRequestKafkaAttributes) *SubscribeRequest
	GetKafkaAttributes() *SubscribeRequestKafkaAttributes
	SetMessageTag(v string) *SubscribeRequest
	GetMessageTag() *string
	SetNotifyContentFormat(v string) *SubscribeRequest
	GetNotifyContentFormat() *string
	SetNotifyStrategy(v string) *SubscribeRequest
	GetNotifyStrategy() *string
	SetPushType(v string) *SubscribeRequest
	GetPushType() *string
	SetStsRoleArn(v string) *SubscribeRequest
	GetStsRoleArn() *string
	SetSubscriptionName(v string) *SubscribeRequest
	GetSubscriptionName() *string
	SetTenantRateLimitPolicy(v *SubscribeRequestTenantRateLimitPolicy) *SubscribeRequest
	GetTenantRateLimitPolicy() *SubscribeRequestTenantRateLimitPolicy
	SetTopicName(v string) *SubscribeRequest
	GetTopicName() *string
}

type SubscribeRequest struct {
	// The dead-letter queue policy.
	DlqPolicy *SubscribeRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// The email push attributes. This parameter is required when PushType is set to dm. The value is in JSON format and contains the following fields:
	//
	// - AccountName: The sender address configured in DirectMail (such as notify@example.com).
	//
	// - Subject: The email subject.
	DmAttributes *SubscribeRequestDmAttributes `json:"DmAttributes,omitempty" xml:"DmAttributes,omitempty" type:"Struct"`
	// The SMS push attributes. This parameter is required when PushType is set to alisms. The value is in JSON format and contains the following fields:
	//
	// - TemplateCode: The SMS template code, which can be obtained from the Short Message Service console.
	//
	// - SignName: The SMS signature name.
	DysmsAttributes *SubscribeRequestDysmsAttributes `json:"DysmsAttributes,omitempty" xml:"DysmsAttributes,omitempty" type:"Struct"`
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
	KafkaAttributes *SubscribeRequestKafkaAttributes `json:"KafkaAttributes,omitempty" xml:"KafkaAttributes,omitempty" type:"Struct"`
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
	TenantRateLimitPolicy *SubscribeRequestTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// topic****1
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SubscribeRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscribeRequest) GoString() string {
	return s.String()
}

func (s *SubscribeRequest) GetDlqPolicy() *SubscribeRequestDlqPolicy {
	return s.DlqPolicy
}

func (s *SubscribeRequest) GetDmAttributes() *SubscribeRequestDmAttributes {
	return s.DmAttributes
}

func (s *SubscribeRequest) GetDysmsAttributes() *SubscribeRequestDysmsAttributes {
	return s.DysmsAttributes
}

func (s *SubscribeRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SubscribeRequest) GetKafkaAttributes() *SubscribeRequestKafkaAttributes {
	return s.KafkaAttributes
}

func (s *SubscribeRequest) GetMessageTag() *string {
	return s.MessageTag
}

func (s *SubscribeRequest) GetNotifyContentFormat() *string {
	return s.NotifyContentFormat
}

func (s *SubscribeRequest) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *SubscribeRequest) GetPushType() *string {
	return s.PushType
}

func (s *SubscribeRequest) GetStsRoleArn() *string {
	return s.StsRoleArn
}

func (s *SubscribeRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *SubscribeRequest) GetTenantRateLimitPolicy() *SubscribeRequestTenantRateLimitPolicy {
	return s.TenantRateLimitPolicy
}

func (s *SubscribeRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *SubscribeRequest) SetDlqPolicy(v *SubscribeRequestDlqPolicy) *SubscribeRequest {
	s.DlqPolicy = v
	return s
}

func (s *SubscribeRequest) SetDmAttributes(v *SubscribeRequestDmAttributes) *SubscribeRequest {
	s.DmAttributes = v
	return s
}

func (s *SubscribeRequest) SetDysmsAttributes(v *SubscribeRequestDysmsAttributes) *SubscribeRequest {
	s.DysmsAttributes = v
	return s
}

func (s *SubscribeRequest) SetEndpoint(v string) *SubscribeRequest {
	s.Endpoint = &v
	return s
}

func (s *SubscribeRequest) SetKafkaAttributes(v *SubscribeRequestKafkaAttributes) *SubscribeRequest {
	s.KafkaAttributes = v
	return s
}

func (s *SubscribeRequest) SetMessageTag(v string) *SubscribeRequest {
	s.MessageTag = &v
	return s
}

func (s *SubscribeRequest) SetNotifyContentFormat(v string) *SubscribeRequest {
	s.NotifyContentFormat = &v
	return s
}

func (s *SubscribeRequest) SetNotifyStrategy(v string) *SubscribeRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SubscribeRequest) SetPushType(v string) *SubscribeRequest {
	s.PushType = &v
	return s
}

func (s *SubscribeRequest) SetStsRoleArn(v string) *SubscribeRequest {
	s.StsRoleArn = &v
	return s
}

func (s *SubscribeRequest) SetSubscriptionName(v string) *SubscribeRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SubscribeRequest) SetTenantRateLimitPolicy(v *SubscribeRequestTenantRateLimitPolicy) *SubscribeRequest {
	s.TenantRateLimitPolicy = v
	return s
}

func (s *SubscribeRequest) SetTopicName(v string) *SubscribeRequest {
	s.TopicName = &v
	return s
}

func (s *SubscribeRequest) Validate() error {
	if s.DlqPolicy != nil {
		if err := s.DlqPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.DmAttributes != nil {
		if err := s.DmAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.DysmsAttributes != nil {
		if err := s.DysmsAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.KafkaAttributes != nil {
		if err := s.KafkaAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.TenantRateLimitPolicy != nil {
		if err := s.TenantRateLimitPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubscribeRequestDlqPolicy struct {
	// The destination queue for dead-letter message delivery.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable dead-letter message delivery.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s SubscribeRequestDlqPolicy) String() string {
	return dara.Prettify(s)
}

func (s SubscribeRequestDlqPolicy) GoString() string {
	return s.String()
}

func (s *SubscribeRequestDlqPolicy) GetDeadLetterTargetQueue() *string {
	return s.DeadLetterTargetQueue
}

func (s *SubscribeRequestDlqPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *SubscribeRequestDlqPolicy) SetDeadLetterTargetQueue(v string) *SubscribeRequestDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *SubscribeRequestDlqPolicy) SetEnabled(v bool) *SubscribeRequestDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *SubscribeRequestDlqPolicy) Validate() error {
	return dara.Validate(s)
}

type SubscribeRequestDmAttributes struct {
	// The sender address.
	//
	// example:
	//
	// notify@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The email subject.
	//
	// example:
	//
	// notify
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s SubscribeRequestDmAttributes) String() string {
	return dara.Prettify(s)
}

func (s SubscribeRequestDmAttributes) GoString() string {
	return s.String()
}

func (s *SubscribeRequestDmAttributes) GetAccountName() *string {
	return s.AccountName
}

func (s *SubscribeRequestDmAttributes) GetSubject() *string {
	return s.Subject
}

func (s *SubscribeRequestDmAttributes) SetAccountName(v string) *SubscribeRequestDmAttributes {
	s.AccountName = &v
	return s
}

func (s *SubscribeRequestDmAttributes) SetSubject(v string) *SubscribeRequestDmAttributes {
	s.Subject = &v
	return s
}

func (s *SubscribeRequestDmAttributes) Validate() error {
	return dara.Validate(s)
}

type SubscribeRequestDysmsAttributes struct {
	// The SMS signature name.
	//
	// example:
	//
	// 阿里云短信测试专用
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The SMS template code.
	//
	// example:
	//
	// 123456
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s SubscribeRequestDysmsAttributes) String() string {
	return dara.Prettify(s)
}

func (s SubscribeRequestDysmsAttributes) GoString() string {
	return s.String()
}

func (s *SubscribeRequestDysmsAttributes) GetSignName() *string {
	return s.SignName
}

func (s *SubscribeRequestDysmsAttributes) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SubscribeRequestDysmsAttributes) SetSignName(v string) *SubscribeRequestDysmsAttributes {
	s.SignName = &v
	return s
}

func (s *SubscribeRequestDysmsAttributes) SetTemplateCode(v string) *SubscribeRequestDysmsAttributes {
	s.TemplateCode = &v
	return s
}

func (s *SubscribeRequestDysmsAttributes) Validate() error {
	return dara.Validate(s)
}

type SubscribeRequestKafkaAttributes struct {
	// The Kafka push type is deprecated.
	//
	// example:
	//
	// Default empty string
	BusinessMode *string `json:"BusinessMode,omitempty" xml:"BusinessMode,omitempty"`
}

func (s SubscribeRequestKafkaAttributes) String() string {
	return dara.Prettify(s)
}

func (s SubscribeRequestKafkaAttributes) GoString() string {
	return s.String()
}

func (s *SubscribeRequestKafkaAttributes) GetBusinessMode() *string {
	return s.BusinessMode
}

func (s *SubscribeRequestKafkaAttributes) SetBusinessMode(v string) *SubscribeRequestKafkaAttributes {
	s.BusinessMode = &v
	return s
}

func (s *SubscribeRequestKafkaAttributes) Validate() error {
	return dara.Validate(s)
}

type SubscribeRequestTenantRateLimitPolicy struct {
	// Specifies whether to enable the throttling policy. Valid values: true and false.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of pushes or consumptions per second.
	//
	// example:
	//
	// 50
	MaxReceivesPerSecond *int32 `json:"MaxReceivesPerSecond,omitempty" xml:"MaxReceivesPerSecond,omitempty"`
}

func (s SubscribeRequestTenantRateLimitPolicy) String() string {
	return dara.Prettify(s)
}

func (s SubscribeRequestTenantRateLimitPolicy) GoString() string {
	return s.String()
}

func (s *SubscribeRequestTenantRateLimitPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *SubscribeRequestTenantRateLimitPolicy) GetMaxReceivesPerSecond() *int32 {
	return s.MaxReceivesPerSecond
}

func (s *SubscribeRequestTenantRateLimitPolicy) SetEnabled(v bool) *SubscribeRequestTenantRateLimitPolicy {
	s.Enabled = &v
	return s
}

func (s *SubscribeRequestTenantRateLimitPolicy) SetMaxReceivesPerSecond(v int32) *SubscribeRequestTenantRateLimitPolicy {
	s.MaxReceivesPerSecond = &v
	return s
}

func (s *SubscribeRequestTenantRateLimitPolicy) Validate() error {
	return dara.Validate(s)
}
