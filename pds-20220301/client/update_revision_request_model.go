// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *UpdateRevisionRequest
	GetDriveId() *string
	SetFileId(v string) *UpdateRevisionRequest
	GetFileId() *string
	SetKeepForever(v bool) *UpdateRevisionRequest
	GetKeepForever() *bool
	SetRevisionDescription(v string) *UpdateRevisionRequest
	GetRevisionDescription() *string
	SetRevisionId(v string) *UpdateRevisionRequest
	GetRevisionId() *string
}

type UpdateRevisionRequest struct {
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
	// Specifies whether to permanently retain a version.
	//
	// By default, this parameter is not specified, which indicates that you do not modify the permanent retention configuration of the version.
	//
	// example:
	//
	// true
	KeepForever *bool `json:"keep_forever,omitempty" xml:"keep_forever,omitempty"`
	// The description of the version. The description can be up to 1,024 characters in length.
	//
	// By default, this parameter is not specified, which indicates that you do not modify the description of the version.
	//
	// example:
	//
	// aaa
	RevisionDescription *string `json:"revision_description,omitempty" xml:"revision_description,omitempty"`
	// The version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40CB7794C929
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s UpdateRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRevisionRequest) GoString() string {
	return s.String()
}

func (s *UpdateRevisionRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *UpdateRevisionRequest) GetFileId() *string {
	return s.FileId
}

func (s *UpdateRevisionRequest) GetKeepForever() *bool {
	return s.KeepForever
}

func (s *UpdateRevisionRequest) GetRevisionDescription() *string {
	return s.RevisionDescription
}

func (s *UpdateRevisionRequest) GetRevisionId() *string {
	return s.RevisionId
}

func (s *UpdateRevisionRequest) SetDriveId(v string) *UpdateRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateRevisionRequest) SetFileId(v string) *UpdateRevisionRequest {
	s.FileId = &v
	return s
}

func (s *UpdateRevisionRequest) SetKeepForever(v bool) *UpdateRevisionRequest {
	s.KeepForever = &v
	return s
}

func (s *UpdateRevisionRequest) SetRevisionDescription(v string) *UpdateRevisionRequest {
	s.RevisionDescription = &v
	return s
}

func (s *UpdateRevisionRequest) SetRevisionId(v string) *UpdateRevisionRequest {
	s.RevisionId = &v
	return s
}

func (s *UpdateRevisionRequest) Validate() error {
	return dara.Validate(s)
}
