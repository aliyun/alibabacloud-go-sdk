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
	// The backup set ID.
	//
	// >  You can call the [DescribeDataBackups](https://help.aliyun.com/document_detail/210093.html) operation to query the IDs of all backup sets in an instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
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
