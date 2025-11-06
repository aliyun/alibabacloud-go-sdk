// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueUpstreamBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListQueueUpstreamBindingsRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *ListQueueUpstreamBindingsRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListQueueUpstreamBindingsRequest
	GetInstanceId() *string
	SetPageSize(v int32) *ListQueueUpstreamBindingsRequest
	GetPageSize() *int32
	SetQueueName(v string) *ListQueueUpstreamBindingsRequest
	GetQueueName() *string
	SetVhostName(v string) *ListQueueUpstreamBindingsRequest
	GetVhostName() *string
}

type ListQueueUpstreamBindingsRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s ListQueueUpstreamBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpstreamBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListQueueUpstreamBindingsRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListQueueUpstreamBindingsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListQueueUpstreamBindingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListQueueUpstreamBindingsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQueueUpstreamBindingsRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *ListQueueUpstreamBindingsRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *ListQueueUpstreamBindingsRequest) SetConsoleSessionId(v string) *ListQueueUpstreamBindingsRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListQueueUpstreamBindingsRequest) SetCurrentPage(v int32) *ListQueueUpstreamBindingsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListQueueUpstreamBindingsRequest) SetInstanceId(v string) *ListQueueUpstreamBindingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueueUpstreamBindingsRequest) SetPageSize(v int32) *ListQueueUpstreamBindingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListQueueUpstreamBindingsRequest) SetQueueName(v string) *ListQueueUpstreamBindingsRequest {
	s.QueueName = &v
	return s
}

func (s *ListQueueUpstreamBindingsRequest) SetVhostName(v string) *ListQueueUpstreamBindingsRequest {
	s.VhostName = &v
	return s
}

func (s *ListQueueUpstreamBindingsRequest) Validate() error {
	return dara.Validate(s)
}
