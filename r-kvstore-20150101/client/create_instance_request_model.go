// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppendonly(v string) *CreateInstanceRequest
	GetAppendonly() *string
	SetAutoRenew(v string) *CreateInstanceRequest
	GetAutoRenew() *string
	SetAutoRenewPeriod(v string) *CreateInstanceRequest
	GetAutoRenewPeriod() *string
	SetAutoUseCoupon(v string) *CreateInstanceRequest
	GetAutoUseCoupon() *string
	SetBackupId(v string) *CreateInstanceRequest
	GetBackupId() *string
	SetBusinessInfo(v string) *CreateInstanceRequest
	GetBusinessInfo() *string
	SetCapacity(v int64) *CreateInstanceRequest
	GetCapacity() *int64
	SetChargeType(v string) *CreateInstanceRequest
	GetChargeType() *string
	SetClusterBackupId(v string) *CreateInstanceRequest
	GetClusterBackupId() *string
	SetConnectionStringPrefix(v string) *CreateInstanceRequest
	GetConnectionStringPrefix() *string
	SetCouponNo(v string) *CreateInstanceRequest
	GetCouponNo() *string
	SetDedicatedHostGroupId(v string) *CreateInstanceRequest
	GetDedicatedHostGroupId() *string
	SetDryRun(v bool) *CreateInstanceRequest
	GetDryRun() *bool
	SetEngineVersion(v string) *CreateInstanceRequest
	GetEngineVersion() *string
	SetGlobalInstance(v bool) *CreateInstanceRequest
	GetGlobalInstance() *bool
	SetGlobalInstanceId(v string) *CreateInstanceRequest
	GetGlobalInstanceId() *string
	SetGlobalSecurityGroupIds(v string) *CreateInstanceRequest
	GetGlobalSecurityGroupIds() *string
	SetInstanceClass(v string) *CreateInstanceRequest
	GetInstanceClass() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetNetworkType(v string) *CreateInstanceRequest
	GetNetworkType() *string
	SetNodeType(v string) *CreateInstanceRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *CreateInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateInstanceRequest
	GetOwnerId() *int64
	SetParamGroupId(v string) *CreateInstanceRequest
	GetParamGroupId() *string
	SetPassword(v string) *CreateInstanceRequest
	GetPassword() *string
	SetPeriod(v string) *CreateInstanceRequest
	GetPeriod() *string
	SetPort(v string) *CreateInstanceRequest
	GetPort() *string
	SetPrivateIpAddress(v string) *CreateInstanceRequest
	GetPrivateIpAddress() *string
	SetReadOnlyCount(v int32) *CreateInstanceRequest
	GetReadOnlyCount() *int32
	SetRecoverConfigMode(v string) *CreateInstanceRequest
	GetRecoverConfigMode() *string
	SetRegionId(v string) *CreateInstanceRequest
	GetRegionId() *string
	SetReplicaCount(v int32) *CreateInstanceRequest
	GetReplicaCount() *int32
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CreateInstanceRequest
	GetRestoreTime() *string
	SetSecondaryZoneId(v string) *CreateInstanceRequest
	GetSecondaryZoneId() *string
	SetSecurityToken(v string) *CreateInstanceRequest
	GetSecurityToken() *string
	SetShardCount(v int32) *CreateInstanceRequest
	GetShardCount() *int32
	SetSlaveReadOnlyCount(v int32) *CreateInstanceRequest
	GetSlaveReadOnlyCount() *int32
	SetSlaveReplicaCount(v int32) *CreateInstanceRequest
	GetSlaveReplicaCount() *int32
	SetSrcDBInstanceId(v string) *CreateInstanceRequest
	GetSrcDBInstanceId() *string
	SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest
	GetTag() []*CreateInstanceRequestTag
	SetToken(v string) *CreateInstanceRequest
	GetToken() *string
	SetVSwitchId(v string) *CreateInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateInstanceRequest
	GetZoneId() *string
}

