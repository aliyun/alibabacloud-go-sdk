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
	// Specifies whether to enable AOF persistence for the new instance. Valid values:
	//
	// - **yes*	- (default): Enables AOF persistence.
	//
	// - **no**: Disables AOF persistence.
	//
	// > This parameter is available only for classic edition instances. AOF persistence cannot be configured for cloud native edition instances at creation.
	//
	// example:
	//
	// yes
	Appendonly *string `json:"Appendonly,omitempty" xml:"Appendonly,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// - **true**: Enables auto-renewal.
	//
	// - **false*	- (default): Disables auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration, in months. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required when **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **true**: Uses a coupon.
	//
	// - **false*	- (default): Does not use a coupon.
	//
	// example:
	//
	// false
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The ID of the backup that you want to use to create the new instance. You can obtain backup IDs by calling the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation. If the source instance has a cluster architecture, you must specify the backup IDs of all its shards, separated by commas (for example, "10\\*\\*,11\\*\\*,15\\*\\*").
	//
	// > If your source instance is a cloud native cluster instance, it is recommended to call [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) to get a cluster backup ID (for example, "cb-xx") and use the `ClusterBackupId` parameter instead. This avoids the need to specify the backup ID for each shard.
	//
	// example:
	//
	// 111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The campaign ID or business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The storage capacity of the instance, in MB.
	//
	// > You must specify either the **Capacity*	- or the **InstanceClass*	- parameter.
	//
	// example:
	//
	// 16384
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method. Valid values:
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid*	- (default): pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the cluster backup. You can get this ID by calling the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation. This parameter is available for some cloud native cluster instances.
	//
	// - This parameter is mutually exclusive with `BackupId`.
	//
	// - If this parameter is not available for your instance, you must specify the backup ID of each shard in the `BackupId` parameter (for example, "2158\\*\\*\\*\\*20,2158\\*\\*\\*\\*22").
	//
	// example:
	//
	// cb-hyxdof5x9kqb****
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The prefix of the connection string. The prefix must be 8 to 40 characters long, start with a lowercase letter, and contain only lowercase letters and digits.
	//
	// > The full connection string is in the format: \\<prefix>.redis.rds.aliyuncs.com.
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
	// The ID of the dedicated host group. This parameter is required when you create a Redis instance in a dedicated host group.
	//
	// example:
	//
	// dhg-uv4fnk6r7zff****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Checks the request for validity without creating the instance. The system verifies required parameters, request format, and service limits. If the request is valid, the `DryRunOperation` error code is returned. If the request is invalid, an error message is returned.
	//
	// - **false*	- (default): Sends the request. If the request is valid, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The Redis engine version. Valid values for **classic edition*	- instances:
	//
	// - **2.8*	- (Not recommended. [Support for this version is scheduled to be discontinued](https://help.aliyun.com/document_detail/2674657.html).)
	//
	// - **4.0*	- (Not recommended.)
	//
	// - **5.0**
	//
	// Valid values for **cloud native edition*	- instances:
	//
	// - **5.0**
	//
	// - **6.0*	- (Recommended)
	//
	// - **7.0**
	//
	// > The default value is **5.0**.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to create the new instance as the first child instance of a distributed instance. Valid values:
	//
	// - **true**: Creates the instance as the first child instance.
	//
	// - **false*	- (default): Does not create the instance as the first child instance.
	//
	// > 	- If you set this parameter to **true**, the new instance must be a Tair memory-enhanced instance that runs Redis 5.0.
	//
	// >
	//
	// > 	- This parameter is available only in Chinese mainland.
	//
	// example:
	//
	// false
	GlobalInstance *bool `json:"GlobalInstance,omitempty" xml:"GlobalInstance,omitempty"`
	// The ID of the distributed instance. This parameter is available only in Chinese mainland.
	//
	// <props="china">
	//
	// This parameter is required to add the new instance as a child of a distributed instance. For more information and the console procedure, see [Add a child instance to a distributed instance](https://help.aliyun.com/document_detail/106885.html).
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The IDs of the security groups to associate with the instance. You can specify multiple security group IDs, separated by commas (,). IDs cannot be repeated.
	//
	// 	Notice: This parameter is available only for cloud native edition instances. Security groups are not supported for classic edition instances.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The instance type. For example, `redis.master.small.default` specifies a 1 GB Community Edition (classic edition) instance with a standard, dual-replica architecture. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/26350.html).
	//
	// > You must specify either the **Capacity*	- or the **InstanceClass*	- parameter.
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The connection endpoint type. This parameter is applicable only when you create a dual-zone, read/write splitting instance of the cloud native edition. If this parameter is not specified, `AzIndependentEndpoint` is used. Valid values:
	//
	// - **AzIndependentEndpoint**: (**Default**) Zone-Independent Endpoint. The primary and secondary zones each provide an independent connection string for zone-local access.
	//
	// - **UnifiedEndpoint**: Unified Endpoint. Provides a single connection string to access nodes in both zones, which may result in cross-zone access.
	//
	// 	Notice:
	//
	// This parameter is applicable only to dual-zone, read/write splitting instances of the cloud native edition. For other instance types, only zone-independent endpoints are supported, and specifying `UnifiedEndpoint` has no effect.
	//
	//
	//
	// 	Notice:
	//
	// The `UnifiedEndpoint` parameter is currently available only to allowlisted users. API calls will fail if you are not on the allowlist. To be added to the allowlist, submit a ticket.
	//
	// example:
	//
	// AzIndependentEndpoint
	InstanceEndpointType *string `json:"InstanceEndpointType,omitempty" xml:"InstanceEndpointType,omitempty"`
	// The name of the instance. The name must be 2 to 80 characters long, start with a letter (uppercase or lowercase) or a Chinese character, and not contain spaces or the characters `@/:=”<>{[]}`.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. Valid values:
	//
	// - **Redis*	- (default)
	//
	// - **Memcache**
	//
	// example:
	//
	// Redis
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The end time of the maintenance window. Specify the time in the *HH:mm*Z format (UTC). For example, to set the end time to 02:00 (UTC+8), specify `18:00Z`.
	//
	// > The duration of the maintenance window must be at least one hour.
	//
	// > If this parameter is not specified, the maintenance window ends at 06:00 (UTC+8), which is 22:00 (UTC).
	//
	// example:
	//
	// 07:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start of the maintenance window. Specify the time in the *HH:mm*Z format (UTC). For example, to set the start time to 01:00 (UTC+8), specify `17:00Z`.
	//
	// > If this parameter is not specified, the maintenance window starts at 02:00 (UTC+8), which is 18:00 (UTC).
	//
	// example:
	//
	// 03:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The network type. Valid value:
	//
	// - **VPC**: Deploys the instance in a Virtual Private Cloud. This is the default value.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The node type. Valid values:
	//
	// - **MASTER_SLAVE**: high-availability (primary-replica)
	//
	// - **STAND_ALONE**: standalone (single-node)
	//
	// - **double**: primary-replica
	//
	// - **single**: standalone (single-node)
	//
	// > Set this parameter to **MASTER_SLAVE*	- or **STAND_ALONE*	- for cloud native edition instances. Set this parameter to **double*	- or **single*	- for classic edition instances.
	//
	// example:
	//
	// STAND_ALONE
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter group. This ID must be globally unique.	Notice:  This parameter is available only for cloud native edition instances.
	//
	// example:
	//
	// rpg-test**
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The password for the instance. The password must be 8 to 32 characters long and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The allowed special characters are `!@#$%^&*()_+-=`.
	//
	// example:
	//
	// Pass!123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription duration, in months. Valid values: **1*	- to **9**, **12**, **24**, **36**, and **60**.
	//
	// > This parameter is available and required only when **ChargeType*	- is set to **PrePaid**.
	//
	// example:
	//
	// 12
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The service port of the instance. The port number must be between **1*	- and **65535**. The default value is **6379**.
	//
	// example:
	//
	// 6379
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the new instance.
	//
	// > The IP address must be within the CIDR block of the specified vSwitch.
	//
	// example:
	//
	// 172.16.0.***
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The number of read-only replicas in the primary zone. This parameter is available only when creating a read/write splitting instance of the cloud native edition.
	//
	// - For a standard-architecture instance, the value must be an integer from 1 to 9.
	//
	// - For a cluster-architecture instance, the value must be an integer from 1 to 4. This specifies the number of read-only replicas for each data shard.
	//
	// > If you create a multi-zone instance, you can use this parameter and `SlaveReadOnlyCount` to customize the number of read-only replicas in the primary and secondary zones.
	//
	// >
	//
	// > - The sum of this parameter and `SlaveReadOnlyCount` cannot exceed 9 for a standard-architecture instance.
	//
	// >
	//
	// > - The sum of this parameter and `SlaveReadOnlyCount` cannot exceed 4 for a cluster-architecture instance.
	//
	// example:
	//
	// 2
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// Specifies which configurations to restore from the backup when creating an instance. Valid values include `account`, `config`, and `whitelist`. For example, to restore account settings, specify `account`. To restore multiple configurations, separate them with commas.
	//
	// By default, this parameter is empty, which means no configurations are restored.
	//
	// > This parameter is applicable only to cloud native edition instances. The source backup must contain the specified configurations. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation and check the `RecoverConfigMode` field in the response to determine which configurations a backup contains.
	//
	// example:
	//
	// whitelist,config,account
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
	// The ID of the region in which to create the instance. Call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to get a list of region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of replicas in the primary zone. This parameter is available only for multi-replica cluster instances of the cloud native edition. You can specify a value from 1 to 4.
	//
	// > When creating a multi-zone instance, you can use this parameter and `SlaveReplicaCount` to customize the number of replicas in the primary and secondary zones. The sum of `ReplicaCount` and `SlaveReplicaCount` cannot exceed 4.
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
	// The point in time to which you want to restore data, specified in the *yyyy-MM-dd*T*HH:mm:ss*Z (UTC) format.
	//
	// example:
	//
	// 2019-06-19T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The ID of the secondary zone. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query the latest list of zones.
	//
	// > The value of this parameter cannot be the same as the value of the `ZoneId` parameter, and you cannot specify a multi-zone ID.
	//
	// example:
	//
	// cn-hangzhou-g
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards. This parameter is available only for cloud native edition instances.
	//
	// - A value of **1*	- creates an instance with a standard architecture.
	//
	// - A value greater than **1*	- creates an instance with a cluster architecture.
	//
	// example:
	//
	// 4
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The number of read-only replicas in the secondary zone.
	//
	// example:
	//
	// 2
	SlaveReadOnlyCount *int32 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// The number of replicas in the secondary zone.
	//
	// example:
	//
	// 2
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// To create an instance from a backup, specify the ID of the source instance.
	//
	// > This parameter must be used in conjunction with one of the following parameters: **BackupId**, **ClusterBackupId*	- (recommended for cloud native, cluster-architecture instances), or **RestoreTime**.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The tags of the instance.
	Tag []*CreateInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// A client-generated token to ensure the idempotence of the request. The token must be unique across requests, case-sensitive, and cannot exceed 64 ASCII characters.
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
	// The ID of the primary zone for the instance. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473763.html) operation to query available zones.
	//
	// > You can also specify a secondary zone by using the `SecondaryZoneId` parameter. The primary and replica nodes are then deployed in the specified primary and secondary zones to create a dual-zone architecture for in-city disaster recovery. For example, you can set the `ZoneId` parameter to "cn-hangzhou-h" and the `SecondaryZoneId` parameter to "cn-hangzhou-g".
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
	// > - `N` represents the sequence number of the tag, from 1 to 20. You can add a maximum of 20 tags to an instance.
	//
	// >
	//
	// > - If the tag key does not exist, it is automatically created.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value for tag `N`.
	//
	// > The N in **Tag.N.Value*	- specifies the sequence number of the tag. For example, **Tag.1.Value*	- specifies the value of the first tag, and **Tag.2.Value*	- specifies the value of the second tag.
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
