// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileListPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *FileListPermissionRequest
	GetDriveId() *string
	SetFileId(v string) *FileListPermissionRequest
	GetFileId() *string
}

type FileListPermissionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43a255edc4463bffa6f5f5d317401
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s FileListPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s FileListPermissionRequest) GoString() string {
	return s.String()
}

func (s *FileListPermissionRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *FileListPermissionRequest) GetFileId() *string {
	return s.FileId
}

func (s *FileListPermissionRequest) SetDriveId(v string) *FileListPermissionRequest {
	s.DriveId = &v
	return s
}

func (s *FileListPermissionRequest) SetFileId(v string) *FileListPermissionRequest {
	s.FileId = &v
	return s
}

func (s *FileListPermissionRequest) Validate() error {
	return dara.Validate(s)
}
