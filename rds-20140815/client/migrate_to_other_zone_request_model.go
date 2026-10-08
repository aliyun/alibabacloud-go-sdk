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
	// The instance edition. Valid values:
	//
	// 	- **Basic**: Basic Edition
	//
	// 	- **HighAvailability**: High-availability Edition
	//
	// 	- **AlwaysOn**: SQL Server Cluster Edition
	//
	// 	- **cluster**: MySQL Cluster Edition
	//
	// 	- **Finance**: RDS Enterprise Edition
	//
	// example:
	//
	// HighAvailability
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The target instance type of the destination instance. Only the instance type can be changed. The storage type cannot be changed.
	//
	// When the **IsModifySpec*	- parameter settings require **true**, you must specify at least one of this parameter and **DBInstanceStorage**.
	//
	// For more information about instance types, see [Primary ApsaraDB RDS for MySQL instance types](https://help.aliyun.com/document_detail/276975.html).
	//
	// example:
	//
	// mysql.x4.xlarge.2
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The destination storage capacity. When the **IsModifySpec*	- parameter settings require **true**, you must specify at least one of this parameter and **DBInstanceClass**.
	//
	// Unit: GB.
	//
	// Valid values: The storage capacity varies based on the instance type. For more information, see [Primary ApsaraDB RDS for MySQL instance types](https://help.aliyun.com/document_detail/276975.html).
	//
	// example:
	//
	// 500
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// - cloud_essd: PL1 ESSD cloud disk.
	//
	// - cloud_essd2: PL2 ESSD cloud disk.
	//
	// - cloud_essd3: PL3 ESSD cloud disk.
	//
	// - cloud_ssd: standard SSD (not recommended because standard SSDs are no longer available for purchase in some regions).
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The effective period. Valid values:
	//
	// 	- **Immediate**: The migration takes effect immediately. This is the default value.
	//
	// 	- **MaintainTime**: The migration takes effect during the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// 	- **ScheduleTime**: The migration takes effect at a custom time.
	//
	// > If you set this parameter to **ScheduleTime**, you must also specify the **SwitchTime*	- parameter.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// Specifies whether to enable the Buffer Pool Extension (BPE) feature for premium performance disks. Valid values:
	//
	//  - **1**: Enable.
	//
	//  - **0**: Disable.
	//
	// > For more information about the BPE feature, see [Buffer Pool Extension (BPE)](https://help.aliyun.com/document_detail/2527067.html).
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to change the instance specifications during zone migration.
	//
	// - **true**: Change the specifications. When this parameter is set to **true**, you must specify at least one of the **DBInstanceClass*	- and **DBInstanceStorage*	- parameters.
	//
	// - **false**: Do not change the specifications. This is the default value.
	//
	// > This parameter is applicable only to ApsaraDB RDS for MySQL instances.
	//
	// example:
	//
	// true
	IsModifySpec         *string `json:"IsModifySpec,omitempty" xml:"IsModifySpec,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The custom time at which the zone switch takes effect. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// > This parameter is used together with the **EffectiveTime*	- parameter and is required only when **EffectiveTime*	- is set to **ScheduleTime**.
	//
	// example:
	//
	// 2021-12-14T15:15:15Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The virtual private cloud (VPC) ID. The VPC cannot be changed during instance migration and must remain the same.
	//
	// - This parameter is required when you migrate a VPC-connected instance to a different zone.
	//
	// - If the instance engine is SQL Server, the VPC can be changed during instance migration.
	//
	// example:
	//
	// vpc-****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// - This parameter is required when you migrate a VPC-connected instance to a different zone. You can invoke DescribeVSwitches to query the vSwitches that have been created.
	//
	// - When you perform instance migration for an ApsaraDB RDS for PostgreSQL or SQL Server instance to a different zone with a secondary zone configured, you can specify multiple vSwitch IDs separated by commas (,), corresponding to the zones.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the destination zone. You can call DescribeRegions to query the zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The secondary zone 1.
	//
	// > This parameter is required for instances that are not of the Basic Edition.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The secondary zone 2.
	//
	// > This parameter is applicable only to RDS Enterprise Edition instances.
	//
	// example:
	//
	// cn-hangzhou-d
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
