// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeSpecBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyNodeSpecBatchRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *ModifyNodeSpecBatchRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *ModifyNodeSpecBatchRequest
	GetClientToken() *string
	SetCouponNo(v string) *ModifyNodeSpecBatchRequest
	GetCouponNo() *string
	SetDBInstanceId(v string) *ModifyNodeSpecBatchRequest
	GetDBInstanceId() *string
	SetEffectiveTime(v string) *ModifyNodeSpecBatchRequest
	GetEffectiveTime() *string
	SetNodesInfo(v string) *ModifyNodeSpecBatchRequest
	GetNodesInfo() *string
	SetOrderType(v string) *ModifyNodeSpecBatchRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *ModifyNodeSpecBatchRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNodeSpecBatchRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyNodeSpecBatchRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNodeSpecBatchRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNodeSpecBatchRequest
	GetResourceOwnerId() *int64
	SetTargetHiddenZoneId(v string) *ModifyNodeSpecBatchRequest
	GetTargetHiddenZoneId() *string
	SetTargetSecondaryZoneId(v string) *ModifyNodeSpecBatchRequest
	GetTargetSecondaryZoneId() *string
	SetTargetVswitchId(v string) *ModifyNodeSpecBatchRequest
	GetTargetVswitchId() *string
	SetTargetZoneId(v string) *ModifyNodeSpecBatchRequest
	GetTargetZoneId() *string
}

type ModifyNodeSpecBatchRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: Automatic payment is enabled. Make sure that your account has a sufficient balance.
	//
	// <props="china">
	//
	// - **false**: Manual payment is enabled. Log on to the ApsaraDB for MongoDB console. In the upper-right corner, choose **Expenses*	- > **Expenses and Costs**. In the navigation pane on the left, choose **Subscription Orders*	- > **My Orders**. On the **Product Orders*	- tab, find the target order and pay for it.
	//
	//
	//
	//
	// <props="intl">
	//
	// - **false**: Manual payment is enabled. Log on to the ApsaraDB for MongoDB console. In the upper-right corner, choose **Expenses*	- > **Expenses and Costs**. In the navigation pane on the left, click **Order Management**. On the **Product Orders*	- page, find the target order and pay for it.
	//
	//
	//
	//
	// Default value: **true**.
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
	// A client token. It is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **default*	- or **null*	- (default): A coupon is used.
	//
	// - **youhuiquan_promotion_option_id_for_blank**: A coupon is not used.
	//
	// example:
	//
	// default
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The ID of the instance for which you want to change configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1337621e8f****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The effective period of the configuration change. Valid values:
	//
	// - **Immediately**: The change takes effect immediately.
	//
	// - **MaintainTime**: The change takes effect during the maintenance window of the instance.
	//
	// > 	- You can call the [ModifyDBInstanceMaintainTime](https://help.aliyun.com/document_detail/62008.html) operation to change the maintenance window of an instance.
	//
	// >
	//
	// > 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to view the maintenance window of an instance.
	//
	// Default value: **Immediately**.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The instance types of the Mongos and shard nodes that you want to change. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Shards":[{"DBInstanceClass":"dds.shard.mid","DBInstanceName":"d-bp14ae4572fd****","Storage":20},{"DBInstanceClass":"dds.shard.mid","DBInstanceName":"d-bp19f4f92dc5****","Storage":30}]}
	NodesInfo *string `json:"NodesInfo,omitempty" xml:"NodesInfo,omitempty"`
	// The type of configuration change. Valid values:
	//
	// - **UPGRADE**: Upgrades the instance configuration.
	//
	// - **DOWNGRADE**: Downgrades the instance configuration.
	//
	// > This parameter is available only when the billing method of the instance is subscription.
	//
	// example:
	//
	// UPGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination zone of the hidden node for a configuration change that involves a zone migration.
	//
	// 	Notice:
	//
	// This parameter applies only to instances that use disks.
	//
	//
	//
	// 	Notice:
	//
	// The value of this parameter cannot be the same as the value of the TargetZoneId or TargetSecondaryZoneId parameter.
	//
	//
	//
	// > - This parameter is required only for a configuration change that involves a zone migration.
	//
	// example:
	//
	// cn-hangzhou-e
	TargetHiddenZoneId *string `json:"TargetHiddenZoneId,omitempty" xml:"TargetHiddenZoneId,omitempty"`
	// The destination zone of the secondary node for a configuration change that involves a zone migration.
	//
	// 	Notice:
	//
	// This parameter applies only to instances that use disks.
	//
	//
	//
	// 	Notice:
	//
	// The value of this parameter cannot be the same as the value of the TargetZoneId or TargetHiddenZoneId parameter.
	//
	//
	//
	// > - This parameter is required only for a configuration change that involves a zone migration.
	//
	// example:
	//
	// cn-hangzhou-j
	TargetSecondaryZoneId *string `json:"TargetSecondaryZoneId,omitempty" xml:"TargetSecondaryZoneId,omitempty"`
	// The ID of the destination virtual switch for a configuration change that involves a zone migration.
	//
	// 	Notice:
	//
	// This parameter applies only to instances that use disks.
	//
	//
	//
	// > - This parameter is required only for a configuration change that involves a zone migration.
	//
	// example:
	//
	// vsw-xxxxxxxx
	TargetVswitchId *string `json:"TargetVswitchId,omitempty" xml:"TargetVswitchId,omitempty"`
	// The destination zone of the primary node for a configuration change that involves a zone migration.
	//
	// 	Notice:
	//
	// This parameter applies only to instances that use disks.
	//
	//
	//
	// 	Notice:
	//
	// The value of this parameter cannot be the same as the value of the TargetSecondaryZoneId or TargetHiddenZoneId parameter.
	//
	//
	//
	// > - This parameter is required only for a configuration change that involves a zone migration.
	//
	// example:
	//
	// cn-hangzhou-h
	TargetZoneId *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
}

