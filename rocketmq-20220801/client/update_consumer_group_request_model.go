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
	// The updated consumption retry policy of the consumer group. For more information, see [Consumption retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// This parameter is required.
	ConsumeRetryPolicy *UpdateConsumerGroupRequestConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	// Deprecated
	//
	// Modification is not supported.
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	// The upper limit of TPS for consuming messages.
	//
	// example:
	//
	// 1000
	MaxReceiveTps *int64 `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The updated remarks of the consumer group.
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
	// If a consumer fails to consume a message and the message still fails after retries, the message is delivered to the dead-letter topic for subsequent business recovery or review. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// The fixed retry interval. Unit: seconds. This parameter takes effect only when the retry policy is set to FixedRetryPolicy (fixed-interval retry).
	//
	// Valid values:
	//
	//   - Concurrent delivery: 10-1800
	//
	//   - Ordered delivery: 1-600
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
	// The retry policy type. For more information, see [Message retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// Valid values:
	//
	// - FixedRetryPolicy: fixed-interval retry. Only ordered delivery is supported.
	//
	// - DefaultRetryPolicy: incremental backoff retry. Only concurrent delivery is supported.
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
