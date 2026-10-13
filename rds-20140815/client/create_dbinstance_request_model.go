// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateDBInstanceRequest
	GetAmount() *int32
	SetAutoCreateProxy(v bool) *CreateDBInstanceRequest
	GetAutoCreateProxy() *bool
	SetAutoPay(v bool) *CreateDBInstanceRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *CreateDBInstanceRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *CreateDBInstanceRequest
	GetAutoUseCoupon() *bool
	SetBabelfishConfig(v string) *CreateDBInstanceRequest
	GetBabelfishConfig() *string
	SetBpeEnabled(v string) *CreateDBInstanceRequest
	GetBpeEnabled() *string
	SetBurstingEnabled(v bool) *CreateDBInstanceRequest
	GetBurstingEnabled() *bool
	SetBusinessInfo(v string) *CreateDBInstanceRequest
	GetBusinessInfo() *string
	SetCategory(v string) *CreateDBInstanceRequest
	GetCategory() *string
	SetClientToken(v string) *CreateDBInstanceRequest
	GetClientToken() *string
	SetColdDataEnabled(v bool) *CreateDBInstanceRequest
	GetColdDataEnabled() *bool
	SetConnectionMode(v string) *CreateDBInstanceRequest
	GetConnectionMode() *string
	SetConnectionString(v string) *CreateDBInstanceRequest
	GetConnectionString() *string
	SetCreateStrategy(v string) *CreateDBInstanceRequest
	GetCreateStrategy() *string
	SetCustomExtraInfo(v string) *CreateDBInstanceRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *CreateDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceRequest
	GetDBInstanceDescription() *string
	SetDBInstanceNetType(v string) *CreateDBInstanceRequest
	GetDBInstanceNetType() *string
	SetDBInstanceStorage(v int32) *CreateDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *CreateDBInstanceRequest
	GetDBInstanceStorageType() *string
	SetDBIsIgnoreCase(v string) *CreateDBInstanceRequest
	GetDBIsIgnoreCase() *string
	SetDBParamGroupId(v string) *CreateDBInstanceRequest
	GetDBParamGroupId() *string
	SetDBTimeZone(v string) *CreateDBInstanceRequest
	GetDBTimeZone() *string
	SetDedicatedHostGroupId(v string) *CreateDBInstanceRequest
	GetDedicatedHostGroupId() *string
	SetDeletionProtection(v bool) *CreateDBInstanceRequest
	GetDeletionProtection() *bool
	SetDryRun(v bool) *CreateDBInstanceRequest
	GetDryRun() *bool
	SetEncryptionKey(v string) *CreateDBInstanceRequest
	GetEncryptionKey() *string
	SetEngine(v string) *CreateDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBInstanceRequest
	GetEngineVersion() *string
	SetExternalReplication(v bool) *CreateDBInstanceRequest
	GetExternalReplication() *bool
	SetInstanceNetworkType(v string) *CreateDBInstanceRequest
	GetInstanceNetworkType() *string
	SetIoAccelerationEnabled(v string) *CreateDBInstanceRequest
	GetIoAccelerationEnabled() *string
	SetOptimizedWrites(v string) *CreateDBInstanceRequest
	GetOptimizedWrites() *string
	SetPayType(v string) *CreateDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceRequest
	GetPeriod() *string
	SetPort(v string) *CreateDBInstanceRequest
	GetPort() *string
	SetPrivateIpAddress(v string) *CreateDBInstanceRequest
	GetPrivateIpAddress() *string
	SetPromotionCode(v string) *CreateDBInstanceRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRoleARN(v string) *CreateDBInstanceRequest
	GetRoleARN() *string
	SetSecurityIPList(v string) *CreateDBInstanceRequest
	GetSecurityIPList() *string
	SetServerlessConfig(v *CreateDBInstanceRequestServerlessConfig) *CreateDBInstanceRequest
	GetServerlessConfig() *CreateDBInstanceRequestServerlessConfig
	SetStorageAutoScale(v string) *CreateDBInstanceRequest
	GetStorageAutoScale() *string
	SetStorageThreshold(v int32) *CreateDBInstanceRequest
	GetStorageThreshold() *int32
	SetStorageUpperBound(v int32) *CreateDBInstanceRequest
	GetStorageUpperBound() *int32
	SetSystemDBCharset(v string) *CreateDBInstanceRequest
	GetSystemDBCharset() *string
	SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest
	GetTag() []*CreateDBInstanceRequestTag
	SetTargetDedicatedHostIdForLog(v string) *CreateDBInstanceRequest
	GetTargetDedicatedHostIdForLog() *string
	SetTargetDedicatedHostIdForMaster(v string) *CreateDBInstanceRequest
	GetTargetDedicatedHostIdForMaster() *string
	SetTargetDedicatedHostIdForSlave(v string) *CreateDBInstanceRequest
	GetTargetDedicatedHostIdForSlave() *string
	SetTargetMinorVersion(v string) *CreateDBInstanceRequest
	GetTargetMinorVersion() *string
	SetUsedTime(v string) *CreateDBInstanceRequest
	GetUsedTime() *string
	SetUserBackupId(v string) *CreateDBInstanceRequest
	GetUserBackupId() *string
	SetVPCId(v string) *CreateDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBInstanceRequest
	GetVSwitchId() *string
	SetWhitelistTemplateList(v string) *CreateDBInstanceRequest
	GetWhitelistTemplateList() *string
	SetZoneId(v string) *CreateDBInstanceRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *CreateDBInstanceRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *CreateDBInstanceRequest
	GetZoneIdSlave2() *string
}