func (s ModifyNodeSpecBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeSpecBatchRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecBatchRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyNodeSpecBatchRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *ModifyNodeSpecBatchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNodeSpecBatchRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *ModifyNodeSpecBatchRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyNodeSpecBatchRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyNodeSpecBatchRequest) GetNodesInfo() *string {
	return s.NodesInfo
}

func (s *ModifyNodeSpecBatchRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ModifyNodeSpecBatchRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNodeSpecBatchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNodeSpecBatchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNodeSpecBatchRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNodeSpecBatchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNodeSpecBatchRequest) GetTargetHiddenZoneId() *string {
	return s.TargetHiddenZoneId
}

func (s *ModifyNodeSpecBatchRequest) GetTargetSecondaryZoneId() *string {
	return s.TargetSecondaryZoneId
}

func (s *ModifyNodeSpecBatchRequest) GetTargetVswitchId() *string {
	return s.TargetVswitchId
}

func (s *ModifyNodeSpecBatchRequest) GetTargetZoneId() *string {
	return s.TargetZoneId
}

func (s *ModifyNodeSpecBatchRequest) SetAutoPay(v bool) *ModifyNodeSpecBatchRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetBusinessInfo(v string) *ModifyNodeSpecBatchRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetClientToken(v string) *ModifyNodeSpecBatchRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetCouponNo(v string) *ModifyNodeSpecBatchRequest {
	s.CouponNo = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetDBInstanceId(v string) *ModifyNodeSpecBatchRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetEffectiveTime(v string) *ModifyNodeSpecBatchRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetNodesInfo(v string) *ModifyNodeSpecBatchRequest {
	s.NodesInfo = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetOrderType(v string) *ModifyNodeSpecBatchRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetOwnerAccount(v string) *ModifyNodeSpecBatchRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetOwnerId(v int64) *ModifyNodeSpecBatchRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetRegionId(v string) *ModifyNodeSpecBatchRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetResourceOwnerAccount(v string) *ModifyNodeSpecBatchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetResourceOwnerId(v int64) *ModifyNodeSpecBatchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetTargetHiddenZoneId(v string) *ModifyNodeSpecBatchRequest {
	s.TargetHiddenZoneId = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetTargetSecondaryZoneId(v string) *ModifyNodeSpecBatchRequest {
	s.TargetSecondaryZoneId = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetTargetVswitchId(v string) *ModifyNodeSpecBatchRequest {
	s.TargetVswitchId = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) SetTargetZoneId(v string) *ModifyNodeSpecBatchRequest {
	s.TargetZoneId = &v
	return s
}

func (s *ModifyNodeSpecBatchRequest) Validate() error {
	return dara.Validate(s)
}
