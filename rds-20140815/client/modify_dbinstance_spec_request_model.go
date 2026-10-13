// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocateStrategy(v string) *ModifyDBInstanceSpecRequest
	GetAllocateStrategy() *string
	SetAllowMajorVersionUpgrade(v bool) *ModifyDBInstanceSpecRequest
	GetAllowMajorVersionUpgrade() *bool
	SetAutoUseCoupon(v bool) *ModifyDBInstanceSpecRequest
	GetAutoUseCoupon() *bool
	SetBurstingEnabled(v bool) *ModifyDBInstanceSpecRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *ModifyDBInstanceSpecRequest
	GetCategory() *string
	SetColdDataEnabled(v bool) *ModifyDBInstanceSpecRequest
	GetColdDataEnabled() *bool
	SetCompressionMode(v string) *ModifyDBInstanceSpecRequest
	GetCompressionMode() *string
	SetDBInstanceClass(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *ModifyDBInstanceSpecRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceStorageType() *string
	SetDedicatedHostGroupId(v string) *ModifyDBInstanceSpecRequest
	GetDedicatedHostGroupId() *string
	SetDirection(v string) *ModifyDBInstanceSpecRequest
	GetDirection() *string
	SetEffectiveTime(v string) *ModifyDBInstanceSpecRequest
	GetEffectiveTime() *string
	SetEngineVersion(v string) *ModifyDBInstanceSpecRequest
	GetEngineVersion() *string
	SetIoAccelerationEnabled(v string) *ModifyDBInstanceSpecRequest
	GetIoAccelerationEnabled() *string
	SetOptimizedWrites(v string) *ModifyDBInstanceSpecRequest
	GetOptimizedWrites() *string
	SetOwnerAccount(v string) *ModifyDBInstanceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceSpecRequest
	GetOwnerId() *int64
	SetPayType(v string) *ModifyDBInstanceSpecRequest
	GetPayType() *string
	SetPromotionCode(v string) *ModifyDBInstanceSpecRequest
	GetPromotionCode() *string
	SetReadOnlyDBInstanceClass(v string) *ModifyDBInstanceSpecRequest
	GetReadOnlyDBInstanceClass() *string
	SetResourceGroupId(v string) *ModifyDBInstanceSpecRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceSpecRequest
	GetResourceOwnerId() *int64
	SetServerlessConfiguration(v *ModifyDBInstanceSpecRequestServerlessConfiguration) *ModifyDBInstanceSpecRequest
	GetServerlessConfiguration() *ModifyDBInstanceSpecRequestServerlessConfiguration
	SetSourceBiz(v string) *ModifyDBInstanceSpecRequest
	GetSourceBiz() *string
	SetSwitchTime(v string) *ModifyDBInstanceSpecRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *ModifyDBInstanceSpecRequest
	GetTargetMinorVersion() *string
	SetUsedTime(v int64) *ModifyDBInstanceSpecRequest
	GetUsedTime() *int64
	SetVSwitchId(v string) *ModifyDBInstanceSpecRequest
	GetVSwitchId() *string
	SetZoneId(v string) *ModifyDBInstanceSpecRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *ModifyDBInstanceSpecRequest
	GetZoneIdSlave1() *string
}

