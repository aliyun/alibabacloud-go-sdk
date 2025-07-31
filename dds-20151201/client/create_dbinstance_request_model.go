// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountPassword(v string) *CreateDBInstanceRequest
	GetAccountPassword() *string
	SetAutoRenew(v string) *CreateDBInstanceRequest
	GetAutoRenew() *string
	SetBackupId(v string) *CreateDBInstanceRequest
	GetBackupId() *string
	SetBusinessInfo(v string) *CreateDBInstanceRequest
	GetBusinessInfo() *string
	SetChargeType(v string) *CreateDBInstanceRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateDBInstanceRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateDBInstanceRequest
	GetClusterId() *string
	SetCouponNo(v string) *CreateDBInstanceRequest
	GetCouponNo() *string
	SetDBInstanceClass(v string) *CreateDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceRequest
	GetDBInstanceDescription() *string
	SetDBInstanceStorage(v int32) *CreateDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetDatabaseNames(v string) *CreateDBInstanceRequest
	GetDatabaseNames() *string
	SetEncrypted(v bool) *CreateDBInstanceRequest
	GetEncrypted() *bool
	SetEncryptionKey(v string) *CreateDBInstanceRequest
	GetEncryptionKey() *string
	SetEngine(v string) *CreateDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateDBInstanceRequest
	GetEngineVersion() *string
	SetGlobalSecurityGroupIds(v string) *CreateDBInstanceRequest
	GetGlobalSecurityGroupIds() *string
	SetHiddenZoneId(v string) *CreateDBInstanceRequest
	GetHiddenZoneId() *string
	SetNetworkType(v string) *CreateDBInstanceRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBInstanceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateDBInstanceRequest
	GetPeriod() *int32
	SetProvisionedIops(v int64) *CreateDBInstanceRequest
	GetProvisionedIops() *int64
	SetReadonlyReplicas(v string) *CreateDBInstanceRequest
	GetReadonlyReplicas() *string
	SetRegionId(v string) *CreateDBInstanceRequest
	GetRegionId() *string
	SetReplicationFactor(v string) *CreateDBInstanceRequest
	GetReplicationFactor() *string
	SetResourceGroupId(v string) *CreateDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CreateDBInstanceRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *CreateDBInstanceRequest
	GetRestoreType() *string
	SetSecondaryZoneId(v string) *CreateDBInstanceRequest
	GetSecondaryZoneId() *string
	SetSecurityIPList(v string) *CreateDBInstanceRequest
	GetSecurityIPList() *string
	SetSrcDBInstanceId(v string) *CreateDBInstanceRequest
	GetSrcDBInstanceId() *string
	SetSrcRegion(v string) *CreateDBInstanceRequest
	GetSrcRegion() *string
	SetStorageEngine(v string) *CreateDBInstanceRequest
	GetStorageEngine() *string
	SetStorageType(v string) *CreateDBInstanceRequest
	GetStorageType() *string
	SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest
	GetTag() []*CreateDBInstanceRequestTag
	SetVSwitchId(v string) *CreateDBInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateDBInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateDBInstanceRequest
	GetZoneId() *string
}

