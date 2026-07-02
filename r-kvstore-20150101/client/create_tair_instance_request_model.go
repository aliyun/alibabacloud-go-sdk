// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateTairInstanceRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *CreateTairInstanceRequest
	GetAutoRenew() *string
	SetAutoRenewPeriod(v string) *CreateTairInstanceRequest
	GetAutoRenewPeriod() *string
	SetAutoUseCoupon(v string) *CreateTairInstanceRequest
	GetAutoUseCoupon() *string
	SetBackupId(v string) *CreateTairInstanceRequest
	GetBackupId() *string
	SetBusinessInfo(v string) *CreateTairInstanceRequest
	GetBusinessInfo() *string
	SetChargeType(v string) *CreateTairInstanceRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateTairInstanceRequest
	GetClientToken() *string
	SetClusterBackupId(v string) *CreateTairInstanceRequest
	GetClusterBackupId() *string
	SetConnectionStringPrefix(v string) *CreateTairInstanceRequest
	GetConnectionStringPrefix() *string
	SetCouponNo(v string) *CreateTairInstanceRequest
	GetCouponNo() *string
	SetDryRun(v bool) *CreateTairInstanceRequest
	GetDryRun() *bool
	SetEngineVersion(v string) *CreateTairInstanceRequest
	GetEngineVersion() *string
	SetGlobalInstanceId(v string) *CreateTairInstanceRequest
	GetGlobalInstanceId() *string
	SetGlobalSecurityGroupIds(v string) *CreateTairInstanceRequest
	GetGlobalSecurityGroupIds() *string
	SetInstanceClass(v string) *CreateTairInstanceRequest
	GetInstanceClass() *string
	SetInstanceEndpointType(v string) *CreateTairInstanceRequest
	GetInstanceEndpointType() *string
	SetInstanceName(v string) *CreateTairInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateTairInstanceRequest
	GetInstanceType() *string
	SetMaintainEndTime(v string) *CreateTairInstanceRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *CreateTairInstanceRequest
	GetMaintainStartTime() *string
	SetOwnerAccount(v string) *CreateTairInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTairInstanceRequest
	GetOwnerId() *int64
	SetParamGroupId(v string) *CreateTairInstanceRequest
	GetParamGroupId() *string
	SetPassword(v string) *CreateTairInstanceRequest
	GetPassword() *string
	SetPeriod(v int32) *CreateTairInstanceRequest
	GetPeriod() *int32
	SetPort(v int32) *CreateTairInstanceRequest
	GetPort() *int32
	SetPrivateIpAddress(v string) *CreateTairInstanceRequest
	GetPrivateIpAddress() *string
	SetReadOnlyCount(v int32) *CreateTairInstanceRequest
	GetReadOnlyCount() *int32
	SetRecoverConfigMode(v string) *CreateTairInstanceRequest
	GetRecoverConfigMode() *string
	SetRegionId(v string) *CreateTairInstanceRequest
	GetRegionId() *string
	SetReplicaCount(v int32) *CreateTairInstanceRequest
	GetReplicaCount() *int32
	SetResourceGroupId(v string) *CreateTairInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateTairInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTairInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CreateTairInstanceRequest
	GetRestoreTime() *string
	SetSecondaryZoneId(v string) *CreateTairInstanceRequest
	GetSecondaryZoneId() *string
	SetSecurityToken(v string) *CreateTairInstanceRequest
	GetSecurityToken() *string
	SetShardCount(v int32) *CreateTairInstanceRequest
	GetShardCount() *int32
	SetShardType(v string) *CreateTairInstanceRequest
	GetShardType() *string
	SetSlaveReadOnlyCount(v int32) *CreateTairInstanceRequest
	GetSlaveReadOnlyCount() *int32
	SetSlaveReplicaCount(v int32) *CreateTairInstanceRequest
	GetSlaveReplicaCount() *int32
	SetSrcDBInstanceId(v string) *CreateTairInstanceRequest
	GetSrcDBInstanceId() *string
	SetStorage(v int32) *CreateTairInstanceRequest
	GetStorage() *int32
	SetStorageType(v string) *CreateTairInstanceRequest
	GetStorageType() *string
	SetTag(v []*CreateTairInstanceRequestTag) *CreateTairInstanceRequest
	GetTag() []*CreateTairInstanceRequestTag
	SetVSwitchId(v string) *CreateTairInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateTairInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateTairInstanceRequest
	GetZoneId() *string
}

