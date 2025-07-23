// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelaySeconds(v int64) *CreateQueueRequest
	GetDelaySeconds() *int64
	SetDlqPolicy(v *CreateQueueRequestDlqPolicy) *CreateQueueRequest
	GetDlqPolicy() *CreateQueueRequestDlqPolicy
	SetEnableLogging(v bool) *CreateQueueRequest
	GetEnableLogging() *bool
	SetMaximumMessageSize(v int64) *CreateQueueRequest
	GetMaximumMessageSize() *int64
	SetMessageRetentionPeriod(v int64) *CreateQueueRequest
	GetMessageRetentionPeriod() *int64
	SetPollingWaitSeconds(v int64) *CreateQueueRequest
	GetPollingWaitSeconds() *int64
	SetQueueName(v string) *CreateQueueRequest
	GetQueueName() *string
	SetTag(v []*CreateQueueRequestTag) *CreateQueueRequest
	GetTag() []*CreateQueueRequestTag
	SetTenantRateLimitPolicy(v *CreateQueueRequestTenantRateLimitPolicy) *CreateQueueRequest
	GetTenantRateLimitPolicy() *CreateQueueRequestTenantRateLimitPolicy
	SetVisibilityTimeout(v int64) *CreateQueueRequest
	GetVisibilityTimeout() *int64
}

type CreateQueueRequest struct {
	// The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter queue policy.
	DlqPolicy *CreateQueueRequestDlqPolicy `json:"DlqPolicy,omitempty" xml:"DlqPolicy,omitempty" type:"Struct"`
	// Specifies whether to enable the log management feature. Valid values:
	//
	// 	- true: enabled.
	//
	// 	- false: disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is consumed. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	//
	// example:
	//
	// 345600
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
	// 06273500-249F-5863-121D-74D51123****
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The tags.
	Tag                   []*CreateQueueRequestTag                 `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TenantRateLimitPolicy *CreateQueueRequestTenantRateLimitPolicy `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty" type:"Struct"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) GetDelaySeconds() *int64 {
	return s.DelaySeconds
}

func (s *CreateQueueRequest) GetDlqPolicy() *CreateQueueRequestDlqPolicy {
	return s.DlqPolicy
}

func (s *CreateQueueRequest) GetEnableLogging() *bool {
	return s.EnableLogging
}

func (s *CreateQueueRequest) GetMaximumMessageSize() *int64 {
	return s.MaximumMessageSize
}

func (s *CreateQueueRequest) GetMessageRetentionPeriod() *int64 {
	return s.MessageRetentionPeriod
}

func (s *CreateQueueRequest) GetPollingWaitSeconds() *int64 {
	return s.PollingWaitSeconds
}

func (s *CreateQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateQueueRequest) GetTag() []*CreateQueueRequestTag {
	return s.Tag
}

func (s *CreateQueueRequest) GetTenantRateLimitPolicy() *CreateQueueRequestTenantRateLimitPolicy {
	return s.TenantRateLimitPolicy
}

func (s *CreateQueueRequest) GetVisibilityTimeout() *int64 {
	return s.VisibilityTimeout
}

func (s *CreateQueueRequest) SetDelaySeconds(v int64) *CreateQueueRequest {
	s.DelaySeconds = &v
	return s
}

func (s *CreateQueueRequest) SetDlqPolicy(v *CreateQueueRequestDlqPolicy) *CreateQueueRequest {
	s.DlqPolicy = v
	return s
}

func (s *CreateQueueRequest) SetEnableLogging(v bool) *CreateQueueRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateQueueRequest) SetMaximumMessageSize(v int64) *CreateQueueRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *CreateQueueRequest) SetMessageRetentionPeriod(v int64) *CreateQueueRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *CreateQueueRequest) SetPollingWaitSeconds(v int64) *CreateQueueRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *CreateQueueRequest) SetQueueName(v string) *CreateQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueRequest) SetTag(v []*CreateQueueRequestTag) *CreateQueueRequest {
	s.Tag = v
	return s
}

func (s *CreateQueueRequest) SetTenantRateLimitPolicy(v *CreateQueueRequestTenantRateLimitPolicy) *CreateQueueRequest {
	s.TenantRateLimitPolicy = v
	return s
}

func (s *CreateQueueRequest) SetVisibilityTimeout(v int64) *CreateQueueRequest {
	s.VisibilityTimeout = &v
	return s
}

func (s *CreateQueueRequest) Validate() error {
	return dara.Validate(s)
}

type CreateQueueRequestDlqPolicy struct {
	// The queue to which dead-letter messages are delivered.
	//
	// example:
	//
	// deadLetterQueue
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

func (s CreateQueueRequestDlqPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueRequestDlqPolicy) GoString() string {
	return s.String()
}

func (s *CreateQueueRequestDlqPolicy) GetDeadLetterTargetQueue() *string {
	return s.DeadLetterTargetQueue
}

func (s *CreateQueueRequestDlqPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateQueueRequestDlqPolicy) GetMaxReceiveCount() *int32 {
	return s.MaxReceiveCount
}

func (s *CreateQueueRequestDlqPolicy) SetDeadLetterTargetQueue(v string) *CreateQueueRequestDlqPolicy {
	s.DeadLetterTargetQueue = &v
	return s
}

func (s *CreateQueueRequestDlqPolicy) SetEnabled(v bool) *CreateQueueRequestDlqPolicy {
	s.Enabled = &v
	return s
}

func (s *CreateQueueRequestDlqPolicy) SetMaxReceiveCount(v int32) *CreateQueueRequestDlqPolicy {
	s.MaxReceiveCount = &v
	return s
}

func (s *CreateQueueRequestDlqPolicy) Validate() error {
	return dara.Validate(s)
}

type CreateQueueRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQueueRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueRequestTag) GoString() string {
	return s.String()
}

func (s *CreateQueueRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateQueueRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateQueueRequestTag) SetKey(v string) *CreateQueueRequestTag {
	s.Key = &v
	return s
}

func (s *CreateQueueRequestTag) SetValue(v string) *CreateQueueRequestTag {
	s.Value = &v
	return s
}

func (s *CreateQueueRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateQueueRequestTenantRateLimitPolicy struct {
	Enabled              *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	MaxReceivesPerSecond *int32 `json:"MaxReceivesPerSecond,omitempty" xml:"MaxReceivesPerSecond,omitempty"`
}

func (s CreateQueueRequestTenantRateLimitPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueRequestTenantRateLimitPolicy) GoString() string {
	return s.String()
}

func (s *CreateQueueRequestTenantRateLimitPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateQueueRequestTenantRateLimitPolicy) GetMaxReceivesPerSecond() *int32 {
	return s.MaxReceivesPerSecond
}

func (s *CreateQueueRequestTenantRateLimitPolicy) SetEnabled(v bool) *CreateQueueRequestTenantRateLimitPolicy {
	s.Enabled = &v
	return s
}

func (s *CreateQueueRequestTenantRateLimitPolicy) SetMaxReceivesPerSecond(v int32) *CreateQueueRequestTenantRateLimitPolicy {
	s.MaxReceivesPerSecond = &v
	return s
}

func (s *CreateQueueRequestTenantRateLimitPolicy) Validate() error {
	return dara.Validate(s)
}
