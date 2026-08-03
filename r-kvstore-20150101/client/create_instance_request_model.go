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
	// The Append Only File (AOF) persistence parameter settings for the new instance. Valid values:
	//
	// - **yes*	- (default): enables AOF persistence.
	//
	// - **no**: disables AOF persistence.
	//
	// > This parameter is applicable to classic instances. Cloud-native instances do not support specifying the AOF parameter.
	//
	// example:
	//
	// yes
	Appendonly *string `json:"Appendonly,omitempty" xml:"Appendonly,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false*	- (default): does not enable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal epoch. Unit: months. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required when **AutoRenew*	- is set to **true**.
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
	// The ID of the backup set of the source instance. The system uses the data stored in the backup set to create the instance. You can invoke [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) to query the BackupId. If the source instance is a cluster instance, specify the backup set IDs of all shards of the source instance, separated by commas (,). Example: "10\\*\\*,11\\*\\*,15\\*\\*".
	//
	// > If your instance is a cloud-native architecture cluster instance, use [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) to query the cluster backup set ID, such as "cb-xx", and specify it in the ClusterBackupId request parameter to clone the cluster instance. This eliminates the need to specify individual shard backup set IDs.
	//
	// example:
	//
	// 111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The activity ID and business information.
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
	// The cluster backup set ID, which is supported by some new cluster architecture instances. You can call [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) to obtain the ID.
	//
	// 	- If supported, specify the cluster backup set ID. You do not need to specify the **BackupId*	- parameter.
	//
	// 	- If not supported, specify the backup set IDs of all shards of the source instance in the BackupId parameter, separated by commas (,). Example: "2158\\*\\*\\*\\*20,2158\\*\\*\\*\\*22".
	//
	// example:
	//
	// cb-hyxdof5x9kqb****
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The prefix of the endpoint. The prefix must consist of lowercase letters and digits, start with a lowercase letter, and be 8 to 40 characters in length.
	//
	// >
	//
	// > The endpoint is in the format of: <prefix>.redis.rds.aliyuncs.com.
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
	// The ID of the dedicated cluster. This parameter is required when you create an instance in a dedicated cluster.
	//
	// example:
	//
	// dhg-uv4fnk6r7zff****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to perform a dry run for this instance creation request. Valid values:
	//
	// 	- **true**: performs a dry run without creating the instance. The system checks items such as the request parameters, request format, service limits, and available resources. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// 	- **false*	- (default): sends the request. After the request passes the check, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Redis **classic*	- engine version. Valid values:
	//
	// 	- **2.8*	- (not recommended, [planned for end of support](https://help.aliyun.com/document_detail/2674657.html))
	//
	// 	- **4.0*	- (not recommended)
	//
	// 	- **5.0**
	//
	// Redis **cloud-native*	- engine version. Valid values:
	//
	// 	- **5.0**
	//
	// 	- **6.0*	- (recommended)
	//
	// 	- **7.0**
	//
	// > Default value: **5.0**.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to use the new instance as the first child instance of a distributed instance. This allows you to create a distributed instance. Valid values:
	//
	// 	- **true**: uses the instance as the first child instance.
	//
	// 	- **false*	- (default): does not use the instance as the first child instance.
	//
	// > 	- To set this parameter to **true**, the new instance must be a Tair memory-optimized instance with a database DPI engine version of 5.0.
	//
	// > 	- This parameter is applicable only to Chinese site (aliyun.com).
	//
	// example:
	//
	// false
	GlobalInstance *bool `json:"GlobalInstance,omitempty" xml:"GlobalInstance,omitempty"`
	// The instance ID of the distributed instance. This parameter is applicable only to Chinese site (aliyun.com).
	//
	// <props="china"> To append the new Redis instance as a child instance of a distributed instance, this parameter is active and required. For more information and console operations, see [Add a child instance to a distributed instance](https://help.aliyun.com/document_detail/106885.html).
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The global IP whitelist templates for the instance. Separate multiple templates with commas (,). Duplicates are not allowed.
	//
	// 	Notice: This parameter is applicable only to cloud-native instances. Classic instances do not support the whitelist template feature.</notice>
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The instance type. For example, redis.master.small.default specifies a Community Edition (classic) standard architecture dual-replica 1 GB instance. For more information, see [Instance type overview](https://help.aliyun.com/document_detail/26350.html).
	//
	// > You must specify at least one of the **Capacity*	- and **InstanceClass*	- parameters when you call this operation.
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The endpoint type used when you create a cloud-native dual-zone deployment read/write splitting instance. If this parameter is not explicitly committed, the default value is AzIndependentEndpoint.
	//
	// - **AzIndependentEndpoint**: **default value**. Zone-independent endpoints. The primary and secondary zones provide independent endpoints, which allow nearest access through different endpoints.
	//
	// - **UnifiedEndpoint**: unified endpoint. A unified endpoint is provided to access nodes in both the primary and secondary zones, but cross-zone access may occur.
	//
	// 	Notice: This parameter is applicable only to cloud-native dual-zone deployment read/write splitting instances. For other instance types, only zone-independent endpoints are supported. Even if UnifiedEndpoint is specified, it does not take effect.</notice>
	//
	// 	Notice: The UnifiedEndpoint option is available only to users on the whitelist. If you are not on the whitelist and specify this parameter, the invocation returns an error. To request access, submit a ticket.</notice>
	//
	// example:
	//
	// AzIndependentEndpoint
	InstanceEndpointType *string `json:"InstanceEndpointType,omitempty" xml:"InstanceEndpointType,omitempty"`
	// The name of the instance. The name must be 2 to 80 characters in length and must start with a letter or a Chinese character. The name cannot contain `@/:="<>{[]}` or spaces.
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
	// The end time of the maintenance window. Specify the time in the <i>HH:mm</i>Z format in UTC. For example, to set the end time to 02:00 (UTC+8), specify `18:00Z`.
	//
	// > The interval between the start time and end time must be at least 1 hour.
	//
	// > If this parameter is not specified, the default value is 06:00 (UTC+8), which is 22:00Z in UTC.
	//
	// example:
	//
	// 07:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. Specify the time in the <i>HH:mm</i>Z format in UTC. For example, to set the start time to 01:00 (UTC+8), specify `17:00Z`.
	//
	// > If this parameter is not specified, the default value is 02:00 (UTC+8), which is 18:00Z in UTC.
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
	// 	- **MASTER_SLAVE**: high availability (dual-replica)
	//
	// 	- **STAND_ALONE**: single replica
	//
	// 	- **double**: dual-replica
	//
	// 	- **single**: single replica
	//
	// > For cloud-native instances, set this parameter to **MASTER_SLAVE*	- or **STAND_ALONE**. For classic instances, set this parameter to **double*	- or **single**.
	//
	// example:
	//
	// STAND_ALONE
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template. The ID is globally unique.	Notice: This parameter is applicable only to cloud-native instances.</notice>
	//
	// example:
	//
	// rpg-test**
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The password of the instance. The password must be 8 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, special characters, and digits. The following special characters are supported: `!@#$%^&*()_+-=`.
	//
	// example:
	//
	// Pass!123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription period. Unit: months. Valid values: **1*	- to **9**, **12**, **24**, **36**, and **60**.
	//
	// > This parameter is available and required only when **ChargeType*	- is set to **PrePaid**.
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
	// The number of read-only nodes in the primary zone. This parameter is applicable only to cloud-native read/write splitting instances.
	//
	// 	- For standard architecture instances, valid values are 1 to 9.
	//
	// 	- For cluster architecture instances, valid values are 1 to 4, which specifies the number of read-only nodes per data shard.
	//
	// > If you create a multi-zone instance, you can use this parameter together with the SlaveReadOnlyCount parameter to customize the number of read-only nodes in the primary and secondary zones.
	//
	// > - For standard architecture instances, the sum of this parameter and SlaveReadOnlyCount cannot exceed 9.
	//
	// > - For cluster architecture instances, the sum of this parameter and SlaveReadOnlyCount cannot exceed 4.
	//
	// example:
	//
	// 2
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// Specifies whether to restore the account, kernel parameter (config), or whitelist information from the original backup set when you create an instance from a specified backup set. For example, to restore account information, set this parameter to `account`.
	//
	// The default value is empty, which indicates that the account, kernel parameter, and whitelist information is not restored from the original backup set.
	//
	// > This parameter is applicable only to cloud-native instances, and the original backup set must contain the account, kernel parameter, or whitelist information. You can call [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) to check whether the RecoverConfigMode parameter of the specified backup set contains the preceding information.
	//
	// example:
	//
	// whitelist,config,account
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) to query available regions. Use this parameter to specify the region in which to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of replica nodes in the primary zone. This parameter is applicable only to cloud-native cluster multi-replica instances. You can use this parameter to customize the number of replica nodes. Valid values: 1 to 4.
	//
	// > If you create a multi-zone instance, you can use this parameter together with the SlaveReplicaCount parameter to customize the number of replica nodes in the primary and secondary zones. The sum of this parameter and the SlaveReplicaCount parameter cannot exceed 4.
	//
	// example:
	//
	// 2
	ReplicaCount *int32 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-resourcegroupid1
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// If flashback is enabled for the source instance, you can specify a point in time within the backup retention period. The system uses the backup data of the source instance at the specified point in time to create the instance. Specify the time in the ISO 8601 standard in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// example:
	//
	// 2019-06-19T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The secondary zone ID. You can call [DescribeZones](https://help.aliyun.com/document_detail/473764.html) to query available zones.
	//
	// > The value of this parameter must be different from the value of ZoneId. You cannot set this parameter to the ID of a multi-zone.
	//
	// example:
	//
	// cn-hangzhou-g
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards. This parameter is applicable only to cloud-native instances. You can use this parameter to customize the number of shards.
	//
	// - 1: creates a non-cluster instance.
	//
	// - A value greater than 1: creates a cluster instance.
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
	// The number of replica nodes in the secondary zone.
	//
	// example:
	//
	// 2
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// To create an instance from a backup set of an existing instance, specify the instance ID of the source instance in this parameter.
	//
	// > Then use the **BackupId**, **ClusterBackupId*	- (recommended for cloud-native cluster instances), or **RestoreTime*	- parameter to specify the backup set or point in time. This parameter must be used together with one of the preceding three parameters. The value is a string, not an array.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The tags of the instance.
	Tag []*CreateInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. The token value is generated by the client and must be unique among different requests. The token is case-sensitive and cannot exceed 64 ASCII characters in length.
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
	// The primary zone ID. You can invoke [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) to query active zones. Use this parameter to specify the zone in which to create the instance.
	//
	// > You can also specify the SecondaryZoneId parameter to set the secondary zone. The primary and secondary nodes are deployed in the specified primary and secondary zones respectively, which implements a dual-center primary/secondary architecture in the same city. For example, set ZoneId to "cn-hangzhou-h" and SecondaryZoneId to "cn-hangzhou-g".
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
	// > 	- **N*	- specifies the sequence number of the tag. A maximum of 20 tags can be attached to a single instance. For example, Tag.1.Key specifies the key of the first tag, and Tag.2.Key specifies the key of the second tag.
	//
	// > 	- If the tag key does not exist, the tag is automatically created.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// > **N*	- specifies the sequence number of the tag. For example, **Tag.1.Value*	- specifies the value of the first tag, and **Tag.2.Value*	- specifies the value of the second tag.
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
