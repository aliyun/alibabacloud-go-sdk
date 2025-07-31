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
}

type ModifyNodeSpecBatchRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: enables automatic payment. Make sure that you have sufficient balance within your account.
	//
	// 	- **false**: disables automatic payment. You can perform the following operations to pay for the instance: Log on to the ApsaraDB for MongoDB console. In the upper-right corner of the page, click **Expenses*	- to go to the **Billing Management*	- console. In the left-side navigation pane, click **Orders**. On the **Orders*	- page, find the order and complete the payment.
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
	// The ID of the instance whose configurations you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1337621e8f****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The time when the changed configurations take effect. Valid values:
	//
	// 	- **Immediately**: The configurations immediately take effect.
	//
	// 	- **MaintainTime**: The configurations take effect during the maintenance window of the instance.
	//
	// >
	//
	// 	- You can call the [ModifyDBInstanceMaintainTime](https://help.aliyun.com/document_detail/62008.html) operation to modify the maintenance window of an instance.
	//
	// 	- You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to view the maintenance window of an instance.
	//
	// Default value: **Immediately**.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The configuration information of the mongos nodes or shard nodes whose configurations you want to change. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Shards":[{"DBInstanceClass":"dds.shard.mid","DBInstanceName":"d-bp14ae4572fd****","Storage":20},{"DBInstanceClass":"dds.shard.mid","DBInstanceName":"d-bp19f4f92dc5****","Storage":30}]}
	NodesInfo *string `json:"NodesInfo,omitempty" xml:"NodesInfo,omitempty"`
	// The type of configuration changes. Valid values:
	//
	// 	- **UPGRADE**
	//
	// 	- **DOWNGRADE**
	//
	// > This parameter is only applicable to instances whose billing method is subscription.
	//
	// example:
	//
	// UPGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *ModifyNodeSpecBatchRequest) Validate() error {
	return dara.Validate(s)
}
