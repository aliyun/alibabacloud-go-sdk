// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDelete(v bool) *CreateQueueRequest
	GetAutoDelete() *bool
	SetAutoExpire(v int64) *CreateQueueRequest
	GetAutoExpire() *int64
	SetConsoleSessionId(v string) *CreateQueueRequest
	GetConsoleSessionId() *string
	SetDeadLetterExchange(v string) *CreateQueueRequest
	GetDeadLetterExchange() *string
	SetDeadLetterRoutingKey(v string) *CreateQueueRequest
	GetDeadLetterRoutingKey() *string
	SetExclusive(v bool) *CreateQueueRequest
	GetExclusive() *bool
	SetInstanceId(v string) *CreateQueueRequest
	GetInstanceId() *string
	SetMaxLength(v int64) *CreateQueueRequest
	GetMaxLength() *int64
	SetMaximunPrioty(v int64) *CreateQueueRequest
	GetMaximunPrioty() *int64
	SetMessageTTL(v int64) *CreateQueueRequest
	GetMessageTTL() *int64
	SetOrdered(v bool) *CreateQueueRequest
	GetOrdered() *bool
	SetQueueName(v string) *CreateQueueRequest
	GetQueueName() *string
	SetRetryInherit(v bool) *CreateQueueRequest
	GetRetryInherit() *bool
	SetRetryInterval(v int32) *CreateQueueRequest
	GetRetryInterval() *int32
	SetRetryTimes(v int32) *CreateQueueRequest
	GetRetryTimes() *int32
	SetSingleActiveConsumer(v bool) *CreateQueueRequest
	GetSingleActiveConsumer() *bool
	SetVhostName(v string) *CreateQueueRequest
	GetVhostName() *string
}

type CreateQueueRequest struct {
	AutoDelete           *bool   `json:"AutoDelete,omitempty" xml:"AutoDelete,omitempty"`
	AutoExpire           *int64  `json:"AutoExpire,omitempty" xml:"AutoExpire,omitempty"`
	ConsoleSessionId     *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	DeadLetterExchange   *string `json:"DeadLetterExchange,omitempty" xml:"DeadLetterExchange,omitempty"`
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" xml:"DeadLetterRoutingKey,omitempty"`
	Exclusive            *bool   `json:"Exclusive,omitempty" xml:"Exclusive,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaximunPrioty        *int64  `json:"MaximunPrioty,omitempty" xml:"MaximunPrioty,omitempty"`
	MessageTTL           *int64  `json:"MessageTTL,omitempty" xml:"MessageTTL,omitempty"`
	Ordered              *bool   `json:"Ordered,omitempty" xml:"Ordered,omitempty"`
	// This parameter is required.
	QueueName            *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	RetryInherit         *bool   `json:"RetryInherit,omitempty" xml:"RetryInherit,omitempty"`
	RetryInterval        *int32  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryTimes           *int32  `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	SingleActiveConsumer *bool   `json:"SingleActiveConsumer,omitempty" xml:"SingleActiveConsumer,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) GetAutoDelete() *bool {
	return s.AutoDelete
}

func (s *CreateQueueRequest) GetAutoExpire() *int64 {
	return s.AutoExpire
}

func (s *CreateQueueRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *CreateQueueRequest) GetDeadLetterExchange() *string {
	return s.DeadLetterExchange
}

func (s *CreateQueueRequest) GetDeadLetterRoutingKey() *string {
	return s.DeadLetterRoutingKey
}

func (s *CreateQueueRequest) GetExclusive() *bool {
	return s.Exclusive
}

func (s *CreateQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateQueueRequest) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *CreateQueueRequest) GetMaximunPrioty() *int64 {
	return s.MaximunPrioty
}

func (s *CreateQueueRequest) GetMessageTTL() *int64 {
	return s.MessageTTL
}

func (s *CreateQueueRequest) GetOrdered() *bool {
	return s.Ordered
}

func (s *CreateQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateQueueRequest) GetRetryInherit() *bool {
	return s.RetryInherit
}

func (s *CreateQueueRequest) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *CreateQueueRequest) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *CreateQueueRequest) GetSingleActiveConsumer() *bool {
	return s.SingleActiveConsumer
}

func (s *CreateQueueRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *CreateQueueRequest) SetAutoDelete(v bool) *CreateQueueRequest {
	s.AutoDelete = &v
	return s
}

func (s *CreateQueueRequest) SetAutoExpire(v int64) *CreateQueueRequest {
	s.AutoExpire = &v
	return s
}

func (s *CreateQueueRequest) SetConsoleSessionId(v string) *CreateQueueRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *CreateQueueRequest) SetDeadLetterExchange(v string) *CreateQueueRequest {
	s.DeadLetterExchange = &v
	return s
}

func (s *CreateQueueRequest) SetDeadLetterRoutingKey(v string) *CreateQueueRequest {
	s.DeadLetterRoutingKey = &v
	return s
}

func (s *CreateQueueRequest) SetExclusive(v bool) *CreateQueueRequest {
	s.Exclusive = &v
	return s
}

func (s *CreateQueueRequest) SetInstanceId(v string) *CreateQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQueueRequest) SetMaxLength(v int64) *CreateQueueRequest {
	s.MaxLength = &v
	return s
}

func (s *CreateQueueRequest) SetMaximunPrioty(v int64) *CreateQueueRequest {
	s.MaximunPrioty = &v
	return s
}

func (s *CreateQueueRequest) SetMessageTTL(v int64) *CreateQueueRequest {
	s.MessageTTL = &v
	return s
}

func (s *CreateQueueRequest) SetOrdered(v bool) *CreateQueueRequest {
	s.Ordered = &v
	return s
}

func (s *CreateQueueRequest) SetQueueName(v string) *CreateQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueRequest) SetRetryInherit(v bool) *CreateQueueRequest {
	s.RetryInherit = &v
	return s
}

func (s *CreateQueueRequest) SetRetryInterval(v int32) *CreateQueueRequest {
	s.RetryInterval = &v
	return s
}

func (s *CreateQueueRequest) SetRetryTimes(v int32) *CreateQueueRequest {
	s.RetryTimes = &v
	return s
}

func (s *CreateQueueRequest) SetSingleActiveConsumer(v bool) *CreateQueueRequest {
	s.SingleActiveConsumer = &v
	return s
}

func (s *CreateQueueRequest) SetVhostName(v string) *CreateQueueRequest {
	s.VhostName = &v
	return s
}

func (s *CreateQueueRequest) Validate() error {
	return dara.Validate(s)
}
