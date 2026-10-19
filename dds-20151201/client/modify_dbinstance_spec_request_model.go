// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlignPrimaryZone(v bool) *ModifyDBInstanceSpecRequest
	GetAlignPrimaryZone() *bool
	SetAutoPay(v bool) *ModifyDBInstanceSpecRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *ModifyDBInstanceSpecRequest
	GetBusinessInfo() *string
	SetCouponNo(v string) *ModifyDBInstanceSpecRequest
	GetCouponNo() *string
	SetDBInstanceClass(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceStorage() *string
	SetEffectiveTime(v string) *ModifyDBInstanceSpecRequest
	GetEffectiveTime() *string
	SetExtraParam(v string) *ModifyDBInstanceSpecRequest
	GetExtraParam() *string
	SetOrderType(v string) *ModifyDBInstanceSpecRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *ModifyDBInstanceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceSpecRequest
	GetOwnerId() *int64
	SetReadonlyReplicas(v string) *ModifyDBInstanceSpecRequest
	GetReadonlyReplicas() *string
	SetReplicationFactor(v string) *ModifyDBInstanceSpecRequest
	GetReplicationFactor() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceSpecRequest
	GetResourceOwnerId() *int64
	SetSearchNodeClass(v string) *ModifyDBInstanceSpecRequest
	GetSearchNodeClass() *string
	SetSearchNodeCount(v int64) *ModifyDBInstanceSpecRequest
	GetSearchNodeCount() *int64
	SetSearchNodeStorage(v int64) *ModifyDBInstanceSpecRequest
	GetSearchNodeStorage() *int64
	SetTargetHiddenZoneId(v string) *ModifyDBInstanceSpecRequest
	GetTargetHiddenZoneId() *string
	SetTargetSecondaryZoneId(v string) *ModifyDBInstanceSpecRequest
	GetTargetSecondaryZoneId() *string
	SetTargetVswitchId(v string) *ModifyDBInstanceSpecRequest
	GetTargetVswitchId() *string
	SetTargetZoneId(v string) *ModifyDBInstanceSpecRequest
	GetTargetZoneId() *string
}

