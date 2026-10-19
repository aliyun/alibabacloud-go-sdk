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
	SetInstanceEndpointType(v string) *CreateInstanceRequest
	GetInstanceEndpointType() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetMaintainEndTime(v string) *CreateInstanceRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *CreateInstanceRequest
	GetMaintainStartTime() *string
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
	// The AOF parameter settings of the new instance. Valid values:
	//
	// - **yes*	- (default): enables AOF persistence.
	//
	// - **no**: shutdown AOF persistence.
	//
	// > This parameter applies to classic instances. Cloud-native instances do not support specifying the AOF parameter.
	//
	// example:
	//
	// yes
	Appendonly *string `json:"Appendonly,omitempty" xml:"Appendonly,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false*	- (default): disables auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal epoch. Unit: months. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required when the AutoRenew parameter is set to **true**.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to use a voucher. Valid values:
	//
	// 	- **true**: uses a voucher.
	//
	// 	- **false*	- (default): does not use a voucher.
	//
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// You can specify the backup set ID of the source instance in this parameter. The system uses the data saved in the backup set to create the new instance. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query the BackupId. If the source instance is a cluster instance, you must enter the backup set IDs of all shards in the source instance and separate them with commas, such as "10**,11**,15**".
	//
	// > If your instance is a cloud-native architecture instance and a cluster instance, we recommend that you call the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation to query the cluster backup set ID, such as "cb-xx", and then enter it in the ClusterBackupId request parameter to clone the cluster instance. This saves you the trouble of entering the backup set ID of each shard.
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
	// > When you call this operation, you must specify at least one of the Capacity and InstanceClass parameters.
	//
	// example:
	//
	// 16384
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid*	- (default): pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the cluster backup set. Some new cluster architectures support cluster backup set IDs. You can call the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation to obtain the ID.
	//
	// 	- If the architecture supports cluster backup sets, you can enter the specified cluster backup set ID without specifying the BackupId parameter.
	//
	// 	- If the architecture does not support cluster backup sets, specify the backup set IDs of all shards in the source instance in the BackupId parameter and separate them with commas, such as "2158\\*\\*\\*\\*20,2158\\*\\*\\*\\*22".
	//
	// example:
	//
	// cb-hyxdof5x9kqb****
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The prefix of the endpoint. The prefix must be 8 to 40 characters in length and can contain lowercase letters and digits. It must start with a lowercase letter.
	//
	// >
	//
	// > The endpoint format is: <Prefix>.redis.rds.aliyuncs.com.
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
	// The ID of the dedicated cluster. This parameter is required when you create an ApsaraDB for Redis instance in a dedicated cluster.
	//
	// example:
	//
	// dhg-uv4fnk6r7zff****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to perform a dry run on the instance creation request. Valid values:
	//
	// 	- **true**: performs a dry run without creating the instance. The system checks the request parameters, request format, service limits, and available resources. If the request fails the dry run, the system returns an error. If the request passes the dry run, the system returns the DryRunOperation error code.
	//
	// 	- **false*	- (default): sends the normal request and creates the instance after the request passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The engine version of the Redis **classic instance**. Valid values:
	//
	// 	- **2.8*	- (Not recommended. [End-of-life is planned](https://help.aliyun.com/document_detail/2674657.html).)
	//
	// 	- **4.0*	- (Not recommended)
	//
	// 	- **5.0**
	//
	// The engine version of the Redis **cloud-native instance**. Valid values:
	//
	// 	- **5.0**
	//
	// 	- **6.0*	- (Recommended)
	//
	// 	- **7.0**
	//
	// > Default value: **5.0**.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to use the new instance as the first child instance of a distributed instance. You can create a distributed instance in this way. Valid values:
	//
	// 	- **true**: uses the new instance as the first child instance.
	//
	// 	- **false*	- (default): does not use the new instance as the first child instance.
	//
	// > 	- If you set this parameter to **true**, the new instance must be a Tair memory-optimized instance and the DPI engine version must be 5.0.
	//
	// > 	- This parameter applies only to the China site.
	//
	// example:
	//
	// false
	GlobalInstance *bool `json:"GlobalInstance,omitempty" xml:"GlobalInstance,omitempty"`
	// The distributed instance ID. This parameter applies only to the China site.
	//
	// <props="china"> If you want to append the new ApsaraDB for Redis instance as a new child instance to a distributed instance, this parameter is active and required. For more information and console operations, see [Add a child instance to a distributed instance](https://help.aliyun.com/document_detail/106885.html).
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The global IP address whitelist templates of the instance. Separate multiple templates with commas (,). Duplicate templates are not allowed.
	//
	// 	Notice: This parameter applies only to cloud-native instances. Classic instances do not support the whitelist template feature.</notice>
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The instance type. For example, redis.master.small.default indicates a 1 GB Community Edition (classic architecture) instance with the standard architecture and two replicas. For more information, see [Overview of instance types](https://help.aliyun.com/document_detail/26350.html).
	//
	// > When you call this operation, you must specify at least one of the Capacity and InstanceClass parameters.
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The type of the endpoint used when you create a cloud-native dual-zone deployment read/write splitting instance. If you do not explicitly specify this parameter, the default value is AzIndependentEndpoint.
	//
	// - **AzIndependentEndpoint**: **Default value**. Active zone-independent endpoints. The primary and secondary active zones provide independent endpoints. You can use different endpoints to implement nearest access for active zones.
	//
	// - **UnifiedEndpoint**: Unified endpoint. A unified endpoint is provided to access nodes in both the primary and secondary active zones, but cross-zone access may occur.
	//
	// 	Notice: This parameter applies only to cloud-native dual-zone deployment read/write splitting instances. For other types of instances, only active zone-independent endpoints are supported. Even if you specify UnifiedEndpoint, it does not take effect.</notice>
	//
	// 	Notice: The UnifiedEndpoint parameter is currently available only to whitelisted users. If a non-whitelisted user specifies this parameter, an invocation error is returned. To apply for the permissions, submit a ticket.</notice>
	//
	// example:
	//
	// AzIndependentEndpoint
	InstanceEndpointType *string `json:"InstanceEndpointType,omitempty" xml:"InstanceEndpointType,omitempty"`
	// The name of the instance. The name must be 2 to 80 characters in length, and must start with a letter or a Chinese character. It cannot contain the following special characters: `@/:=”<>{[]}` and spaces.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. Valid values:
	//
	// 	- **Redis*	- (default)
	//
	// 	- **Memcache**
	//
	// example:
	//
	// Redis
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The end time of the maintenance window for the instance. Specify the time in the <i>HH:mm</i>Z format. The time must be in UTC. For example, if you want the maintenance window to end at 02:00 (UTC+8), set the value to `18:00Z`.
	//
	// > The interval between the start time and the end time must be at least 1 hour.
	//
	// > If you do not specify this parameter, the default value is 22:00 (UTC), which is 06:00 (UTC+8).
	//
	// example:
	//
	// 07:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window for the instance. Specify the time in the <i>HH:mm</i>Z format. The time must be in UTC. For example, if you want the maintenance window to start at 01:00 (UTC+8), set the value to `17:00Z`.
	//
	// > If you do not specify this parameter, the default value is 18:00 (UTC), which is 02:00 (UTC+8).
	//
	// example:
	//
	// 03:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The network type. Valid values:
	//
	// 	- **VPC**: Virtual Private Cloud (VPC). This is the default value.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The node type. Valid values:
	//
	// 	- **MASTER_SLAVE**: high-availability (two replicas)
	//
	// 	- **STAND_ALONE**: single replica
	//
	// 	- **double**: two replicas
	//
	// 	- **single**: single replica
	//
	// > For cloud-native instances, select **MASTER_SLAVE*	- or **STAND_ALONE**. For classic instances, select **double*	- or **single**.
	//
	// example:
	//
	// STAND_ALONE
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The globally unique ID of the parameter template.	Notice: This parameter applies only to cloud-native instances.</notice>
	//
	// example:
	//
	// rpg-test**
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The password of the instance. The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, special characters, and digits. The allowed special characters include `!@#$%^&*()_+-=`.
	//
	// example:
	//
	// Pass!123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription duration. Unit: months. Valid values: **1*	- to **9**, **12**, **24**, **36**, and **60**.
	//
	// > This parameter is required and takes effect only when the ChargeType parameter is set to **PrePaid**.
	//
	// example:
	//
	// 12
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The service port of the instance. Valid values: **1*	- to **65535**. Default value: **6379**.
	//
	// example:
	//
	// 6379
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The internal network IP address of the new instance.
	//
	// > The internal network IP address must be within the vSwitch CIDR block to which the instance belongs.
	//
	// example:
	//
	// 172.16.0.***
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The number of read-only nodes in the primary zone. This parameter applies only to cloud-native read/write splitting instances.
	//
	// 	- If the instance uses the standard architecture, valid values are 1 to 9.
	//
	// 	- If the instance uses the cluster architecture, valid values are 1 to 4, which indicates the number of read-only nodes per shard.
	//
	// > If you create a multi-zone instance, you can use this parameter and the SlaveReadOnlyCount parameter to specify the number of read-only nodes in the primary and secondary zones.
	//
	// > - If the instance uses the standard architecture, the sum of this parameter and the SlaveReadOnlyCount parameter cannot exceed 9.
	//
	// > - If the instance uses the cluster architecture, the sum of this parameter and the SlaveReadOnlyCount parameter cannot exceed 4.
	//
	// example:
	//
	// 2
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// Specifies whether to restore the account, kernel parameters, and whitelist information from the original backup set when you create an instance by using a specified backup set. For example, to restore the account information, set the value to `account`.
	//
	// Default value: empty, which indicates that the account, kernel parameters, and whitelist information are not restored from the original backup set.
	//
	// > This parameter applies only to cloud-native instances, and the original backup set must contain the account, kernel parameters, and whitelist information. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query whether the RecoverConfigMode parameter of the specified backup set contains the preceding information.
	//
	// example:
	//
	// whitelist,config,account
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of secondary nodes in the primary zone. This parameter applies only to cloud-native cluster instances with multiple replicas. You can use this parameter to specify a custom number of secondary nodes. Valid values: 1 to 4.
	//
	// > If you create a multi-zone instance, you can use this parameter and the SlaveReplicaCount parameter to specify the number of secondary nodes in the primary and secondary zones. The sum of this parameter and the SlaveReplicaCount parameter cannot exceed 4.
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
	// The point in time to which you want to restore the source instance. If the data flashback feature is enabled for the source instance, you can specify any point in time within the backup retention period of the source instance. The system uses the backup data of the source instance at the specified point in time to create the new instance. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// example:
	//
	// 2019-06-19T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The secondary zone ID. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query the most recent zone list.
	//
	// > The value of this parameter cannot be the same as the value of the ZoneId parameter, and cannot be the ID of a multi-zone deployment.
	//
	// example:
	//
	// cn-hangzhou-g
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards. This parameter applies only to cloud-native instances. You can use this parameter to specify a custom number of shards.
	//
	// - 1: creates a non-cluster architecture.
	//
	// - A value greater than 1: creates a cluster architecture.
	//
	// example:
	//
	// 4
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The number of read-only nodes in the secondary zone.
	//
	// example:
	//
	// 2
	SlaveReadOnlyCount *int32 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// The number of secondary nodes in the secondary zone.
	//
	// example:
	//
	// 2
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// If you want to create a new instance based on the backup set of an instance, specify the source instance ID in this parameter first.
	//
	// > Then, specify the backup set or point in time by using the BackupId, ClusterBackupId (recommended for cloud-native cluster architecture), or RestoreTime parameter. This parameter must be used in combination with one of the preceding three parameters.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The tag information of the instance.
	Tag []*CreateInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token is case-sensitive and can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The primary active zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent active zone list, and use this parameter to specify the active zone where you want to create the instance.
	//
	// > You can also specify the SecondaryZoneId parameter to set the secondary active zone. The deployment of the primary and secondary nodes is in the specified primary and secondary active zones to implement a primary/secondary architecture. For example, you can set the ZoneId parameter to "cn-hangzhou-h" and the SecondaryZoneId parameter to "cn-hangzhou-g".
	//
	// example:
	//
	// cn-hangzhou-h
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

func (s *CreateInstanceRequest) GetInstanceEndpointType() *string {
	return s.InstanceEndpointType
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *CreateInstanceRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
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

func (s *CreateInstanceRequest) SetInstanceEndpointType(v string) *CreateInstanceRequest {
	s.InstanceEndpointType = &v
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

func (s *CreateInstanceRequest) SetMaintainEndTime(v string) *CreateInstanceRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *CreateInstanceRequest) SetMaintainStartTime(v string) *CreateInstanceRequest {
	s.MaintainStartTime = &v
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
	// The key of the tag.
	//
	// > 	- **N*	- indicates the sequence number of the tag key. You can attach up to 20 tags to a single instance. For example, Tag.1.Key indicates the key of the first tag, and Tag.2.Key indicates the key of the second tag.
	//
	// > 	- If the key of the tag does not exist, the system performs automatic creation.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// > **N*	- indicates the sequence number of the tag value. For example, **Tag.1.Value*	- indicates the value of the first tag, and **Tag.2.Value*	- indicates the value of the second tag.
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
