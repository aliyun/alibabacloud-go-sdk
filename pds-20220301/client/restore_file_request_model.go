// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *RestoreFileRequest
	GetDriveId() *string
	SetFileId(v string) *RestoreFileRequest
	GetFileId() *string
}

type RestoreFileRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the file or folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43a255edc4463bffa6f5f5d317401
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s RestoreFileRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreFileRequest) GoString() string {
	return s.String()
}

func (s *RestoreFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *RestoreFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *RestoreFileRequest) SetDriveId(v string) *RestoreFileRequest {
	s.DriveId = &v
	return s
}

func (s *RestoreFileRequest) SetFileId(v string) *RestoreFileRequest {
	s.FileId = &v
	return s
}

func (s *RestoreFileRequest) Validate() error {
	return dara.Validate(s)
}
