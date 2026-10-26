// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateDBInstanceShrinkRequest
	GetAmount() *int32
	SetAutoCreateProxy(v bool) *CreateDBInstanceShrinkRequest
	GetAutoCreateProxy() *bool
	SetAutoPay(v bool) *CreateDBInstanceShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *CreateDBInstanceShrinkRequest
	GetAutoRenew() *string
	SetAutoUpgradeMinorVersion(v string) *CreateDBInstanceShrinkRequest
	GetAutoUpgradeMinorVersion() *string
	SetAutoUseCoupon(v bool) *CreateDBInstanceShrinkRequest
	GetAutoUseCoupon() *bool
	SetBabelfishConfig(v string) *CreateDBInstanceShrinkRequest
	GetBabelfishConfig() *string
	SetBpeEnabled(v string) *CreateDBInstanceShrinkRequest
	GetBpeEnabled() *string
	SetBurstingEnabled(v bool) *CreateDBInstanceShrinkRequest
	GetBurstingEnabled() *bool
	SetBusinessInfo(v string) *CreateDBInstanceShrinkRequest
	GetBusinessInfo() *string
	SetCategory(v string) *CreateDBInstanceShrinkRequest
	GetCategory() *string
	SetClientToken(v string) *CreateDBInstanceShrinkRequest
	GetClientToken() *string
	SetColdDataEnabled(v bool) *CreateDBInstanceShrinkRequest
	GetColdDataEnabled() *bool
	SetConnectionMode(v string) *CreateDBInstanceShrinkRequest
	GetConnectionMode() *string
	SetConnectionString(v string) *CreateDBInstanceShrinkRequest
	GetConnectionString() *string
	SetCreateStrategy(v string) *CreateDBInstanceShrinkRequest
	GetCreateStrategy() *string
	SetCustomExtraInfo(v string) *CreateDBInstanceShrinkRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *CreateDBInstanceShrinkRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceShrinkRequest
	GetDBInstanceDescription() *string
	SetDBInstanceNetType(v string) *CreateDBInstanceShrinkRequest
	GetDBInstanceNetType() *string
	SetDBInstanceStorage(v int32) *CreateDBInstanceShrinkRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *CreateDBInstanceShrinkRequest
	GetDBInstanceStorageType() *string
	SetDBIsIgnoreCase(v string) *CreateDBInstanceShrinkRequest
	GetDBIsIgnoreCase() *string
	SetDBParamGroupId(v string) *CreateDBInstanceShrinkRequest
	GetDBParamGroupId() *string
	SetDBTimeZone(v string) *CreateDBInstanceShrinkRequest
	GetDBTimeZone() *string
	SetDedicatedHostGroupId(v string) *CreateDBInstanceShrinkRequest
	GetDedicatedHostGroupId() *string
	SetDeletionProtection(v bool) *CreateDBInstanceShrinkRequest
	GetDeletionProtection() *bool
	SetDryRun(v bool) *CreateDBInstanceShrinkRequest
	GetDryRun() *bool
	SetEncryptionKey(v string) *CreateDBInstanceShrinkRequest
	GetEncryptionKey() *string
	SetEngine(v string) *CreateDBInstanceShrinkRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBInstanceShrinkRequest
	GetEngineVersion() *string
	SetExternalReplication(v bool) *CreateDBInstanceShrinkRequest
	GetExternalReplication() *bool
	SetInstanceNetworkType(v string) *CreateDBInstanceShrinkRequest
	GetInstanceNetworkType() *string
	SetIoAccelerationEnabled(v string) *CreateDBInstanceShrinkRequest
	GetIoAccelerationEnabled() *string
	SetOptimizedWrites(v string) *CreateDBInstanceShrinkRequest
	GetOptimizedWrites() *string
	SetPayType(v string) *CreateDBInstanceShrinkRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceShrinkRequest
	GetPeriod() *string
	SetPort(v string) *CreateDBInstanceShrinkRequest
	GetPort() *string
	SetPrivateIpAddress(v string) *CreateDBInstanceShrinkRequest
	GetPrivateIpAddress() *string
	SetPromotionCode(v string) *CreateDBInstanceShrinkRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateDBInstanceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceShrinkRequest
	GetResourceOwnerId() *int64
	SetRoleARN(v string) *CreateDBInstanceShrinkRequest
	GetRoleARN() *string
	SetSecurityIPList(v string) *CreateDBInstanceShrinkRequest
	GetSecurityIPList() *string
	SetServerlessConfigShrink(v string) *CreateDBInstanceShrinkRequest
	GetServerlessConfigShrink() *string
	SetStorageAutoScale(v string) *CreateDBInstanceShrinkRequest
	GetStorageAutoScale() *string
	SetStorageThreshold(v int32) *CreateDBInstanceShrinkRequest
	GetStorageThreshold() *int32
	SetStorageUpperBound(v int32) *CreateDBInstanceShrinkRequest
	GetStorageUpperBound() *int32
	SetSystemDBCharset(v string) *CreateDBInstanceShrinkRequest
	GetSystemDBCharset() *string
	SetTag(v []*CreateDBInstanceShrinkRequestTag) *CreateDBInstanceShrinkRequest
	GetTag() []*CreateDBInstanceShrinkRequestTag
	SetTargetDedicatedHostIdForLog(v string) *CreateDBInstanceShrinkRequest
	GetTargetDedicatedHostIdForLog() *string
	SetTargetDedicatedHostIdForMaster(v string) *CreateDBInstanceShrinkRequest
	GetTargetDedicatedHostIdForMaster() *string
	SetTargetDedicatedHostIdForSlave(v string) *CreateDBInstanceShrinkRequest
	GetTargetDedicatedHostIdForSlave() *string
	SetTargetMinorVersion(v string) *CreateDBInstanceShrinkRequest
	GetTargetMinorVersion() *string
	SetUsedTime(v string) *CreateDBInstanceShrinkRequest
	GetUsedTime() *string
	SetUserBackupId(v string) *CreateDBInstanceShrinkRequest
	GetUserBackupId() *string
	SetVPCId(v string) *CreateDBInstanceShrinkRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBInstanceShrinkRequest
	GetVSwitchId() *string
	SetWhitelistTemplateList(v string) *CreateDBInstanceShrinkRequest
	GetWhitelistTemplateList() *string
	SetZoneId(v string) *CreateDBInstanceShrinkRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *CreateDBInstanceShrinkRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *CreateDBInstanceShrinkRequest
	GetZoneIdSlave2() *string
}

