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
	SetTag(v []*CloneDBInstanceRequestTag) *CloneDBInstanceRequest
	GetTag() []*CloneDBInstanceRequestTag
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
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 1. **true**: enables automatic payment. Make sure that your account balance is sufficient.
	//
	// 1. **false**: generates an unpaid order without deducting fees.
	//
	// > Default value: true. If your payment method has an insufficient balance, you can set the AutoPay parameter to false. In this case, an unpaid order is generated, and you can log on to the ApsaraDB RDS console to pay the order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the backup set.
	//
	// You can call the DescribeBackups operation to query backup sets.
	//
	// >You must specify at least one of the **BackupId*	- and **RestoreTime*	- parameters.
	//
	// example:
	//
	// 902****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **FullBackup**: full backup
	//
	// 	- **IncrementalBackup**: incremental backup
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BpeEnabled *string `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	// Specifies whether to enable the I/O performance burst feature for cloud disks. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
	//
	// > For more information about the I/O performance burst feature, see [What is Premium ESSD?](https://help.aliyun.com/document_detail/2340501.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The instance edition. Valid values:
	//
	// - **Basic**: Basic Edition
	//
	// - **HighAvailability**: High-availability Edition
	//
	// - **AlwaysOn**: Cluster Edition (SQL Server)
	//
	// - **cluster**: Cluster Edition (MySQL)
	//
	// - **Finance**: Finance Edition (supported only on the China site)
	//
	// **Serverless instances**
	//
	// - **serverless_basic**: Serverless Basic Edition (supported only for MySQL and PostgreSQL)
	//
	// - **serverless_standard**: MySQL Serverless High-availability Edition
	//
	// - **serverless_ha**: SQL Server Serverless High-availability Edition
	//
	// > You do not need to specify this parameter. The cloned instance inherits the edition of the source instance.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Ensures the idempotence of the request and prevents duplicate submissions. The value is generated by the client and must be unique across different requests. The value can be up to 64 ASCII characters in length and cannot contain non-ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// >Default value: the instance type of the source instance.
	//
	// example:
	//
	// mysql.n1.micro.1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the instance. The description must be 2 to 255 characters in length. It must start with a Chinese character or a letter and can contain digits, Chinese characters, letters, underscores (_), and hyphens (-).
	//
	// >The description cannot start with http:// or https://.
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
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Instance storage capacity of the instance. Unit: GB. Instance storage capacity increases in increments of 5 GB. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// >Default value: instance storage capacity of the source instance.
	//
	// example:
	//
	// 1000
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **general_essd**: Premium ESSD (recommended)
	//
	// 	- **local_ssd**: local SSD
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: PL1 ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// > Serverless instances support only PL1 ESSD and Premium ESSD.
	//
	// example:
	//
	// general_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The names of the databases to restore. Format: `OriginalDatabaseName1,OriginalDatabaseName2`.
	//
	// example:
	//
	// test1,test2
	DbNames *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	// The ID of the dedicated cluster.
	//
	// example:
	//
	// dhg-7a9****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to enable release protection. Valid values:
	//
	// 	- **true**: enables release protection.
	//
	// 	- **false*	- (default): disables release protection.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: Virtual Private Cloud (VPC)
	//
	// 	- **Classic**: classic network
	//
	// >Default value: the network type of the source instance.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// Specifies whether to enable the Buffer Pool Extension (BPE) feature for cloud disks. Valid values:
	//
	//  - **1**: enables the feature.
	//
	//  - **0**: disables the feature.
	//
	// > For more information about the BPE feature, see [Buffer Pool Extension (BPE)](https://help.aliyun.com/document_detail/2527067.html).
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// 	- **Serverless**: serverless. This billing method is not supported for ApsaraDB RDS for MariaDB instances. For more information, see [Overview of serverless ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/411291.html), [Overview of serverless ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/604344.html), and [Overview of serverless ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/607742.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription period of the subscription instance. Valid values:
	//
	// 	- **Year**: yearly
	//
	// 	- **Month**: monthly
	//
	// > This parameter is required if the **PayType*	- parameter is set to **Prepaid**.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The internal IP address of the new instance. The IP address must be within the IP address range of the specified vSwitch. By default, the system automatically allocates an IP address based on the **VPCId*	- and **VSwitchId*	- parameters.
	//
	// example:
	//
	// 172.XX.XX.69
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to restore individual databases and tables. Set the value to **true*	- to restore individual databases and tables. Otherwise, leave this parameter empty.
	//
	// example:
	//
	// true
	RestoreTable *string `json:"RestoreTable,omitempty" xml:"RestoreTable,omitempty"`
	// The point in time to which you want to restore data. The point in time must be within the backup retention period. Specify the time in the ISO 8601 standard in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// >You must specify at least one of the **BackupId*	- and **RestoreTime*	- parameters.
	//
	// example:
	//
	// 2011-06-11T16:00:00Z
	RestoreTime      *string                                 `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	ServerlessConfig *CloneDBInstanceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The information about the databases and tables to restore. This parameter is required when you restore individual databases and tables. Format:
	//
	// ```[{"type":"db","name":"Original database 1 name","newname":"New database 1 name","tables":[{"type":"table","name":"Table 1 name in database 1","newname":"New table 1 name"},{"type":"table","name":"Table 2 name in database 1","newname":"New table 2 name"}]},{"type":"db","name":"Original database 2 name","newname":"New database 2 name","tables":[{"type":"table","name":"Table 1 name in database 2","newname":"New table 1 name"},{"type":"table","name":"Table 2 name in database 2","newname":"New table 2 name"}]}]```
	//
	// example:
	//
	// [{"type":"db","name":"testdb1","newname":"testdb1_new","tables":[{"type":"table","name":"testdb1table1","newname":"testdb1table1_new"}]}]
	TableMeta *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
	// The list of tags.
	Tag []*CloneDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The subscription duration. Valid values:
	//
	// 	- If the **Period*	- parameter is set to **Year**, the value of the UsedTime parameter ranges from **1*	- to **3**.
	//
	// 	- If the **Period*	- parameter is set to **Month**, the value of the UsedTime parameter ranges from **1*	- to **9**.
	//
	// > This parameter is required if the **PayType*	- parameter is set to **Prepaid**.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The ID of the VPC.
	//
	// >Make sure that the VPC belongs to the specified region.
	//
	// example:
	//
	// vpc-uf6f7l4fg90****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID. The zone where the vSwitch resides must correspond to the zone ID specified in **ZoneId**.
	//
	// - The network type **InstanceNetworkType*	- must be set to **VPC**.
	//
	// - If you specify **ZoneSlaveId1*	- (secondary zone ID), you must specify two vSwitch IDs and separate them with a comma (,).
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The primary zone ID. You can call the DescribeRegions operation to query zone IDs.
	//
	// >Default value: the zone ID of the source instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary node. If this parameter is set to the same value as the **ZoneId*	- parameter, it is a single-zone deployment. If this parameter is set to a different value from the **ZoneId*	- parameter, it is a multi-zone deployment.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// <props="intl">The zone ID of the logger node. If this parameter is set to the same value as the **ZoneId*	- parameter, it is a single-zone deployment. If this parameter is set to a different value from the **ZoneId*	- parameter, it is a multi-zone deployment.
	//
	// <props="china">The zone ID of the secondary node or logger node. If this parameter is set to the same value as the **ZoneId*	- parameter, it is a single-zone deployment. If this parameter is set to a different value from the **ZoneId*	- parameter, it is a multi-zone deployment.
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

