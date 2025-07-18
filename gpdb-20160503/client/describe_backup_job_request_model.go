// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobId(v int64) *DescribeBackupJobRequest
	GetBackupJobId() *int64
	SetDBInstanceId(v string) *DescribeBackupJobRequest
	GetDBInstanceId() *string
}

type DescribeBackupJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	BackupJobId *int64 `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeBackupJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobRequest) GetBackupJobId() *int64 {
	return s.BackupJobId
}

func (s *DescribeBackupJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupJobRequest) SetBackupJobId(v int64) *DescribeBackupJobRequest {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupJobRequest) SetDBInstanceId(v string) *DescribeBackupJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupJobRequest) Validate() error {
	return dara.Validate(s)
}
