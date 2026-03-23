// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateToOtherZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *MigrateToOtherZoneRequest
	GetCategory() *string
	SetCustomExtraInfo(v string) *MigrateToOtherZoneRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *MigrateToOtherZoneRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *MigrateToOtherZoneRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int64) *MigrateToOtherZoneRequest
	GetDBInstanceStorage() *int64
	SetDBInstanceStorageType(v string) *MigrateToOtherZoneRequest
	GetDBInstanceStorageType() *string
	SetEffectiveTime(v string) *MigrateToOtherZoneRequest
	GetEffectiveTime() *string
	SetIoAccelerationEnabled(v string) *MigrateToOtherZoneRequest
	GetIoAccelerationEnabled() *string
	SetIsModifySpec(v string) *MigrateToOtherZoneRequest
	GetIsModifySpec() *string
	SetOwnerAccount(v string) *MigrateToOtherZoneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MigrateToOtherZoneRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MigrateToOtherZoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateToOtherZoneRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *MigrateToOtherZoneRequest
	GetSwitchTime() *string
	SetVPCId(v string) *MigrateToOtherZoneRequest
	GetVPCId() *string
	SetVSwitchId(v string) *MigrateToOtherZoneRequest
	GetVSwitchId() *string
	SetZoneId(v string) *MigrateToOtherZoneRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *MigrateToOtherZoneRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *MigrateToOtherZoneRequest
	GetZoneIdSlave2() *string
}

type MigrateToOtherZoneRequest struct {
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// This parameter is required.
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStorage     *int64  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	EffectiveTime         *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	IsModifySpec          *string `json:"IsModifySpec,omitempty" xml:"IsModifySpec,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchTime            *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	VPCId                 *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	ZoneIdSlave2 *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s MigrateToOtherZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateToOtherZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneRequest) GetCategory() *string {
	return s.Category
}

func (s *MigrateToOtherZoneRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceStorage() *int64 {
	return s.DBInstanceStorage
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *MigrateToOtherZoneRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *MigrateToOtherZoneRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *MigrateToOtherZoneRequest) GetIsModifySpec() *string {
	return s.IsModifySpec
}

func (s *MigrateToOtherZoneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MigrateToOtherZoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateToOtherZoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateToOtherZoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateToOtherZoneRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *MigrateToOtherZoneRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *MigrateToOtherZoneRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *MigrateToOtherZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *MigrateToOtherZoneRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *MigrateToOtherZoneRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *MigrateToOtherZoneRequest) SetCategory(v string) *MigrateToOtherZoneRequest {
	s.Category = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetCustomExtraInfo(v string) *MigrateToOtherZoneRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceClass(v string) *MigrateToOtherZoneRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceId(v string) *MigrateToOtherZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceStorage(v int64) *MigrateToOtherZoneRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceStorageType(v string) *MigrateToOtherZoneRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetEffectiveTime(v string) *MigrateToOtherZoneRequest {
	s.EffectiveTime = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetIoAccelerationEnabled(v string) *MigrateToOtherZoneRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetIsModifySpec(v string) *MigrateToOtherZoneRequest {
	s.IsModifySpec = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerId(v int64) *MigrateToOtherZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetResourceOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetResourceOwnerId(v int64) *MigrateToOtherZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetSwitchTime(v string) *MigrateToOtherZoneRequest {
	s.SwitchTime = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetVPCId(v string) *MigrateToOtherZoneRequest {
	s.VPCId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetVSwitchId(v string) *MigrateToOtherZoneRequest {
	s.VSwitchId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneId(v string) *MigrateToOtherZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneIdSlave1(v string) *MigrateToOtherZoneRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneIdSlave2(v string) *MigrateToOtherZoneRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *MigrateToOtherZoneRequest) Validate() error {
	return dara.Validate(s)
}
