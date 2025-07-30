// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyInstanceSpecRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *ModifyInstanceSpecRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *ModifyInstanceSpecRequest
	GetClientToken() *string
	SetCouponNo(v string) *ModifyInstanceSpecRequest
	GetCouponNo() *string
	SetEffectiveTime(v string) *ModifyInstanceSpecRequest
	GetEffectiveTime() *string
	SetForceTrans(v bool) *ModifyInstanceSpecRequest
	GetForceTrans() *bool
	SetForceUpgrade(v bool) *ModifyInstanceSpecRequest
	GetForceUpgrade() *bool
	SetInstanceClass(v string) *ModifyInstanceSpecRequest
	GetInstanceClass() *string
	SetInstanceId(v string) *ModifyInstanceSpecRequest
	GetInstanceId() *string
	SetMajorVersion(v string) *ModifyInstanceSpecRequest
	GetMajorVersion() *string
	SetNodeType(v string) *ModifyInstanceSpecRequest
	GetNodeType() *string
	SetOrderType(v string) *ModifyInstanceSpecRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *ModifyInstanceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceSpecRequest
	GetOwnerId() *int64
	SetReadOnlyCount(v int32) *ModifyInstanceSpecRequest
	GetReadOnlyCount() *int32
	SetRegionId(v string) *ModifyInstanceSpecRequest
	GetRegionId() *string
	SetReplicaCount(v int32) *ModifyInstanceSpecRequest
	GetReplicaCount() *int32
	SetResourceOwnerAccount(v string) *ModifyInstanceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceSpecRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceSpecRequest
	GetSecurityToken() *string
	SetShardCount(v int32) *ModifyInstanceSpecRequest
	GetShardCount() *int32
	SetSlaveReadOnlyCount(v int32) *ModifyInstanceSpecRequest
	GetSlaveReadOnlyCount() *int32
	SetSlaveReplicaCount(v int32) *ModifyInstanceSpecRequest
	GetSlaveReplicaCount() *int32
	SetSourceBiz(v string) *ModifyInstanceSpecRequest
	GetSourceBiz() *string
	SetStorage(v int32) *ModifyInstanceSpecRequest
	GetStorage() *int32
	SetStorageType(v string) *ModifyInstanceSpecRequest
	GetStorageType() *string
}

type ModifyInstanceSpecRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true*	- (default): enables automatic payment.
	//
	// 	- **false**: disables automatic payment. If you set this parameter to **false**, the instance must be manually renewed before it expires. For more information, see [Renew an instance](https://help.aliyun.com/document_detail/26352.html).
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the promotional event or the business information.
	//
	// example:
	//
	// 000000001
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The time when you want the configurations to be changed. Valid values:
	//
	// 	- **Immediately*	- (default): immediately changes the configurations.
	//
	// 	- **MaintainTime**: changes the configurations within the maintenance window. You can call the [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) operation to change the maintenance window.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// Specifies whether to enable forced transmission during a configuration change. Valid values:
	//
	// 	- **false*	- (default): Before the configuration change, the system checks the minor version of the instance. If the minor version of the instance is outdated, an error is reported. You must update the minor version of the instance and try again.
	//
	// 	- **true**: The system skips the version check and directly performs the configuration change.
	//
	// example:
	//
	// false
	ForceTrans *bool `json:"ForceTrans,omitempty" xml:"ForceTrans,omitempty"`
	// Specifies whether to forcibly change the configurations. Valid values:
	//
	// 	- **false**: The system does not forcefully change the configurations.
	//
	// 	- **true*	- (default): The system forcefully changes the configurations.
	//
	// example:
	//
	// true
	ForceUpgrade *bool `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	// The new instance type. You can call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/473765.html) operation to query the instance types available for configuration change within the zone to which the instance belongs.
	//
	// >  For more information about the instance types, see [Overview](https://help.aliyun.com/document_detail/26350.html).
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The major version of the classic instance that you want to upgrade. Valid values: **2.8**, **4.0**, and **5.0**.
	//
	// >  The **InstanceClass*	- parameter is required when you upgrade the instance version. This parameter indicates that you can upgrade the instance version only when you update the instance specifications. If you only need to upgrade the instance version, call the [ModifyInstanceMajorVersion](https://help.aliyun.com/document_detail/473776.html) operation.
	//
	// example:
	//
	// 5.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- **MASTER_SLAVE**: high availability (master-replica)
	//
	// 	- **STAND_ALONE**: standalone
	//
	// 	- **double**: master-replica
	//
	// 	- **single**: standalone
	//
	// >  To create a cloud-native instance, set this parameter to **MASTER_SLAVE*	- or **STAND_ALONE**. To create a classic instance, set this parameter to **double*	- or **single**.
	//
	// example:
	//
	// MASTER_SLAVE
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The change type. This parameter is required when you change the configurations of a subscription instance. Valid values:
	//
	// 	- **UPGRADE*	- (default): upgrades the configurations of the subscription instance.
	//
	// 	- **DOWNGRADE**: downgrades the configurations of the subscription instance.
	//
	// >
	//
	// 	- To downgrade a subscription instance, you must set this parameter to **DOWNGRADE**.
	//
	// 	- If the price of an instance increases after its configurations are changed, the instance is upgraded. If the price decreases, the instance is downgraded. For example, the price of an 8 GB read/write splitting instance with five read replicas is higher than that of a 16 GB cluster instance. If you want to change a 16 GB cluster instance to an 8 GB read/write splitting instance with five read replicas, you must upgrade the instance.
	//
	// example:
	//
	// DOWNGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read replicas in the primary zone. Valid values: 0 to 5. This parameter applies only to the following scenarios:
	//
	// 	- If the instance is a cloud-native standard instance, you can set this parameter to a value greater than 0 to enable the read/write splitting architecture.
	//
	// 	- If the instance is a cloud-native read/write splitting instance, you can use this parameter to customize the number of read replicas. You can also set this parameter to 0 to disable the read/write splitting architecture and switch the instance to the standard architecture.
	//
	// example:
	//
	// 5
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of replica nodes in the primary zone. This parameter is applicable only to cloud-native multi-replica cluster instances. Valid values: 1 to 4.
	//
	// >
	//
	// 	- The sum of the values of this parameter and the SlaveReplicaCount parameter cannot be greater than 4.
	//
	// 	- You can specify either ReplicaCount or ReadOnlyCount.
	//
	// 	- A master-replica instance cannot contain multiple replica nodes.
	//
	// example:
	//
	// 1
	ReplicaCount         *int32  `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards. This parameter is applicable only to cloud-native cluster instances.
	//
	// >
	//
	// 	- If you want to change a cloud-native cluster instance to a standard instance, you must explicitly set the ShardCount parameter to 1 and specify the specifications of the master-replica instance.
	//
	// 	- To change a cloud-native standard instance to a cluster instance, you must explicitly set the ShardCount parameter to a value greater than 1 and specify the specifications of the cluster instance.
	//
	// example:
	//
	// 8
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The number of read replicas in the secondary zone when you create a read/write splitting instance that is deployed across multiple zones. Valid values: 1 to 9. The sum of the values of this parameter and the ReadOnlyCount parameter cannot be greater than 9.
	//
	// example:
	//
	// 2
	SlaveReadOnlyCount *int32 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// The number of replica nodes in the secondary zone when you create a cloud-native multi-replica cluster instance that is deployed across multiple zones. The sum of the values of this parameter and the ReplicaCount parameter cannot be greater than 4.
	//
	// >  When you create a cloud-native multi-replica cluster instance that is deployed across multiple zones, you must specify both SlaveReplicaCount and SecondaryZoneId.
	//
	// example:
	//
	// 1
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// The source of the operation. This parameter is used only for internal maintenance. You do not need to specify this parameter.
	//
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
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
	// example:
	//
	// essd_pl1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ModifyInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceSpecRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *ModifyInstanceSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceSpecRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *ModifyInstanceSpecRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyInstanceSpecRequest) GetForceTrans() *bool {
	return s.ForceTrans
}

