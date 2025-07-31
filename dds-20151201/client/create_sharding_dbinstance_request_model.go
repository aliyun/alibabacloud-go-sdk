// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShardingDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountPassword(v string) *CreateShardingDBInstanceRequest
	GetAccountPassword() *string
	SetAutoRenew(v string) *CreateShardingDBInstanceRequest
	GetAutoRenew() *string
	SetBackupId(v string) *CreateShardingDBInstanceRequest
	GetBackupId() *string
	SetChargeType(v string) *CreateShardingDBInstanceRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateShardingDBInstanceRequest
	GetClientToken() *string
	SetConfigServer(v []*CreateShardingDBInstanceRequestConfigServer) *CreateShardingDBInstanceRequest
	GetConfigServer() []*CreateShardingDBInstanceRequestConfigServer
	SetDBInstanceDescription(v string) *CreateShardingDBInstanceRequest
	GetDBInstanceDescription() *string
	SetDestRegion(v string) *CreateShardingDBInstanceRequest
	GetDestRegion() *string
	SetEncrypted(v bool) *CreateShardingDBInstanceRequest
	GetEncrypted() *bool
	SetEncryptionKey(v string) *CreateShardingDBInstanceRequest
	GetEncryptionKey() *string
	SetEngine(v string) *CreateShardingDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateShardingDBInstanceRequest
	GetEngineVersion() *string
	SetGlobalSecurityGroupIds(v string) *CreateShardingDBInstanceRequest
	GetGlobalSecurityGroupIds() *string
	SetHiddenZoneId(v string) *CreateShardingDBInstanceRequest
	GetHiddenZoneId() *string
	SetMongos(v []*CreateShardingDBInstanceRequestMongos) *CreateShardingDBInstanceRequest
	GetMongos() []*CreateShardingDBInstanceRequestMongos
	SetNetworkType(v string) *CreateShardingDBInstanceRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateShardingDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateShardingDBInstanceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateShardingDBInstanceRequest
	GetPeriod() *int32
	SetProtocolType(v string) *CreateShardingDBInstanceRequest
	GetProtocolType() *string
	SetProvisionedIops(v int64) *CreateShardingDBInstanceRequest
	GetProvisionedIops() *int64
	SetRegionId(v string) *CreateShardingDBInstanceRequest
	GetRegionId() *string
	SetReplicaSet(v []*CreateShardingDBInstanceRequestReplicaSet) *CreateShardingDBInstanceRequest
	GetReplicaSet() []*CreateShardingDBInstanceRequestReplicaSet
	SetResourceGroupId(v string) *CreateShardingDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateShardingDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateShardingDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CreateShardingDBInstanceRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *CreateShardingDBInstanceRequest
	GetRestoreType() *string
	SetSecondaryZoneId(v string) *CreateShardingDBInstanceRequest
	GetSecondaryZoneId() *string
	SetSecurityIPList(v string) *CreateShardingDBInstanceRequest
	GetSecurityIPList() *string
	SetSrcDBInstanceId(v string) *CreateShardingDBInstanceRequest
	GetSrcDBInstanceId() *string
	SetSrcRegion(v string) *CreateShardingDBInstanceRequest
	GetSrcRegion() *string
	SetStorageEngine(v string) *CreateShardingDBInstanceRequest
	GetStorageEngine() *string
	SetStorageType(v string) *CreateShardingDBInstanceRequest
	GetStorageType() *string
	SetTag(v []*CreateShardingDBInstanceRequestTag) *CreateShardingDBInstanceRequest
	GetTag() []*CreateShardingDBInstanceRequestTag
	SetVSwitchId(v string) *CreateShardingDBInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateShardingDBInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateShardingDBInstanceRequest
	GetZoneId() *string
}

