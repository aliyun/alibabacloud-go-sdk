// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *RestoreRevisionRequest
	GetDriveId() *string
	SetFileId(v string) *RestoreRevisionRequest
	GetFileId() *string
	SetRevisionId(v string) *RestoreRevisionRequest
	GetRevisionId() *string
}

type RestoreRevisionRequest struct {
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
	// The version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40CB7794C929
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s RestoreRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreRevisionRequest) GoString() string {
	return s.String()
}

func (s *RestoreRevisionRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *RestoreRevisionRequest) GetFileId() *string {
	return s.FileId
}

func (s *RestoreRevisionRequest) GetRevisionId() *string {
	return s.RevisionId
}

func (s *RestoreRevisionRequest) SetDriveId(v string) *RestoreRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *RestoreRevisionRequest) SetFileId(v string) *RestoreRevisionRequest {
	s.FileId = &v
	return s
}

func (s *RestoreRevisionRequest) SetRevisionId(v string) *RestoreRevisionRequest {
	s.RevisionId = &v
	return s
}

func (s *RestoreRevisionRequest) Validate() error {
	return dara.Validate(s)
}
