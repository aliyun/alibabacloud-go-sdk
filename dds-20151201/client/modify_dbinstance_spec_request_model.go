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
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true*	- (default): enables automatic payment. Make sure that your Alibaba Cloud account has a sufficient balance.
	//
	// 	- **false**: disables automatic payment. You can perform the following operations to pay for the instance: Log on to the ApsaraDB for MongoDB console. In the upper-right corner of the page, choose **Expenses*	- > **User Center**. In the left-side navigation pane, choose **Order Management*	- > **Order**. On the **Orders for Services*	- tab, find the order and pay for the order.
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
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html). You can also call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation to view instance types.
	//
	// > You must specify at least one of the DBInstanceClass and **DBInstanceStorage*	- parameters.
	//
	// example:
	//
	// dds.sn4.xlarge.1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1ea17b41ab****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The storage capacity of the instance. Valid values: 10 to 3000. The value must be a multiple of 10. Unit: GB. The values that can be specified for this parameter are subject to the instance types. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// >
	//
	// 	- You must specify at least one of the DBInstanceStorage and **DBInstanceClass*	- parameters.
	//
	// 	- Storage capacity can be scaled down only for pay-as-you-go replica set instances. The new storage capacity you specify must be greater than the used storage capacity.
	//
	// example:
	//
	// 50
	DBInstanceStorage *string `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The time when the changed configurations take effect. Valid values:
	//
	// 	- **Immediately*	- (default): The configurations immediately take effect.
	//
	// 	- **MaintainTime**: The configurations take effect during the maintenance window of the instance.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The additional parameter.
	//
	// Valid values:
	//
	// 	- async
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- sync
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// async
	ExtraParam *string `json:"ExtraParam,omitempty" xml:"ExtraParam,omitempty"`
	// The type of the configuration change. Valid values:
	//
	// 	- **UPGRADE**
	//
	// 	- **DOWNGRADE*	- (default)
	//
	// >  This parameter can be configured only when the billing method of the instance is subscription.
	//
	// example:
	//
	// UPGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read-only nodes. Valid values: **0*	- to **5**.
	//
	// If your instance has only **Classic Network*	- and **VPC*	- endpoints, you must apply for a public endpoint or release the classic network endpoint for the instance before you can change the **Read-only Nodes*	- value.
	//
	// > You can go to the **Database Connections*	- page to view the types of networks that are enabled.
	//
	// example:
	//
	// 1
	ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	// The number of nodes in the instance.
	//
	// 	- Valid values for replica set instances: **3**, **5**, and **7**
	//
	// 	- Valid values for standalone instances: **1**
	//
	// >  This parameter is not required for a serverless instance which is only available on the China site (aliyun.com).
	//
	// example:
	//
	// 3
	ReplicationFactor    *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// mdb.shard.2x.xlarge.d
	SearchNodeClass *string `json:"SearchNodeClass,omitempty" xml:"SearchNodeClass,omitempty"`
	// example:
	//
	// 2
	SearchNodeCount *int64 `json:"SearchNodeCount,omitempty" xml:"SearchNodeCount,omitempty"`
	// example:
	//
	// 20
	SearchNodeStorage     *int64  `json:"SearchNodeStorage,omitempty" xml:"SearchNodeStorage,omitempty"`
	TargetHiddenZoneId    *string `json:"TargetHiddenZoneId,omitempty" xml:"TargetHiddenZoneId,omitempty"`
	TargetSecondaryZoneId *string `json:"TargetSecondaryZoneId,omitempty" xml:"TargetSecondaryZoneId,omitempty"`
	TargetVswitchId       *string `json:"TargetVswitchId,omitempty" xml:"TargetVswitchId,omitempty"`
	TargetZoneId          *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
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
