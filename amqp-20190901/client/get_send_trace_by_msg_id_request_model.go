// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByMsgIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetSendTraceByMsgIdRequest
	GetClientToken() *string
	SetConsoleSessionId(v string) *GetSendTraceByMsgIdRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetSendTraceByMsgIdRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *GetSendTraceByMsgIdRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetSendTraceByMsgIdRequest
	GetInstanceId() *string
	SetMsgId(v string) *GetSendTraceByMsgIdRequest
	GetMsgId() *string
	SetPageSize(v int32) *GetSendTraceByMsgIdRequest
	GetPageSize() *int32
	SetQueueName(v string) *GetSendTraceByMsgIdRequest
	GetQueueName() *string
	SetStartTime(v int64) *GetSendTraceByMsgIdRequest
	GetStartTime() *int64
	SetVhostName(v string) *GetSendTraceByMsgIdRequest
	GetVhostName() *string
}

type GetSendTraceByMsgIdRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// This parameter is required.
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetSendTraceByMsgIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *GetSendTraceByMsgIdRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetSendTraceByMsgIdRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetSendTraceByMsgIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSendTraceByMsgIdRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetSendTraceByMsgIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSendTraceByMsgIdRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *GetSendTraceByMsgIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSendTraceByMsgIdRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetSendTraceByMsgIdRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetSendTraceByMsgIdRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetSendTraceByMsgIdRequest) SetClientToken(v string) *GetSendTraceByMsgIdRequest {
	s.ClientToken = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetConsoleSessionId(v string) *GetSendTraceByMsgIdRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetCurrentPage(v int32) *GetSendTraceByMsgIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetEndTime(v int64) *GetSendTraceByMsgIdRequest {
	s.EndTime = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetInstanceId(v string) *GetSendTraceByMsgIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetMsgId(v string) *GetSendTraceByMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetPageSize(v int32) *GetSendTraceByMsgIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetQueueName(v string) *GetSendTraceByMsgIdRequest {
	s.QueueName = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetStartTime(v int64) *GetSendTraceByMsgIdRequest {
	s.StartTime = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) SetVhostName(v string) *GetSendTraceByMsgIdRequest {
	s.VhostName = &v
	return s
}

func (s *GetSendTraceByMsgIdRequest) Validate() error {
	return dara.Validate(s)
}
