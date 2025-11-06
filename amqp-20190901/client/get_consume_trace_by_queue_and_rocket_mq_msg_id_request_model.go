// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumeTraceByQueueAndRocketMqMsgIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest
	GetClientToken() *string
	SetConsoleSessionId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest
	GetConsoleSessionId() *string
	SetEndTime(v int64) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest
	GetInstanceId() *string
	SetMsgId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest
	GetMsgId() *string
	SetQueueName(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest
	GetQueueName() *string
	SetStartTime(v int64) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest
	GetStartTime() *int64
	SetVhostName(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest
	GetVhostName() *string
}

type GetConsumeTraceByQueueAndRocketMqMsgIdRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
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
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetConsumeTraceByQueueAndRocketMqMsgIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GoString() string {
	return s.String()
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) SetClientToken(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest {
	s.ClientToken = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) SetConsoleSessionId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) SetEndTime(v int64) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest {
	s.EndTime = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) SetInstanceId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) SetMsgId(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) SetQueueName(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest {
	s.QueueName = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) SetStartTime(v int64) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest {
	s.StartTime = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) SetVhostName(v string) *GetConsumeTraceByQueueAndRocketMqMsgIdRequest {
	s.VhostName = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdRequest) Validate() error {
	return dara.Validate(s)
}
