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
	// Specifies whether to enable [major engine version upgrade](https://help.aliyun.com/document_detail/127458.html) for the SQL Server instance. Valid values:
	//
	// example:
	//
	// false
	AllowMajorVersionUpgrade *bool `json:"AllowMajorVersionUpgrade,omitempty" xml:"AllowMajorVersionUpgrade,omitempty"`
	// Specifies whether to use coupons to offset fees. Valid values:
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// Specifies whether to enable the [I/O performance burst feature for Premium ESSDs](https://help.aliyun.com/document_detail/2340501.html). Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The [instance edition](https://help.aliyun.com/document_detail/53509.html). Valid values:
	//
	// > This parameter is required if **EngineVersion*	- is set to a SQL Server version number.
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
	// - **Cluster**: MySQL Cluster Edition.
	//
	// - <props="china">**Finance**: Enterprise Edition
	//
	// </details>
	//
	// <details>
	//
	// <summary>Serverless ApsaraDB RDS instances (not supported for MariaDB)</summary>
	//
	// - **serverless_basic**: Serverless Basic Edition (applicable only to MySQL and PostgreSQL)
	//
	// - **serverless_standard**: Serverless High-availability Edition (applicable only to MySQL and PostgreSQL)
	//
	// - **serverless_ha**: Serverless High-availability Edition (applicable only to SQL Server)
	//
	// </details>
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The [cold data archiving feature](https://help.aliyun.com/document_detail/2701832.html) for premium performance disks. Valid values:
	//
	// example:
	//
	// true
	ColdDataEnabled *bool `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	// The MySQL [storage compression feature](https://help.aliyun.com/document_detail/2861985.html). Valid values:
	//
	// example:
	//
	// on
	CompressionMode *string `json:"CompressionMode,omitempty" xml:"CompressionMode,omitempty"`
	// The [target instance type](https://help.aliyun.com/document_detail/26312.html). You can call [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) to query the instance types to which the instance can be changed.
	//
	// example:
	//
	// mysql.n8.large.2c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID. You can call [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The [target storage capacity](https://help.aliyun.com/document_detail/26312.html). Unit: GB. You can call [DescribeAvailableClasses](https://help.aliyun.com/document_detail/610393.html) to query the available storage capacity range for the target instance type.
	//
	// example:
	//
	// 100
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
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
	// The type of specification change. Valid values:
	//
	// - **Up*	- (default): upgrade of a subscription instance or upgrade/downgrade of a pay-as-you-go instance.
	//
	// - **Down**: downgrade of a subscription instance.
	//
	// - **TempUpgrade**: elastic specification change of a subscription ApsaraDB RDS for SQL Server instance. This value is required for elastic specification changes.
	//
	// - **Serverless**: configuration of elastic settings for a serverless instance.
	//
	// > If you want to change only the **DBInstanceStorageType*	- parameter, for example, from standard SSD to ESSD, leave this parameter empty.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The time when the new configurations take effect. Valid values:
	//
	// > **Changing certain configurations may affect the instance**. Read the [impact section in the feature documentation](https://help.aliyun.com/document_detail/96061.html) before configuring this parameter. Perform this operation during off-peak hours.
	//
	// 	- **Immediate*	- (default): The new configurations take effect immediately.
	//
	// 	- **MaintainTime**: The new configurations take effect during the [maintenance window](https://help.aliyun.com/document_detail/610402.html).
	//
	// 	- **ScheduleTime**: The new configurations take effect at a specified time. The specified time must be at least 12 hours later than the current time. The actual switchover time follows the rule: EffectiveTime = ScheduleTime + SwitchTime.
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
	// - MySQL: 5.5, 5.6, 5.7, 8.0
	//
	// - SQL Server: 2008r2, 08r2_ent_ha, 2012, 2012_ent_ha, 2012_std_ha, 2012_web, 2014_std_ha, 2016_ent_ha, 2016_std_ha, 2016_web, 2017_std_ha, 2017_ent, 2019_std_ha, 2019_ent, 2022_web, 2022_std_ha, 2022_ent, 2025_std, 2025_ent
	//
	// - PostgreSQL: 10.0, 11.0, 12.0, 13.0, 14.0, 15.0
	//
	// - MariaDB: 10.3
	//
	// </details>
	//
	// <details>
	//
	// <summary>Serverless ApsaraDB RDS instances (MariaDB is not supported)</summary>
	//
	// - MySQL: 5.7, 8.0
	//
	// - SQL Server: 2016_std_sl, 2017_std_sl, 2019_std_sl
	//
	// - PostgreSQL: 14.0, 15.0, 16.0
	//
	// </details>
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The [Buffer Pool Extension (BPE) feature](https://help.aliyun.com/document_detail/2527067.html) for premium performance disks. Valid values:
	//
	// -  **1**: Enabled.
	//
	// -  **0**: Not enabled.
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to enable the MySQL [16KB atomic write feature](https://help.aliyun.com/document_detail/2858761.html). Valid values:
	//
	// example:
	//
	// optimized
	OptimizedWrites *string `json:"OptimizedWrites,omitempty" xml:"OptimizedWrites,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// - **Serverless*	- (not supported for MariaDB instances): serverless billing method.
	//
	// > To change the billing method to Serverless, you **must configure the following parameters**: automatic start and stop (AutoPause), scaling range (MaxCapacity and MinCapacity), and elastic policy (SwitchForce). For more information, see [Introduction to MySQL Serverless instances](https://help.aliyun.com/document_detail/411291.html), [Introduction to SQL Server Serverless instances](https://help.aliyun.com/document_detail/604344.html), and [Introduction to PostgreSQL Serverless instances](https://help.aliyun.com/document_detail/607742.html).
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
	// The [target instance type of read-only instances](https://help.aliyun.com/document_detail/276980.html) when you perform an Upgrade/Downgrade to change a MySQL high availability (HA) instance with Premium Local SSDs to a cloud disk instance. This parameter is active only when the instance meets the requirements.
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
	// The serverless instance configuration for the specification change.
	ServerlessConfiguration *ModifyDBInstanceSpecRequestServerlessConfiguration `json:"ServerlessConfiguration,omitempty" xml:"ServerlessConfiguration,omitempty" type:"Struct"`
	// A deprecated parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// test
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
	// The time at which the specification change is performed. **Perform the specification change during off-peak hours.**
	//
	// example:
	//
	// 2019-07-10T13:15:12Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The [minor engine version](https://help.aliyun.com/document_detail/126002.html) of the PostgreSQL instance. If the specification change fails because the minor engine version is not supported, specify this parameter to **upgrade the minor engine version during the specification change**.
	//
	// example:
	//
	// rds_postgres_1200_20200830
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The duration of the SQL Server [elastic upgrade](https://help.aliyun.com/document_detail/95665.html). Unit: days.
	//
	// example:
	//
	// 3
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch ID. The zone of the vSwitch must correspond to the zone ID specified in **ZoneId**.
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
	// The zone ID of the secondary node. If this value is the same as **ZoneId**, the instance uses single-zone deployment. If this value is different from **ZoneId**, the instance uses multi-zone deployment.
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
	// The [intelligent suspension and startup](https://help.aliyun.com/document_detail/2838448.html) feature for MySQL Serverless or PostgreSQL Serverless instances. Valid values:
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AutoPause *bool `json:"AutoPause,omitempty" xml:"AutoPause,omitempty"`
	// The **maximum*	- value of the automatic scaling range for RCUs of the serverless instance. Valid values:
	//
	// example:
	//
	// 8
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The **minimum*	- value of the automatic scaling range for RCUs of the serverless instance. Valid values:
	//
	// example:
	//
	// 0.5
	MinCapacity *float64 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
	// Specifies whether to enable forced scaling for MySQL Serverless or PostgreSQL Serverless instances. Elastic scaling of instance RCUs usually takes effect immediately, but in certain special cases (such as during large transaction execution), scaling cannot be completed instantly. In such cases, you can enable this parameter to force scaling. Valid values:
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
