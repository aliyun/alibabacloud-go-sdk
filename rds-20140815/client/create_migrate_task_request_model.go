// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMode(v string) *CreateMigrateTaskRequest
	GetBackupMode() *string
	SetCheckDBMode(v string) *CreateMigrateTaskRequest
	GetCheckDBMode() *string
	SetDBInstanceId(v string) *CreateMigrateTaskRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateMigrateTaskRequest
	GetDBName() *string
	SetIsOnlineDB(v string) *CreateMigrateTaskRequest
	GetIsOnlineDB() *string
	SetMigrateTaskId(v string) *CreateMigrateTaskRequest
	GetMigrateTaskId() *string
	SetOSSUrls(v string) *CreateMigrateTaskRequest
	GetOSSUrls() *string
	SetOssObjectPositions(v string) *CreateMigrateTaskRequest
	GetOssObjectPositions() *string
	SetOwnerId(v int64) *CreateMigrateTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateMigrateTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMigrateTaskRequest
	GetResourceOwnerId() *int64
}

type CreateMigrateTaskRequest struct {
	// This parameter is required.
	BackupMode  *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	CheckDBMode *string `json:"CheckDBMode,omitempty" xml:"CheckDBMode,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// This parameter is required.
	IsOnlineDB           *string `json:"IsOnlineDB,omitempty" xml:"IsOnlineDB,omitempty"`
	MigrateTaskId        *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	OSSUrls              *string `json:"OSSUrls,omitempty" xml:"OSSUrls,omitempty"`
	OssObjectPositions   *string `json:"OssObjectPositions,omitempty" xml:"OssObjectPositions,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMigrateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMigrateTaskRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *CreateMigrateTaskRequest) GetCheckDBMode() *string {
	return s.CheckDBMode
}

func (s *CreateMigrateTaskRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateMigrateTaskRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateMigrateTaskRequest) GetIsOnlineDB() *string {
	return s.IsOnlineDB
}

func (s *CreateMigrateTaskRequest) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *CreateMigrateTaskRequest) GetOSSUrls() *string {
	return s.OSSUrls
}

func (s *CreateMigrateTaskRequest) GetOssObjectPositions() *string {
	return s.OssObjectPositions
}

func (s *CreateMigrateTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMigrateTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMigrateTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMigrateTaskRequest) SetBackupMode(v string) *CreateMigrateTaskRequest {
	s.BackupMode = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetCheckDBMode(v string) *CreateMigrateTaskRequest {
	s.CheckDBMode = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetDBInstanceId(v string) *CreateMigrateTaskRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetDBName(v string) *CreateMigrateTaskRequest {
	s.DBName = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetIsOnlineDB(v string) *CreateMigrateTaskRequest {
	s.IsOnlineDB = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetMigrateTaskId(v string) *CreateMigrateTaskRequest {
	s.MigrateTaskId = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetOSSUrls(v string) *CreateMigrateTaskRequest {
	s.OSSUrls = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetOssObjectPositions(v string) *CreateMigrateTaskRequest {
	s.OssObjectPositions = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetOwnerId(v int64) *CreateMigrateTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetResourceOwnerAccount(v string) *CreateMigrateTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMigrateTaskRequest) SetResourceOwnerId(v int64) *CreateMigrateTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMigrateTaskRequest) Validate() error {
	return dara.Validate(s)
}
