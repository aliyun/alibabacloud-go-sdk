// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilePushStatusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListFilePushStatusesRequest
	GetEndTime() *string
	SetFileId(v string) *ListFilePushStatusesRequest
	GetFileId() *string
	SetFileName(v string) *ListFilePushStatusesRequest
	GetFileName() *string
	SetPageNumber(v int64) *ListFilePushStatusesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListFilePushStatusesRequest
	GetPageSize() *int64
	SetRenderingInstanceId(v string) *ListFilePushStatusesRequest
	GetRenderingInstanceId() *string
	SetStartTime(v string) *ListFilePushStatusesRequest
	GetStartTime() *string
}

type ListFilePushStatusesRequest struct {
	// The start time for the time range filter. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-06-23T02:13:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the file, which uniquely identifies the file.
	//
	// example:
	//
	// f-1671accd4dafdag3er256cvgewt13f7141db2f7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// myfile
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The page number of the list to query. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The Settings for the number of entries per page during paging in a paged query. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The instance ID of the cloud application service, which is used to query the files installed on the specified instance.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The start time for the time range filter. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-06-17T12:16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListFilePushStatusesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFilePushStatusesRequest) GoString() string {
	return s.String()
}

func (s *ListFilePushStatusesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListFilePushStatusesRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListFilePushStatusesRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListFilePushStatusesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListFilePushStatusesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFilePushStatusesRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListFilePushStatusesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListFilePushStatusesRequest) SetEndTime(v string) *ListFilePushStatusesRequest {
	s.EndTime = &v
	return s
}

func (s *ListFilePushStatusesRequest) SetFileId(v string) *ListFilePushStatusesRequest {
	s.FileId = &v
	return s
}

func (s *ListFilePushStatusesRequest) SetFileName(v string) *ListFilePushStatusesRequest {
	s.FileName = &v
	return s
}

func (s *ListFilePushStatusesRequest) SetPageNumber(v int64) *ListFilePushStatusesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFilePushStatusesRequest) SetPageSize(v int64) *ListFilePushStatusesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFilePushStatusesRequest) SetRenderingInstanceId(v string) *ListFilePushStatusesRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListFilePushStatusesRequest) SetStartTime(v string) *ListFilePushStatusesRequest {
	s.StartTime = &v
	return s
}

func (s *ListFilePushStatusesRequest) Validate() error {
	return dara.Validate(s)
}