func (s *ModifyInstanceSpecRequest) GetForceUpgrade() *bool {
	return s.ForceUpgrade
}

func (s *ModifyInstanceSpecRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *ModifyInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceSpecRequest) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *ModifyInstanceSpecRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ModifyInstanceSpecRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ModifyInstanceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceSpecRequest) GetReadOnlyCount() *int32 {
	return s.ReadOnlyCount
}

func (s *ModifyInstanceSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceSpecRequest) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *ModifyInstanceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceSpecRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceSpecRequest) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *ModifyInstanceSpecRequest) GetSlaveReadOnlyCount() *int32 {
	return s.SlaveReadOnlyCount
}

func (s *ModifyInstanceSpecRequest) GetSlaveReplicaCount() *int32 {
	return s.SlaveReplicaCount
}

func (s *ModifyInstanceSpecRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ModifyInstanceSpecRequest) GetStorage() *int32 {
	return s.Storage
}

func (s *ModifyInstanceSpecRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ModifyInstanceSpecRequest) SetAutoPay(v bool) *ModifyInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetBusinessInfo(v string) *ModifyInstanceSpecRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetClientToken(v string) *ModifyInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetCouponNo(v string) *ModifyInstanceSpecRequest {
	s.CouponNo = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetEffectiveTime(v string) *ModifyInstanceSpecRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetForceTrans(v bool) *ModifyInstanceSpecRequest {
	s.ForceTrans = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetForceUpgrade(v bool) *ModifyInstanceSpecRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceClass(v string) *ModifyInstanceSpecRequest {
	s.InstanceClass = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceId(v string) *ModifyInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetMajorVersion(v string) *ModifyInstanceSpecRequest {
	s.MajorVersion = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetNodeType(v string) *ModifyInstanceSpecRequest {
	s.NodeType = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetOrderType(v string) *ModifyInstanceSpecRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetOwnerAccount(v string) *ModifyInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetOwnerId(v int64) *ModifyInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetReadOnlyCount(v int32) *ModifyInstanceSpecRequest {
	s.ReadOnlyCount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetRegionId(v string) *ModifyInstanceSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetReplicaCount(v int32) *ModifyInstanceSpecRequest {
	s.ReplicaCount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetSecurityToken(v string) *ModifyInstanceSpecRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetShardCount(v int32) *ModifyInstanceSpecRequest {
	s.ShardCount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetSlaveReadOnlyCount(v int32) *ModifyInstanceSpecRequest {
	s.SlaveReadOnlyCount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetSlaveReplicaCount(v int32) *ModifyInstanceSpecRequest {
	s.SlaveReplicaCount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetSourceBiz(v string) *ModifyInstanceSpecRequest {
	s.SourceBiz = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetStorage(v int32) *ModifyInstanceSpecRequest {
	s.Storage = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetStorageType(v string) *ModifyInstanceSpecRequest {
	s.StorageType = &v
	return s
}

func (s *ModifyInstanceSpecRequest) Validate() error {
	return dara.Validate(s)
}