type CreateDBInstanceRequest struct {
	// The number of ApsaraDB RDS for MySQL instances to create. This parameter is applicable only when you create multiple ApsaraDB RDS for MySQL instances at a time.
	//
	// Valid values: **1*	- to **20**. Default value: **1**.
	//
	// > - When you create multiple ApsaraDB RDS for MySQL instances, you can use the **Tag.Key*	- and **Tag.Value*	- parameters to add tags to the instances in the same batch. This facilitates instance management by tag after creation.
	//
	// > - After multiple instances are created, the API operation returns only the **TaskId**, **RequestId**, and **Message*	- parameters. Other detail parameters are not returned. To query the details of a specific instance, call the DescribeDBInstanceAttribute operation.
	//
	// > - If **Engine*	- is not set to **MySQL*	- and the value of this parameter is greater than **1**, the API operation fails and returns the error code `InvalidParam.Engine`.
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
	// - **true**: enables automatic payment. Make sure that your account balance is sufficient.
	//
	// - **false**: generates an order without deducting fees.
	//
	// > The default value is true. If your payment method has an insufficient balance, you can set AutoPay to false. In this case, an unpaid order is generated, and you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
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
	// Specifies whether to use a voucher. Valid values:
	//
	// 	- **true**: uses a voucher.
	//
	// 	- **false**: does not use a voucher. This is the default value.
	//
	// > If you use a voucher and later need to downgrade the instance configuration, the amount deducted by the voucher will not be refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The configuration information for Babelfish for ApsaraDB RDS for PostgreSQL.
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
	// Specifies whether to enable the I/O performance burst feature for Premium ESSD cloud disks. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
	//
	// > For more information about the I/O performance burst feature for Premium ESSD cloud disks, see [What is a Premium ESSD cloud disk?](https://help.aliyun.com/document_detail/2340501.html).
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
	//     > This parameter is required when you create a SQL Server Enterprise Edition Cluster<props="china">, Basic Edition Standard, or Basic Edition Enterprise instance. For example, when you create a Basic Edition 2022 Enterprise Edition Cluster (2022_ent) instance, you must set this parameter to Basic.
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
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the [data archiving](https://help.aliyun.com/document_detail/2701832.html) feature for Premium ESSD cloud disks. Valid values:
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
	// 	- **Atomicity**: atomicity. All instances in the same batch are created successfully, or none of them are created if one fails. This is the default value.
	//
	// 	- **Partial**: non-atomicity. The creation of an instance is not affected by other instances in the same batch.
	//
	// example:
	//
	// Atomicity
	CreateStrategy  *string `json:"CreateStrategy,omitempty" xml:"CreateStrategy,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type. You can specify a standard or Yitian instance type. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
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
	// The name of the instance. The name must be 2 to 255 characters in length. It must start with a letter and can contain digits, letters, and hyphens (-).
	//
	// > The name cannot start with http:// or https://.
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
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: high-performance local SSD (recommended).
	//
	// 	- **general_essd**: Premium ESSD cloud disk (recommended).
	//
	// 	- **cloud_essd**: PL1 ESSD.
	//
	// 	- **cloud_essd2**: PL2 ESSD.
	//
	// 	- **cloud_essd3**: PL3 ESSD.
	//
	// 	- **cloud_ssd**: standard SSD (not recommended, discontinued in some regions).
	//
	// The default value of this parameter is automatically determined based on the instance type specified by **DBInstanceClass**:
	//
	// 	- If the instance type is a high-performance local SSD type, the default value is **local_ssd**.
	//
	// 	- If the instance type is a cloud disk type, the default value is **cloud_essd**.
	//
	// > Serverless instances support only PL1 ESSDs and Premium ESSD cloud disks.
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
	//     - This parameter specifies the UTC time zone. Valid values: **-12:59*	- to **+13:00**.
	//
	//     - For an instance with Premium Local SSDs, you can specify a named time zone, such as Asia/Hong_Kong. For more information about named time zones, see [Named time zones](https://help.aliyun.com/document_detail/297356.html).
	//
	// - If **Engine*	- is set to **PostgreSQL**:
	//
	//     - This parameter specifies a named time zone. UTC time zones are not supported. For more information about named time zones, see [Named time zones](https://help.aliyun.com/document_detail/297356.html).
	//
	//     - This parameter can be configured only for ApsaraDB RDS for PostgreSQL instances with cloud disks.
	//
	// > - You can set the time zone when you purchase a primary instance. Read-only instances do not support parameter settings for the time zone and inherit the time zone of the primary instance.
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
	// Specifies whether to enable release protection for the instance. This feature is supported only for pay-as-you-go instances. Valid values:
	//
	// 	- **true**: enables release protection.
	//
	// 	- **false**: disables release protection. This is the default value.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to perform a precheck before creating the instance. Valid values:
	//
	// 	- **true**: performs a precheck but does not create the instance. The system checks the request parameters, request format, service limits, and inventory.
	//
	// 	- **false**: sends a normal request and creates the instance after the request passes the precheck. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the key used for cloud disk encryption in the same region. If you specify this parameter, cloud disk encryption is enabled and cannot be disabled. You must also specify **RoleARN**.
	//
	// You can view the key ID or create a new key in the Key Management Service (KMS) console. For more information, see [Create a key](https://help.aliyun.com/document_detail/181610.html).
	//
	// > - For ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and ApsaraDB RDS for SQL Server, you do not need to specify this parameter. You only need to specify **RoleARN*	- to use the service key to create an instance with cloud disk encryption enabled.
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
	//     	- SQL Server: **08r2_ent_ha*	- (cloud disk edition, discontinued), **2008r2*	- (high-performance local SSD, discontinued), **2012*	- (Enterprise Edition Standalone), **2012_ent_ha**, **2012_std_ha**, **2012_web**, **2014_ent_ha**, **2014_std_ha**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_ent**, **2017_std_ha**, **2017_web**, **2019_ent**, **2019_std_ha**, **2019_web**, **2022_ent**, **2022_std_ha**, **2022_web**, **2025_ent**, and **2025_std**
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
	// > - For SQL Server instances, `_ent` indicates the Enterprise Edition Cluster, `_ent_ha` indicates the Enterprise Edition, `_std_ha` indicates the Standard Edition, and `_web` indicates the Web Edition.
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
	// - **ON**: enables native replication.
	//
	// - **OFF**: disables native replication.
	//
	// example:
	//
	// ON
	ExternalReplication *bool `json:"ExternalReplication,omitempty" xml:"ExternalReplication,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: Virtual Private Cloud (VPC).
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
	// Specifies whether to enable the [Buffer Pool Extension (BPE)](https://help.aliyun.com/document_detail/2527067.html) feature for Premium ESSD cloud disks. Valid values:
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
	// - **Serverless**: Serverless billing method. MariaDB instances are not supported. For more information, see [Overview of ApsaraDB RDS for MySQL Serverless instances](https://help.aliyun.com/document_detail/411291.html), [Overview of ApsaraDB RDS for SQL Server Serverless instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of ApsaraDB RDS for PostgreSQL Serverless instances](https://help.aliyun.com/document_detail/607742.html).
	//
	// > The system automatically generates an order and completes the payment. You do not need to manually confirm the payment.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription period unit. Valid values:
	//
	// 	- **Year**: year.
	//
	// 	- **Month**: month.
	//
	// > This parameter is required when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The port used to initialize the instance upon creation. Valid values:
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The global resource descriptor (ARN) used for the authorization of the ApsaraDB RDS service account by the primary account to access KMS. You can call the CheckCloudResourceAuthorized operation to query the descriptor.
	//
	// 	Notice: You must specify **RoleARN*	- when you enable cloud disk encryption.
	//
	// example:
	//
	// acs:ram::1406****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The [IP address whitelist](https://help.aliyun.com/document_detail/43185.html) of the instance. Separate multiple IP addresses or CIDR blocks with commas (,). Each IP address or CIDR block must be unique. You can add up to 1,000 IP addresses or CIDR blocks to the whitelist of a single instance. The following formats are supported:
	//
	// 	- IP addresses, such as 10.10.XX.XX.
	//
	// 	- CIDR blocks, such as 10.10.XX.XX/24. In this example, 24 indicates that the prefix of the CIDR block is 24 bits long. The prefix length ranges from 1 to 32 bits.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.XX.XX/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The settings for the Serverless instance. This parameter is required when you create a Serverless instance.
	//
	// > MariaDB does not support Serverless instances.
	ServerlessConfig *CreateDBInstanceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// Specifies whether to enable automatic storage expansion. This parameter is supported only by MySQL and PostgreSQL. Valid values:
	//
	// 	- **Enable**: enables automatic storage expansion.
	//
	// 	- **Disable**: disables automatic storage expansion. This is the default value.
	//
	// > You can also call the ModifyDasInstanceConfig operation to adjust this setting after the instance is created. For more information, see [Configure automatic storage expansion](https://help.aliyun.com/document_detail/173826.html).
	//
	// example:
	//
	// Disable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The threshold that triggers automatic storage expansion. Unit: %. Valid values:
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
	// The maximum storage capacity after automatic storage expansion. Automatic storage expansion does not increase the total storage capacity of the instance beyond this value. Unit: GB.
	//
	// > - The value must be greater than or equal to 0.
	//
	// > - This parameter is required when **StorageAutoScale*	- is set to **Enable**.
	//
	// example:
	//
	// 2000
	StorageUpperBound *int32 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and does not need to be configured.
	//
	// example:
	//
	// gbk
	SystemDBCharset *string `json:"SystemDBCharset,omitempty" xml:"SystemDBCharset,omitempty"`
	// The list of tags.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the host for the logger node in the dedicated cluster.
	//
	// This parameter is required when you create an RDS Enterprise Edition instance in a dedicated cluster. If you do not specify this parameter, the system automatically allocates a host by default.
	//
	// - You can call the DescribeDedicatedHosts operation to query the host information in the dedicated cluster.
	//
	// - If you do not have a host, you can call the CreateDedicatedHost operation to add one.
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
	// - You can call the DescribeDedicatedHosts operation to query the host information in the dedicated cluster.
	//
	// - If you do not have a host, you can call the CreateDedicatedHost operation to add one.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForSlave *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	// The minor engine version of the RDS instance to be created. This parameter is required only when you create a MySQL or PostgreSQL instance.
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
	//     > You can call the DescribeDBMiniEngineVersions operation to query the numeric version number. For more information about the differences between versions, see [AliSQL release notes](https://help.aliyun.com/document_detail/96060.html).
	//
	// 	- PostgreSQL: `rds_postgres_<Major version>00_<Minor engine version number>`. Example: `rds_postgres_1400_20220830`. Description:
	//
	//     	- 1400: PostgreSQL major version 14.
	//
	//     	- 20220830: AliPG minor engine version. You can call the DescribeDBMiniEngineVersions operation to query the minor engine version. For more information about the differences between versions, see [PostgreSQL release notes](https://help.aliyun.com/document_detail/126002.html).
	//
	//     > If Babelfish is enabled in **BabelfishConfig**, the minor engine version format for ApsaraDB RDS for PostgreSQL is: `rds_postgres_<Major version>00_<AliPG minor engine version>_babelfish`.
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
	// To specify this parameter, the following limits apply:
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
	// The ID of the VPC.
	//
	// > This parameter takes effect only when the network type **InstanceNetworkType*	- is set to **VPC**.
	//
	// example:
	//
	// vpc-****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch.
	//
	// - **Zone mapping**: The zone of the vSwitch must match the zone of the primary node (**ZoneId**) and the zone of the secondary node (**ZoneIdSlave1**). If you specify two vSwitch IDs, their order must match the order of **ZoneId*	- and **ZoneIdSlave1**.
	//
	// - **Network type limit**: **InstanceNetworkType*	- must be set to **VPC**.
	//
	// - **Multiple vSwitches**: If you specify **ZoneIdSlave1*	- and the value is not **Auto**, you must specify two vSwitch IDs and separate them with a comma (,).
	//
	// - **Character limit**: **VSwitchId*	- cannot contain special characters such as spaces, exclamation marks (!), number signs (#), currency symbols (¥), ampersands (&), and percent signs (%).
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The list of whitelists.
	//
	// To configure multiple IP addresses, separate them with commas (,). Do not add spaces before or after the commas. Example: `192.168.0.1,172.16.213.9`.
	//
	// example:
	//
	// 192.168.0.1,172.16.213.9
	WhitelistTemplateList *string `json:"WhitelistTemplateList,omitempty" xml:"WhitelistTemplateList,omitempty"`
	// The zone ID of the primary node.
	//
	// - If you specify a VPC and a vSwitch, you must specify the zone ID of the vSwitch. Otherwise, the instance cannot be created.
	//
	// - For a high-availability series instance, you must also specify **ZoneIdSlave1*	- to determine whether the instance uses a single-zone deployment or multi-zone deployment.
	//
	// - For an RDS Enterprise Edition instance, you must also specify **ZoneIdSlave1*	- and **ZoneIdSlave2*	- to determine whether the instance uses a single-zone deployment or multi-zone deployment.
	//
	// - For an ApsaraDB RDS Cluster Edition instance, you must also specify **ZoneIdSlave1*	- for a two-node instance, and both **ZoneIdSlave1*	- and **ZoneIdSlave2*	- for a three-node instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary node.
	//
	// - If you set this parameter to **Auto**, the instance uses a multi-zone deployment and the system automatically selects a zone for the secondary node.
	//
	// - If you set this parameter to the same value as **ZoneId**, the instance uses a single-zone deployment.
	//
	// - If you set this parameter to a value different from **ZoneId**, the instance uses a multi-zone deployment.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The zone ID of the second secondary node. ApsaraDB RDS for MySQL Cluster Edition instances support the creation of one or two secondary nodes when you create an instance. You can use this parameter to specify the zone of the second secondary node.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneIdSlave2 *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateDBInstanceRequest) GetAutoCreateProxy() *bool {
	return s.AutoCreateProxy
}

func (s *CreateDBInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateDBInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateDBInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateDBInstanceRequest) GetBabelfishConfig() *string {
	return s.BabelfishConfig
}

func (s *CreateDBInstanceRequest) GetBpeEnabled() *string {
	return s.BpeEnabled
}

func (s *CreateDBInstanceRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateDBInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateDBInstanceRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *CreateDBInstanceRequest) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *CreateDBInstanceRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDBInstanceRequest) GetCreateStrategy() *string {
	return s.CreateStrategy
}

