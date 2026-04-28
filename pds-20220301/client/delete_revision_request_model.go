// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *DeleteRevisionRequest
	GetDriveId() *string
	SetFileId(v string) *DeleteRevisionRequest
	GetFileId() *string
	SetRevisionId(v string) *DeleteRevisionRequest
	GetRevisionId() *string
}

type DeleteRevisionRequest struct {
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

func (s DeleteRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRevisionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRevisionRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *DeleteRevisionRequest) GetFileId() *string {
	return s.FileId
}

func (s *DeleteRevisionRequest) GetRevisionId() *string {
	return s.RevisionId
}

func (s *DeleteRevisionRequest) SetDriveId(v string) *DeleteRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteRevisionRequest) SetFileId(v string) *DeleteRevisionRequest {
	s.FileId = &v
	return s
}

func (s *DeleteRevisionRequest) SetRevisionId(v string) *DeleteRevisionRequest {
	s.RevisionId = &v
	return s
}

func (s *DeleteRevisionRequest) Validate() error {
	return dara.Validate(s)
}
