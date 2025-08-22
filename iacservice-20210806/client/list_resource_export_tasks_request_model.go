// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskId(v string) *ListResourceExportTasksRequest
	GetExportTaskId() *string
	SetKeyword(v string) *ListResourceExportTasksRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListResourceExportTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceExportTasksRequest
	GetPageSize() *int32
}

type ListResourceExportTasksRequest struct {
	// example:
	//
	// ex-al1c11jl9g2tbte727otp85
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListResourceExportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksRequest) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *ListResourceExportTasksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListResourceExportTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceExportTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceExportTasksRequest) SetExportTaskId(v string) *ListResourceExportTasksRequest {
	s.ExportTaskId = &v
	return s
}

func (s *ListResourceExportTasksRequest) SetKeyword(v string) *ListResourceExportTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListResourceExportTasksRequest) SetPageNumber(v int32) *ListResourceExportTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceExportTasksRequest) SetPageSize(v int32) *ListResourceExportTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceExportTasksRequest) Validate() error {
	return dara.Validate(s)
}
