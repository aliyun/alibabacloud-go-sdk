// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkRule interface {
	dara.Model
	String() string
	GoString() string
	SetLinkType(v string) *LinkRule
	GetLinkType() *string
	SetSrcDriveId(v string) *LinkRule
	GetSrcDriveId() *string
	SetSrcDriveName(v string) *LinkRule
	GetSrcDriveName() *string
	SetSrcFileId(v string) *LinkRule
	GetSrcFileId() *string
	SetSrcFileName(v string) *LinkRule
	GetSrcFileName() *string
	SetSrcValid(v bool) *LinkRule
	GetSrcValid() *bool
}

type LinkRule struct {
	LinkType     *string `json:"link_type,omitempty" xml:"link_type,omitempty"`
	SrcDriveId   *string `json:"src_drive_id,omitempty" xml:"src_drive_id,omitempty"`
	SrcDriveName *string `json:"src_drive_name,omitempty" xml:"src_drive_name,omitempty"`
	SrcFileId    *string `json:"src_file_id,omitempty" xml:"src_file_id,omitempty"`
	SrcFileName  *string `json:"src_file_name,omitempty" xml:"src_file_name,omitempty"`
	SrcValid     *bool   `json:"src_valid,omitempty" xml:"src_valid,omitempty"`
}

func (s LinkRule) String() string {
	return dara.Prettify(s)
}

func (s LinkRule) GoString() string {
	return s.String()
}

func (s *LinkRule) GetLinkType() *string {
	return s.LinkType
}

func (s *LinkRule) GetSrcDriveId() *string {
	return s.SrcDriveId
}

func (s *LinkRule) GetSrcDriveName() *string {
	return s.SrcDriveName
}

func (s *LinkRule) GetSrcFileId() *string {
	return s.SrcFileId
}

func (s *LinkRule) GetSrcFileName() *string {
	return s.SrcFileName
}

func (s *LinkRule) GetSrcValid() *bool {
	return s.SrcValid
}

func (s *LinkRule) SetLinkType(v string) *LinkRule {
	s.LinkType = &v
	return s
}

func (s *LinkRule) SetSrcDriveId(v string) *LinkRule {
	s.SrcDriveId = &v
	return s
}

func (s *LinkRule) SetSrcDriveName(v string) *LinkRule {
	s.SrcDriveName = &v
	return s
}

func (s *LinkRule) SetSrcFileId(v string) *LinkRule {
	s.SrcFileId = &v
	return s
}

func (s *LinkRule) SetSrcFileName(v string) *LinkRule {
	s.SrcFileName = &v
	return s
}

func (s *LinkRule) SetSrcValid(v bool) *LinkRule {
	s.SrcValid = &v
	return s
}

func (s *LinkRule) Validate() error {
	return dara.Validate(s)
}
