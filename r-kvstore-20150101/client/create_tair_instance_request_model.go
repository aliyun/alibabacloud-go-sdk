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
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false**: disables auto-renewal. This is the default value.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal epoch. Unit: months. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required when the **AutoRenew*	- parameter is set to **true**.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *string `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// Specifies whether to use a voucher. Valid values:
	//
	// 	- **true**: uses a voucher.
	//
	// 	- **false**: does not use a voucher. This is the default value.
	//
	// example:
	//
	// true
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The ID of the backup set of the source instance. The system uses the data in the specified backup set to create the instance. You can invoke the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query the backup set ID. If the source instance is a cluster instance, you must enter the backup set IDs of all shards in the source instance and separate them with commas, such as "10**,11**,15**".
	//
	// > If your instance is a cloud-native architecture instance, we recommend that you invoke the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation to query the cluster backup set ID, such as "cb-xx", and then specify the ClusterBackupId request parameter to clone the cluster instance. This saves you the trouble of entering the backup set ID of each shard.
	//
	// example:
	//
	// 2158****20
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The activity ID and business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method. Valid values:
	//
	// - **PrePaid**: subscription. This is the default value.
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token is case-sensitive and can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cluster backup set. Some new cluster architectures support cluster backup set IDs. You can call the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation to obtain the ID.
	//
	// 	- If the cluster architecture supports cluster backup sets, you can enter the specified cluster backup set ID. You do not need to specify the **BackupId*	- parameter.
	//
	// 	- If the cluster architecture does not support cluster backup sets, enter the backup set IDs of all shards in the source instance in the BackupId parameter and separate them with commas, such as "2158\\*\\*\\*\\*20,2158\\*\\*\\*\\*22".
	//
	// example:
	//
	// cb-hyxdof5x9kqb****
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The prefix of the endpoint. The prefix must be 8 to 40 characters in length and can contain lowercase letters and digits. It must start with a lowercase letter.
	//
	// >
	//
	// > The endpoint is in the following format: <Prefix-Instance ID>.redis.rds.aliyuncs.com.
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
	// Specifies whether to perform a dry run before you create the instance. Valid values:
	//
	// 	- **true**: performs a dry run without creating the instance. The system checks the request parameters, request format, service limits, and available resources. If the request fails the dry run, the system returns an error. If the request passes the dry run, the system returns the `DryRunOperation` error code.
	//
	// 	- **false**: sends the request. If the request passes the check, the instance is created. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The database engine version. Default value: **1.0**. The value varies based on the Tair instance type:
	//
	// - Tair_rdb: Tair memory-enhanced instances are compatible with Redis 5.0, Redis 6.0, and Redis 7.0 protocols. Set the value to **5.0**, **6.0**, or **7.0**.
	//
	// - Tair_scm: Tair persistent memory-optimized instances are compatible with the Redis 6.0 protocol. Set the value to **1.0**.
	//
	// - Tair_essd: Tair disk-based (ESSD/SSD) instances are compatible with the Redis 6.0 protocol. Set the value to **1.0*	- for ESSD-based instances and **2.0*	- for SSD-based instances.
	//
	// example:
	//
	// 1.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to use the new instance as a child instance in a distributed instance. You can use this parameter to create a distributed instance.
	//
	// 	- To use the new instance as the first child instance, set the value to **true**.
	//
	// 	- To use the new instance as the second or third child instance, set the value to the ID of the distributed instance, such as gr-bp14rkqrhac****.
	//
	// 	- If you do not want to use the new instance as a child instance in a distributed instance, leave this parameter empty.
	//
	// > The new instance must be a Tair memory-enhanced instance if you want to use it as a child instance in a distributed instance.
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The global IP address whitelist templates of the instance. Separate multiple whitelist templates with commas (,). Each template must be unique.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The instance type. For more information, see:
	//
	// - [Memory-enhanced instance types](https://help.aliyun.com/document_detail/2527112.html).
	//
	// - [Persistent memory-optimized instance types](https://help.aliyun.com/document_detail/2527110.html).
	//
	// - [Disk-based instance types](https://help.aliyun.com/document_detail/2527111.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// tair.scm.standard.4m.32d
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The endpoint type used when you create a cloud-native Dual-active-zone deployment read/write splitting instance. If you do not specify this parameter, the default value is AzIndependentEndpoint.
	//
	// - **AzIndependentEndpoint**: active zone-independent endpoint. This is the **default value**. Independent endpoints are provided for the primary and secondary active zones. You can use different endpoints to achieve nearest access.
	//
	// - **UnifiedEndpoint**: unified endpoint. A unified endpoint is provided to access nodes in both the primary and secondary active zones. However, cross-active-zone access may occur.
	//
	// 	Notice: This parameter applies only to cloud-native Dual-active-zone deployment read/write splitting instances. For other types of instances, only active zone-independent endpoints are supported. Even if you specify UnifiedEndpoint, it does not take effect.
	//
	// 	Notice: The UnifiedEndpoint parameter is currently available only to whitelisted users. If you are not a whitelisted user and specify this parameter, an invoke error is returned. To request access, commit a ticket.
	//
	// example:
	//
	// AzIndependentEndpoint
	InstanceEndpointType *string `json:"InstanceEndpointType,omitempty" xml:"InstanceEndpointType,omitempty"`
	// The instance name. The name must meet the following requirements:
	//
	// 	- The name must be 2 to 80 characters in length.
	//
	// 	- The name must start with a letter or a Chinese character. It cannot contain spaces or the following special characters: `@/:=”<>{[]}`.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The storage medium. Valid values:
	//
	// 	- **tair_rdb**: memory-enhanced.
	//
	// 	- **tair_scm**: persistent memory-optimized.
	//
	// 	- **tair_essd**: disk-based.
	//
	// This parameter is required.
	//
	// example:
	//
	// tair_scm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The end time of the maintenance window. Specify the time in the <i>HH:mm</i>Z format. The time must be in UTC. For example, if you want the maintenance window to end at 02:00 (UTC+8), set the value to `18:00Z`.
	//
	// > The interval between the start time and the end time must be at least 1 hour.
	//
	// > If you do not specify this parameter, the default value is 22:00 (UTC), which is 06:00 (UTC+8).
	//
	// example:
	//
	// 18:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. Specify the time in the <i>HH:mm</i>Z format. The time must be in UTC. For example, if you want the maintenance window to start at 01:00 (UTC+8), set the value to `17:00Z`.
	//
	// > If you do not specify this parameter, the default value is 18:00 (UTC), which is 02:00 (UTC+8).
	//
	// example:
	//
	// 17:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template. The instance is created based on the parameters in the specified template. Each template ID must be unique.
	//
	// example:
	//
	// g-50npzjcqb1ua6q6j****
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The instance password. The password must meet the following requirements:
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, special characters, and digits. Supported special characters are `!@#$%^&*()_+-=`.
	//
	// example:
	//
	// Pass!123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription duration. Unit: months. Valid values: **1*	- to **9**, **12**, **24**, **36**, and **60**.
	//
	// > This parameter is required when the **ChargeType*	- parameter is set to **PrePaid**.
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
	// The internal IP address of the instance.
	//
	// > The IP address must fall within the CIDR block of the vSwitch to which the instance belongs. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation of the VPC API to query the CIDR block information.
	//
	// example:
	//
	// 172.16.88.***
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The number of read-only nodes in the primary zone. This parameter applies only when you create a cloud-native read/write splitting instance.
	//
	// 	- If the instance uses the standard architecture, valid values are 1 to 9.
	//
	// 	- If the instance uses the cluster architecture, valid values are 1 to 4, which indicate the number of read-only nodes per data shard.
	//
	// > If you create a multi-zone instance, you can use this parameter and the SlaveReadOnlyCount parameter to specify the number of read-only nodes in the primary and secondary zones.
	//
	// > 	- If the instance uses the standard architecture, the sum of the values of this parameter and the SlaveReadOnlyCount parameter cannot exceed 9.
	//
	// > 	- If the instance uses the cluster architecture, the sum of the values of this parameter and the SlaveReadOnlyCount parameter cannot exceed 4.
	//
	// example:
	//
	// 5
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// Specifies whether to restore the account, kernel parameters, and whitelist from the original backup set when you create an instance by using a specified backup set. For example, to restore the account information, set the value to account.
	//
	// Default value: empty. An empty value indicates that the account, kernel parameters, and whitelist are not restored from the original backup set.
	//
	// > This parameter applies only to cloud-native instances and requires that the original backup set contains the account, kernel parameters, and whitelist information. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query whether the RecoverConfigMode parameter of a specified backup set contains the preceding information.
	//
	// example:
	//
	// whitelist,config,account
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list. Use this parameter to specify the region where you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of secondary nodes in the primary zone. This parameter applies only when you create a cloud-native cluster instance with multiple replicas. You can use this parameter to specify the number of secondary nodes. Valid values: 1 to 4.
	//
	// > If you create a multi-zone instance, you can use this parameter and the SlaveReplicaCount parameter to specify the number of secondary nodes in the primary and secondary zones. The sum of the values of this parameter and the SlaveReplicaCount parameter cannot exceed 4.
	//
	// example:
	//
	// 2
	ReplicaCount *int32 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	// The ID of the destination resource group.
	//
	// > 	- You can invoke the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation or use the console to obtain the resource group ID list. For related operations, see [View the basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// > 	- Before you change the resource group to which the instance belongs, you can invoke the [ListResources](https://help.aliyun.com/document_detail/158866.html) operation to view the current resource group of the instance.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore the source instance. If the data flashback feature is enabled for the source instance, you can specify a point in time within the backup retention period of the source instance. The system uses the backup data of the source instance at the specified point in time to create the instance. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// example:
	//
	// 2021-07-06T07:25:57Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The secondary zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent zone list.
	//
	// > The value of this parameter cannot be the same as the value of the ZoneId parameter, and you cannot specify the ID of a multi-zone deployment.
	//
	// example:
	//
	// cn-hangzhou-g
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of data nodes in the instance. Valid values:
	//
	// 	- **1**: creates a [standard instance](https://help.aliyun.com/document_detail/52228.html) that has only one data node. This is the default value.
	//
	// 	- **2*	- to **32**: creates a [cluster instance](https://help.aliyun.com/document_detail/52228.html) that has the specified number of data nodes.
	//
	// > You can set this parameter to a value from **2*	- to **32*	- only when the **InstanceType*	- parameter is set to **tair_rdb*	- or **tair_scm**. This means that only memory-enhanced and persistent memory-optimized instances support the cluster architecture.
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The instance edition. Valid values:
	//
	// - **MASTER_SLAVE**: High-availability Edition. This edition uses a primary/secondary architecture to ensure high availability. This is the default value.
	//
	// - **STAND_ALONE**: Standalone Edition. This edition uses a single-node architecture. If a node fails, data is lost and the system automatically creates a new empty instance. This edition is supported only in **single-zone*	- deployments and does not support cluster or read/write splitting architectures.
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
	// The number of secondary nodes in the secondary zone.
	//
	// example:
	//
	// 2
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// Instance ID of the source instance. If you want to create an instance based on the backup set of an existing instance, specify instance ID of the source instance in this parameter.
	//
	// > Then, specify the backup set or point in time by using the **BackupId**, **ClusterBackupId*	- (recommended for cloud-native cluster architectures), or **RestoreTime*	- parameter. This parameter must be used in combination with one of the preceding three parameters.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The storage capacity of the disk-based instance. The valid values vary based on the instance type. For more information, see [Disk-based instance types](https://help.aliyun.com/document_detail/2527111.html).
	//
	// > This parameter is required only when the **InstanceType*	- parameter is set to **tair_essd*	- and you create an ESSD-based instance. If you create a Tair disk-based **SSD*	- instance, the capacity field is defined as a fixed value based on the instance type, and you do not need to specify this field.
	//
	// example:
	//
	// 60
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The storage type. Valid values: **essd_pl1**, **essd_pl2**, and **essd_pl3**.
	//
	// > This parameter is required only when the **InstanceType*	- parameter is set to **tair_essd*	- and you create an ESSD-based instance.
	//
	// example:
	//
	// essd_pl1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tag information of the instance.
	Tag []*CreateTairInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch in the VPC. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation of the VPC API to obtain the vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC). You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation of the VPC API to query the VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The primary active zone ID. You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent active zone list. Use this parameter to specify the active zone where you want to create the instance.
	//
	// > You can also specify the SecondaryZoneId parameter to set the secondary active zone. The primary and secondary nodes will be placed in the specified primary and secondary active zones to implement a dual-active-zone primary/secondary architecture deployment. For example, set the ZoneId parameter to "cn-hangzhou-h" and the SecondaryZoneId parameter to "cn-hangzhou-g".
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
	// The tag key. A tag key and a tag value form a key-value pair.
	//
	// > You can specify up to 5 tag key-value pairs at a time.
	//
	// example:
	//
	// key1_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// > **N*	- indicates the sequence number of the tag. For example, **Tag.1.Value*	- indicates the value of the first tag, and **Tag.2.Value*	- indicates the value of the second tag.
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
