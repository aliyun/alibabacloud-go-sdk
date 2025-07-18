// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMode(v string) *ListBackupJobsRequest
	GetBackupMode() *string
	SetDBInstanceId(v string) *ListBackupJobsRequest
	GetDBInstanceId() *string
}

type ListBackupJobsRequest struct {
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ListBackupJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBackupJobsRequest) GoString() string {
	return s.String()
}

func (s *ListBackupJobsRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *ListBackupJobsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListBackupJobsRequest) SetBackupMode(v string) *ListBackupJobsRequest {
	s.BackupMode = &v
	return s
}

func (s *ListBackupJobsRequest) SetDBInstanceId(v string) *ListBackupJobsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListBackupJobsRequest) Validate() error {
	return dara.Validate(s)
}
