// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupType(v string) *CreateBackupRequest
	GetBackupType() *string
	SetDBInstanceName(v string) *CreateBackupRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateBackupRequest
	GetRegionId() *string
}

type CreateBackupRequest struct {
	// example:
	//
	// 0
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *CreateBackupRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateBackupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBackupRequest) SetBackupType(v string) *CreateBackupRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupRequest) SetDBInstanceName(v string) *CreateBackupRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateBackupRequest) SetRegionId(v string) *CreateBackupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBackupRequest) Validate() error {
	return dara.Validate(s)
}
