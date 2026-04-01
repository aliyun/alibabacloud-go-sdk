// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *MigrateDBInstanceRequest
	GetDBInstanceId() *string
	SetDedicatedHostGroupId(v string) *MigrateDBInstanceRequest
	GetDedicatedHostGroupId() *string
	SetEffectiveTime(v string) *MigrateDBInstanceRequest
	GetEffectiveTime() *string
	SetOwnerId(v int64) *MigrateDBInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *MigrateDBInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *MigrateDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateDBInstanceRequest
	GetResourceOwnerId() *int64
	SetSpecifiedTime(v string) *MigrateDBInstanceRequest
	GetSpecifiedTime() *string
	SetTargetDedicatedHostIdForMaster(v string) *MigrateDBInstanceRequest
	GetTargetDedicatedHostIdForMaster() *string
	SetTargetDedicatedHostIdForSlave(v string) *MigrateDBInstanceRequest
	GetTargetDedicatedHostIdForSlave() *string
	SetZoneIdForFollower(v string) *MigrateDBInstanceRequest
	GetZoneIdForFollower() *string
	SetZoneIdForLog(v string) *MigrateDBInstanceRequest
	GetZoneIdForLog() *string
}

type MigrateDBInstanceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The dedicated cluster ID. You can call the DescribeDedicatedHostGroups operation to query the dedicated cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dhg-4n******
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The time when you want the system to start the migration. Valid values:
	//
	// 	- **Immediately**: The system immediately starts the migration. This is the default value.
	//
	// 	- **MaintainTime**: The system starts the migration during the specified maintenance window.
	//
	// 	- **Specified**: The system starts the migration at the specified point in time.
	//
	// example:
	//
	// MaintainTime
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time when you want the system to start the migration. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > This parameter must be specified when you set **EffectiveTime*	- to **Specified**.
	//
	// example:
	//
	// 2019-10-21T10:00:00Z
	SpecifiedTime *string `json:"SpecifiedTime,omitempty" xml:"SpecifiedTime,omitempty"`
	// The ID of the host to which you want to migrate the primary instance. You can call the DescribeDedicatedHosts operation to query the host ID.
	//
	// example:
	//
	// i-bp******
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// The ID of the host to which you want to migrate the secondary instance. You can call the DescribeDedicatedHosts operation to query the host ID.
	//
	// example:
	//
	// i-bp******
	TargetDedicatedHostIdForSlave *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	// The zone ID of the secondary node.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneIdForFollower *string `json:"ZoneIdForFollower,omitempty" xml:"ZoneIdForFollower,omitempty"`
	// The zone ID of the logger instance.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneIdForLog *string `json:"ZoneIdForLog,omitempty" xml:"ZoneIdForLog,omitempty"`
}

func (s MigrateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *MigrateDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateDBInstanceRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *MigrateDBInstanceRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *MigrateDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MigrateDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateDBInstanceRequest) GetSpecifiedTime() *string {
	return s.SpecifiedTime
}

func (s *MigrateDBInstanceRequest) GetTargetDedicatedHostIdForMaster() *string {
	return s.TargetDedicatedHostIdForMaster
}

func (s *MigrateDBInstanceRequest) GetTargetDedicatedHostIdForSlave() *string {
	return s.TargetDedicatedHostIdForSlave
}

func (s *MigrateDBInstanceRequest) GetZoneIdForFollower() *string {
	return s.ZoneIdForFollower
}

func (s *MigrateDBInstanceRequest) GetZoneIdForLog() *string {
	return s.ZoneIdForLog
}

func (s *MigrateDBInstanceRequest) SetDBInstanceId(v string) *MigrateDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetDedicatedHostGroupId(v string) *MigrateDBInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetEffectiveTime(v string) *MigrateDBInstanceRequest {
	s.EffectiveTime = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetOwnerId(v int64) *MigrateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetRegionId(v string) *MigrateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetResourceOwnerAccount(v string) *MigrateDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetResourceOwnerId(v int64) *MigrateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetSpecifiedTime(v string) *MigrateDBInstanceRequest {
	s.SpecifiedTime = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetTargetDedicatedHostIdForMaster(v string) *MigrateDBInstanceRequest {
	s.TargetDedicatedHostIdForMaster = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetTargetDedicatedHostIdForSlave(v string) *MigrateDBInstanceRequest {
	s.TargetDedicatedHostIdForSlave = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetZoneIdForFollower(v string) *MigrateDBInstanceRequest {
	s.ZoneIdForFollower = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetZoneIdForLog(v string) *MigrateDBInstanceRequest {
	s.ZoneIdForLog = &v
	return s
}

func (s *MigrateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