type CreateTairInstanceRequest struct {
	// Specifies whether to enable automatic payment. Set the value to **true**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// - **true**: Enable auto-renewal.
	//
	// - **false*	- (default): Disable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: month. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required only when the **AutoRenew*	- parameter is set to **true**.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **true**: Use a coupon.
	//
	// - **false*	- (default): Do not use a coupon.
	//
	// example:
	//
	// true
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The ID of the backup set from the source instance. The system creates a new instance based on the data in this backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query the backup set ID. If the source instance is a cluster instance, you must specify the backup ID for each shard, separated by commas, for example, "10\\*\\*,11\\*\\*,15\\*\\*".
	//
	// > If your instance is a cloud-native cluster instance, we recommend that you call the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation to query the cluster backup ID, such as `cb-xx`. Then, specify the cluster backup ID for the `ClusterBackupId` parameter to clone the cluster instance. This avoids the need to specify the backup ID of each shard.
	//
	// example:
	//
	// 2158****20
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The business information. This can be the ID of a promotion or a business context.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PrePaid*	- (default): The subscription billing method.
	//
	// - **PostPaid**: The pay-as-you-go billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// A client-generated token that ensures the idempotence of the request. The token must be unique among different requests. It is case-sensitive and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cluster backup set. Some instances that use the cluster architecture support cluster backup sets. You can call the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation to query for cluster backup set IDs.
	//
	// - If this feature is supported, you can specify this parameter and leave the **BackupId*	- parameter empty.
	//
	// - If this feature is not supported, you must specify the backup ID of each shard of the source instance for the `BackupId` parameter. Separate the backup IDs with commas, for example, "2158\\*\\*\\*\\*20,2158\\*\\*\\*\\*22".
	//
	// example:
	//
	// cb-hyxdof5x9kqb****
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The prefix of the connection string. It must start with a lowercase letter, consist of lowercase letters and digits, and be 8 to 40 characters in length.
	//
	// > The full connection string is in the format of `<prefix>-<instance ID>.redis.rds.aliyuncs.com`.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// Specifies whether to perform a precheck for this request. Valid values:
	//
	// - **true**: Performs a precheck and does not create the instance. The system checks the request parameters, request format, service limits, and available inventory. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request and creates the instance after the request passes the precheck.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The database version. Default value: **1.0**. The valid values depend on the Tair instance series:
	//
	// - **tair_rdb**: Tair memory-enhanced instances are compatible with Redis 5.0, Redis 6.0, and Redis 7.0. Set the value to **5.0**, **6.0**, or **7.0**.
	//
	// - **tair_scm**: Tair persistent memory-optimized instances are compatible with Redis 6.0. Set the value to **1.0**.
	//
	// - **tair_essd**: Tair disk-based instances (ESSD/SSD) are compatible with Redis 6.0. Set the value to **1.0*	- to create an ESSD-based instance or **2.0*	- to create an SSD-based instance.
	//
	// example:
	//
	// 1.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to create the instance as a child instance in a distributed instance. By using this parameter, you can create a distributed instance.
	//
	// - To create the first child instance, set this parameter to **true**.
	//
	// - To create the second or third child instance, specify the ID of the distributed instance, such as `gr-bp14rkqrhac****`.
	//
	// - If you do not want to create a distributed instance, do not specify this parameter.
	//
	// > To be created as a child instance of a distributed instance, the new instance must be a Tair memory-enhanced instance.
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The IDs of the global IP whitelist templates for the instance. To specify multiple template IDs, separate them with commas. The IDs cannot be repeated.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The instance type. For more information, see the following topics:
	//
	// - [Memory-enhanced instance types](https://help.aliyun.com/document_detail/2527112.html)
	//
	// - [Persistent memory-optimized instance types](https://help.aliyun.com/document_detail/2527110.html)
	//
	// - [Disk-based instance types](https://help.aliyun.com/document_detail/2527111.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// tair.scm.standard.4m.32d
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The type of connection string to use when creating a cloud-native, dual-zone instance with the read/write splitting architecture. If you do not specify this parameter, the default value `AzIndependentEndpoint` is used.
	//
	// - **AzIndependentEndpoint*	- (**default**): Zone-specific connection string. The primary and secondary zones each provide an independent connection string, allowing clients to connect to the nearest zone.
	//
	// - **UnifiedEndpoint**: Unified connection string. A single connection string is provided to access nodes in both the primary and secondary zones, but this may cause cross-zone access.
	//
	// 	Notice:
	//
	// This parameter applies only to cloud-native, dual-zone instances with the read/write splitting architecture. Other instance types support only zone-specific connection strings.
	//
	//
	//
	// 	Notice:
	//
	// The `UnifiedEndpoint` option is available only to users on a whitelist. If a non-whitelisted user specifies this value, the request fails. To request access, submit a ticket.
	//
	// example:
	//
	// AzIndependentEndpoint
	InstanceEndpointType *string `json:"InstanceEndpointType,omitempty" xml:"InstanceEndpointType,omitempty"`
	// The name of the instance. The name must meet the following requirements:
	//
	// - It must be 2 to 80 characters in length.
	//
	// - It must start with an uppercase or lowercase letter or a Chinese character. It cannot contain spaces or the following special characters: `@/:=”<>{[]}`.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The Tair instance series, which determines the storage medium. Valid values:
	//
	// - **tair_rdb**: memory-enhanced
	//
	// - **tair_scm**: persistent memory-optimized
	//
	// - **tair_essd**: disk-based
	//
	// This parameter is required.
	//
	// example:
	//
	// tair_scm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The end time of the maintenance window. Specify the time in the *HH:mm*Z format (UTC). For example, to end the maintenance at 02:00 (UTC+8), set this parameter to `18:00Z`.
	//
	// > The maintenance window must be at least one hour long.
	//
	// > If this parameter is not specified, the maintenance window ends at 22:00 UTC (06:00 UTC+8) by default.
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. Specify the time in the *HH:mm*Z format (UTC). For example, to start the maintenance at 01:00 (UTC+8), set this parameter to `17:00Z`.
	//
	// > If this parameter is not specified, the maintenance window starts at 18:00 UTC (02:00 UTC+8) by default.
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template. The instance is created by using the parameters defined in this template.
	//
	// example:
	//
	// g-50npzjcqb1ua6q6j****
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The password of the instance. The password must meet the following requirements:
	//
	// - It must be 8 to 32 characters in length.
	//
	// - It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The supported special characters are `!@#$%^&*()_+-=`.
	//
	// example:
	//
	// Pass!123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription duration, in months. Valid values: **1**, **2**, **3**, **4**, **5**, **6**, 7, 8, 9, 12, 24, 36, and 60.
	//
	// > This parameter is required only when you set the `ChargeType` parameter to `PrePaid`.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The service port of the instance. Valid values: 1 to 65535. Default value: 6379.
	//
	// example:
	//
	// 6379
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the instance.
	//
	// > The IP address must be within the CIDR block of the vSwitch to which the instance belongs. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query the CIDR block information.
	//
	// example:
	//
	// 172.16.88.***
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The number of read-only nodes in the primary zone. This parameter is applicable only to cloud-native instances that use the read/write splitting architecture.
	//
	// - If the instance uses the standard architecture, the valid values are 1 to 9.
	//
	// - If the instance uses the cluster architecture, specify the number of read-only nodes per shard. The valid values are 1 to 4.
	//
	// > If you create a multi-zone instance, you can use this parameter and the `SlaveReadOnlyCount` parameter to customize the number of read-only nodes in the primary and secondary zones.
	//
	// >
	//
	// > - If the instance uses the standard architecture, the sum of `ReadOnlyCount` and `SlaveReadOnlyCount` cannot exceed 9.
	//
	// >
	//
	// > - If the instance uses the cluster architecture, the sum of `ReadOnlyCount` and `SlaveReadOnlyCount` cannot exceed 4.
	//
	// example:
	//
	// 5
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// When creating an instance from a backup set, specifies whether to restore configurations such as account information (`account`), kernel parameters (`config`), and whitelists (`whitelist`) from the source backup set. To restore a specific configuration, specify its keyword. To restore multiple configurations, separate the keywords with commas.
	//
	// If this parameter is not specified, no configurations are restored from the source backup set.
	//
	// > This parameter applies only to cloud-native instances, and the source backup set must contain the specified configuration information. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation and check the `RecoverConfigMode` parameter in the response to check if the backup set contains the information.
	//
	// example:
	//
	// whitelist,config,account
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
	// The ID of the region where you want to create the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of replica nodes in the primary zone. This parameter is applicable only to cloud-native, multi-replica cluster instances. You can use this parameter to customize the number of replica nodes. Valid values: 1 to 4.
	//
	// > If you create a multi-zone instance, you can use this parameter and the `SlaveReplicaCount` parameter to customize the number of replica nodes in the primary and secondary zones. The sum of `ReplicaCount` and `SlaveReplicaCount` cannot exceed 4.
	//
	// example:
	//
	// 2
	ReplicaCount *int32 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// > - You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation or use the Resource Management console to query the IDs of resource groups. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// >
	//
	// > - Before you change the resource group of an instance, you can call the [ListResources](https://help.aliyun.com/document_detail/158866.html) operation to view the current resource group of the instance.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// If point-in-time recovery (PITR) is enabled for the source instance, you can specify a point in time within the backup retention period. The system creates a new instance by using the backup data of the source instance at that point in time. Specify the time in the *yyyy-MM-dd*T*HH:mm:ss*Z format (UTC).
	//
	// example:
	//
	// 2021-07-06T07:25:57Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The ID of the secondary zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query available zones.
	//
	// > The value of this parameter cannot be the same as the value of the `ZoneId` parameter. You cannot specify a multi-zone ID.
	//
	// example:
	//
	// cn-hangzhou-g
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards in the instance. Valid values:
	//
	// - **1*	- (default): Creates a standard architecture instance with a single shard.
	//
	// - From **2*	- to **32**: Creates a cluster architecture instance with the specified number of shards.
	//
	// > You can specify a value from **2*	- to **32*	- for this parameter only when you set the **InstanceType*	- parameter to `tair_rdb` or `tair_scm`. Only memory-enhanced and persistent memory-optimized instances support the cluster architecture.
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The architecture type of the instance. Valid values:
	//
	// - **MASTER_SLAVE*	- (default): The primary/replica architecture, which provides high availability.
	//
	// - **STAND_ALONE**: single-replica. This architecture uses a single node. If the node fails, data is lost, and the system automatically creates a new, empty instance. This architecture is supported only for **single-zone*	- deployments and does not support cluster or read/write splitting architectures.
	//
	// example:
	//
	// MASTER_SLAVE
	ShardType *string `json:"ShardType,omitempty" xml:"ShardType,omitempty"`
	// The number of read-only nodes in the secondary zone.
	//
	// example:
	//
	// 1
	SlaveReadOnlyCount *int32 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// The number of replica nodes in the secondary zone.
	//
	// example:
	//
	// 2
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// To create an instance from a backup set of an existing instance, specify the ID of the source instance.
	//
	// > You must also specify the backup data by using one of the following parameters: **BackupId**, **ClusterBackupId**, or **RestoreTime**. We recommend that you use `ClusterBackupId` for cloud-native instances that use a cluster architecture.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The storage space of the disk-based instance. The valid values of this parameter vary based on the instance type. For more information, see [Disk-based instance types](https://help.aliyun.com/document_detail/2527111.html).
	//
	// > This parameter is required only when you set the **InstanceType*	- parameter to `tair_essd` to create a Tair instance that uses an ESSD. For Tair instances that use standard `SSD`s, the storage capacity is determined by the instance type and you do not need to specify this parameter.
	//
	// example:
	//
	// 60
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The storage type. Valid values: **essd_pl1**, **essd_pl2**, and **essd_pl3**.
	//
	// > This parameter is required only when you set the **InstanceType*	- parameter to `tair_essd` to create a Tair instance that uses an Enhanced SSD (ESSD).
	//
	// example:
	//
	// essd_pl1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags of the instance.
	Tag []*CreateTairInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch in the specified VPC. You can call the VPC API operation [DescribeVSwitches](https://help.aliyun.com/document_detail/35739.html) to obtain the vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the Virtual Private Cloud (VPC) where you want to create the instance. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query available VPCs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the primary zone where you want to create the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query available zones.
	//
	// > You can also specify a secondary zone by using the `SecondaryZoneId` parameter. This deploys the primary and replica nodes in different zones within the same region for a high-availability primary/replica architecture. For example, you can set `ZoneId` to `cn-hangzhou-h` and `SecondaryZoneId` to `cn-hangzhou-g`.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTairInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateTairInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateTairInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateTairInstanceRequest) GetAutoRenewPeriod() *string {
	return s.AutoRenewPeriod
}

func (s *CreateTairInstanceRequest) GetAutoUseCoupon() *string {
	return s.AutoUseCoupon
}

func (s *CreateTairInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateTairInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateTairInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTairInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTairInstanceRequest) GetClusterBackupId() *string {
	return s.ClusterBackupId
}

func (s *CreateTairInstanceRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CreateTairInstanceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *CreateTairInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTairInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateTairInstanceRequest) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *CreateTairInstanceRequest) GetGlobalSecurityGroupIds() *string {
	return s.GlobalSecurityGroupIds
}

