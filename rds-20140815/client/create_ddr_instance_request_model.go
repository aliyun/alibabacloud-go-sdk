// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDdrInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *CreateDdrInstanceRequest
	GetBackupSetId() *string
	SetBackupSetRegion(v string) *CreateDdrInstanceRequest
	GetBackupSetRegion() *string
	SetClientToken(v string) *CreateDdrInstanceRequest
	GetClientToken() *string
	SetConnectionMode(v string) *CreateDdrInstanceRequest
	GetConnectionMode() *string
	SetDBInstanceClass(v string) *CreateDdrInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateDdrInstanceRequest
	GetDBInstanceDescription() *string
	SetDBInstanceNetType(v string) *CreateDdrInstanceRequest
	GetDBInstanceNetType() *string
	SetDBInstanceStorage(v int32) *CreateDdrInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *CreateDdrInstanceRequest
	GetDBInstanceStorageType() *string
	SetEncryptionKey(v string) *CreateDdrInstanceRequest
	GetEncryptionKey() *string
	SetEngine(v string) *CreateDdrInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDdrInstanceRequest
	GetEngineVersion() *string
	SetInstanceNetworkType(v string) *CreateDdrInstanceRequest
	GetInstanceNetworkType() *string
	SetOwnerAccount(v string) *CreateDdrInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDdrInstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateDdrInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDdrInstanceRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *CreateDdrInstanceRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *CreateDdrInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDdrInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDdrInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDdrInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CreateDdrInstanceRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *CreateDdrInstanceRequest
	GetRestoreType() *string
	SetRoleARN(v string) *CreateDdrInstanceRequest
	GetRoleARN() *string
	SetSecurityIPList(v string) *CreateDdrInstanceRequest
	GetSecurityIPList() *string
	SetSourceDBInstanceName(v string) *CreateDdrInstanceRequest
	GetSourceDBInstanceName() *string
	SetSourceRegion(v string) *CreateDdrInstanceRequest
	GetSourceRegion() *string
	SetSystemDBCharset(v string) *CreateDdrInstanceRequest
	GetSystemDBCharset() *string
	SetUsedTime(v string) *CreateDdrInstanceRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateDdrInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDdrInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDdrInstanceRequest
	GetZoneId() *string
}

type CreateDdrInstanceRequest struct {
	BackupSetId           *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	BackupSetRegion       *string `json:"BackupSetRegion,omitempty" xml:"BackupSetRegion,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionMode        *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// This parameter is required.
	DBInstanceNetType     *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	DBInstanceStorage     *int32  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	EncryptionKey         *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	EngineVersion       *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period           *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// This parameter is required.
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	RoleARN     *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// This parameter is required.
	SecurityIPList       *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	SourceRegion         *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	SystemDBCharset      *string `json:"SystemDBCharset,omitempty" xml:"SystemDBCharset,omitempty"`
	UsedTime             *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDdrInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDdrInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDdrInstanceRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateDdrInstanceRequest) GetBackupSetRegion() *string {
	return s.BackupSetRegion
}

func (s *CreateDdrInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDdrInstanceRequest) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *CreateDdrInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateDdrInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDdrInstanceRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *CreateDdrInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CreateDdrInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CreateDdrInstanceRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateDdrInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDdrInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDdrInstanceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CreateDdrInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDdrInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDdrInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDdrInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDdrInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateDdrInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDdrInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDdrInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDdrInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDdrInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateDdrInstanceRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CreateDdrInstanceRequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *CreateDdrInstanceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDdrInstanceRequest) GetSourceDBInstanceName() *string {
	return s.SourceDBInstanceName
}

func (s *CreateDdrInstanceRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *CreateDdrInstanceRequest) GetSystemDBCharset() *string {
	return s.SystemDBCharset
}

func (s *CreateDdrInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDdrInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDdrInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDdrInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDdrInstanceRequest) SetBackupSetId(v string) *CreateDdrInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetBackupSetRegion(v string) *CreateDdrInstanceRequest {
	s.BackupSetRegion = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetClientToken(v string) *CreateDdrInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetConnectionMode(v string) *CreateDdrInstanceRequest {
	s.ConnectionMode = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceClass(v string) *CreateDdrInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceDescription(v string) *CreateDdrInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceNetType(v string) *CreateDdrInstanceRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceStorage(v int32) *CreateDdrInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceStorageType(v string) *CreateDdrInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetEncryptionKey(v string) *CreateDdrInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetEngine(v string) *CreateDdrInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetEngineVersion(v string) *CreateDdrInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetInstanceNetworkType(v string) *CreateDdrInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetOwnerAccount(v string) *CreateDdrInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetOwnerId(v int64) *CreateDdrInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetPayType(v string) *CreateDdrInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetPeriod(v string) *CreateDdrInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetPrivateIpAddress(v string) *CreateDdrInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetRegionId(v string) *CreateDdrInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetResourceGroupId(v string) *CreateDdrInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetResourceOwnerAccount(v string) *CreateDdrInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetResourceOwnerId(v int64) *CreateDdrInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetRestoreTime(v string) *CreateDdrInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetRestoreType(v string) *CreateDdrInstanceRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetRoleARN(v string) *CreateDdrInstanceRequest {
	s.RoleARN = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetSecurityIPList(v string) *CreateDdrInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetSourceDBInstanceName(v string) *CreateDdrInstanceRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetSourceRegion(v string) *CreateDdrInstanceRequest {
	s.SourceRegion = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetSystemDBCharset(v string) *CreateDdrInstanceRequest {
	s.SystemDBCharset = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetUsedTime(v string) *CreateDdrInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetVPCId(v string) *CreateDdrInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetVSwitchId(v string) *CreateDdrInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetZoneId(v string) *CreateDdrInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDdrInstanceRequest) Validate() error {
	return dara.Validate(s)
}
