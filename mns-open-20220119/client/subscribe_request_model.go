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
	DlqPolicy       *SubscribeRequestDlqPolicy       `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	DmAttributes    *SubscribeRequestDmAttributes    `json:"DmAttributes,omitempty" xml:"DmAttributes,omitempty" type:"Struct"`
	DysmsAttributes *SubscribeRequestDysmsAttributes `json:"DysmsAttributes,omitempty" xml:"DysmsAttributes,omitempty" type:"Struct"`
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
	Endpoint        *string                          `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	KafkaAttributes *SubscribeRequestKafkaAttributes `json:"KafkaAttributes,omitempty" xml:"KafkaAttributes,omitempty" type:"Struct"`
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
	SubscriptionName      *string                                `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	TenantRateLimitPolicy *SubscribeRequestTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
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
	return dara.Validate(s)
}

type SubscribeRequestDlqPolicy struct {
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
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Subject     *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
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
	SignName     *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
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
	Enabled              *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