func (s *CreateTairInstanceRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateTairInstanceRequest) GetInstanceEndpointType() *string {
	return s.InstanceEndpointType
}

func (s *CreateTairInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateTairInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateTairInstanceRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *CreateTairInstanceRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *CreateTairInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTairInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTairInstanceRequest) GetParamGroupId() *string {
	return s.ParamGroupId
}

func (s *CreateTairInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateTairInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateTairInstanceRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreateTairInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateTairInstanceRequest) GetReadOnlyCount() *int32 {
	return s.ReadOnlyCount
}

func (s *CreateTairInstanceRequest) GetRecoverConfigMode() *string {
	return s.RecoverConfigMode
}

func (s *CreateTairInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairInstanceRequest) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *CreateTairInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTairInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTairInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTairInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateTairInstanceRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *CreateTairInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateTairInstanceRequest) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *CreateTairInstanceRequest) GetShardType() *string {
	return s.ShardType
}

func (s *CreateTairInstanceRequest) GetSlaveReadOnlyCount() *int32 {
	return s.SlaveReadOnlyCount
}

func (s *CreateTairInstanceRequest) GetSlaveReplicaCount() *int32 {
	return s.SlaveReplicaCount
}

func (s *CreateTairInstanceRequest) GetSrcDBInstanceId() *string {
	return s.SrcDBInstanceId
}

