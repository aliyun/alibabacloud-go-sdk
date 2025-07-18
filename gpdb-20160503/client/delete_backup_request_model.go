// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DeleteBackupRequest
	GetBackupId() *string
	SetDBInstanceId(v string) *DeleteBackupRequest
	GetDBInstanceId() *string
}

type DeleteBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DeleteBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DeleteBackupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteBackupRequest) SetBackupId(v string) *DeleteBackupRequest {
	s.BackupId = &v
	return s
}

func (s *DeleteBackupRequest) SetDBInstanceId(v string) *DeleteBackupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteBackupRequest) Validate() error {
	return dara.Validate(s)
}
