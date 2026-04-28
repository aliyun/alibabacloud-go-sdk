// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileIDInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *FileIDInfo
	GetDriveId() *string
	SetFileId(v string) *FileIDInfo
	GetFileId() *string
	SetType(v string) *FileIDInfo
	GetType() *string
}

type FileIDInfo struct {
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId  *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FileIDInfo) String() string {
	return dara.Prettify(s)
}

func (s FileIDInfo) GoString() string {
	return s.String()
}

func (s *FileIDInfo) GetDriveId() *string {
	return s.DriveId
}

func (s *FileIDInfo) GetFileId() *string {
	return s.FileId
}

func (s *FileIDInfo) GetType() *string {
	return s.Type
}

func (s *FileIDInfo) SetDriveId(v string) *FileIDInfo {
	s.DriveId = &v
	return s
}

func (s *FileIDInfo) SetFileId(v string) *FileIDInfo {
	s.FileId = &v
	return s
}

func (s *FileIDInfo) SetType(v string) *FileIDInfo {
	s.Type = &v
	return s
}

func (s *FileIDInfo) Validate() error {
	return dara.Validate(s)
}
