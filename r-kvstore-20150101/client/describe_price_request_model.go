// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessInfo(v string) *DescribePriceRequest
	GetBusinessInfo() *string
	SetCapacity(v int64) *DescribePriceRequest
	GetCapacity() *int64
	SetChargeType(v string) *DescribePriceRequest
	GetChargeType() *string
	SetCouponNo(v string) *DescribePriceRequest
	GetCouponNo() *string
	SetEngineVersion(v string) *DescribePriceRequest
	GetEngineVersion() *string
	SetForceUpgrade(v bool) *DescribePriceRequest
	GetForceUpgrade() *bool
	SetInstanceClass(v string) *DescribePriceRequest
	GetInstanceClass() *string
	SetInstanceId(v string) *DescribePriceRequest
	GetInstanceId() *string
	SetInstances(v string) *DescribePriceRequest
	GetInstances() *string
	SetNeedUnavailableCoupon(v bool) *DescribePriceRequest
	GetNeedUnavailableCoupon() *bool
	SetNodeType(v string) *DescribePriceRequest
	GetNodeType() *string
	SetOrderParamOut(v string) *DescribePriceRequest
	GetOrderParamOut() *string
	SetOrderType(v string) *DescribePriceRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribePriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePriceRequest
	GetOwnerId() *int64
	SetPeriod(v int64) *DescribePriceRequest
	GetPeriod() *int64
	SetQuantity(v int64) *DescribePriceRequest
	GetQuantity() *int64
	SetRegionId(v string) *DescribePriceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePriceRequest
	GetResourceOwnerId() *int64
	SetSecondaryZoneId(v string) *DescribePriceRequest
	GetSecondaryZoneId() *string
	SetSecurityToken(v string) *DescribePriceRequest
	GetSecurityToken() *string
	SetShardCount(v int32) *DescribePriceRequest
	GetShardCount() *int32
	SetZoneId(v string) *DescribePriceRequest
	GetZoneId() *string
}

