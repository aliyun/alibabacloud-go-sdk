// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAckInfoOfMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetAckInfoOfMessageRequest
	GetClientToken() *string
	SetConsoleSessionId(v string) *GetAckInfoOfMessageRequest
	GetConsoleSessionId() *string
	SetConsumeStatus(v string) *GetAckInfoOfMessageRequest
	GetConsumeStatus() *string
	SetDeliveryTag(v string) *GetAckInfoOfMessageRequest
	GetDeliveryTag() *string
	SetEndTime(v int64) *GetAckInfoOfMessageRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetAckInfoOfMessageRequest
	GetInstanceId() *string
	SetMsgId(v string) *GetAckInfoOfMessageRequest
	GetMsgId() *string
	SetQueueName(v string) *GetAckInfoOfMessageRequest
	GetQueueName() *string
	SetStartTime(v int64) *GetAckInfoOfMessageRequest
	GetStartTime() *int64
	SetTimeStamp(v string) *GetAckInfoOfMessageRequest
	GetTimeStamp() *string
	SetVhostName(v string) *GetAckInfoOfMessageRequest
	GetVhostName() *string
}

type GetAckInfoOfMessageRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	ConsumeStatus *string `json:"ConsumeStatus,omitempty" xml:"ConsumeStatus,omitempty"`
	// This parameter is required.
	DeliveryTag *string `json:"DeliveryTag,omitempty" xml:"DeliveryTag,omitempty"`
	// This parameter is required.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetAckInfoOfMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoOfMessageRequest) GoString() string {
	return s.String()
}

func (s *GetAckInfoOfMessageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetAckInfoOfMessageRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetAckInfoOfMessageRequest) GetConsumeStatus() *string {
	return s.ConsumeStatus
}

func (s *GetAckInfoOfMessageRequest) GetDeliveryTag() *string {
	return s.DeliveryTag
}

func (s *GetAckInfoOfMessageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAckInfoOfMessageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAckInfoOfMessageRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *GetAckInfoOfMessageRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetAckInfoOfMessageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAckInfoOfMessageRequest) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *GetAckInfoOfMessageRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetAckInfoOfMessageRequest) SetClientToken(v string) *GetAckInfoOfMessageRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetConsoleSessionId(v string) *GetAckInfoOfMessageRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetConsumeStatus(v string) *GetAckInfoOfMessageRequest {
	s.ConsumeStatus = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetDeliveryTag(v string) *GetAckInfoOfMessageRequest {
	s.DeliveryTag = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetEndTime(v int64) *GetAckInfoOfMessageRequest {
	s.EndTime = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetInstanceId(v string) *GetAckInfoOfMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetMsgId(v string) *GetAckInfoOfMessageRequest {
	s.MsgId = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetQueueName(v string) *GetAckInfoOfMessageRequest {
	s.QueueName = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetStartTime(v int64) *GetAckInfoOfMessageRequest {
	s.StartTime = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetTimeStamp(v string) *GetAckInfoOfMessageRequest {
	s.TimeStamp = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) SetVhostName(v string) *GetAckInfoOfMessageRequest {
	s.VhostName = &v
	return s
}

func (s *GetAckInfoOfMessageRequest) Validate() error {
	return dara.Validate(s)
}
