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
	SetMaxReceiveTps(v int64) *CreateConsumerGroupRequest
	GetMaxReceiveTps() *int64
	SetRemark(v string) *CreateConsumerGroupRequest
	GetRemark() *string
}

type CreateConsumerGroupRequest struct {
	// The consumption retry policy of the consumer group. For more information, see [Consumption retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// This parameter is required.
	ConsumeRetryPolicy *CreateConsumerGroupRequestConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	// The message delivery method of the consumer group.
	//
	// Valid values:
	//
	// 	- Concurrently: concurrent delivery
	//
	// 	- Orderly: ordered delivery
	//
	// This parameter is required.
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	// The maximum number of messages that can be processed by consumers per second.
	//
	// example:
	//
	// 300
	MaxReceiveTps *int64 `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The description of the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
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

func (s *CreateConsumerGroupRequest) GetMaxReceiveTps() *int64 {
	return s.MaxReceiveTps
}

func (s *CreateConsumerGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateConsumerGroupRequest) SetConsumeRetryPolicy(v *CreateConsumerGroupRequestConsumeRetryPolicy) *CreateConsumerGroupRequest {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *CreateConsumerGroupRequest) SetDeliveryOrderType(v string) *CreateConsumerGroupRequest {
	s.DeliveryOrderType = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetMaxReceiveTps(v int64) *CreateConsumerGroupRequest {
	s.MaxReceiveTps = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRemark(v string) *CreateConsumerGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateConsumerGroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateConsumerGroupRequestConsumeRetryPolicy struct {
	// The dead-letter topic.
	//
	// If a message still fails to be consumed after the maximum number of retries specified in the consumption retry policy is reached, the message is delivered to the dead-letter topic for subsequent business recovery or backtracking. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// The fixed interval. Unit: seconds. This parameter takes effect only if you set retryPolicy to FixedRetryPolicy. Valid values:
	//
	// 	- Concurrent delivery: 10 to 1800
	//
	// 	- Ordered delivery: 1 to 600
	//
	// example:
	//
	// 10
	FixedIntervalRetryTime *int32 `json:"fixedIntervalRetryTime,omitempty" xml:"fixedIntervalRetryTime,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 16
	MaxRetryTimes *int32 `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	// The retry policy. For more information, see [Message retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// Valid values:
	//
	// 	- FixedRetryPolicy: fixed-interval retry. This value is valid only if you set deliveryOrderType to Orderly.
	//
	// 	- DefaultRetryPolicy: exponential backoff retry. This value is valid only if you set deliveryOrderType to Concurrently.
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