type ModifyDBInstanceSpecRequest struct {
	AllocateStrategy *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	// Specifies whether to perform a [major engine version upgrade](https://help.aliyun.com/document_detail/127458.html) for the ApsaraDB RDS for SQL Server instance. Valid values:
	//
	// - **true**: performs the upgrade.
	//
	// - **false*	- (default): does not perform the upgrade.
	//
	// > - When you perform a major engine version upgrade, you must also specify other required parameters, including DBInstanceId, EngineVersion, DBInstanceClass, Category, ZoneId, and VSwitchId.
	//
	// > - If you want to upgrade the instance to the High-availability Edition or Cluster Edition, you must also specify the ZoneIdSlave1 parameter.
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
	// Specifies whether to enable the [I/O performance burst feature for Premium ESSDs](https://help.aliyun.com/document_detail/2340501.html). Valid values:
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
	// - **HighAvailability**: High-availability Edition
	//
	// - **AlwaysOn**: SQL Server Cluster Edition
	//
	// - **Cluster**: MySQL Cluster Edition
	//
	// - <props="china">**Finance**: Enterprise Edition
	//
	// </details>
	//
	// <details>
	//
	// <summary>Serverless ApsaraDB RDS instances (not supported for MariaDB)</summary>
	//
	// - **serverless_basic**: Serverless Basic Edition (MySQL and PostgreSQL only)
	//
	// - **serverless_standard**: Serverless High-availability Edition (MySQL and PostgreSQL only)
	//
	// - **serverless_ha**: Serverless High-availability Edition (SQL Server only)
	//
	// </details>
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to enable the [data archiving feature for Premium ESSDs](https://help.aliyun.com/document_detail/2701832.html). Valid values:
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
	// The [target instance type](https://help.aliyun.com/document_detail/26312.html). You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) operation to query the available instance types for specification changes.
	//
	// > 	- You must specify at least one of this parameter and **DBInstanceStorage**.
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html) operation to view the current instance type.
	//
	// example:
	//
	// mysql.n8.large.2c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The [target storage capacity](https://help.aliyun.com/document_detail/26312.html). Unit: GB. You can call the [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) operation to query the available storage capacity range for the target instance type.
	//
	// > 	- You must specify at least one of this parameter and **DBInstanceClass**.
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html) operation to view the current storage capacity of the instance.
	//
	// example:
	//
	// 100
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **local_ssd**: premium local SSD
	//
	// 	- **cloud_ssd**: standard SSD (not recommended. Sales have stopped in some regions)
	//
	// 	- **cloud_essd**: PL1 ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// 	- **general_essd**: Premium ESSD
	//
	// If you want to change the instance storage type, take note of the following items:
	//
	// <props="china">	- If you change the instance storage type from ESSD to Premium ESSD, other parameters such as the instance type and storage capacity must remain the same as the original values and cannot be modified.
	//
	// <props="china">	- For ApsaraDB RDS for MySQL or MariaDB instances whose instance storage type is standard SSD, you can set this parameter to `cloud_essd` to [upgrade the instance storage type to ESSD](https://help.aliyun.com/document_detail/314678.html). To perform the upgrade, you must specify both **DBInstanceClass*	- and **DBInstanceStorage**, and their values must be the same as the current values of the instance.
	//
	// <props="china">	- For ApsaraDB RDS for PostgreSQL instances, you can upgrade standard SSDs to [ESSDs of any PL](https://help.aliyun.com/document_detail/96750.html), but **you cannot downgrade ESSDs to standard SSDs**.
	//
	// <props="intl">For ApsaraDB RDS for PostgreSQL instances, you can upgrade standard SSDs to [ESSDs of any PL](https://help.aliyun.com/document_detail/96750.html), but **you cannot downgrade ESSDs to standard SSDs**.
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The dedicated cluster ID.
	//
	// example:
	//
	// dhg-7a9****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The type of the specification change. Valid values:
	//
	// - **Up*	- (default): upgrade for subscription instances, and upgrade or downgrade for pay-as-you-go instances.
	//
	// - **Down**: downgrade for subscription instances.
	//
	// - **TempUpgrade**: elastic specification change for subscription ApsaraDB RDS for SQL Server instances. This value is required for elastic specification changes.
	//
	// - **Serverless**: configuration for adjusting the elastic settings of serverless instances.
	//
	// > If you only change the **DBInstanceStorageType*	- parameter, such as changing standard SSDs to ESSDs, leave this parameter empty.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The effective time of the new configurations. Valid values:
	//
	// > **Changing some configurations may affect the instance**. Read the [Impact section in the feature documentation](https://help.aliyun.com/document_detail/96061.html) carefully before you configure this parameter. We recommend that you do this during off-peak hours.
	//
	// 	- **Immediate*	- (default): takes effect immediately.
	//
	// 	- **MaintainTime**: takes effect within the [maintenance window](https://help.aliyun.com/document_detail/610402.html).
	//
	// 	- **ScheduleTime**: takes effect at a specified time. The specified time must be at least 12 hours later than the current time, and the actual switchover time follows the formula: EffectiveTime = ScheduleTime + SwitchTime.
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
	// Specifies whether to enable the [Buffer Pool Extension (BPE) feature for Premium ESSDs](https://help.aliyun.com/document_detail/2527067.html). Valid values:
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
	// - **Serverless*	- (not supported for ApsaraDB RDS for MariaDB instances): serverless.
	//
	// > To change the billing method to serverless, **you must configure*	- the auto-start and auto-stop (AutoPause), scaling range (MaxCapacity and MinCapacity), and elastic policy (SwitchForce) in subsequent parameters. For more information, see [Overview of serverless ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/411291.html), [Overview of serverless ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of serverless ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/607742.html).
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
	// The [target instance type of the read-only instance](https://help.aliyun.com/document_detail/276980.html) when you change an ApsaraDB RDS for MySQL High-availability Edition instance with premium local SSDs to a cloud disk instance.
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
	// The configurations for changing the serverless instance.
	ServerlessConfiguration *ModifyDBInstanceSpecRequestServerlessConfiguration `json:"ServerlessConfiguration,omitempty" xml:"ServerlessConfiguration,omitempty" type:"Struct"`
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
	// > - The specified time **must be later than the current time*	- (the time when the call is initiated). Otherwise, the specification change task will fail. After the task fails, you must wait for the order to be automatically canceled before you can initiate the call again.
	//
	// > - If you only increase the storage capacity or change the ESSD storage type, the specification change is performed immediately after submission because it does not affect your business. You do not need to configure this parameter.
	//
	// example:
	//
	// 2019-07-10T13:15:12Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The [minor engine version](https://help.aliyun.com/document_detail/126002.html) of the ApsaraDB RDS for PostgreSQL instance. When you change the instance type and an error indicates that the minor engine version is not supported, you must specify the minor engine version **to upgrade the minor engine version during the instance type change**.
	//
	// Format: `rds_postgres_<Major version>00_<Minor version>`. For example, for version 12 and minor version 20200830: `rds_postgres_1200_20200830`.
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
	// - The network type (**InstanceNetworkType**) must be **VPC**.
	//
	// - If you specify ZoneSlaveId1 (the secondary zone ID), you must specify two vSwitch IDs and separate them with a comma (,).
	//
	// > You only need to specify the vSwitch when you upgrade the major engine version of an ApsaraDB RDS for SQL Server instance (AllowMajorVersionUpgrade) or when you switch the vSwitch.
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
	// > You only need to specify the secondary zone when you upgrade the major engine version of an ApsaraDB RDS for SQL Server instance (AllowMajorVersionUpgrade), or you need to configure this parameter when switching the secondary zone.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
}

