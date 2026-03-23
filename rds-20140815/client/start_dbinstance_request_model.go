// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *StartDBInstanceRequest
	GetDBInstanceId() *string
	SetDBInstanceTransType(v int32) *StartDBInstanceRequest
	GetDBInstanceTransType() *int32
	SetDedicatedHostGroupId(v string) *StartDBInstanceRequest
	GetDedicatedHostGroupId() *string
	SetEffectiveTime(v string) *StartDBInstanceRequest
	GetEffectiveTime() *string
	SetEngineVersion(v string) *StartDBInstanceRequest
	GetEngineVersion() *string
	SetOwnerId(v int64) *StartDBInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartDBInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartDBInstanceRequest
	GetResourceOwnerId() *int64
	SetSpecifiedTime(v string) *StartDBInstanceRequest
	GetSpecifiedTime() *string
	SetStorage(v int32) *StartDBInstanceRequest
	GetStorage() *int32
	SetTargetDBInstanceClass(v string) *StartDBInstanceRequest
	GetTargetDBInstanceClass() *string
	SetTargetDedicatedHostIdForLog(v string) *StartDBInstanceRequest
	GetTargetDedicatedHostIdForLog() *string
	SetTargetDedicatedHostIdForMaster(v string) *StartDBInstanceRequest
	GetTargetDedicatedHostIdForMaster() *string
	SetTargetDedicatedHostIdForSlave(v string) *StartDBInstanceRequest
	GetTargetDedicatedHostIdForSlave() *string
	SetVSwitchId(v string) *StartDBInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *StartDBInstanceRequest
	GetZoneId() *string
}

type StartDBInstanceRequest struct {
	// This parameter is required.
	DBInstanceId                   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceTransType            *int32  `json:"DBInstanceTransType,omitempty" xml:"DBInstanceTransType,omitempty"`
	DedicatedHostGroupId           *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	EffectiveTime                  *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	EngineVersion                  *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId                        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount           *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpecifiedTime                  *string `json:"SpecifiedTime,omitempty" xml:"SpecifiedTime,omitempty"`
	Storage                        *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
	TargetDBInstanceClass          *string `json:"TargetDBInstanceClass,omitempty" xml:"TargetDBInstanceClass,omitempty"`
	TargetDedicatedHostIdForLog    *string `json:"TargetDedicatedHostIdForLog,omitempty" xml:"TargetDedicatedHostIdForLog,omitempty"`
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	TargetDedicatedHostIdForSlave  *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	VSwitchId                      *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s StartDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *StartDBInstanceRequest) GetDBInstanceTransType() *int32 {
	return s.DBInstanceTransType
}

func (s *StartDBInstanceRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *StartDBInstanceRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *StartDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *StartDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartDBInstanceRequest) GetSpecifiedTime() *string {
	return s.SpecifiedTime
}

func (s *StartDBInstanceRequest) GetStorage() *int32 {
	return s.Storage
}

func (s *StartDBInstanceRequest) GetTargetDBInstanceClass() *string {
	return s.TargetDBInstanceClass
}

func (s *StartDBInstanceRequest) GetTargetDedicatedHostIdForLog() *string {
	return s.TargetDedicatedHostIdForLog
}

func (s *StartDBInstanceRequest) GetTargetDedicatedHostIdForMaster() *string {
	return s.TargetDedicatedHostIdForMaster
}

func (s *StartDBInstanceRequest) GetTargetDedicatedHostIdForSlave() *string {
	return s.TargetDedicatedHostIdForSlave
}

func (s *StartDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *StartDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *StartDBInstanceRequest) SetDBInstanceId(v string) *StartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StartDBInstanceRequest) SetDBInstanceTransType(v int32) *StartDBInstanceRequest {
	s.DBInstanceTransType = &v
	return s
}

func (s *StartDBInstanceRequest) SetDedicatedHostGroupId(v string) *StartDBInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *StartDBInstanceRequest) SetEffectiveTime(v string) *StartDBInstanceRequest {
	s.EffectiveTime = &v
	return s
}

func (s *StartDBInstanceRequest) SetEngineVersion(v string) *StartDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *StartDBInstanceRequest) SetOwnerId(v int64) *StartDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StartDBInstanceRequest) SetRegionId(v string) *StartDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartDBInstanceRequest) SetResourceOwnerAccount(v string) *StartDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartDBInstanceRequest) SetResourceOwnerId(v int64) *StartDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartDBInstanceRequest) SetSpecifiedTime(v string) *StartDBInstanceRequest {
	s.SpecifiedTime = &v
	return s
}

func (s *StartDBInstanceRequest) SetStorage(v int32) *StartDBInstanceRequest {
	s.Storage = &v
	return s
}

func (s *StartDBInstanceRequest) SetTargetDBInstanceClass(v string) *StartDBInstanceRequest {
	s.TargetDBInstanceClass = &v
	return s
}

func (s *StartDBInstanceRequest) SetTargetDedicatedHostIdForLog(v string) *StartDBInstanceRequest {
	s.TargetDedicatedHostIdForLog = &v
	return s
}

func (s *StartDBInstanceRequest) SetTargetDedicatedHostIdForMaster(v string) *StartDBInstanceRequest {
	s.TargetDedicatedHostIdForMaster = &v
	return s
}

func (s *StartDBInstanceRequest) SetTargetDedicatedHostIdForSlave(v string) *StartDBInstanceRequest {
	s.TargetDedicatedHostIdForSlave = &v
	return s
}

func (s *StartDBInstanceRequest) SetVSwitchId(v string) *StartDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *StartDBInstanceRequest) SetZoneId(v string) *StartDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *StartDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
