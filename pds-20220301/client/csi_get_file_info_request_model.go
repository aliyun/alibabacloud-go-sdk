// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCsiGetFileInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *CsiGetFileInfoRequest
	GetDriveId() *string
	SetFileId(v string) *CsiGetFileInfoRequest
	GetFileId() *string
	SetUrlExpireSec(v int32) *CsiGetFileInfoRequest
	GetUrlExpireSec() *int32
}

type CsiGetFileInfoRequest struct {
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
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 100
	UrlExpireSec *int32 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
}

func (s CsiGetFileInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s CsiGetFileInfoRequest) GoString() string {
	return s.String()
}

func (s *CsiGetFileInfoRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *CsiGetFileInfoRequest) GetFileId() *string {
	return s.FileId
}

func (s *CsiGetFileInfoRequest) GetUrlExpireSec() *int32 {
	return s.UrlExpireSec
}

func (s *CsiGetFileInfoRequest) SetDriveId(v string) *CsiGetFileInfoRequest {
	s.DriveId = &v
	return s
}

func (s *CsiGetFileInfoRequest) SetFileId(v string) *CsiGetFileInfoRequest {
	s.FileId = &v
	return s
}

func (s *CsiGetFileInfoRequest) SetUrlExpireSec(v int32) *CsiGetFileInfoRequest {
	s.UrlExpireSec = &v
	return s
}

func (s *CsiGetFileInfoRequest) Validate() error {
	return dara.Validate(s)
}
