// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *MoveFileResponseBody
	GetAsyncTaskId() *string
	SetDomainId(v string) *MoveFileResponseBody
	GetDomainId() *string
	SetDriveId(v string) *MoveFileResponseBody
	GetDriveId() *string
	SetExist(v bool) *MoveFileResponseBody
	GetExist() *bool
	SetFileId(v string) *MoveFileResponseBody
	GetFileId() *string
	SetFileName(v string) *MoveFileResponseBody
	GetFileName() *string
	SetRevisionId(v string) *MoveFileResponseBody
	GetRevisionId() *string
	SetUpdatedAt(v string) *MoveFileResponseBody
	GetUpdatedAt() *string
}

type MoveFileResponseBody struct {
	// The ID of the asynchronous task.
	//
	// If an empty string is returned, the file is moved.
	//
	// If a non-empty string is returned, an asynchronous task is required. You can call the GetAsyncTask operation to obtain the information about an asynchronous task based on the task ID.
	//
	// example:
	//
	// 23ebd1a24dba4166b1527add476ef2866051b4d5del106
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
	// Indicates whether the file already exists in the destination directory.
	//
	// example:
	//
	// false
	Exist *bool `json:"exist,omitempty" xml:"exist,omitempty"`
	// The file ID.
	//
	// example:
	//
	// fileid1
	FileId     *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	FileName   *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	UpdatedAt  *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s MoveFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveFileResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFileResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *MoveFileResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *MoveFileResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *MoveFileResponseBody) GetExist() *bool {
	return s.Exist
}

func (s *MoveFileResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *MoveFileResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *MoveFileResponseBody) GetRevisionId() *string {
	return s.RevisionId
}

func (s *MoveFileResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *MoveFileResponseBody) SetAsyncTaskId(v string) *MoveFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *MoveFileResponseBody) SetDomainId(v string) *MoveFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *MoveFileResponseBody) SetDriveId(v string) *MoveFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *MoveFileResponseBody) SetExist(v bool) *MoveFileResponseBody {
	s.Exist = &v
	return s
}

func (s *MoveFileResponseBody) SetFileId(v string) *MoveFileResponseBody {
	s.FileId = &v
	return s
}

func (s *MoveFileResponseBody) SetFileName(v string) *MoveFileResponseBody {
	s.FileName = &v
	return s
}

func (s *MoveFileResponseBody) SetRevisionId(v string) *MoveFileResponseBody {
	s.RevisionId = &v
	return s
}

func (s *MoveFileResponseBody) SetUpdatedAt(v string) *MoveFileResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *MoveFileResponseBody) Validate() error {
	return dara.Validate(s)
}
