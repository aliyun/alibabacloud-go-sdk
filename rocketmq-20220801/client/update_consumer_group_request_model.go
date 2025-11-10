// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumeRetryPolicy(v *UpdateConsumerGroupRequestConsumeRetryPolicy) *UpdateConsumerGroupRequest
	GetConsumeRetryPolicy() *UpdateConsumerGroupRequestConsumeRetryPolicy
	SetDeliveryOrderType(v string) *UpdateConsumerGroupRequest
	GetDeliveryOrderType() *string
	SetMaxReceiveTps(v int64) *UpdateConsumerGroupRequest
	GetMaxReceiveTps() *int64
	SetRemark(v string) *UpdateConsumerGroupRequest
	GetRemark() *string
}

type UpdateConsumerGroupRequest struct {
	// The new consumption retry policy of the consumer group. For more information, see [Consumption retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// This parameter is required.
	ConsumeRetryPolicy *UpdateConsumerGroupRequestConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	// The new message delivery method of the consumer group.
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
	// 100
	MaxReceiveTps *int64 `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The new description of the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateConsumerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequest) GetConsumeRetryPolicy() *UpdateConsumerGroupRequestConsumeRetryPolicy {
	return s.ConsumeRetryPolicy
}

func (s *UpdateConsumerGroupRequest) GetDeliveryOrderType() *string {
	return s.DeliveryOrderType
}

func (s *UpdateConsumerGroupRequest) GetMaxReceiveTps() *int64 {
	return s.MaxReceiveTps
}

func (s *UpdateConsumerGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateConsumerGroupRequest) SetConsumeRetryPolicy(v *UpdateConsumerGroupRequestConsumeRetryPolicy) *UpdateConsumerGroupRequest {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *UpdateConsumerGroupRequest) SetDeliveryOrderType(v string) *UpdateConsumerGroupRequest {
	s.DeliveryOrderType = &v
	return s
}

func (s *UpdateConsumerGroupRequest) SetMaxReceiveTps(v int64) *UpdateConsumerGroupRequest {
	s.MaxReceiveTps = &v
	return s
}

func (s *UpdateConsumerGroupRequest) SetRemark(v string) *UpdateConsumerGroupRequest {
	s.Remark = &v
	return s
}

func (s *UpdateConsumerGroupRequest) Validate() error {
	if s.ConsumeRetryPolicy != nil {
		if err := s.ConsumeRetryPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateConsumerGroupRequestConsumeRetryPolicy struct {
	// The dead-letter topic.
	//
	// If a message still fails to be consumed after the maximum number of retries specified in the consumption retry policy is reached, the message is delivered to the dead-letter topic for subsequent business recovery or backtracking. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// Fixed retry interval, unit: seconds.This option is effective when retryPolicy is FixedRetryPolicy.Value range：
	//
	//   - Concurrently:10-1800
	//
	//   - Orderly:1-600
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

func (s UpdateConsumerGroupRequestConsumeRetryPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerGroupRequestConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) GetDeadLetterTargetTopic() *string {
	return s.DeadLetterTargetTopic
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) GetFixedIntervalRetryTime() *int32 {
	return s.FixedIntervalRetryTime
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) GetMaxRetryTimes() *int32 {
	return s.MaxRetryTimes
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) GetRetryPolicy() *string {
	return s.RetryPolicy
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetDeadLetterTargetTopic(v string) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.DeadLetterTargetTopic = &v
	return s
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetFixedIntervalRetryTime(v int32) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.FixedIntervalRetryTime = &v
	return s
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetMaxRetryTimes(v int32) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetRetryPolicy(v string) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) Validate() error {
	return dara.Validate(s)
}
