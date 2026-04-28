// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *CopyFileResponseBody
	GetAsyncTaskId() *string
	SetDomainId(v string) *CopyFileResponseBody
	GetDomainId() *string
	SetDriveId(v string) *CopyFileResponseBody
	GetDriveId() *string
	SetFileId(v string) *CopyFileResponseBody
	GetFileId() *string
}

type CopyFileResponseBody struct {
	// The ID of the asynchronous task.
	//
	// If a file is copied, this parameter is not returned. If a folder is copied, the folder is asynchronously copied in the background and this parameter is returned. You can call the GetAsyncTask operation to query the information about the asynchronous task based on the task ID.
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
	// The ID of the copied file or folder.
	//
	// example:
	//
	// 4221bf6e6ab43a255edc4463bffa6f5f5d317401
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s CopyFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyFileResponseBody) GoString() string {
	return s.String()
}

func (s *CopyFileResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *CopyFileResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *CopyFileResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *CopyFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *CopyFileResponseBody) SetAsyncTaskId(v string) *CopyFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *CopyFileResponseBody) SetDomainId(v string) *CopyFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *CopyFileResponseBody) SetDriveId(v string) *CopyFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *CopyFileResponseBody) SetFileId(v string) *CopyFileResponseBody {
	s.FileId = &v
	return s
}

func (s *CopyFileResponseBody) Validate() error {
	return dara.Validate(s)
}