func (s *CreateDBInstanceRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *CreateDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *CreateDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CreateDBInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CreateDBInstanceRequest) GetDBIsIgnoreCase() *string {
	return s.DBIsIgnoreCase
}

func (s *CreateDBInstanceRequest) GetDBParamGroupId() *string {
	return s.DBParamGroupId
}

func (s *CreateDBInstanceRequest) GetDBTimeZone() *string {
	return s.DBTimeZone
}

func (s *CreateDBInstanceRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *CreateDBInstanceRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateDBInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDBInstanceRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateDBInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceRequest) GetExternalReplication() *bool {
	return s.ExternalReplication
}

func (s *CreateDBInstanceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CreateDBInstanceRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *CreateDBInstanceRequest) GetOptimizedWrites() *string {
	return s.OptimizedWrites
}

func (s *CreateDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceRequest) GetPort() *string {
	return s.Port
}

func (s *CreateDBInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateDBInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceRequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *CreateDBInstanceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDBInstanceRequest) GetServerlessConfig() *CreateDBInstanceRequestServerlessConfig {
	return s.ServerlessConfig
}

func (s *CreateDBInstanceRequest) GetStorageAutoScale() *string {
	return s.StorageAutoScale
}

func (s *CreateDBInstanceRequest) GetStorageThreshold() *int32 {
	return s.StorageThreshold
}

