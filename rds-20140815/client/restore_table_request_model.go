// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *RestoreTableRequest
	GetBackupId() *string
	SetClientToken(v string) *RestoreTableRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RestoreTableRequest
	GetDBInstanceId() *string
	SetInstantRecovery(v bool) *RestoreTableRequest
	GetInstantRecovery() *bool
	SetOwnerAccount(v string) *RestoreTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestoreTableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestoreTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestoreTableRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *RestoreTableRequest
	GetRestoreTime() *string
	SetTableMeta(v string) *RestoreTableRequest
	GetTableMeta() *string
}

type RestoreTableRequest struct {
	BackupId    *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	InstantRecovery      *bool   `json:"InstantRecovery,omitempty" xml:"InstantRecovery,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// This parameter is required.
	TableMeta *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
}

func (s RestoreTableRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreTableRequest) GoString() string {
	return s.String()
}

func (s *RestoreTableRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *RestoreTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestoreTableRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestoreTableRequest) GetInstantRecovery() *bool {
	return s.InstantRecovery
}

func (s *RestoreTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestoreTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestoreTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestoreTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestoreTableRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *RestoreTableRequest) GetTableMeta() *string {
	return s.TableMeta
}

func (s *RestoreTableRequest) SetBackupId(v string) *RestoreTableRequest {
	s.BackupId = &v
	return s
}

func (s *RestoreTableRequest) SetClientToken(v string) *RestoreTableRequest {
	s.ClientToken = &v
	return s
}

func (s *RestoreTableRequest) SetDBInstanceId(v string) *RestoreTableRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestoreTableRequest) SetInstantRecovery(v bool) *RestoreTableRequest {
	s.InstantRecovery = &v
	return s
}

func (s *RestoreTableRequest) SetOwnerAccount(v string) *RestoreTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestoreTableRequest) SetOwnerId(v int64) *RestoreTableRequest {
	s.OwnerId = &v
	return s
}

func (s *RestoreTableRequest) SetResourceOwnerAccount(v string) *RestoreTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestoreTableRequest) SetResourceOwnerId(v int64) *RestoreTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreTableRequest) SetRestoreTime(v string) *RestoreTableRequest {
	s.RestoreTime = &v
	return s
}

func (s *RestoreTableRequest) SetTableMeta(v string) *RestoreTableRequest {
	s.TableMeta = &v
	return s
}

func (s *RestoreTableRequest) Validate() error {
	return dara.Validate(s)
}
