// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListValidateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListValidateFileRequest
	GetEndTime() *string
	SetFileKeyword(v string) *ListValidateFileRequest
	GetFileKeyword() *string
	SetPage(v int32) *ListValidateFileRequest
	GetPage() *int32
	SetPageSize(v int32) *ListValidateFileRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListValidateFileRequest
	GetStartTime() *string
}

type ListValidateFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-19T20:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// keyword
	FileKeyword *string `json:"FileKeyword,omitempty" xml:"FileKeyword,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-12-19T08:30:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListValidateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ListValidateFileRequest) GoString() string {
	return s.String()
}

func (s *ListValidateFileRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListValidateFileRequest) GetFileKeyword() *string {
	return s.FileKeyword
}

func (s *ListValidateFileRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListValidateFileRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListValidateFileRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListValidateFileRequest) SetEndTime(v string) *ListValidateFileRequest {
	s.EndTime = &v
	return s
}

func (s *ListValidateFileRequest) SetFileKeyword(v string) *ListValidateFileRequest {
	s.FileKeyword = &v
	return s
}

func (s *ListValidateFileRequest) SetPage(v int32) *ListValidateFileRequest {
	s.Page = &v
	return s
}

func (s *ListValidateFileRequest) SetPageSize(v int32) *ListValidateFileRequest {
	s.PageSize = &v
	return s
}

func (s *ListValidateFileRequest) SetStartTime(v string) *ListValidateFileRequest {
	s.StartTime = &v
	return s
}

func (s *ListValidateFileRequest) Validate() error {
	return dara.Validate(s)
}
