// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *CloneDBInstanceRequest
	GetBackupId() *string
	SetDBInstanceId(v string) *CloneDBInstanceRequest
	GetDBInstanceId() *string
	SetSrcDbInstanceName(v string) *CloneDBInstanceRequest
	GetSrcDbInstanceName() *string
}

type CloneDBInstanceRequest struct {
	// The ID of the backup set.
	//
	// > You can call the [DescribeDataBackups](https://help.aliyun.com/document_detail/210093.html) operation to query the IDs of all backup sets of the instance. Only snapshot backup sets are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the new instance.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the source instance.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	SrcDbInstanceName *string `json:"SrcDbInstanceName,omitempty" xml:"SrcDbInstanceName,omitempty"`
}

func (s CloneDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CloneDBInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CloneDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CloneDBInstanceRequest) GetSrcDbInstanceName() *string {
	return s.SrcDbInstanceName
}

func (s *CloneDBInstanceRequest) SetBackupId(v string) *CloneDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDBInstanceId(v string) *CloneDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetSrcDbInstanceName(v string) *CloneDBInstanceRequest {
	s.SrcDbInstanceName = &v
	return s
}

func (s *CloneDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
