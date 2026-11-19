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
	// The attributes for Direct Mail push. This parameter is required when the PushType parameter is set to dm. The value is in JSON format and contains the following fields:
	//
	// - AccountName: the sender address, which is the sender address configured in Direct Mail, such as notify@example.com.
	//
	// - Subject: the email subject.
	DmAttributes *SubscribeRequestDmAttributes `json:"DmAttributes,omitempty" xml:"DmAttributes,omitempty" type:"Struct"`
	// The attributes for SMS push. This parameter is required when the PushType parameter is set to alisms. The value is in JSON format and contains the following fields:
	//
	// - TemplateCode: the code of the SMS template, which can be obtained from the SMS console.
	//
	// - SignName: the name of the SMS signature.
	DysmsAttributes *SubscribeRequestDysmsAttributes `json:"DysmsAttributes,omitempty" xml:"DysmsAttributes,omitempty" type:"Struct"`
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
	KafkaAttributes *SubscribeRequestKafkaAttributes `json:"KafkaAttributes,omitempty" xml:"KafkaAttributes,omitempty" type:"Struct"`
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
	TenantRateLimitPolicy *SubscribeRequestTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The name of the topic. The name must be 1 to 255 characters in length. It must start with a letter or a digit, and can contain only letters, digits, and hyphens (-). The name cannot start with mns-reserved-.
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
	// The destination queue for dead-letter message delivery. The queue must already exist in the current region.
	//
	// example:
	//
	// deadLetterTargetQueue
	DeadLetterTargetQueue *string `json:"DeadLetterTargetQueue,omitempty" xml:"DeadLetterTargetQueue,omitempty"`
	// Specifies whether to enable dead-letter message delivery. Default value: false.
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
	// The name of the SMS signature.
	//
	// example:
	//
	// Alibaba Cloud SMS test
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The code of the SMS template.
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
	// Specifies whether to enable the throttling policy. Valid values:
	//
	// - true: enabled
	//
	// - false: disabled (default)
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum number of pushes or consumptions per second. The value must be a positive integer and is valid only when the PushType parameter is set to http or https. The specific minimum value varies by region.
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