func (s *CreateTairInstanceRequest) GetStorage() *int32 {
	return s.Storage
}

func (s *CreateTairInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateTairInstanceRequest) GetTag() []*CreateTairInstanceRequestTag {
	return s.Tag
}

func (s *CreateTairInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTairInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTairInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairInstanceRequest) SetAutoPay(v bool) *CreateTairInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateTairInstanceRequest) SetAutoRenew(v string) *CreateTairInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateTairInstanceRequest) SetAutoRenewPeriod(v string) *CreateTairInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateTairInstanceRequest) SetAutoUseCoupon(v string) *CreateTairInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateTairInstanceRequest) SetBackupId(v string) *CreateTairInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetBusinessInfo(v string) *CreateTairInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateTairInstanceRequest) SetChargeType(v string) *CreateTairInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTairInstanceRequest) SetClientToken(v string) *CreateTairInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTairInstanceRequest) SetClusterBackupId(v string) *CreateTairInstanceRequest {
	s.ClusterBackupId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetConnectionStringPrefix(v string) *CreateTairInstanceRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateTairInstanceRequest) SetCouponNo(v string) *CreateTairInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *CreateTairInstanceRequest) SetDryRun(v bool) *CreateTairInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTairInstanceRequest) SetEngineVersion(v string) *CreateTairInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateTairInstanceRequest) SetGlobalInstanceId(v string) *CreateTairInstanceRequest {
	s.GlobalInstanceId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetGlobalSecurityGroupIds(v string) *CreateTairInstanceRequest {
	s.GlobalSecurityGroupIds = &v
	return s
}

func (s *CreateTairInstanceRequest) SetInstanceClass(v string) *CreateTairInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateTairInstanceRequest) SetInstanceEndpointType(v string) *CreateTairInstanceRequest {
	s.InstanceEndpointType = &v
	return s
}

