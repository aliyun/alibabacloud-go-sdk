// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListFilesRequest
	GetEndTime() *string
	SetFileId(v string) *ListFilesRequest
	GetFileId() *string
	SetFileName(v string) *ListFilesRequest
	GetFileName() *string
	SetPageNumber(v int64) *ListFilesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListFilesRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListFilesRequest
	GetStartTime() *string
}

type ListFilesRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// f-1671accd4dafdag3er256cvgewt13f7141db2f7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// myfile
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListFilesRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListFilesRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListFilesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListFilesRequest) SetEndTime(v string) *ListFilesRequest {
	s.EndTime = &v
	return s
}

func (s *ListFilesRequest) SetFileId(v string) *ListFilesRequest {
	s.FileId = &v
	return s
}

func (s *ListFilesRequest) SetFileName(v string) *ListFilesRequest {
	s.FileName = &v
	return s
}

func (s *ListFilesRequest) SetPageNumber(v int64) *ListFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFilesRequest) SetPageSize(v int64) *ListFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFilesRequest) SetStartTime(v string) *ListFilesRequest {
	s.StartTime = &v
	return s
}

func (s *ListFilesRequest) Validate() error {
	return dara.Validate(s)
}
