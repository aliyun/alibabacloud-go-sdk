// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAckInfoByIntervalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetAckInfoByIntervalRequest
	GetClientToken() *string
	SetConsoleSessionId(v string) *GetAckInfoByIntervalRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetAckInfoByIntervalRequest
	GetCurrentPage() *int32
	SetEndTime(v int32) *GetAckInfoByIntervalRequest
	GetEndTime() *int32
	SetInstanceId(v string) *GetAckInfoByIntervalRequest
	GetInstanceId() *string
	SetIntervalSec(v int32) *GetAckInfoByIntervalRequest
	GetIntervalSec() *int32
	SetPageSize(v int32) *GetAckInfoByIntervalRequest
	GetPageSize() *int32
	SetQueueName(v string) *GetAckInfoByIntervalRequest
	GetQueueName() *string
	SetStartTime(v int32) *GetAckInfoByIntervalRequest
	GetStartTime() *int32
	SetVhostName(v string) *GetAckInfoByIntervalRequest
	GetVhostName() *string
}

type GetAckInfoByIntervalRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	IntervalSec *int32 `json:"IntervalSec,omitempty" xml:"IntervalSec,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetAckInfoByIntervalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoByIntervalRequest) GoString() string {
	return s.String()
}

func (s *GetAckInfoByIntervalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetAckInfoByIntervalRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetAckInfoByIntervalRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAckInfoByIntervalRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *GetAckInfoByIntervalRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAckInfoByIntervalRequest) GetIntervalSec() *int32 {
	return s.IntervalSec
}

func (s *GetAckInfoByIntervalRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAckInfoByIntervalRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetAckInfoByIntervalRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *GetAckInfoByIntervalRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetAckInfoByIntervalRequest) SetClientToken(v string) *GetAckInfoByIntervalRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetConsoleSessionId(v string) *GetAckInfoByIntervalRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetCurrentPage(v int32) *GetAckInfoByIntervalRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetEndTime(v int32) *GetAckInfoByIntervalRequest {
	s.EndTime = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetInstanceId(v string) *GetAckInfoByIntervalRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetIntervalSec(v int32) *GetAckInfoByIntervalRequest {
	s.IntervalSec = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetPageSize(v int32) *GetAckInfoByIntervalRequest {
	s.PageSize = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetQueueName(v string) *GetAckInfoByIntervalRequest {
	s.QueueName = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetStartTime(v int32) *GetAckInfoByIntervalRequest {
	s.StartTime = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) SetVhostName(v string) *GetAckInfoByIntervalRequest {
	s.VhostName = &v
	return s
}

func (s *GetAckInfoByIntervalRequest) Validate() error {
	return dara.Validate(s)
}