func (s *CreateTairInstanceRequest) SetInstanceName(v string) *CreateTairInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateTairInstanceRequest) SetInstanceType(v string) *CreateTairInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateTairInstanceRequest) SetMaintainEndTime(v string) *CreateTairInstanceRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *CreateTairInstanceRequest) SetMaintainStartTime(v string) *CreateTairInstanceRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *CreateTairInstanceRequest) SetOwnerAccount(v string) *CreateTairInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetOwnerId(v int64) *CreateTairInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetParamGroupId(v string) *CreateTairInstanceRequest {
	s.ParamGroupId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetPassword(v string) *CreateTairInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateTairInstanceRequest) SetPeriod(v int32) *CreateTairInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateTairInstanceRequest) SetPort(v int32) *CreateTairInstanceRequest {
	s.Port = &v
	return s
}

func (s *CreateTairInstanceRequest) SetPrivateIpAddress(v string) *CreateTairInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateTairInstanceRequest) SetReadOnlyCount(v int32) *CreateTairInstanceRequest {
	s.ReadOnlyCount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetRecoverConfigMode(v string) *CreateTairInstanceRequest {
	s.RecoverConfigMode = &v
	return s
}

func (s *CreateTairInstanceRequest) SetRegionId(v string) *CreateTairInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetReplicaCount(v int32) *CreateTairInstanceRequest {
	s.ReplicaCount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetResourceGroupId(v string) *CreateTairInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetResourceOwnerAccount(v string) *CreateTairInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetResourceOwnerId(v int64) *CreateTairInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetRestoreTime(v string) *CreateTairInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateTairInstanceRequest) SetSecondaryZoneId(v string) *CreateTairInstanceRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetSecurityToken(v string) *CreateTairInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTairInstanceRequest) SetShardCount(v int32) *CreateTairInstanceRequest {
	s.ShardCount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetShardType(v string) *CreateTairInstanceRequest {
	s.ShardType = &v
	return s
}

func (s *CreateTairInstanceRequest) SetSlaveReadOnlyCount(v int32) *CreateTairInstanceRequest {
	s.SlaveReadOnlyCount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetSlaveReplicaCount(v int32) *CreateTairInstanceRequest {
	s.SlaveReplicaCount = &v
	return s
}

func (s *CreateTairInstanceRequest) SetSrcDBInstanceId(v string) *CreateTairInstanceRequest {
	s.SrcDBInstanceId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetStorage(v int32) *CreateTairInstanceRequest {
	s.Storage = &v
	return s
}

func (s *CreateTairInstanceRequest) SetStorageType(v string) *CreateTairInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateTairInstanceRequest) SetTag(v []*CreateTairInstanceRequestTag) *CreateTairInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateTairInstanceRequest) SetVSwitchId(v string) *CreateTairInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetVpcId(v string) *CreateTairInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateTairInstanceRequest) SetZoneId(v string) *CreateTairInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateTairInstanceRequest) Validate() error {
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

type CreateTairInstanceRequestTag struct {
	// The key of the tag.
	//
	// > A single request can contain up to five key-value pairs.
	//
	// example:
	//
	// key1_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// > **N*	- specifies the Nth tag in the request. For example, **Tag.1.Value*	- specifies the value of the first tag, and **Tag.2.Value*	- specifies the value of the second tag.
	//
	// example:
	//
	// value1_test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTairInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTairInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTairInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTairInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTairInstanceRequestTag) SetKey(v string) *CreateTairInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTairInstanceRequestTag) SetValue(v string) *CreateTairInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTairInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
