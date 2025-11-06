// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListQueueRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *ListQueueRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListQueueRequest
	GetInstanceId() *string
	SetPageSize(v int32) *ListQueueRequest
	GetPageSize() *int32
	SetQueueNamePrefix(v string) *ListQueueRequest
	GetQueueNamePrefix() *string
	SetVhostName(v string) *ListQueueRequest
	GetVhostName() *string
}

type ListQueueRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueueNamePrefix *string `json:"QueueNamePrefix,omitempty" xml:"QueueNamePrefix,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s ListQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueueRequest) GoString() string {
	return s.String()
}

func (s *ListQueueRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListQueueRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListQueueRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQueueRequest) GetQueueNamePrefix() *string {
	return s.QueueNamePrefix
}

func (s *ListQueueRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *ListQueueRequest) SetConsoleSessionId(v string) *ListQueueRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListQueueRequest) SetCurrentPage(v int32) *ListQueueRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListQueueRequest) SetInstanceId(v string) *ListQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *ListQueueRequest) SetPageSize(v int32) *ListQueueRequest {
	s.PageSize = &v
	return s
}

func (s *ListQueueRequest) SetQueueNamePrefix(v string) *ListQueueRequest {
	s.QueueNamePrefix = &v
	return s
}

func (s *ListQueueRequest) SetVhostName(v string) *ListQueueRequest {
	s.VhostName = &v
	return s
}

func (s *ListQueueRequest) Validate() error {
	return dara.Validate(s)
}
