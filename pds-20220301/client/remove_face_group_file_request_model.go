// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFaceGroupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *RemoveFaceGroupFileRequest
	GetDriveId() *string
	SetFaceGroupId(v string) *RemoveFaceGroupFileRequest
	GetFaceGroupId() *string
	SetFileId(v string) *RemoveFaceGroupFileRequest
	GetFileId() *string
}

type RemoveFaceGroupFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Cluster-abc
	FaceGroupId *string `json:"face_group_id,omitempty" xml:"face_group_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s RemoveFaceGroupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveFaceGroupFileRequest) GoString() string {
	return s.String()
}

func (s *RemoveFaceGroupFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *RemoveFaceGroupFileRequest) GetFaceGroupId() *string {
	return s.FaceGroupId
}

func (s *RemoveFaceGroupFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *RemoveFaceGroupFileRequest) SetDriveId(v string) *RemoveFaceGroupFileRequest {
	s.DriveId = &v
	return s
}

func (s *RemoveFaceGroupFileRequest) SetFaceGroupId(v string) *RemoveFaceGroupFileRequest {
	s.FaceGroupId = &v
	return s
}

func (s *RemoveFaceGroupFileRequest) SetFileId(v string) *RemoveFaceGroupFileRequest {
	s.FileId = &v
	return s
}

func (s *RemoveFaceGroupFileRequest) Validate() error {
	return dara.Validate(s)
}
