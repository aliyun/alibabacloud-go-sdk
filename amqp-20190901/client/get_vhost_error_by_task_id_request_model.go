// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVhostErrorByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetVhostErrorByTaskIdRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetVhostErrorByTaskIdRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetVhostErrorByTaskIdRequest
	GetPageSize() *int32
	SetTaskId(v int64) *GetVhostErrorByTaskIdRequest
	GetTaskId() *int64
}

type GetVhostErrorByTaskIdRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetVhostErrorByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVhostErrorByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetVhostErrorByTaskIdRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetVhostErrorByTaskIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetVhostErrorByTaskIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetVhostErrorByTaskIdRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetVhostErrorByTaskIdRequest) SetConsoleSessionId(v string) *GetVhostErrorByTaskIdRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetVhostErrorByTaskIdRequest) SetCurrentPage(v int32) *GetVhostErrorByTaskIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetVhostErrorByTaskIdRequest) SetPageSize(v int32) *GetVhostErrorByTaskIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetVhostErrorByTaskIdRequest) SetTaskId(v int64) *GetVhostErrorByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *GetVhostErrorByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
