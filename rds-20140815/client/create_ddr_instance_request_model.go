// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDdrInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *CreateDdrInstanceRequest
	GetBackupSetId() *string
	SetBackupSetRegion(v string) *CreateDdrInstanceRequest
	GetBackupSetRegion() *string
	SetClientToken(v string) *CreateDdrInstanceRequest
	GetClientToken() *string
	SetConnectionMode(v string) *CreateDdrInstanceRequest
	GetConnectionMode() *string
	SetDBInstanceClass(v string) *CreateDdrInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateDdrInstanceRequest
	GetDBInstanceDescription() *string
	SetDBInstanceNetType(v string) *CreateDdrInstanceRequest
	GetDBInstanceNetType() *string
	SetDBInstanceStorage(v int32) *CreateDdrInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *CreateDdrInstanceRequest
	GetDBInstanceStorageType() *string
	SetEncryptionKey(v string) *CreateDdrInstanceRequest
	GetEncryptionKey() *string
	SetEngine(v string) *CreateDdrInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDdrInstanceRequest
	GetEngineVersion() *string
	SetInstanceNetworkType(v string) *CreateDdrInstanceRequest
	GetInstanceNetworkType() *string
	SetOwnerAccount(v string) *CreateDdrInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDdrInstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateDdrInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDdrInstanceRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *CreateDdrInstanceRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *CreateDdrInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDdrInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDdrInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDdrInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CreateDdrInstanceRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *CreateDdrInstanceRequest
	GetRestoreType() *string
	SetRoleARN(v string) *CreateDdrInstanceRequest
	GetRoleARN() *string
	SetSecurityIPList(v string) *CreateDdrInstanceRequest
	GetSecurityIPList() *string
	SetSourceDBInstanceName(v string) *CreateDdrInstanceRequest
	GetSourceDBInstanceName() *string
	SetSourceRegion(v string) *CreateDdrInstanceRequest
	GetSourceRegion() *string
	SetSystemDBCharset(v string) *CreateDdrInstanceRequest
	GetSystemDBCharset() *string
	SetUsedTime(v string) *CreateDdrInstanceRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateDdrInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDdrInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDdrInstanceRequest
	GetZoneId() *string
}

