// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyNodeSpecRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *ModifyNodeSpecRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *ModifyNodeSpecRequest
	GetClientToken() *string
	SetCouponNo(v string) *ModifyNodeSpecRequest
	GetCouponNo() *string
	SetDBInstanceId(v string) *ModifyNodeSpecRequest
	GetDBInstanceId() *string
	SetEffectiveTime(v string) *ModifyNodeSpecRequest
	GetEffectiveTime() *string
	SetFromApp(v string) *ModifyNodeSpecRequest
	GetFromApp() *string
	SetNodeClass(v string) *ModifyNodeSpecRequest
	GetNodeClass() *string
	SetNodeId(v string) *ModifyNodeSpecRequest
	GetNodeId() *string
	SetNodeStorage(v int32) *ModifyNodeSpecRequest
	GetNodeStorage() *int32
	SetOrderType(v string) *ModifyNodeSpecRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *ModifyNodeSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNodeSpecRequest
	GetOwnerId() *int64
	SetReadonlyReplicas(v int32) *ModifyNodeSpecRequest
	GetReadonlyReplicas() *int32
	SetResourceOwnerAccount(v string) *ModifyNodeSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNodeSpecRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyNodeSpecRequest
	GetSwitchTime() *string
	SetTargetHiddenZoneId(v string) *ModifyNodeSpecRequest
	GetTargetHiddenZoneId() *string
	SetTargetSecondaryZoneId(v string) *ModifyNodeSpecRequest
	GetTargetSecondaryZoneId() *string
	SetTargetVswitchId(v string) *ModifyNodeSpecRequest
	GetTargetVswitchId() *string
	SetTargetZoneId(v string) *ModifyNodeSpecRequest
	GetTargetZoneId() *string
}

type ModifyNodeSpecRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true*	- (default): enables automatic payment. Make sure that you have sufficient balance within your account.
	//
	// 	- **false**: disables automatic payment. In this case, you must manually pay for the instance.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The business information. This is an additional parameter.
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
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1c0b990184****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The time when the changed configurations take effect. Valid values:
	//
	// 	- **Immediately*	- (default): The new configurations immediately take effect
	//
	// 	- **MaintainTime**: The new configurations take effect during the maintenance window of the instance.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The source of the request. Valid values:
	//
	// 	- **OpenApi**: the ApsaraDB for MongoDB API
	//
	// 	- **mongo_buy**: the ApsaraDB for MongoDB console
	//
	// example:
	//
	// OpenApi
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The specifications of the shard or mongos node. For more information, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
	//
	// example:
	//
	// dds.mongos.standard
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The ID of the shard or mongos node in the sharded cluster instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the node ID.
	//
	// > If you set this parameter to the ID of the shard node, you must also specify the **NodeStorage*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp143e1b1637****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The storage capacity of the shard node. Unit: GB.
	//
	// 	- Valid values are **10*	- to **2000*	- if the instance uses local SSDs.
	//
	// 	- Valid values are **20*	- to **16000*	- if the instance uses enhanced SSDs (ESSDs) at PL1.
	//
	// > The value must be a multiple of 10.
	//
	// example:
	//
	// 20
	NodeStorage *int32 `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	// The order type. Valid values:
	//
	// 	- **UPGRADE**
	//
	// 	- **DOWNGRADE**
	//
	// example:
	//
	// UPGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read-only nodes in the shard node.
	//
	// Valid values: **0*	- to **5**. The value must be an integer. Default value: **0**.
	//
	// example:
	//
	// 5
	ReadonlyReplicas     *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The execution time. Specify the time in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-05T03:18:53Z
	SwitchTime            *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	TargetHiddenZoneId    *string `json:"TargetHiddenZoneId,omitempty" xml:"TargetHiddenZoneId,omitempty"`
	TargetSecondaryZoneId *string `json:"TargetSecondaryZoneId,omitempty" xml:"TargetSecondaryZoneId,omitempty"`
	TargetVswitchId       *string `json:"TargetVswitchId,omitempty" xml:"TargetVswitchId,omitempty"`
	TargetZoneId          *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
}

func (s ModifyNodeSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyNodeSpecRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *ModifyNodeSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNodeSpecRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *ModifyNodeSpecRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyNodeSpecRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyNodeSpecRequest) GetFromApp() *string {
	return s.FromApp
}

func (s *ModifyNodeSpecRequest) GetNodeClass() *string {
	return s.NodeClass
}

func (s *ModifyNodeSpecRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyNodeSpecRequest) GetNodeStorage() *int32 {
	return s.NodeStorage
}

func (s *ModifyNodeSpecRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ModifyNodeSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNodeSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNodeSpecRequest) GetReadonlyReplicas() *int32 {
	return s.ReadonlyReplicas
}

func (s *ModifyNodeSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNodeSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNodeSpecRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyNodeSpecRequest) GetTargetHiddenZoneId() *string {
	return s.TargetHiddenZoneId
}

func (s *ModifyNodeSpecRequest) GetTargetSecondaryZoneId() *string {
	return s.TargetSecondaryZoneId
}

func (s *ModifyNodeSpecRequest) GetTargetVswitchId() *string {
	return s.TargetVswitchId
}

func (s *ModifyNodeSpecRequest) GetTargetZoneId() *string {
	return s.TargetZoneId
}

func (s *ModifyNodeSpecRequest) SetAutoPay(v bool) *ModifyNodeSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetBusinessInfo(v string) *ModifyNodeSpecRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetClientToken(v string) *ModifyNodeSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetCouponNo(v string) *ModifyNodeSpecRequest {
	s.CouponNo = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetDBInstanceId(v string) *ModifyNodeSpecRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetEffectiveTime(v string) *ModifyNodeSpecRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetFromApp(v string) *ModifyNodeSpecRequest {
	s.FromApp = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetNodeClass(v string) *ModifyNodeSpecRequest {
	s.NodeClass = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetNodeId(v string) *ModifyNodeSpecRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetNodeStorage(v int32) *ModifyNodeSpecRequest {
	s.NodeStorage = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetOrderType(v string) *ModifyNodeSpecRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetOwnerAccount(v string) *ModifyNodeSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetOwnerId(v int64) *ModifyNodeSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetReadonlyReplicas(v int32) *ModifyNodeSpecRequest {
	s.ReadonlyReplicas = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetResourceOwnerAccount(v string) *ModifyNodeSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetResourceOwnerId(v int64) *ModifyNodeSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetSwitchTime(v string) *ModifyNodeSpecRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetTargetHiddenZoneId(v string) *ModifyNodeSpecRequest {
	s.TargetHiddenZoneId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetTargetSecondaryZoneId(v string) *ModifyNodeSpecRequest {
	s.TargetSecondaryZoneId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetTargetVswitchId(v string) *ModifyNodeSpecRequest {
	s.TargetVswitchId = &v
	return s
}

func (s *ModifyNodeSpecRequest) SetTargetZoneId(v string) *ModifyNodeSpecRequest {
	s.TargetZoneId = &v
	return s
}

func (s *ModifyNodeSpecRequest) Validate() error {
	return dara.Validate(s)
}
