// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRename(v bool) *CopyFileRequest
	GetAutoRename() *bool
	SetDriveId(v string) *CopyFileRequest
	GetDriveId() *string
	SetFileId(v string) *CopyFileRequest
	GetFileId() *string
	SetShareId(v string) *CopyFileRequest
	GetShareId() *string
	SetToDriveId(v string) *CopyFileRequest
	GetToDriveId() *string
	SetToParentFileId(v string) *CopyFileRequest
	GetToParentFileId() *string
}

type CopyFileRequest struct {
	// Specifies whether to automatically rename the file if the file name already exists in the destination folder. Default value: false.
	//
	// example:
	//
	// true
	AutoRename *bool `json:"auto_rename,omitempty" xml:"auto_rename,omitempty"`
	// The drive ID.
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
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID. If you want to manage a file by using a share link, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The ID of the drive to which you want to copy the file or folder. Default value: the value of drive_id.
	//
	// example:
	//
	// 1
	ToDriveId *string `json:"to_drive_id,omitempty" xml:"to_drive_id,omitempty"`
	// The ID of the destination parent folder. If you want to copy the file or folder to a root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6520943DC261
	ToParentFileId *string `json:"to_parent_file_id,omitempty" xml:"to_parent_file_id,omitempty"`
}

func (s CopyFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyFileRequest) GoString() string {
	return s.String()
}

func (s *CopyFileRequest) GetAutoRename() *bool {
	return s.AutoRename
}

func (s *CopyFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *CopyFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *CopyFileRequest) GetShareId() *string {
	return s.ShareId
}

func (s *CopyFileRequest) GetToDriveId() *string {
	return s.ToDriveId
}

func (s *CopyFileRequest) GetToParentFileId() *string {
	return s.ToParentFileId
}

func (s *CopyFileRequest) SetAutoRename(v bool) *CopyFileRequest {
	s.AutoRename = &v
	return s
}

func (s *CopyFileRequest) SetDriveId(v string) *CopyFileRequest {
	s.DriveId = &v
	return s
}

func (s *CopyFileRequest) SetFileId(v string) *CopyFileRequest {
	s.FileId = &v
	return s
}

func (s *CopyFileRequest) SetShareId(v string) *CopyFileRequest {
	s.ShareId = &v
	return s
}

func (s *CopyFileRequest) SetToDriveId(v string) *CopyFileRequest {
	s.ToDriveId = &v
	return s
}

func (s *CopyFileRequest) SetToParentFileId(v string) *CopyFileRequest {
	s.ToParentFileId = &v
	return s
}

func (s *CopyFileRequest) Validate() error {
	return dara.Validate(s)
}
