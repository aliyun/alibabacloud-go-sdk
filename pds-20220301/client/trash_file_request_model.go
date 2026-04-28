// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrashFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *TrashFileRequest
	GetDriveId() *string
	SetFileId(v string) *TrashFileRequest
	GetFileId() *string
}

type TrashFileRequest struct {
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
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s TrashFileRequest) String() string {
	return dara.Prettify(s)
}

func (s TrashFileRequest) GoString() string {
	return s.String()
}

func (s *TrashFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *TrashFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *TrashFileRequest) SetDriveId(v string) *TrashFileRequest {
	s.DriveId = &v
	return s
}

func (s *TrashFileRequest) SetFileId(v string) *TrashFileRequest {
	s.FileId = &v
	return s
}

func (s *TrashFileRequest) Validate() error {
	return dara.Validate(s)
}
