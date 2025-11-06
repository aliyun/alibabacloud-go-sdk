// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMsgIdListByQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetMsgIdListByQueueRequest
	GetClientToken() *string
	SetConsoleSessionId(v string) *GetMsgIdListByQueueRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetMsgIdListByQueueRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *GetMsgIdListByQueueRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetMsgIdListByQueueRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetMsgIdListByQueueRequest
	GetPageSize() *int32
	SetQueueName(v string) *GetMsgIdListByQueueRequest
	GetQueueName() *string
	SetStartTime(v int64) *GetMsgIdListByQueueRequest
	GetStartTime() *int64
	SetVhostName(v string) *GetMsgIdListByQueueRequest
	GetVhostName() *string
}

type GetMsgIdListByQueueRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetMsgIdListByQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMsgIdListByQueueRequest) GoString() string {
	return s.String()
}

func (s *GetMsgIdListByQueueRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetMsgIdListByQueueRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetMsgIdListByQueueRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetMsgIdListByQueueRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetMsgIdListByQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMsgIdListByQueueRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMsgIdListByQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetMsgIdListByQueueRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetMsgIdListByQueueRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetMsgIdListByQueueRequest) SetClientToken(v string) *GetMsgIdListByQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) SetConsoleSessionId(v string) *GetMsgIdListByQueueRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) SetCurrentPage(v int32) *GetMsgIdListByQueueRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) SetEndTime(v int64) *GetMsgIdListByQueueRequest {
	s.EndTime = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) SetInstanceId(v string) *GetMsgIdListByQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) SetPageSize(v int32) *GetMsgIdListByQueueRequest {
	s.PageSize = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) SetQueueName(v string) *GetMsgIdListByQueueRequest {
	s.QueueName = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) SetStartTime(v int64) *GetMsgIdListByQueueRequest {
	s.StartTime = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) SetVhostName(v string) *GetMsgIdListByQueueRequest {
	s.VhostName = &v
	return s
}

func (s *GetMsgIdListByQueueRequest) Validate() error {
	return dara.Validate(s)
}
