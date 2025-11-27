// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowMajorVersionUpgrade(v bool) *ModifyDBInstanceSpecShrinkRequest
	GetAllowMajorVersionUpgrade() *bool
	SetAutoUseCoupon(v bool) *ModifyDBInstanceSpecShrinkRequest
	GetAutoUseCoupon() *bool
	SetBurstingEnabled(v bool) *ModifyDBInstanceSpecShrinkRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *ModifyDBInstanceSpecShrinkRequest
	GetCategory() *string
	SetColdDataEnabled(v bool) *ModifyDBInstanceSpecShrinkRequest
	GetColdDataEnabled() *bool
	SetCompressionMode(v string) *ModifyDBInstanceSpecShrinkRequest
	GetCompressionMode() *string
	SetDBInstanceClass(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *ModifyDBInstanceSpecShrinkRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDBInstanceStorageType() *string
	SetDedicatedHostGroupId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDedicatedHostGroupId() *string
	SetDirection(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDirection() *string
	SetEffectiveTime(v string) *ModifyDBInstanceSpecShrinkRequest
	GetEffectiveTime() *string
	SetEngineVersion(v string) *ModifyDBInstanceSpecShrinkRequest
	GetEngineVersion() *string
	SetIoAccelerationEnabled(v string) *ModifyDBInstanceSpecShrinkRequest
	GetIoAccelerationEnabled() *string
	SetOptimizedWrites(v string) *ModifyDBInstanceSpecShrinkRequest
	GetOptimizedWrites() *string
	SetOwnerAccount(v string) *ModifyDBInstanceSpecShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceSpecShrinkRequest
	GetOwnerId() *int64
	SetPayType(v string) *ModifyDBInstanceSpecShrinkRequest
	GetPayType() *string
	SetPromotionCode(v string) *ModifyDBInstanceSpecShrinkRequest
	GetPromotionCode() *string
	SetReadOnlyDBInstanceClass(v string) *ModifyDBInstanceSpecShrinkRequest
	GetReadOnlyDBInstanceClass() *string
	SetResourceGroupId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceSpecShrinkRequest
	GetResourceOwnerId() *int64
	SetServerlessConfigurationShrink(v string) *ModifyDBInstanceSpecShrinkRequest
	GetServerlessConfigurationShrink() *string
	SetSourceBiz(v string) *ModifyDBInstanceSpecShrinkRequest
	GetSourceBiz() *string
	SetSwitchTime(v string) *ModifyDBInstanceSpecShrinkRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *ModifyDBInstanceSpecShrinkRequest
	GetTargetMinorVersion() *string
	SetUsedTime(v int64) *ModifyDBInstanceSpecShrinkRequest
	GetUsedTime() *int64
	SetVSwitchId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *ModifyDBInstanceSpecShrinkRequest
	GetZoneIdSlave1() *string
}

type ModifyDBInstanceSpecShrinkRequest struct {
	// Specifies whether to upgrade the major engine version of an ApsaraDB RDS for SQL Server instance. For more information, see [Upgrade the major engine version](https://help.aliyun.com/document_detail/127458.html). Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > 	- When you upgrade the major engine version, you must also specify the required parameters such as DBInstanceId, EngineVersion, DBInstanceClass, Category, ZoneId, and VSwitchId.
	//
	// > 	- If you want to upgrade the instance edition to RDS High-availability Edition or RDS Cluster Edition, you must specify ZoneIdSlave1.
	//
	// example:
	//
	// false
	AllowMajorVersionUpgrade *bool `json:"AllowMajorVersionUpgrade,omitempty" xml:"AllowMajorVersionUpgrade,omitempty"`
	// Specifies whether to use vouchers to offset fees. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// An invalid parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The RDS edition of the instance. Valid values:
	//
	// >  If you set **EngineVersion*	- to an SQL Server version number, you must also specify this parameter.
	//
	// **Regular RDS instances**
	//
	// 	- **Basic**: RDS Basic Edition.
	//
	// 	- **HighAvailability**: RDS High-availability Edition.
	//
	// 	- **AlwaysOn**: RDS Cluster Edition for ApsaraDB RDS for SQL Server.
	//
	// 	- **Cluster**: RDS Cluster Edition for ApsaraDB RDS for MySQL.
	//
	// **Serverless instances. ApsaraDB RDS for MariaDB does not support serverless instances.**
	//
	// 	- **serverless_basic**: RDS Basic Edition. This edition is available only for serverless instances that run MySQL and PostgreSQL.
	//
	// 	- **serverless_standard**: RDS High-availability Edition. This edition is available only for serverless instances that run MySQL and PostgreSQL.
	//
	// 	- **serverless_ha**: RDS High-availability Edition for serverless instances. This edition is available only for instances that run SQL Server.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// true
	ColdDataEnabled *bool `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	// Specifies whether to enable the storage compression feature for the ApsaraDB RDS for MySQL instance. For more information, see [Use the storage compression feature](https://help.aliyun.com/document_detail/2861985.html). Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	CompressionMode *string `json:"CompressionMode,omitempty" xml:"CompressionMode,omitempty"`
	// The instance type of the new instance. For more information, see [Specifications](https://help.aliyun.com/document_detail/26312.html). You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) operation to query the instance types.
	//
	// > 	- You must specify at least one of DBInstanceClass and **DBInstanceStorage**.
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html) operation to query the current instance type of the instance.
	//
	// example:
	//
	// rds.mys2.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) operation to query the instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The storage capacity of the new instance. Unit: GB. For more information, see [Storage types](https://help.aliyun.com/document_detail/26312.html). You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) operation to query the storage capacity range that is supported by the new instance type.
	//
	// > 	- You must specify at least one of DBInstanceStorage and **DBInstanceClass**.
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html) operation to query the current storage capacity of the instance.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the new instance. Valid values:
	//
	// 	- **local_ssd**: local SSD.
	//
	// 	- **cloud_ssd**: SSD cloud disks. This storage medium is not recommended and is unavailable in specific Alibaba Cloud regions.
	//
	// 	- **cloud_essd**: performance level 1 (PL1) Enterprise SSD (ESSD).
	//
	// 	- **cloud_essd2**: PL2 ESSD.
	//
	// 	- **cloud_essd3**: PL3 ESSD.
	//
	// To change the storage type, take note of the following items:
	//
	// If the instance runs PostgreSQL, you can upgrade the storage type of the instance from standard SSDs to ESSDs. However, you cannot downgrade the storage type of the instance from ESSDs to standard SSDs. ESSDs provide the following PLs: ESSDs of PL1, ESSDs of PL2, and ESSDs of PL3. You can upgrade or downgrade the storage type between ESSD of PL1, ESSD of PL2, and ESSD of PL3. For more information, see [Configuration items](https://help.aliyun.com/document_detail/96750.html).
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The ID of the dedicated cluster.
	//
	// example:
	//
	// dhg-7a9********
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The type of change that you want to perform on the instance. Valid values:
	//
	// 	- **Up*	- (default): upgrades a subscription instance, or upgrades or downgrades a pay-as-you-go instance.
	//
	// 	- **Down**: downgrades a subscription instance.
	//
	// 	- **TempUpgrade**: performs auto scaling on a subscription instance that runs SQL Server. This value is required for auto scaling.
	//
	// 	- **Serverless**: modifies the auto scaling settings of a serverless instance.
	//
	// >  If you specify only **DBInstanceStorageType**, you can leave Direction empty. For example, if you want to change only the storage type of the instance from standard SSD to Enterprise SSD (ESSD), you do not need to specify Direction.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The time when the new specifications take effect. Valid values:
	//
	// >  **Specific changes may affect the instance**. Read the [Impact](https://help.aliyun.com/document_detail/96061.html) section before you specify this parameter. We recommend that you specify this parameter during off-peak hours.
	//
	// 	- **Immediate*	- (default): The changes immediately take effect.
	//
	// 	- **MaintainTime**: The changes take effect during the [maintenance window](https://help.aliyun.com/document_detail/610402.html) of the instance.
	//
	// 	- **ScheduleTime**: The changes take effect at the point in time that you specify. This time must be at least 12 hours later than the current time. The actual effective time is calculated based on the following formula: EffectiveTime = ScheduleTime + SwitchTime.
	//
	// example:
	//
	// MaintainTime
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The database engine version of the instance. Valid values:
	//
	// **Regular RDS instances**
	//
	// 	- Valid values when Engine is set to MySQL: 5.5, 5.6, 5.7, and 8.0.
	//
	// 	- Valid values when Engine is set to SQLServer: 2008r2, 08r2_ent_ha, 2012, 2012_ent_ha, 2012_std_ha, 2012_web, 2014_std_ha, 2016_ent_ha, 2016_std_ha, 2016_web, 2017_std_ha, 2017_ent, 2019_std_ha, and 2019_ent.
	//
	// 	- Valid values when Engine is set to PostgreSQL: 10.0, 11.0, 12.0, 13.0, 14.0, and 15.0.
	//
	// 	- Valid value when Engine is set to MariaDB: 10.3.
	//
	// **Serverless instances. ApsaraDB RDS for MariaDB does not support serverless instances.**
	//
	// 	- Valid values when Engine is set to MySQL: 5.7 and 8.0.
	//
	// 	- Valid values when Engine is set to SQL Server: 2016_std_sl, 2017_std_sl, and 2019_std_sl.
	//
	// 	- Valid values when Engine is set to PostgreSQL: 14.0, 15.0, and 16.0.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to enable the write optimization feature for the ApsaraDB RDS for MySQL instance. For more information, see [Use the write optimization feature](https://help.aliyun.com/document_detail/2858761.html). Valid values:
	//
	// 	- **optimized**: enables the feature.
	//
	// 	- **none**: disables the feature.
	//
	// example:
	//
	// optimized
	OptimizedWrites *string `json:"OptimizedWrites,omitempty" xml:"OptimizedWrites,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// 	- **Serverless**: serverless. This value is not supported for ApsaraDB RDS for MariaDB instances.
	//
	// >  If you want to set this parameter to Serverless, **you must specify **AutoPause, MaxCapacity, MinCapacity, and SwitchForce. For more information, see [Overview of serverless ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/411291.html), [Overview of serverless ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of serverless ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/607742.html).
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 723298850895
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The specification of the read-only instance when you change the storage type of the ApsaraDB RDS for MySQL instance that runs RDS High-availability Edition from cloud disk to local disk.
	//
	// example:
	//
	// mysqlro.n2.large.c
	ReadOnlyDBInstanceClass *string `json:"ReadOnlyDBInstanceClass,omitempty" xml:"ReadOnlyDBInstanceClass,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy**********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The specifications that you want to change for a serverless instance.
	ServerlessConfigurationShrink *string `json:"ServerlessConfiguration,omitempty" xml:"ServerlessConfiguration,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// Specifies whether to enable the automatic suspension feature.
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
	// The time at which you want to change the specifications. **We recommend that you perform the specification changes during off-peak hours.**
	//
	// Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > 	- The time at which you want to change the specifications **must be later than the current time**. Otherwise, the specification change task fails. If the specification change task fails, you must wait for the order to be automatically canceled, and then call this operation again.
	//
	// > 	- If you want to increase the storage capacity or change the ESSD storage type between different PLs, the specification change immediately takes effect and does not affect your workloads. You do not need to specify this parameter.
	//
	// example:
	//
	// 2019-07-10T13:15:12Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The minor engine version number of the ApsaraDB RDS for PostgreSQL instance. For more information, see [Update the minor engine version](https://help.aliyun.com/document_detail/126002.html). If the minor engine version does not support changing the instance type, you must specify the minor engine version to **update the minor engine version when you change the instance type**.
	//
	// Format: `rds_postgres_<Major engine version>00_<Minor engine version>`. For example, if the instance runs PostgreSQL 12, set this parameter to `rds_postgres_1200_20200830`.
	//
	// example:
	//
	// rds_postgres_1200_20200830
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The validity period of the specification changes on an ApsaraDB RDS for SQL Server instance. At the end of the validity period, the specifications of the instance are restored to the specifications that are used before an [elastic upgrade](https://help.aliyun.com/document_detail/95665.html) is performed. Unit: days.
	//
	// example:
	//
	// 3
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch ID. The vSwitch must belong to the zone that is specified by **ZoneId**.
	//
	// 	- If you set **InstanceNetworkType*	- to **VPC**, you must also specify this parameter.
	//
	// 	- If you specify ZoneSlaveId1, you must specify the IDs of two vSwitches for this parameter and separate the IDs with a comma (,).
	//
	// >  If you want to upgrade the major engine version of an ApsaraDB RDS for SQL Server instance by specifying AllowMajorVersionUpgrade or change the vSwitch, you must specify this parameter.
	//
	// example:
	//
	// vsw-bp1oxflciovg9l7163lr7
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition.
	//
	// 	- **HighAvailability**: RDS High-availability Edition.
	//
	// 	- **AlwaysOn**: RDS Cluster Edition for SQL Server.
	//
	// 	- **Finance**: RDS Enterprise Edition. This edition is available only on the China site (aliyun.com).
	//
	// > If you set **EngineVersion*	- to an SQL Server version number, you must also specify this parameter.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary instance. If you set this parameter to the same value as **ZoneId**, the single-zone deployment method is used. If you set this parameter to a different value from **ZoneId**, the multi-zone deployment method is used.
	//
	// >  If you want to upgrade the major engine version of an ApsaraDB RDS for SQL Server instance by specifying AllowMajorVersionUpgrade or change the secondary zone, you must specify this parameter.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
}

func (s ModifyDBInstanceSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetAllowMajorVersionUpgrade() *bool {
	return s.AllowMajorVersionUpgrade
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetCompressionMode() *string {
	return s.CompressionMode
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetOptimizedWrites() *string {
	return s.OptimizedWrites
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetReadOnlyDBInstanceClass() *string {
	return s.ReadOnlyDBInstanceClass
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetServerlessConfigurationShrink() *string {
	return s.ServerlessConfigurationShrink
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetAllowMajorVersionUpgrade(v bool) *ModifyDBInstanceSpecShrinkRequest {
	s.AllowMajorVersionUpgrade = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetAutoUseCoupon(v bool) *ModifyDBInstanceSpecShrinkRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetBurstingEnabled(v bool) *ModifyDBInstanceSpecShrinkRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetCategory(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetColdDataEnabled(v bool) *ModifyDBInstanceSpecShrinkRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetCompressionMode(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.CompressionMode = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDBInstanceClass(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDBInstanceId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDBInstanceStorage(v int32) *ModifyDBInstanceSpecShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDBInstanceStorageType(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDedicatedHostGroupId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDirection(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.Direction = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetEffectiveTime(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetEngineVersion(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetIoAccelerationEnabled(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetOptimizedWrites(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.OptimizedWrites = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetOwnerAccount(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetOwnerId(v int64) *ModifyDBInstanceSpecShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetPayType(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.PayType = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetPromotionCode(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetReadOnlyDBInstanceClass(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ReadOnlyDBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetResourceGroupId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSpecShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetServerlessConfigurationShrink(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ServerlessConfigurationShrink = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetSourceBiz(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.SourceBiz = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetSwitchTime(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetTargetMinorVersion(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetUsedTime(v int64) *ModifyDBInstanceSpecShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetVSwitchId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetZoneId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetZoneIdSlave1(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
