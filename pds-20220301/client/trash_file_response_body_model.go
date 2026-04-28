// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrashFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *TrashFileResponseBody
	GetAsyncTaskId() *string
	SetDomainId(v string) *TrashFileResponseBody
	GetDomainId() *string
	SetDriveId(v string) *TrashFileResponseBody
	GetDriveId() *string
	SetFileId(v string) *TrashFileResponseBody
	GetFileId() *string
}

type TrashFileResponseBody struct {
	// The ID of the asynchronous task.
	//
	// If an empty string is returned, the file or folder is moved to the recycle bin.
	//
	// If a non-empty string is returned, an asynchronous task is required. You can call the GetAsyncTask operation to obtain the information about an asynchronous task based on the task ID.
	//
	// example:
	//
	// 13ebd3a24dba4166b1527add676ef2866051b4d5dele16
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
	// The ID of the file or folder.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s TrashFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrashFileResponseBody) GoString() string {
	return s.String()
}

func (s *TrashFileResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *TrashFileResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *TrashFileResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *TrashFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *TrashFileResponseBody) SetAsyncTaskId(v string) *TrashFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *TrashFileResponseBody) SetDomainId(v string) *TrashFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *TrashFileResponseBody) SetDriveId(v string) *TrashFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *TrashFileResponseBody) SetFileId(v string) *TrashFileResponseBody {
	s.FileId = &v
	return s
}

func (s *TrashFileResponseBody) Validate() error {
	return dara.Validate(s)
}
