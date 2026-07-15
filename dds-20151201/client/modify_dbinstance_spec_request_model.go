// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// Specifies whether to enable automatic payment for the instance. Valid values:
	//
	// - **true**: enables automatic payment. Make sure that your account has a sufficient balance. This is the default value.
	//
	// <props="china">
	//
	// - **false**: disables automatic payment. You can log on to the ApsaraDB for MongoDB console to pay for the instance. In the upper-right corner of the page, choose **Billing Management*	- > **Billing Management**. In the left-side navigation pane, choose **Orders*	- > **My Orders**. On the **Product Orders*	- tab, find the order and complete the payment.
	//
	//
	//
	//
	// <props="intl">
	//
	// - **false**: disables automatic payment. You can log on to the ApsaraDB for MongoDB console to pay for the instance. In the upper-right corner of the page, choose **Billing Management*	- > **Billing Management**. In the left-side navigation pane, click **Orders**. On the **Product Orders*	- page, find the order and complete the payment.
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
	// - **default*	- or **null*	- (default): A coupon is used.
	//
	// - **youhuiquan_promotion_option_id_for_blank**: No coupon is used.
	//
	// example:
	//
	// default
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance type. <props="intl">For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html). You can also call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to query instance types.<props="china">
	//
	// - For a standalone instance or a replica set instance, this parameter specifies the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html). You can also call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to query the instance types of standalone and replica set instances.
	//
	// - For a serverless instance, this parameter specifies the computing capacity of the instance. Valid values: 100 to 8000.
	//
	// > You must configure one of the **DBInstanceStorage*	- and DBInstanceClass parameters.
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
	// The storage capacity of the instance. <props="intl">The value must be an integer that is greater than or equal to 10. The value increases in increments of 10. Unit: GB. The values that can be specified for this parameter are subject to the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// <props="china">
	//
	// - The storage capacity of a standalone instance or a replica set instance must be a multiple of 10. The valid values are 10 to 3000. Unit: GB. The values that can be specified for this parameter are subject to the instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// - The storage capacity of a serverless instance must be a multiple of 1. The valid values are 1 to 100. Unit: GB.
	//
	//
	//
	// > - You must configure one of the **DBInstanceClass*	- and DBInstanceStorage parameters.
	//
	// >
	//
	// > - You cannot decrease the storage capacity of an instance.
	//
	// example:
	//
	// 50
	DBInstanceStorage *string `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The effective time of the configuration change. Valid values:
	//
	// - **Immediately**: The configuration change immediately takes effect. This is the default value.
	//
	// - **MaintainTime**: The configuration change takes effect during the maintenance window of the instance.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// Additional parameters. Valid values:
	//
	// - **async**: The result is returned after the specification change order is created.
	//
	// - **sync**: The result is returned after the instance specification change is delivered.
	//
	// example:
	//
	// async
	ExtraParam *string `json:"ExtraParam,omitempty" xml:"ExtraParam,omitempty"`
	// The specification change type. Valid values:
	//
	// - **UPGRADE**: upgrades the specifications. This is the default value.
	//
	// - **DOWNGRADE**: downgrades the specifications.
	//
	// > This parameter is available only for subscription instances.
	//
	// example:
	//
	// UPGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read-only nodes. Valid values: **0*	- to **5**.
	//
	// If the network type of the instance is set to only **classic network*	- and **VPC**, you need to enable public access or release the classic network endpoint before you can change the **number of read-only nodes**.
	//
	// > You can log on to the ApsaraDB for MongoDB console and go to the **Database Connections*	- page to view the network types that have been enabled.
	//
	// example:
	//
	// 1
	ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The number of nodes in the instance. Default value: 3.
	//
	// - Valid values for replica set instances: **3**, **5**, and **7**.
	//
	// - The value for standalone instances is fixed at **1**.
	//
	// - The value for replica set instances with shared storage (available only in the China site) is fixed at **2**.
	//
	// > This parameter is not required for serverless instances (available only in the China site).
	//
	// example:
	//
	// 3
	ReplicationFactor    *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The specifications of the Search node to be changed.
	//
	// example:
	//
	// mdb.shard.2x.xlarge.d
	SearchNodeClass *string `json:"SearchNodeClass,omitempty" xml:"SearchNodeClass,omitempty"`
	// The number of Search nodes to be changed.
	//
	// example:
	//
	// 2
	SearchNodeCount *int64 `json:"SearchNodeCount,omitempty" xml:"SearchNodeCount,omitempty"`
	// The capacity of the Search node to be changed.
	//
	// example:
	//
	// 20
	SearchNodeStorage *int64 `json:"SearchNodeStorage,omitempty" xml:"SearchNodeStorage,omitempty"`
	// The destination zone for the hidden node when you change the specifications and migrate the instance across zones.
	//
	// 	Notice: This parameter applies only to cloud disk instances.
	//
	// 	Notice: The value of this parameter cannot be the same as the value of the TargetZoneId or TargetSecondaryZoneId parameter.
	//
	// > - You must specify this parameter only when you change the specifications and migrate the instance across zones.
	//
	// >
	//
	// > - This parameter is available only for multi-zone migration.
	//
	// >
	//
	// > - The destination zone and the current zone must be in the same region.
	//
	// >
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-i
	TargetHiddenZoneId *string `json:"TargetHiddenZoneId,omitempty" xml:"TargetHiddenZoneId,omitempty"`
	// The destination secondary zone for the secondary node when you change the specifications and migrate the instance across zones.
	//
	// 	Notice: This parameter applies only to cloud disk instances.
	//
	// 	Notice: The value of this parameter cannot be the same as the value of the TargetZoneId or TargetHiddenZoneId parameter.
	//
	// > - You must specify this parameter only when you change the specifications and migrate the instance across zones.
	//
	// >
	//
	// > - This parameter is available only for multi-zone migration.
	//
	// >
	//
	// > - The destination zone and the current zone must be in the same region.
	//
	// >
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-h
	TargetSecondaryZoneId *string `json:"TargetSecondaryZoneId,omitempty" xml:"TargetSecondaryZoneId,omitempty"`
	// The destination vSwitch ID when you change the specifications and migrate the instance across zones.
	//
	// 	Notice: This parameter applies only to cloud disk instances.
	//
	// > - You must specify this parameter only when you change the specifications and migrate the instance across zones.
	//
	// example:
	//
	// vsw-bp1buy0h9myt5i9e7****
	TargetVswitchId *string `json:"TargetVswitchId,omitempty" xml:"TargetVswitchId,omitempty"`
	// The destination zone to which you want to migrate the instance when you change the specifications and migrate the instance across zones.
	//
	// 	Notice: This parameter applies only to cloud disk instances.
	//
	// 	Notice: The value of this parameter cannot be the same as the value of the TargetSecondaryZoneId or TargetHiddenZoneId parameter.
	//
	// > - You must specify this parameter only when you change the specifications and migrate the instance across zones.
	//
	// >
	//
	// > - The destination zone and the current zone must be in the same region.
	//
	// >
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
