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
