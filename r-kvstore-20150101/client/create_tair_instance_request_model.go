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
	SetInstanceName(v string) *CreateTairInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateTairInstanceRequest
	GetInstanceType() *string
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
	// true
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// You can set the BackupId parameter to the backup set ID of the source instance. The system uses the data stored in the backup set to create an instance. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query the backup set ID. If the source instance is a cluster instance, set the BackupId parameter to the backup set IDs of all shards of the source instance, separated by commas (,). Example: "10\\*\\*,11\\*\\*,15\\*\\*".
	//
	// >  If your instance is a cloud-native cluster instance, we recommend that you use [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) to query the backup set ID of the cluster instance, such as cb-xx. Then, set the ClusterBackupId request parameter to the backup set ID to clone the cluster instance. This eliminates the need to specify the backup set ID of each shard.
	//
	// example:
	//
	// 11111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the promotional event or the business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid*	- (default): subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests and is case-sensitive. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is supported for specific new cluster instances. You can query the backup set ID by calling the [DescribeClusterBackupList](https://help.aliyun.com/document_detail/2679168.html) operation.
	//
	// 	- If this parameter is supported, you can specify the backup set ID. In this case, you do not need to specify the **BackupId*	- parameter.
	//
	// 	- If this parameter is not supported, set the BackupId parameter to the IDs of backup sets in all shards of the source instance, separated by commas (,). Example: "2158\\*\\*\\*\\*20,2158\\*\\*\\*\\*22".
	//
	// example:
	//
	// cb-hyxdof5x9kqb****
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The prefix of the endpoint. The prefix must be 8 to 40 characters in length and can contain lowercase letters and digits. It must start with a lowercase letter.
	//
	// >  The endpoint must be in the \\<prefix>.redis.rds.aliyuncs.com format.
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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run and does not create the instance. The system prechecks the request parameters, request format, service limits, and available resources. If the request fails the dry run, an error code is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (false): performs a dry run and performs the actual request. If the request passes the dry run, the instance is directly created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The database engine version. Default value: **1.0**. The parameter value varies based on the Tair instance series.
	//
	// 	- To create a Tair DRAM-based instance (Tair_rdb) that is compatible with Redis 5.0, 6.0, or 7.0, set this parameter to **5.0**, **6.0**, or **7.0**.
	//
	// 	- To create a Tair persistent memory-optimized instance (tair_scm) that is compatible with Redis 6.0, set this parameter to **1.0**.
	//
	// 	- To create a Tair ESSD-based instance (tair_essd) that is compatible with Redis 6.0, set this parameter to **1.0**. To create a Tair SSD-based instance that is compatible with Redis 6.0, set this parameter to **2.0**.
	//
	// example:
	//
	// 1.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to use the created instance as a child instance of a distributed instance.
	//
	// 	- If you want the created instance to be used as the first child instance, enter **true**.
	//
	// 	- If you want the created instance to be used as the second or third child instance, enter the ID of the distributed instance, such as gr-bp14rkqrhac\\*\\*\\*\\*.
	//
	// 	- If you do not want the created instance to be used as a distributed instance, leave the parameter empty.
	//
	// >  If you want the created instance to be used as a distributed instance, the created instance must be a Tair DRAM-based instance.
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The global IP whitelist templates of the instance. Separate multiple IP whitelist templates with commas (,). Each IP whitelist template must be unique.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The instance series. For more information, see the following topics:
	//
	// 	- [DRAM-based instances](https://help.aliyun.com/document_detail/2527112.html)
	//
	// 	- [Persistent memory-optimized instances](https://help.aliyun.com/document_detail/2527110.html)
	//
	// 	- [ESSD/SSD-based instances](https://help.aliyun.com/document_detail/2527111.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// tair.scm.standard.4m.32d
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The name of the instance. The name must meet the following requirements:
	//
	// 	- The name must be 2 to 80 characters in length.
	//
	// 	- The name must start with a letter and cannot contain spaces or special characters. Special characters include `@ / : = " < > { [ ] }`
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance series. Valid values:
	//
	// 	- **tair_rdb**: Tair DRAM-based instance
	//
	// 	- **tair_scm**: Tair persistent memory-optimized instance
	//
	// 	- **tair_essd**: Tair ESSD/SSD-based instance
	//
	// This parameter is required.
	//
	// example:
	//
	// tair_scm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template. The instance is created based on the parameters in the parameter template. The ID must be unique.
	//
	// example:
	//
	// g-50npzjcqb1ua6q6j****
	ParamGroupId *string `json:"ParamGroupId,omitempty" xml:"ParamGroupId,omitempty"`
	// The password that is used to connect to the instance. The password must meet the following requirements:
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// 	- The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ % ^ & 	- ( ) _ + - =`
	//
	// example:
	//
	// Pass!123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The subscription duration. Valid values: **1**, 2, 3, 4, 5, 6, 7, 8, **9**, **12**, **24**,**36**, and **60**. Unit: month.
	//
	// >  This parameter is required only if the **ChargeType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The service port number of the instance. Valid values: 1024 to 65535. Default value: 6379.
	//
	// example:
	//
	// 6379
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The internal IP address of the instance.
	//
	// >  The IP address must be within the CIDR block of the vSwitch to which you want the instance to connect. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation of VPC to query the CIDR block information.
	//
	// example:
	//
	// 172.16.88.***
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The number of read replicas in the primary zone. This parameter applies only to cloud-native read/write splitting instances. Valid values: 1 to 9.
	//
	// >  The sum of the values of this parameter and the SlaveReadOnlyCount parameter cannot exceed 9.
	//
	// example:
	//
	// 5
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// Specifies whether to restore the account, kernel parameter, and whitelist information from the original backup set when you create an instance from the specified backup set. For example, if you want to restore the account information, set the parameter to `{"account":true}`.
	//
	// This parameter is empty by default, which indicates that the account, kernel parameter, and whitelist information is not restored from the original backup set.
	//
	// >  This parameter applies only to cloud-native cluster instances. The account, kernel parameter, and whitelist information must be stored in the original backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to check whether the RecoverConfigMode configurations in the specified backup set contain the preceding information.
	//
	// example:
	//
	// {"whitelist":true,"config":true,"account":true}
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of replica nodes in the primary zone. This parameter applies only to cloud-native multi-replica cluster instances. Valid values: 1 to 4.
	//
	// >
	//
	// 	- The sum of the values of this parameter and the SlaveReplicaCount parameter cannot exceed 4.
	//
	// 	- You can specify only one of the ReplicaCount and ReadOnlyCount parameters.
	//
	// 	- Master-replica instances do not support multiple replicas.
	//
	// example:
	//
	// 2
	ReplicaCount *int32 `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	// The ID of the resource group that you want to manage.
	//
	// >
	//
	// 	- You can query resource group IDs in the console or by calling the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation. For more information, see [View the basic information about a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// 	- Before you modify the resource group to which an instance belongs, you can call the [ListResources](https://help.aliyun.com/document_detail/158866.html) operation to view the current resource group of the instance.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// If data flashback is enabled for the source instance, you can use this parameter to specify a point in time within the backup retention period of the source instance. The system uses the backup data of the source instance at the point in time to create an instance. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2021-07-06T07:25:57Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The ID of the secondary zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the ID of the secondary zone.
	//
	// >  You cannot specify multiple zone IDs or set this parameter to a value that is the same as that of the ZoneId parameter.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of data nodes in the instance. Valid values:
	//
	// 	- **1*	- (default): You can create a [standard instance](https://help.aliyun.com/document_detail/52228.html) that contains only one data node.
	//
	// 	- **2*	- to **32**: You can create a [cluster instance](https://help.aliyun.com/document_detail/52228.html) that contains the specified number of data nodes.
	//
	// >  When the **InstanceType*	- parameter is set to **tair_rdb*	- or **tair_scm**, this parameter can be set to a value in the range of **2*	- to **32**. Only DRAM-based and persistent memory-optimized instances support the cluster architecture.
	//
	// example:
	//
	// 1
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The shard type of the instance. Valid values:
	//
	// 	- **MASTER_SLAVE*	- (default): runs in a master-replica architecture that provides high availability.
	//
	// 	- **STAND_ALONE**: runs in a standalone architecture. If the only node fails, the system creates a new instance and switches the workloads to the new instance. This may cause data loss. You can set the ShardType parameter to this value only if the instance uses the **single-zone*	- deployment mode. If you set the ShardType parameter to this value, you cannot create cluster or read/write splitting instances.
	//
	// example:
	//
	// MASTER_SLAVE
	ShardType *string `json:"ShardType,omitempty" xml:"ShardType,omitempty"`
	// The number of read replicas in the secondary zone when you create a multi-zone read/write splitting instance. The sum of the values of this parameter and the ReadOnlyCount parameter cannot exceed 9.
	//
	// > When you create a multi-zone read/write splitting instance, you must specify both SlaveReadOnlyCount and SecondaryZoneId.
	//
	// example:
	//
	// 1
	SlaveReadOnlyCount *int32 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// The number of replica nodes in the secondary zone when you create a cloud-native multi-replica cluster instance deployed across multiple zones. The sum of the values of this parameter and the ReplicaCount parameter cannot exceed 4.
	//
	// >  When you create a cloud-native multi-replica cluster instance deployed across multiple zones, you must specify both SlaveReplicaCount and SecondaryZoneId.
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
	// The storage capacity of the ESSD/SSD-based instance. The valid values vary based on the instance type. For more information, see [ESSD/SSD-based instances](https://help.aliyun.com/document_detail/2527111.html).
	//
	// >  This parameter is required only when you set the **InstanceType*	- parameter to **tair_essd*	- to create an ESSD-based instance. If you create a Tair **SSD**-based instance, the Storage parameter is automatically specified based on predefined specifications. You do not need to specify this parameter.
	//
	// example:
	//
	// 60
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The storage type. Valid values: **essd_pl1**, **essd_pl2**, and **essd_pl3**.
	//
	// >  This parameter is required only when you set the **InstanceType*	- parameter to **tair_essd*	- to create an ESSD-based instance.
	//
	// Enumerated values:
	//
	// 	- essd_pl0
	//
	// 	- essd_pl1
	//
	// 	- essd_pl2
	//
	// 	- essd_pl3
	//
	// example:
	//
	// essd_pl1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Details of the tags.
	Tag []*CreateTairInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch that belongs to the VPC. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query vSwitch IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query VPC IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the primary zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent zone list.
	//
	// >  You can also set the SecondaryZoneId parameter to specify the secondary zone. The primary and secondary nodes will then be deployed in the specified primary and secondary zones to implement the master-replica zone-disaster recovery architecture. For example, you can set the ZoneId parameter to cn-hangzhou-h and the SecondaryZoneId parameter to cn-hangzhou-g.
	//
	// example:
	//
	// cn-hangzhou-e
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

func (s *CreateTairInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateTairInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
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

func (s *CreateTairInstanceRequest) SetInstanceName(v string) *CreateTairInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateTairInstanceRequest) SetInstanceType(v string) *CreateTairInstanceRequest {
	s.InstanceType = &v
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
	return dara.Validate(s)
}

type CreateTairInstanceRequestTag struct {
	// The tag key. A tag is a key-value pair.
	//
	// >  A maximum of five key-value pairs can be specified at a time.
	//
	// example:
	//
	// key1_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// >  **N*	- specifies the value of the nth tag. For example, **Tag.1.Value*	- specifies the value of the first tag, and **Tag.2.Value*	- specifies the value of the second tag.
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
