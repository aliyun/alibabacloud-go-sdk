// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDBInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CloneDBInstanceShrinkRequest
	GetAutoPay() *bool
	SetBackupId(v string) *CloneDBInstanceShrinkRequest
	GetBackupId() *string
	SetBackupType(v string) *CloneDBInstanceShrinkRequest
	GetBackupType() *string
	SetBpeEnabled(v string) *CloneDBInstanceShrinkRequest
	GetBpeEnabled() *string
	SetBurstingEnabled(v bool) *CloneDBInstanceShrinkRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *CloneDBInstanceShrinkRequest
	GetCategory() *string
	SetClientToken(v string) *CloneDBInstanceShrinkRequest
	GetClientToken() *string
	SetCustomExtraInfo(v string) *CloneDBInstanceShrinkRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *CloneDBInstanceShrinkRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CloneDBInstanceShrinkRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *CloneDBInstanceShrinkRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *CloneDBInstanceShrinkRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *CloneDBInstanceShrinkRequest
	GetDBInstanceStorageType() *string
	SetDbNames(v string) *CloneDBInstanceShrinkRequest
	GetDbNames() *string
	SetDedicatedHostGroupId(v string) *CloneDBInstanceShrinkRequest
	GetDedicatedHostGroupId() *string
	SetDeletionProtection(v bool) *CloneDBInstanceShrinkRequest
	GetDeletionProtection() *bool
	SetInstanceNetworkType(v string) *CloneDBInstanceShrinkRequest
	GetInstanceNetworkType() *string
	SetIoAccelerationEnabled(v string) *CloneDBInstanceShrinkRequest
	GetIoAccelerationEnabled() *string
	SetPayType(v string) *CloneDBInstanceShrinkRequest
	GetPayType() *string
	SetPeriod(v string) *CloneDBInstanceShrinkRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *CloneDBInstanceShrinkRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *CloneDBInstanceShrinkRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *CloneDBInstanceShrinkRequest
	GetResourceOwnerId() *int64
	SetRestoreTable(v string) *CloneDBInstanceShrinkRequest
	GetRestoreTable() *string
	SetRestoreTime(v string) *CloneDBInstanceShrinkRequest
	GetRestoreTime() *string
	SetServerlessConfigShrink(v string) *CloneDBInstanceShrinkRequest
	GetServerlessConfigShrink() *string
	SetTableMeta(v string) *CloneDBInstanceShrinkRequest
	GetTableMeta() *string
	SetUsedTime(v int32) *CloneDBInstanceShrinkRequest
	GetUsedTime() *int32
	SetVPCId(v string) *CloneDBInstanceShrinkRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CloneDBInstanceShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CloneDBInstanceShrinkRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *CloneDBInstanceShrinkRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *CloneDBInstanceShrinkRequest
	GetZoneIdSlave2() *string
}

type CloneDBInstanceShrinkRequest struct {
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
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
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

func (s CloneDBInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CloneDBInstanceShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CloneDBInstanceShrinkRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CloneDBInstanceShrinkRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *CloneDBInstanceShrinkRequest) GetBpeEnabled() *string {
	return s.BpeEnabled
}

func (s *CloneDBInstanceShrinkRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CloneDBInstanceShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *CloneDBInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CloneDBInstanceShrinkRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *CloneDBInstanceShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CloneDBInstanceShrinkRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CloneDBInstanceShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CloneDBInstanceShrinkRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CloneDBInstanceShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CloneDBInstanceShrinkRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *CloneDBInstanceShrinkRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *CloneDBInstanceShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CloneDBInstanceShrinkRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CloneDBInstanceShrinkRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *CloneDBInstanceShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *CloneDBInstanceShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CloneDBInstanceShrinkRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CloneDBInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloneDBInstanceShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloneDBInstanceShrinkRequest) GetRestoreTable() *string {
	return s.RestoreTable
}

func (s *CloneDBInstanceShrinkRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CloneDBInstanceShrinkRequest) GetServerlessConfigShrink() *string {
	return s.ServerlessConfigShrink
}

func (s *CloneDBInstanceShrinkRequest) GetTableMeta() *string {
	return s.TableMeta
}

func (s *CloneDBInstanceShrinkRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CloneDBInstanceShrinkRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CloneDBInstanceShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CloneDBInstanceShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CloneDBInstanceShrinkRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *CloneDBInstanceShrinkRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *CloneDBInstanceShrinkRequest) SetAutoPay(v bool) *CloneDBInstanceShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetBackupId(v string) *CloneDBInstanceShrinkRequest {
	s.BackupId = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetBackupType(v string) *CloneDBInstanceShrinkRequest {
	s.BackupType = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetBpeEnabled(v string) *CloneDBInstanceShrinkRequest {
	s.BpeEnabled = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetBurstingEnabled(v bool) *CloneDBInstanceShrinkRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetCategory(v string) *CloneDBInstanceShrinkRequest {
	s.Category = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetClientToken(v string) *CloneDBInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetCustomExtraInfo(v string) *CloneDBInstanceShrinkRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetDBInstanceClass(v string) *CloneDBInstanceShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetDBInstanceDescription(v string) *CloneDBInstanceShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetDBInstanceId(v string) *CloneDBInstanceShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetDBInstanceStorage(v int32) *CloneDBInstanceShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetDBInstanceStorageType(v string) *CloneDBInstanceShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetDbNames(v string) *CloneDBInstanceShrinkRequest {
	s.DbNames = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetDedicatedHostGroupId(v string) *CloneDBInstanceShrinkRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetDeletionProtection(v bool) *CloneDBInstanceShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetInstanceNetworkType(v string) *CloneDBInstanceShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetIoAccelerationEnabled(v string) *CloneDBInstanceShrinkRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetPayType(v string) *CloneDBInstanceShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetPeriod(v string) *CloneDBInstanceShrinkRequest {
	s.Period = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetPrivateIpAddress(v string) *CloneDBInstanceShrinkRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetRegionId(v string) *CloneDBInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetResourceOwnerId(v int64) *CloneDBInstanceShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetRestoreTable(v string) *CloneDBInstanceShrinkRequest {
	s.RestoreTable = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetRestoreTime(v string) *CloneDBInstanceShrinkRequest {
	s.RestoreTime = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetServerlessConfigShrink(v string) *CloneDBInstanceShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetTableMeta(v string) *CloneDBInstanceShrinkRequest {
	s.TableMeta = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetUsedTime(v int32) *CloneDBInstanceShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetVPCId(v string) *CloneDBInstanceShrinkRequest {
	s.VPCId = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetVSwitchId(v string) *CloneDBInstanceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetZoneId(v string) *CloneDBInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetZoneIdSlave1(v string) *CloneDBInstanceShrinkRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) SetZoneIdSlave2(v string) *CloneDBInstanceShrinkRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *CloneDBInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
