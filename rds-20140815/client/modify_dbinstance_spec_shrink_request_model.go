// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocateStrategy(v string) *ModifyDBInstanceSpecShrinkRequest
	GetAllocateStrategy() *string
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
	// The resource allocation strategy. Valid values: ForceLocal and Optional. You cannot set this parameter to ForceLocal when you downgrade the instance.
	//
	// example:
	//
	// Optional
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// Specifies whether to perform a [major engine version upgrade](https://help.aliyun.com/document_detail/127458.html) for the ApsaraDB RDS for SQL Server instance. Valid values:
	//
	// - **true**: performs the upgrade.
	//
	// - **false*	- (default): does not perform the upgrade.
	//
	// > - When you perform a major engine version upgrade, you must also pass the other required parameters, including DBInstanceId, EngineVersion, DBInstanceClass, Category, ZoneId, and VSwitchId.
	//
	// > - If you want to upgrade the instance to the high availability series or Cluster Edition, you must also specify the ZoneIdSlave1 parameter.
	//
	// example:
	//
	// false
	AllowMajorVersionUpgrade *bool `json:"AllowMajorVersionUpgrade,omitempty" xml:"AllowMajorVersionUpgrade,omitempty"`
	// Specifies whether to use vouchers to offset fees. Valid values:
	//
	// - **true*	- (default): uses vouchers.
	//
	// - **false**: does not use vouchers.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// Specifies whether to enable the [I/O performance burst feature for premium performance disks](https://help.aliyun.com/document_detail/2340501.html). Valid values:
	//
	// - **true**: enables the feature.
	//
	// - **false**: disables the feature.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The [instance edition](https://help.aliyun.com/document_detail/53509.html). Valid values:
	//
	// > This parameter is required when **EngineVersion*	- is set to a SQL Server version number.
	//
	// <details>
	//
	// <summary>Regular ApsaraDB RDS instances</summary>
	//
	// - **Basic**: Basic Edition
	//
	// - **HighAvailability**: high availability series
	//
	// - **AlwaysOn**: Cluster Edition (for SQL Server)
	//
	// - **Cluster**: Cluster Edition (for MySQL)
	//
	// - <props="china">**Finance**: Finance Edition
	//
	// </details>
	//
	// <details>
	//
	// <summary>Serverless ApsaraDB RDS instances (not supported for MariaDB)</summary>
	//
	// - **serverless_basic**: Serverless Basic Edition (only for MySQL and PostgreSQL)
	//
	// - **serverless_standard**: Serverless high availability series (only for MySQL and PostgreSQL)
	//
	// - **serverless_ha**: Serverless high availability series (only for SQL Server)
	//
	// </details>
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to enable the [data archiving feature](https://help.aliyun.com/document_detail/2701832.html) for premium performance disks. Valid values:
	//
	// - **true**: enables the feature.
	//
	// - **false**: disables the feature.
	//
	// example:
	//
	// true
	ColdDataEnabled *bool `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	// Specifies whether to enable the [storage compression feature](https://help.aliyun.com/document_detail/2861985.html) for ApsaraDB RDS for MySQL. Valid values:
	//
	// - **on**: enables the feature.
	//
	// - **off**: disables the feature.
	//
	// example:
	//
	// on
	CompressionMode *string `json:"CompressionMode,omitempty" xml:"CompressionMode,omitempty"`
	// The [target instance type](https://help.aliyun.com/document_detail/26312.html). You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) operation to query the available instance types.
	//
	// > 	- You must specify at least one of the **DBInstanceClass*	- and **DBInstanceStorage*	- parameters.
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html) operation to view the current instance type.
	//
	// example:
	//
	// mysql.n8.large.2c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// Instance ID of the instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The [target storage capacity](https://help.aliyun.com/document_detail/26312.html). Unit: GB. You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) operation to query the available storage capacity range for the target instance type.
	//
	// > 	- You must specify at least one of the **DBInstanceStorage*	- and **DBInstanceClass*	- parameters.
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html) operation to view the current storage capacity of the instance.
	//
	// example:
	//
	// 100
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: premium local SSD
	//
	// 	- **cloud_ssd**: standard SSD (not recommended; no longer available in some regions)
	//
	// 	- **cloud_essd**: PL1 ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// 	- **general_essd**: premium performance disk
	//
	// When you change the storage type of the instance, note the following items:
	//
	// <props="china">	- If you change the storage type from ESSD to premium performance disk, other parameters such as the instance type and storage capacity must remain unchanged and cannot be modified.
	//
	// <props="china">	- For ApsaraDB RDS for MySQL or ApsaraDB RDS for MariaDB instances whose storage type is standard SSD, you can set this parameter to `cloud_essd` to [upgrade the storage type to ESSD](https://help.aliyun.com/document_detail/314678.html). To perform the upgrade, you must specify both **DBInstanceClass*	- and **DBInstanceStorage**, and their values must be the same as the current values of the instance.
	//
	// <props="china">	- For ApsaraDB RDS for PostgreSQL instances, you can upgrade standard SSDs to [ESSDs of any PL](https://help.aliyun.com/document_detail/96750.html), but **you cannot downgrade ESSDs to standard SSDs**.
	//
	// <props="intl">For ApsaraDB RDS for PostgreSQL instances, you can upgrade standard SSDs to [ESSDs of any PL](https://help.aliyun.com/document_detail/96750.html), but **you cannot downgrade ESSDs to standard SSDs**.
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The ID of the dedicated cluster.
	//
	// example:
	//
	// dhg-7a9****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The type of specification change. Valid values:
	//
	// - **Up*	- (default): upgrade for subscription instances, and upgrade or downgrade for pay-as-you-go instances.
	//
	// - **Down**: downgrade for subscription instances.
	//
	// - **TempUpgrade**: elastic specification change for subscription ApsaraDB RDS for SQL Server instances. This value is required for elastic specification changes.
	//
	// - **Serverless**: elastic setting adjustment for Serverless instances.
	//
	// > If you only change the **DBInstanceStorageType*	- parameter, such as changing standard SSDs to ESSDs, leave this parameter empty.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The effective period of the new configurations. Valid values:
	//
	// > **Changing some configurations may affect the instance**. Read the [Impact section in the feature documentation](https://help.aliyun.com/document_detail/96061.html) before you configure this parameter. We recommend that you perform the operation during off-peak hours.
	//
	// 	- **Immediate*	- (default): The new configurations immediately take effect.
	//
	// 	- **MaintainTime**: The new configurations take effect within the [maintenance window](https://help.aliyun.com/document_detail/610402.html).
	//
	// 	- **ScheduleTime**: The new configurations take effect at a specified time. The specified time must be at least 12 hours later than the current time. The actual switchover time follows the formula: EffectiveTime = ScheduleTime + SwitchTime.
	//
	// example:
	//
	// MaintainTime
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The database engine version. Valid values:
	//
	// <details>
	//
	// <summary>Regular ApsaraDB RDS instances</summary>
	//
	// - MySQL: 5.5, 5.6, 5.7, and 8.0
	//
	// - SQL Server: 2008r2, 08r2_ent_ha, 2012, 2012_ent_ha, 2012_std_ha, 2012_web, 2014_std_ha, 2016_ent_ha, 2016_std_ha, 2016_web, 2017_std_ha, 2017_ent, 2019_std_ha, 2019_ent, 2022_web, 2022_std_ha, 2022_ent, 2025_std, and 2025_ent
	//
	// - PostgreSQL: 10.0, 11.0, 12.0, 13.0, 14.0, and 15.0
	//
	// - MariaDB: 10.3
	//
	// </details>
	//
	// <details>
	//
	// <summary>Serverless ApsaraDB RDS instances (not supported for MariaDB)</summary>
	//
	// - MySQL: 5.7 and 8.0
	//
	// - SQL Server: 2016_std_sl, 2017_std_sl, and 2019_std_sl
	//
	// - PostgreSQL: 14.0, 15.0, and 16.0
	//
	// </details>
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to enable the [Buffer Pool Extension (BPE) feature](https://help.aliyun.com/document_detail/2527067.html) for premium performance disks. Valid values:
	//
	// -  **1**: enables the feature.
	//
	// -  **0**: disables the feature.
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to enable the [16KB atomic write feature](https://help.aliyun.com/document_detail/2858761.html) for ApsaraDB RDS for MySQL. Valid values:
	//
	// - **optimized**: enables the feature.
	//
	// - **none**: disables the feature.
	//
	// example:
	//
	// optimized
	OptimizedWrites *string `json:"OptimizedWrites,omitempty" xml:"OptimizedWrites,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The current billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// - **Serverless*	- (not supported for ApsaraDB RDS for MariaDB instances): Serverless.
	//
	// > To change the billing method to Serverless, **you must configure*	- the auto-start and auto-stop (AutoPause), scaling range (MaxCapacity and MinCapacity), and elastic policy (SwitchForce) parameters. For more information, see [Overview of Serverless ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/411291.html), [Overview of Serverless ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of Serverless ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/607742.html).
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 72329885****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The [target instance type of the read-only instance](https://help.aliyun.com/document_detail/276980.html) when you change the specifications of an ApsaraDB RDS for MySQL High-availability Edition instance with Premium Local SSDs to a cloud disk instance.
	//
	// example:
	//
	// mysqlro.n2.large.1c
	ReadOnlyDBInstanceClass *string `json:"ReadOnlyDBInstanceClass,omitempty" xml:"ReadOnlyDBInstanceClass,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The configurations for changing the Serverless instance.
	ServerlessConfigurationShrink *string `json:"ServerlessConfiguration,omitempty" xml:"ServerlessConfiguration,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and does not need to be configured.
	//
	// example:
	//
	// test
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
	// The time when the specification change is performed. **We recommend that you perform the specification change during off-peak hours**.
	//
	// Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// > - The specified time **must be later than the current time*	- (the time when the API call is initiated). Otherwise, the specification change task fails. After the task fails, you must wait for the order to be automatically canceled before you can initiate the call again.
	//
	// > - If you only increase the storage capacity or change the ESSD storage type, the specification change is performed immediately after submission because it does not affect your business. You do not need to configure this parameter.
	//
	// example:
	//
	// 2019-07-10T13:15:12Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The [minor engine version](https://help.aliyun.com/document_detail/126002.html) of the ApsaraDB RDS for PostgreSQL instance. When you change the instance type and an error indicating that the minor engine version is not supported is returned, you must specify the minor engine version **to upgrade the minor engine version during the instance type change**.
	//
	// Format: `rds_postgres_<Major version number>00_<Minor version number>`. Example for version 12 released on August 30, 2020: `rds_postgres_1200_20200830`.
	//
	// example:
	//
	// rds_postgres_1200_20200830
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The duration of the [elastic upgrade](https://help.aliyun.com/document_detail/95665.html) for the ApsaraDB RDS for SQL Server instance. Unit: days.
	//
	// example:
	//
	// 3
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch ID. The zone where the vSwitch resides must correspond to the zone ID specified in **ZoneId**.
	//
	// - The network type **InstanceNetworkType*	- must be set to **VPC**.
	//
	// - If you specify ZoneSlaveId1 (the secondary zone ID), you must specify two vSwitch IDs and separate them with a comma (,).
	//
	// > You only need to specify the vSwitch when you upgrade the major engine version of an ApsaraDB RDS for SQL Server instance (AllowMajorVersionUpgrade), or when you need to set this parameter to switch the vSwitch.
	//
	// example:
	//
	// vsw-bp1oxflciovg9l7******
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary node. If this value is the same as the value of **ZoneId**, it is a single-zone deployment. If this value is different from the value of **ZoneId**, it is a multi-zone deployment.
	//
	// > You only need to specify the secondary zone when you upgrade the major engine version of an ApsaraDB RDS for SQL Server instance (AllowMajorVersionUpgrade), or when you need to set this parameter to switch the secondary zone.
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

func (s *ModifyDBInstanceSpecShrinkRequest) GetAllocateStrategy() *string {
	return s.AllocateStrategy
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

func (s *ModifyDBInstanceSpecShrinkRequest) SetAllocateStrategy(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.AllocateStrategy = &v
	return s
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