func (s ModifyDBInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecRequest) GetAllocateStrategy() *string {
	return s.AllocateStrategy
}

func (s *ModifyDBInstanceSpecRequest) GetAllowMajorVersionUpgrade() *bool {
	return s.AllowMajorVersionUpgrade
}

func (s *ModifyDBInstanceSpecRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBInstanceSpecRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDBInstanceSpecRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBInstanceSpecRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *ModifyDBInstanceSpecRequest) GetCompressionMode() *string {
	return s.CompressionMode
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBInstanceSpecRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *ModifyDBInstanceSpecRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyDBInstanceSpecRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBInstanceSpecRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ModifyDBInstanceSpecRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *ModifyDBInstanceSpecRequest) GetOptimizedWrites() *string {
	return s.OptimizedWrites
}

func (s *ModifyDBInstanceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceSpecRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyDBInstanceSpecRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBInstanceSpecRequest) GetReadOnlyDBInstanceClass() *string {
	return s.ReadOnlyDBInstanceClass
}

func (s *ModifyDBInstanceSpecRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceSpecRequest) GetServerlessConfiguration() *ModifyDBInstanceSpecRequestServerlessConfiguration {
	return s.ServerlessConfiguration
}

func (s *ModifyDBInstanceSpecRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ModifyDBInstanceSpecRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceSpecRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *ModifyDBInstanceSpecRequest) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *ModifyDBInstanceSpecRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceSpecRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBInstanceSpecRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *ModifyDBInstanceSpecRequest) SetAllocateStrategy(v string) *ModifyDBInstanceSpecRequest {
	s.AllocateStrategy = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetAllowMajorVersionUpgrade(v bool) *ModifyDBInstanceSpecRequest {
	s.AllowMajorVersionUpgrade = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetAutoUseCoupon(v bool) *ModifyDBInstanceSpecRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetBurstingEnabled(v bool) *ModifyDBInstanceSpecRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetCategory(v string) *ModifyDBInstanceSpecRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetColdDataEnabled(v bool) *ModifyDBInstanceSpecRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetCompressionMode(v string) *ModifyDBInstanceSpecRequest {
	s.CompressionMode = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceClass(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceId(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceStorage(v int32) *ModifyDBInstanceSpecRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceStorageType(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDedicatedHostGroupId(v string) *ModifyDBInstanceSpecRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDirection(v string) *ModifyDBInstanceSpecRequest {
	s.Direction = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetEffectiveTime(v string) *ModifyDBInstanceSpecRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetEngineVersion(v string) *ModifyDBInstanceSpecRequest {
	s.EngineVersion = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetIoAccelerationEnabled(v string) *ModifyDBInstanceSpecRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOptimizedWrites(v string) *ModifyDBInstanceSpecRequest {
	s.OptimizedWrites = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOwnerAccount(v string) *ModifyDBInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOwnerId(v int64) *ModifyDBInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetPayType(v string) *ModifyDBInstanceSpecRequest {
	s.PayType = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetPromotionCode(v string) *ModifyDBInstanceSpecRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetReadOnlyDBInstanceClass(v string) *ModifyDBInstanceSpecRequest {
	s.ReadOnlyDBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceGroupId(v string) *ModifyDBInstanceSpecRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetServerlessConfiguration(v *ModifyDBInstanceSpecRequestServerlessConfiguration) *ModifyDBInstanceSpecRequest {
	s.ServerlessConfiguration = v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetSourceBiz(v string) *ModifyDBInstanceSpecRequest {
	s.SourceBiz = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetSwitchTime(v string) *ModifyDBInstanceSpecRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetTargetMinorVersion(v string) *ModifyDBInstanceSpecRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetUsedTime(v int64) *ModifyDBInstanceSpecRequest {
	s.UsedTime = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetVSwitchId(v string) *ModifyDBInstanceSpecRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetZoneId(v string) *ModifyDBInstanceSpecRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetZoneIdSlave1(v string) *ModifyDBInstanceSpecRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) Validate() error {
	if s.ServerlessConfiguration != nil {
		if err := s.ServerlessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBInstanceSpecRequestServerlessConfiguration struct {
	// Specifies whether to enable the [smart pause and start](https://help.aliyun.com/document_detail/2838448.html) feature for the serverless ApsaraDB RDS for MySQL or PostgreSQL instance. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false*	- (default): disables the feature.
	//
	// > If no connections are established within 10 minutes, the instance enters the paused state. When a connection is established, the instance is automatically woken up.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AutoPause *bool `json:"AutoPause,omitempty" xml:"AutoPause,omitempty"`
	// The **maximum*	- number of ApsaraDB RDS capacity units (RCUs) for the automatic scaling of the serverless instance. Valid values:
	//
	// - MySQL: **1 to 32**
	//
	// - SQL Server: **2 to 16*	- (integers only)
	//
	// - PostgreSQL: **1 to 14**
	//
	// > The value must be greater than or equal to the value of **MinCapacity**.
	//
	// example:
	//
	// 8
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The **minimum*	- number of RCUs for the automatic scaling of the serverless instance. Valid values:
	//
	// - MySQL: **0.5 to 32**
	//
	// - SQL Server: **2 to 16*	- (integers only)
	//
	// - PostgreSQL: **0.5 to 14**
	//
	// > The value must be less than or equal to the value of MaxCapacity.
	//
	// example:
	//
	// 0.5
	MinCapacity *float64 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
	// Specifies whether to enable forced elastic scaling for the serverless ApsaraDB RDS for MySQL or PostgreSQL instance. Elastic scaling of RCUs usually takes effect immediately. However, in some special cases (such as during the execution of large transactions), the scaling cannot be completed immediately. In this case, you can enable this parameter to force the scaling. Valid values:
	//
	// - **true**: enables forced elastic scaling.
	//
	// - **false*	- (default): disables forced elastic scaling.
	//
	// > After this parameter is enabled, **the instance will be unavailable for 30 to 120 seconds during forced scaling**. Use this parameter with caution based on your actual situation.
	//
	// example:
	//
	// false
	SwitchForce *bool `json:"SwitchForce,omitempty" xml:"SwitchForce,omitempty"`
}

func (s ModifyDBInstanceSpecRequestServerlessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecRequestServerlessConfiguration) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) GetAutoPause() *bool {
	return s.AutoPause
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) GetMaxCapacity() *float64 {
	return s.MaxCapacity
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) GetMinCapacity() *float64 {
	return s.MinCapacity
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) GetSwitchForce() *bool {
	return s.SwitchForce
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) SetAutoPause(v bool) *ModifyDBInstanceSpecRequestServerlessConfiguration {
	s.AutoPause = &v
	return s
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) SetMaxCapacity(v float64) *ModifyDBInstanceSpecRequestServerlessConfiguration {
	s.MaxCapacity = &v
	return s
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) SetMinCapacity(v float64) *ModifyDBInstanceSpecRequestServerlessConfiguration {
	s.MinCapacity = &v
	return s
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) SetSwitchForce(v bool) *ModifyDBInstanceSpecRequestServerlessConfiguration {
	s.SwitchForce = &v
	return s
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) Validate() error {
	return dara.Validate(s)
}
