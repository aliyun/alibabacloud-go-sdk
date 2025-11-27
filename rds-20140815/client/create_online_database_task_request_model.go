// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineDatabaseTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckDBMode(v string) *CreateOnlineDatabaseTaskRequest
	GetCheckDBMode() *string
	SetClientToken(v string) *CreateOnlineDatabaseTaskRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CreateOnlineDatabaseTaskRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateOnlineDatabaseTaskRequest
	GetDBName() *string
	SetMigrateTaskId(v string) *CreateOnlineDatabaseTaskRequest
	GetMigrateTaskId() *string
	SetOwnerAccount(v string) *CreateOnlineDatabaseTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateOnlineDatabaseTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateOnlineDatabaseTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateOnlineDatabaseTaskRequest
	GetResourceOwnerId() *int64
}

type CreateOnlineDatabaseTaskRequest struct {
	// The consistency check method after the database is open. Valid values:
	//
	// 	- **SyncExecuteDBCheck**: synchronous database check
	//
	// 	- **AsyncExecuteDBCheck**: asynchronous database check
	//
	// > The check methods are supported for RDS instances that run SQL Server 2008 R2.
	//
	// This parameter is required.
	//
	// example:
	//
	// AsyncExecuteDBCheck
	CheckDBMode *string `json:"CheckDBMode,omitempty" xml:"CheckDBMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The ID of the migration task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5652255443
	MigrateTaskId        *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateOnlineDatabaseTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineDatabaseTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOnlineDatabaseTaskRequest) GetCheckDBMode() *string {
	return s.CheckDBMode
}

func (s *CreateOnlineDatabaseTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOnlineDatabaseTaskRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateOnlineDatabaseTaskRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateOnlineDatabaseTaskRequest) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *CreateOnlineDatabaseTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateOnlineDatabaseTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateOnlineDatabaseTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateOnlineDatabaseTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateOnlineDatabaseTaskRequest) SetCheckDBMode(v string) *CreateOnlineDatabaseTaskRequest {
	s.CheckDBMode = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) SetClientToken(v string) *CreateOnlineDatabaseTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) SetDBInstanceId(v string) *CreateOnlineDatabaseTaskRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) SetDBName(v string) *CreateOnlineDatabaseTaskRequest {
	s.DBName = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) SetMigrateTaskId(v string) *CreateOnlineDatabaseTaskRequest {
	s.MigrateTaskId = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) SetOwnerAccount(v string) *CreateOnlineDatabaseTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) SetOwnerId(v int64) *CreateOnlineDatabaseTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) SetResourceOwnerAccount(v string) *CreateOnlineDatabaseTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) SetResourceOwnerId(v int64) *CreateOnlineDatabaseTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateOnlineDatabaseTaskRequest) Validate() error {
	return dara.Validate(s)
}
