// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *DeleteFileResponseBody
	GetAsyncTaskId() *string
	SetDomainId(v string) *DeleteFileResponseBody
	GetDomainId() *string
	SetDriveId(v string) *DeleteFileResponseBody
	GetDriveId() *string
	SetFileId(v string) *DeleteFileResponseBody
	GetFileId() *string
}

type DeleteFileResponseBody struct {
	// The ID of the asynchronous task. This parameter is returned only in asynchronous processing scenarios. You can call the [GetAsyncTask](https://help.aliyun.com/document_detail/440456.html) operation to query the information about the asynchronous task based on the task ID.
	//
	// example:
	//
	// 000e89fb-cf8f-11e9-8ab4-b6e980803a3b
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *DeleteFileResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *DeleteFileResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *DeleteFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *DeleteFileResponseBody) SetAsyncTaskId(v string) *DeleteFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *DeleteFileResponseBody) SetDomainId(v string) *DeleteFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *DeleteFileResponseBody) SetDriveId(v string) *DeleteFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *DeleteFileResponseBody) SetFileId(v string) *DeleteFileResponseBody {
	s.FileId = &v
	return s
}

func (s *DeleteFileResponseBody) Validate() error {
	return dara.Validate(s)
}