type CreateDBInstanceRequest struct {
	// The password of the root account. The password must meet the following requirements:
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- Special characters include ! # $ % ^ & \\	- ( ) _ + - =
	//
	// 	- The password of the account must be 8 to 32 characters in length.
	//
	// example:
	//
	// 123456Aa
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Default value: false. Valid values:
	//
	// 	- **true**: The instance is automatically renewed.
	//
	// 	- **false**: The instance is manually renewed.
	//
	// > This parameter is valid and optional when the **ChargeType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/62172.html) operation to query the backup set ID.
	//
	// > When you call this operation to clone an instance based on the backup set, this parameter is required. The **SrcDBInstanceId*	- parameter is also required.
	//
	// example:
	//
	// 32994****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The business information. This is an additional parameter.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PostPaid**: pay-as-you-go. This is the default value.
	//
	// 	- **PrePaid**: subscription.
	//
	// > If you set this parameter to **PrePaid**, you must also specify the **Period*	- parameter.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the dedicated cluster to which the instance belongs.
	//
	// example:
	//
	// dhg-2x78****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to use coupons. Default value: null. Valid values:
	//
	// - **default*	- or **null**: uses coupons.
	//
	// - **youhuiquan_promotion_option_id_for_blank**: does not use coupons.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance type. You can also call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to query the instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds.mongo.standard
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The name of the instance. The name of the instance must meet the following requirements:
	//
	// 	- The name must start with a letter.
	//
	// 	- The name can contain digits, letters, underscores (_), and hyphens (-).
	//
	// 	- The name must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The storage capacity of the instance. Unit: GB.
	//
	// The values that can be specified for this parameter vary based on the instance types. For more information, see [Replica set instance types](https://help.aliyun.com/document_detail/311410.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The name of the database.
	//
	// > When you call this operation to clone an instance, you can set this parameter to specify the database to clone. Otherwise, all databases of the instance are cloned.
	//
	// example:
	//
	// mongodbtest
	DatabaseNames *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// true
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the custom key.
	//
	// example:
	//
	// 2axxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The database engine of the instance. Set the value to **MongoDB**.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine. Valid values:
	//
	// 	- **6.0**
	//
	// 	- **5.0**
	//
	// 	- **4.4**
	//
	// 	- **4.2**
	//
	// 	- **4.0**
	//
	// > When you call this operation to clone an instance or restore an instance from the recycle bin, set the value of this parameter to the engine version of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The global IP address whitelist template name of the instance. Multiple IP address whitelist template names are separated by commas (,) and each template name must be unique. (The template feature is available only in canary release.)
	//
	// example:
	//
	// g-qxieqf40xjst1ngp****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The zone where the hidden node resides for multi-zone deployment. Valid values:
	//
	// 	- **cn-hangzhou-g**: Hangzhou Zone G.
	//
	// 	- **cn-hangzhou-h**: Hangzhou Zone H.
	//
	// 	- **cn-hangzhou-i**: Hangzhou Zone I.
	//
	// 	- **cn-hongkong-b**: Hongkong Zone B.
	//
	// 	- **cn-hongkong-c**: Hongkong Zone C.
	//
	// 	- **cn-hongkong-d**: Hongkong Zone D.
	//
	// 	- **cn-wulanchabu-a**: Ulanqab Zone A.
	//
	// 	- **cn-wulanchabu-b**: Ulanqab Zone B.
	//
	// 	- **cn-wulanchabu-c**: Ulanqab Zone C.
	//
	// 	- **ap-southeast-1a**: Singapore Zone A.
	//
	// 	- **ap-southeast-1b**: Singapore Zone B.
	//
	// 	- **ap-southeast-1c**: Singapore Zone C.
	//
	// 	- **ap-southeast-5a**: Jakarta Zone A.
	//
	// 	- **ap-southeast-5b**: Jakarta Zone B.
	//
	// 	- **ap-southeast-5c**: Jakarta Zone C.
	//
	// 	- **eu-central-1a**: Frankfurt Zone A.
	//
	// 	- **eu-central-1b**: Frankfurt Zone B.
	//
	// 	- **eu-central-1c**: Frankfurt Zone C.
	//
	// >  	- This parameter is valid and required when the **EngineVersion*	- parameter is set to **4.4*	- or **5.0**.
	//
	// >  	- The value of this parameter cannot be the same as the value of the **ZoneId*	- or **SecondaryZoneId*	- parameter.
	//
	// example:
	//
	// cn-hangzhou-i
	HiddenZoneId *string `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	// The network type of the instance. Valid value:
	//
	// **VPC**: Virtual Private Cloud (VPC)
	//
	// example:
	//
	// VPC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription period of the instance. Unit: months.
	//
	// Valid values: **1*	- to **9**, **12**, **24**, **36**, and **60**.
	//
	// > When you set the **ChargeType*	- parameter to **PrePaid**, this parameter is valid and required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The provisioned IOPS. Valid values: 0 to 50000.
	//
	// example:
	//
	// 1960
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The number of **read-only nodes*	- in the replica set instance. Default value: **0**. Valid values: **0*	- to **5**.
	//
	// example:
	//
	// 0
	ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of **nodes*	- in the replica set instance. Default value: 3. Valid values:
	//
	// 	- **3**
	//
	// 	- **5**
	//
	// 	- **7**
	//
	// example:
	//
	// 3
	ReplicationFactor *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which the instance is restored, which must be within seven days. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format (UTC time).
	//
	// > When you call this operation to restore an instance to the specified time, this parameter is required. The **SrcDBInstanceId*	- parameter is also required.
	//
	// example:
	//
	// 2022-03-13T12:11:14Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The backup restore type of the instance.
	//
	// - 0: restore an instance to the specified backup set.
	//
	// - 1:  restore an instance to the specified time.
	//
	// - 2: restore an  released instance to the specified backup set.
	//
	// - 3：restore an instance to the specified cross-regional backup set.
	//
	// example:
	//
	// 0
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The zone where the secondary node resides for multi-zone deployment. Valid values:
	//
	// 	- **cn-hangzhou-g**: Hangzhou Zone G.
	//
	// 	- **cn-hangzhou-h**: Hangzhou Zone H.
	//
	// 	- **cn-hangzhou-i**: Hangzhou Zone I.
	//
	// 	- **cn-hongkong-b**: Hongkong Zone B.
	//
	// 	- **cn-hongkong-c**: Hongkong Zone C.
	//
	// 	- **cn-hongkong-d**: Hongkong Zone D.
	//
	// 	- **cn-wulanchabu-a**: Ulanqab Zone A.
	//
	// 	- **cn-wulanchabu-b**: Ulanqab Zone B.
	//
	// 	- **cn-wulanchabu-c**: Ulanqab Zone C.
	//
	// 	- **ap-southeast-1a**: Singapore Zone A.
	//
	// 	- **ap-southeast-1b**: Singapore Zone B.
	//
	// 	- **ap-southeast-1c**: Singapore Zone C.
	//
	// 	- **ap-southeast-5a**: Jakarta Zone A.
	//
	// 	- **ap-southeast-5b**: Jakarta Zone B.
	//
	// 	- **ap-southeast-5c**: Jakarta Zone C.
	//
	// 	- **eu-central-1a**: Frankfurt Zone A.
	//
	// 	- **eu-central-1b**: Frankfurt Zone B.
	//
	// 	- **eu-central-1c**: Frankfurt Zone C.
	//
	// >  	- This parameter is valid and required when the **EngineVersion*	- parameter is set to **4.4*	- or **5.0**.
	//
	// >  	- The value of this parameter cannot be the same as the value of the **ZoneId*	- or **HiddenZoneId*	- parameter.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The IP addresses in an IP address whitelist. Multiple IP addresses are separated by commas (,), and each IP address in the IP address whitelist must be unique. The following types of values are supported:
	//
	// 	- 0.0.0.0/0
	//
	// 	- IP addresses, such as 10.23.12.24.
	//
	// 	- Classless Inter-Domain Routing (CIDR) blocks, such as 10.23.12.0/24. In this case, /24 indicates that the prefix of each IP address is 24-bit long. You can replace 24 with a value within the range of 1 to 32.
	//
	// > 	- A maximum of 1,000 IP addresses or CIDR blocks can be configured for each instance.
	//
	// > 	- If you enter 0.0.0.0/0, all IP addresses can access the instance. This may introduce security risks to the instance. Proceed with caution.
	//
	// example:
	//
	// 192.168.xx.xx,192.168.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The ID of the source instance.
	//
	// > When you call this operation to clone an instance, this parameter is required. The **BackupId*	- or **RestoreTime*	- parameter is also required. When you call this operation to restore an instance from the recycle bin, this parameter is required. The **BackupId*	- or **RestoreTime*	- parameter is not required.
	//
	// example:
	//
	// dds-bp1ee12ad351****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The region ID of the instance.
	//
	// > -  This parameter is required when restore type is set to 2 or 3.
	//
	// example:
	//
	// 2
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The storage engine of the instance. Set the value to **WiredTiger**.
	//
	// > 	- If you call this operation to clone an instance or restore an instance from the recycle bin, set this parameter to the storage engine of the source instance.
	//
	// > 	- For more information about the limits on database versions and storage engines of an instance, see [MongoDB versions and storage engines](https://help.aliyun.com/document_detail/61906.html).
	//
	// example:
	//
	// WiredTiger
	StorageEngine *string `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_essd1*	- :ESSD PL1.
	//
	// 	- **cloud_essd2**: ESSD PL2.
	//
	// 	- **cloud_essd3**: ESSD PL3.
	//
	// 	- **local_ssd**: local SSD.
	//
	// example:
	//
	// cloud_essd1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The custom tags added to the instance.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the instance is connected.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateDBInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateDBInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateDBInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateDBInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateDBInstanceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CreateDBInstanceRequest) GetDatabaseNames() *string {
	return s.DatabaseNames
}

