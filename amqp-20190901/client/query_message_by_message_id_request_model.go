// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageByMessageIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryMessageByMessageIdRequest
	GetBeginTime() *int64
	SetConsoleSessionId(v string) *QueryMessageByMessageIdRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *QueryMessageByMessageIdRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *QueryMessageByMessageIdRequest
	GetEndTime() *int64
	SetInstanceId(v string) *QueryMessageByMessageIdRequest
	GetInstanceId() *string
	SetMessageId(v string) *QueryMessageByMessageIdRequest
	GetMessageId() *string
	SetPageSize(v int32) *QueryMessageByMessageIdRequest
	GetPageSize() *int32
	SetQueueName(v string) *QueryMessageByMessageIdRequest
	GetQueueName() *string
	SetVhostName(v string) *QueryMessageByMessageIdRequest
	GetVhostName() *string
}

type QueryMessageByMessageIdRequest struct {
	BeginTime        *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	CurrentPage      *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s QueryMessageByMessageIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByMessageIdRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageByMessageIdRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryMessageByMessageIdRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *QueryMessageByMessageIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMessageByMessageIdRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMessageByMessageIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMessageByMessageIdRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryMessageByMessageIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMessageByMessageIdRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *QueryMessageByMessageIdRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *QueryMessageByMessageIdRequest) SetBeginTime(v int64) *QueryMessageByMessageIdRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) SetConsoleSessionId(v string) *QueryMessageByMessageIdRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) SetCurrentPage(v int32) *QueryMessageByMessageIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) SetEndTime(v int64) *QueryMessageByMessageIdRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) SetInstanceId(v string) *QueryMessageByMessageIdRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) SetMessageId(v string) *QueryMessageByMessageIdRequest {
	s.MessageId = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) SetPageSize(v int32) *QueryMessageByMessageIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) SetQueueName(v string) *QueryMessageByMessageIdRequest {
	s.QueueName = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) SetVhostName(v string) *QueryMessageByMessageIdRequest {
	s.VhostName = &v
	return s
}

func (s *QueryMessageByMessageIdRequest) Validate() error {
	return dara.Validate(s)
}
