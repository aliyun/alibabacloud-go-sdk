// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPunishFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionCode(v string) *PunishFileRequest
	GetActionCode() *string
	SetDriveId(v string) *PunishFileRequest
	GetDriveId() *string
	SetFileId(v string) *PunishFileRequest
	GetFileId() *string
	SetPunishReason(v string) *PunishFileRequest
	GetPunishReason() *string
}

type PunishFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pds_file_meta_frozen
	ActionCode *string `json:"action_code,omitempty" xml:"action_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2498DZ2
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// illegal
	PunishReason *string `json:"punish_reason,omitempty" xml:"punish_reason,omitempty"`
}

func (s PunishFileRequest) String() string {
	return dara.Prettify(s)
}

func (s PunishFileRequest) GoString() string {
	return s.String()
}

func (s *PunishFileRequest) GetActionCode() *string {
	return s.ActionCode
}

func (s *PunishFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *PunishFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *PunishFileRequest) GetPunishReason() *string {
	return s.PunishReason
}

func (s *PunishFileRequest) SetActionCode(v string) *PunishFileRequest {
	s.ActionCode = &v
	return s
}

func (s *PunishFileRequest) SetDriveId(v string) *PunishFileRequest {
	s.DriveId = &v
	return s
}

func (s *PunishFileRequest) SetFileId(v string) *PunishFileRequest {
	s.FileId = &v
	return s
}

func (s *PunishFileRequest) SetPunishReason(v string) *PunishFileRequest {
	s.PunishReason = &v
	return s
}

func (s *PunishFileRequest) Validate() error {
	return dara.Validate(s)
}
