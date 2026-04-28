// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonFileItem interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *CommonFileItem
	GetDriveId() *string
	SetFileId(v string) *CommonFileItem
	GetFileId() *string
	SetRevisionId(v string) *CommonFileItem
	GetRevisionId() *string
}

type CommonFileItem struct {
	DriveId    *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId     *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s CommonFileItem) String() string {
	return dara.Prettify(s)
}

func (s CommonFileItem) GoString() string {
	return s.String()
}

func (s *CommonFileItem) GetDriveId() *string {
	return s.DriveId
}

func (s *CommonFileItem) GetFileId() *string {
	return s.FileId
}

func (s *CommonFileItem) GetRevisionId() *string {
	return s.RevisionId
}

func (s *CommonFileItem) SetDriveId(v string) *CommonFileItem {
	s.DriveId = &v
	return s
}

func (s *CommonFileItem) SetFileId(v string) *CommonFileItem {
	s.FileId = &v
	return s
}

func (s *CommonFileItem) SetRevisionId(v string) *CommonFileItem {
	s.RevisionId = &v
	return s
}

func (s *CommonFileItem) Validate() error {
	return dara.Validate(s)
}
