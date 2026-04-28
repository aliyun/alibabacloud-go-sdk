// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *RestoreFileResponseBody
	GetAsyncTaskId() *string
	SetDomainId(v string) *RestoreFileResponseBody
	GetDomainId() *string
	SetDriveId(v string) *RestoreFileResponseBody
	GetDriveId() *string
	SetFileId(v string) *RestoreFileResponseBody
	GetFileId() *string
}

type RestoreFileResponseBody struct {
	// The ID of the asynchronous task.
	//
	// If an empty string is returned, the file or folder is restored.
	//
	// If a non-empty string is returned, an asynchronous task is required. You can call the GetAsyncTask operation to obtain the information about an asynchronous task based on the task ID.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
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
	// 4221bf6e6ab43a255edc4463bffa6f5f5d317401
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s RestoreFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreFileResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreFileResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *RestoreFileResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *RestoreFileResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *RestoreFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *RestoreFileResponseBody) SetAsyncTaskId(v string) *RestoreFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *RestoreFileResponseBody) SetDomainId(v string) *RestoreFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *RestoreFileResponseBody) SetDriveId(v string) *RestoreFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *RestoreFileResponseBody) SetFileId(v string) *RestoreFileResponseBody {
	s.FileId = &v
	return s
}

func (s *RestoreFileResponseBody) Validate() error {
	return dara.Validate(s)
}
