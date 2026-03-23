// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceMajorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowDDL(v bool) *UpgradeDBInstanceMajorVersionRequest
	GetAllowDDL() *bool
	SetCollectStatMode(v string) *UpgradeDBInstanceMajorVersionRequest
	GetCollectStatMode() *string
	SetCustomExtraInfo(v string) *UpgradeDBInstanceMajorVersionRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *UpgradeDBInstanceMajorVersionRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *UpgradeDBInstanceMajorVersionRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *UpgradeDBInstanceMajorVersionRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *UpgradeDBInstanceMajorVersionRequest
	GetDBInstanceStorageType() *string
	SetInstanceNetworkType(v string) *UpgradeDBInstanceMajorVersionRequest
	GetInstanceNetworkType() *string
	SetPayType(v string) *UpgradeDBInstanceMajorVersionRequest
	GetPayType() *string
	SetPeriod(v string) *UpgradeDBInstanceMajorVersionRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *UpgradeDBInstanceMajorVersionRequest
	GetPrivateIpAddress() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceMajorVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchOver(v string) *UpgradeDBInstanceMajorVersionRequest
	GetSwitchOver() *string
	SetSwitchTime(v string) *UpgradeDBInstanceMajorVersionRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *UpgradeDBInstanceMajorVersionRequest
	GetSwitchTimeMode() *string
	SetTargetMajorVersion(v string) *UpgradeDBInstanceMajorVersionRequest
	GetTargetMajorVersion() *string
	SetUpgradeMode(v string) *UpgradeDBInstanceMajorVersionRequest
	GetUpgradeMode() *string
	SetUsedTime(v string) *UpgradeDBInstanceMajorVersionRequest
	GetUsedTime() *string
	SetVPCId(v string) *UpgradeDBInstanceMajorVersionRequest
	GetVPCId() *string
	SetVSwitchId(v string) *UpgradeDBInstanceMajorVersionRequest
	GetVSwitchId() *string
	SetZoneId(v string) *UpgradeDBInstanceMajorVersionRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *UpgradeDBInstanceMajorVersionRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *UpgradeDBInstanceMajorVersionRequest
	GetZoneIdSlave2() *string
}

type UpgradeDBInstanceMajorVersionRequest struct {
	AllowDDL              *bool   `json:"AllowDDL,omitempty" xml:"AllowDDL,omitempty"`
	CollectStatMode       *string `json:"CollectStatMode,omitempty" xml:"CollectStatMode,omitempty"`
	CustomExtraInfo       *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStorage     *int32  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	InstanceNetworkType   *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// This parameter is required.
	PayType            *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PrivateIpAddress   *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerId    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchOver         *string `json:"SwitchOver,omitempty" xml:"SwitchOver,omitempty"`
	SwitchTime         *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode     *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	UpgradeMode        *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
	UsedTime           *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId              *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneIdSlave1       *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	ZoneIdSlave2       *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s UpgradeDBInstanceMajorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceMajorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetAllowDDL() *bool {
	return s.AllowDDL
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetCollectStatMode() *string {
	return s.CollectStatMode
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetPayType() *string {
	return s.PayType
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetPeriod() *string {
	return s.Period
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetSwitchOver() *string {
	return s.SwitchOver
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetTargetMajorVersion() *string {
	return s.TargetMajorVersion
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetAllowDDL(v bool) *UpgradeDBInstanceMajorVersionRequest {
	s.AllowDDL = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetCollectStatMode(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.CollectStatMode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetCustomExtraInfo(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetDBInstanceClass(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetDBInstanceStorage(v int32) *UpgradeDBInstanceMajorVersionRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetDBInstanceStorageType(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetInstanceNetworkType(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetPayType(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.PayType = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetPeriod(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.Period = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetPrivateIpAddress(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceMajorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetSwitchOver(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.SwitchOver = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetSwitchTime(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetTargetMajorVersion(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.TargetMajorVersion = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetUpgradeMode(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.UpgradeMode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetUsedTime(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.UsedTime = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetVPCId(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.VPCId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetVSwitchId(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetZoneId(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.ZoneId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetZoneIdSlave1(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetZoneIdSlave2(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) Validate() error {
	return dara.Validate(s)
}
