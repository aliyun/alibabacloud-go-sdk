// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreDdrTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *RestoreDdrTableRequest
	GetBackupId() *string
	SetClientToken(v string) *RestoreDdrTableRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RestoreDdrTableRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *RestoreDdrTableRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RestoreDdrTableRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RestoreDdrTableRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RestoreDdrTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestoreDdrTableRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *RestoreDdrTableRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *RestoreDdrTableRequest
	GetRestoreType() *string
	SetSourceDBInstanceName(v string) *RestoreDdrTableRequest
	GetSourceDBInstanceName() *string
	SetSourceRegion(v string) *RestoreDdrTableRequest
	GetSourceRegion() *string
	SetTableMeta(v string) *RestoreDdrTableRequest
	GetTableMeta() *string
}

type RestoreDdrTableRequest struct {
	BackupId    *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// This parameter is required.
	RestoreType          *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	SourceRegion         *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// This parameter is required.
	TableMeta *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
}

func (s RestoreDdrTableRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreDdrTableRequest) GoString() string {
	return s.String()
}

func (s *RestoreDdrTableRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *RestoreDdrTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestoreDdrTableRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestoreDdrTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestoreDdrTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestoreDdrTableRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RestoreDdrTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestoreDdrTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestoreDdrTableRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *RestoreDdrTableRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *RestoreDdrTableRequest) GetSourceDBInstanceName() *string {
	return s.SourceDBInstanceName
}

func (s *RestoreDdrTableRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *RestoreDdrTableRequest) GetTableMeta() *string {
	return s.TableMeta
}

func (s *RestoreDdrTableRequest) SetBackupId(v string) *RestoreDdrTableRequest {
	s.BackupId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetClientToken(v string) *RestoreDdrTableRequest {
	s.ClientToken = &v
	return s
}

func (s *RestoreDdrTableRequest) SetDBInstanceId(v string) *RestoreDdrTableRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetOwnerId(v int64) *RestoreDdrTableRequest {
	s.OwnerId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetRegionId(v string) *RestoreDdrTableRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetResourceGroupId(v string) *RestoreDdrTableRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetResourceOwnerAccount(v string) *RestoreDdrTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestoreDdrTableRequest) SetResourceOwnerId(v int64) *RestoreDdrTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetRestoreTime(v string) *RestoreDdrTableRequest {
	s.RestoreTime = &v
	return s
}

func (s *RestoreDdrTableRequest) SetRestoreType(v string) *RestoreDdrTableRequest {
	s.RestoreType = &v
	return s
}

func (s *RestoreDdrTableRequest) SetSourceDBInstanceName(v string) *RestoreDdrTableRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *RestoreDdrTableRequest) SetSourceRegion(v string) *RestoreDdrTableRequest {
	s.SourceRegion = &v
	return s
}

func (s *RestoreDdrTableRequest) SetTableMeta(v string) *RestoreDdrTableRequest {
	s.TableMeta = &v
	return s
}

func (s *RestoreDdrTableRequest) Validate() error {
	return dara.Validate(s)
}
