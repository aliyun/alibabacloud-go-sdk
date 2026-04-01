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
	// The number of ApsaraDB RDS for MySQL instances that you want to create. The parameter takes effect only when you create multiple ApsaraDB RDS for MySQL instances at a time by using a single request.
	//
	// Valid values: **1*	- to **20**. Default value: **1**.
	//
	// > 	- If you want to create multiple ApsaraDB RDS for MySQL instances at a time by using a single request, you can add tags to all the instances by using the **Tag.Key*	- parameter and the **Tag.Value*	- parameter. After the instances are created, you can manage the instances based on the tags.
	//
	// > 	- After you submit a request to create multiple ApsaraDB RDS for MySQL instances, this operation returns **TaskId**, **RequestId**, and **Message**. You can call the DescribeDBInstanceAttribute operation to query the information about an instance.
	//
	// > 	- If the value of the **Engine*	- parameter is not **MySQL*	- and the value of the Amount parameter is greater than **1**, this operation fails and returns an error code `InvalidParam.Engine`.
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to automatically create a proxy. Valid values:
	//
	// 	- **true**: automatically creates a database proxy. By default, a general-purpose database proxy is created.
	//
	// 	- **false**: does not automatically create a database proxy.
	//
	// example:
	//
	// false
	AutoCreateProxy *bool `json:"AutoCreateProxy,omitempty" xml:"AutoCreateProxy,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: enables the feature. Make sure that your account balance is sufficient when you enable automatic payment.
	//
	// 	- **false**: does not automatically complete the payment. An unpaid order is generated.
	//
	// >  Default value: true. If your account balance is insufficient, you can set AutoPay to false to generate an unpaid order. Then, you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. You must specify this parameter only if the instance uses the subscription billing method. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > 	- The auto-renewal cycle is one month for a monthly subscription.
	//
	// > 	- The auto-renewal cycle is one year for a yearly subscription.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use a coupon. Default value: false. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If you downgrade the specifications of an instance after you use coupons, the used coupons cannot be refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The configuration of the Babelfish feature for the instance that runs PostgreSQL.
	//
	// Format:{"babelfishEnabled":"true","migrationMode":"xxxxxxx","masterUsername":"xxxxxxx","masterUserPassword":"xxxxxxxx"}
	//
	// The following list describes the fields in the format:
	//
	// 	- **babelfishEnabled**: specifies whether to enable Babelfish for the instance. If you set this field to **true**, you enable Babelfish for the instance. If you leave this parameter empty, Babelfish is disabled for the instance.
	//
	// 	- **migrationMode**: The migration mode of the instance. Valid values: **single-db*	- and **multi-db**.
	//
	// 	- **masterUsername**: The username of the administrator account. The username can contain lowercase letters, digits, and underscores (_). It must start with a letter and end with a letter or digit. It can be up to 63 characters in length and cannot start with pg.
	//
	// 	- **masterUserPassword**: The password of the administrator account. The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. It must be 8 to 32 characters in length. The password can contain any of the following characters: `! @ # $ % ^ & 	- ( ) _ + - =`.
	//
	// > This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://help.aliyun.com/document_detail/428613.html).
	//
	// example:
	//
	// {"babelfishEnabled":"true","migrationMode":"single-db","masterUsername":"babelfish_user","masterUserPassword":"Babelfish123!"}
	BabelfishConfig *string `json:"BabelfishConfig,omitempty" xml:"BabelfishConfig,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// false
	BpeEnabled *string `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	// Specifies whether to enable the I/O burst feature of Premium ESSDs. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  For more information about the I/O burst feature of general ESSDs, see [What are Premium ESSDs?](https://help.aliyun.com/document_detail/2340501.html)
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The additional business information about the instance.
	//
	// example:
	//
	// 121436975448952
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The RDS edition of the instance. Valid values:
	//
	// 	- Regular RDS instance
	//
	//     	- **Basic**: RDS Basic Edition
	//
	//     	- **HighAvailability**: RDS High-availability Edition
	//
	//     	- **cluster**: RDS Cluster Edition for ApsaraDB RDS for MySQL or PostgreSQL
	//
	//     	- **AlwaysOn**: RDS Cluster Edition for ApsaraDB RDS for SQL Server
	//
	//     	- **Finance**: RDS Basic Edition for serverless instances
	//
	// 	- Serverless RDS instance
	//
	//     	- **serverless_basic**: RDS Basic Edition for serverless instances. This edition is available only for instances that run MySQL and PostgreSQL.
	//
	//     	- **serverless_standard**: RDS High-availability Edition for serverless instances. This edition is available only for instances that run MySQL and PostgreSQL.
	//
	//     	- **serverless_ha**: RDS High-availability Edition for serverless instances. This edition is available only for instances that run SQL Server.
	//
	// > This parameter is required if PayType is set to Serverless.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the data archiving feature of Premium ESSDs. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  For more information about the data archiving feature of Premium ESSDs, see [Use the data archiving feature](https://help.aliyun.com/document_detail/2701832.html).
	//
	// example:
	//
	// false
	ColdDataEnabled *bool `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	// The connection mode of the instance. Valid values:
	//
	// 	- **Standard**: standard mode
	//
	// 	- **Safe**: database proxy mode
	//
	// ApsaraDB RDS automatically assigns a connection mode to the instance.
	//
	// > SQL Server 2012, SQL Server 2016, and SQL Server 2017 support only the standard mode.
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The internal endpoint that is used to connect to the instance.
	//
	// example:
	//
	// rm-uf6wjk5*****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The policy based on which multiple instances are created. The parameter takes effect only when the value of the **Amount*	- parameter is greater than 1. Valid values:
	//
	// 	- **Atomicity*	- (default): atomicity. The instances are all created together. If one instance cannot be created, none of the instances are created.
	//
	// 	- **Partial**: non-atomicity. Each instance is independently created. The failure in creating an instance does not affect the creation of the other instances.
	//
	// example:
	//
	// Atomicity
	CreateStrategy  *string `json:"CreateStrategy,omitempty" xml:"CreateStrategy,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type of the instance. You can specify an instance type of the standard or YiTian product type. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// To create a serverless instance, set this parameter to one of the following values:
	//
	// 	- If you want to create a serverless instance that runs MySQL on RDS Basic Edition, set this parameter to **mysql.n2.serverless.1c**.
	//
	// 	- If you want to create a serverless instance that runs MySQL on RDS High-availability Edition, set this parameter to **mysql.n2.serverless.2c**.
	//
	// 	- If you want to create a serverless instance that runs SQL Server, set this parameter to **mssql.mem2.serverless.s2**.
	//
	// 	- If you want to create a serverless instance that runs PostgreSQL on RDS Basic Edition, set this parameter to **pg.n2.serverless.1c**.
	//
	// 	- If you want to create a serverless instance that runs PostgreSQL on RDS High-availability Edition, set this parameter to **pg.n2.serverless.2c**.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds.mysql.s1.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance name. The value must be 2 to 255 characters in length The name can contain letters, digits, and hyphens (-) and must start with a letter.
	//
	// >  The value cannot start with http:// or https://.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The network connection type of the instance. The value of this parameter is fixed as **Intranet**, indicating an internal network connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// Internet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The storage capacity of the instance. Unit: GB. The storage capacity increases in increments of 5 GB. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: Premium Local SSD (recommended)
	//
	// 	- **general_essd**: Premium Enterprise SSD (ESSD) (recommend)
	//
	// 	- **cloud_essd**: PL1 ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// 	- **cloud_ssd**: standard SSD. This storage type is not recommended. Standard SSDs are no longer available for purchase in some Alibaba Cloud regions.
	//
	// The default value of this parameter is determined by the instance type specified by the **DBInstanceClass*	- parameter.
	//
	// 	- If the instance type specifies the Premium Local SSD storage type, the default value of this parameter is **local_ssd**.
	//
	// 	- If the instance type specifies the cloud disk storage type, the default value of this parameter is **cloud_essd**.
	//
	// >  Serverless instances support only PL1 ESSDs and Premium ESSDs.
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// Specifies whether the table name is case-sensitive. Valid values:
	//
	// 	- **true**: Table names are not case-sensitive. This is the default value.
	//
	// 	- **false**: Table names are case-sensitive.
	//
	// example:
	//
	// true
	DBIsIgnoreCase *string `json:"DBIsIgnoreCase,omitempty" xml:"DBIsIgnoreCase,omitempty"`
	// The parameter template ID. You can call the DescribeParameterGroups operation to query the parameter template ID.
	//
	// >  This parameter is available if you want to create an instance that runs MySQL or PostgreSQL. If you do not configure this parameter, the default parameter template is used. If you want to use a custom parameter template, you can customize a parameter template and set this parameter to the ID of the custom template.
	//
	// example:
	//
	// rpg-sys-*****
	DBParamGroupId *string `json:"DBParamGroupId,omitempty" xml:"DBParamGroupId,omitempty"`
	// The time zone of the instance. This parameter takes effect only when you set **Engine*	- to **MySQL*	- or **PostgreSQL**.
	//
	// 	- **Engine*	- is set to **MySQL**:
	//
	//     	- This time zone is in UTC. Valid values: \\*\\*-12:59\\*\\	- to **+13:00**.
	//
	//     	- If the instance uses Premium Local SSDs, you can specify the name of the time zone. For example, you can specify the Asia/Hong_Kong time zone. For more information, see [Time zones](https://help.aliyun.com/document_detail/297356.html).
	//
	// 	- **Engine*	- is set to **PostgreSQL**.
	//
	//     	- This time zone is not in UTC. For more information, see [Time zones](https://help.aliyun.com/document_detail/297356.html).
	//
	//     	- You can configure this parameter only when the RDS instance uses cloud disks.
	//
	// > 	- You can specify the time zone when you create a primary instance. You cannot specify the time zone when you create a read-only instance. Read-only instances inherit the time zone of their primary instance.
	//
	// > 	- If you do not specify this parameter, the system automatically assigns the default time zone of the region in which the instance resides.
	//
	// example:
	//
	// +08:00
	DBTimeZone *string `json:"DBTimeZone,omitempty" xml:"DBTimeZone,omitempty"`
	// The ID of the dedicated cluster to which the instance belongs.
	//
	// If you create the instance in a dedicated cluster, you must specify this parameter.
	//
	// 	- You can call the DescribeDedicatedHostGroups operation to query the information about the dedicated cluster.
	//
	// 	- If no dedicated clusters are created, you can call the CreateDedicatedHostGroup operation to create a dedicated cluster.
	//
	// example:
	//
	// dhg-4n*****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to enable the release protection feature for the instance. This feature is available only for pay-as-you-go instances. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false*	- (default): disables the feature.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to perform a dry run. Default value: false. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and insufficient inventory errors.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the key that is used for cloud disk encryption in the region in which the instance is deployed. If this parameter is specified, cloud disk encryption is enabled and you must also specify the **RoleARN*	- parameter. Cloud disk encryption cannot be disabled after it is enabled.
	//
	// You can obtain the ID of the key in the Key Management Service (KMS) console or create a key. For more information, see [Create a key](https://help.aliyun.com/document_detail/181610.html).
	//
	// > 	- This parameter is not required when you create an instance that runs MySQL, PostgreSQL, or SQL Server. You need to only specify the **RoleARN*	- parameter to create an instance that has cloud disk encryption enabled by using the obtained key ID.
	//
	// > 	- You can configure RAM authorization to require a RAM user to enable cloud disk encryption when the RAM user is used to create an instance. If cloud disk encryption is disabled during the instance creation, the creation operation fails. To complete the configuration, you can attach the following policy to the RAM user: `{"Version":"1","Statement":[{"Effect":"Deny","Action":"rds:CreateDBInstance","Resource":"*","Condition":{"StringEquals":{"rds:DiskEncryptionRequired":"false"}}}]}`
	//
	//
	// 	Warning: The configuration also affects the CreateOrder operation that is called to create instances in the console.
	//
	// example:
	//
	// 0d24*****-da7b-4786-b981-9a164dxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The database engine of the instance. Valid values:
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
	// The database engine version of the instance.
	//
	// 	- Regular RDS instance
	//
	//     	- Valid values when you set Engine to MySQL: **5.5**, **5.6**, **5.7**, and **8.0**
	//
	//     	- Valid values when you set Engine to SQLServer: **08r2_ent_ha**(cloud disks, discontinued), **2008r2**(premium local disks, discontinued), **2012**(SQL Server EE Basic), **2012_ent_ha**, **2012_std_ha**, **2012_web**, **2014_ent_ha**, **2014_std_ha**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_ent**, **2017_std_ha**, **2017_web**, **2019_ent**, **2019_std_ha**, **2019_web**, **2022_ent**, **2022_std_ha**, and **2022_web**
	//
	//     	- Valid values when you set Engine to PostgreSQL: **10.0**, **11.0**, **12.0**, **13.0**, **14.0**, **15.0**, **16.0**, and **17.0**
	//
	//     	- Valid values when you set Engine to MariaDB: **10.3*	- and **10.6**
	//
	// 	- Serverless RDS instance
	//
	//     	- Valid values when you set Engine to MySQL: **5.7*	- and **8.0**
	//
	//     	- Valid values when you set Engine to SQLServer: **2016_std_sl**, **2017_std_sl**, and **2019_std_sl**
	//
	//     	- Valid values when you set Engine to PostgreSQL: **14.0**, **15.0**, **16.0**, and **17.0**
	//
	// >
	//
	// 	- ApsaraDB RDS for MariaDB does not support serverless instances.
	//
	// 	- RDS instances that run SQL Server: `_ent` specifies SQL Server EE (Always On), `_ent_ha` specifies SQL Server EE, `_std_ha` specifies SQL Server SE, and `_web` specifies SQL Server Web.
	//
	// 	- RDS instances that run SQL Server 2014 are not available for purchase on the international site (alibabacloud.com).
	//
	// 	- Babelfish is supported only for RDS instances that run PostgreSQL 15.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	EngineVersion       *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExternalReplication *bool   `json:"ExternalReplication,omitempty" xml:"ExternalReplication,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **Classic**: classic network
	//
	// >
	//
	// 	- If the instance runs MySQL and uses cloud disks, you must set this parameter to **VPC**.
	//
	// 	- If the instance runs PostgreSQL or MariaDB, you must set this parameter to **VPC**.
	//
	// 	- If the instance runs SQL Server Basic or SQL Server Web, you can set this parameter to VPC or Classic. If the instance runs other database engines, you must set this parameter to **VPC**.
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// Specifies whether to enable Buffer Pool Extension (BPE) of Premium ESSDs. Valid values:
	//
	// 	- **1**: enables BPE.
	//
	// 	- **0**: disables BPE.
	//
	// >  For more information about Buffer Pool Extension(BPE) of Premium ESSDs, see [Buffer Pool Extension(BPE)](https://help.aliyun.com/document_detail/2527067.html).
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to enable the 16K atomic write feature. Valid values:
	//
	// 	- **optimized**: enables the 16K atomic write feature.
	//
	// 	- **none*	- (default): does not enable the 16K atomic write feature.
	//
	// >  For more information, see [Use the 16K atomic write feature](https://help.aliyun.com/document_detail/2858761.html).
	//
	// example:
	//
	// optimized
	OptimizedWrites *string `json:"OptimizedWrites,omitempty" xml:"OptimizedWrites,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// 	- **Serverless**: serverless. This value is not supported for instances that run MariaDB. For more information, see [Overview of serverless ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/411291.html), [Overview of serverless ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of serverless ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/607742.html).
	//
	// > The system automatically generates a purchase order and completes the payment.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// >  If you set the PayType parameter to **Prepaid**, you must specify this parameter.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The port. You can initialize the port when you create the instance.
	//
	// 	- Valid values if the instance runs MySQL: 1000 to 65534
	//
	// 	- Valid values if the instance runs PostgreSQL, SQL Server, or MariaDB: 1000 to 5999
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the instance. The private IP address must be within the CIDR block that is supported by the specified vSwitch. ApsaraDB RDS automatically assigns a private IP address to the instance based on the values of the **VPCId*	- and **vSwitchId*	- parameters.
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
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) operation to query the most recent region list.
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
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) that is provided by your Alibaba Cloud account for Resource Access Management (RAM) users. RAM users can use the ARN to connect to ApsaraDB RDS to Key Management Service (KMS). You can call the CheckCloudResourceAuthorized operation to query the ARN.
	//
	// >  When you enable the encryption, you must specify the RoleARN.
	//
	// example:
	//
	// acs:ram::1406xxxxxx:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The IP address whitelist of the instance. For more information, see [Configure an IP address whitelist](https://help.aliyun.com/document_detail/43185.html). Separate multiple IP addresses or CIDR blocks with commas (,). You can add up to 1,000 IP addresses or CIDR blocks to the whitelist. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as 10.10.XX.XX.
	//
	// 	- CIDR blocks, such as 10.10.XX.XX/24. In this example, 24 indicates that the prefix of each IP address in the IP address whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.XX.XX/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The settings of the serverless instance. These parameters are required only when you create a serverless instance.
	//
	// >  ApsaraDB RDS for MariaDB does not support serverless instances.
	ServerlessConfig *CreateDBInstanceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// Specifies whether to enable the automatic storage expansion feature for the instance. This feature is supported if the instance runs MySQL or PostgreSQL. Valid values:
	//
	// 	- **Enable**: enables the feature.
	//
	// 	- **Disable*	- (default): disables the feature.
	//
	// >  After the instance is created, you can call the ModifyDasInstanceConfig operation to adjust the settings. For more information, see [Configure automatic storage expansion](https://help.aliyun.com/document_detail/173826.html).
	//
	// example:
	//
	// Disable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The threshold in percentage based on which automatic storage expansion is triggered.
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
	// >  If you set the **StorageAutoScale*	- parameter to **Enable**, you must specify this parameter.
	//
	// example:
	//
	// 50
	StorageThreshold *int32 `json:"StorageThreshold,omitempty" xml:"StorageThreshold,omitempty"`
	// The maximum storage capacity that is allowed for automatic storage expansion. The storage capacity of the instance cannot exceed the maximum storage capacity. Unit: GB.
	//
	// > 	- Valid values: an integer greater than or equal to 0.
	//
	// > 	- If you set **StorageAutoScale*	- to **Enable**, you must specify this parameter.
	//
	// example:
	//
	// 2000
	StorageUpperBound *int32 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// gbk
	SystemDBCharset *string `json:"SystemDBCharset,omitempty" xml:"SystemDBCharset,omitempty"`
	// The tags that are added to instances.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the host to which the logger instance belongs in the specified dedicated cluster.
	//
	// If you want to create an instance that runs RDS Enterprise Edition in a dedicated cluster, you must specify this parameter. If you do not specify this parameter, the system automatically assigns a host.
	//
	// 	- You can call the DescribeDedicatedHosts operation to query the host in the dedicated cluster.
	//
	// 	- If no hosts are created, you can call the CreateDedicatedHost operation to create a host.
	//
	// example:
	//
	// i-bp*****3
	TargetDedicatedHostIdForLog *string `json:"TargetDedicatedHostIdForLog,omitempty" xml:"TargetDedicatedHostIdForLog,omitempty"`
	// The ID of the host to which the instance belongs in the specified dedicated cluster.
	//
	// If you create the instance in a dedicated cluster, you must specify this parameter. If you do not specify this parameter, the system automatically assigns a host.
	//
	// 	- You can call the DescribeDedicatedHosts operation to query the host in the dedicated cluster.
	//
	// 	- If no hosts are created, you can call the CreateDedicatedHost operation to create a host.
	//
	// example:
	//
	// i-bp*****1
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// The ID of the host to which the secondary instance belongs in the specified dedicated cluster.
	//
	// If you want to create an instance that runs RDS High-availability Edition or RDS Enterprise Edition in a dedicated cluster, you must specify this parameter. If you do not specify this parameter, the system automatically assigns a host.
	//
	// 	- You can call the DescribeDedicatedHosts operation to query the host in the dedicated cluster.
	//
	// 	- If no hosts are created, you can call the CreateDedicatedHost operation to create a host.
	//
	// example:
	//
	// i-bp*****2
	TargetDedicatedHostIdForSlave *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	// The minor engine version of the instance. This parameter is required only when you create an instance that runs MySQL or PostgreSQL. The value format varies based on the database engine of the instance.
	//
	// 	- If you create an instance that runs MySQL, the value is in the following format: `<RDS edition>_<Minor engine version>`. Examples: `rds_20200229`, `xcluster_20200229`, and `xcluster80_20200229`.
	//
	//     	- rds: The instance runs RDS Basic Edition or RDS High-availability Edition.
	//
	//     	- xcluster: The instance runs MySQL 5.7 on RDS Enterprise Edition.
	//
	//     	- xcluster80: The instance runs MySQL 8.0 on RDS Enterprise Edition.
	//
	//     > You can call the DescribeDBMiniEngineVersions operation to query the minor engine version. For more information about the differences between minor engine versions of AliSQL, see [Release notes](https://help.aliyun.com/document_detail/96060.html).
	//
	// 	- If you create an instance that runs PostgreSQL, the value is in the following format: `rds_postgres_<Major engine version>00_<Minor engine version>`. Example: `rds_postgres_1400_20220830`.
	//
	//     	- 1400: The major engine version is PostgreSQL 14.
	//
	//     	- 20220830: the AliPG version. You can call the DescribeDBMiniEngineVersions operation to query the AliPG version. For more information about minor engine versions, see [Release notes for AliPG](https://help.aliyun.com/document_detail/126002.html).
	//
	//     > If you configure the **BabelfishConfig*	- parameter for your instance that runs PostgreSQL and set the babelfishEnabled field to true, the value of this parameter is in the following format: `rds_postgres_Major engine version00_AliPG version_babelfish`.
	//
	// example:
	//
	// rds_20200229
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The subscription duration of the instance. Valid values:
	//
	// 	- If you set the **Period*	- parameter to **Year**, the value of the **UsedTime*	- parameter ranges from **1 to 5**.
	//
	// 	- If you set the **Period*	- parameter to **Month**, the value of the **UsedTime*	- parameter ranges from **1 to 11**.
	//
	// >  If you set the PayType parameter to **Prepaid**, you must also specify this parameter.
	//
	// example:
	//
	// 2
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The ID of the full backup file. You can call the ListUserBackupFiles operation to query the ID of the full backup file. If you want to create an instance by using the data of a backup file, you must specify this parameter.
	//
	// This parameter is supported only when the following requirements are met:
	//
	// 	- The **PayType*	- parameter is set to **Postpaid**.
	//
	// 	- The **Engine*	- parameter is set to **MySQL**.
	//
	// 	- The **EngineVersion*	- parameter is set to **5.7**.
	//
	// 	- The **Category*	- parameter is set to **Basic**.
	//
	// example:
	//
	// 67798*****
	UserBackupId *string `json:"UserBackupId,omitempty" xml:"UserBackupId,omitempty"`
	// The ID of the VPC to which the instance belongs.
	//
	// > This parameter is available when you set the **InstanceNetworkType*	- parameter to **VPC**.
	//
	// example:
	//
	// vpc-*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// 	- **Relations with zones**: Specify the vSwitch ID based on the zones in which the vSwitch belongs to. If you specify two vSwitch IDs, make sure that the vSwitch IDs match the zone IDs specified by the ZoneId and ZoneIdSlave1 parameters.
	//
	// 	- **Limits on the network type**: Set **InstanceNetworkType*	- to **VPC**.
	//
	// 	- **Limits on multiple vSwitch IDs**: If you set **ZoneSlaveId1*	- to a value that is not **Auto**, you must specify the IDs of two vSwitches for this parameter and separate the IDs with a comma (,).
	//
	// 	- **Limits on characters**: The value cannot contain `spaces` or the following characters: `!` `#` `￥` `&` `%`
	//
	// example:
	//
	// vsw-*****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The entries in the whitelist. If you enter multiple IP addresses or CIDR blocks, you must separate the IP addresses or CIDR blocks with commas (,). Do not add spaces preceding or following the commas. Example: `192.168.0.1,172.16.213.9`.
	//
	// example:
	//
	// 192.XXX.XX.1,172.XXX.XX.9
	WhitelistTemplateList *string `json:"WhitelistTemplateList,omitempty" xml:"WhitelistTemplateList,omitempty"`
	// The zone ID of the primary instance.
	//
	// 	- If you specify a virtual private cloud (VPC) and a vSwitch, you must specify the ID of the zone to which the specified vSwitch belongs. Otherwise, the instance cannot be created.
	//
	// 	- If the instance runs RDS High-availability Edition, you must specify the **ZoneIdSlave1*	- parameter. The ZoneIdSlave1 parameter specifies whether to use the single-zone deployment method or the multi-zone deployment method.
	//
	// 	- If the instance runs RDS Enterprise Edition, you must specify the **ZoneIdSlave1*	- and **ZoneIdSlave2*	- parameters. The ZoneIdSlave1 and ZoneIdSlave2 parameters specify whether to use the single-zone deployment method or the multi-zone deployment method.
	//
	// 	- If the instance runs MySQL on RDS Cluster Edition, you must specify the **ZoneIdSlave1*	- parameter for the RDS cluster that has two nodes and the **ZoneIdSlave1*	- and **ZoneIdSlave2*	- parameters for the RDS cluster that has three nodes.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary instance.
	//
	// 	- If you set this parameter to **Auto**, the multi-zone deployment method is used and the zone of the secondary instance is automatically configured.
	//
	// 	- If you set this parameter to the same value as the **ZoneId*	- parameter, the single-zone deployment method is used.
	//
	// 	- If you set this parameter to a value that is different from the value of the **ZoneId*	- parameter, the multiple-zone deployment method is used.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The zone ID of the other secondary node. When you create an ApsaraDB RDS for MySQL cluster, you can create one to two secondary nodes for the cluster. This parameter applies if you create a cluster that contains two secondary nodes.
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
	// Specifies whether to enable the automatic startup and stop feature for the serverless instance. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false*	- (default): disables the feature.
	//
	// >  This parameter is required only for serverless instances that run MySQL and PostgreSQL. After the automatic start and stop feature is enabled, if no connections to the instance are established within 10 minutes, the instance is suspended. After a connection to the instance is established, the instance is resumed.
	//
	// example:
	//
	// true
	AutoPause *bool `json:"AutoPause,omitempty" xml:"AutoPause,omitempty"`
	// The maximum number of RDS Capacity Units (RCUs). Valid values:
	//
	// 	- Serverless ApsaraDB RDS for MySQL instances: **1 to 32**
	//
	// 	- Serverless ApsaraDB RDS for SQL Server instances: **2 to 16**
	//
	// 	- Serverless ApsaraDB RDS for PostgreSQL instances: **1 to 14**
	//
	// >  The value of this parameter must be greater than or equal to the value of the **MinCapacity*	- parameter and can be set only to an **integer**.
	//
	// example:
	//
	// 8
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The minimum number of RCUs. Valid values:
	//
	// 	- Serverless ApsaraDB RDS for MySQL instances: **0.5 to 32**.
	//
	// 	- Serverless ApsaraDB RDS for SQL Server instances: **2 to 16**. Only integers are supported.
	//
	// 	- Serverless ApsaraDB RDS for PostgreSQL instances: **0.5 to 14**
	//
	// >  The value of this parameter must be less than or equal to the value of the **MaxCapacity*	- parameter.
	//
	// example:
	//
	// 0.5
	MinCapacity *float64 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
	// Specifies whether to enable the forced scaling feature for the serverless instance. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false*	- (default): disables the feature.
	//
	// >
	//
	// 	- This parameter is required only for serverless instances that run MySQL and PostgreSQL. If you set this parameter to true, a service interruption that lasts approximately 30 to 120 seconds occurs during forced scaling. Process with caution.
	//
	// 	- The RCU scaling for a serverless instance immediately takes effect. In some cases, such as the execution of large transactions, the scaling does not immediately take effect. In this case, you can enable this feature to forcefully scale the RCUs of the instance.
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
	// The tag key. You can use this parameter to add tags to the instance.
	//
	// 	- If the specified tag key is an existing key, the system directly adds the tag key to the instance. You can call the ListTagResources to query the existing tag.
	//
	// 	- If the specified tag key does not exist, the system creates the tag key and adds the tag key to the instance.
	//
	// 	- The value cannot be an empty string.
	//
	// 	- This parameter must be used together with the **Tag.Value*	- parameter.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can use this parameter to add tags to the instance.
	//
	// 	- If the specified tag value is found in the specified tag key, the system directly adds the tag value to the instance. You can call the ListTagResources to query the existing tag.
	//
	// 	- If the specified tag value is not found in the specified tag key, the system creates the tag value and adds the tag value to the instance.
	//
	// 	- This parameter must be used together with the **Tag.Key*	- parameter.
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
