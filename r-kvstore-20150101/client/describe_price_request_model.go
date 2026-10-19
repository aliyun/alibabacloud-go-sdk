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
	// The extended information such as the promotional event ID and business information.
	//
	// example:
	//
	// 000000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The storage capacity used to specify the instance type, in MB. This parameter is used only to query Redis Community Edition instances in classic deployment mode. We recommend that you use the **InstanceClass*	- parameter to specify an exact instance type.
	//
	// > After you specify the **InstanceClass*	- parameter, you do not need to specify this parameter.
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
	// Specifies whether to forcibly change the configuration. Valid values:
	//
	// 	- **false**: yes.
	//
	// 	- **true*	- (default): no.
	//
	// example:
	//
	// true
	ForceUpgrade *bool `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	// The instance type specified by using the **InstanceClass*	- code. To find the **InstanceClass*	- code:
	//
	// 1. In the [instance type navigation](https://help.aliyun.com/document_detail/26350.html), select the desired instance type and click the value in the **Architecture*	- column to open the specification document.
	//
	// 2. In the instance type table of the specification document, find the desired instance type and locate the code in the **InstanceClass*	- column.
	//
	// > The following example describes the parameters required for querying the price of a new instance. In addition to other key parameters, specify at least the following parameters:
	//
	// > 	- Classic Edition instance: specify this parameter.
	//
	// > 	- Cloud-native Edition standard architecture instance: specify this parameter.
	//
	// > 	- Cloud-native Edition cluster architecture instance: specify this parameter and the **ShardCount*	- parameter.
	//
	// > 	- Cloud-native Edition read/write splitting architecture instance: specify this parameter and the **Instances*	- parameter.
	//
	// > 	- Multiple instances of different types, or Tair ESSD-based instances (custom storage type and storage capacity): specify the Instances parameter (this parameter is not required).
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
	// If you want to query the price of cloud-native read/write splitting architecture instances, Tair ESSD-based instances, or multiple instances of different types, specify this parameter. The value is a JSON-formatted string. For more information, see **Additional description of the Instances parameter*	- below.
	//
	// example:
	//
	// [{"RegionId": "cn-hangzhou","ZoneId": "cn-hangzhou-b","ShardClass": "tair.scm.with.proxy.standard.2m.8d","ShardCount": "3","Period": "1","Quantity": "1"}]
	Instances *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The node type. Valid values:
	//
	// - **STAND_ALONE**: standalone (single replica).
	//
	// - **MASTER_SLAVE*	- (default): high availability (dual replicas).
	//
	// example:
	//
	// MASTER_SLAVE
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// Specifies whether to return order parameters. Valid values:
	//
	// 	- **false*	- (default): does not return order parameters.
	//
	// 	- **true**: returns order parameters.
	//
	// example:
	//
	// true
	OrderParamOut *string `json:"OrderParamOut,omitempty" xml:"OrderParamOut,omitempty"`
	// The order type. Valid values:
	//
	// 	- **BUY**: new purchase.
	//
	// 	- **UPGRADE**: configuration change.
	//
	// 	- **RENEW**: renewal.
	//
	// 	- **CONVERT**: billing method conversion.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUY
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration, in months. Valid values: **1*	- to **9**, **12**, **24**, and **36**.
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
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) to query available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The secondary zone ID. You can call [DescribeZones](https://help.aliyun.com/document_detail/473764.html) to query available zones.
	//
	// > The value of this parameter must be different from the value of the ZoneId parameter.
	//
	// example:
	//
	// cn-hangzhou-g
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of data shard nodes in a cloud-native cluster instance.
	//
	// example:
	//
	// 2
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The zone ID. You can call [DescribeZones](https://help.aliyun.com/document_detail/473764.html) to query available zones.
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
