// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *RecoveryDBInstanceRequest
	GetBackupId() *string
	SetDBInstanceClass(v string) *RecoveryDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *RecoveryDBInstanceRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *RecoveryDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *RecoveryDBInstanceRequest
	GetDBInstanceStorageType() *string
	SetDbNames(v string) *RecoveryDBInstanceRequest
	GetDbNames() *string
	SetInstanceNetworkType(v string) *RecoveryDBInstanceRequest
	GetInstanceNetworkType() *string
	SetPayType(v string) *RecoveryDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *RecoveryDBInstanceRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *RecoveryDBInstanceRequest
	GetPrivateIpAddress() *string
	SetResourceOwnerId(v int64) *RecoveryDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *RecoveryDBInstanceRequest
	GetRestoreTime() *string
	SetTargetDBInstanceId(v string) *RecoveryDBInstanceRequest
	GetTargetDBInstanceId() *string
	SetUsedTime(v string) *RecoveryDBInstanceRequest
	GetUsedTime() *string
	SetVPCId(v string) *RecoveryDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *RecoveryDBInstanceRequest
	GetVSwitchId() *string
}

type RecoveryDBInstanceRequest struct {
	BackupId              *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStorage     *int32  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// This parameter is required.
	DbNames             *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	PayType             *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period              *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PrivateIpAddress    *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerId     *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestoreTime         *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	TargetDBInstanceId  *string `json:"TargetDBInstanceId,omitempty" xml:"TargetDBInstanceId,omitempty"`
	UsedTime            *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId               *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s RecoveryDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoveryDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RecoveryDBInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *RecoveryDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *RecoveryDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RecoveryDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *RecoveryDBInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *RecoveryDBInstanceRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *RecoveryDBInstanceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *RecoveryDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *RecoveryDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *RecoveryDBInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *RecoveryDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RecoveryDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *RecoveryDBInstanceRequest) GetTargetDBInstanceId() *string {
	return s.TargetDBInstanceId
}

func (s *RecoveryDBInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *RecoveryDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *RecoveryDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RecoveryDBInstanceRequest) SetBackupId(v string) *RecoveryDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDBInstanceClass(v string) *RecoveryDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDBInstanceId(v string) *RecoveryDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDBInstanceStorage(v int32) *RecoveryDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDBInstanceStorageType(v string) *RecoveryDBInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDbNames(v string) *RecoveryDBInstanceRequest {
	s.DbNames = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetInstanceNetworkType(v string) *RecoveryDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetPayType(v string) *RecoveryDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetPeriod(v string) *RecoveryDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetPrivateIpAddress(v string) *RecoveryDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetResourceOwnerId(v int64) *RecoveryDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetRestoreTime(v string) *RecoveryDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetTargetDBInstanceId(v string) *RecoveryDBInstanceRequest {
	s.TargetDBInstanceId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetUsedTime(v string) *RecoveryDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetVPCId(v string) *RecoveryDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetVSwitchId(v string) *RecoveryDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
