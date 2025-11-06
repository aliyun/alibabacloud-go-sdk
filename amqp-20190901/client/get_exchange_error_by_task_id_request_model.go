// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExchangeErrorByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetExchangeErrorByTaskIdRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetExchangeErrorByTaskIdRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetExchangeErrorByTaskIdRequest
	GetPageSize() *int32
	SetTaskId(v int64) *GetExchangeErrorByTaskIdRequest
	GetTaskId() *int64
}

type GetExchangeErrorByTaskIdRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetExchangeErrorByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeErrorByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetExchangeErrorByTaskIdRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetExchangeErrorByTaskIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetExchangeErrorByTaskIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetExchangeErrorByTaskIdRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetExchangeErrorByTaskIdRequest) SetConsoleSessionId(v string) *GetExchangeErrorByTaskIdRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetExchangeErrorByTaskIdRequest) SetCurrentPage(v int32) *GetExchangeErrorByTaskIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetExchangeErrorByTaskIdRequest) SetPageSize(v int32) *GetExchangeErrorByTaskIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetExchangeErrorByTaskIdRequest) SetTaskId(v int64) *GetExchangeErrorByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *GetExchangeErrorByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