type CreateInstanceRequest struct {
	// Specifies whether to enable append-only file (AOF) persistence for the instance. Valid values:
	//
	// 	- **yes*	- (default): enables AOF persistence.
	//
	// 	- **no**: disables AOF persistence.
	//
	// **
	//
	// **Description*	- This parameter is applicable to classic instances, and is unavailable for cloud-native instances.
	//
	// example:
	//
	// yes
	Appendonly *string `json:"Appendonly,omitempty" xml:"Appendonly,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false*	- (default): disables auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The subscription duration that is supported by auto-renewal. Unit: month. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// >  This parameter is required if the **AutoRenew*	- parameter is set to **true**.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// 	- **true**: uses a coupon.
	//
	// 	- **false*	- (default): does not use a coupon.
	//
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// If your instance is a cloud-native cluster instance, we recommend that you use [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) to query the backup set ID of the cluster instance, such as cb-xx. Then, set the ClusterBackupId request parameter to the backup set ID to clone the cluster instance. This eliminates the need to specify the backup set ID of each shard.
	//
	// You can set the BackupId parameter to the backup set ID of the source instance. The system uses the data stored in the backup set to create an instance. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query backup set IDs. If the source instance is a cluster instance, set the BackupId parameter to the backup set IDs of all shards of the source instance, separated by commas (,). Example: "10\\*\\*,11\\*\\*,15\\*\\*".
	//
	// example:
	//
	// 111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the promotional event or business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The storage capacity of the instance. Unit: MB.
	//
	// > You must specify at least one of the **Capacity*	- and **InstanceClass*	- parameters when you call this operation.
	//
	// example:
	//
	// 16384
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid*	- (default): pay-as-you-go
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// This parameter is supported for specific new cluster instances. You can query the backup set ID by using the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation.
	//
	// 	- If this parameter is supported, you can specify the backup set ID. In this case, you do not need to specify the **BackupId*	- parameter.
	//
	// 	- If this parameter is not supported, set the BackupId parameter to the IDs of backup sets for all shards of the source instance, separated by commas (,). Example: "2158\\*\\*\\*\\*20,2158\\*\\*\\*\\*22".
	//
	// example:
	//
	// cb-hyxdof5x9kqbtust
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The operation that you want to perform. Set the value to **AllocateInstancePublicConnection**.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The coupon code. Default value: `default`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The ID of the dedicated cluster. This parameter is required if you create an instance in a dedicated cluster.
	//
	// example:
	//
	// dhg-uv4fnk6r7zff****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run and does not create the instance. The system prechecks the request parameters, request format, service limits, and available resources. If the request fails to pass the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and sends the request. If the request passes the dry run, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The engine version. Valid values for **classic instances**:
	//
	// 	- **2.8*	- (not recommended due to [scheduled EOFS](https://help.aliyun.com/document_detail/2674657.html))
	//
	// 	- **4.0*	- (not recommended)
	//
	// 	- **5.0**
	//
	// Valid values for **cloud-native instances**:
	//
	// 	- **5.0**
	//
	// 	- **6.0*	- (recommended)
	//
	// 	- **7.0**
	//
	// >  The default value is **5.0**.
	//
	// example:
	//
	// 4.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to use the new instance as the first child instance of a distributed instance. Valid values:
	//
	// 	- **true**: uses the new instance as the first child instance.
	//
	// 	- **false*	- (default): does not use the new instance as the first child instance.
	//
	// >
	//
	// 	- If you want to create a Tair DRAM-based instance that runs Redis 5.0, you must set this parameter to **true**.
	//
	// 	- This parameter is available only on the China site (aliyun.com).
	//
	// example:
	//
	// false
	GlobalInstance *bool `json:"GlobalInstance,omitempty" xml:"GlobalInstance,omitempty"`
	// The ID of the distributed instance. This parameter is available only on the China site (aliyun.com).
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The global IP whitelist template for the instance. Multiple IP whitelist templates should be separated by English commas (,) and cannot be duplicated.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The instance type. For example, redis.master.small.default indicates a Community Edition standard master-replica instance that has 1 GB of memory. For more information, see [Overview](https://help.aliyun.com/document_detail/26350.html).
	//
	// **
	//
	// **Description*	- You must specify at least one of the **Capacity*	- and **InstanceClass*	- parameters when you call the CreateInstance operation.
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The name of the instance. The name must be 2 to 80 characters in length and must start with a letter. It cannot contain spaces or specific special characters. These special characters include `@ / : = " < > { [ ] }`
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- **Redis*	- (default)
	//
	// 	- **Memcache**
	//
	// example:
	//
	// Redis
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The network type. Valid value:
	//
	// 	- **VPC*	- (default)
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The node type. Valid values:
	//
	// 	- **MASTER_SLAVE**: high availability (master-replica)
	//
	// 	- **STAND_ALONE**: standalone
	//
	// 	- **double**: master-replica
	//
	// 	- **single**: standalone
	//
	// >  To create a cloud-native instance, set this parameter to **MASTER_SLAVE*	- or **STAND_ALONE**. To create a classic instance, set this parameter to **double*	- or **single**.
	//
	// example:
	//
	// STAND_ALONE
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID, which must be globally unique.
	//
	// example:
	//
	// rpg-test**
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The password that is used to connect to the instance. The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and specific special characters. These special characters include `! @ # $ % ^ & 	- ( ) _ + - =`
	//
	// example:
	//
	// Pass!123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription duration. Valid values: **1**, 2, 3, 4, 5, 6, 7, 8, **9**, **12**, **24**,**36**, and **60**. Unit: months.
	//
	// > This parameter is available and required only if the **ChargeType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// 12
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The port number that is used to connect to the instance. Valid values: **1024*	- to **65535**. Default value: **6379**.
	//
	// example:
	//
	// 6379
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the instance.
	//
	// > The private IP address must be available within the CIDR block of the vSwitch to which to connect the instance.
	//
	// example:
	//
	// 172.16.0.***
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The number of read replicas in the primary zone. This parameter applies only to read/write splitting instances that use cloud disks. You can use this parameter to customize the number of read replicas. Valid values: 1 to 9.
	//
	// >  The sum of the values of this parameter and SlaveReadOnlyCount cannot be greater than 9.
	//
	// example:
	//
	// 5
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// When creating an instance using a specified backup set, whether to restore account, kernel parameter (whitelist), and whitelist (config) information from the original backup set. For example, if you need to restore account information, the value should be `{"account":true}`.
	//
	// By default, it is empty, indicating that no account, kernel parameter, or whitelist information will be restored from the original backup set.
	//
	// > This parameter applies only to cloud-native instances and requires that the original backup set has saved the account, kernel parameter, and whitelist information. You can use the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) API to check if the RecoverConfigMode parameter in the specified backup set contains the above information.
	//
	// example:
	//
	// {"whitelist":true,"config":true,"account":true}
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
	// The ID of the region where you want to create the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of slave replicas in the primary availability zone. This parameter is applicable only for creating cloud-native cluster edition multi-replica instances, allowing you to customize the number of slave replicas. The value range is 1 to 4.
	//
	// > > - The sum of this parameter and SlaveReplicaCount cannot exceed 4.
	//
	// >> - Only one of this parameter and ReadOnlyCount can be passed; there are no instances that simultaneously include both replicas and read-only nodes.
	//
	// >> - Primary-secondary instances do not support multiple replicas.
	//
	// example:
	//
	// 2
	ReplicaCount *int32 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-resourcegroupid1
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// If data flashback is enabled for the source instance, you can use this parameter to specify a point in time within the backup retention period of the source instance. The system uses the backup data of the source instance at the point in time to create an instance. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2019-06-19T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The secondary zone ID of the instance. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query the most recent zone list.
	//
	// > If you specify this parameter, the master node and replica node of the instance can be deployed in different zones and disaster recovery is implemented across zones. The instance can withstand failures in data centers.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// 系统自动生成的安全 Token，无需传入
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards. This parameter applies only to cloud-native cluster instances.
	//
	// example:
	//
	// 4
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The number of read replicas in the secondary zone. This parameter is used to create a read/write splitting instance that is deployed across multiple zones. The sum of the values of this parameter and ReadOnlyCount cannot be greater than 9.
	//
	// > When you create a multi-zone read/write splitting instance, you must specify both SlaveReadOnlyCount and SecondaryZoneId.
	//
	// example:
	//
	// 2
	SlaveReadOnlyCount *int32 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// Used for specifying the number of slave replicas in the secondary availability zone when creating a multi-AZ cloud-native cluster edition with multiple replicas. The sum of this parameter and ReplicaCount cannot exceed 4. <notice>When creating a multi-AZ cloud-native cluster edition with multiple replicas, both SlaveReplicaCount and SecondaryZoneId parameters must be specified.</notice>
	//
	// example:
	//
	// 2
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// If you want to create an instance based on the backup set of an existing instance, set this parameter to the ID of the source instance.
	//
	// >  After you specify the SrcDBInstanceId parameter, use the **BackupId**, **ClusterBackupId*	- (recommended for cloud-native cluster instances), or **RestoreTime*	- parameter to specify the backup set or the specific point in time that you want to use to create an instance. The SrcDBInstanceId parameter must be used in combination with one of the preceding three parameters.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The tags of the instance.
	Tag []*CreateInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the token is unique among different requests. The token is case-sensitive. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the vSwitch to which you want the instance to connect.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The primary zone ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetAppendonly() *string {
	return s.Appendonly
}

func (s *CreateInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetAutoRenewPeriod() *string {
	return s.AutoRenewPeriod
}

func (s *CreateInstanceRequest) GetAutoUseCoupon() *string {
	return s.AutoUseCoupon
}

func (s *CreateInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateInstanceRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CreateInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstanceRequest) GetClusterBackupId() *string {
	return s.ClusterBackupId
}

func (s *CreateInstanceRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CreateInstanceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateInstanceRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *CreateInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateInstanceRequest) GetGlobalInstance() *bool {
	return s.GlobalInstance
}

func (s *CreateInstanceRequest) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *CreateInstanceRequest) GetGlobalSecurityGroupIds() *string {
	return s.GlobalSecurityGroupIds
}

func (s *CreateInstanceRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateInstanceRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstanceRequest) GetParamGroupId() *string {
	return s.ParamGroupId
}

func (s *CreateInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateInstanceRequest) GetPort() *string {
	return s.Port
}

func (s *CreateInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateInstanceRequest) GetReadOnlyCount() *int32 {
	return s.ReadOnlyCount
}

func (s *CreateInstanceRequest) GetRecoverConfigMode() *string {
	return s.RecoverConfigMode
}

func (s *CreateInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceRequest) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateInstanceRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *CreateInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateInstanceRequest) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *CreateInstanceRequest) GetSlaveReadOnlyCount() *int32 {
	return s.SlaveReadOnlyCount
}

