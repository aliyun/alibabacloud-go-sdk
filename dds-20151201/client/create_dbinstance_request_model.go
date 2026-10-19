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
	// - Must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - Special characters include !@#$%^&*()_+-=.
	//
	// - Must be 8 to 32 characters in length.
	//
	// > For information about how to database connection failures caused by special characters in the password, see [How do I database connection failures caused by special characters in the connection string?](https://help.aliyun.com/document_detail/471568.html).
	//
	// example:
	//
	// 123456Aa
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// - **true**: enables auto-renewal.
	//
	// - **false**: default value. Disables auto-renewal. You must manually renew the instance.
	//
	// > This parameter is available and optional when **ChargeType*	- is set to **PrePaid**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The backup ID. You can invoke the [DescribeBackups](https://help.aliyun.com/document_detail/62172.html) operation to query the backup ID.
	//
	// > This parameter is required only when you invoke this operation to clone a clone instance from a backup point. You must also configure the **SrcDBInstanceId*	- parameter.
	//
	// example:
	//
	// 32994****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The additional information about the business.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PostPaid**: default value. Pay-as-you-go.
	//
	// - **PrePaid**: Subscription.
	//
	// > If you set this parameter to **PrePaid**, you must also configure the **Period*	- parameter.
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
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **default*	- or **null*	- (default): uses a coupon.
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
	// The instance name. The name must meet the following requirements:
	//
	// - Must start with a Chinese character or a letter.
	//
	// - Can contain digits, Chinese characters, letters, underscores (_), periods (.), and hyphens (-).
	//
	// - Must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance storage capacity. Unit: GB.
	//
	// The valid values depend on the instance type. For details, see [Replica set instance specifications](https://help.aliyun.com/document_detail/311410.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The name of the database.
	//
	// > When you invoke this operation to clone a clone instance, you can specify this parameter to clone specific databases. If you do not specify this parameter, all databases are cloned.
	//
	// example:
	//
	// mongodbtest
	DatabaseNames *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	// Specifies whether to enable cloud disk encryption.
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
	// The database engine. The value is fixed as **MongoDB**.
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
	// > If you call this operation to clone clone instance or restore clone instance from the recycle bin, set this parameter to the same value as that of the source instance.
	//
	// 	Warning: Versions 3.4 and earlier are no longer available for purchase.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The global IP whitelist templates for the instance. Separate multiple templates with commas (,). Duplicate templates are not allowed. (This feature is in canary release.)
	//
	// example:
	//
	// g-qxieqf40xjst1ngpr3jz
	GlobalSecurityGroupIds *string `json:"GlobalSecurityGroupIds,omitempty" xml:"GlobalSecurityGroupIds,omitempty"`
	// The zone where the hidden node is deployed to implement multi-zone deployment. Valid values:
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
	// > - This parameter is available when the instance uses cloud disks.
	//
	// > - The value of this parameter cannot be the same as the values of **ZoneId*	- and **SecondaryZoneId**.
	//
	// example:
	//
	// cn-hangzhou-i
	HiddenZoneId *string `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	// The network type of the instance. Valid values:
	//
	// **VPC**: Virtual Private Cloud.
	//
	// example:
	//
	// VPC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the instance. Unit: months.
	//
	// Valid values: **1*	- to **9*	- (integers), **12**, **24**, **36**, and **60**.
	//
	// > This parameter is available and required when **ChargeType*	- is set to **PrePaid**.
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
	// The **number of read-only nodes*	- in the replica set instance. Valid values: **0*	- to **5*	- (integers). Default value: **0**.
	//
	// example:
	//
	// 0
	ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
	//
	// > If you call this operation to clone clone instance or restore clone instance from the recycle bin, set this parameter to the same value as that of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The **number of primary and secondary nodes*	- in the replica set instance. Valid values:
	//
	// - **3*	- (default)
	//
	// - **5**
	//
	// - **7**
	//
	// 	Notice: This parameter is not required for single-node instances.
	//
	// example:
	//
	// 3
	ReplicationFactor    *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to recover the instance to. You can specify any point in time within the last 7 days. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// > This parameter is required only when you invoke this operation to clone a clone instance to a point in time. You must also configure the **SrcDBInstanceId*	- parameter.
	//
	// example:
	//
	// 2022-03-13T12:11:14Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The backup restoration type. Valid values:
	//
	// - 0: recovers the instance to a specified backup set.
	//
	// - 1: recovers the instance to a specified point in time.
	//
	// - 2: recovers a released instance to a specified backup set.
	//
	// - 3: recovers the instance to a specified geo-redundancy backup set.
	//
	// example:
	//
	// 0
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The zone where the secondary node is deployed to implement multi-zone deployment. Valid values:
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
	// > - This parameter is available when the instance uses cloud disks.
	//
	// > - The value of this parameter cannot be the same as the values of **ZoneId*	- and **HiddenZoneId**.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The IP whitelist of the instance. Separate multiple IP addresses with commas (,). Duplicate IP addresses are not allowed. The following formats are supported:
	//
	// - 0.0.0.0/0
	//
	// - IP address format, such as 10.23.12.24.
	//
	// - CIDR block format, such as 10.23.12.0/24. In this example, 24 indicates that the prefix of the CIDR block is 24 bits in length. The prefix length ranges from 1 to 32.
	//
	// > - Each instance supports a maximum of 1,000 IP addresses or CIDR blocks across all IP whitelist groups.
	//
	// > - 0.0.0.0/0 allows all IP addresses to access the instance database. This is a high-risk setting. Proceed with caution.
	//
	// example:
	//
	// 192.168.xx.xx,192.168.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The source instance ID.
	//
	// > If you call this operation to clone clone instance, you must configure this parameter along with the **BackupId*	- or **RestoreTime*	- parameter. If you call this operation to restore clone instance from the recycle bin, you only need to configure this parameter. The **BackupId*	- and **RestoreTime*	- parameters are not required.
	//
	// example:
	//
	// dds-bp1ee12ad351****
	SrcDBInstanceId *string `json:"SrcDBInstanceId,omitempty" xml:"SrcDBInstanceId,omitempty"`
	// The region where the source instance resides.
	//
	// > This parameter is required when the restoration type is 2 or 3.
	//
	// example:
	//
	// 2
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The storage engine of the instance. The value is fixed as **WiredTiger**.
	//
	// > - If you call this operation to clone clone instance or restore clone instance from the recycle bin, set this parameter to the same value as that of the source instance.
	//
	// > - For information about the constraints on storage engines and versions, see [Versions and storage engines](https://help.aliyun.com/document_detail/61906.html).
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
	// - **cloud_auto**: ESSD AutoPL disk.
	//
	// - **local_ssd**: local SSD.
	//
	// > - For single-node instances, if you specify cloud_essd1, the instance uses a cloud_essd disk.
	//
	// > - ESSD AutoPL disks are currently available only on the China site.
	//
	// > - For instances running version 4.4 or later, the default value is **cloud_essd1**.
	//
	// > - For instances running version 4.2 or earlier, the default value is **local_ssd**.
	//
	// example:
	//
	// cloud_essd1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The user-defined tags.
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
	// The tag key.
	//
	// > - **N*	- indicates the sequence number of the tag key. For example, **Tag.1.Key*	- indicates the key of the first tag, and **Tag.2.Key*	- indicates the key of the second tag.
	//
	// example:
	//
	// testdatabase
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// > **N*	- indicates the sequence number of the tag value. For example, **Tag.1.Value*	- indicates the value of the first tag, and **Tag.2.Value*	- indicates the value of the second tag.
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
