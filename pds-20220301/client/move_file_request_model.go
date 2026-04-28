// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckNameMode(v string) *MoveFileRequest
	GetCheckNameMode() *string
	SetDriveId(v string) *MoveFileRequest
	GetDriveId() *string
	SetFileId(v string) *MoveFileRequest
	GetFileId() *string
	SetToParentFileId(v string) *MoveFileRequest
	GetToParentFileId() *string
}

type MoveFileRequest struct {
	// The processing method that is used if the file that you want to move has the same name as an existing file in the destination directory. Valid values:
	//
	// ignore: allows you to move the file by using the same name as an existing file in the destination directory.
	//
	// auto_rename: automatically renames the file that has the same name exists in the destination directory. By default, the current point in time is added to the end of the file name. Example: xxx_20060102_150405.
	//
	// refuse: does not move the file that you want to move but returns the information about the file that has the same name in the destination directory.
	//
	// Default value: ignore.
	//
	// example:
	//
	// ignore
	CheckNameMode *string `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
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
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The ID of the destination parent directory to which you want to move a file or folder. If you want to move a file or folder to the root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6520943DC261
	ToParentFileId *string `json:"to_parent_file_id,omitempty" xml:"to_parent_file_id,omitempty"`
}

func (s MoveFileRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveFileRequest) GoString() string {
	return s.String()
}

func (s *MoveFileRequest) GetCheckNameMode() *string {
	return s.CheckNameMode
}

func (s *MoveFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *MoveFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *MoveFileRequest) GetToParentFileId() *string {
	return s.ToParentFileId
}

func (s *MoveFileRequest) SetCheckNameMode(v string) *MoveFileRequest {
	s.CheckNameMode = &v
	return s
}

func (s *MoveFileRequest) SetDriveId(v string) *MoveFileRequest {
	s.DriveId = &v
	return s
}

func (s *MoveFileRequest) SetFileId(v string) *MoveFileRequest {
	s.FileId = &v
	return s
}

func (s *MoveFileRequest) SetToParentFileId(v string) *MoveFileRequest {
	s.ToParentFileId = &v
	return s
}

func (s *MoveFileRequest) Validate() error {
	return dara.Validate(s)
}
