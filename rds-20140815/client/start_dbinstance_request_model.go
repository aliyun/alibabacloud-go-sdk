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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The data migration method of the instance. This parameter is available only for instances that are created in dedicated clusters. Valid values:
	//
	// 	- **0*	- (default): The system preferentially upgrades or downgrades the instance without a migration. If the resources on the host on which the instance resides are insufficient, the system migrates the instance to another suitable host.
	//
	// 	- **1**: The system upgrades or downgrades the instance without a migration. If the upgrade or downgrade is not supported, the system reports an error.
	//
	// 	- **2**: The system migrates the data of the instance from the host on which the instance resides to another host. You must also specify **DedicatedHostGroupId**, **TargetDedicatedHostIdForMaster**, and **TargetDedicatedHostIdForSlave**. If you set DBInstanceTransType to 2, you cannot migrate the data of the instance to the host on which the instance resides. If you migrate the data of the instance to the host on which the instance resides, the migration fails.
	//
	// example:
	//
	// 0
	DBInstanceTransType *int32 `json:"DBInstanceTransType,omitempty" xml:"DBInstanceTransType,omitempty"`
	// The dedicated cluster ID. This parameter is supported if you call this operation to suspend an RDS instance in the dedicated cluster. You can call the DescribeDedicatedHostGroups operation to query the dedicated cluster ID.
	//
	// example:
	//
	// dhg-39****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The effective time. This parameter is available only for instances that are created in dedicated clusters.
	//
	// 	- **Immediate**
	//
	// 	- **MaintainTime**: The change takes effect during the planned maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// 	- **SpecificTime**: The change takes effect at a specified point in time.
	//
	// Default value: MaintainTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The database engine version of the instance. This parameter is available only for instances that are created in dedicated clusters.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The switching time. This parameter is available only for instances that are created in dedicated clusters. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > This parameter must be specified when **EffectiveTime*	- is set to **Specified**.
	//
	// example:
	//
	// 2019-10-21T10:00:00Z
	SpecifiedTime *string `json:"SpecifiedTime,omitempty" xml:"SpecifiedTime,omitempty"`
	// The storage capacity of the instance. This parameter is available only for instances that are created in dedicated clusters. Valid values: **5 to 2000**. Unit: GB. If you do not specify this parameter, the storage capacity of the instance remains unchanged.
	//
	// example:
	//
	// 1000
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The instance type of the required instance. This parameter is available only for instances that are created in dedicated clusters.
	//
	// example:
	//
	// rds.ebmhfc6.20xlarge
	TargetDBInstanceClass *string `json:"TargetDBInstanceClass,omitempty" xml:"TargetDBInstanceClass,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// dh-bp****
	TargetDedicatedHostIdForLog *string `json:"TargetDedicatedHostIdForLog,omitempty" xml:"TargetDedicatedHostIdForLog,omitempty"`
	// The ID of the host on which the primary instance is created. This parameter is available only for instances that are created in dedicated clusters.
	//
	// > This parameter must be specified when **DBInstanceTransType*	- is set to **2**.
	//
	// example:
	//
	// dh-bp****
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// The ID of the host on which the secondary instance is created. This parameter is available only for instances that are created in dedicated clusters.
	//
	// > This parameter must be specified when **DBInstanceTransType*	- is set to **2**.
	//
	// example:
	//
	// dh-bp****
	TargetDedicatedHostIdForSlave *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	// The vSwitch ID. This parameter is available only for instances that are created in dedicated clusters.
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID. This parameter is available only for instances that are created in dedicated clusters.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
