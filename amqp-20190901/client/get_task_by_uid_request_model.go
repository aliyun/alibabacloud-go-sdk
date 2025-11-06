// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskByUidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetTaskByUidRequest
	GetConsoleSessionId() *string
	SetCurrentPage(v int32) *GetTaskByUidRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetTaskByUidRequest
	GetPageSize() *int32
}

type GetTaskByUidRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetTaskByUidRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUidRequest) GoString() string {
	return s.String()
}

func (s *GetTaskByUidRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetTaskByUidRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetTaskByUidRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTaskByUidRequest) SetConsoleSessionId(v string) *GetTaskByUidRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetTaskByUidRequest) SetCurrentPage(v int32) *GetTaskByUidRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetTaskByUidRequest) SetPageSize(v int32) *GetTaskByUidRequest {
	s.PageSize = &v
	return s
}

func (s *GetTaskByUidRequest) Validate() error {
	return dara.Validate(s)
}
