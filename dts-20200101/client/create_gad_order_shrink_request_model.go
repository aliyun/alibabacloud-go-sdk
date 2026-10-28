// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateGadOrderShrinkRequest
	GetAmount() *int32
	SetAutoCreateProxy(v bool) *CreateGadOrderShrinkRequest
	GetAutoCreateProxy() *bool
	SetAutoPay(v bool) *CreateGadOrderShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *CreateGadOrderShrinkRequest
	GetAutoRenew() *string
	SetBabelfishConfig(v string) *CreateGadOrderShrinkRequest
	GetBabelfishConfig() *string
	SetBpeEnabled(v string) *CreateGadOrderShrinkRequest
	GetBpeEnabled() *string
	SetBurstingEnabled(v bool) *CreateGadOrderShrinkRequest
	GetBurstingEnabled() *bool
	SetBusinessInfo(v string) *CreateGadOrderShrinkRequest
	GetBusinessInfo() *string
	SetCategory(v string) *CreateGadOrderShrinkRequest
	GetCategory() *string
	SetClientToken(v string) *CreateGadOrderShrinkRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateGadOrderShrinkRequest
	GetClusterId() *string
	SetColdDataEnabled(v bool) *CreateGadOrderShrinkRequest
	GetColdDataEnabled() *bool
	SetConnectionMode(v string) *CreateGadOrderShrinkRequest
	GetConnectionMode() *string
	SetConnectionString(v string) *CreateGadOrderShrinkRequest
	GetConnectionString() *string
	SetCreateStrategy(v string) *CreateGadOrderShrinkRequest
	GetCreateStrategy() *string
	SetDBInstanceClass(v string) *CreateGadOrderShrinkRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateGadOrderShrinkRequest
	GetDBInstanceDescription() *string
	SetDBInstanceNetType(v string) *CreateGadOrderShrinkRequest
	GetDBInstanceNetType() *string
	SetDBInstanceStorage(v int32) *CreateGadOrderShrinkRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *CreateGadOrderShrinkRequest
	GetDBInstanceStorageType() *string
	SetDBIsIgnoreCase(v string) *CreateGadOrderShrinkRequest
	GetDBIsIgnoreCase() *string
	SetDBParamGroupId(v string) *CreateGadOrderShrinkRequest
	GetDBParamGroupId() *string
	SetDBTimeZone(v string) *CreateGadOrderShrinkRequest
	GetDBTimeZone() *string
	SetDbList(v string) *CreateGadOrderShrinkRequest
	GetDbList() *string
	SetDbListMode(v int32) *CreateGadOrderShrinkRequest
	GetDbListMode() *int32
	SetDedicatedHostGroupId(v string) *CreateGadOrderShrinkRequest
	GetDedicatedHostGroupId() *string
	SetDeletionProtection(v bool) *CreateGadOrderShrinkRequest
	GetDeletionProtection() *bool
	SetDryRun(v bool) *CreateGadOrderShrinkRequest
	GetDryRun() *bool
	SetEncryptionKey(v string) *CreateGadOrderShrinkRequest
	GetEncryptionKey() *string
	SetEngine(v string) *CreateGadOrderShrinkRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateGadOrderShrinkRequest
	GetEngineVersion() *string
	SetInstanceNetworkType(v string) *CreateGadOrderShrinkRequest
	GetInstanceNetworkType() *string
	SetIoAccelerationEnabled(v string) *CreateGadOrderShrinkRequest
	GetIoAccelerationEnabled() *string
	SetMasterAccountExist(v bool) *CreateGadOrderShrinkRequest
	GetMasterAccountExist() *bool
	SetMasterAccountName(v string) *CreateGadOrderShrinkRequest
	GetMasterAccountName() *string
	SetMasterAccountPassword(v string) *CreateGadOrderShrinkRequest
	GetMasterAccountPassword() *string
	SetMasterDatabaseName(v string) *CreateGadOrderShrinkRequest
	GetMasterDatabaseName() *string
	SetMasterEngineArchType(v string) *CreateGadOrderShrinkRequest
	GetMasterEngineArchType() *string
	SetMasterInstanceId(v string) *CreateGadOrderShrinkRequest
	GetMasterInstanceId() *string
	SetMasterRegion(v string) *CreateGadOrderShrinkRequest
	GetMasterRegion() *string
	SetMasterShardAccountName(v string) *CreateGadOrderShrinkRequest
	GetMasterShardAccountName() *string
	SetMasterShardAccountPassword(v string) *CreateGadOrderShrinkRequest
	GetMasterShardAccountPassword() *string
	SetNeedCreateSlave(v bool) *CreateGadOrderShrinkRequest
	GetNeedCreateSlave() *bool
	SetOwnerId(v string) *CreateGadOrderShrinkRequest
	GetOwnerId() *string
	SetPayType(v string) *CreateGadOrderShrinkRequest
	GetPayType() *string
	SetPeriod(v string) *CreateGadOrderShrinkRequest
	GetPeriod() *string
	SetPort(v string) *CreateGadOrderShrinkRequest
	GetPort() *string
	SetPrivateIpAddress(v string) *CreateGadOrderShrinkRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *CreateGadOrderShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateGadOrderShrinkRequest
	GetResourceGroupId() *string
	SetRoleARN(v string) *CreateGadOrderShrinkRequest
	GetRoleARN() *string
	SetSecurityIPList(v string) *CreateGadOrderShrinkRequest
	GetSecurityIPList() *string
	SetServerlessConfigShrink(v string) *CreateGadOrderShrinkRequest
	GetServerlessConfigShrink() *string
	SetSlaveAccountExist(v bool) *CreateGadOrderShrinkRequest
	GetSlaveAccountExist() *bool
	SetSlaveAccountName(v string) *CreateGadOrderShrinkRequest
	GetSlaveAccountName() *string
	SetSlaveAccountPassword(v string) *CreateGadOrderShrinkRequest
	GetSlaveAccountPassword() *string
	SetSlaveDatabaseName(v string) *CreateGadOrderShrinkRequest
	GetSlaveDatabaseName() *string
	SetSlaveEngineArchType(v string) *CreateGadOrderShrinkRequest
	GetSlaveEngineArchType() *string
	SetSlaveInstanceId(v string) *CreateGadOrderShrinkRequest
	GetSlaveInstanceId() *string
	SetSlaveReadOnly(v bool) *CreateGadOrderShrinkRequest
	GetSlaveReadOnly() *bool
	SetSlaveRegion(v string) *CreateGadOrderShrinkRequest
	GetSlaveRegion() *string
	SetStorageAutoScale(v string) *CreateGadOrderShrinkRequest
	GetStorageAutoScale() *string
	SetStorageThreshold(v int32) *CreateGadOrderShrinkRequest
	GetStorageThreshold() *int32
	SetStorageUpperBound(v int32) *CreateGadOrderShrinkRequest
	GetStorageUpperBound() *int32
	SetSyncAccountType(v string) *CreateGadOrderShrinkRequest
	GetSyncAccountType() *string
	SetSyncArchitecture(v string) *CreateGadOrderShrinkRequest
	GetSyncArchitecture() *string
	SetSyncConflict(v string) *CreateGadOrderShrinkRequest
	GetSyncConflict() *string
	SetSystemDBCharset(v string) *CreateGadOrderShrinkRequest
	GetSystemDBCharset() *string
	SetTag(v []*CreateGadOrderShrinkRequestTag) *CreateGadOrderShrinkRequest
	GetTag() []*CreateGadOrderShrinkRequestTag
	SetTargetDedicatedHostIdForLog(v string) *CreateGadOrderShrinkRequest
	GetTargetDedicatedHostIdForLog() *string
	SetTargetDedicatedHostIdForMaster(v string) *CreateGadOrderShrinkRequest
	GetTargetDedicatedHostIdForMaster() *string
	SetTargetDedicatedHostIdForSlave(v string) *CreateGadOrderShrinkRequest
	GetTargetDedicatedHostIdForSlave() *string
	SetTargetMinorVersion(v string) *CreateGadOrderShrinkRequest
	GetTargetMinorVersion() *string
	SetUsedTime(v string) *CreateGadOrderShrinkRequest
	GetUsedTime() *string
	SetUserBackupId(v string) *CreateGadOrderShrinkRequest
	GetUserBackupId() *string
	SetVSwitchId(v string) *CreateGadOrderShrinkRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateGadOrderShrinkRequest
	GetVpcId() *string
	SetWhitelistTemplateList(v string) *CreateGadOrderShrinkRequest
	GetWhitelistTemplateList() *string
	SetZoneId(v string) *CreateGadOrderShrinkRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *CreateGadOrderShrinkRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *CreateGadOrderShrinkRequest
	GetZoneIdSlave2() *string
}

