// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumeRetryPolicy(v *CreateConsumerGroupRequestConsumeRetryPolicy) *CreateConsumerGroupRequest
	GetConsumeRetryPolicy() *CreateConsumerGroupRequestConsumeRetryPolicy
	SetDeliveryOrderType(v string) *CreateConsumerGroupRequest
	GetDeliveryOrderType() *string
	SetExclusive(v bool) *CreateConsumerGroupRequest
	GetExclusive() *bool
	SetMaxReceiveTps(v int64) *CreateConsumerGroupRequest
	GetMaxReceiveTps() *int64
	SetMessageModel(v string) *CreateConsumerGroupRequest
	GetMessageModel() *string
	SetRemark(v string) *CreateConsumerGroupRequest
	GetRemark() *string
	SetTopicName(v string) *CreateConsumerGroupRequest
	GetTopicName() *string
}

type CreateConsumerGroupRequest struct {
	// The consumption retry policy of the consumer group to be created. For more information, see [Consumption retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// This parameter is required.
	ConsumeRetryPolicy *CreateConsumerGroupRequestConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	// The delivery order type of the consumer group to be created.
	//
	// Valid parameter values:
	//
	// - Concurrently: concurrent delivery
	//
	// - Orderly: ordered delivery
	//
	// This parameter is required.
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	// Specifies whether to enable exclusive subscription, which means a LiteTopic can only be subscribed to by one consumer. Set to true to enable or false to shutdown.
	Exclusive *bool `json:"exclusive,omitempty" xml:"exclusive,omitempty"`
	// The maximum TPS for consuming messages.
	//
	// example:
	//
	// 1000
	MaxReceiveTps *int64 `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The subscribe pattern for consumers:
	//
	// - CLUSTERING: clustering consumption pattern. All consumers in the ConsumerGroup share the same subscription.
	//
	// - LITE_SELECTIVE: lightweight consumption pattern. Consumers in the ConsumerGroup can have different subscriptions.
	//
	// example:
	//
	// LITE_SELECTIVE
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The remarks for the consumer group to be created.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The name of the lite-type topic to subscribe to when using the LITE_SELECTIVE consumption pattern.
	//
	// example:
	//
	// liteTopicTest
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) GetConsumeRetryPolicy() *CreateConsumerGroupRequestConsumeRetryPolicy {
	return s.ConsumeRetryPolicy
}

func (s *CreateConsumerGroupRequest) GetDeliveryOrderType() *string {
	return s.DeliveryOrderType
}

func (s *CreateConsumerGroupRequest) GetExclusive() *bool {
	return s.Exclusive
}

func (s *CreateConsumerGroupRequest) GetMaxReceiveTps() *int64 {
	return s.MaxReceiveTps
}

func (s *CreateConsumerGroupRequest) GetMessageModel() *string {
	return s.MessageModel
}

func (s *CreateConsumerGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateConsumerGroupRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *CreateConsumerGroupRequest) SetConsumeRetryPolicy(v *CreateConsumerGroupRequestConsumeRetryPolicy) *CreateConsumerGroupRequest {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *CreateConsumerGroupRequest) SetDeliveryOrderType(v string) *CreateConsumerGroupRequest {
	s.DeliveryOrderType = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetExclusive(v bool) *CreateConsumerGroupRequest {
	s.Exclusive = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetMaxReceiveTps(v int64) *CreateConsumerGroupRequest {
	s.MaxReceiveTps = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetMessageModel(v string) *CreateConsumerGroupRequest {
	s.MessageModel = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRemark(v string) *CreateConsumerGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetTopicName(v string) *CreateConsumerGroupRequest {
	s.TopicName = &v
	return s
}

func (s *CreateConsumerGroupRequest) Validate() error {
	if s.ConsumeRetryPolicy != nil {
		if err := s.ConsumeRetryPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConsumerGroupRequestConsumeRetryPolicy struct {
	// The dead-letter topic.
	//
	// If a consumer fails to consume a message and the message still failed after retries, the message is delivered to the dead-letter topic for subsequent business recovery or backtracking. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// The fixed retry interval, in seconds. This parameter takes effect when the retry policy is set to FixedRetryPolicy.
	//
	// Valid values:
	//
	// - Concurrent delivery: 10 to 1800
	//
	// - Ordered delivery: 1 to 600
	//
	// example:
	//
	// 10
	FixedIntervalRetryTime *int32 `json:"fixedIntervalRetryTime,omitempty" xml:"fixedIntervalRetryTime,omitempty"`
	// The maximum number of retry attempts.
	//
	// example:
	//
	// 16
	MaxRetryTimes *int32 `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	// The retry policy type. For more information, see [Message retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// Valid parameter values:
	//
	// - FixedRetryPolicy: fixed-interval retry, supported only for ordered delivery
	//
	// - DefaultRetryPolicy: progressive backoff retry, supported only for concurrent delivery
	//
	// This parameter is required.
	//
	// example:
	//
	// DefaultRetryPolicy
	RetryPolicy *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
}

func (s CreateConsumerGroupRequestConsumeRetryPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupRequestConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) GetDeadLetterTargetTopic() *string {
	return s.DeadLetterTargetTopic
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) GetFixedIntervalRetryTime() *int32 {
	return s.FixedIntervalRetryTime
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) GetMaxRetryTimes() *int32 {
	return s.MaxRetryTimes
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) GetRetryPolicy() *string {
	return s.RetryPolicy
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetDeadLetterTargetTopic(v string) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.DeadLetterTargetTopic = &v
	return s
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetFixedIntervalRetryTime(v int32) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.FixedIntervalRetryTime = &v
	return s
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetMaxRetryTimes(v int32) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetRetryPolicy(v string) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) Validate() error {
	return dara.Validate(s)
}
