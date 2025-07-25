// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *GetBackupRequest
	GetBackupId() *string
}

type GetBackupRequest struct {
	// The ID of the backup.
	//
	// example:
	//
	// backup-fdb897sdf****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s GetBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBackupRequest) GoString() string {
	return s.String()
}

func (s *GetBackupRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *GetBackupRequest) SetBackupId(v string) *GetBackupRequest {
	s.BackupId = &v
	return s
}

func (s *GetBackupRequest) Validate() error {
	return dara.Validate(s)
}
