// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *ResetBackupRequest
	GetBackupId() *string
}

type ResetBackupRequest struct {
	// The ID of the backup.
	//
	// example:
	//
	// backup-fdb897sdfg5****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s ResetBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetBackupRequest) GoString() string {
	return s.String()
}

func (s *ResetBackupRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ResetBackupRequest) SetBackupId(v string) *ResetBackupRequest {
	s.BackupId = &v
	return s
}

func (s *ResetBackupRequest) Validate() error {
	return dara.Validate(s)
}
