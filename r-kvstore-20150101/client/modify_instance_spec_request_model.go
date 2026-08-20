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
	SetSecondaryZoneId(v string) *ModifyInstanceSpecRequest
	GetSecondaryZoneId() *string
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
	// 	- **true**: Automatic payment is enabled. This is the default value.
	//
	// 	- **false**: Automatic payment is disabled. If you set this parameter to **false**, you must manually renew the instance before the instance expires in the console. For details, see [Manual renewal](https://help.aliyun.com/document_detail/26352.html).
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The activity ID and business information.
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
	// The time when the specification change takes effect. Valid values:
	//
	// 	- **Immediately**: The specification change takes effect immediately. This is the default value.
	//
	// 	- **MaintainTime**: The specification change takes effect during the maintenance window of the instance. You can call [ModifyInstanceMaintainTime](https://help.aliyun.com/document_detail/473775.html) to modify the maintenance window.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// Specifies whether to enable forced transmission. Valid values:
	//
	// - **false*	- (default): Before the specification change, the system checks the current minor engine version of the instance. If the minor engine version is too old, an error is returned. You must upgrade the minor engine version and retry.
	//
	// - **true**: Skips the check and directly executes the specification change operation.
	//
	// example:
	//
	// false
	ForceTrans *bool `json:"ForceTrans,omitempty" xml:"ForceTrans,omitempty"`
	// Specifies whether to forcibly change the specifications. Valid values:
	//
	// 	- **false**: does not forcibly change the specifications.
	//
	// 	- **true**: forcibly changes the specifications. This is the default value.
	//
	// example:
	//
	// true
	ForceUpgrade *bool `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	// The new instance type. You can call [DescribeAvailableResource](https://help.aliyun.com/document_detail/473765.html) to query the instance types available for specification changes in the zone where the instance resides.
	//
	// > For more information about instance types, see [Instance type navigation](https://help.aliyun.com/document_detail/26350.html).
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID. You can call [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The major engine version for upgrading a classic instance. Valid values: **2.8**, **4.0**, and **5.0**.
	//
	// > When you upgrade the version, the **InstanceClass*	- parameter is required. This indicates that this operation supports version upgrades only when the instance specifications are also changed. To upgrade only the instance version, call [ModifyInstanceMajorVersion](https://help.aliyun.com/document_detail/473776.html).
	//
	// example:
	//
	// 5.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
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
	// > For cloud-native instances, select **MASTER_SLAVE*	- or **STAND_ALONE**. For classic instances, select **double*	- or **single**.
	//
	// example:
	//
	// MASTER_SLAVE
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The type of specification change. This parameter is required when you change the specifications of a subscription instance. Valid values:
	//
	// 	- **UPGRADE**: Upgrade. This is the default value.
	//
	// 	- **DOWNGRADE**: Downgrade.
	//
	// > 	- You must set this parameter to **DOWNGRADE*	- when you downgrade a subscription instance.
	//
	// > 	- If the price of the target instance type is higher than that of the current instance type, the change is an upgrade. Otherwise, the change is a downgrade. For example, the price of the read/write splitting 8 GB edition (5 read-only nodes) is higher than that of the 16 GB cluster edition. Changing from the latter to the former is an upgrade.
	//
	// example:
	//
	// DOWNGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read-only nodes in the primary zone. This parameter is applicable only to cloud-native read/write splitting instances.
	//
	// 	- For standard architecture instances, valid values are 0 to 9. A value of 0 indicates that read/write splitting is shutdown and the instance is switched to the standard architecture.
	//
	// 	- For cluster architecture instances, valid values are 1 to 4, which specifies the number of read-only nodes per data shard.
	//
	// > For multi-zone instances, you can use this parameter together with the SlaveReadOnlyCount parameter to specify the number of read-only nodes in the primary and secondary zones.
	//
	// > 	- For standard architecture instances, the sum of this parameter and SlaveReadOnlyCount cannot exceed 9.
	//
	// > 	- For cluster architecture instances, the sum of this parameter and SlaveReadOnlyCount cannot exceed 4.
	//
	// example:
	//
	// 5
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) to query available regions. Use this parameter to specify the region of the instance whose specifications you want to change.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of replica nodes in the primary zone. This parameter is applicable only to cloud-native cluster multi-replica instances. You can use this parameter to specify a custom number of replica nodes. Valid values: 1 to 4.
	//
	// > For multi-zone instances, you can use this parameter together with the SlaveReplicaCount parameter to specify the number of replica nodes in the primary and secondary zones. The sum of this parameter and the SlaveReplicaCount parameter cannot exceed 4.
	//
	// example:
	//
	// 1
	ReplicaCount         *int32  `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The secondary zone ID. This parameter is required when you change the specifications of a single-zone instance and migrate it to a multi-zone deployment. You can call [DescribeZones](https://help.aliyun.com/document_detail/473764.html) to query available zones.
	//
	// > The value of this parameter must be different from the value of the ZoneId parameter. Do not set this parameter to the ID of a multi-zone.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of shards. This parameter is applicable only to cloud-native cluster instances. You can use this parameter to specify a custom number of shards.
	//
	// >
	//
	// > - To change a cloud-native cluster instance to a standard architecture, set ShardCount to 1 and set the instance type to a standard instance type.
	//
	// > - To change a cloud-native standard instance to a cluster architecture, set ShardCount to a value greater than 1 and set the instance type to a cluster instance type.
	//
	// example:
	//
	// 8
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
	// 1
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// The source of the request. This parameter is used only for internal maintenance and does not need to be specified.
	//
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
	// The storage capacity of a cloud disk instance. The valid values vary based on the instance type. For more information, see [Cloud disk-based instance types](https://help.aliyun.com/document_detail/2527111.html).
	//
	// > This parameter is required only when **InstanceType*	- is set to **tair_essd*	- and you are creating a Tair ESSD-based cloud disk instance. For Tair SSD-based cloud disk instances, the storage capacity is a fixed value based on the instance type, and you do not need to specify this parameter.
	//
	// example:
	//
	// 60
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The storage type. Valid values: **essd_pl1**, **essd_pl2**, and **essd_pl3**.
	//
	// > This parameter is required only when **InstanceType*	- is set to **tair_essd*	- and the instance is a Tair ESSD-based cloud disk instance.
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

func (s *ModifyInstanceSpecRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
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

func (s *ModifyInstanceSpecRequest) SetSecondaryZoneId(v string) *ModifyInstanceSpecRequest {
	s.SecondaryZoneId = &v
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