type DescribePriceRequest struct {
	// The extended information, such as the promotion ID and business information.
	//
	// example:
	//
	// 000000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// Specifies the instance type by using the storage capacity of the instance. Unit: MB. This parameter applies only to instances of ApsaraDB for Redis Community Edition in the classic deployment mode. We recommend that you use the **InstanceClass*	- parameter to specify an exact instance type.
	//
	// > If you specify the **InstanceClass*	- parameter, you do not need to specify this parameter.
	//
	// example:
	//
	// 1024
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PostPaid*	- (default): pay-as-you-go.
	//
	// 	- **PrePaid**: subscription.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The coupon code. Default value: youhuiquan_promotion_option_id_for_blank, which indicates that no coupon code is used.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// <props="china">Database engine version of the Redis instance. Valid values: **2.8**, **4.0**, **5.0**, and **6.0**.
	//
	// <props="intl">Database engine version of the Redis instance. Valid values: **2.8**, **4.0**, and **5.0**.
	//
	// <props="partner">Database engine version of the Redis instance. Valid values: **2.8**, **4.0**, and **5.0**.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether to forcibly upgrade or downgrade the instance. Valid values:
	//
	// 	- **false**: Yes.
	//
	// 	- **true*	- (default): No.
	//
	// example:
	//
	// true
	ForceUpgrade *bool `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	// Specifies the instance type by using the **InstanceClass*	- code. To query the **InstanceClass*	- code, perform the following steps:
	//
	// 1. On the [Instance type query](https://help.aliyun.com/document_detail/26350.html) page, select the instance type as needed and click the value in the **Architecture Type*	- column to open the instance type documentation.
	//
	// 2. In the instance type table of the documentation, select the required instance type and find the code in the **InstanceClass*	- column.
	//
	// > For example, to query pricing for a new purchase, you must specify at least the following parameters in addition to other key parameters:
	//
	// > 	- To query a classic instance: specify this parameter.
	//
	// > 	- To query a cloud-native standard architecture instance: specify this parameter.
	//
	// > 	- To query a cloud-native cluster architecture instance: specify this parameter and the **ShardCount*	- parameter.
	//
	// > 	- To query a cloud-native read/write splitting architecture instance: specify this parameter and the **Instances*	- parameter.
	//
	// > 	- To query multiple instances of different types, or to query Tair ESSD instances (with custom storage types and storage space): specify the Instances parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID.
	//
	// > This parameter is required when **OrderType*	- is set to **UPGRADE*	- or **RENEW**.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instances to query. This parameter is required when you query cloud-native read/write splitting architecture instances, Tair ESSD instances, or multiple instances of different types. The value is a JSON string. For more information, see **Supplementary description of the Instances parameter*	- below.
	//
	// example:
	//
	// [{"RegionId": "cn-hangzhou","ZoneId": "cn-hangzhou-b","ShardClass": "tair.scm.with.proxy.standard.2m.8d","ShardCount": "3","Period": "1","Quantity": "1"}]
	Instances *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// Specifies whether to return unavailable coupons.
	NeedUnavailableCoupon *bool `json:"NeedUnavailableCoupon,omitempty" xml:"NeedUnavailableCoupon,omitempty"`
	// The node type. Valid values:
	//
	// - **STAND_ALONE**: Single-replica.
	//
	// - **MASTER_SLAVE*	- (default): High-availability (dual-replica).
	//
	// example:
	//
	// MASTER_SLAVE
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// Specifies whether to return the order parameters. Valid values:
	//
	// 	- **false*	- (default): Not returned.
	//
	// 	- **true**: Returned.
	//
	// example:
	//
	// true
	OrderParamOut *string `json:"OrderParamOut,omitempty" xml:"OrderParamOut,omitempty"`
	// The order type. Valid values:
	//
	// 	- **BUY**: New purchase.
	//
	// 	- **UPGRADE**: Upgrade or downgrade.
	//
	// 	- **RENEW**: Renewal.
	//
	// 	- **CONVERT**: Billing method conversion.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUY
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Unit: months. Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// example:
	//
	// 3
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The number of instances to purchase. Valid values: **1*	- to **30**. Default value: **1**.
	//
	// example:
	//
	// 1
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The secondary zone ID. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query available zones.
	//
	// > The value of this parameter cannot be the same as the value of the ZoneId parameter.
	//
	// example:
	//
	// cn-hangzhou-g
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of data shard nodes in the cloud-native cluster instance.
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The zone ID. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query available zones.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *DescribePriceRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribePriceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribePriceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribePriceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribePriceRequest) GetForceUpgrade() *bool {
	return s.ForceUpgrade
}

func (s *DescribePriceRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePriceRequest) GetInstances() *string {
	return s.Instances
}

func (s *DescribePriceRequest) GetNeedUnavailableCoupon() *bool {
	return s.NeedUnavailableCoupon
}

func (s *DescribePriceRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribePriceRequest) GetOrderParamOut() *string {
	return s.OrderParamOut
}

func (s *DescribePriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePriceRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribePriceRequest) GetQuantity() *int64 {
	return s.Quantity
}

func (s *DescribePriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePriceRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *DescribePriceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePriceRequest) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *DescribePriceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePriceRequest) SetBusinessInfo(v string) *DescribePriceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *DescribePriceRequest) SetCapacity(v int64) *DescribePriceRequest {
	s.Capacity = &v
	return s
}

func (s *DescribePriceRequest) SetChargeType(v string) *DescribePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetCouponNo(v string) *DescribePriceRequest {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceRequest) SetEngineVersion(v string) *DescribePriceRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribePriceRequest) SetForceUpgrade(v bool) *DescribePriceRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceClass(v string) *DescribePriceRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceId(v string) *DescribePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePriceRequest) SetInstances(v string) *DescribePriceRequest {
	s.Instances = &v
	return s
}

func (s *DescribePriceRequest) SetNeedUnavailableCoupon(v bool) *DescribePriceRequest {
	s.NeedUnavailableCoupon = &v
	return s
}

func (s *DescribePriceRequest) SetNodeType(v string) *DescribePriceRequest {
	s.NodeType = &v
	return s
}

func (s *DescribePriceRequest) SetOrderParamOut(v string) *DescribePriceRequest {
	s.OrderParamOut = &v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerAccount(v string) *DescribePriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerId(v int64) *DescribePriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int64) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetQuantity(v int64) *DescribePriceRequest {
	s.Quantity = &v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerAccount(v string) *DescribePriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerId(v int64) *DescribePriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetSecondaryZoneId(v string) *DescribePriceRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *DescribePriceRequest) SetSecurityToken(v string) *DescribePriceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePriceRequest) SetShardCount(v int32) *DescribePriceRequest {
	s.ShardCount = &v
	return s
}

func (s *DescribePriceRequest) SetZoneId(v string) *DescribePriceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	return dara.Validate(s)
}