type CreateDBInstanceShrinkRequest struct {
	// The number of ApsaraDB RDS for MySQL instances to create. This parameter is applicable only when you create multiple ApsaraDB RDS for MySQL instances at a time.
	//
	// Valid values: **1*	- to **20**. Default value: **1**.
	//
	// > - When you create multiple ApsaraDB RDS for MySQL instances, you can use the **Tag.Key*	- and **Tag.Value*	- parameters to add tags to the instances in the same batch. This facilitates instance management by tag after creation.
	//
	// > - After multiple instances are created, the operation returns only the **TaskId**, **RequestId**, and **Message*	- parameters, and does not return other detail parameters. To query the details of a specific instance, you can call the DescribeDBInstanceAttribute operation.
	//
	// > - If **Engine*	- is not set to **MySQL*	- and the value of this parameter is greater than **1**, the operation fails and returns the error code `InvalidParam.Engine`.
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to automatically create a proxy. Valid values:
	//
	// - **true**: enables automatic creation. By default, a general-purpose proxy is created.
	//
	// - **false**: disables automatic creation.
	//
	// example:
	//
	// false
	AutoCreateProxy *bool `json:"AutoCreateProxy,omitempty" xml:"AutoCreateProxy,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: enables automatic payment. You must ensure that your account balance is sufficient.
	//
	// - **false**: generates an order without deducting fees.
	//
	// > The default value is true. If your payment method has an insufficient balance, you can set AutoPay to false. In this case, an unpaid order is generated, and you can log on to the ApsaraDB RDS console to pay the order manually.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. This parameter is valid only when you create a subscription instance. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// > - If you purchase the instance on a monthly basis, the auto-renewal epoch is one month.
	//
	// > - If you purchase the instance on a yearly basis, the auto-renewal epoch is one year.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to enable or disable automatic minor version upgrades. Valid values:
	//
	// - **Auto**: Enabled.
	//
	// - **Manual**: Disabled.
	//
	// example:
	//
	// Auto
	AutoUpgradeMinorVersion *string `json:"AutoUpgradeMinorVersion,omitempty" xml:"AutoUpgradeMinorVersion,omitempty"`
	// Specifies whether to use vouchers. Valid values:
	//
	// 	- **true**: uses vouchers.
	//
	// 	- **false**: does not use vouchers. This is the default value.
	//
	// > If you use vouchers and later need to downgrade the instance configuration, the amount deducted by the vouchers will not be refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The configuration information of Babelfish for ApsaraDB RDS for PostgreSQL.
	//
	// Configuration format: {"babelfishEnabled":"true","migrationMode":"xxxxxxx","masterUsername":"xxxxxxx","masterUserPassword":"xxxxxxxx"}
	//
	// Parameter description:
	//
	// - **babelfishEnabled**: specifies whether to enable Babelfish. Set the value to **true*	- to enable Babelfish. By default, Babelfish is disabled.
	//
	// - **migrationMode**: the database mode. Set the value to **single-db*	- for the single-database mode or **multi-db*	- for the multi-database mode.
	//
	// - **masterUsername**: the initial administrator account. The username must be 1 to 63 characters in length and can contain lowercase letters, digits, and underscores (_). It must start with a letter and end with a letter or digit. It cannot start with pg.
	//
	// - **masterUserPassword**: the password of the administrator account. The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & 	- () _ + - =`.
	//
	// > This parameter is applicable only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Babelfish overview](https://help.aliyun.com/document_detail/428613.html).
	//
	// example:
	//
	// {"babelfishEnabled":"true","migrationMode":"single-db","masterUsername":"babelfish_user","masterUserPassword":"Babelfish123!"}
	BabelfishConfig *string `json:"BabelfishConfig,omitempty" xml:"BabelfishConfig,omitempty"`
	BpeEnabled      *string `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	// Specifies whether to enable the I/O performance burst feature for premium cloud disks. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
	//
	// > For more information about the I/O performance burst feature for premium cloud disks, see [What is a premium cloud disk?](https://help.aliyun.com/document_detail/2340501.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The extended business parameters.
	//
	// example:
	//
	// 121436975448952
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- Regular instances
	//
	//     	- **Basic**: Basic Edition.
	//
	//     	- **HighAvailability**: high-availability series.
	//
	//     	- **cluster**: Cluster Edition for MySQL or PostgreSQL.
	//
	//     	- **AlwaysOn**: Cluster Edition for SQL Server.
	//
	//     	- **Finance**: RDS Enterprise Edition.
	//
	//     > This parameter is required when you create a SQL Server RDS Enterprise Edition<props="china">, Basic Edition Standard Edition, or Basic Edition Enterprise Edition instance. For example, when you create a Basic Edition 2022 RDS Enterprise Edition (2022_ent) instance, you must set this parameter to Basic.
	//
	// 	- Serverless instances
	//
	//     	- **serverless_basic**: Serverless Basic Edition. This value is applicable only to MySQL and PostgreSQL.
	//
	//     	- **serverless_standard**: Serverless high-availability series. This value is applicable only to MySQL and PostgreSQL.
	//
	//     	- **serverless_ha**: Serverless high-availability series for SQL Server.
	//
	//     > This parameter is required when PayType is set to Serverless.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can use the client to generate the value, but you must ensure that the value is unique across different requests. The value can contain a maximum of 64 ASCII characters and cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the [data archiving](https://help.aliyun.com/document_detail/2701832.html) feature for premium cloud disks. Valid values:
	//
	// - **true**: enables the feature.
	//
	// - **false**: disables the feature.
	//
	// example:
	//
	// false
	ColdDataEnabled *bool `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	// The access mode of the instance. Valid values:
	//
	// 	- **Standard**: standard access mode.
	//
	// 	- **Safe**: database proxy mode.
	//
	// The default value is allocated by the RDS system.
	//
	// > SQL Server 2012, 2016, and 2017 instances support only the standard access mode.
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The internal endpoint of the database.
	//
	// Endpoint format: `xxx.mysql.rds.aliyuncs.com`, where `xxx` is the prefix of the instance ID, such as rm-uf6wjk5***.
	//
	// example:
	//
	// rm-uf6wjk5****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The strategy for creating multiple instances at a time. This parameter takes effect only when **Amount*	- is greater than 1. Valid values:
	//
	// 	- **Atomicity**: atomicity. All instances in the same batch are created successfully, or if one instance fails to be created, all instances fail to be created. This is the default value.
	//
	// 	- **Partial**: non-atomicity. The creation of an instance is not affected by other instances in the same batch.
	//
	// example:
	//
	// Atomicity
	CreateStrategy  *string `json:"CreateStrategy,omitempty" xml:"CreateStrategy,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type. You can specify a standard or YiTian instance type. For details, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// To create a Serverless instance, specify one of the following values:
	//
	// - MySQL Basic Edition: **mysql.n2.serverless.1c**
	//
	// - MySQL high-availability series: **mysql.n2.serverless.2c**
	//
	// - SQL Server: **mssql.mem2.serverless.s2**
	//
	// - PostgreSQL Basic Edition: **pg.n2.serverless.1c**
	//
	// - PostgreSQL high-availability series: **pg.n2.serverless.2c**
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql.n2.medium.2c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the instance. The description must be 2 to 255 characters in length. It must start with a letter and can contain letters, digits, Chinese characters, and hyphens (-).
	//
	// > The description cannot start with http:// or https://.
	//
	// example:
	//
	// testInstance
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The network connectivity type of the instance. Set the value to **Intranet**, which indicates an internal network connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// Intranet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// Instance storage capacity of the instance. Unit: GB. Instance storage capacity increases in increments of 5 GB. For more information about the valid values, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **local_ssd**: premium local SSD (recommended).
	//
	// 	- **general_essd**: premium cloud disk (recommended).
	//
	// 	- **cloud_essd**: PL1 ESSD.
	//
	// 	- **cloud_essd2**: PL2 ESSD.
	//
	// 	- **cloud_essd3**: PL3 ESSD.
	//
	// 	- **cloud_ssd**: standard SSD (not recommended, and no longer available in some regions).
	//
	// The default value of this parameter is automatically determined based on the instance type specified by **DBInstanceClass**:
	//
	// 	- If the instance type code indicates a premium local SSD, the default value is **local_ssd**.
	//
	// 	- If the instance type code indicates a cloud disk, the default value is **cloud_essd**.
	//
	// > Serverless instances support only PL1 ESSDs and premium cloud disks.
	//
	// example:
	//
	// general_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// Specifies whether table names are case-sensitive. Valid values:
	//
	// 	- **true**: Table names are not case-sensitive. This is the default value.
	//
	// 	- **false**: Table names are case-sensitive.
	//
	// example:
	//
	// true
	DBIsIgnoreCase *string `json:"DBIsIgnoreCase,omitempty" xml:"DBIsIgnoreCase,omitempty"`
	// The ID of the parameter template. You can call the DescribeParameterGroups operation to query the parameter template ID.
	//
	// > This parameter is supported only by MySQL and PostgreSQL. If you do not specify this parameter, the default parameter template is used. You can also create a custom parameter template and specify it here.
	//
	// example:
	//
	// rpg-sys-****
	DBParamGroupId *string `json:"DBParamGroupId,omitempty" xml:"DBParamGroupId,omitempty"`
	// The time zone of the instance. This parameter takes effect only when **Engine*	- is set to **MySQL*	- or **PostgreSQL**.
	//
	// - If **Engine*	- is set to **MySQL**:
	//
	//     - This parameter configures the UTC time zone. Valid values: **-12:59*	- to **+13:00**.
	//
	//     - For an instance with premium local SSDs, you can specify a named time zone, such as Asia/Hong_Kong. For more information about named time zones, see [Named time zones](https://help.aliyun.com/document_detail/297356.html).
	//
	// - If **Engine*	- is set to **PostgreSQL**:
	//
	//     - This parameter configures a named time zone. UTC time zones are not supported. For more information about named time zones, see [Named time zones](https://help.aliyun.com/document_detail/297356.html).
	//
	//     - This parameter can be configured only for ApsaraDB RDS for PostgreSQL instances with cloud disks.
	//
	// > - You can set the time zone when you purchase a primary instance. Read-only instances do not support time zone settings and inherit the time zone of the primary instance.
	//
	// > - If you do not specify this parameter, the system selects the default time zone based on the region where the instance is created.
	//
	// example:
	//
	// +08:00
	DBTimeZone *string `json:"DBTimeZone,omitempty" xml:"DBTimeZone,omitempty"`
	// The ID of the dedicated cluster host group.
	//
	// This parameter is required when you create an instance in a dedicated cluster.
	//
	// - You can call the DescribeDedicatedHostGroups operation to query the host group information.
	//
	// - If you have not created a host group, you can call the CreateDedicatedHostGroup operation to create one.
	//
	// example:
	//
	// dhg-4n****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to enable the release protection feature for the ApsaraDB RDS instance. This feature is supported only for pay-as-you-go instances. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature. This is the default value.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to perform a precheck for the instance creation operation. Valid values:
	//
	// 	- **true**: performs a precheck but does not create the instance. The check items include request parameters, request format, business limits, and inventory.
	//
	// 	- **false**: sends a normal request and creates the instance directly after the check passes. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the key used for cloud disk encryption in the same region. If you specify this parameter, cloud disk encryption is enabled and cannot be disabled. You must also specify **RoleARN**.
	//
	// You can view the key ID in the Key Management Service (KMS) console or create a new key. For more information, see [Create a CMK](https://help.aliyun.com/document_detail/181610.html).
	//
	// > - For ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and ApsaraDB RDS for SQL Server instances, you do not need to specify this parameter. You only need to specify **RoleARN*	- to use the service key to create an instance with cloud disk encryption enabled.
	//
	// > - If RAM authorization requires RAM users to enable cloud disk encryption when creating instances, the instances cannot be created if cloud disk encryption is not enabled. Configure the following RAM policy for the RAM user:
	//
	// `{"Version":"1","Statement":[{"Effect":"Deny","Action":"rds:CreateDBInstance","Resource":"*","Condition":{"StringEquals":{"rds:DiskEncryptionRequired":"false"}}}]}`
	//
	// 	Warning: This configuration also affects the CreateOrder operation for creating instances in the console.
	//
	// example:
	//
	// 0d24*****-da7b-4786-b981-9a164dxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **MariaDB**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version. Valid values:
	//
	// 	- Regular instances
	//
	//     	- MySQL: **5.5**, **5.6**, **5.7**, and **8.0**
	//
	//     	- SQL Server: **08r2_ent_ha*	- (cloud disk edition, discontinued), **2008r2*	- (premium local SSD, discontinued), **2012*	- (Enterprise Edition standalone), **2012_ent_ha**, **2012_std_ha**, **2012_web**, **2014_ent_ha**, **2014_std_ha**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_ent**, **2017_std_ha**, **2017_web**, **2019_ent**, **2019_std_ha**, **2019_web**, **2022_ent**, **2022_std_ha**, **2022_web**, **2025_ent**, and **2025_std**
	//
	//     	- PostgreSQL: **10.0**, **11.0**, **12.0**, **13.0**, **14.0**, **15.0**, **16.0**, **17.0**, and **18.0**
	//
	//     	- MariaDB: **10.3*	- and **10.6**
	//
	// 	- Serverless instances
	//
	//     	- MySQL: **5.7*	- and **8.0**
	//
	//     	- SQL Server: **2016_std_sl**, **2017_std_sl**, and **2019_std_sl**
	//
	//     	- PostgreSQL: **14.0**, **15.0**, **16.0**, **17.0**, and **18.0**
	//
	// > - MariaDB does not support Serverless instances.
	//
	// > - For SQL Server instances, `_ent` indicates Enterprise Cluster Edition, `_ent_ha` indicates Enterprise Edition, `_std_ha` indicates Standard Edition, and `_web` indicates Web Edition.
	//
	// > - SQL Server 2014 instances are not available on the international site.
	//
	// > - Babelfish for ApsaraDB RDS for PostgreSQL instances support only major version 15.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to enable [native replication for ApsaraDB RDS for MySQL](https://help.aliyun.com/document_detail/2856526.html). Valid values:
	//
	// - **ON**: enables the feature.
	//
	// - **OFF**: disables the feature.
	//
	// example:
	//
	// ON
	ExternalReplication *bool `json:"ExternalReplication,omitempty" xml:"ExternalReplication,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: VPC.
	//
	// 	- **Classic**: classic network.
	//
	// > 	- ApsaraDB RDS for MySQL cloud disk instances support only VPCs. You must set this parameter to **VPC**.
	//
	// > 	- ApsaraDB RDS for PostgreSQL and ApsaraDB RDS for MariaDB instances support only VPCs. You must set this parameter to **VPC**.
	//
	// > 	- ApsaraDB RDS for SQL Server Basic Edition and Web Edition instances support both the classic network and VPCs. Other instances support only VPCs. You must set this parameter to **VPC**.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// Specifies whether to enable the [Buffer Pool Extension (BPE)](https://help.aliyun.com/document_detail/2527067.html) feature for premium cloud disks. Valid values:
	//
	//  - **1**: enables the feature.
	//
	//  - **0**: disables the feature.
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to enable the [16KB atomic write](https://help.aliyun.com/document_detail/2858761.html) feature. Valid values:
	//
	// - **optimized**: enables the feature.
	//
	// - **none**: disables the feature. This is the default value.
	//
	// example:
	//
	// optimized
	OptimizedWrites *string `json:"OptimizedWrites,omitempty" xml:"OptimizedWrites,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// - **Serverless**: serverless billing method. MariaDB instances are not supported. For more information, see [Overview of ApsaraDB RDS for MySQL Serverless instances](https://help.aliyun.com/document_detail/411291.html), [Overview of ApsaraDB RDS for SQL Server Serverless instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of ApsaraDB RDS for PostgreSQL Serverless instances](https://help.aliyun.com/document_detail/607742.html).
	//
	// > The system automatically generates an order and completes the payment. You do not need to manually confirm the payment.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription period of the subscription instance. Valid values:
	//
	// 	- **Year**: yearly subscription.
	//
	// 	- **Month**: monthly subscription.
	//
	// > This parameter is required when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The port used to initialize the instance during creation. Valid values:
	//
	// - MySQL: 1000 to 65534
	//
	// - PostgreSQL, SQL Server, and MariaDB: 1000 to 5999
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The internal IP address of the instance. The IP address must be within the address range of the specified vSwitch. By default, the system automatically allocates an IP address based on **VPCId*	- and **VSwitchId**.
	//
	// example:
	//
	// 172.16.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// aliwood-1688-mobile-promotion
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The global resource descriptor (ARN) used for the authorization of the ApsaraDB RDS service account to access KMS. You can invoke the CheckCloudResourceAuthorized operation to query the ARN.
	//
	// 	Notice: When you enable cloud disk encryption, you must specify **RoleARN**.
	//
	// example:
	//
	// acs:ram::1406****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The [IP address whitelist](https://help.aliyun.com/document_detail/43185.html) of the instance. Separate multiple IP addresses or CIDR blocks with commas (,). Duplicate IP addresses or CIDR blocks are not allowed. You can add up to 1,000 IP addresses or CIDR blocks to the whitelist of a single instance. The following formats are supported:
	//
	// 	- IP addresses, such as 10.10.XX.XX.
	//
	// 	- CIDR blocks, such as 10.10.XX.XX/24. In this example, 24 indicates the length of the prefix in the CIDR block, and the prefix length ranges from 1 to 32.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.XX.XX/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The settings for the ApsaraDB RDS Serverless instance. This parameter is required when you create a Serverless instance.
	//
	// > MariaDB does not support Serverless instances.
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// Specifies whether to enable automatic storage expansion. This parameter is supported only by MySQL and PostgreSQL. Valid values:
	//
	// 	- **Enable**: enables automatic storage expansion.
	//
	// 	- **Disable**: disables automatic storage expansion. This is the default value.
	//
	// > You can also call the ModifyDasInstanceConfig operation to adjust this setting after the instance is created. For more information, see [Enable automatic storage expansion](https://help.aliyun.com/document_detail/173826.html).
	//
	// example:
	//
	// Disable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The threshold percentage that triggers automatic storage expansion. Valid values:
	//
	// 	- **10**
	//
	// 	- **20**
	//
	// 	- **30**
	//
	// 	- **40**
	//
	// 	- **50**
	//
	// > This parameter is required when **StorageAutoScale*	- is set to **Enable**.
	//
	// example:
	//
	// 50
	StorageThreshold *int32 `json:"StorageThreshold,omitempty" xml:"StorageThreshold,omitempty"`
	// The upper limit of the total storage capacity for automatic storage expansion. The automatic expansion does not cause the total storage capacity of the instance to exceed this value. Unit: GB.
	//
	// > - The value must be greater than or equal to 0.
	//
	// > - This parameter is required when **StorageAutoScale*	- is set to **Enable**.
	//
	// example:
	//
	// 2000
	StorageUpperBound *int32 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// Deprecated. You do not need to configure this parameter.
	//
	// example:
	//
	// gbk
	SystemDBCharset *string `json:"SystemDBCharset,omitempty" xml:"SystemDBCharset,omitempty"`
	// The tags.
	Tag []*CreateDBInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the host on which the logger instance is created in the dedicated cluster.
	//
	// This parameter is required when you create an RDS Enterprise Edition instance in a dedicated cluster. If you do not specify this parameter, the system automatically allocates a host by default.
	//
	// - You can invoke the DescribeDedicatedHosts operation to query the host information in the dedicated cluster.
	//
	// - If you do not have a host, you can invoke the CreateDedicatedHost operation to add one.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForLog *string `json:"TargetDedicatedHostIdForLog,omitempty" xml:"TargetDedicatedHostIdForLog,omitempty"`
	// The ID of the host on which the primary instance is created in the dedicated cluster.
	//
	// This parameter is required when you create an instance in a dedicated cluster. If you do not specify this parameter, the system automatically assigns a host.
	//
	// - You can call the DescribeDedicatedHosts operation to query the host information in the host group.
	//
	// - If you do not have a host, you can call the CreateDedicatedHost operation to add one.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// The ID of the host on which the secondary instance is created in the dedicated cluster.
	//
	// This parameter is required when you create a high-availability series or RDS Enterprise Edition instance in a dedicated cluster. If you do not specify this parameter, the system automatically allocates a host by default.
	//
	// - You can invoke the DescribeDedicatedHosts operation to query the host information in the dedicated cluster.
	//
	// - If you do not have a host, you can invoke the CreateDedicatedHost operation to add one.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForSlave *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	// The minor engine version of the ApsaraDB RDS instance. This parameter is required only when you create a MySQL or PostgreSQL instance.
	//
	// Format:
	//
	// 	- MySQL: `<Instance version>_<Numeric version number>`. Examples: `rds_20200229`, `xcluster_20200229`, and `xcluster80_20200229`. Description:
	//
	//     	- rds: high-availability series or Basic Edition.
	//
	//     	- xcluster: MySQL 5.7 RDS Enterprise Edition.
	//
	//     	- xcluster80: MySQL 8.0 RDS Enterprise Edition.
	//
	//     > You can invoke the DescribeDBMiniEngineVersions operation to query the numeric version number. For more information about the differences between versions, see [AliSQL release notes](https://help.aliyun.com/document_detail/96060.html).
	//
	// 	- PostgreSQL: `rds_postgres_<Major version>00_<Minor version number>`. Example: `rds_postgres_1400_20220830`. Description:
	//
	//     	- 1400: PostgreSQL major version 14.
	//
	//     	- 20220830: AliPG minor engine version. You can invoke the DescribeDBMiniEngineVersions operation to query the minor engine version. For more information about the differences between versions, see [PostgreSQL release notes](https://help.aliyun.com/document_detail/126002.html).
	//
	//     > If Babelfish is enabled in **BabelfishConfig**, the minor engine version format for the ApsaraDB RDS for PostgreSQL instance is `rds_postgres_<Major version>00_<AliPG minor engine version>_babelfish`.
	//
	// example:
	//
	// rds_20200229
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- If **Period*	- is set to **Year**, the valid values of **UsedTime*	- are **1*	- to **5**.
	//
	// 	- If **Period*	- is set to **Month**, the valid values of **UsedTime*	- are **1*	- to **11**.
	//
	// > This parameter is required when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// 2
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The ID of the user backup. You can call the ListUserBackupFiles operation to query the backup ID. If you specify this parameter, the instance is created based on the user backup.
	//
	// If you specify this parameter, the following limits apply:
	//
	// - **PayType*	- must be set to **Postpaid**.
	//
	// - **Engine*	- must be set to **MySQL**.
	//
	// - **EngineVersion*	- must be set to **5.7**.
	//
	// - **Category*	- must be set to **Basic**.
	//
	// example:
	//
	// 67798****
	UserBackupId *string `json:"UserBackupId,omitempty" xml:"UserBackupId,omitempty"`
	// The VPC ID.
	//
	// > This parameter takes effect only when the network type **InstanceNetworkType*	- is set to **VPC**.
	//
	// example:
	//
	// vpc-****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// - **Zone mapping**: The zone of the vSwitch specified by VSwitchId must correspond to the zone of the primary node (ZoneId) and the zone of the secondary node (ZoneIdSlave1). If you specify two vSwitch IDs, their sequence must match the sequence of ZoneId and ZoneIdSlave1.
	//
	// - **Network type limit**: **InstanceNetworkType*	- must be set to **VPC**.
	//
	// - **Multiple vSwitches**: If you specify **ZoneIdSlave1*	- and the value is not **Auto**, you must specify two vSwitch IDs and separate them with a comma (,).
	//
	// - **Character limit**: VSwitchId cannot contain special characters such as spaces, exclamation marks (!), number signs (#), yen signs (¥), ampersands (&), or percent signs (%).
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The whitelist templates.
	//
	// When you need to configure multiple IP addresses, separate multiple IP addresses or CIDR blocks with commas (,). Do not include spaces before or after the commas. Example: `192.168.0.1,172.16.213.9`.
	//
	// example:
	//
	// 192.168.0.1,172.16.213.9
	WhitelistTemplateList *string `json:"WhitelistTemplateList,omitempty" xml:"WhitelistTemplateList,omitempty"`
	// The zone ID of the primary node.
	//
	// - If you specify a VPC and a vSwitch, you must specify the zone ID of the target vSwitch. Otherwise, the instance cannot be created.
	//
	// - For a high-availability series instance, you must also specify **ZoneIdSlave1*	- to determine whether the instance uses a single-zone deployment or a multi-zone deployment.
	//
	// - For an RDS Enterprise Edition instance, you must also specify **ZoneIdSlave1*	- and **ZoneIdSlave2*	- to determine whether the instance uses a single-zone deployment or a multi-zone deployment.
	//
	// - For an ApsaraDB RDS for MySQL Cluster Edition instance, you must also specify **ZoneIdSlave1*	- for a two-node instance, and **ZoneIdSlave1*	- and **ZoneIdSlave2*	- for a three-node instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary node.
	//
	// - If you set this parameter to **Auto**, the instance uses a multi-zone deployment, and the system automatically selects a zone for the secondary node.
	//
	// - If you set this parameter to the same value as **ZoneId**, the instance uses a single-zone deployment.
	//
	// - If you set this parameter to a value different from **ZoneId**, the instance uses a multi-zone deployment.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The zone ID of the second secondary node. ApsaraDB RDS for MySQL Cluster Edition instances support the creation of one or two secondary nodes when you create the instance. If you want to create a second secondary node, you can use this parameter to specify the zone of the second secondary node.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneIdSlave2 *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s CreateDBInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateDBInstanceShrinkRequest) GetAutoCreateProxy() *bool {
	return s.AutoCreateProxy
}

func (s *CreateDBInstanceShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateDBInstanceShrinkRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateDBInstanceShrinkRequest) GetAutoUpgradeMinorVersion() *string {
	return s.AutoUpgradeMinorVersion
}

func (s *CreateDBInstanceShrinkRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateDBInstanceShrinkRequest) GetBabelfishConfig() *string {
	return s.BabelfishConfig
}

func (s *CreateDBInstanceShrinkRequest) GetBpeEnabled() *string {
	return s.BpeEnabled
}

func (s *CreateDBInstanceShrinkRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateDBInstanceShrinkRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateDBInstanceShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateDBInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceShrinkRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *CreateDBInstanceShrinkRequest) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *CreateDBInstanceShrinkRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDBInstanceShrinkRequest) GetCreateStrategy() *string {
	return s.CreateStrategy
}

func (s *CreateDBInstanceShrinkRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CreateDBInstanceShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CreateDBInstanceShrinkRequest) GetDBIsIgnoreCase() *string {
	return s.DBIsIgnoreCase
}

func (s *CreateDBInstanceShrinkRequest) GetDBParamGroupId() *string {
	return s.DBParamGroupId
}

func (s *CreateDBInstanceShrinkRequest) GetDBTimeZone() *string {
	return s.DBTimeZone
}

func (s *CreateDBInstanceShrinkRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *CreateDBInstanceShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateDBInstanceShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDBInstanceShrinkRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateDBInstanceShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBInstanceShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceShrinkRequest) GetExternalReplication() *bool {
	return s.ExternalReplication
}

func (s *CreateDBInstanceShrinkRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CreateDBInstanceShrinkRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *CreateDBInstanceShrinkRequest) GetOptimizedWrites() *string {
	return s.OptimizedWrites
}

func (s *CreateDBInstanceShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceShrinkRequest) GetPort() *string {
	return s.Port
}

func (s *CreateDBInstanceShrinkRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateDBInstanceShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateDBInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceShrinkRequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *CreateDBInstanceShrinkRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDBInstanceShrinkRequest) GetServerlessConfigShrink() *string {
	return s.ServerlessConfigShrink
}

func (s *CreateDBInstanceShrinkRequest) GetStorageAutoScale() *string {
	return s.StorageAutoScale
}

func (s *CreateDBInstanceShrinkRequest) GetStorageThreshold() *int32 {
	return s.StorageThreshold
}

func (s *CreateDBInstanceShrinkRequest) GetStorageUpperBound() *int32 {
	return s.StorageUpperBound
}

func (s *CreateDBInstanceShrinkRequest) GetSystemDBCharset() *string {
	return s.SystemDBCharset
}

func (s *CreateDBInstanceShrinkRequest) GetTag() []*CreateDBInstanceShrinkRequestTag {
	return s.Tag
}

func (s *CreateDBInstanceShrinkRequest) GetTargetDedicatedHostIdForLog() *string {
	return s.TargetDedicatedHostIdForLog
}

func (s *CreateDBInstanceShrinkRequest) GetTargetDedicatedHostIdForMaster() *string {
	return s.TargetDedicatedHostIdForMaster
}

func (s *CreateDBInstanceShrinkRequest) GetTargetDedicatedHostIdForSlave() *string {
	return s.TargetDedicatedHostIdForSlave
}

func (s *CreateDBInstanceShrinkRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *CreateDBInstanceShrinkRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDBInstanceShrinkRequest) GetUserBackupId() *string {
	return s.UserBackupId
}

func (s *CreateDBInstanceShrinkRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceShrinkRequest) GetWhitelistTemplateList() *string {
	return s.WhitelistTemplateList
}

func (s *CreateDBInstanceShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceShrinkRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *CreateDBInstanceShrinkRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *CreateDBInstanceShrinkRequest) SetAmount(v int32) *CreateDBInstanceShrinkRequest {
	s.Amount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetAutoCreateProxy(v bool) *CreateDBInstanceShrinkRequest {
	s.AutoCreateProxy = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetAutoPay(v bool) *CreateDBInstanceShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetAutoRenew(v string) *CreateDBInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetAutoUpgradeMinorVersion(v string) *CreateDBInstanceShrinkRequest {
	s.AutoUpgradeMinorVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetAutoUseCoupon(v bool) *CreateDBInstanceShrinkRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetBabelfishConfig(v string) *CreateDBInstanceShrinkRequest {
	s.BabelfishConfig = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetBpeEnabled(v string) *CreateDBInstanceShrinkRequest {
	s.BpeEnabled = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetBurstingEnabled(v bool) *CreateDBInstanceShrinkRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetBusinessInfo(v string) *CreateDBInstanceShrinkRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCategory(v string) *CreateDBInstanceShrinkRequest {
	s.Category = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClientToken(v string) *CreateDBInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetColdDataEnabled(v bool) *CreateDBInstanceShrinkRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetConnectionMode(v string) *CreateDBInstanceShrinkRequest {
	s.ConnectionMode = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetConnectionString(v string) *CreateDBInstanceShrinkRequest {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCreateStrategy(v string) *CreateDBInstanceShrinkRequest {
	s.CreateStrategy = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCustomExtraInfo(v string) *CreateDBInstanceShrinkRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceClass(v string) *CreateDBInstanceShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceDescription(v string) *CreateDBInstanceShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceNetType(v string) *CreateDBInstanceShrinkRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceStorage(v int32) *CreateDBInstanceShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBInstanceStorageType(v string) *CreateDBInstanceShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBIsIgnoreCase(v string) *CreateDBInstanceShrinkRequest {
	s.DBIsIgnoreCase = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBParamGroupId(v string) *CreateDBInstanceShrinkRequest {
	s.DBParamGroupId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBTimeZone(v string) *CreateDBInstanceShrinkRequest {
	s.DBTimeZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDedicatedHostGroupId(v string) *CreateDBInstanceShrinkRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDeletionProtection(v bool) *CreateDBInstanceShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDryRun(v bool) *CreateDBInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEncryptionKey(v string) *CreateDBInstanceShrinkRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngine(v string) *CreateDBInstanceShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngineVersion(v string) *CreateDBInstanceShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetExternalReplication(v bool) *CreateDBInstanceShrinkRequest {
	s.ExternalReplication = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetInstanceNetworkType(v string) *CreateDBInstanceShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetIoAccelerationEnabled(v string) *CreateDBInstanceShrinkRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetOptimizedWrites(v string) *CreateDBInstanceShrinkRequest {
	s.OptimizedWrites = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPayType(v string) *CreateDBInstanceShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPeriod(v string) *CreateDBInstanceShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPort(v string) *CreateDBInstanceShrinkRequest {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPrivateIpAddress(v string) *CreateDBInstanceShrinkRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPromotionCode(v string) *CreateDBInstanceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetRegionId(v string) *CreateDBInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetResourceOwnerId(v int64) *CreateDBInstanceShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetRoleARN(v string) *CreateDBInstanceShrinkRequest {
	s.RoleARN = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSecurityIPList(v string) *CreateDBInstanceShrinkRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetServerlessConfigShrink(v string) *CreateDBInstanceShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetStorageAutoScale(v string) *CreateDBInstanceShrinkRequest {
	s.StorageAutoScale = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetStorageThreshold(v int32) *CreateDBInstanceShrinkRequest {
	s.StorageThreshold = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetStorageUpperBound(v int32) *CreateDBInstanceShrinkRequest {
	s.StorageUpperBound = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSystemDBCharset(v string) *CreateDBInstanceShrinkRequest {
	s.SystemDBCharset = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTag(v []*CreateDBInstanceShrinkRequestTag) *CreateDBInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTargetDedicatedHostIdForLog(v string) *CreateDBInstanceShrinkRequest {
	s.TargetDedicatedHostIdForLog = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTargetDedicatedHostIdForMaster(v string) *CreateDBInstanceShrinkRequest {
	s.TargetDedicatedHostIdForMaster = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTargetDedicatedHostIdForSlave(v string) *CreateDBInstanceShrinkRequest {
	s.TargetDedicatedHostIdForSlave = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTargetMinorVersion(v string) *CreateDBInstanceShrinkRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetUsedTime(v string) *CreateDBInstanceShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetUserBackupId(v string) *CreateDBInstanceShrinkRequest {
	s.UserBackupId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVPCId(v string) *CreateDBInstanceShrinkRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVSwitchId(v string) *CreateDBInstanceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetWhitelistTemplateList(v string) *CreateDBInstanceShrinkRequest {
	s.WhitelistTemplateList = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneId(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneIdSlave1(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneIdSlave2(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDBInstanceShrinkRequestTag struct {
	// The tag key. You can specify this parameter to add a tag to the instance.
	//
	// 	- If the specified tag key already exists, the tag key is directly added to the instance. You can call the ListTagResources operation to query existing tags.
	//
	// 	- If the specified tag key does not exist, the tag key is created and then added to the instance.
	//
	// 	- Empty strings are not allowed.
	//
	// 	- This parameter must be used together with **Tag.Value**.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key. You can specify this parameter to add a tag to the instance.
	//
	// 	- If the specified tag value exists for the corresponding tag key, the tag value is directly added to the instance. You can call the ListTagResources operation to query existing tags.
	//
	// 	- If the specified tag value does not exist for the corresponding tag key, the tag value is created and then added to the instance.
	//
	// 	- This parameter must be used together with **Tag.Key**.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBInstanceShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDBInstanceShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDBInstanceShrinkRequestTag) SetKey(v string) *CreateDBInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBInstanceShrinkRequestTag) SetValue(v string) *CreateDBInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDBInstanceShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