type ModifyDBInstanceSpecRequest struct {
	// Specifies whether to automatically align the primary node to the primary zone for a multi-zone instance after operations such as configuration changes. If you choose not to align the node, one primary/secondary switchover (transient disconnection) is avoided. Default value: true (aligned).
	AlignPrimaryZone *bool `json:"AlignPrimaryZone,omitempty" xml:"AlignPrimaryZone,omitempty"`
	// Specifies whether to enable automatic payment for the instance. Valid values:
	//
	// - **true**: enables automatic payment. This is the default value. Make sure that your account has sufficient balance.
	//
	// <props="china">- **false**: disables automatic payment. You must manually pay for the order in the console. Log on to the ApsaraDB for MongoDB console. In the upper-right corner of the page, choose **Expenses*	- > **Expenses and Costs**. In the left-side navigation pane, choose **Orders*	- > **My Orders**. On the **Orders*	- tab, find the order and complete the payment.
	//
	// <props="intl">- **false**: disables automatic payment. You must manually pay for the order in the console. Log on to the ApsaraDB for MongoDB console. In the upper-right corner of the page, choose **Expenses*	- > **Expenses and Costs**. In the left-side navigation pane, click **Order Management**. On the **Orders*	- page, find the order and complete the payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The business information.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
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
	// The instance type. <props="intl"><ph>For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html). You can also call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to query instance types.</ph>
	//
	// <props="china">
	//
	// - For standalone instances and replica set instances, this parameter specifies the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html). You can also call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to query the instance types of standalone instances and replica set instances.
	//
	// - For serverless instances, this parameter specifies the compute specifications of the instance. Valid values: 100 to 8000.
	//
	//
	// > You must specify at least one of the **DBInstanceClass*	- and **DBInstanceStorage*	- parameters.
	//
	// example:
	//
	// dds.sn4.xlarge.1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1ea17b41ab****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Instance storage space of the instance. <props="intl"><ph>Valid values: 10 to 3000. Unit: GB. The step size is 10 GB. The valid values are subject to the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).</ph>
	//
	// <props="china">
	//
	// - For standalone instances and replica set instances, valid values are 10 to 3000. Unit: GB. The step size is 10 GB. The valid values are subject to the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// - For serverless instances, valid values are 1 to 100. Unit: GB. The step size is 1 GB.
	//
	//
	// > - You must specify at least one of the **DBInstanceStorage*	- and **DBInstanceClass*	- parameters.
	//
	// > - You cannot downgrade instance storage space.
	//
	// example:
	//
	// 50
	DBInstanceStorage *string `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The time when the configuration change takes effect. Valid values:
	//
	// - **Immediately**: takes effect immediately. This is the default value.
	//
	// - **MaintainTime**: takes effect during the maintenance window of the instance.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The additional parameter. Valid values:
	//
	// - **async**: returns the result after the configuration change order is created.
	//
	// - **sync**: returns the result when the configuration change is delivered after the order is created.
	//
	// example:
	//
	// async
	ExtraParam *string `json:"ExtraParam,omitempty" xml:"ExtraParam,omitempty"`
	// The type of configuration change. Valid values:
	//
	// - **UPGRADE**: upgrades the configurations. This is the default value.
	//
	// - **DOWNGRADE**: downgrades the configurations.
	//
	// > You can specify this parameter only when the instance uses the subscription billing method.
	//
	// example:
	//
	// UPGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read-only nodes. Valid values: **0*	- to **5**.
	//
	// If the instance is connected only over the **classic network*	- and a **VPC**, you must enable Internet access or release the classic network before you can change the **number of read-only nodes**.
	//
	// > You can log on to the ApsaraDB for MongoDB console and view the enabled network types on the **Database Connections*	- page.
	//
	// example:
	//
	// 1
	ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The number of nodes in the instance. Default value: 3.
	//
	// - For a replica set instance, valid values are **3**, **5**, and **7**.
	//
	// - For a standalone instance, the value is fixed to **1**.
	//
	// - For a replica set instance that uses shared storage (supported only on the Alibaba Cloud China site), the value is fixed to **2**.
	//
	// > You do not need to specify this parameter for a serverless instance (supported only on the Alibaba Cloud China site).
	//
	// example:
	//
	// 3
	ReplicationFactor    *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The specifications of the search node for the configuration change.
	//
	// example:
	//
	// mdb.shard.2x.xlarge.d
	SearchNodeClass *string `json:"SearchNodeClass,omitempty" xml:"SearchNodeClass,omitempty"`
	// The number of search nodes for the configuration change.
	//
	// example:
	//
	// 2
	SearchNodeCount *int64 `json:"SearchNodeCount,omitempty" xml:"SearchNodeCount,omitempty"`
	// The storage capacity of the search node for the configuration change.
	//
	// example:
	//
	// 20
	SearchNodeStorage *int64 `json:"SearchNodeStorage,omitempty" xml:"SearchNodeStorage,omitempty"`
	// The destination zone for the hidden node when you cloud disk instance across zones and change its configurations at the same time.
	//
	// 	Notice: This parameter is supported only for instances that use cloud disks.</notice>
	//
	// 	Notice: The value of this parameter cannot be the same as the value of the TargetZoneId or TargetSecondaryZoneId parameter.</notice>
	//
	// >
	//
	// > - You must specify this parameter only when you cloud disk instance across zones and change its configurations at the same time.
	//
	// > - This parameter is available only for multi-zone migration.
	//
	// > - The destination zone and the current zone of the instance must be in the same region.
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-i
	TargetHiddenZoneId *string `json:"TargetHiddenZoneId,omitempty" xml:"TargetHiddenZoneId,omitempty"`
	// The destination zone for the secondary node when you cloud disk instance across zones and change its configurations at the same time.
	//
	// 	Notice: This parameter is supported only for instances that use cloud disks.</notice>
	//
	// 	Notice: The value of this parameter cannot be the same as the value of the TargetZoneId or TargetHiddenZoneId parameter.</notice>
	//
	// >
	//
	// > - You must specify this parameter only when you cloud disk instance across zones and change its configurations at the same time.
	//
	// > - This parameter is available only for multi-zone migration.
	//
	// > - The destination zone and the current zone of the instance must be in the same region.
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-h
	TargetSecondaryZoneId *string `json:"TargetSecondaryZoneId,omitempty" xml:"TargetSecondaryZoneId,omitempty"`
	// The ID of the destination vSwitch when you cloud disk instance across zones and change its configurations at the same time.
	//
	// 	Notice: This parameter is supported only for instances that use cloud disks.</notice>
	//
	// >
	//
	// > - You must specify this parameter only when you cloud disk instance across zones and change its configurations at the same time.
	//
	// example:
	//
	// vsw-bp1buy0h9myt5i9e7****
	TargetVswitchId *string `json:"TargetVswitchId,omitempty" xml:"TargetVswitchId,omitempty"`
	// The destination zone for instance migration when you migrate the instance across zones and change its configurations at the same time.
	//
	// 	Notice: This parameter is supported only for cloud disk instances.</notice>
	//
	// 	Notice: The value of this parameter cannot be the same as the value of the TargetSecondaryZoneId or TargetHiddenZoneId parameter.</notice>
	//
	// > - You must specify this parameter only when you migrate the instance across zones and change its configurations at the same time.
	//
	// > - The destination zone and the current zone of the instance must be in the same region.
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-j
	TargetZoneId *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
}

func (s ModifyDBInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecRequest) GetAlignPrimaryZone() *bool {
	return s.AlignPrimaryZone
}

func (s *ModifyDBInstanceSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDBInstanceSpecRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *ModifyDBInstanceSpecRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceStorage() *string {
	return s.DBInstanceStorage
}

func (s *ModifyDBInstanceSpecRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBInstanceSpecRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *ModifyDBInstanceSpecRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ModifyDBInstanceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceSpecRequest) GetReadonlyReplicas() *string {
	return s.ReadonlyReplicas
}

func (s *ModifyDBInstanceSpecRequest) GetReplicationFactor() *string {
	return s.ReplicationFactor
}

func (s *ModifyDBInstanceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceSpecRequest) GetSearchNodeClass() *string {
	return s.SearchNodeClass
}

func (s *ModifyDBInstanceSpecRequest) GetSearchNodeCount() *int64 {
	return s.SearchNodeCount
}

func (s *ModifyDBInstanceSpecRequest) GetSearchNodeStorage() *int64 {
	return s.SearchNodeStorage
}

func (s *ModifyDBInstanceSpecRequest) GetTargetHiddenZoneId() *string {
	return s.TargetHiddenZoneId
}

func (s *ModifyDBInstanceSpecRequest) GetTargetSecondaryZoneId() *string {
	return s.TargetSecondaryZoneId
}

func (s *ModifyDBInstanceSpecRequest) GetTargetVswitchId() *string {
	return s.TargetVswitchId
}

func (s *ModifyDBInstanceSpecRequest) GetTargetZoneId() *string {
	return s.TargetZoneId
}

func (s *ModifyDBInstanceSpecRequest) SetAlignPrimaryZone(v bool) *ModifyDBInstanceSpecRequest {
	s.AlignPrimaryZone = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetAutoPay(v bool) *ModifyDBInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetBusinessInfo(v string) *ModifyDBInstanceSpecRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetCouponNo(v string) *ModifyDBInstanceSpecRequest {
	s.CouponNo = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceClass(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceId(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceStorage(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetEffectiveTime(v string) *ModifyDBInstanceSpecRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetExtraParam(v string) *ModifyDBInstanceSpecRequest {
	s.ExtraParam = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOrderType(v string) *ModifyDBInstanceSpecRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOwnerAccount(v string) *ModifyDBInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOwnerId(v int64) *ModifyDBInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetReadonlyReplicas(v string) *ModifyDBInstanceSpecRequest {
	s.ReadonlyReplicas = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetReplicationFactor(v string) *ModifyDBInstanceSpecRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetSearchNodeClass(v string) *ModifyDBInstanceSpecRequest {
	s.SearchNodeClass = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetSearchNodeCount(v int64) *ModifyDBInstanceSpecRequest {
	s.SearchNodeCount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetSearchNodeStorage(v int64) *ModifyDBInstanceSpecRequest {
	s.SearchNodeStorage = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetTargetHiddenZoneId(v string) *ModifyDBInstanceSpecRequest {
	s.TargetHiddenZoneId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetTargetSecondaryZoneId(v string) *ModifyDBInstanceSpecRequest {
	s.TargetSecondaryZoneId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetTargetVswitchId(v string) *ModifyDBInstanceSpecRequest {
	s.TargetVswitchId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetTargetZoneId(v string) *ModifyDBInstanceSpecRequest {
	s.TargetZoneId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) Validate() error {
	return dara.Validate(s)
}