func (s *CreateInstanceRequest) GetSlaveReplicaCount() *int32 {
	return s.SlaveReplicaCount
}

func (s *CreateInstanceRequest) GetSrcDBInstanceId() *string {
	return s.SrcDBInstanceId
}

func (s *CreateInstanceRequest) GetTag() []*CreateInstanceRequestTag {
	return s.Tag
}

func (s *CreateInstanceRequest) GetToken() *string {
	return s.Token
}

func (s *CreateInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequest) SetAppendonly(v string) *CreateInstanceRequest {
	s.Appendonly = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v string) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v string) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoUseCoupon(v string) *CreateInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateInstanceRequest) SetBackupId(v string) *CreateInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateInstanceRequest) SetBusinessInfo(v string) *CreateInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateInstanceRequest) SetCapacity(v int64) *CreateInstanceRequest {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetClusterBackupId(v string) *CreateInstanceRequest {
	s.ClusterBackupId = &v
	return s
}

func (s *CreateInstanceRequest) SetConnectionStringPrefix(v string) *CreateInstanceRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateInstanceRequest) SetCouponNo(v string) *CreateInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateInstanceRequest) SetDedicatedHostGroupId(v string) *CreateInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetDryRun(v bool) *CreateInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateInstanceRequest) SetEngineVersion(v string) *CreateInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateInstanceRequest) SetGlobalInstance(v bool) *CreateInstanceRequest {
	s.GlobalInstance = &v
	return s
}

