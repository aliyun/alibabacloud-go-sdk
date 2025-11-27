// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CloneDBInstanceRequest
	GetAutoPay() *bool
	SetBackupId(v string) *CloneDBInstanceRequest
	GetBackupId() *string
	SetBackupType(v string) *CloneDBInstanceRequest
	GetBackupType() *string
	SetBpeEnabled(v string) *CloneDBInstanceRequest
	GetBpeEnabled() *string
	SetBurstingEnabled(v bool) *CloneDBInstanceRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *CloneDBInstanceRequest
	GetCategory() *string
	SetClientToken(v string) *CloneDBInstanceRequest
	GetClientToken() *string
	SetCustomExtraInfo(v string) *CloneDBInstanceRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *CloneDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CloneDBInstanceRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *CloneDBInstanceRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *CloneDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *CloneDBInstanceRequest
	GetDBInstanceStorageType() *string
	SetDbNames(v string) *CloneDBInstanceRequest
	GetDbNames() *string
	SetDedicatedHostGroupId(v string) *CloneDBInstanceRequest
	GetDedicatedHostGroupId() *string
	SetDeletionProtection(v bool) *CloneDBInstanceRequest
	GetDeletionProtection() *bool
	SetInstanceNetworkType(v string) *CloneDBInstanceRequest
	GetInstanceNetworkType() *string
	SetIoAccelerationEnabled(v string) *CloneDBInstanceRequest
	GetIoAccelerationEnabled() *string
	SetPayType(v string) *CloneDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CloneDBInstanceRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *CloneDBInstanceRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *CloneDBInstanceRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *CloneDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTable(v string) *CloneDBInstanceRequest
	GetRestoreTable() *string
	SetRestoreTime(v string) *CloneDBInstanceRequest
	GetRestoreTime() *string
	SetServerlessConfig(v *CloneDBInstanceRequestServerlessConfig) *CloneDBInstanceRequest
	GetServerlessConfig() *CloneDBInstanceRequestServerlessConfig
	SetTableMeta(v string) *CloneDBInstanceRequest
	GetTableMeta() *string
	SetUsedTime(v int32) *CloneDBInstanceRequest
	GetUsedTime() *int32
	SetVPCId(v string) *CloneDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CloneDBInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CloneDBInstanceRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *CloneDBInstanceRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *CloneDBInstanceRequest
	GetZoneIdSlave2() *string
}