func (s *CreateDBInstanceRequest) GetEncrypted() *bool {
	return s.Encrypted
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

func (s *CreateDBInstanceRequest) GetGlobalSecurityGroupIds() *string {
	return s.GlobalSecurityGroupIds
}

func (s *CreateDBInstanceRequest) GetHiddenZoneId() *string {
	return s.HiddenZoneId
}

func (s *CreateDBInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateDBInstanceRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateDBInstanceRequest) GetReadonlyReplicas() *string {
	return s.ReadonlyReplicas
}

func (s *CreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceRequest) GetReplicationFactor() *string {
	return s.ReplicationFactor
}

func (s *CreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateDBInstanceRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CreateDBInstanceRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *CreateDBInstanceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDBInstanceRequest) GetSrcDBInstanceId() *string {
	return s.SrcDBInstanceId
}

func (s *CreateDBInstanceRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *CreateDBInstanceRequest) GetStorageEngine() *string {
	return s.StorageEngine
}

func (s *CreateDBInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDBInstanceRequest) GetTag() []*CreateDBInstanceRequestTag {
	return s.Tag
}

func (s *CreateDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequest) SetAccountPassword(v string) *CreateDBInstanceRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v string) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBackupId(v string) *CreateDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBusinessInfo(v string) *CreateDBInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateDBInstanceRequest) SetChargeType(v string) *CreateDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClusterId(v string) *CreateDBInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCouponNo(v string) *CreateDBInstanceRequest {
	s.CouponNo = &v
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

func (s *CreateDBInstanceRequest) SetDBInstanceStorage(v int32) *CreateDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDatabaseNames(v string) *CreateDBInstanceRequest {
	s.DatabaseNames = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncrypted(v bool) *CreateDBInstanceRequest {
	s.Encrypted = &v
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

func (s *CreateDBInstanceRequest) SetGlobalSecurityGroupIds(v string) *CreateDBInstanceRequest {
	s.GlobalSecurityGroupIds = &v
	return s
}

func (s *CreateDBInstanceRequest) SetHiddenZoneId(v string) *CreateDBInstanceRequest {
	s.HiddenZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetNetworkType(v string) *CreateDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerAccount(v string) *CreateDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v int32) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetProvisionedIops(v int64) *CreateDBInstanceRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateDBInstanceRequest) SetReadonlyReplicas(v string) *CreateDBInstanceRequest {
	s.ReadonlyReplicas = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetReplicationFactor(v string) *CreateDBInstanceRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerId(v int64) *CreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRestoreTime(v string) *CreateDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRestoreType(v string) *CreateDBInstanceRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecondaryZoneId(v string) *CreateDBInstanceRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSrcDBInstanceId(v string) *CreateDBInstanceRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSrcRegion(v string) *CreateDBInstanceRequest {
	s.SrcRegion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageEngine(v string) *CreateDBInstanceRequest {
	s.StorageEngine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageType(v string) *CreateDBInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTag(v []*CreateDBInstanceRequestTag) *CreateDBInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVpcId(v string) *CreateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDBInstanceRequestTag struct {
	// The key of the tag.
	//
	// > **N*	- specifies the serial number of the tag. For example, **Tag.1.Key*	- specifies the key of the first tag and **Tag.2.Key*	- specifies the key of the second tag.
	//
	// example:
	//
	// testdatabase
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// > **N*	- specifies the serial number of the tag. For example, **Tag.1.Value*	- specifies the value of the first tag and **Tag.2.Value*	- specifies the value of the second tag.
	//
	// example:
	//
	// apitest
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
