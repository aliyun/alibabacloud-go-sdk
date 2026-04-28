// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditLogDetail interface {
	dara.Model
	String() string
	GoString() string
	SetDriveLogDetail(v *DriveLogDetail) *AuditLogDetail
	GetDriveLogDetail() *DriveLogDetail
	SetFileLogDetail(v *FileLogDetail) *AuditLogDetail
	GetFileLogDetail() *FileLogDetail
	SetUserLogDetail(v *UserLogDetail) *AuditLogDetail
	GetUserLogDetail() *UserLogDetail
}

type AuditLogDetail struct {
	DriveLogDetail *DriveLogDetail `json:"drive_log_detail,omitempty" xml:"drive_log_detail,omitempty"`
	FileLogDetail  *FileLogDetail  `json:"file_log_detail,omitempty" xml:"file_log_detail,omitempty"`
	UserLogDetail  *UserLogDetail  `json:"user_log_detail,omitempty" xml:"user_log_detail,omitempty"`
}

func (s AuditLogDetail) String() string {
	return dara.Prettify(s)
}

func (s AuditLogDetail) GoString() string {
	return s.String()
}

func (s *AuditLogDetail) GetDriveLogDetail() *DriveLogDetail {
	return s.DriveLogDetail
}

func (s *AuditLogDetail) GetFileLogDetail() *FileLogDetail {
	return s.FileLogDetail
}

func (s *AuditLogDetail) GetUserLogDetail() *UserLogDetail {
	return s.UserLogDetail
}

func (s *AuditLogDetail) SetDriveLogDetail(v *DriveLogDetail) *AuditLogDetail {
	s.DriveLogDetail = v
	return s
}

func (s *AuditLogDetail) SetFileLogDetail(v *FileLogDetail) *AuditLogDetail {
	s.FileLogDetail = v
	return s
}

func (s *AuditLogDetail) SetUserLogDetail(v *UserLogDetail) *AuditLogDetail {
	s.UserLogDetail = v
	return s
}

func (s *AuditLogDetail) Validate() error {
	if s.DriveLogDetail != nil {
		if err := s.DriveLogDetail.Validate(); err != nil {
			return err
		}
	}
	if s.FileLogDetail != nil {
		if err := s.FileLogDetail.Validate(); err != nil {
			return err
		}
	}
	if s.UserLogDetail != nil {
		if err := s.UserLogDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
