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
	// The ID of the backup set used for restoration from a backup set. You can call the DescribeCrossRegionBackups operation to query backup set IDs.
	//
	// > This parameter is required when **RestoreType*	- is set to **BackupSet**.
	//
	// example:
	//
	// 14****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The region where the backup set resides.
	//
	// example:
	//
	// cn-beijing
	BackupSetRegion *string `json:"BackupSetRegion,omitempty" xml:"BackupSetRegion,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The access mode of the target instance. Valid values:
	//
	// - **Standard*	- (default): standard access mode
	//
	// - **Safe**: database proxy mode
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The instance type of the target instance. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// rds.mysql.s1.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The name of the target instance. The name must be 2 to 256 characters in length. The name must start with a letter or a Chinese character and can contain digits, Chinese characters, letters, underscores (_), and hyphens (-).
	//
	// > The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testdb
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The network connectivity type of the target instance. Valid values:
	//
	// 	- **Internet**: public network connection
	//
	// 	- **Intranet**: internal network connection
	//
	// This parameter is required.
	//
	// example:
	//
	// Intranet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The instance storage capacity of the target instance. Valid values: **5 to 2000**. The value is incremented in steps of 5 GB. Unit: GB. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type of the target instance. Valid values:
	//
	// > Use the same storage type as the source instance.
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for MySQL</summary>
	//
	// - local_ssd: Premium Local SSDs (default)
	//
	// - cloud_essd: PL1 ESSD cloud disk
	//
	// - cloud_essd2: PL2 ESSD cloud disk
	//
	// - cloud_essd3: PL3 ESSD cloud disk
	//
	// - cloud_ssd: standard SSD cloud disk (discontinued)
	//
	// </details>
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for SQL Server</summary>
	//
	// - cloud_essd: PL1 ESSD cloud disk
	//
	// - cloud_essd2: PL2 ESSD cloud disk
	//
	// - cloud_essd3: PL3 ESSD cloud disk
	//
	// - local_ssd: Premium Local SSDs (discontinued)
	//
	// - cloud_ssd: standard SSD cloud disk (discontinued)
	//
	// </details>
	//
	// <details>
	//
	// <summary>ApsaraDB RDS for PostgreSQL</summary>
	//
	// - cloud_essd: PL1 ESSD cloud disk
	//
	// - cloud_essd2: PL2 ESSD cloud disk
	//
	// - cloud_essd3: PL3 ESSD cloud disk
	//
	// - local_ssd: Premium Local SSDs (discontinued)
	//
	// - cloud_ssd: standard SSD cloud disk (discontinued)
	//
	// </details>
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The ID of the custom key used for cloud disk encryption for **SQL Server instances**. Specifying this parameter enables cloud disk encryption (which cannot be disabled after it is enabled). You must also specify **RoleARN**.
	//
	// You can view the key ID in the Key Management Service (KMS) console or [create a new key](https://help.aliyun.com/document_detail/181610.html).
	//
	// > You can also leave this parameter empty and specify only **RoleARN*	- to set the cloud disk encryption type to the RDS-managed service key (Default Service CMK).
	//
	// example:
	//
	// 749c1df7-****-****-****-****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The type of the destination database engine. Valid values:
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
	// The version of the destination database engine. The valid values vary based on the value of **Engine**:
	//
	// - MySQL: **5.5/5.6/5.7/8.0**
	//
	// - SQL Server: **2008r2 (Premium Local SSDs, discontinued)/08r2_ent_ha (cloud disks, discontinued)/2012/2012_ent_ha/2012_std_ha/2012_web/2014_std_ha/2016_ent_ha/2016_std_ha/2016_web/2017_std_ha/2017_ent/2019_std_ha/2019_ent**
	//
	// - PostgreSQL: **10.0/11.0/12.0/13.0/14.0/15.0**
	//
	// > For SQL Server instances, `_ent` indicates Cluster Edition, `_ent_ha` indicates Enterprise Edition, `_std_ha` indicates Standard Edition, and `_web` indicates Web Edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The network type of the target instance. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **Classic**: classic network (offline)
	//
	// > If you set this parameter to **VPC**, you must also specify the **VpcId*	- and **VSwitchId*	- parameters.
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the target instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: upfront (subscription)
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the upfront subscription duration for the target instance. Valid values:
	//
	// 	- **Year**: yearly subscription
	//
	// 	- **Month**: monthly subscription
	//
	// > This parameter is required when PayType is set to **Prepaid**.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// Settings for the internal network IP address of the target instance. The IP address must be within the IP address range of the specified vSwitch. By default, the system automatically allocates an internal network IP address based on the values of **VPCId*	- and **VSwitchId**.
	//
	// example:
	//
	// 172.XX.XX.69
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the destination region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query region IDs.
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
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore data when you restore data to a point in time. The point in time must be earlier than the current time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > This parameter is required when **RestoreType*	- is set to **BackupTime**.
	//
	// example:
	//
	// 2019-05-30T03:29:10Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restoration method. Valid values:
	//
	// - **BackupSet**: restores data from a backup set. The data in the backup set is restored to the new instance. You must also specify the **BackupSetId*	- parameter.
	//
	// - **BackupTime**: restores data to a point in time within the log backup retention period. You must also specify the **RestoreTime**, **SourceRegion**, and **SourceDBInstanceName*	- parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// BackupSet
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The global resource descriptor (ARN) that provides authorization for the RDS cloud service account to access Key Management Service (KMS) for **SQL Server instances**. You can call the [CheckCloudResourceAuthorized](https://help.aliyun.com/document_detail/2628797.html) operation to query the ARN.
	//
	// example:
	//
	// acs:ram::1406****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The [IP whitelist](https://help.aliyun.com/document_detail/43185.html) of the target instance. Separate multiple IP addresses with commas (,). IP addresses cannot be duplicated. You can specify up to 1,000 IP addresses. The following two formats are supported:
	//
	// 	- IP address format, such as 10.23.12.24.
	//
	// 	- CIDR format, such as 10.23.12.24/24 (Classless Inter-Domain Routing. 24 indicates the length of the prefix in the address. The value ranges from 1 to 32).
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The ID of the source instance for point-in-time restoration.
	//
	// > This parameter is required when **RestoreType*	- is set to **BackupTime**.
	//
	// example:
	//
	// rm-uf6wjk5****
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	// The ID of the source region for point-in-time restoration.
	//
	// > This parameter is required when **RestoreType*	- is set to **BackupTime**.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The character set of the target instance. Valid values:
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
	// The subscription duration. Valid values:
	//
	// 	- If **Period*	- is set to **Year**, the valid values of UsedTime are **1 to 3**.
	//
	// 	- If **Period*	- is set to **Month**, the valid values of UsedTime are **1 to 9**.
	//
	// > This parameter is required when PayType is set to **Prepaid**.
	//
	// example:
	//
	// 2
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID of the target instance.
	//
	// > - This parameter is required when **InstanceNetworkType*	- is set to **VPC**.
	//
	// > - If you specify this parameter, you must also specify the **ZoneId*	- parameter.
	//
	// example:
	//
	// vpc-****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the target instance. Separate multiple values with commas (,).
	//
	// > - This parameter is required when **InstanceNetworkType*	- is set to **VPC**.
	//
	// > - If you specify this parameter, you must also specify the **ZoneId*	- parameter.
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The active zone ID of the target instance. Separate multiple zones with colons (:).
	//
	// > If you specify a VPC and a vSwitch, this parameter is required to match the zone of the specified vSwitch.
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
