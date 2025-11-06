// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueErrorByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetQueueErrorByTaskIdRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetQueueErrorByTaskIdRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetQueueErrorByTaskIdRequest
	GetPageSize() *int32
	SetTaskId(v int64) *GetQueueErrorByTaskIdRequest
	GetTaskId() *int64
}

type GetQueueErrorByTaskIdRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetQueueErrorByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueErrorByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetQueueErrorByTaskIdRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetQueueErrorByTaskIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetQueueErrorByTaskIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueueErrorByTaskIdRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetQueueErrorByTaskIdRequest) SetConsoleSessionId(v string) *GetQueueErrorByTaskIdRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetQueueErrorByTaskIdRequest) SetCurrentPage(v int32) *GetQueueErrorByTaskIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetQueueErrorByTaskIdRequest) SetPageSize(v int32) *GetQueueErrorByTaskIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueueErrorByTaskIdRequest) SetTaskId(v int64) *GetQueueErrorByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *GetQueueErrorByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
