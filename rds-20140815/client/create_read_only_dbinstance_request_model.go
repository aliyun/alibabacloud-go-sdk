// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReadOnlyDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateProxy(v bool) *CreateReadOnlyDBInstanceRequest
	GetAutoCreateProxy() *bool
	SetAutoPay(v bool) *CreateReadOnlyDBInstanceRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *CreateReadOnlyDBInstanceRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *CreateReadOnlyDBInstanceRequest
	GetAutoUseCoupon() *bool
	SetBpeEnabled(v string) *CreateReadOnlyDBInstanceRequest
	GetBpeEnabled() *string
	SetBurstingEnabled(v bool) *CreateReadOnlyDBInstanceRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *CreateReadOnlyDBInstanceRequest
	GetCategory() *string
	SetClientToken(v string) *CreateReadOnlyDBInstanceRequest
	GetClientToken() *string
	SetCustomExtraInfo(v string) *CreateReadOnlyDBInstanceRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *CreateReadOnlyDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CreateReadOnlyDBInstanceRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *CreateReadOnlyDBInstanceRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *CreateReadOnlyDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *CreateReadOnlyDBInstanceRequest
	GetDBInstanceStorageType() *string
	SetDedicatedHostGroupId(v string) *CreateReadOnlyDBInstanceRequest
	GetDedicatedHostGroupId() *string
	SetDeletionProtection(v bool) *CreateReadOnlyDBInstanceRequest
	GetDeletionProtection() *bool
	SetEngineVersion(v string) *CreateReadOnlyDBInstanceRequest
	GetEngineVersion() *string
	SetGdnInstanceName(v string) *CreateReadOnlyDBInstanceRequest
	GetGdnInstanceName() *string
	SetInstanceNetworkType(v string) *CreateReadOnlyDBInstanceRequest
	GetInstanceNetworkType() *string
	SetInstructionSetArch(v string) *CreateReadOnlyDBInstanceRequest
	GetInstructionSetArch() *string
	SetIoAccelerationEnabled(v string) *CreateReadOnlyDBInstanceRequest
	GetIoAccelerationEnabled() *string
	SetIsAnalyticReadOnlyIns(v bool) *CreateReadOnlyDBInstanceRequest
	GetIsAnalyticReadOnlyIns() *bool
	SetOwnerAccount(v string) *CreateReadOnlyDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateReadOnlyDBInstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateReadOnlyDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateReadOnlyDBInstanceRequest
	GetPeriod() *string
	SetPort(v string) *CreateReadOnlyDBInstanceRequest
	GetPort() *string
	SetPrivateIpAddress(v string) *CreateReadOnlyDBInstanceRequest
	GetPrivateIpAddress() *string
	SetPromotionCode(v string) *CreateReadOnlyDBInstanceRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateReadOnlyDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateReadOnlyDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateReadOnlyDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateReadOnlyDBInstanceRequest
	GetResourceOwnerId() *int64
	SetTargetDedicatedHostIdForMaster(v string) *CreateReadOnlyDBInstanceRequest
	GetTargetDedicatedHostIdForMaster() *string
	SetTddlBizType(v string) *CreateReadOnlyDBInstanceRequest
	GetTddlBizType() *string
	SetTddlRegionConfig(v string) *CreateReadOnlyDBInstanceRequest
	GetTddlRegionConfig() *string
	SetUsedTime(v string) *CreateReadOnlyDBInstanceRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateReadOnlyDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateReadOnlyDBInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateReadOnlyDBInstanceRequest
	GetZoneId() *string
}

