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
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is supported only for dedicated cluster instances. The migration method of the instance. Valid values:
	//
	// 	- **0**: Default value. The system preferentially performs a local specification change. If local resources are insufficient, a cross-instance migration is performed.
	//
	// 	- **1**: Local specification change. If the system determines that the instance does not support a local specification change, an error is returned.
	//
	// 	- **2**: Cross-instance migration. The instance is migrated to a specified host. You must specify **DedicatedHostGroupId**, **TargetDedicatedHostIdForMaster**, and **TargetDedicatedHostIdForSlave**. The instance cannot be migrated to the host on which it currently resides. Otherwise, the migration fails.
	//
	// example:
	//
	// 0
	DBInstanceTransType *int32 `json:"DBInstanceTransType,omitempty" xml:"DBInstanceTransType,omitempty"`
	// This operation also supports starting an ApsaraDB RDS instance in a dedicated cluster. In this case, specify the dedicated cluster ID. You can call DescribeDedicatedHostGroups to query the dedicated cluster ID.
	//
	// example:
	//
	// dhg-39****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// This parameter is supported only for dedicated cluster instances. The effective period. Valid values:
	//
	// 	- **Immediate**: The operation takes effect immediately.
	//
	// 	- **MaintainTime**: The operation takes effect during the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// 	- **SpecificTime**: The operation takes effect at a specified time.
	//
	// Default value: MaintainTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// This parameter is supported only for dedicated cluster instances. The database engine version.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call DescribeRegions to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is supported only for dedicated cluster instances. The specified switchover time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// > This parameter is required when **EffectiveTime*	- is set to **Specified**.
	//
	// example:
	//
	// 2019-10-21T10:00:00Z
	SpecifiedTime *string `json:"SpecifiedTime,omitempty" xml:"SpecifiedTime,omitempty"`
	// This parameter is supported only for dedicated cluster instances. The custom storage capacity. Valid values: **5 to 2000**. Unit: GB. If you do not specify this parameter, the storage capacity remains unchanged.
	//
	// example:
	//
	// 1000
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// This parameter is supported only for dedicated cluster instances. The instance type of the target instance.
	//
	// example:
	//
	// rds.ebmhfc6.20xlarge
	TargetDBInstanceClass *string `json:"TargetDBInstanceClass,omitempty" xml:"TargetDBInstanceClass,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and does not need to be configured.
	//
	// example:
	//
	// dh-bp****
	TargetDedicatedHostIdForLog *string `json:"TargetDedicatedHostIdForLog,omitempty" xml:"TargetDedicatedHostIdForLog,omitempty"`
	// This parameter is supported only for dedicated cluster instances. Specifies the ID of the destination host for the primary node.
	//
	// > This parameter is required when **DBInstanceTransType*	- is set to **2**.
	//
	// example:
	//
	// dh-bp****
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// This parameter is supported only for dedicated cluster instances. Specifies the ID of the destination host for the secondary node.
	//
	// > This parameter is required when **DBInstanceTransType*	- is set to **2**.
	//
	// example:
	//
	// dh-bp****
	TargetDedicatedHostIdForSlave *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	// This parameter is supported only for dedicated cluster instances. The vSwitch ID.
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is supported only for dedicated cluster instances. The zone ID.
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