func (s *CloneDBInstanceRequest) GetTag() []*CloneDBInstanceRequestTag {
	return s.Tag
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

func (s *CloneDBInstanceRequest) SetTag(v []*CloneDBInstanceRequestTag) *CloneDBInstanceRequest {
	s.Tag = v
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

type CloneDBInstanceRequestServerlessConfig struct {
	AutoPause   *bool    `json:"AutoPause,omitempty" xml:"AutoPause,omitempty"`
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	MinCapacity *float64 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
	SwitchForce *bool    `json:"SwitchForce,omitempty" xml:"SwitchForce,omitempty"`
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

type CloneDBInstanceRequestTag struct {
	// The tag key. You can specify this parameter to bind a tag to the instance.
	//
	// 	- If the specified tag key already exists, the tag key is directly bound to the instance. You can call the ListTagResources operation to query created tags.
	//
	// 	- If the specified tag key does not exist, the tag key is created and then bound to the instance.
	//
	// 	- The tag key cannot be an empty string.
	//
	// 	- This parameter must be used together with the **Tag.Value*	- parameter.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key. You can specify this parameter to bind a tag to the instance.
	//
	// 	- If the specified tag value already exists for the corresponding tag key, the tag value is directly bound to the instance. You can call the ListTagResources operation to query created tags.
	//
	// 	- If the specified tag value does not exist for the corresponding tag key, the tag value is created and then bound to the instance.
	//
	// 	- This parameter must be used together with the **Tag.Key*	- parameter.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CloneDBInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CloneDBInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CloneDBInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CloneDBInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CloneDBInstanceRequestTag) SetKey(v string) *CloneDBInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CloneDBInstanceRequestTag) SetValue(v string) *CloneDBInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CloneDBInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