type CloneDBInstanceRequest struct {
	// Specifies whether to enable the automatic payment feature for the new instance. Valid values:
	//
	// 1.  **true**: enables the feature. You must make sure that your account balance is sufficient.
	//
	// 2.  **false**: disables the feature. An unpaid order is generated.
	//
	// >  Default value: true. If your account balance is insufficient, you can set the AutoPay parameter to false to generate an unpaid order. Then, you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The backup set ID.
	//
	// You can call the DescribeBackups operation to query the backup set ID.
	//
	// >  You must specify at least one of the **BackupId*	- or **RestoreTime*	- parameters.
	//
	// example:
	//
	// 902****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The type of backup that is used to restore the data of the original instance. Valid values:
	//
	// 	- **FullBackup**
	//
	// 	- **IncrementalBackup**
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// false
	BpeEnabled *string `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	// An invalid parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition.
	//
	// 	- **HighAvailability**: RDS High-availability Edition.
	//
	// 	- **AlwaysOn**: RDS Cluster Edition for ApsaraDB RDS for SQL Server.
	//
	// 	- **cluster**: RDS Cluster Edition for ApsaraDB RDS for MySQL.
	//
	// 	- **Finance**: RDS Enterprise Edition. This edition is available only on the China site (aliyun.com).
	//
	// **Serverless instances**
	//
	// 	- **serverless_basic**: RDS Basic Edition. This edition is available only for serverless instances that run MySQL and PostgreSQL.
	//
	// 	- **serverless_standard**: RDS High-availability Edition for ApsaraDB RDS for MySQL
	//
	// 	- **serverless_ha**: RDS High-availability Edition for ApsaraDB RDS for SQL Server
	//
	// >  You do not need to configure this parameter. The value of this parameter is the same as that of the original instance.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type of the new instance. For information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// > By default, the new instance uses the same instance type as the original primary instance.
	//
	// example:
	//
	// mysql.n1.micro.1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance name. The value must be 2 to 255 characters in length The value can contain letters, digits, underscores (_), and hyphens (-), and must start with a letter.
	//
	// >  The value cannot start with http:// or https://.
	//
	// example:
	//
	// testInstance
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The storage capacity of the new instance. Unit: GB. You can increase the storage capacity in increments of 5 GB. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// > By default, the new instance has the same storage capacity as the original primary instance.
	//
	// example:
	//
	// 1000
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the new instance. Valid values:
	//
	// 	- **general_essd*	- (recommend): general Enterprise SSD (ESSD)
	//
	// 	- **local_ssd**: local SSD
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: performance level 1 (PL1) ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// >  Serverless instances support only PL1 ESSDs and general ESSDs.
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The name of the database. If you specify more than one database, the value is in the following format: `Original database name 1,Original database name 2`.
	//
	// example:
	//
	// test1,test2
	DbNames *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	// The ID of the dedicated cluster.
	//
	// example:
	//
	// dhg-7a9xxxxxxxx
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to enable the release protection feature for the new instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The network type of the new instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **Classic**
	//
	// > By default, the new instance has the same network type as the original primary instance.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// 	- **Serverless**: serverless. This value is not supported for instances that run MariaDB. For more information, see [Overview of serverless ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/411291.html), [Overview of serverless ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of serverless ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/607742.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit that is used to calculate the billing cycle of the new instance. This parameter takes effect only when you select the subscription billing method for the new instance. Valid values:
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
	// The internal IP address of the new instance, which must be within the CIDR block supported by the specified vSwitch. The system automatically assigns an internal IP address based on the values of the **VPCId*	- and **VSwitchId*	- parameters.
	//
	// example:
	//
	// 172.XX.XXX.69
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to restore only the databases and tables that you specify. The value **1*	- specifies to restore only the specified databases and tables. If you do not want to restore only the specified databases or tables, you do not need to specify this parameter.
	//
	// example:
	//
	// 1
	RestoreTable *string `json:"RestoreTable,omitempty" xml:"RestoreTable,omitempty"`
	// The point in time to which you want to restore data. The point in time must fall within the specified backup retention period. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > You must specify at least one of the **BackupId*	- and **RestoreTime*	- parameters.
	//
	// example:
	//
	// 2011-06-11T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The specifications for the serverless instance. You must specify this parameter only when you restore data to a new serverless instance.
	//
	// >  This parameter is available only on the China site (aliyun.com).
	ServerlessConfig *CloneDBInstanceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The information about the database and table that you want to restore. The value is in the following format: `[{"type":"db","name":"Name of Database 1","newname":"New name of Database 1","tables":[{"type":"table","name":"Name of Table 1 in Database 1","newname":"New name of Table 1"},{"type":"table","name":"Name of Table 2 in Database 1","newname":"New name of Table 2"}]},{"type":"db","name":"Name of Database 2","newname":"New name of Database 2","tables":[{"type":"table","name":"Name of Table 1 in Database 2","newname":"New name of Table 1"},{"type":"table","name":"Name of Table 2 in Database 2","newname":"New name of Table 2"}]}]`
	//
	// example:
	//
	// [{"type":"db","name":"testdb1","newname":"testdb1_new","tables":[{"type":"table","name":"testdb1table1","newname":"testdb1table1_new"}]}]
	TableMeta *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
	// The subscription duration of the new instance. Valid values:
	//
	// 	- If you set the **Period*	- parameter to **Year**, the value of the UsedTime parameter ranges from **1 to 3**.
	//
	// 	- If you set the **Period*	- parameter to **Month**, the value of the UsedTime parameter ranges from **1 to 9**.
	//
	// > If you set the PayType parameter to **Prepaid**, you must also specify this parameter.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// >  Make sure that the VPC belongs to the required region.
	//
	// example:
	//
	// vpc-uf6f7l4fg90xxxxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch. The vSwitch must belong to the zone that is specified by **ZoneId**.
	//
	// 	- If you set **InstanceNetworkType*	- to **VPC**, you must also specify this parameter.
	//
	// 	- If you specify the **ZoneSlaveId1*	- parameter, you must specify the IDs of two vSwitches for this parameter and separate the IDs with a comma (,).
	//
	// example:
	//
	// vsw-uf6adz52c2pxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the primary instance. You can call the DescribeRegions operation to query the zone ID.
	//
	// >  Set this value to the zone ID of the original instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary instance. If you set the ZoneIdSlave1 parameter and the **ZoneId*	- parameter to the same value, the single-zone deployment method is used. If you set the ZoneIdSlave1 parameter and the **ZoneId*	- parameter to different values, the multi-zone deployment method is used.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The zone ID of the logger instance. If you set the ZoneIdSlave2 parameter to the same value as the **ZoneId*	- parameter, the single-zone deployment method is used. If you set the ZoneIdSlave2 parameter to a different value from the **ZoneId*	- parameter, the multi-zone deployment method is used.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneIdSlave2 *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s CloneDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CloneDBInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CloneDBInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CloneDBInstanceRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *CloneDBInstanceRequest) GetBpeEnabled() *string {
	return s.BpeEnabled
}

func (s *CloneDBInstanceRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CloneDBInstanceRequest) GetCategory() *string {
	return s.Category
}

func (s *CloneDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CloneDBInstanceRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *CloneDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CloneDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CloneDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CloneDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CloneDBInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CloneDBInstanceRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *CloneDBInstanceRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *CloneDBInstanceRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CloneDBInstanceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CloneDBInstanceRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *CloneDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CloneDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CloneDBInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CloneDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloneDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloneDBInstanceRequest) GetRestoreTable() *string {
	return s.RestoreTable
}

func (s *CloneDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CloneDBInstanceRequest) GetServerlessConfig() *CloneDBInstanceRequestServerlessConfig {
	return s.ServerlessConfig
}

func (s *CloneDBInstanceRequest) GetTableMeta() *string {
	return s.TableMeta
}

func (s *CloneDBInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CloneDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CloneDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CloneDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CloneDBInstanceRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *CloneDBInstanceRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *CloneDBInstanceRequest) SetAutoPay(v bool) *CloneDBInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CloneDBInstanceRequest) SetBackupId(v string) *CloneDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetBackupType(v string) *CloneDBInstanceRequest {
	s.BackupType = &v
	return s
}

func (s *CloneDBInstanceRequest) SetBpeEnabled(v string) *CloneDBInstanceRequest {
	s.BpeEnabled = &v
	return s
}

func (s *CloneDBInstanceRequest) SetBurstingEnabled(v bool) *CloneDBInstanceRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CloneDBInstanceRequest) SetCategory(v string) *CloneDBInstanceRequest {
	s.Category = &v
	return s
}

func (s *CloneDBInstanceRequest) SetClientToken(v string) *CloneDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CloneDBInstanceRequest) SetCustomExtraInfo(v string) *CloneDBInstanceRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDBInstanceClass(v string) *CloneDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDBInstanceDescription(v string) *CloneDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDBInstanceId(v string) *CloneDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDBInstanceStorage(v int32) *CloneDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDBInstanceStorageType(v string) *CloneDBInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDbNames(v string) *CloneDBInstanceRequest {
	s.DbNames = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDedicatedHostGroupId(v string) *CloneDBInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetDeletionProtection(v bool) *CloneDBInstanceRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CloneDBInstanceRequest) SetInstanceNetworkType(v string) *CloneDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CloneDBInstanceRequest) SetIoAccelerationEnabled(v string) *CloneDBInstanceRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *CloneDBInstanceRequest) SetPayType(v string) *CloneDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CloneDBInstanceRequest) SetPeriod(v string) *CloneDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CloneDBInstanceRequest) SetPrivateIpAddress(v string) *CloneDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CloneDBInstanceRequest) SetRegionId(v string) *CloneDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetResourceOwnerId(v int64) *CloneDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetRestoreTable(v string) *CloneDBInstanceRequest {
	s.RestoreTable = &v
	return s
}

func (s *CloneDBInstanceRequest) SetRestoreTime(v string) *CloneDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CloneDBInstanceRequest) SetServerlessConfig(v *CloneDBInstanceRequestServerlessConfig) *CloneDBInstanceRequest {
	s.ServerlessConfig = v
	return s
}

func (s *CloneDBInstanceRequest) SetTableMeta(v string) *CloneDBInstanceRequest {
	s.TableMeta = &v
	return s
}

func (s *CloneDBInstanceRequest) SetUsedTime(v int32) *CloneDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CloneDBInstanceRequest) SetVPCId(v string) *CloneDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetVSwitchId(v string) *CloneDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetZoneId(v string) *CloneDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CloneDBInstanceRequest) SetZoneIdSlave1(v string) *CloneDBInstanceRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *CloneDBInstanceRequest) SetZoneIdSlave2(v string) *CloneDBInstanceRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *CloneDBInstanceRequest) Validate() error {
	if s.ServerlessConfig != nil {
		if err := s.ServerlessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloneDBInstanceRequestServerlessConfig struct {
	// Specifies whether to enable the automatic start and stop feature for the serverless ApsaraDB RDS for MySQL instance. After the automatic start and stop feature is enabled, if no connections to the instance are established within 10 minutes, the instance is suspended. After a connection is established to the instance, the instance is automatically resumed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > 	- This parameter is supported only for serverless ApsaraDB RDS for MySQL instances.
	//
	// > 	- This parameter is available only on the China site (aliyun.com).
	//
	// example:
	//
	// true
	AutoPause *bool `json:"AutoPause,omitempty" xml:"AutoPause,omitempty"`
	// The maximum number of RDS Capacity Units (RCUs). Valid values:
	//
	// 	- Serverless ApsaraDB RDS for MySQL instances: **1 to 8**
	//
	// 	- Serverless ApsaraDB RDS for SQL Server instances: **2 to 8**
	//
	// 	- Serverless ApsaraDB RDS for PostgreSQL instances: **1 to 12**
	//
	// > 	- The value of this parameter must be greater than or equal to the value of **MinCapacity*	- and can be specified only to an **integer**.
	//
	// > 	- This parameter is available only on the China site (aliyun.com).
	//
	// example:
	//
	// 8
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The minimum number of RCUs. Valid values:
	//
	// 	- Serverless ApsaraDB RDS for MySQL instances: **0.5 to 8**.
	//
	// 	- Serverless ApsaraDB RDS for SQL Server instances: **2 to 8**. Only integers are supported.
	//
	// 	- Serverless ApsaraDB RDS for PostgreSQL instances: **0.5 to 12**.
	//
	// > 	- The value of this parameter must be less than or equal to the value of **MaxCapacity**.
	//
	// > 	- This parameter is available only on the China site (aliyun.com).
	//
	// example:
	//
	// 0.5
	MinCapacity *float64 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
	// Specifies whether to enable the forced scaling feature for the serverless ApsaraDB RDS for MySQL instance. In most cases, ApsaraDB RDS automatically scales in or out the RCUs of a serverless instance based on business requirements in real time. In rare cases, the scaling does not take effect in real time. You can enable the forced scaling feature to forcefully scales in or out the RCUs of the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >
	//
	// 	- This parameter is supported only for serverless ApsaraDB RDS for MySQL instances.
	//
	// 	- This parameter is available only on the China site (aliyun.com).
	//
	// example:
	//
	// false
	SwitchForce *bool `json:"SwitchForce,omitempty" xml:"SwitchForce,omitempty"`
}

func (s CloneDBInstanceRequestServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s CloneDBInstanceRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *CloneDBInstanceRequestServerlessConfig) GetAutoPause() *bool {
	return s.AutoPause
}

func (s *CloneDBInstanceRequestServerlessConfig) GetMaxCapacity() *float64 {
	return s.MaxCapacity
}

func (s *CloneDBInstanceRequestServerlessConfig) GetMinCapacity() *float64 {
	return s.MinCapacity
}

func (s *CloneDBInstanceRequestServerlessConfig) GetSwitchForce() *bool {
	return s.SwitchForce
}

func (s *CloneDBInstanceRequestServerlessConfig) SetAutoPause(v bool) *CloneDBInstanceRequestServerlessConfig {
	s.AutoPause = &v
	return s
}

func (s *CloneDBInstanceRequestServerlessConfig) SetMaxCapacity(v float64) *CloneDBInstanceRequestServerlessConfig {
	s.MaxCapacity = &v
	return s
}

func (s *CloneDBInstanceRequestServerlessConfig) SetMinCapacity(v float64) *CloneDBInstanceRequestServerlessConfig {
	s.MinCapacity = &v
	return s
}

func (s *CloneDBInstanceRequestServerlessConfig) SetSwitchForce(v bool) *CloneDBInstanceRequestServerlessConfig {
	s.SwitchForce = &v
	return s
}

func (s *CloneDBInstanceRequestServerlessConfig) Validate() error {
	return dara.Validate(s)
}
