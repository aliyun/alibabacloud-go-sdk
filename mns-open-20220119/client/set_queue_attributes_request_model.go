// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetQueueAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelaySeconds(v int64) *SetQueueAttributesRequest
	GetDelaySeconds() *int64
	SetDlqPolicy(v *SetQueueAttributesRequestDlqPolicy) *SetQueueAttributesRequest
	GetDlqPolicy() *SetQueueAttributesRequestDlqPolicy
	SetEnableLogging(v bool) *SetQueueAttributesRequest
	GetEnableLogging() *bool
	SetEnableSSE(v bool) *SetQueueAttributesRequest
	GetEnableSSE() *bool
	SetKmsKeyId(v string) *SetQueueAttributesRequest
	GetKmsKeyId() *string
	SetMaximumMessageSize(v int64) *SetQueueAttributesRequest
	GetMaximumMessageSize() *int64
	SetMessageRetentionPeriod(v int64) *SetQueueAttributesRequest
	GetMessageRetentionPeriod() *int64
	SetPollingWaitSeconds(v int64) *SetQueueAttributesRequest
	GetPollingWaitSeconds() *int64
	SetQueueName(v string) *SetQueueAttributesRequest
	GetQueueName() *string
	SetSseAlgorithm(v string) *SetQueueAttributesRequest
	GetSseAlgorithm() *string
	SetSseType(v string) *SetQueueAttributesRequest
	GetSseType() *string
	SetTenantRateLimitPolicy(v *SetQueueAttributesRequestTenantRateLimitPolicy) *SetQueueAttributesRequest
	GetTenantRateLimitPolicy() *SetQueueAttributesRequestTenantRateLimitPolicy
	SetVisibilityTimeout(v int64) *SetQueueAttributesRequest
	GetVisibilityTimeout() *int64
}

type SetQueueAttributesRequest struct {
	// The delay time for all messages sent to this queue. Messages sent to the queue can be consumed only after the delay time specified by this parameter has elapsed.
	//
	// Valid values: 0 to 604800. Unit: seconds.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *SetQueueAttributesRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// Specifies whether to enable the log management feature. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableLogging *bool   `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	EnableSSE     *bool   `json:"EnableSSE,omitempty" xml:"EnableSSE,omitempty"`
	KmsKeyId      *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The maximum length of the message body sent to this queue.
	//
	// Valid values: 1024 to 65536. Unit: bytes.
	//
	// Default value: 65536.
	//
	// example:
	//
	// 1024
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in this queue. After the time specified by this parameter has elapsed since the message was sent to the queue, the message is deleted regardless of whether it has been consumed.
	//
	// Valid values: 60 to 604800. Unit: seconds.
	//
	// Default value: 345600.
	//
	// example:
	//
	// 120
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum wait time for a ReceiveMessage request on this queue when no messages are available in the queue.
	//
	// Valid values: 0 to 30. Unit: seconds.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// testqueue
	QueueName             *string                                         `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	SseAlgorithm          *string                                         `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	SseType               *string                                         `json:"SseType,omitempty" xml:"SseType,omitempty"`
	TenantRateLimitPolicy *SetQueueAttributesRequestTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The duration for which a message stays in the Inactive state after it is consumed from the queue and changes from the Active state to the Inactive state.
	//
	// Valid values: 1 to 43200. Unit: seconds.
	//
	// Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s SetQueueAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetQueueAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesRequest) GetDelaySeconds() *int64 {
	return s.DelaySeconds
}

func (s *SetQueueAttributesRequest) GetDlqPolicy() *SetQueueAttributesRequestDlqPolicy {
	return s.DlqPolicy
}

func (s *SetQueueAttributesRequest) GetEnableLogging() *bool {
	return s.EnableLogging
}

func (s *SetQueueAttributesRequest) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *SetQueueAttributesRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *SetQueueAttributesRequest) GetMaximumMessageSize() *int64 {
	return s.MaximumMessageSize
}

func (s *SetQueueAttributesRequest) GetMessageRetentionPeriod() *int64 {
	return s.MessageRetentionPeriod
}

func (s *SetQueueAttributesRequest) GetPollingWaitSeconds() *int64 {
	return s.PollingWaitSeconds
}

func (s *SetQueueAttributesRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *SetQueueAttributesRequest) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *SetQueueAttributesRequest) GetSseType() *string {
	return s.SseType
}

func (s *SetQueueAttributesRequest) GetTenantRateLimitPolicy() *SetQueueAttributesRequestTenantRateLimitPolicy {
	return s.TenantRateLimitPolicy
}

func (s *SetQueueAttributesRequest) GetVisibilityTimeout() *int64 {
	return s.VisibilityTimeout
}