type CreateReadOnlyDBInstanceRequest struct {
	// Specifies whether to automatically create a database proxy. Valid values:
	//
	// - **true**: enables automatic creation. By default, a general-purpose database proxy is created.
	//
	// - **false**: does not enable automatic creation of a database proxy.
	//
	// example:
	//
	// false
	AutoCreateProxy *bool `json:"AutoCreateProxy,omitempty" xml:"AutoCreateProxy,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: enables automatic payment. Make sure that your account balance is sufficient.
	//
	// - **false**: generates an order without charging your account.
	//
	//
	//
	//
	// > The default value is true. If your payment method has an insufficient balance, set AutoPay to false. In this case, an unpaid order is generated. You can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter is required only for subscription instances. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false**: disables auto-renewal.
	//
	// > 	- If you purchase the instance on a monthly basis, the auto-renewal cycle is one month.
	//
	// > 	- If you purchase the instance on a yearly basis, the auto-renewal cycle is one year.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use coupons. Valid values:
	//
	// 	- **true**: uses coupons.
	//
	// 	- **false*	- (default): does not use coupons.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	BpeEnabled    *string `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	// Specifies whether to enable the I/O performance burst feature for [Premium ESSDs](https://help.aliyun.com/document_detail/2340501.html). Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- **Basic**: Basic Edition
	//
	// 	- **HighAvailability**: High-availability Edition (default)
	//
	// 	- **AlwaysOn**: Cluster Edition
	//
	// <props="china">	- **Finance**: Finance Edition
	//
	// > The read-only instances of ApsaraDB RDS for PostgreSQL cloud disk instances use the Basic Edition. You must set this parameter to **Basic**.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// None
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type. For more information, see [Read-only instance types](https://help.aliyun.com/document_detail/145759.html). We recommend that the specifications of the read-only instance be equal to or higher than those of the primary instance. Otherwise, the read-only instance may experience high latency and heavy loads.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysqlro.n2.small.1c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance description. The description must be 2 to 256 characters in length and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	//
	// > The description cannot start with http:// or https://.
	//
	// example:
	//
	// testReadOnly
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The primary instance ID. You can call [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Instance storage capacity. Instance storage capacity of the read-only instance must be greater than or equal to that of the primary instance. For more information, see the **Storage capacity*	- column in [Read-only instance types](https://help.aliyun.com/document_detail/145759.html). The value is incremented in units of 5 GB. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: Premium Local SSDs
	//
	// 	- **cloud_ssd**: standard SSDs
	//
	// 	- **cloud_essd**: PL1 ESSDs
	//
	// 	- **cloud_essd2**: PL2 ESSDs
	//
	// 	- **cloud_essd3**: PL3 ESSDs
	//
	// 	- **general_essd**: Premium ESSDs
	//
	//
	// > 	- If the primary ApsaraDB RDS for MySQL instance uses Premium Local SSDs, only **local_ssd*	- is supported. If the primary ApsaraDB RDS for MySQL instance uses cloud disks, premium performance disk storage types are supported.
	//
	// > 	- ApsaraDB RDS for SQL Server supports premium performance disk storage types.
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The dedicated cluster ID. This parameter is required when you create a read-only instance in a dedicated cluster.
	//
	// example:
	//
	// dhg-4n****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to enable the release protection feature for the instance. Valid values:
	//
	// 	- **true**: enables release protection.
	//
	// 	- **false**: disables release protection. (default)
	//
	// > This feature is supported only when the **billing method*	- is **pay-as-you-go**.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The database engine version. The version must be the same as that of the primary instance.
	//
	// 	- Valid values for MySQL: **5.6**, **5.7**, and **8.0**.
	//
	// 	- Valid values for SQL Server: **2017_ent, 2019_ent, and 2022_ent**.
	//
	// 	- Valid values for PostgreSQL: **10.0, 11.0, 12.0, 13.0, 14.0, and 15.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// test
	GdnInstanceName *string `json:"GdnInstanceName,omitempty" xml:"GdnInstanceName,omitempty"`
	// The network type of the read-only instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **Classic**: classic network
	//
	// By default, a VPC-connected instance is created. You must also specify **VPCId*	- and **VSwitchId**.
	//
	// > The network type of the read-only instance can be different from that of the primary instance.
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// test
	InstructionSetArch *string `json:"InstructionSetArch,omitempty" xml:"InstructionSetArch,omitempty"`
	// Specifies whether to enable the [Buffer Pool Extension (BPE)](https://help.aliyun.com/document_detail/2527067.html) feature for Premium ESSDs. Valid values:
	//
	//  - **1**: enables the feature.
	//
	//  - **0**: does not enable the feature.
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to create a DuckDB-based analytical instance. Valid values:
	//
	// - **true**: creates a DuckDB-based analytical instance.
	//
	// - **false**: does not create a DuckDB-based analytical instance.
	//
	// > Only ApsaraDB RDS for MySQL and ApsaraDB RDS for PostgreSQL support DuckDB-based analytical instances.
	IsAnalyticReadOnlyIns *bool   `json:"IsAnalyticReadOnlyIns,omitempty" xml:"IsAnalyticReadOnlyIns,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the instance. Valid values:
	//
	// 	- **Year**: yearly subscription
	//
	// 	- **Month**: monthly subscription
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The port that is initialized when you create a read-only instance for an ApsaraDB RDS for MySQL primary instance.
	//
	// Valid values: 1000 to 65534.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The internal IP address of the read-only instance. The IP address must be within the address range of the specified vSwitch. The system automatically allocates an internal IP address based on the values of **VPCId*	- and **VSwitchId*	- by default.
	//
	// example:
	//
	// 172.16.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 71744626****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID. The read-only instance must reside in the same region as the primary instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The host ID of the primary instance in the dedicated cluster. This parameter is required when you create a read-only instance in a dedicated cluster.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// test
	TddlBizType *string `json:"TddlBizType,omitempty" xml:"TddlBizType,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// test
	TddlRegionConfig *string `json:"TddlRegionConfig,omitempty" xml:"TddlRegionConfig,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- If **Period*	- is set to **Year**, the valid values of **UsedTime*	- are **1*	- to **5**.
	//
	// 	- If **Period*	- is set to **Month**, the valid values of **UsedTime*	- are **1*	- to **9**.
	//
	// > This parameter is required when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID of the read-only instance. This parameter is required when **InstanceNetworkType*	- is left empty or set to **VPC**.
	//
	// > 	- If the storage type of the primary instance is Premium Local SSDs, the read-only instance can use any VPC.
	//
	// > 	- If the storage type of the primary instance is cloud disks, the VPC of the read-only instance must be the same as that of the primary instance.
	//
	// example:
	//
	// vpc-uf6f7l4fg90****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the read-only instance. This parameter is required when **InstanceNetworkType*	- is left empty or set to **VPC**.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) to query the most recent zone list.
	//
	// - For single-zone deployment, specify one zone ID, such as `cn-hangzhou-b`.
	//
	// - For multi-zone deployment, specify multiple zone IDs separated by colons (:), such as `cn-hangzhou-b:cn-hangzhou-c`.
	//
	// - The number of specified zones must be less than or equal to the number of nodes in the read-only instance. A Basic Edition read-only instance contains only one node. A High-availability Edition read-only instance contains two nodes (one primary node and one secondary node).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateReadOnlyDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReadOnlyDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateReadOnlyDBInstanceRequest) GetAutoCreateProxy() *bool {
	return s.AutoCreateProxy
}

func (s *CreateReadOnlyDBInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateReadOnlyDBInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *CreateReadOnlyDBInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateReadOnlyDBInstanceRequest) GetBpeEnabled() *string {
	return s.BpeEnabled
}

func (s *CreateReadOnlyDBInstanceRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateReadOnlyDBInstanceRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateReadOnlyDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateReadOnlyDBInstanceRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *CreateReadOnlyDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CreateReadOnlyDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateReadOnlyDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateReadOnlyDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CreateReadOnlyDBInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CreateReadOnlyDBInstanceRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *CreateReadOnlyDBInstanceRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateReadOnlyDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateReadOnlyDBInstanceRequest) GetGdnInstanceName() *string {
	return s.GdnInstanceName
}

func (s *CreateReadOnlyDBInstanceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CreateReadOnlyDBInstanceRequest) GetInstructionSetArch() *string {
	return s.InstructionSetArch
}

func (s *CreateReadOnlyDBInstanceRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *CreateReadOnlyDBInstanceRequest) GetIsAnalyticReadOnlyIns() *bool {
	return s.IsAnalyticReadOnlyIns
}

func (s *CreateReadOnlyDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateReadOnlyDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateReadOnlyDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateReadOnlyDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateReadOnlyDBInstanceRequest) GetPort() *string {
	return s.Port
}

func (s *CreateReadOnlyDBInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateReadOnlyDBInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateReadOnlyDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateReadOnlyDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateReadOnlyDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateReadOnlyDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateReadOnlyDBInstanceRequest) GetTargetDedicatedHostIdForMaster() *string {
	return s.TargetDedicatedHostIdForMaster
}

func (s *CreateReadOnlyDBInstanceRequest) GetTddlBizType() *string {
	return s.TddlBizType
}

func (s *CreateReadOnlyDBInstanceRequest) GetTddlRegionConfig() *string {
	return s.TddlRegionConfig
}

func (s *CreateReadOnlyDBInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateReadOnlyDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateReadOnlyDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateReadOnlyDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateReadOnlyDBInstanceRequest) SetAutoCreateProxy(v bool) *CreateReadOnlyDBInstanceRequest {
	s.AutoCreateProxy = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetAutoPay(v bool) *CreateReadOnlyDBInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetAutoRenew(v string) *CreateReadOnlyDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetAutoUseCoupon(v bool) *CreateReadOnlyDBInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetBpeEnabled(v string) *CreateReadOnlyDBInstanceRequest {
	s.BpeEnabled = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetBurstingEnabled(v bool) *CreateReadOnlyDBInstanceRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetCategory(v string) *CreateReadOnlyDBInstanceRequest {
	s.Category = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetClientToken(v string) *CreateReadOnlyDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetCustomExtraInfo(v string) *CreateReadOnlyDBInstanceRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetDBInstanceClass(v string) *CreateReadOnlyDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetDBInstanceDescription(v string) *CreateReadOnlyDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetDBInstanceId(v string) *CreateReadOnlyDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetDBInstanceStorage(v int32) *CreateReadOnlyDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetDBInstanceStorageType(v string) *CreateReadOnlyDBInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetDedicatedHostGroupId(v string) *CreateReadOnlyDBInstanceRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetDeletionProtection(v bool) *CreateReadOnlyDBInstanceRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetEngineVersion(v string) *CreateReadOnlyDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetGdnInstanceName(v string) *CreateReadOnlyDBInstanceRequest {
	s.GdnInstanceName = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetInstanceNetworkType(v string) *CreateReadOnlyDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetInstructionSetArch(v string) *CreateReadOnlyDBInstanceRequest {
	s.InstructionSetArch = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetIoAccelerationEnabled(v string) *CreateReadOnlyDBInstanceRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetIsAnalyticReadOnlyIns(v bool) *CreateReadOnlyDBInstanceRequest {
	s.IsAnalyticReadOnlyIns = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetOwnerAccount(v string) *CreateReadOnlyDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetOwnerId(v int64) *CreateReadOnlyDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetPayType(v string) *CreateReadOnlyDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetPeriod(v string) *CreateReadOnlyDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetPort(v string) *CreateReadOnlyDBInstanceRequest {
	s.Port = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetPrivateIpAddress(v string) *CreateReadOnlyDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetPromotionCode(v string) *CreateReadOnlyDBInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetRegionId(v string) *CreateReadOnlyDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetResourceGroupId(v string) *CreateReadOnlyDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateReadOnlyDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetResourceOwnerId(v int64) *CreateReadOnlyDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetTargetDedicatedHostIdForMaster(v string) *CreateReadOnlyDBInstanceRequest {
	s.TargetDedicatedHostIdForMaster = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetTddlBizType(v string) *CreateReadOnlyDBInstanceRequest {
	s.TddlBizType = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetTddlRegionConfig(v string) *CreateReadOnlyDBInstanceRequest {
	s.TddlRegionConfig = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetUsedTime(v string) *CreateReadOnlyDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetVPCId(v string) *CreateReadOnlyDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetVSwitchId(v string) *CreateReadOnlyDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) SetZoneId(v string) *CreateReadOnlyDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateReadOnlyDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