func (s *CreateInstanceRequest) SetGlobalInstanceId(v string) *CreateInstanceRequest {
	s.GlobalInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetGlobalSecurityGroupIds(v string) *CreateInstanceRequest {
	s.GlobalSecurityGroupIds = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceClass(v string) *CreateInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetNetworkType(v string) *CreateInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateInstanceRequest) SetNodeType(v string) *CreateInstanceRequest {
	s.NodeType = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerAccount(v string) *CreateInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetParamGroupId(v string) *CreateInstanceRequest {
	s.ParamGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetPassword(v string) *CreateInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v string) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPort(v string) *CreateInstanceRequest {
	s.Port = &v
	return s
}

func (s *CreateInstanceRequest) SetPrivateIpAddress(v string) *CreateInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetReadOnlyCount(v int32) *CreateInstanceRequest {
	s.ReadOnlyCount = &v
	return s
}

func (s *CreateInstanceRequest) SetRecoverConfigMode(v string) *CreateInstanceRequest {
	s.RecoverConfigMode = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetReplicaCount(v int32) *CreateInstanceRequest {
	s.ReplicaCount = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerAccount(v string) *CreateInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerId(v int64) *CreateInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetRestoreTime(v string) *CreateInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateInstanceRequest) SetSecondaryZoneId(v string) *CreateInstanceRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *CreateInstanceRequest) SetSecurityToken(v string) *CreateInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateInstanceRequest) SetShardCount(v int32) *CreateInstanceRequest {
	s.ShardCount = &v
	return s
}

func (s *CreateInstanceRequest) SetSlaveReadOnlyCount(v int32) *CreateInstanceRequest {
	s.SlaveReadOnlyCount = &v
	return s
}

func (s *CreateInstanceRequest) SetSlaveReplicaCount(v int32) *CreateInstanceRequest {
	s.SlaveReplicaCount = &v
	return s
}

func (s *CreateInstanceRequest) SetSrcDBInstanceId(v string) *CreateInstanceRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateInstanceRequest) SetToken(v string) *CreateInstanceRequest {
	s.Token = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
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

type CreateInstanceRequestTag struct {
	// The keys of the tags that are added to the instance.
	//
	// > 	- **N*	- specifies the serial number of the tag. Up to 20 tags can be added to a single instance. For example, Tag.1.Key specifies the key of the first tag and Tag.2.Key specifies the key of the second tag.
	//
	// > 	- If the key of the tag does not exist, the tag is automatically created.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the tags that are added to the instance.
	//
	// > **N*	- specifies the serial number of the tag. For example, **Tag.1.Value*	- specifies the value of the first tag and **Tag.2.Value*	- specifies the value of the second tag.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTag) SetKey(v string) *CreateInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTag) SetValue(v string) *CreateInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