type CreateShardingDBInstanceRequest struct {
	// The password of the root account. The password must meet the following requirements:
	//
	// 	- The password contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- The following special characters are supported: ! @ # $ % ^ & \\	- ( ) _ + - =.
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// >  For more information about how to resolve failed database connections due to special characters, see [What do I do if my instance is not connected due to special characters in the password in the connection string of the instance?](https://help.aliyun.com/document_detail/471568.html)
	//
	// example:
	//
	// 123456Aa
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > This parameter is available and optional if you set the value of **ChargeType*	- to **PrePaid**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the backup set.
	//
	// > When you call this operation to clone an instance based on the backup set, this parameter is required. The **SrcDBInstanceId*	- parameter is also required.
	//
	// example:
	//
	// cb-xxx
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PostPaid*	- (default): pay-as-you-go
	//
	// 	- **PrePaid**: subscription
	//
	// >  If you set this parameter to **PrePaid**, you must also configure the **Period*	- parameter.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ConfigServer nodes of the instance.
	//
	// This parameter is required.
	ConfigServer []*CreateShardingDBInstanceRequestConfigServer `json:"ConfigServer,omitempty" xml:"ConfigServer,omitempty" type:"Repeated"`
	// The name of the instance. The name of the instance must meet the following requirements:
	//
	// 	- The name must start with a letter.
	//
	// 	- It can contain digits, letters, underscores (_), and hyphens (-).
	//
	// 	- It must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The region of the backup set used for the cross-region backup and restoration.
	//
	// >  This parameter is required when you set the RestoreType parameter to 3.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. Valid values:
	//
	// 	- **7.0**
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
	// >
	//
	// 	- For more information about the limits on database versions and storage engines, see [MongoDB versions and storage engines](https://help.aliyun.com/document_detail/61906.html).
	//
	// 	- If you call this operation to clone an instance, set the value of this parameter to the database engine version of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The global IP address whitelist template of the instance. Separate multiple templates with commas (,). The template name must be globally unique.
	//
	// example:
	//
	// g-qxieqf40xjst1ngp****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The ID of secondary zone 2 for multi-zone deployment. Valid values:
	//
	// 	- **cn-hangzhou-g**: Hangzhou Zone G
	//
	// 	- **cn-hangzhou-h**: Hangzhou Zone H
	//
	// 	- **cn-hangzhou-i**: Hangzhou Zone I
	//
	// 	- **cn-hongkong-b**: Hong Kong Zone B
	//
	// 	- **cn-hongkong-c**: Hong Kong Zone C
	//
	// 	- **cn-hongkong-d**: Hong Kong Zone D
	//
	// 	- **cn-wulanchabu-a**: Ulanqab Zone A
	//
	// 	- **cn-wulanchabu-b**: Ulanqab Zone B
	//
	// 	- **cn-wulanchabu-c**: Ulanqab Zone C
	//
	// 	- **ap-southeast-1a**: Singapore Zone A
	//
	// 	- **ap-southeast-1b**: Singapore Zone B
	//
	// 	- **ap-southeast-1c**: Singapore Zone C
	//
	// 	- **ap-southeast-5a**: Jakarta Zone A
	//
	// 	- **ap-southeast-5b**: Jakarta Zone B
	//
	// 	- **ap-southeast-5c**: Jakarta Zone C
	//
	// 	- **eu-central-1a**: Frankfurt Zone A
	//
	// 	- **eu-central-1b**: Frankfurt Zone B
	//
	// 	- **eu-central-1c**: Frankfurt Zone C
	//
	// > 	- This parameter is available and required if you set the value of **EngineVersion*	- to **4.4*	- or **5.0**.
	//
	// > 	- The value of this parameter cannot be the same as the value of **ZoneId*	- or **SecondaryZoneId**.
	//
	// > 	- For more information about the multi-zone deployment policy of a sharded cluster instance, see [Create a multi-zone sharded cluster instance](https://help.aliyun.com/document_detail/117865.html).
	//
	// example:
	//
	// cn-hangzhou-i
	HiddenZoneId *string `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	// The mongos nodes of the instance.
	//
	// This parameter is required.
	Mongos []*CreateShardingDBInstanceRequestMongos `json:"Mongos,omitempty" xml:"Mongos,omitempty" type:"Repeated"`
	// The network type of the instance. Set the value to VPC.
	//
	// ****
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
	// The access protocol type of the instance. Valid values:
	//
	// 	- **mongodb**
	//
	// 	- **dynamodb**
	//
	// example:
	//
	// mongodb
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The provisioned IOPS of the instance:
	//
	// example:
	//
	// 1960
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information of the shard component.
	//
	// This parameter is required.
	ReplicaSet []*CreateShardingDBInstanceRequestReplicaSet `json:"ReplicaSet,omitempty" xml:"ReplicaSet,omitempty" type:"Repeated"`
	// The resource group ID. For more information, see [View the basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to restore the instance, which must be within seven days. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in Coordinated Universal Time (UTC).
	//
	// > This parameter is required only if you call this operation to clone an instance. If you specify this parameter, you must also specify **SrcDBInstanceId**.
	//
	// example:
	//
	// 2022-03-08T02:30:25Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restoration type of the instance. Valid values:
	//
	// 	- 1: restores the instance data to the specified point in time.
	//
	// 	- 2: restores the data of the released instance to the specified backup set.
	//
	// 	- 3: restores the instance data to the specified cross-region backup set.
	//
	// example:
	//
	// 1
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The ID of secondary zone 1 for multi-zone deployment. Valid values:
	//
	// 	- **cn-hangzhou-g**: Hangzhou Zone G
	//
	// 	- **cn-hangzhou-h**: Hangzhou Zone H
	//
	// 	- **cn-hangzhou-i**: Hangzhou Zone I
	//
	// 	- **cn-hongkong-b**: Hong Kong Zone B
	//
	// 	- **cn-hongkong-c**: Hong Kong Zone C
	//
	// 	- **cn-hongkong-d**: Hong Kong Zone D
	//
	// 	- **cn-wulanchabu-a**: Ulanqab Zone A
	//
	// 	- **cn-wulanchabu-b**: Ulanqab Zone B
	//
	// 	- **cn-wulanchabu-c**: Ulanqab Zone C
	//
	// 	- **ap-southeast-1a**: Singapore Zone A
	//
	// 	- **ap-southeast-1b**: Singapore Zone B
	//
	// 	- **ap-southeast-1c**: Singapore Zone C
	//
	// 	- **ap-southeast-5a**: Jakarta Zone A
	//
	// 	- **ap-southeast-5b**: Jakarta Zone B
	//
	// 	- **ap-southeast-5c**: Jakarta Zone C
	//
	// 	- **eu-central-1a**: Frankfurt Zone A
	//
	// 	- **eu-central-1b**: Frankfurt Zone B
	//
	// 	- **eu-central-1c**: Frankfurt Zone C
	//
	// > 	- This parameter is available and required if you set the value of **EngineVersion*	- to **4.4*	- or **5.0**.
	//
	// > 	- The value of this parameter cannot be the same as the value of **ZoneId*	- or **HiddenZoneId**.
	//
	// > 	- For more information about the multi-zone deployment policy of a sharded cluster instance, see [Create a multi-zone sharded cluster instance](https://help.aliyun.com/document_detail/117865.html).
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The IP addresses in an IP address whitelist of the instance. Multiple IP addresses are separated by commas (,), and each IP address in the IP address whitelist must be unique. The following types of values are supported:
	//
	// 	- 0.0.0.0/0
	//
	// 	- IP addresses, such as 10.23.12.24.
	//
	// 	- CIDR blocks, such as 10.23.12.0/24. In this case, 24 indicates that the prefix of each IP address is 24-bit long. You can replace 24 with a value within the range of 1 to 32.
	//
	// > 	- A maximum of 1,000 IP addresses and CIDR blocks can be configured for each instance.
	//
	// > 	- If you enter 0.0.0.0/0, all IP addresses can access the instance. This may introduce security risks to the instance. Proceed with caution.
	//
	// example:
	//
	// 192.168.xx.xx,192.168.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The source instance ID.
	//
	// > This parameter is required only if you call this operation to clone an instance. If you specify this parameter, you must also specify **RestoreTime**.
	//
	// example:
	//
	// dds-bp11483712c1****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The region ID of the instance.
	//
	// > This parameter is required when restore type is set to 2 or 3.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The storage engine of the instance. Set the value to **WiredTiger**.
	//
	// >
	//
	// 	- If you call this operation to clone an instance, set the value of this parameter to the storage engine of the source instance.
	//
	// 	- For more information about the limits on database versions and storage engines, see [MongoDB versions and storage engines](https://help.aliyun.com/document_detail/61906.html).
	//
	// example:
	//
	// WiredTiger
	StorageEngine *string `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_essd1**: ESSD PL1
	//
	// 	- **cloud_essd2**: ESSD PL2
	//
	// 	- **cloud_essd3**: ESSD PL3
	//
	// 	- **local_ssd**: local SSD
	//
	// > 	- Instances of MongoDB 4.4 and later support only cloud disks. **cloud_essd1*	- is selected if you leave this parameter empty.
	//
	// > 	- Instances of MongoDB 4.2 and earlier support only local disks. **local_ssd*	- is selected if you leave this parameter empty.
	//
	// example:
	//
	// cloud_essd1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The custom tags that you want to add to the instance.
	Tag []*CreateShardingDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID of the instance.
	//
	// example:
	//
	// vsw-bp1vj604nj5a9zz74****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateShardingDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateShardingDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateShardingDBInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateShardingDBInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateShardingDBInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateShardingDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateShardingDBInstanceRequest) GetConfigServer() []*CreateShardingDBInstanceRequestConfigServer {
	return s.ConfigServer
}

