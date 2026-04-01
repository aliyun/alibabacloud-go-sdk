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
	// Specifies whether to automatically create database proxies. Valid values:
	//
	// 	- **true**: automatically creates database proxies. By default, general-purpose database proxies are created.
	//
	// 	- **false**: does not automatically create database proxies.
	//
	// example:
	//
	// false
	AutoCreateProxy *bool `json:"AutoCreateProxy,omitempty" xml:"AutoCreateProxy,omitempty"`
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 1.  **true**: automatically completes the payment. Make sure that your account balance is sufficient.
	//
	// 2.  **false**: does not automatically complete the payment. An unpaid order is generated.
	//
	// >  Default value: true. If your account balance is insufficient, you can set the AutoPay parameter to false to generate an unpaid order. Then, you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable the auto-renewal feature for the read-only instance. If you set the PayType parameter to Prepaid, you must also specify this parameter. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
	//
	// > 	- If you set the Period parameter to Month, the auto-renewal cycle is one month.
	//
	// > 	- If you set the Period parameter to Year, the auto-renewal cycle is one year.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// 	- **true**: uses a coupon.
	//
	// 	- **false*	- (default): does not use a coupon.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// false
	BpeEnabled *string `json:"BpeEnabled,omitempty" xml:"BpeEnabled,omitempty"`
	// An invalid parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability*	- (default): RDS High-availability Edition
	//
	// 	- **AlwaysOn**: RDS Cluster Edition
	//
	// >  The read-only instances of the primary instance that run PostgreSQL and use cloud disks run RDS Basic Edition. Therefore, set this parameter to **Basic**.
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
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type of the read-only instance. For more information, see [Read-only instance types](https://help.aliyun.com/document_detail/145759.html). We recommend that you specify an instance type whose specifications are higher than or equal to the specifications of the instance type of the primary instance. If the specifications of the read-only instance are lower than the specifications of the primary instance, the read-only instance may encounter issues such as high latency and heavy load.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds.mys2.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the read-only instance. The description must be 2 to 256 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The value must start with a letter
	//
	// > The value cannot start with [http:// or https://.](http://https://。)
	//
	// example:
	//
	// Test read-only instance
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The primary instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The storage capacity of the read-only instance. The storage capacity of the read-only instance must be greater than or equal to that of the primary instance. For more information, see the **Storage capacity*	- column in [Read-only instance types](https://help.aliyun.com/document_detail/145759.html). This value must be a multiple of 5. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: local SSDs
	//
	// 	- **cloud_ssd**: standard SSDs
	//
	// 	- **cloud_essd**: enhanced SSDs (ESSDs) of performance level 1 (PL1)
	//
	// 	- **cloud_essd2**: ESSDs of PL2
	//
	// 	- **cloud_essd3**: ESSDs of PL3
	//
	// > 	- If the primary instance runs MySQL with local disks, you must set this parameter to **local_ssd**. If the primary instance runs MySQL with cloud disks, you must set this parameter to cloud_ssd, cloud_essd, cloud_essd2, or cloud_essd3.
	//
	// > 	- If the primary instance runs SQL Server, you must set this parameter to cloud_ssd, cloud_essd, cloud_essd2, or cloud_essd3.
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The ID of the dedicated cluster to which the read-only instance belongs. This parameter is valid when you create the read-only instance in a dedicated cluster.
	//
	// example:
	//
	// dhg-4n****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// Specifies whether to enable the release protection feature for the read-only instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  You can enable the release protection feature for the read-only instance only when you set the **PayType*	- parameter to **Postpaid**.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The version of the database engine. The read-only instance and the primary instance must run the same major engine version.
	//
	// 	- If the read-only instance runs MySQL, set this parameter to **5.6**, **5.7**, or **8.0**.
	//
	// 	- If the read-only instance runs MySQL, set this parameter to **2017_ent, 2019_ent, or 2022_ent**.
	//
	// 	- If the read-only instance runs PostgreSQL, set this parameter to **10.0, 11.0, 12.0, 13.0, 14.0, or 15.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	GdnInstanceName *string `json:"GdnInstanceName,omitempty" xml:"GdnInstanceName,omitempty"`
	// The network type of the read-only instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **Classic**
	//
	// Default value: VPC. If you set this parameter to VPC, you must also specify the **VPCId*	- and **VSwitchId*	- parameters.
	//
	// >  The network type of the read-only instance can be different from the network type of the primary instance.
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	InstructionSetArch *string `json:"InstructionSetArch,omitempty" xml:"InstructionSetArch,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	IsAnalyticReadOnlyIns *bool   `json:"IsAnalyticReadOnlyIns,omitempty" xml:"IsAnalyticReadOnlyIns,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the read-only instance. Valid values:
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
	// The renewal cycle of the read-only instance. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The port that can be initialized when you create a read-only ApsaraDB RDS for MySQL instance.
	//
	// Valid values: 1000 to 65534.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the read-only instance. The private IP address must be within the CIDR block that is supported by the specified vSwitch. The system assigns a private IP address to the read-only instance based on the values of the **VPCId*	- and **VSwitchId*	- parameters.
	//
	// example:
	//
	// 172.16.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 717446260784
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID. The read-only instance and the primary instance must reside in the same region. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the host on which the primary instance resides. This parameter is valid when you create the read-only instance in a dedicated cluster.
	//
	// example:
	//
	// i-bp****
	TargetDedicatedHostIdForMaster *string `json:"TargetDedicatedHostIdForMaster,omitempty" xml:"TargetDedicatedHostIdForMaster,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	TddlBizType *string `json:"TddlBizType,omitempty" xml:"TddlBizType,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	TddlRegionConfig *string `json:"TddlRegionConfig,omitempty" xml:"TddlRegionConfig,omitempty"`
	// The subscription duration of the read-only instance. Valid values:
	//
	// 	- If you set the **Period*	- parameter to **Year**, the value of the **UsedTime*	- parameter ranges from **1*	- to **5**.
	//
	// 	- If you set the **Period*	- parameter to **Month**, the value of the **UsedTime*	- parameter ranges from **1*	- to **9**.
	//
	// > If you set the **PayType*	- parameter to **Prepaid**, you must specify the UsedTime parameter.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The virtual private cloud (VPC) ID of the read-only instance. If you leave the **InstanceNetworkType*	- parameter empty or set it to **VPC**, you must also specify this parameter.
	//
	// > 	- If the primary instance uses local disks, the read-only instance and the primary instance can belong to the same VPC or different VPCs.
	//
	// > 	- If the primary instance uses cloud disks, the read-only instance and the primary instance must belong to the same VPC.
	//
	// example:
	//
	// vpc-uf6f7l4fg90****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the read-only instance. If you leave the **InstanceNetworkType*	- parameter empty or set it to **VPC**, you must specify the VSwitchId parameter.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID. You can call the DescribeRegions operation to query the zone ID.
	//
	// 	- If you use the single-zone deployment method, set this parameter to the ID of one zone. Example: `cn-hangzhou-b`.
	//
	// 	- If you use the multi-zone deployment method, set this parameter to the IDs of multiple zones and separate the IDs with colons (:). Example: `cn-hangzhou-b:cn-hangzhou-c`.
	//
	// 	- The number of zone IDs that you specify must be less than or equal to the number of nodes created for the read-only instance. If you create a read-only instance that runs RDS Basic Edition, only one node is provisioned. If you create a read-only instance that runs RDS High-availability Edition, one primary node and one secondary node are provisioned.
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
