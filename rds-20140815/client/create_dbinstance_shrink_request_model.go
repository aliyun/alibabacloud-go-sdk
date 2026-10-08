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
	// The number of ApsaraDB RDS for MySQL instances to create. This parameter applies only to batch creation of ApsaraDB RDS for MySQL instances.
	//
	// Valid values: **1*	- to **20**. Default value: **1**.
	//
	// > - When creating multiple ApsaraDB RDS for MySQL instances, consider using **Tag.Key*	- and **Tag.Value*	- to tag all instances in the same batch, so that you can manage them by tag after creation.
	//
	// > - After multiple ApsaraDB RDS for MySQL instances are created, the operation returns only **TaskId**, **RequestId**, and **Message**. Other details are not returned. To query the details of individual instances, call DescribeDBInstanceAttribute.
	//
	// > - If **engine*	- is not set to **MySQL*	- and this parameter is set to a value greater than **1**, the operation fails and returns the error code `InvalidParam.Engine`.
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to automatically create a proxy. Valid values:
	//
	// - **true**: enables automatic automatic creation. The default proxy type is general-purpose.
	//
	// - **false**: disables automatic automatic creation.
	//
	// example:
	//
	// false
	AutoCreateProxy *bool `json:"AutoCreateProxy,omitempty" xml:"AutoCreateProxy,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: enables automatic payment. Make sure that your account balance is sufficient.
	//
	// - **false**: generates an order without deducting fees.
	//
	//
	//
	//
	// > The default value is true. If your payment method has insufficient balance, set AutoPay to false. This generates an unpaid order, which you can pay for in the ApsaraDB RDS console.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. This parameter is valid only for subscription instances. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// > - If you purchase the instance on a monthly basis, the auto-renewal cycle is one month.
	//
	// > - If you purchase the instance on a yearly basis, the auto-renewal cycle is one year.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// 	- **true**: uses a coupon.
	//
	// 	- **false*	- (default): does not use a coupon.
	//
	// > If you use a coupon and then perform a downgrade, the amount offset by the coupon is not refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The Babelfish configuration for ApsaraDB RDS for PostgreSQL instances.
	//
	// Configuration format: {"babelfishEnabled":"true","migrationMode":"xxxxxxx","masterUsername":"xxxxxxx","masterUserPassword":"xxxxxxxx"}
	//
	// The parameters are described as follows:
	//
	// - **babelfishEnabled**: specifies whether to enable Babelfish. Set to **true*	- to enable. Babelfish is disabled by default if this parameter is not configured.
	//
	// - **migrationMode**: the database mode. Set to **single-db*	- for single-database mode or **multi-db*	- for multi-database mode.
	//
	// - **masterUsername**: the initial administrator account name. The name can contain lowercase letters, digits, and underscores (_), must start with a letter, must end with a letter or digit, can be up to 63 characters in length, and cannot start with pg.
	//
	// - **masterUserPassword**: the password of the administrator account. The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The password must be 8 to 32 characters in length. Special characters include `! @ # $ % ^ & 	- () _ + - =`.
	//
	// > This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://help.aliyun.com/document_detail/428613.html).
	//
	// example:
	//
	// {"babelfishEnabled":"true","migrationMode":"single-db","masterUsername":"babelfish_user","masterUserPassword":"Babelfish123!"}
	BabelfishConfig *string `json:"BabelfishConfig,omitempty" xml:"BabelfishConfig,omitempty"`
	BpeEnabled      *string `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	// Specifies whether to enable the I/O performance burst feature for premium performance disks (cloud disks). Valid values:
	//
	// 	- **true**: enabled.
	//
	// 	- **false**: disabled.
	//
	// > For more information about the I/O performance burst feature for premium performance disks, see [What is a premium performance disk](https://help.aliyun.com/document_detail/2340501.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The business extension parameter.
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
	//     	- **HighAvailability**: High-availability Edition.
	//
	//     	- **cluster**: MySQL or PostgreSQL Cluster Edition.
	//
	//     	- **AlwaysOn**: SQL Server Cluster Edition.
	//
	//     	- **Finance**: RDS Enterprise Edition.
	//
	//     > This parameter is required when you create a SQL Server Enterprise Cluster Edition<props="china">, Basic Edition Standard Edition, or Basic Edition Enterprise Edition instance. For example, to create a Basic Edition 2022 Enterprise Cluster Edition (2022_ent) instance, set this parameter to Basic.
	//
	// 	- Serverless instances
	//
	//     	- **serverless_basic**: Serverless Basic Edition. (Applicable to MySQL and PostgreSQL only.)
	//
	//     	- **serverless_standard**: Serverless High-availability Edition. (Applicable to MySQL and PostgreSQL only.)
	//
	//     	- **serverless_ha**: SQL Server Serverless High-availability Edition.
	//
	//     > This parameter is required when PayType is set to Serverless.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotency of the request. The token is generated by the client and must be unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the [cold data archiving](https://help.aliyun.com/document_detail/2701832.html) feature for premium performance disks (cloud disks). Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
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
	// > SQL Server 2012, 2016, and 2017 support only standard access mode.
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The internal endpoint of the database.
	//
	// The endpoint format is `xxx.mysql.rds.aliyuncs.com`, where `xxx` is the prefix of the instance ID, such as rm-uf6wjk5***.
	//
	// example:
	//
	// rm-uf6wjk5****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The batch instance creation strategy. This parameter takes effect only when **Amount*	- is greater than 1. Valid values:
	//
	// 	- **Atomicity*	- (default): atomic. All instances in the same batch must be created successfully. If any instance fails to be created, all instances in the batch fail.
	//
	// 	- **Partial**: non-atomic. The creation of each instance is independent of other instances in the same batch.
	//
	// example:
	//
	// Atomicity
	CreateStrategy  *string `json:"CreateStrategy,omitempty" xml:"CreateStrategy,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type. You can specify a standard or YiTian instance type. For details, see [Primary instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// To create a serverless instance, use one of the following values:
	//
	// - MySQL Basic Edition: **mysql.n2.serverless.1c**
	//
	// - MySQL High-availability Edition: **mysql.n2.serverless.2c**
	//
	// - SQL Server: **mssql.mem2.serverless.s2**
	//
	// - PostgreSQL Basic Edition: **pg.n2.serverless.1c**
	//
	// - PostgreSQL High-availability Edition: **pg.n2.serverless.2c**
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql.n2.medium.2c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance name. The name must be 2 to 255 characters in length. It must start with a Chinese character or an English letter, and can contain digits, Chinese characters, English letters, and hyphens (-).
	//
	// >The name cannot start with http:// or https://.
	//
	// example:
	//
	// testInstance
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The network connectivity type of the instance. Set this parameter to **Intranet**, which indicates an internal network connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// Intranet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The instance storage capacity. Unit: GB. The value increments in steps of 5 GB. For the valid values, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **local_ssd**: instance with Premium Local SSDs (recommended).
	//
	// 	- **general_essd**: premium performance disk (recommended).
	//
	// 	- **cloud_essd**: PL1 ESSD.
	//
	// 	- **cloud_essd2**: PL2 ESSD.
	//
	// 	- **cloud_essd3**: PL3 ESSD.
	//
	// 	- **cloud_ssd**: standard SSD (not recommended. No longer available in some regions).
	//
	// The default value of this parameter is automatically determined based on the instance type specified in **DBInstanceClass**:
	//
	// 	- If the instance type is an instance with Premium Local SSDs, the default value is **local_ssd**.
	//
	// 	- If the instance type is a cloud disk type, the default value is **cloud_essd**.
	//
	// > Serverless instances support only PL1 ESSDs and premium performance disks.
	//
	// example:
	//
	// general_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// Specifies whether table names are case-insensitive. Valid values:
	//
	// 	- **true**: case-insensitive (default).
	//
	// 	- **false**: case-sensitive.
	//
	// example:
	//
	// true
	DBIsIgnoreCase *string `json:"DBIsIgnoreCase,omitempty" xml:"DBIsIgnoreCase,omitempty"`
	// The parameter template ID. You can call DescribeParameterGroups to query the ID.
	//
	// > This parameter is supported only for MySQL and PostgreSQL instances. If you do not specify this parameter, the system default parameter template is used. You can also create a custom parameter template and specify it here.
	//
	// example:
	//
	// rpg-sys-****
	DBParamGroupId *string `json:"DBParamGroupId,omitempty" xml:"DBParamGroupId,omitempty"`
	// The time zone of the instance. This parameter takes effect only when **Engine*	- is set to **MySQL*	- or **PostgreSQL**.
	//
	// - When **Engine*	- is **MySQL**:
	//
	//     - This parameter configures the UTC time zone. Valid values: **-12:59*	- to **+13:00**.
	//
	//     - Instances with Premium Local SSDs support named time zones, such as Asia/Hong_Kong. For more information about named time zones, see [Named time zone reference](https://help.aliyun.com/document_detail/297356.html).
	//
	// - When **Engine*	- is **PostgreSQL**:
	//
	//     - This parameter configures a named time zone. UTC time zones are not supported. For more information about named time zones, see [Named time zone reference](https://help.aliyun.com/document_detail/297356.html).
	//
	//     - This parameter can be configured only for PostgreSQL instances with cloud disks.
	//
	// > - You can configure the time zone when creating a primary instance. Read-only instances do not support custom time zones and inherit the time zone of the primary instance.
	//
	// > - If you do not specify this parameter, the system selects a default time zone based on the region where you purchase the instance.
	//
	// example:
	//
	// +08:00
	DBTimeZone *string `json:"DBTimeZone,omitempty" xml:"DBTimeZone,omitempty"`
	// The ID of the dedicated host group.
	//
	// This parameter is required when you create an ApsaraDB RDS instance in a dedicated cluster.
	//
	// - You can call DescribeDedicatedHostGroups to query the host group information.
	//
	// - If you have not created a host group, call CreateDedicatedHostGroup to create one.
	//
	// example:
	//
	// dhg-4n****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to enable the release protection feature for the RDS instance. This parameter is supported only for pay-as-you-go instances. Valid values:
	//
	// 	- **true**: enables release protection.
	//
	// 	- **false**: disables release protection (default).
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to perform a dry run for this instance creation operation. Valid values:
	//
	// 	- **true**: performs a dry run without creating the instance. The dry run checks the request parameters, request format, business limits, and resource availability.
	//
	// 	- **false**: sends a normal request and creates the instance directly after the check passes (default).
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the cloud disk encryption key in the same region. Specifying this parameter enables cloud disk encryption (which cannot be disabled after it is enabled) and requires you to also specify **RoleARN**.
	//
	// You can view the key ID in the Key Management Service console or create a new key. For more information, see [Create a key](https://help.aliyun.com/document_detail/181610.html).
	//
	// > - For ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and ApsaraDB RDS for SQL Server instances, you can omit this parameter and specify only **RoleARN*	- to create a cloud disk-encrypted instance using a service key.
	//
	// > - To allow RAM users to create instances only when cloud disk encryption is enabled, configure the following RAM authorization policy. If cloud disk encryption is not enabled, the RAM user cannot create instances:
	//
	// `{"Version":"1","Statement":[{"Effect":"Deny","Action":"rds:CreateDBInstance","Resource":"*","Condition":{"StringEquals":{"rds:DiskEncryptionRequired":"false"}}}]}`
	//
	// 	Warning: This configuration also affects the CreateOrder operation that is called when you create an instance in the console.
	//
	// example:
	//
	// 0d24*****-da7b-4786-b981-9a164dxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The database engine type. Valid values:
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
	//     	- MySQL: **5.5**, **5.6**, **5.7**, **8.0**
	//
	//     	- SQL Server: **08r2_ent_ha*	- (cloud disk, discontinued), **2008r2*	- (Premium Local SSD, discontinued), **2012*	- (Enterprise Edition single-node), **2012_ent_ha**, **2012_std_ha**, **2012_web**, **2014_ent_ha**, **2014_std_ha**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_ent**, **2017_std_ha**, **2017_web**, **2019_ent**, **2019_std_ha**, **2019_web**, **2022_ent**, **2022_std_ha**, **2022_web**, **2025_ent**, **2025_std**
	//
	//     	- PostgreSQL: **10.0**, **11.0**, **12.0**, **13.0**, **14.0**, **15.0**, **16.0**, **17.0**, **18.0**
	//
	//     	- MariaDB: **10.3**, **10.6**
	//
	// 	- Serverless instances
	//
	//     	- MySQL: **5.7**, **8.0**
	//
	//     	- SQL Server: **2016_std_sl**, **2017_std_sl**, **2019_std_sl**
	//
	//     	- PostgreSQL: **14.0**, **15.0**, **16.0**, **17.0**, **18.0**
	//
	// > - MariaDB does not support serverless instances.
	//
	// > - In SQL Server instance versions, `_ent` indicates Enterprise Cluster Edition, `_ent_ha` indicates Enterprise Edition, `_std_ha` indicates Standard Edition, and `_web` indicates Web Edition.
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
	// Specifies whether to enable [ApsaraDB RDS for MySQL native replication](https://help.aliyun.com/document_detail/2856526.html). Valid values:
	//
	// - **ON**: enabled.
	//
	// - **OFF**: disabled.
	//
	// example:
	//
	// ON
	ExternalReplication *bool `json:"ExternalReplication,omitempty" xml:"ExternalReplication,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud.
	//
	// 	- **Classic**: classic network.
	//
	// > 	- ApsaraDB RDS for MySQL cloud disk instances support only VPCs. Set this parameter to **VPC**.
	//
	// > 	- ApsaraDB RDS for PostgreSQL and MariaDB instances support only VPCs. Set this parameter to **VPC**.
	//
	// > 	- ApsaraDB RDS for SQL Server Basic Edition and Web Edition instances support both classic networks and VPCs. All other instances support only VPCs. Set this parameter to **VPC**.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// Specifies whether to enable the [Buffer Pool Extension (BPE)](https://help.aliyun.com/document_detail/2527067.html) feature for premium performance disks (cloud disks). Valid values:
	//
	//  - **1**: enabled.
	//
	//  - **0**: disabled.
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to enable the [16KB atomic write](https://help.aliyun.com/document_detail/2858761.html) feature. Valid values:
	//
	// - **optimized**: enabled.
	//
	// - **none*	- (default): disabled.
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
	// - **Serverless**: serverless billing method. MariaDB instances do not support this billing method. For more information, see [Overview of MySQL Serverless instances](https://help.aliyun.com/document_detail/411291.html), [Overview of SQL Server Serverless instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of PostgreSQL Serverless instances](https://help.aliyun.com/document_detail/607742.html).
	//
	// >The system automatically generates and pays for the order. No manual payment confirmation is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the prepaid instance. Valid values:
	//
	// 	- **Year**: subscription on a yearly basis.
	//
	// 	- **Month**: subscription on a monthly basis.
	//
	// > This parameter is required if the billing method is **Prepaid**.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The port to initialize when creating the ApsaraDB RDS instance. Valid values:
	//
	// - MySQL: 1000 to 65534
	//
	// - PostgreSQL, SQL Server, MariaDB: 1000 to 5999
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Settings for the internal network IP address of the instance. The IP address must be within the address range of the specified vSwitch. By default, the system automatically allocates an IP address based on **VPCId*	- and **vSwitchId**.
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
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) to query the region ID.
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
	// The global resource descriptor (ARN) that grants the RDS service account authorization to access KMS on behalf of the primary account. You can call CheckCloudResourceAuthorized to query the ARN information.
	//
	// 	Notice: You must specify **RoleARN*	- when you enable cloud disk encryption.
	//
	// example:
	//
	// acs:ram::1406****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The [IP whitelist](https://help.aliyun.com/document_detail/43185.html) of the instance. Separate multiple entries with commas (,). Duplicate entries are not allowed. You can add up to 1,000 IP addresses or CIDR blocks to a single instance. The following formats are supported:
	//
	// 	- IP address format, for example: 10.10.XX.XX.
	//
	// 	- CIDR block format, for example: 10.10.XX.XX/24 (classless inter-domain routing, where 24 indicates the length of the prefix in the address, ranging from 1 to 32).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.XX.XX/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The settings for the serverless ApsaraDB RDS instance. This parameter is required when you create a serverless instance.
	//
	// >MariaDB does not support serverless instances.
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// Specifies whether to enable automatic storage expansion. This parameter is supported only for MySQL and PostgreSQL instances. Valid values:
	//
	// 	- **Enable**: enables automatic storage expansion.
	//
	// 	- **Disable**: disables automatic storage expansion (default).
	//
	// >You can also call ModifyDasInstanceConfig after the instance is created to adjust this setting. For more information, see [Configure automatic storage expansion](https://help.aliyun.com/document_detail/173826.html).
	//
	// example:
	//
	// Disable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The threshold (percentage) that triggers automatic storage expansion. Valid values:
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
	// >This parameter is required when **StorageAutoScale*	- is set to **Enable**.
	//
	// example:
	//
	// 50
	StorageThreshold *int32 `json:"StorageThreshold,omitempty" xml:"StorageThreshold,omitempty"`
	// The maximum total storage capacity allowed for automatic storage expansion. Automatic storage expansion does not cause the total storage capacity of the instance to exceed this value. Unit: GB.
	//
	// > - The value must be greater than or equal to 0.
	//
	// > - This parameter is required when **StorageAutoScale*	- is set to **Enable**.
	//
	// example:
	//
	// 2000
	StorageUpperBound *int32 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// This parameter is deprecated. You do not need to configure it.
	//
	// example:
	//
	// gbk
	SystemDBCharset *string `json:"SystemDBCharset,omitempty" xml:"SystemDBCharset,omitempty"`
	// The list of tags.
	Tag []*CreateDBInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The host ID of the logger instance in the dedicated cluster.
	//
	// This parameter is required when you create an ApsaraDB RDS Enterprise Edition instance in a dedicated cluster. If you do not specify this parameter, the system automatically assigns a host.
	//
	// - You can call DescribeDedicatedHosts to query the host information in the dedicated cluster.
	//
	// - If you have not added a host, call CreateDedicatedHost to add one.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForLog *string `json:"TargetDedicatedHostIdForLog,omitempty" xml:"TargetDedicatedHostIdForLog,omitempty"`
	// The host ID of the primary instance in the dedicated cluster.
	//
	// This parameter is required when you create an ApsaraDB RDS instance in a dedicated cluster. If you do not specify this parameter, the system automatically assigns a host.
	//
	// - You can call DescribeDedicatedHosts to query the host information in the host group.
	//
	// - If you have not added a host, call CreateDedicatedHost to add one.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// The host ID of the secondary instance in the dedicated cluster.
	//
	// This parameter is required when you create an ApsaraDB RDS High-availability Edition or RDS Enterprise Edition instance in a dedicated cluster. If you do not specify this parameter, the system automatically allocates a host by default.
	//
	// - You can call DescribeDedicatedHosts to query the host information in the dedicated cluster.
	//
	// - If you have not added a host, call CreateDedicatedHost to add one.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForSlave *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	// The minor engine version of the RDS instance to create. This parameter is required only when you create a MySQL or PostgreSQL instance.
	//
	// Format:
	//
	// 	- MySQL: `<instance version>_<numeric version number>`. For example, `rds_20200229`, `xcluster_20200229`, or `xcluster80_20200229`. The prefixes are described as follows:
	//
	//     	- rds: high availability series or Basic Edition.
	//
	//     	- xcluster: MySQL 5.7 RDS Enterprise Edition.
	//
	//     	- xcluster80: MySQL 8.0 RDS Enterprise Edition.
	//
	//     > You can call DescribeDBMiniEngineVersions to query the numeric version number. For differences between versions, see [AliSQL minor version release notes](https://help.aliyun.com/document_detail/96060.html).
	//
	// 	- PostgreSQL: `rds_postgres_<major version>00_<minor version number>`. For example, `rds_postgres_1400_20220830`. The fields are described as follows:
	//
	//     	- 1400: PostgreSQL major version 14.
	//
	//     	- 20220830: AliPG minor engine version. You can call DescribeDBMiniEngineVersions to query the minor version number. For differences between versions, see [PostgreSQL minor version release notes](https://help.aliyun.com/document_detail/126002.html).
	//
	//     > If Babelfish is enabled in **BabelfishConfig**, the minor version format for ApsaraDB RDS for PostgreSQL instances is: `rds_postgres_<major version>00_<AliPG minor version>_babelfish`.
	//
	// example:
	//
	// rds_20200229
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- If **Period*	- is set to **Year**, **UsedTime*	- can be set to **1 to 5**.
	//
	// 	- If **Period*	- is set to **Month**, **UsedTime*	- can be set to **1 to 11**.
	//
	// > This parameter is required if the billing method is **Prepaid**.
	//
	// example:
	//
	// 2
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The user backup ID. You can call ListUserBackupFiles to query the ID. Specifying this parameter creates an instance from a user backup.
	//
	// The following restrictions apply when you specify this parameter:
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
	// >This parameter takes effect only when **InstanceNetworkType*	- is set to **VPC**, which indicates the network type is VPC.
	//
	// example:
	//
	// vpc-****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// - **Zone correspondence**: The zone of the vSwitch must correspond to the zone of the primary node (ZoneId) and the zone of the secondary node (ZoneIdSlave1). If you specify two vSwitch IDs, their order must match the order of ZoneId and ZoneSlaveId1.
	//
	// - **Network type requirement**: **InstanceNetworkType*	- must be set to **VPC**.
	//
	// - **Multiple vSwitch requirement**: If you specify **ZoneSlaveId1*	- (the zone ID of the secondary node) and it is not set to **Auto**, you must specify two vSwitch IDs separated by a comma (,).
	//
	// - **Character restriction**: VSwitchId cannot contain special characters such as spaces, `!`, `#`, `￥`, `&`, or `%`.
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The whitelist. If you need to configure multiple IP addresses, separate them with commas (,) without spaces before or after the commas. Example: `192.168.0.1,172.16.213.9`.
	//
	// example:
	//
	// 192.168.0.1,172.16.213.9
	WhitelistTemplateList *string `json:"WhitelistTemplateList,omitempty" xml:"WhitelistTemplateList,omitempty"`
	// The zone ID of the primary node.
	//
	// - If you specify a VPC and a vSwitch, you must set this parameter to the zone ID of the vSwitch. Otherwise, the instance cannot be created.
	//
	// - For high availability series instances, you must also specify **ZoneIdSlave1*	- to determine whether the instance uses single-zone or multi-zone deployment.
	//
	// - For RDS Enterprise Edition instances, you must also specify **ZoneIdSlave1*	- and **ZoneIdSlave2*	- to determine whether the instance uses single-zone or multi-zone deployment.
	//
	// - For RDS Cluster Edition instances, two-node clusters require **ZoneIdSlave1**, and three-node clusters require both **ZoneIdSlave1*	- and **ZoneIdSlave2**.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary node.
	//
	// - If you set this parameter to **Auto**, the instance uses multi-zone deployment and the system automatically selects a zone for the secondary node.
	//
	// - If this parameter is the same as **ZoneId**, the instance uses single-zone deployment.
	//
	// - If this parameter is different from **ZoneId**, the instance uses multi-zone deployment.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The zone ID of the second secondary node. ApsaraDB RDS for MySQL Cluster Edition instances support creating one or two secondary nodes when you create the instance. If you need this, use this parameter to specify the zone of the second secondary node.
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
	// The tag key. Specifying this parameter binds a tag to the instance.
	//
	// 	- If the specified tag key already exists, the tag is directly bound to the instance. You can call ListTagResources to query existing tags.
	//
	// 	- If the specified tag key does not exist, the tag key is created and then bound to the instance.
	//
	// 	- Empty strings are not allowed.
	//
	// 	- This parameter must be used together with **Tag.Value**.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value corresponding to the tag key. Specifying this parameter binds a tag to the instance.
	//
	// 	- If the specified tag value already exists under the corresponding tag key, the tag value is directly bound to the instance. You can call ListTagResources to query existing tags.
	//
	// 	- If the specified tag value does not exist under the corresponding tag key, the tag value is created and then bound to the instance.
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
