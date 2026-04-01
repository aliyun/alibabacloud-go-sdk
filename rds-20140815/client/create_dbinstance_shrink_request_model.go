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
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
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
	Tag []*CreateDBInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
