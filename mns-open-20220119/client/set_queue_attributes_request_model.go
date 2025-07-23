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
	SetMaximumMessageSize(v int64) *SetQueueAttributesRequest
	GetMaximumMessageSize() *int64
	SetMessageRetentionPeriod(v int64) *SetQueueAttributesRequest
	GetMessageRetentionPeriod() *int64
	SetPollingWaitSeconds(v int64) *SetQueueAttributesRequest
	GetPollingWaitSeconds() *int64
	SetQueueName(v string) *SetQueueAttributesRequest
	GetQueueName() *string
	SetTenantRateLimitPolicy(v *SetQueueAttributesRequestTenantRateLimitPolicy) *SetQueueAttributesRequest
	GetTenantRateLimitPolicy() *SetQueueAttributesRequestTenantRateLimitPolicy
	SetVisibilityTimeout(v int64) *SetQueueAttributesRequest
	GetVisibilityTimeout() *int64
}

type SetQueueAttributesRequest struct {
	// The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *SetQueueAttributesRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled. Default value: false.
	//
	// example:
	//
	// True
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 1024
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	//
	// example:
	//
	// 120
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0
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
	TenantRateLimitPolicy *SetQueueAttributesRequestTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
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

func (s *SetQueueAttributesRequest) SetTenantRateLimitPolicy(v *SetQueueAttributesRequestTenantRateLimitPolicy) *SetQueueAttributesRequest {
	s.TenantRateLimitPolicy = v
	return s
}

func (s *SetQueueAttributesRequest) SetVisibilityTimeout(v int64) *SetQueueAttributesRequest {
	s.VisibilityTimeout = &v
	return s
}

func (s *SetQueueAttributesRequest) Validate() error {
	return dara.Validate(s)
}

type SetQueueAttributesRequestDlqPolicy struct {
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
	// The maximum number of retries.
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
