// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetSendTraceByQueueRequest
	GetClientToken() *string
	SetConsoleSessionId(v string) *GetSendTraceByQueueRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetSendTraceByQueueRequest
	GetCurrentPage() *int32
	SetEndTime(v int32) *GetSendTraceByQueueRequest
	GetEndTime() *int32
	SetInstanceId(v string) *GetSendTraceByQueueRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetSendTraceByQueueRequest
	GetPageSize() *int32
	SetQueueName(v string) *GetSendTraceByQueueRequest
	GetQueueName() *string
	SetStartTime(v int32) *GetSendTraceByQueueRequest
	GetStartTime() *int32
	SetVhostName(v string) *GetSendTraceByQueueRequest
	GetVhostName() *string
}

type GetSendTraceByQueueRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetSendTraceByQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByQueueRequest) GoString() string {
	return s.String()
}

func (s *GetSendTraceByQueueRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetSendTraceByQueueRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetSendTraceByQueueRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSendTraceByQueueRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *GetSendTraceByQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSendTraceByQueueRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSendTraceByQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetSendTraceByQueueRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *GetSendTraceByQueueRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetSendTraceByQueueRequest) SetClientToken(v string) *GetSendTraceByQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *GetSendTraceByQueueRequest) SetConsoleSessionId(v string) *GetSendTraceByQueueRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetSendTraceByQueueRequest) SetCurrentPage(v int32) *GetSendTraceByQueueRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSendTraceByQueueRequest) SetEndTime(v int32) *GetSendTraceByQueueRequest {
	s.EndTime = &v
	return s
}

func (s *GetSendTraceByQueueRequest) SetInstanceId(v string) *GetSendTraceByQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSendTraceByQueueRequest) SetPageSize(v int32) *GetSendTraceByQueueRequest {
	s.PageSize = &v
	return s
}

func (s *GetSendTraceByQueueRequest) SetQueueName(v string) *GetSendTraceByQueueRequest {
	s.QueueName = &v
	return s
}

func (s *GetSendTraceByQueueRequest) SetStartTime(v int32) *GetSendTraceByQueueRequest {
	s.StartTime = &v
	return s
}

func (s *GetSendTraceByQueueRequest) SetVhostName(v string) *GetSendTraceByQueueRequest {
	s.VhostName = &v
	return s
}

func (s *GetSendTraceByQueueRequest) Validate() error {
	return dara.Validate(s)
}