func (s *CreateShardingDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateShardingDBInstanceRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *CreateShardingDBInstanceRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateShardingDBInstanceRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateShardingDBInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateShardingDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateShardingDBInstanceRequest) GetGlobalSecurityGroupIds() *string {
	return s.GlobalSecurityGroupIds
}

func (s *CreateShardingDBInstanceRequest) GetHiddenZoneId() *string {
	return s.HiddenZoneId
}

func (s *CreateShardingDBInstanceRequest) GetMongos() []*CreateShardingDBInstanceRequestMongos {
	return s.Mongos
}

func (s *CreateShardingDBInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateShardingDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateShardingDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateShardingDBInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateShardingDBInstanceRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *CreateShardingDBInstanceRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateShardingDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateShardingDBInstanceRequest) GetReplicaSet() []*CreateShardingDBInstanceRequestReplicaSet {
	return s.ReplicaSet
}

func (s *CreateShardingDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateShardingDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateShardingDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateShardingDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateShardingDBInstanceRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CreateShardingDBInstanceRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *CreateShardingDBInstanceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateShardingDBInstanceRequest) GetSrcDBInstanceId() *string {
	return s.SrcDBInstanceId
}

func (s *CreateShardingDBInstanceRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *CreateShardingDBInstanceRequest) GetStorageEngine() *string {
	return s.StorageEngine
}

