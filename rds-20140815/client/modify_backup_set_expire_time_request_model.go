// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSetExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v int64) *ModifyBackupSetExpireTimeRequest
	GetBackupId() *int64
	SetDBInstanceId(v string) *ModifyBackupSetExpireTimeRequest
	GetDBInstanceId() *string
	SetExpectExpireTime(v string) *ModifyBackupSetExpireTimeRequest
	GetExpectExpireTime() *string
	SetResourceOwnerId(v int64) *ModifyBackupSetExpireTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyBackupSetExpireTimeRequest struct {
	// The backup set ID. You can call the DescribeBackups operation to query the backup set ID. The backup set must meet the following requirements:
	//
	// 	- The Engine parameter is SQLServer
	//
	// 	- The BackupMode parameter is set to Manual.
	//
	// 	- The BackupMethod parameter is set to Physical.
	//
	// 	- The BackupType parameter is set to FullBackup.
	//
	// 	- The BackupStatus parameter is set to Success.
	//
	// This parameter is required.
	//
	// example:
	//
	// b-n8tpg24c6i0v****
	BackupId *int64 `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The point in time to which you want to extend the expiration time of the backup set. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// The time cannot be earlier than the current expiration time. You can call the DescribeBackups operation to view the current expiration time of the backup set.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-17T12:10:23Z
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupSetExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSetExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetExpireTimeRequest) GetBackupId() *int64 {
	return s.BackupId
}

func (s *ModifyBackupSetExpireTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyBackupSetExpireTimeRequest) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *ModifyBackupSetExpireTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupSetExpireTimeRequest) SetBackupId(v int64) *ModifyBackupSetExpireTimeRequest {
	s.BackupId = &v
	return s
}

func (s *ModifyBackupSetExpireTimeRequest) SetDBInstanceId(v string) *ModifyBackupSetExpireTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupSetExpireTimeRequest) SetExpectExpireTime(v string) *ModifyBackupSetExpireTimeRequest {
	s.ExpectExpireTime = &v
	return s
}

func (s *ModifyBackupSetExpireTimeRequest) SetResourceOwnerId(v int64) *ModifyBackupSetExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupSetExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