type CreateDdrInstanceRequest struct {
	// The backup set ID that you want to use for the restoration. You can call the DescribeCrossRegionBackups operation to query backup set ID.
	//
	// >  This parameter is required when you set the **RestoreType*	- parameter to **BackupSet**.
	//
	// example:
	//
	// 14***
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The region where the backup set is located.
	//
	// example:
	//
	// cn-beijing
	BackupSetRegion *string `json:"BackupSetRegion,omitempty" xml:"BackupSetRegion,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connection mode of the destination instance. Valid values:
	//
	// 	- **Standard**: standard mode
	//
	// 	- **Safe**: database proxy mode
	//
	// Default value: **Standard**.
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The instance type of the destination instance. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// rds.mysql.s1.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance name. The name must be 2 to 256 characters in length. The value can contain letters, digits, underscores (_), and hyphens (-), and must start with a letter.
	//
	// >  The value cannot start with http:// or https://.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The network connection type of the destination instance. Valid values:
	//
	// 	- **Internet**
	//
	// 	- **Intranet**
	//
	// This parameter is required.
	//
	// example:
	//
	// Intranet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The storage capacity of the destination instance. Valid values: **5 to 2000**. Unit: GB. You can increase the storage capacity at a step size of 5 GB. For more information, see [Primary instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the destination instance. Only the local SSD storage type is supported. Default value: **local_ssd**.
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The ID of the customer master key (CMK) for cloud disk encryption. If this parameter is specified, cloud disk encryption is enabled and you must also specify the **RoleARN*	- parameter. Cloud disk encryption cannot be disabled after it is enabled. You can obtain the ID of the key in the KMS console or create a key. For more information, see [Create a key](https://help.aliyun.com/document_detail/181610.html).
	//
	// **
	//
	// **Notes**
	//
	// 	- This parameter is applicable only to ApsaraDB RDS for SQL Server instances.
	//
	// 	- You can leave this parameter empty. If you do not specify this parameter, you only need to specify the **RoleARN*	- to use the service key that is managed by ApsaraDB RDS to encrypt cloud disks.
	//
	// example:
	//
	// 749c1df7-****-****-****-****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The database engine of the destination instance. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The major engine version of the destination instance. The value of this parameter varies based on the value of **Engine**.
	//
	// 	- Valid values when Engine is set to MySQL: **5.5, 5.6, 5.7, and 8.0**
	//
	// 	- Valid values when Engine is set to SQLServer: **2008r2, 08r2_ent_ha, 2012, 2012_ent_ha, 2012_std_ha, 2012_web, 2014_std_ha, 2016_ent_ha, 2016_std_ha, 2016_web, 2017_std_ha, 2017_ent, 2019_std_ha, and 2019_ent**
	//
	// 	- Valid values when Engine is set to PostgreSQL: **9.4, 10.0, 11.0, 12.0, and 13.0**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **Classic**
	//
	// Default value: Classic.
	//
	// > If you set this parameter to **VPC**, you must also specify **VpcId*	- and **VSwitchId**.
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit that is used to measure the subscription duration of the destination instance. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// > If you set PayType to **Prepaid**, you must specify UsedTime.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The private IP address of the destination instance. The private IP address must be within the CIDR block that is supported by the specified vSwitch. The system automatically assigns an internal IP address based on the values of the **VPCId*	- and **VSwitchId*	- parameters.
	//
	// example:
	//
	// 172.XXX.XXX.69
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID of the destination instance. You can call the DescribeRegions operation to query the most recent region list.
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
	// rg-acfmyxxxxxxxxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore data. The point in time that you specify must be earlier than the current time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > If **RestoreType*	- is set to **BackupTime**, you must specify this parameter.
	//
	// example:
	//
	// 2019-05-30T03:29:10Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restoration method that you want to use. Valid values:
	//
	// 	- **BackupSet**: restores data from a backup set. If you use this value, you must also specify **BackupSetId**.
	//
	// 	- **BackupTime**: restores data to a point in time. If you use this value, you must also specify **RestoreTime**, **SourceRegion**, and **SourceDBInstanceName**.
	//
	// This parameter is required.
	//
	// example:
	//
	// BackupSet
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) that is provided by your Alibaba Cloud account for Resource Access Management (RAM) users. RAM users can use the ARN to connect to ApsaraDB RDS to Key Management Service (KMS). You can call the [CheckCloudResourceAuthorized](https://help.aliyun.com/document_detail/2628797.html) operation to query the ARN.
	//
	// >  This parameter is applicable only to ApsaraDB RDS for SQL Server instances.
	//
	// example:
	//
	// acs:ram::1406****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The IP address whitelist of the destination instance. If you want to add more than one entry to the IP address whitelist, separate the entries with commas (,). Each entry must be unique. You can add a maximum of 1,000 entries. For more information, see [Configure an IP address whitelist for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/43185.html). The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP address. Example: 10.23.12.24.
	//
	// 	- CIDR block. Example: 10.23.12.24/24. In this example, 24 indicates that the prefix of the CIDR block is 24 bits in length. You can replace 24 with a value that ranges from 1 to 32.
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The source instance ID, which is used if you want to restore data to a point in time.
	//
	// >  This parameter is required when you set the **RestoreType*	- parameter to **BackupTime**.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	// The region ID of the source instance if you want to restore data to a point in time.
	//
	// > If you set **RestoreType*	- to **BackupTime**, you must specify this parameter.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The character set of the destination instance. Valid values:
	//
	// 	- **utf8**
	//
	// 	- **gbk**
	//
	// 	- **latin1**
	//
	// 	- **utf8mb4**
	//
	// example:
	//
	// uft8
	SystemDBCharset *string `json:"SystemDBCharset,omitempty" xml:"SystemDBCharset,omitempty"`
	// The subscription duration of the instance.
	//
	// 	- If you set **Period*	- to **Year**, the value of UsedTime ranges from **1 to 3**.
	//
	// 	- If you set **Period*	- to **Month**, the value of UsedTime ranges from **1 to 9**.
	//
	// > If you set PayType to **Prepaid**, you must specify UsedTime.
	//
	// example:
	//
	// 2
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID of the destination instance. This parameter is available only when you set the **InstanceNetworkType*	- parameter to **VPC**.
	//
	// >  If you specify this parameter, you must also specify the **ZoneId*	- parameter.
	//
	// example:
	//
	// vpc-xxxxxxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the destination instance. If you specify more than one vSwitch, separate the IDs of the vSwitches with commas (,). This parameter is available only when you set the **InstanceNetworkType*	- parameter to **VPC**.
	//
	// >  If you specify this parameter, you must also specify the **ZoneId*	- parameter.
	//
	// example:
	//
	// vsw-xxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the destination instance. If the destination instance is deployed in multiple zones, separate the IDs of the zones with colons (:).
	//
	// > If you specify a virtual private cloud (VPC) and a vSwitch, you must specify this parameter to identify the zone for the vSwitch.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDdrInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDdrInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDdrInstanceRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateDdrInstanceRequest) GetBackupSetRegion() *string {
	return s.BackupSetRegion
}

func (s *CreateDdrInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDdrInstanceRequest) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *CreateDdrInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateDdrInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDdrInstanceRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *CreateDdrInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CreateDdrInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CreateDdrInstanceRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateDdrInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDdrInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDdrInstanceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CreateDdrInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDdrInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDdrInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDdrInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDdrInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateDdrInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDdrInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDdrInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDdrInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDdrInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateDdrInstanceRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CreateDdrInstanceRequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *CreateDdrInstanceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDdrInstanceRequest) GetSourceDBInstanceName() *string {
	return s.SourceDBInstanceName
}

func (s *CreateDdrInstanceRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *CreateDdrInstanceRequest) GetSystemDBCharset() *string {
	return s.SystemDBCharset
}

func (s *CreateDdrInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDdrInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDdrInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDdrInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDdrInstanceRequest) SetBackupSetId(v string) *CreateDdrInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetBackupSetRegion(v string) *CreateDdrInstanceRequest {
	s.BackupSetRegion = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetClientToken(v string) *CreateDdrInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetConnectionMode(v string) *CreateDdrInstanceRequest {
	s.ConnectionMode = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceClass(v string) *CreateDdrInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceDescription(v string) *CreateDdrInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceNetType(v string) *CreateDdrInstanceRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceStorage(v int32) *CreateDdrInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetDBInstanceStorageType(v string) *CreateDdrInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetEncryptionKey(v string) *CreateDdrInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetEngine(v string) *CreateDdrInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetEngineVersion(v string) *CreateDdrInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetInstanceNetworkType(v string) *CreateDdrInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetOwnerAccount(v string) *CreateDdrInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetOwnerId(v int64) *CreateDdrInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetPayType(v string) *CreateDdrInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetPeriod(v string) *CreateDdrInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetPrivateIpAddress(v string) *CreateDdrInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetRegionId(v string) *CreateDdrInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetResourceGroupId(v string) *CreateDdrInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetResourceOwnerAccount(v string) *CreateDdrInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetResourceOwnerId(v int64) *CreateDdrInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetRestoreTime(v string) *CreateDdrInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetRestoreType(v string) *CreateDdrInstanceRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetRoleARN(v string) *CreateDdrInstanceRequest {
	s.RoleARN = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetSecurityIPList(v string) *CreateDdrInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetSourceDBInstanceName(v string) *CreateDdrInstanceRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetSourceRegion(v string) *CreateDdrInstanceRequest {
	s.SourceRegion = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetSystemDBCharset(v string) *CreateDdrInstanceRequest {
	s.SystemDBCharset = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetUsedTime(v string) *CreateDdrInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetVPCId(v string) *CreateDdrInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetVSwitchId(v string) *CreateDdrInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDdrInstanceRequest) SetZoneId(v string) *CreateDdrInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDdrInstanceRequest) Validate() error {
	return dara.Validate(s)
}