func (s *CreateShardingDBInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateShardingDBInstanceRequest) GetTag() []*CreateShardingDBInstanceRequestTag {
	return s.Tag
}

func (s *CreateShardingDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateShardingDBInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateShardingDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateShardingDBInstanceRequest) SetAccountPassword(v string) *CreateShardingDBInstanceRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetAutoRenew(v string) *CreateShardingDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetBackupId(v string) *CreateShardingDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetChargeType(v string) *CreateShardingDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetClientToken(v string) *CreateShardingDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetConfigServer(v []*CreateShardingDBInstanceRequestConfigServer) *CreateShardingDBInstanceRequest {
	s.ConfigServer = v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetDBInstanceDescription(v string) *CreateShardingDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetDestRegion(v string) *CreateShardingDBInstanceRequest {
	s.DestRegion = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetEncrypted(v bool) *CreateShardingDBInstanceRequest {
	s.Encrypted = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetEncryptionKey(v string) *CreateShardingDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetEngine(v string) *CreateShardingDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetEngineVersion(v string) *CreateShardingDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetGlobalSecurityGroupIds(v string) *CreateShardingDBInstanceRequest {
	s.GlobalSecurityGroupIds = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetHiddenZoneId(v string) *CreateShardingDBInstanceRequest {
	s.HiddenZoneId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetMongos(v []*CreateShardingDBInstanceRequestMongos) *CreateShardingDBInstanceRequest {
	s.Mongos = v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetNetworkType(v string) *CreateShardingDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetOwnerAccount(v string) *CreateShardingDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetOwnerId(v int64) *CreateShardingDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetPeriod(v int32) *CreateShardingDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetProtocolType(v string) *CreateShardingDBInstanceRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetProvisionedIops(v int64) *CreateShardingDBInstanceRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetRegionId(v string) *CreateShardingDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetReplicaSet(v []*CreateShardingDBInstanceRequestReplicaSet) *CreateShardingDBInstanceRequest {
	s.ReplicaSet = v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetResourceGroupId(v string) *CreateShardingDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateShardingDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetResourceOwnerId(v int64) *CreateShardingDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetRestoreTime(v string) *CreateShardingDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetRestoreType(v string) *CreateShardingDBInstanceRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetSecondaryZoneId(v string) *CreateShardingDBInstanceRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetSecurityIPList(v string) *CreateShardingDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetSrcDBInstanceId(v string) *CreateShardingDBInstanceRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetSrcRegion(v string) *CreateShardingDBInstanceRequest {
	s.SrcRegion = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetStorageEngine(v string) *CreateShardingDBInstanceRequest {
	s.StorageEngine = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetStorageType(v string) *CreateShardingDBInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetTag(v []*CreateShardingDBInstanceRequestTag) *CreateShardingDBInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetVSwitchId(v string) *CreateShardingDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetVpcId(v string) *CreateShardingDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) SetZoneId(v string) *CreateShardingDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateShardingDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateShardingDBInstanceRequestConfigServer struct {
	// The instance type of the ConfigServer node. Valid values:
	//
	// 	- **mdb.shard.2x.xlarge.d**: 4 cores, 8 GB (dedicated). Only instances that run MongoDB 4.4 and later support this instance type.
	//
	// 	- **dds.cs.mid*	- :1 core, 2 GB (general-purpose). Only instances that run MongoDB 4.2 and earlier support this instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// mdb.shard.2x.xlarge.d
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// The storage space of the ConfigServer node. Unit: GB.
	//
	// > The values that can be specified for this parameter vary based on the instance types. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s CreateShardingDBInstanceRequestConfigServer) String() string {
	return dara.Prettify(s)
}

func (s CreateShardingDBInstanceRequestConfigServer) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequestConfigServer) GetClass() *string {
	return s.Class
}

func (s *CreateShardingDBInstanceRequestConfigServer) GetStorage() *int32 {
	return s.Storage
}

func (s *CreateShardingDBInstanceRequestConfigServer) SetClass(v string) *CreateShardingDBInstanceRequestConfigServer {
	s.Class = &v
	return s
}

func (s *CreateShardingDBInstanceRequestConfigServer) SetStorage(v int32) *CreateShardingDBInstanceRequestConfigServer {
	s.Storage = &v
	return s
}

func (s *CreateShardingDBInstanceRequestConfigServer) Validate() error {
	return dara.Validate(s)
}

type CreateShardingDBInstanceRequestMongos struct {
	// The instance type of the mongos node. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// > 	- **N*	- specifies the serial number of the mongos node for which the instance type is specified. For example, **Mongos.2.Class*	- specifies the instance type of the second mongos node.
	//
	// > 	- Valid values for **N**: **2*	- to **32**.
	//
	// This parameter is required.
	//
	// example:
	//
	// mdb.shard.2x.xlarge.d
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
}

func (s CreateShardingDBInstanceRequestMongos) String() string {
	return dara.Prettify(s)
}

func (s CreateShardingDBInstanceRequestMongos) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequestMongos) GetClass() *string {
	return s.Class
}

func (s *CreateShardingDBInstanceRequestMongos) SetClass(v string) *CreateShardingDBInstanceRequestMongos {
	s.Class = &v
	return s
}

func (s *CreateShardingDBInstanceRequestMongos) Validate() error {
	return dara.Validate(s)
}

type CreateShardingDBInstanceRequestReplicaSet struct {
	// The instance type of the shard component. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// >
	//
	// 	- **N*	- specifies the serial number of the shard component for which the instance type is specified. For example, **ReplicaSet.2.Class*	- specifies the instance type of the second shard component.
	//
	// 	- Valid values of **N**: **2*	- to **32**.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds.shard.standard
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// The number of read-only nodes in the shard component.
	//
	// Valid values: **0**, **1, 2, 3, 4, and 5**. Default value: **0**.
	//
	// >  **N*	- specifies the serial number of the shard component for which you want to set the number of read-only nodes. **ReplicaSet.2.ReadonlyReplicas*	- specifies the number of read-only nodes in the second shard component.
	//
	// example:
	//
	// 0
	ReadonlyReplicas *int32 `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The storage capacity of the shard component. Unit: GB.
	//
	// >
	//
	// 	- The values that can be specified for this parameter vary based on the instance types. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// 	- **N*	- specifies the serial number of the shard component for which the storage capacity is specified. For example, **ReplicaSet.2.Storage*	- specifies the storage capacity of the second shard component.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s CreateShardingDBInstanceRequestReplicaSet) String() string {
	return dara.Prettify(s)
}

func (s CreateShardingDBInstanceRequestReplicaSet) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequestReplicaSet) GetClass() *string {
	return s.Class
}

func (s *CreateShardingDBInstanceRequestReplicaSet) GetReadonlyReplicas() *int32 {
	return s.ReadonlyReplicas
}

func (s *CreateShardingDBInstanceRequestReplicaSet) GetStorage() *int32 {
	return s.Storage
}

func (s *CreateShardingDBInstanceRequestReplicaSet) SetClass(v string) *CreateShardingDBInstanceRequestReplicaSet {
	s.Class = &v
	return s
}

func (s *CreateShardingDBInstanceRequestReplicaSet) SetReadonlyReplicas(v int32) *CreateShardingDBInstanceRequestReplicaSet {
	s.ReadonlyReplicas = &v
	return s
}

func (s *CreateShardingDBInstanceRequestReplicaSet) SetStorage(v int32) *CreateShardingDBInstanceRequestReplicaSet {
	s.Storage = &v
	return s
}

func (s *CreateShardingDBInstanceRequestReplicaSet) Validate() error {
	return dara.Validate(s)
}

type CreateShardingDBInstanceRequestTag struct {
	// The tag key.
	//
	// >  **N*	- specifies the serial number of the tag. For example, **Tag.1.Key*	- specifies the key of the first tag and **Tag.2.Key*	- specifies the key of the second tag.
	//
	// example:
	//
	// testdatabase
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// >  **N*	- specifies the serial number of the tag. For example, **Tag.1.Value*	- specifies the value of the first tag and Tag.2.Value specifies the value of the second tag.
	//
	// example:
	//
	// apitest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateShardingDBInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateShardingDBInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateShardingDBInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateShardingDBInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateShardingDBInstanceRequestTag) SetKey(v string) *CreateShardingDBInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateShardingDBInstanceRequestTag) SetValue(v string) *CreateShardingDBInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateShardingDBInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
