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
	// - It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Special characters include !@#$%^&\\*()_+-=
	//
	// - It must be 8 to 32 characters in length.
	//
	// > For information about how to resolve connection failures caused by special characters in passwords, see [How do I fix connection failures caused by special characters in a password?](https://help.aliyun.com/document_detail/471568.html).
	//
	// example:
	//
	// 123456Aa
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// - **true**: Auto-renewal is enabled.
	//
	// - **false**: Auto-renewal is disabled. You must manually renew the instance. This is the default value.
	//
	// > This parameter is optional and takes effect only when you set the **ChargeType*	- parameter to **PrePaid**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The cluster backup ID.
	//
	// > - This parameter is required only when RestoreType is set to 2 or 3.
	//
	// example:
	//
	// cb-xxx
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PostPaid**: pay-as-you-go. This is the default value.
	//
	// - **PrePaid**: subscription.
	//
	// > If you set this parameter to **PrePaid**, you must also specify the **Period*	- parameter.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information of Configserver nodes.
	//
	// This parameter is required.
	ConfigServer []*CreateShardingDBInstanceRequestConfigServer `json:"ConfigServer,omitempty" xml:"ConfigServer,omitempty" type:"Repeated"`
	// The name of the instance. The name must meet the following requirements:
	//
	// - It must start with a Chinese character or a letter.
	//
	// - It can contain digits, Chinese characters, letters, underscores (_), periods (.), and hyphens (-).
	//
	// - It must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The region where the geo-redundant backup is stored.
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
	// The custom key ID.
	//
	// example:
	//
	// 2axxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The database engine. Set the value to **MongoDB**.
	//
	// This parameter is required.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database version. Valid values:
	//
	// - **8.0**
	//
	// - **7.0**
	//
	// - **6.0**
	//
	// - **5.0**
	//
	// - **4.4**
	//
	// - **4.2**
	//
	// - **4.0**
	//
	// > 	- For more information about the constraints on storage engines and database versions, see [Versions and storage engines](https://help.aliyun.com/document_detail/61906.html).
	//
	// >
	//
	// > 	- When you clone an instance by calling this operation, the value of this parameter must be the same as that of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The global IP address whitelist templates of the instance. Separate multiple templates with commas (,). Each template must be unique.
	//
	// example:
	//
	// g-qxieqf40xjst1ngpr3jz
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The secondary zone 2 for multi-zone deployment. Valid values:
	//
	// - **cn-hangzhou-g**: Hangzhou Zone G.
	//
	// - **cn-hangzhou-h**: Hangzhou Zone H.
	//
	// - **cn-hangzhou-i**: Hangzhou Zone I.
	//
	// - **cn-hongkong-b**: Hong Kong (China) Zone B.
	//
	// - **cn-hongkong-c**: Hong Kong (China) Zone C.
	//
	// - **cn-hongkong-d**: Hong Kong (China) Zone D.
	//
	// - **cn-wulanchabu-a**: Ulanqab Zone A.
	//
	// - **cn-wulanchabu-b**: Ulanqab Zone B.
	//
	// - **cn-wulanchabu-c**: Ulanqab Zone C.
	//
	// - **ap-southeast-1a**: Singapore Zone A.
	//
	// - **ap-southeast-1b**: Singapore Zone B.
	//
	// - **ap-southeast-1c**: Singapore Zone C.
	//
	// - **ap-southeast-5a**: Jakarta Zone A.
	//
	// - **ap-southeast-5b**: Jakarta Zone B.
	//
	// - **ap-southeast-5c**: Jakarta Zone C.
	//
	// - **eu-central-1a**: Frankfurt Zone A.
	//
	// - **eu-central-1b**: Frankfurt Zone B.
	//
	// - **eu-central-1c**: Frankfurt Zone C.
	//
	// > 	- This parameter is available for disk-based instances.
	//
	// >
	//
	// > 	- The value of this parameter cannot be the same as the value of **ZoneId*	- or **SecondaryZoneId**.
	//
	// >
	//
	// > 	- For more information about the multi-zone deployment policy for sharded cluster instances, see [Create a multi-zone sharded cluster instance](https://help.aliyun.com/document_detail/117865.html).
	//
	// example:
	//
	// cn-hangzhou-i
	HiddenZoneId *string `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	// The information of Mongos nodes.
	//
	// This parameter is required.
	Mongos []*CreateShardingDBInstanceRequestMongos `json:"Mongos,omitempty" xml:"Mongos,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// **VPC**: virtual private cloud.
	//
	// example:
	//
	// VPC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the instance. Unit: month.
	//
	// Valid values: **1*	- to **9*	- (integer), **12**, **24**, **36**, and **60**.
	//
	// > This parameter is required and takes effect only when you set the **ChargeType*	- parameter to **PrePaid**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The protocol type of the instance. Valid values:
	//
	// - **mongodb**: MongoDB protocol.
	//
	// - **dynamodb**: DynamoDB protocol.
	//
	// example:
	//
	// mongodb
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The provisioned IOPS.
	//
	// example:
	//
	// 1960
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information of shard nodes.
	//
	// This parameter is required.
	ReplicaSet           []*CreateShardingDBInstanceRequestReplicaSet `json:"ReplicaSet,omitempty" xml:"ReplicaSet,omitempty" type:"Repeated"`
	ResourceGroupId      *string                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore data. You can specify any point in time within the last seven days. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is in Coordinated Universal Time (UTC).
	//
	// > This parameter is required only when you clone an instance by calling this operation. You must also specify the **SrcDBInstanceId*	- parameter.
	//
	// example:
	//
	// 2022-03-08T02:30:25Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The backup-based instance restoration method.
	//
	// - 1: Restore the instance to a specific point in time.
	//
	// - 2: Restore a released instance from a specific backup set.
	//
	// - 3: Restore the instance from a specific geo-redundant backup set.
	//
	// example:
	//
	// 1
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The secondary zone 1 for multi-zone deployment. Valid values:
	//
	// - **cn-hangzhou-g**: Hangzhou Zone G.
	//
	// - **cn-hangzhou-h**: Hangzhou Zone H.
	//
	// - **cn-hangzhou-i**: Hangzhou Zone I.
	//
	// - **cn-hongkong-b**: Hong Kong (China) Zone B.
	//
	// - **cn-hongkong-c**: Hong Kong (China) Zone C.
	//
	// - **cn-hongkong-d**: Hong Kong (China) Zone D.
	//
	// - **cn-wulanchabu-a**: Ulanqab Zone A.
	//
	// - **cn-wulanchabu-b**: Ulanqab Zone B.
	//
	// - **cn-wulanchabu-c**: Ulanqab Zone C.
	//
	// - **ap-southeast-1a**: Singapore Zone A.
	//
	// - **ap-southeast-1b**: Singapore Zone B.
	//
	// - **ap-southeast-1c**: Singapore Zone C.
	//
	// - **ap-southeast-5a**: Jakarta Zone A.
	//
	// - **ap-southeast-5b**: Jakarta Zone B.
	//
	// - **ap-southeast-5c**: Jakarta Zone C.
	//
	// - **eu-central-1a**: Frankfurt Zone A.
	//
	// - **eu-central-1b**: Frankfurt Zone B.
	//
	// - **eu-central-1c**: Frankfurt Zone C.
	//
	// > 	- This parameter is available for disk-based instances.
	//
	// >
	//
	// > 	- The value of this parameter cannot be the same as the value of **ZoneId*	- or **HiddenZoneId**.
	//
	// >
	//
	// > 	- For more information about the multi-zone deployment policy for sharded cluster instances, see [Create a multi-zone sharded cluster instance](https://help.aliyun.com/document_detail/117865.html).
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The IP address whitelist of the instance. Separate multiple IP addresses with commas (,). Each IP address in the whitelist must be unique. The following formats are supported:
	//
	// - 0.0.0.0/0
	//
	// - IP addresses, such as 10.23.12.24.
	//
	// - CIDR blocks, such as 10.23.12.0/24. The /24 part indicates the prefix length of the CIDR block. The prefix length ranges from 1 to 32.
	//
	// > 	- You can add a maximum of 1,000 IP addresses or CIDR blocks to all IP address whitelists.
	//
	// >
	//
	// > 	- The 0.0.0.0/0 entry allows access from all IP addresses. This is a high-risk setting. Configure it with caution.
	//
	// example:
	//
	// 192.168.xx.xx,192.168.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The source instance ID.
	//
	// > This parameter is required only when you clone an instance by calling this operation. You must also specify the **RestoreTime*	- parameter.
	//
	// example:
	//
	// dds-bp11483712c1****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The region of the source instance.
	//
	// > - This parameter is required when you recreate a released instance from a backup.
	//
	// >
	//
	// > - This parameter is required when you clone an instance from a geo-redundant backup.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The storage engine of the instance. Set the value to **WiredTiger**.
	//
	// > - When you clone an instance by calling this operation, the value of this parameter must be the same as that of the source instance.
	//
	// >
	//
	// > - For more information about the constraints on storage engines and database versions, see [Versions and storage engines](https://help.aliyun.com/document_detail/61906.html).
	//
	// example:
	//
	// WiredTiger
	StorageEngine *string `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	// The storage type. Valid values:
	//
	// - **cloud_essd1**: enhanced SSD (ESSD) PL1.
	//
	// - **cloud_essd2**: ESSD PL2.
	//
	// - **cloud_essd3**: ESSD PL3.
	//
	// - **local_ssd**: local SSD.
	//
	// > 	- Instances that run MongoDB 4.4 or later support only disks. If you do not specify this parameter, **cloud_essd1*	- is used.
	//
	// >
	//
	// > 	- Instances that run MongoDB 4.2 or earlier support only local disks. If you do not specify this parameter, **local_ssd*	- is used.
	//
	// example:
	//
	// cloud_essd1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The custom tags.
	Tag []*CreateShardingDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The virtual switch ID.
	//
	// example:
	//
	// vsw-bp1vj604nj5a9zz74****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the zone ID.
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
	if s.ConfigServer != nil {
		for _, item := range s.ConfigServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Mongos != nil {
		for _, item := range s.Mongos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReplicaSet != nil {
		for _, item := range s.ReplicaSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreateShardingDBInstanceRequestConfigServer struct {
	// The instance type of the Configserver node. Valid values:
	//
	// - **mdb.shard.2x.xlarge.d**: 4-core 8 GB (dedicated). This instance type is available only for instances that run MongoDB 4.4 or later.
	//
	// - **dds.cs.mid**: 1-core 2 GB (general-purpose). This instance type is available only for instances that run MongoDB 4.2 or earlier.
	//
	// This parameter is required.
	//
	// example:
	//
	// mdb.shard.2x.xlarge.d
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// The storage space of the Configserver node. Unit: GB.
	//
	// > The value of this parameter is constrained by the instance type. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
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
	// The instance type of the Mongos node. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// > - **N*	- in the parameter name specifies the serial number of the Mongos node. For example, **Mongos.2.Class*	- specifies the instance type of the second Mongos node.
	//
	// >
	//
	// > - The value of **N*	- ranges from **2*	- to **32**.
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
	// The instance type of the shard node. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// > - **N*	- in the parameter name specifies the serial number of the shard node. For example, **ReplicaSet.2.Class*	- specifies the instance type of the second shard node.
	//
	// >
	//
	// > - The value of **N*	- ranges from **2*	- to **32**.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds.shard.standard
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// The number of read-only nodes in the shard node.
	//
	// Valid values: **0*	- to **5**. The default value is **0**.
	//
	// > **N*	- in the parameter name specifies the serial number of the shard node. For example, **ReplicaSet.2.ReadonlyReplicas*	- specifies the number of read-only nodes in the second shard node.
	//
	// example:
	//
	// 0
	ReadonlyReplicas *int32 `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The storage space of the shard node. Unit: GB.
	//
	// > - The value of this parameter is constrained by the instance type. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
	//
	// >
	//
	// > - **N*	- in the parameter name specifies the serial number of the shard node. For example, **ReplicaSet.2.Storage*	- specifies the storage space of the second shard node.
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
	// The key of the tag.
	//
	// > - **N*	- specifies the serial number of the tag. For example, **Tag.1.Key*	- specifies the key of the first tag, and **Tag.2.Key*	- specifies the key of the second tag.
	//
	// example:
	//
	// testdatabase
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// > **N*	- specifies the serial number of the tag. For example, **Tag.1.Value*	- specifies the value of the first tag, and **Tag.2.Value*	- specifies the value of the second tag.
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
