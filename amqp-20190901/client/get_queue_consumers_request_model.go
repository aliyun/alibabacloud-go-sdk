// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetQueueConsumersRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetQueueConsumersRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *GetQueueConsumersRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetQueueConsumersRequest
	GetPageSize() *int32
	SetQueueName(v string) *GetQueueConsumersRequest
	GetQueueName() *string
	SetVhostName(v string) *GetQueueConsumersRequest
	GetVhostName() *string
}

type GetQueueConsumersRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	CurrentPage      *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetQueueConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueConsumersRequest) GoString() string {
	return s.String()
}

func (s *GetQueueConsumersRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetQueueConsumersRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetQueueConsumersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueueConsumersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueueConsumersRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetQueueConsumersRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetQueueConsumersRequest) SetConsoleSessionId(v string) *GetQueueConsumersRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetQueueConsumersRequest) SetCurrentPage(v int32) *GetQueueConsumersRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetQueueConsumersRequest) SetInstanceId(v string) *GetQueueConsumersRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueueConsumersRequest) SetPageSize(v int32) *GetQueueConsumersRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueueConsumersRequest) SetQueueName(v string) *GetQueueConsumersRequest {
	s.QueueName = &v
	return s
}

func (s *GetQueueConsumersRequest) SetVhostName(v string) *GetQueueConsumersRequest {
	s.VhostName = &v
	return s
}

func (s *GetQueueConsumersRequest) Validate() error {
	return dara.Validate(s)
}
