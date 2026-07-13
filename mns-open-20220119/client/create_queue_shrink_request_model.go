// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelaySeconds(v int64) *CreateQueueShrinkRequest
	GetDelaySeconds() *int64
	SetDlqPolicyShrink(v string) *CreateQueueShrinkRequest
	GetDlqPolicyShrink() *string
	SetEnableLogging(v bool) *CreateQueueShrinkRequest
	GetEnableLogging() *bool
	SetEnableSSE(v bool) *CreateQueueShrinkRequest
	GetEnableSSE() *bool
	SetKmsKeyId(v string) *CreateQueueShrinkRequest
	GetKmsKeyId() *string
	SetMaximumMessageSize(v int64) *CreateQueueShrinkRequest
	GetMaximumMessageSize() *int64
	SetMessageRetentionPeriod(v int64) *CreateQueueShrinkRequest
	GetMessageRetentionPeriod() *int64
	SetPollingWaitSeconds(v int64) *CreateQueueShrinkRequest
	GetPollingWaitSeconds() *int64
	SetQueueName(v string) *CreateQueueShrinkRequest
	GetQueueName() *string
	SetQueueType(v string) *CreateQueueShrinkRequest
	GetQueueType() *string
	SetSseAlgorithm(v string) *CreateQueueShrinkRequest
	GetSseAlgorithm() *string
	SetSseType(v string) *CreateQueueShrinkRequest
	GetSseType() *string
	SetTag(v []*CreateQueueShrinkRequestTag) *CreateQueueShrinkRequest
	GetTag() []*CreateQueueShrinkRequestTag
	SetTenantRateLimitPolicyShrink(v string) *CreateQueueShrinkRequest
	GetTenantRateLimitPolicyShrink() *string
	SetVisibilityTimeout(v int64) *CreateQueueShrinkRequest
	GetVisibilityTimeout() *int64
}

type CreateQueueShrinkRequest struct {
	// The delay period for all messages sent to the queue. A message sent to the queue can be consumed only after the delay period specified by this parameter elapses. Unit: seconds.
	//
	// Valid values: 0 to 604800.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The dead-letter policy.
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
	// The maximum size of a message body that can be sent to the queue. Unit: bytes.
	//
	// Valid values: 1024 to 65536.
	//
	// Default value: 65536.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified duration elapses from the time the message is sent to the queue, the message is deleted regardless of whether it has been consumed. Unit: seconds.
	//
	// Valid values: 60 to 604800.
	//
	// Default value: 345600.
	//
	// example:
	//
	// 345600
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum wait time for a ReceiveMessage request when the queue is empty. Unit: seconds.
	//
	// Valid values: 0 to 30.
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
	// 06273500-249F-5863-121D-74D51123****
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The type of the queue. Valid values:
	//
	//    	- normal: standard queue.
	//
	//    	- fifo: FIFO queue.
	//
	// example:
	//
	// normal
	QueueType    *string `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
	SseAlgorithm *string `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	SseType      *string `json:"SseType,omitempty" xml:"SseType,omitempty"`
	// The list of resource tags.
	Tag []*CreateQueueShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The rate limiting policy.
	TenantRateLimitPolicyShrink *string `json:"TenantRateLimitPolicy,omitempty" xml:"TenantRateLimitPolicy,omitempty"`
	// The duration for which a consumed message stays in the Inactive state after it is changed from the Active state. Unit: seconds.
	//
	// Valid values: 1 to 43200.
	//
	// Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s CreateQueueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueShrinkRequest) GetDelaySeconds() *int64 {
	return s.DelaySeconds
}

func (s *CreateQueueShrinkRequest) GetDlqPolicyShrink() *string {
	return s.DlqPolicyShrink
}

func (s *CreateQueueShrinkRequest) GetEnableLogging() *bool {
	return s.EnableLogging
}

func (s *CreateQueueShrinkRequest) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *CreateQueueShrinkRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateQueueShrinkRequest) GetMaximumMessageSize() *int64 {
	return s.MaximumMessageSize
}

func (s *CreateQueueShrinkRequest) GetMessageRetentionPeriod() *int64 {
	return s.MessageRetentionPeriod
}

func (s *CreateQueueShrinkRequest) GetPollingWaitSeconds() *int64 {
	return s.PollingWaitSeconds
}

func (s *CreateQueueShrinkRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateQueueShrinkRequest) GetQueueType() *string {
	return s.QueueType
}

func (s *CreateQueueShrinkRequest) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *CreateQueueShrinkRequest) GetSseType() *string {
	return s.SseType
}

func (s *CreateQueueShrinkRequest) GetTag() []*CreateQueueShrinkRequestTag {
	return s.Tag
}

func (s *CreateQueueShrinkRequest) GetTenantRateLimitPolicyShrink() *string {
	return s.TenantRateLimitPolicyShrink
}

func (s *CreateQueueShrinkRequest) GetVisibilityTimeout() *int64 {
	return s.VisibilityTimeout
}

func (s *CreateQueueShrinkRequest) SetDelaySeconds(v int64) *CreateQueueShrinkRequest {
	s.DelaySeconds = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetDlqPolicyShrink(v string) *CreateQueueShrinkRequest {
	s.DlqPolicyShrink = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetEnableLogging(v bool) *CreateQueueShrinkRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetEnableSSE(v bool) *CreateQueueShrinkRequest {
	s.EnableSSE = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetKmsKeyId(v string) *CreateQueueShrinkRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetMaximumMessageSize(v int64) *CreateQueueShrinkRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetMessageRetentionPeriod(v int64) *CreateQueueShrinkRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetPollingWaitSeconds(v int64) *CreateQueueShrinkRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetQueueName(v string) *CreateQueueShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetQueueType(v string) *CreateQueueShrinkRequest {
	s.QueueType = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetSseAlgorithm(v string) *CreateQueueShrinkRequest {
	s.SseAlgorithm = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetSseType(v string) *CreateQueueShrinkRequest {
	s.SseType = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetTag(v []*CreateQueueShrinkRequestTag) *CreateQueueShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateQueueShrinkRequest) SetTenantRateLimitPolicyShrink(v string) *CreateQueueShrinkRequest {
	s.TenantRateLimitPolicyShrink = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetVisibilityTimeout(v int64) *CreateQueueShrinkRequest {
	s.VisibilityTimeout = &v
	return s
}

func (s *CreateQueueShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateQueueShrinkRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQueueShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateQueueShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateQueueShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateQueueShrinkRequestTag) SetKey(v string) *CreateQueueShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateQueueShrinkRequestTag) SetValue(v string) *CreateQueueShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateQueueShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
