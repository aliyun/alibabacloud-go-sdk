// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *DeleteFileRequest
	GetDriveId() *string
	SetFileId(v string) *DeleteFileRequest
	GetFileId() *string
}

type DeleteFileRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID or folder ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *DeleteFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteFileRequest) SetDriveId(v string) *DeleteFileRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteFileRequest) SetFileId(v string) *DeleteFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteFileRequest) Validate() error {
	return dara.Validate(s)
}