type CreateGadOrderShrinkRequest struct {
	// The number of ApsaraDB RDS for MySQL instances to be created. This parameter is applicable only when you create multiple ApsaraDB RDS for MySQL instances at a time.
	//
	// Valid values: 1 to 20. Default value: 1.
	//
	// example:
	//
	// 1
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
	// Specifies whether to enable auto-renewal upon expiration. Valid values:
	//
	// - **false**: no. This is the default value.
	//
	// - **true**: yes.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - true: enables auto-renewal.
	//
	// - false: disables auto-renewal.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The configuration information of Babelfish for ApsaraDB RDS for PostgreSQL.
	//
	// Configuration format: {"babelfishEnabled":"true","migrationMode":"xxxxxxx","masterUsername":"xxxxxxx","masterUserPassword":"xxxxxxxx"}
	//
	// Parameter descriptions:
	//
	// - **babelfishEnabled**: specifies whether to enable Babelfish. Set the value to **true*	- to enable Babelfish. By default, Babelfish is disabled if this parameter is not configured.
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
	// {\\"isAnalyticReadOnlyIns\\":\\"true\\"}
	BabelfishConfig *string `json:"BabelfishConfig,omitempty" xml:"BabelfishConfig,omitempty"`
	BpeEnabled      *string `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	// Specifies whether to enable the I/O performance burst feature for ESSD AutoPL cloud disks. Valid values:
	//
	// - **true**: enables the feature.
	//
	// - **false**: disables the feature. This is the default value.
	//
	// > This parameter is supported only when StorageType is set to ESSDAUTOPL.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The additional parameter that specifies business information.
	//
	// example:
	//
	// ***
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The instance edition. Valid values:
	//
	// General instances:
	//
	// Basic: Basic Edition.
	//
	// HighAvailability: high-availability series.
	//
	// cluster: Cluster Edition for MySQL or PostgreSQL.
	//
	// AlwaysOn: Cluster Edition for SQL Server.
	//
	// Finance: RDS Enterprise Edition.
	//
	// Note: This parameter is required when you create an ApsaraDB RDS for SQL Server instance of the RDS Enterprise Edition, Basic Edition Standard, or Basic Edition Enterprise. For example, when you create a Basic Edition instance of the 2022 RDS Enterprise Edition (2022_ent), you must set this parameter to Basic.
	//
	// Serverless instances:
	//
	// serverless_basic: Serverless Basic Edition. This edition is active only for MySQL and PostgreSQL.
	//
	// serverless_standard: Serverless high availability series. This edition is active only for MySQL and PostgreSQL.
	//
	// serverless_ha: SQL Server Serverless high availability series.
	//
	// Note: This parameter is required when PayType is set to Serverless.
	//
	// example:
	//
	// AlwaysOn
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The portal rewrites this parameter as a random UUID.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gad-x32115vv105f427
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable the [data archiving](https://help.aliyun.com/document_detail/2701832.html) feature for premium performance disks. Valid values:
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
	// Default value: allocated by the RDS system.
	//
	// > SQL Server 2012, 2016, and 2017 support only the standard access mode.
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The database connection address.
	//
	// example:
	//
	// Mid-end service
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The policy for creating multiple instances at a time. This parameter takes effect only when the **Amount*	- parameter is set to a value greater than 1. Valid values:
	//
	// 	- **Atomicity*	- (default): atomicity. All instances in the same batch are created successfully, or if one instance fails to be created, all instances fail to be created.
	//
	// 	- **Partial**: non-atomicity. The creation of an instance is not affected by other instances in the same batch.
	//
	// example:
	//
	// Atomicity
	CreateStrategy *string `json:"CreateStrategy,omitempty" xml:"CreateStrategy,omitempty"`
	// The instance type. Valid values:
	//
	// - **selectdb.xlarge**: 4 cores, 32 GB.
	//
	// - **selectdb.2xlarge**: 8 cores, 64 GB.
	//
	// - **selectdb.4xlarge**: 16 cores, 128 GB.
	//
	// - **selectdb.8xlarge**: 32 cores, 256 GB.
	//
	// - **selectdb.16xlarge**: 64 cores, 512 GB.
	//
	// - **selectdb.24xlarge**: 96 cores, 768 GB.
	//
	// - **selectdb.32xlarge**: 128 cores, 1024 GB.
	//
	// example:
	//
	// mssql.x4.medium.e2
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The name of the instance. The name must be 2 to 255 characters in length. It must start with a letter or Chinese character and can contain letters, Chinese characters, digits, and hyphens (-).
	//
	// example:
	//
	// testInstance
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The network connection type of the instance. Set the value to Intranet, which indicates an internal network connection.
	//
	// example:
	//
	// Intranet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// Instance storage capacity of the instance. Instance storage capacity of a read-only instance must be greater than or equal to that of the primary instance. Valid values: See the **Storage Capacity*	- column in [Read-only instance types](https://help.aliyun.com/document_detail/145759.html). The value must be a multiple of 5. Unit: GB.
	//
	// example:
	//
	// 100
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: premium local SSD.
	//
	// 	- **cloud_ssd**: standard SSD.
	//
	// 	- **cloud_essd**: PL1 ESSD.
	//
	// 	- **cloud_essd2**: PL2 ESSD.
	//
	// 	- **cloud_essd3**: PL3 ESSD.
	//
	// 	- **general_essd**: premium performance disk.
	//
	//
	// > 	- If the primary ApsaraDB RDS for MySQL instance uses premium local SSDs, you can set this parameter only to **local_ssd**. If the primary instance uses cloud disks, you can set this parameter only to a cloud disk storage type.
	//
	// > 	- ApsaraDB RDS for SQL Server supports only cloud disk storage types.
	//
	// example:
	//
	// cloud_auto
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
	// The ID of the parameter template. You can call the DescribeParameterGroups operation to query the ID.
	//
	// > This parameter is supported only for MySQL and PostgreSQL. If you do not specify this parameter, the default parameter template is used. You can also create a custom parameter template and use it here.
	//
	// example:
	//
	// rpg-sys-01040402010***
	DBParamGroupId *string `json:"DBParamGroupId,omitempty" xml:"DBParamGroupId,omitempty"`
	// The time zone of the instance. This parameter takes effect only when **Engine*	- is set to **MySQL*	- or **PostgreSQL**.
	//
	// - When **Engine*	- is set to **MySQL**:
	//
	//     - The parameter settings configure the UTC time zone. Valid values: **-12:59*	- to **+13:00**.
	//
	//     - An instance with Premium Local SSDs can use a named time zone, such as Asia/Hong_Kong. For more information about named time zones, see [Named time zones](https://help.aliyun.com/document_detail/297356.html).
	//
	// - When **Engine*	- is set to **PostgreSQL**:
	//
	//     - The parameter settings configure a named time zone. UTC time zones are not supported. For more information about named time zones, see [Named time zones](https://help.aliyun.com/document_detail/297356.html).
	//
	//     - This parameter can be configured only for PostgreSQL instances that use cloud disks.
	//
	// > - You can make parameter settings for the time zone when you purchase a primary instance. Read-only instances do not support time zone Settings and inherit the time zone of the primary instance.
	//
	// > - If you do not make parameter settings for this parameter, the system selects the default time zone based on the region where the instance is created.
	//
	// example:
	//
	// +08:00
	DBTimeZone *string `json:"DBTimeZone,omitempty" xml:"DBTimeZone,omitempty"`
	// The objects to be migrated or synchronized, in JSON format. For more information, see [Description of migration, synchronization, or change tracking objects](https://help.aliyun.com/document_detail/209545.html).
	//
	// - The DbList parameter supports up to 1 MB of data.
	//
	// - If DbList contains filter conditions, the total length of DbList, including the filter conditions, cannot exceed 1 MB.
	//
	// - For distributed tasks (such as migration or synchronization tasks with PolarDB-X 1.0 as the source), DbList is split based on physical table shards to generate multiple subtasks. The DbList of each subtask supports up to 1 MB of data.
	//
	// example:
	//
	// {\\"dtstest\\":{\\"name\\":\\"dtstest\\",\\"all\\":true,\\"conflict\\":\\"overwrite\\"}}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The DbList mode for SQL Server disaster recovery. Set this parameter to 0. Valid values are 0, 1, and 2, and the value must be consistent within the same group.
	//
	// example:
	//
	// 0
	DbListMode *int32 `json:"DbListMode,omitempty" xml:"DbListMode,omitempty"`
	// The ID of the existing dedicated cluster.
	//
	// example:
	//
	// dhg-4kv98863455z***
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to enable deletion protection for the resource stack. Valid values:
	//
	// - Enabled: enables deletion protection for the resource stack.
	//
	// - Disabled (default): disables deletion protection for the resource stack. In this case, you can release the resource stack by using the console or calling the DeleteStack operation.
	//
	// > The deletion protection settings for a nested resource stack are the same as those for the root resource stack.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to perform only a precheck. Valid values:
	//
	// - **true**: yes. After the precheck succeeds, no changes are made to the instance.
	//
	// - **false*	- (default): no. After the precheck succeeds, the database instance of the DTS task is modified and the task is run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the encryption key. If encryption is disabled, this parameter is empty.
	//
	// > After disk encryption is enabled for a cloud disk instance, it cannot be disabled.
	//
	// example:
	//
	// 07609cc3-3109-408f-a35e-c548e776***
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The database engine type.
	//
	// This parameter is required.
	//
	// example:
	//
	// SQLServer
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. The valid values vary based on the database engine type.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The network type of the instance. Set the value to **VPC**.
	//
	// > - Only VPCs are supported in Alibaba Cloud.
	//
	// > - If you do not specify this parameter, the default value VPC is used.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// Specifies whether to enable the Buffer Pool Extension (BPE) feature for premium performance disks. Valid values:
	//
	//  - **1**: enables the feature.
	//
	//  - **0**: disables the feature.
	//
	// > For more information about the Buffer Pool Extension (BPE) feature, see [Buffer Pool Extension (BPE)](https://help.aliyun.com/document_detail/2527067.html).
	//
	// example:
	//
	// 1
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether the primary instance account exists.
	//
	// example:
	//
	// true
	MasterAccountExist *bool `json:"MasterAccountExist,omitempty" xml:"MasterAccountExist,omitempty"`
	// The privileged account of the primary instance.
	//
	// example:
	//
	// gadtest
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The password of the primary instance account.
	//
	// example:
	//
	// <password>
	MasterAccountPassword *string `json:"MasterAccountPassword,omitempty" xml:"MasterAccountPassword,omitempty"`
	// This parameter is specific to MongoDB. The architecture type of the primary database. Valid values:
	//
	// - 1: replica set.
	//
	// - 2: sharded cluster.
	//
	// example:
	//
	// 1
	MasterDatabaseName *string `json:"MasterDatabaseName,omitempty" xml:"MasterDatabaseName,omitempty"`
	// This parameter is specific to MongoDB. The architecture type of the primary database. Valid values:
	//
	// - 1: replica set.
	//
	// - 2: sharded cluster.
	//
	// example:
	//
	// 1
	MasterEngineArchType *string `json:"MasterEngineArchType,omitempty" xml:"MasterEngineArchType,omitempty"`
	// The primary instance.
	//
	// example:
	//
	// rm-j6c685vh0j0y1k***
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// The region where the primary instance resides.
	//
	// example:
	//
	// cn-hongkong
	MasterRegion *string `json:"MasterRegion,omitempty" xml:"MasterRegion,omitempty"`
	// The shard account.
	//
	// example:
	//
	// db_test
	MasterShardAccountName *string `json:"MasterShardAccountName,omitempty" xml:"MasterShardAccountName,omitempty"`
	// The password of the shard account.
	//
	// example:
	//
	// <password>
	MasterShardAccountPassword *string `json:"MasterShardAccountPassword,omitempty" xml:"MasterShardAccountPassword,omitempty"`
	// Specifies whether to purchase a new secondary instance. If this parameter is set to false, an existing instance is used.
	//
	// example:
	//
	// false
	NeedCreateSlave *bool   `json:"NeedCreateSlave,omitempty" xml:"NeedCreateSlave,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// > This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing method for subscription instances. Valid values:
	//
	// - **Year**: yearly.
	//
	// - **Month**: monthly.
	//
	// > This parameter is valid and required only when the billing method is **PrePaid*	- (subscription).
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The database service port of the destination instance.
	//
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Settings for the internal network IP address of the instance. The IP address must be within the IP address range of the specified vSwitch. By default, the system automatically allocates an IP address based on the VpcId and VSwitchId parameters.
	//
	// example:
	//
	// 10.60.22.91
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID of the primary instance.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmx4gs2dcym***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The global resource descriptor (ARN) for the authorization of the RDS cloud service account by the primary account to access KMS. You can call the CheckCloudResourceAuthorized operation to query the ARN.
	//
	// 	Notice: When you enable cloud disk encryption, you must specify **RoleARN**.</notice>
	//
	// example:
	//
	// acs:ram::1929332604914597:role/aliyunr***tanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The IP addresses that are added to the whitelist of the instance. Separate multiple IP addresses with commas (,). Note: If you set this parameter to 127.0.0.1, access from all IP addresses is denied. If you set this parameter to 192.168.0.0/24, access from all IP addresses in the 192.168.0.XX format is allowed.
	//
	// example:
	//
	// 172.16.0.0/12
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The settings of the Serverless instance. This parameter is required when you create a Serverless instance.
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// Specifies whether the secondary instance account exists.
	//
	// example:
	//
	// true
	SlaveAccountExist *bool `json:"SlaveAccountExist,omitempty" xml:"SlaveAccountExist,omitempty"`
	// The account with the same name on the secondary instance.
	//
	// example:
	//
	// gadtest
	SlaveAccountName *string `json:"SlaveAccountName,omitempty" xml:"SlaveAccountName,omitempty"`
	// The password of the secondary instance account.
	//
	// example:
	//
	// <password>
	SlaveAccountPassword *string `json:"SlaveAccountPassword,omitempty" xml:"SlaveAccountPassword,omitempty"`
	// The name of the secondary database.
	//
	// example:
	//
	// slave_database
	SlaveDatabaseName *string `json:"SlaveDatabaseName,omitempty" xml:"SlaveDatabaseName,omitempty"`
	// This parameter is specific to MongoDB. The architecture type of the secondary database. Valid values:
	//
	// - 1: replica set.
	//
	// - 2: sharded cluster.
	//
	// example:
	//
	// 1
	SlaveEngineArchType *string `json:"SlaveEngineArchType,omitempty" xml:"SlaveEngineArchType,omitempty"`
	// The secondary instance.
	//
	// example:
	//
	// rm-t4nb481a8qn0b3***
	SlaveInstanceId *string `json:"SlaveInstanceId,omitempty" xml:"SlaveInstanceId,omitempty"`
	// Specifies whether the secondary instance is read-only.
	//
	// example:
	//
	// true
	SlaveReadOnly *bool `json:"SlaveReadOnly,omitempty" xml:"SlaveReadOnly,omitempty"`
	// The region where the secondary instance resides.
	//
	// example:
	//
	// ap-southeast-1
	SlaveRegion *string `json:"SlaveRegion,omitempty" xml:"SlaveRegion,omitempty"`
	// Specifies whether to enable Automatic storage scaling for the Standard Edition cluster. Valid values:
	//
	// - Enable: enables Automatic storage scaling.
	//
	// - Disable: shutdown Automatic storage scaling.
	//
	// example:
	//
	// Disable
	StorageAutoScale *string `json:"StorageAutoScale,omitempty" xml:"StorageAutoScale,omitempty"`
	// The threshold (in percentage) that triggers automatic storage expansion. Valid values:
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
	// 60
	StorageThreshold *int32 `json:"StorageThreshold,omitempty" xml:"StorageThreshold,omitempty"`
	// The Settings for the upper limit of Automatic storage scaling for the Standard Edition cluster. Unit: GB.
	//
	// > The maximum value is 32000.
	//
	// example:
	//
	// 64000
	StorageUpperBound *int32 `json:"StorageUpperBound,omitempty" xml:"StorageUpperBound,omitempty"`
	// The synchronization account type.
	//
	// example:
	//
	// user
	SyncAccountType *string `json:"SyncAccountType,omitempty" xml:"SyncAccountType,omitempty"`
	// The synchronization topology. Valid values:
	//
	// - **oneway**: one-way synchronization. This is the default value.
	//
	// - **bidirectional**: two-way synchronization.
	//
	// example:
	//
	// oneway
	SyncArchitecture *string `json:"SyncArchitecture,omitempty" xml:"SyncArchitecture,omitempty"`
	// The conflict resolution policy.
	//
	// example:
	//
	// ignore
	SyncConflict *string `json:"SyncConflict,omitempty" xml:"SyncConflict,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and does not need to be configured.
	//
	// example:
	//
	// utf8
	SystemDBCharset *string `json:"SystemDBCharset,omitempty" xml:"SystemDBCharset,omitempty"`
	// The tags of the data migration instance. This parameter is used as a filter. If this parameter is not empty, only the instance tasks that have the specified tags are returned.
	Tag []*CreateGadOrderShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the host on which the logger instance is created in the dedicated cluster.
	//
	// This parameter is required when you create an RDS Enterprise Edition instance in a dedicated cluster. If you do not specify this parameter, the system automatically allocates a host by default.
	//
	// - You can invoke the DescribeDedicatedHosts operation to query the host information in the dedicated cluster.
	//
	// - If no hosts are active, you can invoke the CreateDedicatedHost operation to create a host.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForLog *string `json:"TargetDedicatedHostIdForLog,omitempty" xml:"TargetDedicatedHostIdForLog,omitempty"`
	// The ID of the host on which the primary instance is created in the dedicated cluster.
	//
	// This parameter is required when you create an instance in a dedicated cluster. If you do not specify this parameter, the system automatically assigns a host.
	//
	// You can call the DescribeDedicatedHosts operation to query the host information in the dedicated cluster.
	//
	// If no hosts are available, you can call the CreateDedicatedHost operation to create a host.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// The ID of the host on which the secondary instance is created in the dedicated cluster.
	//
	// This parameter is required when you create a high availability series or RDS Enterprise Edition instance in a dedicated cluster. If you do not specify this parameter, the system automatically allocates a host by default.
	//
	// - You can invoke the DescribeDedicatedHosts operation to query the host information in the dedicated cluster.
	//
	// - If no hosts are active, you can invoke the CreateDedicatedHost operation to create a host.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForSlave *string `json:"TargetDedicatedHostIdForSlave,omitempty" xml:"TargetDedicatedHostIdForSlave,omitempty"`
	// The minor engine version of the ApsaraDB RDS instance. This parameter is required only when you create a MySQL or PostgreSQL instance.
	//
	// example:
	//
	// rds_postgres_1700_20260***
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The purchase duration for subscription instances.
	//
	// - When **Period*	- is set to **Month**, valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// - When **Period*	- is set to **Year**, valid values are 1, 2, 3, and 5.
	//
	// > - This parameter is valid and required only when **PayType*	- is set to **PrePaid**.
	//
	// - For the billing method of subscription instances, you can specify it in the **Period*	- parameter.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The ID of the user backup. You can call the ListUserBackupFiles operation to query the ID. If you specify this parameter, the instance is created based on the user backup.
	//
	// If you specify this parameter, the following limits apply:
	//
	// - The PayType parameter must be set to Postpaid.
	//
	// - The Engine parameter must be set to MySQL.
	//
	// - The EngineVersion parameter must be set to 5.7.
	//
	// - The Category parameter must be set to Basic.
	//
	// example:
	//
	// 67798****
	UserBackupId *string `json:"UserBackupId,omitempty" xml:"UserBackupId,omitempty"`
	// The ID of the vSwitch in the VPC.
	//
	// example:
	//
	// vsw-wz9gttyapd3i9dmzn***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-m5ehjekt8mzvyyx5a1***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The whitelist. To configure multiple IP addresses, separate the IP addresses or CIDR blocks with commas (,) without spaces before or after the commas. Example: `192.168.0.1,172.16.213.9`.
	//
	// example:
	//
	// 192.168.0.1,172.16.213.9
	WhitelistTemplateList *string `json:"WhitelistTemplateList,omitempty" xml:"WhitelistTemplateList,omitempty"`
	// The zone ID of the primary node.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary node.
	//
	// - If you set this parameter to **Auto**, it indicates a multi-zone deployment and the system automatically selects an active zone for the secondary node.
	//
	// - If you set this parameter to the same value as **ZoneId**, it is a single-zone deployment.
	//
	// - If you set this parameter to a value different from **ZoneId**, it is a multi-zone deployment.
	//
	// example:
	//
	// cn-beijing-i
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The secondary zone 2.
	//
	// > This parameter is active only for instances of the RDS Enterprise Edition.
	//
	// example:
	//
	// cn-chengdu-b
	ZoneIdSlave2 *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s CreateGadOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGadOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGadOrderShrinkRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateGadOrderShrinkRequest) GetAutoCreateProxy() *bool {
	return s.AutoCreateProxy
}

func (s *CreateGadOrderShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateGadOrderShrinkRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateGadOrderShrinkRequest) GetBabelfishConfig() *string {
	return s.BabelfishConfig
}

func (s *CreateGadOrderShrinkRequest) GetBpeEnabled() *string {
	return s.BpeEnabled
}

func (s *CreateGadOrderShrinkRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateGadOrderShrinkRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateGadOrderShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateGadOrderShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGadOrderShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateGadOrderShrinkRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *CreateGadOrderShrinkRequest) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *CreateGadOrderShrinkRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateGadOrderShrinkRequest) GetCreateStrategy() *string {
	return s.CreateStrategy
}

func (s *CreateGadOrderShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateGadOrderShrinkRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateGadOrderShrinkRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *CreateGadOrderShrinkRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CreateGadOrderShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CreateGadOrderShrinkRequest) GetDBIsIgnoreCase() *string {
	return s.DBIsIgnoreCase
}

func (s *CreateGadOrderShrinkRequest) GetDBParamGroupId() *string {
	return s.DBParamGroupId
}

func (s *CreateGadOrderShrinkRequest) GetDBTimeZone() *string {
	return s.DBTimeZone
}

func (s *CreateGadOrderShrinkRequest) GetDbList() *string {
	return s.DbList
}

func (s *CreateGadOrderShrinkRequest) GetDbListMode() *int32 {
	return s.DbListMode
}

func (s *CreateGadOrderShrinkRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *CreateGadOrderShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateGadOrderShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateGadOrderShrinkRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateGadOrderShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateGadOrderShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateGadOrderShrinkRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CreateGadOrderShrinkRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *CreateGadOrderShrinkRequest) GetMasterAccountExist() *bool {
	return s.MasterAccountExist
}

func (s *CreateGadOrderShrinkRequest) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *CreateGadOrderShrinkRequest) GetMasterAccountPassword() *string {
	return s.MasterAccountPassword
}

func (s *CreateGadOrderShrinkRequest) GetMasterDatabaseName() *string {
	return s.MasterDatabaseName
}

func (s *CreateGadOrderShrinkRequest) GetMasterEngineArchType() *string {
	return s.MasterEngineArchType
}

func (s *CreateGadOrderShrinkRequest) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *CreateGadOrderShrinkRequest) GetMasterRegion() *string {
	return s.MasterRegion
}

func (s *CreateGadOrderShrinkRequest) GetMasterShardAccountName() *string {
	return s.MasterShardAccountName
}

func (s *CreateGadOrderShrinkRequest) GetMasterShardAccountPassword() *string {
	return s.MasterShardAccountPassword
}

func (s *CreateGadOrderShrinkRequest) GetNeedCreateSlave() *bool {
	return s.NeedCreateSlave
}

func (s *CreateGadOrderShrinkRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateGadOrderShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateGadOrderShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateGadOrderShrinkRequest) GetPort() *string {
	return s.Port
}

func (s *CreateGadOrderShrinkRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateGadOrderShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGadOrderShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGadOrderShrinkRequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *CreateGadOrderShrinkRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateGadOrderShrinkRequest) GetServerlessConfigShrink() *string {
	return s.ServerlessConfigShrink
}

func (s *CreateGadOrderShrinkRequest) GetSlaveAccountExist() *bool {
	return s.SlaveAccountExist
}

func (s *CreateGadOrderShrinkRequest) GetSlaveAccountName() *string {
	return s.SlaveAccountName
}

func (s *CreateGadOrderShrinkRequest) GetSlaveAccountPassword() *string {
	return s.SlaveAccountPassword
}

func (s *CreateGadOrderShrinkRequest) GetSlaveDatabaseName() *string {
	return s.SlaveDatabaseName
}

func (s *CreateGadOrderShrinkRequest) GetSlaveEngineArchType() *string {
	return s.SlaveEngineArchType
}

func (s *CreateGadOrderShrinkRequest) GetSlaveInstanceId() *string {
	return s.SlaveInstanceId
}

func (s *CreateGadOrderShrinkRequest) GetSlaveReadOnly() *bool {
	return s.SlaveReadOnly
}

func (s *CreateGadOrderShrinkRequest) GetSlaveRegion() *string {
	return s.SlaveRegion
}

func (s *CreateGadOrderShrinkRequest) GetStorageAutoScale() *string {
	return s.StorageAutoScale
}

func (s *CreateGadOrderShrinkRequest) GetStorageThreshold() *int32 {
	return s.StorageThreshold
}

func (s *CreateGadOrderShrinkRequest) GetStorageUpperBound() *int32 {
	return s.StorageUpperBound
}

func (s *CreateGadOrderShrinkRequest) GetSyncAccountType() *string {
	return s.SyncAccountType
}

func (s *CreateGadOrderShrinkRequest) GetSyncArchitecture() *string {
	return s.SyncArchitecture
}

func (s *CreateGadOrderShrinkRequest) GetSyncConflict() *string {
	return s.SyncConflict
}

func (s *CreateGadOrderShrinkRequest) GetSystemDBCharset() *string {
	return s.SystemDBCharset
}

func (s *CreateGadOrderShrinkRequest) GetTag() []*CreateGadOrderShrinkRequestTag {
	return s.Tag
}

func (s *CreateGadOrderShrinkRequest) GetTargetDedicatedHostIdForLog() *string {
	return s.TargetDedicatedHostIdForLog
}

func (s *CreateGadOrderShrinkRequest) GetTargetDedicatedHostIdForMaster() *string {
	return s.TargetDedicatedHostIdForMaster
}

func (s *CreateGadOrderShrinkRequest) GetTargetDedicatedHostIdForSlave() *string {
	return s.TargetDedicatedHostIdForSlave
}

func (s *CreateGadOrderShrinkRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *CreateGadOrderShrinkRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateGadOrderShrinkRequest) GetUserBackupId() *string {
	return s.UserBackupId
}

func (s *CreateGadOrderShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateGadOrderShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGadOrderShrinkRequest) GetWhitelistTemplateList() *string {
	return s.WhitelistTemplateList
}

func (s *CreateGadOrderShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateGadOrderShrinkRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *CreateGadOrderShrinkRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *CreateGadOrderShrinkRequest) SetAmount(v int32) *CreateGadOrderShrinkRequest {
	s.Amount = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetAutoCreateProxy(v bool) *CreateGadOrderShrinkRequest {
	s.AutoCreateProxy = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetAutoPay(v bool) *CreateGadOrderShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetAutoRenew(v string) *CreateGadOrderShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetBabelfishConfig(v string) *CreateGadOrderShrinkRequest {
	s.BabelfishConfig = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetBpeEnabled(v string) *CreateGadOrderShrinkRequest {
	s.BpeEnabled = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetBurstingEnabled(v bool) *CreateGadOrderShrinkRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetBusinessInfo(v string) *CreateGadOrderShrinkRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetCategory(v string) *CreateGadOrderShrinkRequest {
	s.Category = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetClientToken(v string) *CreateGadOrderShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetClusterId(v string) *CreateGadOrderShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetColdDataEnabled(v bool) *CreateGadOrderShrinkRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetConnectionMode(v string) *CreateGadOrderShrinkRequest {
	s.ConnectionMode = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetConnectionString(v string) *CreateGadOrderShrinkRequest {
	s.ConnectionString = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetCreateStrategy(v string) *CreateGadOrderShrinkRequest {
	s.CreateStrategy = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDBInstanceClass(v string) *CreateGadOrderShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDBInstanceDescription(v string) *CreateGadOrderShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDBInstanceNetType(v string) *CreateGadOrderShrinkRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDBInstanceStorage(v int32) *CreateGadOrderShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDBInstanceStorageType(v string) *CreateGadOrderShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDBIsIgnoreCase(v string) *CreateGadOrderShrinkRequest {
	s.DBIsIgnoreCase = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDBParamGroupId(v string) *CreateGadOrderShrinkRequest {
	s.DBParamGroupId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDBTimeZone(v string) *CreateGadOrderShrinkRequest {
	s.DBTimeZone = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDbList(v string) *CreateGadOrderShrinkRequest {
	s.DbList = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDbListMode(v int32) *CreateGadOrderShrinkRequest {
	s.DbListMode = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDedicatedHostGroupId(v string) *CreateGadOrderShrinkRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDeletionProtection(v bool) *CreateGadOrderShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetDryRun(v bool) *CreateGadOrderShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetEncryptionKey(v string) *CreateGadOrderShrinkRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetEngine(v string) *CreateGadOrderShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetEngineVersion(v string) *CreateGadOrderShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetInstanceNetworkType(v string) *CreateGadOrderShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetIoAccelerationEnabled(v string) *CreateGadOrderShrinkRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterAccountExist(v bool) *CreateGadOrderShrinkRequest {
	s.MasterAccountExist = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterAccountName(v string) *CreateGadOrderShrinkRequest {
	s.MasterAccountName = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterAccountPassword(v string) *CreateGadOrderShrinkRequest {
	s.MasterAccountPassword = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterDatabaseName(v string) *CreateGadOrderShrinkRequest {
	s.MasterDatabaseName = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterEngineArchType(v string) *CreateGadOrderShrinkRequest {
	s.MasterEngineArchType = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterInstanceId(v string) *CreateGadOrderShrinkRequest {
	s.MasterInstanceId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterRegion(v string) *CreateGadOrderShrinkRequest {
	s.MasterRegion = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterShardAccountName(v string) *CreateGadOrderShrinkRequest {
	s.MasterShardAccountName = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetMasterShardAccountPassword(v string) *CreateGadOrderShrinkRequest {
	s.MasterShardAccountPassword = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetNeedCreateSlave(v bool) *CreateGadOrderShrinkRequest {
	s.NeedCreateSlave = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetOwnerId(v string) *CreateGadOrderShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetPayType(v string) *CreateGadOrderShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetPeriod(v string) *CreateGadOrderShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetPort(v string) *CreateGadOrderShrinkRequest {
	s.Port = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetPrivateIpAddress(v string) *CreateGadOrderShrinkRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetRegionId(v string) *CreateGadOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetResourceGroupId(v string) *CreateGadOrderShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetRoleARN(v string) *CreateGadOrderShrinkRequest {
	s.RoleARN = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSecurityIPList(v string) *CreateGadOrderShrinkRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetServerlessConfigShrink(v string) *CreateGadOrderShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSlaveAccountExist(v bool) *CreateGadOrderShrinkRequest {
	s.SlaveAccountExist = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSlaveAccountName(v string) *CreateGadOrderShrinkRequest {
	s.SlaveAccountName = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSlaveAccountPassword(v string) *CreateGadOrderShrinkRequest {
	s.SlaveAccountPassword = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSlaveDatabaseName(v string) *CreateGadOrderShrinkRequest {
	s.SlaveDatabaseName = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSlaveEngineArchType(v string) *CreateGadOrderShrinkRequest {
	s.SlaveEngineArchType = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSlaveInstanceId(v string) *CreateGadOrderShrinkRequest {
	s.SlaveInstanceId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSlaveReadOnly(v bool) *CreateGadOrderShrinkRequest {
	s.SlaveReadOnly = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSlaveRegion(v string) *CreateGadOrderShrinkRequest {
	s.SlaveRegion = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetStorageAutoScale(v string) *CreateGadOrderShrinkRequest {
	s.StorageAutoScale = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetStorageThreshold(v int32) *CreateGadOrderShrinkRequest {
	s.StorageThreshold = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetStorageUpperBound(v int32) *CreateGadOrderShrinkRequest {
	s.StorageUpperBound = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSyncAccountType(v string) *CreateGadOrderShrinkRequest {
	s.SyncAccountType = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSyncArchitecture(v string) *CreateGadOrderShrinkRequest {
	s.SyncArchitecture = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSyncConflict(v string) *CreateGadOrderShrinkRequest {
	s.SyncConflict = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetSystemDBCharset(v string) *CreateGadOrderShrinkRequest {
	s.SystemDBCharset = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetTag(v []*CreateGadOrderShrinkRequestTag) *CreateGadOrderShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetTargetDedicatedHostIdForLog(v string) *CreateGadOrderShrinkRequest {
	s.TargetDedicatedHostIdForLog = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetTargetDedicatedHostIdForMaster(v string) *CreateGadOrderShrinkRequest {
	s.TargetDedicatedHostIdForMaster = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetTargetDedicatedHostIdForSlave(v string) *CreateGadOrderShrinkRequest {
	s.TargetDedicatedHostIdForSlave = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetTargetMinorVersion(v string) *CreateGadOrderShrinkRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetUsedTime(v string) *CreateGadOrderShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetUserBackupId(v string) *CreateGadOrderShrinkRequest {
	s.UserBackupId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetVSwitchId(v string) *CreateGadOrderShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetVpcId(v string) *CreateGadOrderShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetWhitelistTemplateList(v string) *CreateGadOrderShrinkRequest {
	s.WhitelistTemplateList = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetZoneId(v string) *CreateGadOrderShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetZoneIdSlave1(v string) *CreateGadOrderShrinkRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) SetZoneIdSlave2(v string) *CreateGadOrderShrinkRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *CreateGadOrderShrinkRequest) Validate() error {
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

type CreateGadOrderShrinkRequestTag struct {
	// The tag key. You can call the [ListTagResources](https://help.aliyun.com/document_detail/191187.html) operation to query the tag key.
	//
	// > - N indicates the sequence number of the tag key. For example, Tag.1.Key indicates the key of the first tag, and Tag.2.Key indicates the key of the second tag. You can specify 1 to 20 tag keys at a time.
	//
	// - This parameter cannot be set to an empty string.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can call the [ListTagResources](https://help.aliyun.com/document_detail/191187.html) operation to query the tag value.
	//
	// > - N indicates the sequence number of the tag value. For example, Tag.1.Value indicates the value of the first tag, and Tag.2.Value indicates the value of the second tag. You can specify 1 to 20 tag values at a time.
	//
	// - This parameter can be set to an empty string.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGadOrderShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateGadOrderShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateGadOrderShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateGadOrderShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateGadOrderShrinkRequestTag) SetKey(v string) *CreateGadOrderShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateGadOrderShrinkRequestTag) SetValue(v string) *CreateGadOrderShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateGadOrderShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
