// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateDdrDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *CheckCreateDdrDBInstanceRequest
	GetBackupSetId() *string
	SetDBInstanceClass(v string) *CheckCreateDdrDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceStorage(v int32) *CheckCreateDdrDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetEngine(v string) *CheckCreateDdrDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CheckCreateDdrDBInstanceRequest
	GetEngineVersion() *string
	SetOwnerId(v int64) *CheckCreateDdrDBInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckCreateDdrDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CheckCreateDdrDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CheckCreateDdrDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckCreateDdrDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CheckCreateDdrDBInstanceRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *CheckCreateDdrDBInstanceRequest
	GetRestoreType() *string
	SetSourceDBInstanceName(v string) *CheckCreateDdrDBInstanceRequest
	GetSourceDBInstanceName() *string
	SetSourceRegion(v string) *CheckCreateDdrDBInstanceRequest
	GetSourceRegion() *string
}

type CheckCreateDdrDBInstanceRequest struct {
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// This parameter is required.
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// This parameter is required.
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// This parameter is required.
	RestoreType          *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	SourceRegion         *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
}

func (s CheckCreateDdrDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateDdrDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CheckCreateDdrDBInstanceRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CheckCreateDdrDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CheckCreateDdrDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CheckCreateDdrDBInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CheckCreateDdrDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CheckCreateDdrDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCreateDdrDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCreateDdrDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckCreateDdrDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckCreateDdrDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckCreateDdrDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CheckCreateDdrDBInstanceRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CheckCreateDdrDBInstanceRequest) GetSourceDBInstanceName() *string {
	return s.SourceDBInstanceName
}

func (s *CheckCreateDdrDBInstanceRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *CheckCreateDdrDBInstanceRequest) SetBackupSetId(v string) *CheckCreateDdrDBInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetDBInstanceClass(v string) *CheckCreateDdrDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetDBInstanceStorage(v int32) *CheckCreateDdrDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetEngine(v string) *CheckCreateDdrDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetEngineVersion(v string) *CheckCreateDdrDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetOwnerId(v int64) *CheckCreateDdrDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetRegionId(v string) *CheckCreateDdrDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetResourceGroupId(v string) *CheckCreateDdrDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetResourceOwnerAccount(v string) *CheckCreateDdrDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetResourceOwnerId(v int64) *CheckCreateDdrDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetRestoreTime(v string) *CheckCreateDdrDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetRestoreType(v string) *CheckCreateDdrDBInstanceRequest {
	s.RestoreType = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetSourceDBInstanceName(v string) *CheckCreateDdrDBInstanceRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetSourceRegion(v string) *CheckCreateDdrDBInstanceRequest {
	s.SourceRegion = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
