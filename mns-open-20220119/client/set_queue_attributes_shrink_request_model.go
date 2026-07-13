// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetQueueAttributesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelaySeconds(v int64) *SetQueueAttributesShrinkRequest
	GetDelaySeconds() *int64
	SetDlqPolicyShrink(v string) *SetQueueAttributesShrinkRequest
	GetDlqPolicyShrink() *string
	SetEnableLogging(v bool) *SetQueueAttributesShrinkRequest
	GetEnableLogging() *bool
	SetEnableSSE(v bool) *SetQueueAttributesShrinkRequest
	GetEnableSSE() *bool
	SetKmsKeyId(v string) *SetQueueAttributesShrinkRequest
	GetKmsKeyId() *string
	SetMaximumMessageSize(v int64) *SetQueueAttributesShrinkRequest
	GetMaximumMessageSize() *int64
	SetMessageRetentionPeriod(v int64) *SetQueueAttributesShrinkRequest
	GetMessageRetentionPeriod() *int64
	SetPollingWaitSeconds(v int64) *SetQueueAttributesShrinkRequest
	GetPollingWaitSeconds() *int64
	SetQueueName(v string) *SetQueueAttributesShrinkRequest
	GetQueueName() *string
	SetSseAlgorithm(v string) *SetQueueAttributesShrinkRequest
	GetSseAlgorithm() *string
	SetSseType(v string) *SetQueueAttributesShrinkRequest
	GetSseType() *string
	SetTenantRateLimitPolicyShrink(v string) *SetQueueAttributesShrinkRequest
	GetTenantRateLimitPolicyShrink() *string
	SetVisibilityTimeout(v int64) *SetQueueAttributesShrinkRequest
	GetVisibilityTimeout() *int64
}

type SetQueueAttributesShrinkRequest struct {
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
	DlqPolicyShrink *string `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty"`
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
	QueueName                   *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	SseAlgorithm                *string `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	SseType                     *string `json:"SseType,omitempty" xml:"SseType,omitempty"`
	TenantRateLimitPolicyShrink *string `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty"`
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

func (s SetQueueAttributesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetQueueAttributesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesShrinkRequest) GetDelaySeconds() *int64 {
	return s.DelaySeconds
}

func (s *SetQueueAttributesShrinkRequest) GetDlqPolicyShrink() *string {
	return s.DlqPolicyShrink
}

func (s *SetQueueAttributesShrinkRequest) GetEnableLogging() *bool {
	return s.EnableLogging
}

func (s *SetQueueAttributesShrinkRequest) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *SetQueueAttributesShrinkRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *SetQueueAttributesShrinkRequest) GetMaximumMessageSize() *int64 {
	return s.MaximumMessageSize
}

func (s *SetQueueAttributesShrinkRequest) GetMessageRetentionPeriod() *int64 {
	return s.MessageRetentionPeriod
}

func (s *SetQueueAttributesShrinkRequest) GetPollingWaitSeconds() *int64 {
	return s.PollingWaitSeconds
}

func (s *SetQueueAttributesShrinkRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *SetQueueAttributesShrinkRequest) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *SetQueueAttributesShrinkRequest) GetSseType() *string {
	return s.SseType
}

func (s *SetQueueAttributesShrinkRequest) GetTenantRateLimitPolicyShrink() *string {
	return s.TenantRateLimitPolicyShrink
}

func (s *SetQueueAttributesShrinkRequest) GetVisibilityTimeout() *int64 {
	return s.VisibilityTimeout
}

func (s *SetQueueAttributesShrinkRequest) SetDelaySeconds(v int64) *SetQueueAttributesShrinkRequest {
	s.DelaySeconds = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetDlqPolicyShrink(v string) *SetQueueAttributesShrinkRequest {
	s.DlqPolicyShrink = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetEnableLogging(v bool) *SetQueueAttributesShrinkRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetEnableSSE(v bool) *SetQueueAttributesShrinkRequest {
	s.EnableSSE = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetKmsKeyId(v string) *SetQueueAttributesShrinkRequest {
	s.KmsKeyId = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetMaximumMessageSize(v int64) *SetQueueAttributesShrinkRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetMessageRetentionPeriod(v int64) *SetQueueAttributesShrinkRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetPollingWaitSeconds(v int64) *SetQueueAttributesShrinkRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetQueueName(v string) *SetQueueAttributesShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetSseAlgorithm(v string) *SetQueueAttributesShrinkRequest {
	s.SseAlgorithm = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetSseType(v string) *SetQueueAttributesShrinkRequest {
	s.SseType = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetTenantRateLimitPolicyShrink(v string) *SetQueueAttributesShrinkRequest {
	s.TenantRateLimitPolicyShrink = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) SetVisibilityTimeout(v int64) *SetQueueAttributesShrinkRequest {
	s.VisibilityTimeout = &v
	return s
}

func (s *SetQueueAttributesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
