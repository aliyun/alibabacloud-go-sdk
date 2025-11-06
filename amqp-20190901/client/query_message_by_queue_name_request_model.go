// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageByQueueNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryMessageByQueueNameRequest
	GetBeginTime() *int64
	SetConsoleSessionId(v string) *QueryMessageByQueueNameRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *QueryMessageByQueueNameRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *QueryMessageByQueueNameRequest
	GetEndTime() *int64
	SetInstanceId(v string) *QueryMessageByQueueNameRequest
	GetInstanceId() *string
	SetPageSize(v int32) *QueryMessageByQueueNameRequest
	GetPageSize() *int32
	SetQueueName(v string) *QueryMessageByQueueNameRequest
	GetQueueName() *string
	SetTaskId(v string) *QueryMessageByQueueNameRequest
	GetTaskId() *string
	SetVhostName(v string) *QueryMessageByQueueNameRequest
	GetVhostName() *string
}

type QueryMessageByQueueNameRequest struct {
	// This parameter is required.
	BeginTime        *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s QueryMessageByQueueNameRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByQueueNameRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageByQueueNameRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryMessageByQueueNameRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *QueryMessageByQueueNameRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMessageByQueueNameRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMessageByQueueNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMessageByQueueNameRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMessageByQueueNameRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *QueryMessageByQueueNameRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryMessageByQueueNameRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *QueryMessageByQueueNameRequest) SetBeginTime(v int64) *QueryMessageByQueueNameRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) SetConsoleSessionId(v string) *QueryMessageByQueueNameRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) SetCurrentPage(v int32) *QueryMessageByQueueNameRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) SetEndTime(v int64) *QueryMessageByQueueNameRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) SetInstanceId(v string) *QueryMessageByQueueNameRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) SetPageSize(v int32) *QueryMessageByQueueNameRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) SetQueueName(v string) *QueryMessageByQueueNameRequest {
	s.QueueName = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) SetTaskId(v string) *QueryMessageByQueueNameRequest {
	s.TaskId = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) SetVhostName(v string) *QueryMessageByQueueNameRequest {
	s.VhostName = &v
	return s
}

func (s *QueryMessageByQueueNameRequest) Validate() error {
	return dara.Validate(s)
}
