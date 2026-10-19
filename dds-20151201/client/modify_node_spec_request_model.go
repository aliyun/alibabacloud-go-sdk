// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlignPrimaryZone(v bool) *ModifyNodeSpecRequest
	GetAlignPrimaryZone() *bool
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
	// Specifies whether to automatically align the primary node to the primary zone after operations such as specification changes for multi-zone instances. If you choose not to align the node, one primary/secondary switchover (transient disconnection) is avoided. Default value: true.
	AlignPrimaryZone *bool `json:"AlignPrimaryZone,omitempty" xml:"AlignPrimaryZone,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true*	- (default): Automatic payment is enabled. Make sure that your account has sufficient balance.
	//
	// - **false**: Manual payment is enabled.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The additional parameter that contains business information.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. Make sure that the value is unique across different requests and contains a maximum of 64 ASCII characters. The value cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **default*	- or **null*	- (default): Use a coupon.
	//
	// - **youhuiquan_promotion_option_id_for_blank**: Do not use a coupon.
	//
	// example:
	//
	// default
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1c0b990184****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The effective period of the configuration change. Valid values:
	//
	// - **Immediately*	- (default): The configuration change takes effect immediately.
	//
	// - **MaintainTime**: The configuration change takes effect within the maintenance window of the instance.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The source of the request. Valid values:
	//
	// - **OpenApi**: The request is sent from OpenAPI.
	//
	// - **mongo_buy**: The request is sent from the console.
	//
	// example:
	//
	// OpenApi
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The specifications of the shard or mongos node. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/57141.html).
	//
	// example:
	//
	// dds.mongos.standard
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The ID of the shard or mongos node in the sharded cluster instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the node ID.
	//
	// > If the value is a shard node ID, you must also specify the **NodeStorage*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp143e1b1637****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The storage space of the shard node. The step size is 10 GB.
	//
	// - For local SSDs, valid values: **10*	- to **2000**.
	//
	// - For PL1 ESSDs, valid values: **20*	- to **16000**.
	//
	// example:
	//
	// 20
	NodeStorage *int32 `json:"NodeStorage,omitempty" xml:"NodeStorage,omitempty"`
	// The order type. Valid values:
	//
	// - **UPGRADE**: Upgrade.
	//
	// - **DOWNGRADE**: Downgrade.
	//
	// example:
	//
	// UPGRADE
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read-only nodes in the shard node.
	//
	// Valid values: **0*	- to **5**. The value must be an integer.
	//
	// example:
	//
	// 5
	ReadonlyReplicas     *int32  `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time when the configuration change is executed. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// > **[Deprecated]*	- This parameter is deprecated. Use the EffectiveTime parameter instead.
	//
	// example:
	//
	// 2022-01-05T03:18:53Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The destination hidden zone for zone migration.
	//
	// 	Notice: This parameter is supported only for cloud disk instances.</notice>
	//
	// example:
	//
	// cn-hangzhou-j
	TargetHiddenZoneId *string `json:"TargetHiddenZoneId,omitempty" xml:"TargetHiddenZoneId,omitempty"`
	// The destination secondary zone for zone migration.
	//
	// 	Notice: This parameter is supported only for cloud disk instances.</notice>
	//
	// example:
	//
	// cn-hangzhou-e
	TargetSecondaryZoneId *string `json:"TargetSecondaryZoneId,omitempty" xml:"TargetSecondaryZoneId,omitempty"`
	// The ID of the destination vSwitch for zone migration.
	//
	// 	Notice: This parameter is supported only for cloud disk instances.</notice>
	//
	// example:
	//
	// vsw-xxxxxxxx
	TargetVswitchId *string `json:"TargetVswitchId,omitempty" xml:"TargetVswitchId,omitempty"`
	// The destination primary zone for zone migration.
	//
	// 	Notice: This parameter is supported only for cloud disk instances.</notice>
	//
	// example:
	//
	// cn-hangzhou-h
	TargetZoneId *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
}

func (s ModifyNodeSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeSpecRequest) GetAlignPrimaryZone() *bool {
	return s.AlignPrimaryZone
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

func (s *ModifyNodeSpecRequest) SetAlignPrimaryZone(v bool) *ModifyNodeSpecRequest {
	s.AlignPrimaryZone = &v
	return s
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
