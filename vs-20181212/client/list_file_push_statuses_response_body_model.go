// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilePushStatusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListFilePushStatusesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListFilePushStatusesResponseBody
	GetPageSize() *int64
	SetPushStatuses(v []*ListFilePushStatusesResponseBodyPushStatuses) *ListFilePushStatusesResponseBody
	GetPushStatuses() []*ListFilePushStatusesResponseBodyPushStatuses
	SetRequestId(v string) *ListFilePushStatusesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListFilePushStatusesResponseBody
	GetTotalCount() *int64
}

type ListFilePushStatusesResponseBody struct {
	// The page number of the returned results.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of file push status records.
	PushStatuses []*ListFilePushStatusesResponseBodyPushStatuses `json:"PushStatuses,omitempty" xml:"PushStatuses,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of matching file push records.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFilePushStatusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFilePushStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilePushStatusesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListFilePushStatusesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFilePushStatusesResponseBody) GetPushStatuses() []*ListFilePushStatusesResponseBodyPushStatuses {
	return s.PushStatuses
}

func (s *ListFilePushStatusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFilePushStatusesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFilePushStatusesResponseBody) SetPageNumber(v int64) *ListFilePushStatusesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFilePushStatusesResponseBody) SetPageSize(v int64) *ListFilePushStatusesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFilePushStatusesResponseBody) SetPushStatuses(v []*ListFilePushStatusesResponseBodyPushStatuses) *ListFilePushStatusesResponseBody {
	s.PushStatuses = v
	return s
}

func (s *ListFilePushStatusesResponseBody) SetRequestId(v string) *ListFilePushStatusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFilePushStatusesResponseBody) SetTotalCount(v int64) *ListFilePushStatusesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFilePushStatusesResponseBody) Validate() error {
	if s.PushStatuses != nil {
		for _, item := range s.PushStatuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFilePushStatusesResponseBodyPushStatuses struct {
	// The file ID.
	//
	// example:
	//
	// f-1671accd4dafdag3er256cvgewt13f7141db2f7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file name.
	//
	// example:
	//
	// myfile
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The scheduled push start time.
	//
	// example:
	//
	// 2024-03-26T16:32:20+08:00
	PushTime *string `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	// The cloud application service instance ID.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The file push status. Valid values:
	//
	// 1. Created
	//
	// 2. Doing
	//
	// 3. Success: desired state
	//
	// 4. Failed: desired state
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status description.
	//
	// example:
	//
	// push success
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// The latest update time of the status.
	//
	// example:
	//
	// 2024-03-26T17:02:10+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListFilePushStatusesResponseBodyPushStatuses) String() string {
	return dara.Prettify(s)
}

func (s ListFilePushStatusesResponseBodyPushStatuses) GoString() string {
	return s.String()
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) GetFileId() *string {
	return s.FileId
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) GetFileName() *string {
	return s.FileName
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) GetPushTime() *string {
	return s.PushTime
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) GetStatus() *string {
	return s.Status
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) SetFileId(v string) *ListFilePushStatusesResponseBodyPushStatuses {
	s.FileId = &v
	return s
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) SetFileName(v string) *ListFilePushStatusesResponseBodyPushStatuses {
	s.FileName = &v
	return s
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) SetPushTime(v string) *ListFilePushStatusesResponseBodyPushStatuses {
	s.PushTime = &v
	return s
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) SetRenderingInstanceId(v string) *ListFilePushStatusesResponseBodyPushStatuses {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) SetStatus(v string) *ListFilePushStatusesResponseBodyPushStatuses {
	s.Status = &v
	return s
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) SetStatusDescription(v string) *ListFilePushStatusesResponseBodyPushStatuses {
	s.StatusDescription = &v
	return s
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) SetUpdateTime(v string) *ListFilePushStatusesResponseBodyPushStatuses {
	s.UpdateTime = &v
	return s
}

func (s *ListFilePushStatusesResponseBodyPushStatuses) Validate() error {
	return dara.Validate(s)
}