func (s *CreateDBInstanceRequest) GetStorageUpperBound() *int32 {
	return s.StorageUpperBound
}

func (s *CreateDBInstanceRequest) GetSystemDBCharset() *string {
	return s.SystemDBCharset
}

func (s *CreateDBInstanceRequest) GetTag() []*CreateDBInstanceRequestTag {
	return s.Tag
}

func (s *CreateDBInstanceRequest) GetTargetDedicatedHostIdForLog() *string {
	return s.TargetDedicatedHostIdForLog
}

func (s *CreateDBInstanceRequest) GetTargetDedicatedHostIdForMaster() *string {
	return s.TargetDedicatedHostIdForMaster
}

func (s *CreateDBInstanceRequest) GetTargetDedicatedHostIdForSlave() *string {
	return s.TargetDedicatedHostIdForSlave
}

func (s *CreateDBInstanceRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *CreateDBInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDBInstanceRequest) GetUserBackupId() *string {
	return s.UserBackupId
}

func (s *CreateDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceRequest) GetWhitelistTemplateList() *string {
	return s.WhitelistTemplateList
}

func (s *CreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *CreateDBInstanceRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *CreateDBInstanceRequest) SetAmount(v int32) *CreateDBInstanceRequest {
	s.Amount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetAutoCreateProxy(v bool) *CreateDBInstanceRequest {
	s.AutoCreateProxy = &v
	return s
}

func (s *CreateDBInstanceRequest) SetAutoPay(v bool) *CreateDBInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v string) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetAutoUseCoupon(v bool) *CreateDBInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBabelfishConfig(v string) *CreateDBInstanceRequest {
	s.BabelfishConfig = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBpeEnabled(v string) *CreateDBInstanceRequest {
	s.BpeEnabled = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBurstingEnabled(v bool) *CreateDBInstanceRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBusinessInfo(v string) *CreateDBInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCategory(v string) *CreateDBInstanceRequest {
	s.Category = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetColdDataEnabled(v bool) *CreateDBInstanceRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *CreateDBInstanceRequest) SetConnectionMode(v string) *CreateDBInstanceRequest {
	s.ConnectionMode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetConnectionString(v string) *CreateDBInstanceRequest {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCreateStrategy(v string) *CreateDBInstanceRequest {
	s.CreateStrategy = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCustomExtraInfo(v string) *CreateDBInstanceRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceClass(v string) *CreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceNetType(v string) *CreateDBInstanceRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceStorage(v int32) *CreateDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceStorageType(v string) *CreateDBInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBIsIgnoreCase(v string) *CreateDBInstanceRequest {
	s.DBIsIgnoreCase = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBParamGroupId(v string) *CreateDBInstanceRequest {
	s.DBParamGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBTimeZone(v string) *CreateDBInstanceRequest {
	s.DBTimeZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDedicatedHostGroupId(v string) *CreateDBInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDeletionProtection(v bool) *CreateDBInstanceRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDryRun(v bool) *CreateDBInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionKey(v string) *CreateDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngine(v string) *CreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetExternalReplication(v bool) *CreateDBInstanceRequest {
	s.ExternalReplication = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceNetworkType(v string) *CreateDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetIoAccelerationEnabled(v string) *CreateDBInstanceRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOptimizedWrites(v string) *CreateDBInstanceRequest {
	s.OptimizedWrites = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPort(v string) *CreateDBInstanceRequest {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrivateIpAddress(v string) *CreateDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPromotionCode(v string) *CreateDBInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerId(v int64) *CreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRoleARN(v string) *CreateDBInstanceRequest {
	s.RoleARN = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetServerlessConfig(v *CreateDBInstanceRequestServerlessConfig) *CreateDBInstanceRequest {
	s.ServerlessConfig = v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageAutoScale(v string) *CreateDBInstanceRequest {
	s.StorageAutoScale = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageThreshold(v int32) *CreateDBInstanceRequest {
	s.StorageThreshold = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageUpperBound(v int32) *CreateDBInstanceRequest {
	s.StorageUpperBound = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSystemDBCharset(v string) *CreateDBInstanceRequest {
	s.SystemDBCharset = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateDBInstanceRequest) SetTargetDedicatedHostIdForLog(v string) *CreateDBInstanceRequest {
	s.TargetDedicatedHostIdForLog = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTargetDedicatedHostIdForMaster(v string) *CreateDBInstanceRequest {
	s.TargetDedicatedHostIdForMaster = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTargetDedicatedHostIdForSlave(v string) *CreateDBInstanceRequest {
	s.TargetDedicatedHostIdForSlave = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTargetMinorVersion(v string) *CreateDBInstanceRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v string) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUserBackupId(v string) *CreateDBInstanceRequest {
	s.UserBackupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetWhitelistTemplateList(v string) *CreateDBInstanceRequest {
	s.WhitelistTemplateList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneIdSlave1(v string) *CreateDBInstanceRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneIdSlave2(v string) *CreateDBInstanceRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *CreateDBInstanceRequest) Validate() error {
	if s.ServerlessConfig != nil {
		if err := s.ServerlessConfig.Validate(); err != nil {
			return err
		}
	}
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

type CreateDBInstanceRequestServerlessConfig struct {
	// Specifies whether to enable the smart pause and resume feature for the Serverless instance. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature. This is the default value.
	//
	// > This parameter is required only for MySQL and PostgreSQL Serverless instances. If no connections are established within 10 minutes, the instance enters the paused state and automatically resumes when a new connection is established.
	//
	// example:
	//
	// true
	AutoPause *bool `json:"AutoPause,omitempty" xml:"AutoPause,omitempty"`
	// The maximum number of RDS capacity units (RCUs) for automatic scaling. Valid values:
	//
	// - MySQL: **1*	- to **32**
	//
	// - SQL Server: **2*	- to **16**
	//
	// - PostgreSQL: **1*	- to **14**
	//
	// > The value of this parameter must be greater than or equal to the value of **MinCapacity**. Only **integers*	- are supported.
	//
	// example:
	//
	// 8
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The minimum number of RCUs for automatic scaling. Valid values:
	//
	// - MySQL: **0.5*	- to **32**
	//
	// - SQL Server: **2*	- to **16*	- (integers only)
	//
	// - PostgreSQL: **0.5*	- to **14**
	//
	// > The value of this parameter must be less than or equal to the value of **MaxCapacity**.
	//
	// example:
	//
	// 0.5
	MinCapacity *float64 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
	// Specifies whether to enable forced elastic scaling for the Serverless instance. Valid values:
	//
	// 	- **true**: enables forced elastic scaling.
	//
	// 	- **false**: disables forced elastic scaling. This is the default value.
	//
	// > 	- This parameter is required only for MySQL and PostgreSQL Serverless instances. If you enable this parameter, the instance will be unavailable for 30 to 120 seconds during forced scaling. Use this parameter with caution based on your business requirements.
	//
	// > 	- Elastic scaling of RCUs usually takes effect immediately. However, in some special cases (such as during the execution of a large transaction), scaling cannot be completed immediately. In this case, you can enable this parameter to force scaling.
	//
	// example:
	//
	// false
	SwitchForce *bool `json:"SwitchForce,omitempty" xml:"SwitchForce,omitempty"`
}

func (s CreateDBInstanceRequestServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestServerlessConfig) GetAutoPause() *bool {
	return s.AutoPause
}

func (s *CreateDBInstanceRequestServerlessConfig) GetMaxCapacity() *float64 {
	return s.MaxCapacity
}

func (s *CreateDBInstanceRequestServerlessConfig) GetMinCapacity() *float64 {
	return s.MinCapacity
}

func (s *CreateDBInstanceRequestServerlessConfig) GetSwitchForce() *bool {
	return s.SwitchForce
}

func (s *CreateDBInstanceRequestServerlessConfig) SetAutoPause(v bool) *CreateDBInstanceRequestServerlessConfig {
	s.AutoPause = &v
	return s
}

func (s *CreateDBInstanceRequestServerlessConfig) SetMaxCapacity(v float64) *CreateDBInstanceRequestServerlessConfig {
	s.MaxCapacity = &v
	return s
}

func (s *CreateDBInstanceRequestServerlessConfig) SetMinCapacity(v float64) *CreateDBInstanceRequestServerlessConfig {
	s.MinCapacity = &v
	return s
}

func (s *CreateDBInstanceRequestServerlessConfig) SetSwitchForce(v bool) *CreateDBInstanceRequestServerlessConfig {
	s.SwitchForce = &v
	return s
}

func (s *CreateDBInstanceRequestServerlessConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDBInstanceRequestTag struct {
	// The tag key. Specify this parameter to add a tag to the instance.
	//
	// 	- If the specified tag key already exists, the tag key is directly added to the instance. You can call the ListTagResources operation to query existing tags.
	//
	// 	- If the specified tag key does not exist, the system creates the tag key and adds it to the instance.
	//
	// 	- Empty strings are not allowed.
	//
	// 	- This parameter must be used together with **Tag.Value**.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key. Specify this parameter to add a tag to the instance.
	//
	// 	- If the specified tag value exists for the corresponding tag key, the tag value is directly added to the instance. You can call the ListTagResources operation to query existing tags.
	//
	// 	- If the specified tag value does not exist for the corresponding tag key, the system creates the tag value and adds it to the instance.
	//
	// 	- This parameter must be used together with **Tag.Key**.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDBInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDBInstanceRequestTag) SetKey(v string) *CreateDBInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBInstanceRequestTag) SetValue(v string) *CreateDBInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDBInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