func (s *SetQueueAttributesRequest) SetDelaySeconds(v int64) *SetQueueAttributesRequest {
	s.DelaySeconds = &v
	return s
}

func (s *SetQueueAttributesRequest) SetDlqPolicy(v *SetQueueAttributesRequestDlqPolicy) *SetQueueAttributesRequest {
	s.DlqPolicy = v
	return s
}

func (s *SetQueueAttributesRequest) SetEnableLogging(v bool) *SetQueueAttributesRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetQueueAttributesRequest) SetEnableSSE(v bool) *SetQueueAttributesRequest {
	s.EnableSSE = &v
	return s
}

func (s *SetQueueAttributesRequest) SetKmsKeyId(v string) *SetQueueAttributesRequest {
	s.KmsKeyId = &v
	return s
}

func (s *SetQueueAttributesRequest) SetMaximumMessageSize(v int64) *SetQueueAttributesRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *SetQueueAttributesRequest) SetMessageRetentionPeriod(v int64) *SetQueueAttributesRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *SetQueueAttributesRequest) SetPollingWaitSeconds(v int64) *SetQueueAttributesRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *SetQueueAttributesRequest) SetQueueName(v string) *SetQueueAttributesRequest {
	s.QueueName = &v
	return s
}

func (s *SetQueueAttributesRequest) SetSseAlgorithm(v string) *SetQueueAttributesRequest {
	s.SseAlgorithm = &v
	return s
}

func (s *SetQueueAttributesRequest) SetSseType(v string) *SetQueueAttributesRequest {
	s.SseType = &v
	return s
}

func (s *SetQueueAttributesRequest) SetTenantRateLimitPolicy(v *SetQueueAttributesRequestTenantRateLimitPolicy) *SetQueueAttributesRequest {
	s.TenantRateLimitPolicy = v
	return s
}

func (s *SetQueueAttributesRequest) SetVisibilityTimeout(v int64) *SetQueueAttributesRequest {
	s.VisibilityTimeout = &v
	return s
}

func (s *SetQueueAttributesRequest) Validate() error {
	if s.DlqPolicy != nil {
		if err := s.DlqPolicy.Validate(); err != nil {
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

type SetQueueAttributesRequestDlqPolicy struct {
	// The target queue for dead-letter message delivery.
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
	// The maximum number of times a message can be delivered.
	//
	// example:
	//
	// 3
	MaxReceiveCount *int32 `json:"MaxReceiveCount,omitempty" xml:"MaxReceiveCount,omitempty"`
}

func (s SetQueueAttributesRequestDlqPolicy) String() string {
	return dara.Prettify(s)
}

func (s SetQueueAttributesRequestDlqPolicy) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesRequestDlqPolicy) GetDeadLetterTargetQueue() *string {
	return s.DeadLetterTargetQueue
}

func (s *SetQueueAttributesRequestDlqPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *SetQueueAttributesRequestDlqPolicy) GetMaxReceiveCount() *int32 {
	return s.MaxReceiveCount
}

func (s *SetQueueAttributesRequestDlqPolicy) SetDeadLetterTargetQueue(v string) *SetQueueAttributesRequestDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *SetQueueAttributesRequestDlqPolicy) SetEnabled(v bool) *SetQueueAttributesRequestDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *SetQueueAttributesRequestDlqPolicy) SetMaxReceiveCount(v int32) *SetQueueAttributesRequestDlqPolicy {
	s.MaxReceiveCount = &v
	return s
}

func (s *SetQueueAttributesRequestDlqPolicy) Validate() error {
	return dara.Validate(s)
}

type SetQueueAttributesRequestTenantRateLimitPolicy struct {
	Enabled              *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	MaxReceivesPerSecond *int32 `json:"MaxReceivesPerSecond,omitempty" xml:"MaxReceivesPerSecond,omitempty"`
}

func (s SetQueueAttributesRequestTenantRateLimitPolicy) String() string {
	return dara.Prettify(s)
}

func (s SetQueueAttributesRequestTenantRateLimitPolicy) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesRequestTenantRateLimitPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *SetQueueAttributesRequestTenantRateLimitPolicy) GetMaxReceivesPerSecond() *int32 {
	return s.MaxReceivesPerSecond
}

func (s *SetQueueAttributesRequestTenantRateLimitPolicy) SetEnabled(v bool) *SetQueueAttributesRequestTenantRateLimitPolicy {
	s.Enabled = &v
	return s
}

func (s *SetQueueAttributesRequestTenantRateLimitPolicy) SetMaxReceivesPerSecond(v int32) *SetQueueAttributesRequestTenantRateLimitPolicy {
	s.MaxReceivesPerSecond = &v
	return s
}

func (s *SetQueueAttributesRequestTenantRateLimitPolicy) Validate() error {
	return dara.Validate(s)
}
