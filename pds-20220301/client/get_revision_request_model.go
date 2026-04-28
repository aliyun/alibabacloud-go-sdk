// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *GetRevisionRequest
	GetDriveId() *string
	SetFields(v string) *GetRevisionRequest
	GetFields() *string
	SetFileId(v string) *GetRevisionRequest
	GetFileId() *string
	SetRevisionId(v string) *GetRevisionRequest
	GetRevisionId() *string
	SetUrlExpireSec(v int64) *GetRevisionRequest
	GetUrlExpireSec() *int64
}

type GetRevisionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Specifies the returned fields.
	//
	// By default, this parameter is left empty. If you set this parameter to \\*, all fields are returned. If you leave this parameter empty, the creator of the file is not returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
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
	// The validity period of the file download or preview. Valid values: 10 to 86400.
	//
	// Default value: 900. Unit: seconds.
	//
	// example:
	//
	// 900
	UrlExpireSec *int64 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
}

func (s GetRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRevisionRequest) GoString() string {
	return s.String()
}

func (s *GetRevisionRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetRevisionRequest) GetFields() *string {
	return s.Fields
}

func (s *GetRevisionRequest) GetFileId() *string {
	return s.FileId
}

func (s *GetRevisionRequest) GetRevisionId() *string {
	return s.RevisionId
}

func (s *GetRevisionRequest) GetUrlExpireSec() *int64 {
	return s.UrlExpireSec
}

func (s *GetRevisionRequest) SetDriveId(v string) *GetRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *GetRevisionRequest) SetFields(v string) *GetRevisionRequest {
	s.Fields = &v
	return s
}

func (s *GetRevisionRequest) SetFileId(v string) *GetRevisionRequest {
	s.FileId = &v
	return s
}

func (s *GetRevisionRequest) SetRevisionId(v string) *GetRevisionRequest {
	s.RevisionId = &v
	return s
}

func (s *GetRevisionRequest) SetUrlExpireSec(v int64) *GetRevisionRequest {
	s.UrlExpireSec = &v
	return s
}

func (s *GetRevisionRequest) Validate() error {
	return dara.Validate(s)
}
