// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingErrorByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetBindingErrorByTaskIdRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetBindingErrorByTaskIdRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetBindingErrorByTaskIdRequest
	GetPageSize() *int32
	SetTaskId(v int64) *GetBindingErrorByTaskIdRequest
	GetTaskId() *int64
}

type GetBindingErrorByTaskIdRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetBindingErrorByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBindingErrorByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetBindingErrorByTaskIdRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetBindingErrorByTaskIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetBindingErrorByTaskIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetBindingErrorByTaskIdRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetBindingErrorByTaskIdRequest) SetConsoleSessionId(v string) *GetBindingErrorByTaskIdRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetBindingErrorByTaskIdRequest) SetCurrentPage(v int32) *GetBindingErrorByTaskIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetBindingErrorByTaskIdRequest) SetPageSize(v int32) *GetBindingErrorByTaskIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetBindingErrorByTaskIdRequest) SetTaskId(v int64) *GetBindingErrorByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *GetBindingErrorByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
