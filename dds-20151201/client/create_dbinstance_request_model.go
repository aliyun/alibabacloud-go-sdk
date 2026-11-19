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
	SetDryRun(v bool) *CreateDBInstanceRequest
	GetDryRun() *bool
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
	SetMaintainEndTime(v string) *CreateDBInstanceRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *CreateDBInstanceRequest
	GetMaintainStartTime() *string
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
	SetTTLEnabled(v bool) *CreateDBInstanceRequest
	GetTTLEnabled() *bool
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
	// - It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Special characters include !@#$%^&*()_+-=.
	//
	// - It must be 8 to 32 characters in length.
	//
	// > For more information about how to resolve the issue where database connection fails due to special characters in the account password in the connection string, see [What do I do if the connection to an ApsaraDB for MongoDB instance fails because the username or password in the connection string contains special characters?](https://help.aliyun.com/document_detail/471568.html).
	//
	// example:
	//
	// 123456Aa
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// - **true**: enables auto-renewal.
	//
	// - **false**: disables auto-renewal. In this case, you must manually renew the instance. This is the default value.
	//
	// > This parameter is valid and optional only when the **ChargeType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/62172.html) operation to query the backup set ID.
	//
	// > This parameter is required only when you call this operation to clone an instance from a backup set. You must also specify the **SrcDBInstanceId*	- parameter.
	//
	// example:
	//
	// 32994****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The additional parameter for business information.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PostPaid**: pay-as-you-go. This is the default value.
	//
	// - **PrePaid**: subscription.
	//
	// > When this parameter is set to **PrePaid**, you must also specify the **Period*	- parameter.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and can be up to 64 ASCII characters in length. The value cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **default*	- or **null**: uses a coupon. This is the default value.
	//
	// - **youhuiquan_promotion_option_id_for_blank**: does not use a coupon.
	//
	// example:
	//
	// default
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance type. You can call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to query the available instance types.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds.mongo.standard
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The name of the instance. The name must meet the following requirements:
	//
	// - It must start with a letter or a Chinese character.
	//
	// - It can contain digits, letters, Chinese characters, underscores (_), periods (.), and hyphens (-).
	//
	// - It must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance storage. Unit: GB.
	//
	// The valid values are subject to the instance type. For details, see [Replica set instance types](https://help.aliyun.com/document_detail/311410.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The name of the database.
	//
	// > When you call this operation to clone an instance, you can specify this parameter to specify the databases to be cloned. If you do not specify this parameter, all databases of the instance are cloned.
	//
	// example:
	//
	// mongodbtest
	DatabaseNames *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: performs a dry run without calling the operation. The system checks items such as request parameters, request format, business limits, and inventory. If the check fails, the corresponding error is returned. If the check succeeds, the DryRunOperation error code is returned.
	//
	// - false (default): sends a normal request. After the check succeeds, the operation is directly executed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable cloud disk encryption.
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
	// The database engine. Set the value to **MongoDB**.
	//
	// example:
	//
	// MongoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version. Valid values:
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
	// > When you call this operation to clone clone instance or restore clone instance from the recycle bin, the value of this parameter must be the same as that of the source instance.
	//
	// 	Warning: Versions 3.4 and earlier are no longer available for sale.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The global IP address whitelist templates of the instance. Separate multiple templates with commas (,). Duplicate templates are not allowed. This feature is in canary release.
	//
	// example:
	//
	// g-qxieqf40xjst1ngpr3jz
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The zone where the hidden node resides. You can specify this parameter for multi-zone deployment. Valid values:
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
	// > - This parameter is active only when the instance uses cloud disks.
	//
	// > - The value of this parameter cannot be the same as the value of the **ZoneId*	- or **SecondaryZoneId*	- parameter.
	//
	// example:
	//
	// cn-hangzhou-i
	HiddenZoneId *string `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 22:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The network type of the instance. Valid values:
	//
	// **VPC**: Virtual Private Cloud (VPC).
	//
	// example:
	//
	// VPC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the instance. Unit: month.
	//
	// Valid values: **1*	- to **9*	- (integers), **12**, **24**, **36**, and **60**.
	//
	// > This parameter is required and valid only when the **ChargeType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The provisioned performance (IOPS). Valid values: 0 to 50000.
	//
	// example:
	//
	// 1960
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The number of **read-only nodes*	- in the replica set instance. Valid values: **0*	- to **5*	- (integers). Default value: **0**.
	//
	// example:
	//
	// 0
	ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
	//
	// > When you call this operation to clone clone instance or restore clone instance from the recycle bin, the value of this parameter must be the same as that of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of **primary and secondary nodes*	- in the replica set instance. Valid values:
	//
	// - **3*	- (default)
	//
	// - **5**
	//
	// - **7**
	//
	// 	Notice: You do not need to specify this parameter for a standalone instance.
	//
	// example:
	//
	// 3
	ReplicationFactor    *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore the instance. You can specify any point in time within the last seven days. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// > This parameter is required only when you call this operation to clone an instance to a specific point in time. You must also specify the **SrcDBInstanceId*	- parameter.
	//
	// example:
	//
	// 2022-03-13T12:11:14Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The backup restoration type of the instance. Valid values:
	//
	// - 0: restores the instance to a specified backup set
	//
	// - 1: restores the instance to a specified point in time
	//
	// - 2: restores a released instance to a specified backup set
	//
	// - 3: restores the instance to a specified geo-redundancy backup set
	//
	// example:
	//
	// 0
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The zone where the secondary node resides. You can specify this parameter for multi-zone deployment. Valid values:
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
	// > - This parameter is active only when the instance uses cloud disks.
	//
	// > - The value of this parameter cannot be the same as the value of the **ZoneId*	- or **HiddenZoneId*	- parameter.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The IP address whitelist of the instance. Separate multiple IP addresses with commas (,). Duplicate IP addresses are not allowed. The following three formats are supported:
	//
	// - 0.0.0.0/0
	//
	// - IP address format, such as 10.23.12.24.
	//
	// - CIDR format, such as 10.23.12.0/24. 24 indicates the length of the prefix in the CIDR block, and the valid range is 1 to 32.
	//
	// > - You can add up to 1,000 IP addresses or CIDR blocks to each instance. The total number of IP addresses or CIDR blocks in all IP address whitelist groups cannot exceed 1,000.
	//
	// > - 0.0.0.0/0 indicates that all IP addresses can access the databases of the instance. This is a high-risk setting. Proceed with caution.
	//
	// example:
	//
	// 192.168.xx.xx,192.168.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The source instance ID.
	//
	// > When you call this operation to clone an instance, you must specify this parameter and the **BackupId*	- or **RestoreTime*	- parameter. When you call this operation to restore a revoked instance, you only need to specify this parameter, and you do not need to specify the **BackupId*	- or **RestoreTime*	- parameter.
	//
	// example:
	//
	// dds-bp1ee12ad351****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The region where the source instance resides.
	//
	// >- This parameter is required when the backup restoration type is 2 or 3.
	//
	// example:
	//
	// 2
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The storage engine of the instance. Set the value to **WiredTiger**.
	//
	// > - When you call this operation to clone clone instance or restore clone instance from the recycle bin, the value of this parameter must be the same as that of the source instance.
	//
	// > - For more information about the constraints on the selection of storage engines and versions, see [Versions and storage engines](https://help.aliyun.com/document_detail/61906.html).
	//
	// example:
	//
	// WiredTiger
	StorageEngine *string `json:"StorageEngine,omitempty" xml:"StorageEngine,omitempty"`
	// The storage type. Valid values:
	//
	// - **cloud_essd1**: PL1 ESSD.
	//
	// - **cloud_essd2**: PL2 ESSD.
	//
	// - **cloud_essd3**: PL3 ESSD.
	//
	// - **cloud_auto**: ESSD AutoPL cloud disk.
	//
	// - **local_ssd**: local SSD.
	//
	// > - If you specify cloud_essd1 when you purchase a standalone instance, the instance actually uses a cloud_essd cloud disk.
	//
	// > - ESSD AutoPL cloud disks are supported only on the Alibaba Cloud China site.
	//
	// > - The default value for instances of version 4.4 and later is **cloud_essd1**.
	//
	// > - The default value for instances of version 4.2 and earlier is **local_ssd**.
	//
	// example:
	//
	// cloud_essd1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Specifies whether to enable TTL index cleanup for the restored instance. A value of true or an empty value indicates that cleanup is enabled.
	TTLEnabled *bool `json:"TTLEnabled,omitempty" xml:"TTLEnabled,omitempty"`
	// The custom tags.
	Tag []*CreateDBInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the zone ID.
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

func (s *CreateDBInstanceRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *CreateDBInstanceRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *CreateDBInstanceRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
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

func (s *CreateDBInstanceRequest) GetTTLEnabled() *bool {
	return s.TTLEnabled
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

func (s *CreateDBInstanceRequest) SetDryRun(v bool) *CreateDBInstanceRequest {
	s.DryRun = &v
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

func (s *CreateDBInstanceRequest) SetMaintainEndTime(v string) *CreateDBInstanceRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetMaintainStartTime(v string) *CreateDBInstanceRequest {
	s.MaintainStartTime = &v
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

func (s *CreateDBInstanceRequest) SetTTLEnabled(v bool) *CreateDBInstanceRequest {
	s.TTLEnabled = &v
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

type CreateDBInstanceRequestTag struct {
	// The key of the tag.
	//
	// > - **N*	- specifies the sequence number of the tag. For example, **Tag.1.Key*	- specifies the key of the first tag, and **Tag.2.Key*	- specifies the key of the second tag.
	//
	// example:
	//
	// testdatabase
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// > **N*	- specifies the sequence number of the tag. For example, **Tag.1.Value*	- specifies the value of the first tag, and **Tag.2.Value*	- specifies the value of the second tag.
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
