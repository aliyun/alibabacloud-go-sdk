// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAdditionalBandwidthRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAutoPay(v bool) *EnableAdditionalBandwidthRequest
  GetAutoPay() *bool 
  SetAutoRenew(v bool) *EnableAdditionalBandwidthRequest
  GetAutoRenew() *bool 
  SetAutoRenewPeriod(v int32) *EnableAdditionalBandwidthRequest
  GetAutoRenewPeriod() *int32 
  SetBandWidthBurst(v bool) *EnableAdditionalBandwidthRequest
  GetBandWidthBurst() *bool 
  SetBandwidth(v string) *EnableAdditionalBandwidthRequest
  GetBandwidth() *string 
  SetChargeType(v string) *EnableAdditionalBandwidthRequest
  GetChargeType() *string 
  SetCouponNo(v string) *EnableAdditionalBandwidthRequest
  GetCouponNo() *string 
  SetInstanceId(v string) *EnableAdditionalBandwidthRequest
  GetInstanceId() *string 
  SetNodeId(v string) *EnableAdditionalBandwidthRequest
  GetNodeId() *string 
  SetOrderTimeLength(v string) *EnableAdditionalBandwidthRequest
  GetOrderTimeLength() *string 
  SetOwnerAccount(v string) *EnableAdditionalBandwidthRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableAdditionalBandwidthRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnableAdditionalBandwidthRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableAdditionalBandwidthRequest
  GetResourceOwnerId() *int64 
  SetSecurityToken(v string) *EnableAdditionalBandwidthRequest
  GetSecurityToken() *string 
  SetSourceBiz(v string) *EnableAdditionalBandwidthRequest
  GetSourceBiz() *string 
}

type EnableAdditionalBandwidthRequest struct {
  // Specifies whether to enable automatic payment. Default value: true. Valid values:
  // 
  // 	- **true**: enables automatic payment. Make sure that you have sufficient balance within your account.
  // 
  // 	- **false**: disables automatic payment. If automatic payment is disabled, you must perform the following steps to complete the payment in the Tair (Redis OSS-compatible) console: In the top navigation bar, choose **Expenses*	- > **Renewal Management**. In the left-side navigation pane, click **Orders**. On the **Orders*	- page, find the order and complete the payment.
  // 
  // example:
  // 
  // true
  AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
  // Specifies whether to enable auto-renewal. Valid values:
  // 
  // 	- **true**: enables auto-renewal.
  // 
  // 	- **false**: disables auto-renewal. This is the default value.
  // 
  // example:
  // 
  // false
  AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
  // The auto-renewal cycle based on which Tair (Redis OSS-compatible) automatically renews the purchased bandwidth. Unit: months. Valid values: **1**, **2**, **3**, **4**, **5**, **6**, **7**, **8**, **9**, **12**, **24**, **36**, and **60**.
  // 
  // > 	- This parameter takes effect and must be specified only when you set the **AutoRenew*	- parameter to **true**.
  // 
  // > 	- You cannot query the auto-renewal cycle by calling an API operation. To obtain the auto-renewal cycle, you can perform the following procedure: In the top navigation bar of the Tair (Redis OSS-compatible) console, choose **Expenses*	- > **Renewal Management**. On the page that appears, enter the ID of the instance and the `-bw` suffix in the **Instance ID*	- field. Example: r-bp1zxszhcgatnx****-bw.
  // 
  // example:
  // 
  // 1
  AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
  BandWidthBurst *bool `json:"BandWidthBurst,omitempty" xml:"BandWidthBurst,omitempty"`
  // The amount of extra bandwidth that you want to purchase. Unit: Mbit/s. The value must be an integer greater than or equal to **0**. The maximum value can be up to six times the default bandwidth of the instance or a single shard, but cannot exceed 192 Mbit/s. For example, if the default bandwidth of an instance is 10 Mbit/s, the value range of this parameter is **0*	- to **60**.
  // 
  // > 	- You can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/473782.html) operation to obtain the default maximum bandwidth returned by the **DefaultBandWidth*	- response parameter. For more information about instance types, see [Overview](https://help.aliyun.com/document_detail/26350.html).
  // 
  // > -   If you specify multiple data shard IDs in the **NodeId*	- parameter, you must specify the amount of bandwidth that you want to purchase for each specified data shard in the Bandwidth parameter. The bandwidth values that you specify in the Bandwidth parameter must be in the same sequence as the data shard IDs that you specify in the NodeId parameter. In addition, you must separate the bandwidth values with commas (,).
  // 
  // example:
  // 
  // 20
  Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
  // The billing method of the bandwidth instance. Default value: PostPaid. Valid values:
  // 
  // - PrePaid: subscription
  // 
  // - PostPaid: pay-as-you-go
  // 
  // example:
  // 
  // PostPaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // The coupon ID.
  // 
  // example:
  // 
  // youhuiquan_promotion_option_id_for_blank
  CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
  // The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the IDs of instances.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // r-bp1zxszhcgatnx****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the data shard for which you want to purchase a specific amount of bandwidth. You can call the [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) operation to query the IDs of the data shards in an instance. If you specify multiple data shard IDs, separate the data shard IDs with commas (,). You can also set this parameter to **All**, which specifies all the data shards of the instance.
  // 
  // >  This parameter is valid and required only if the instance is a [cluster](https://help.aliyun.com/document_detail/52228.html) instance or [read/write splitting](https://help.aliyun.com/document_detail/62870.html) instance.
  // 
  // example:
  // 
  // r-bp1zxszhcgatnx****-db-0
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  // The validity period of the bandwidth that you purchase. Unit: day. Valid values: **1**, **2**, **3**, **7**, **14**, **30**, **60**, **90**, **180**, **365**, **730**, **1095**, and **1825**.
  // 
  // > If you want to continue using the purchased bandwidth after the specified period of time elapses, you must call the [RenewAdditionalBandwidth](https://help.aliyun.com/document_detail/473804.html) operation to submit a renewal order.
  // 
  // example:
  // 
  // 30
  OrderTimeLength *string `json:"OrderTimeLength,omitempty" xml:"OrderTimeLength,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
  // The source of the operation. This parameter is used only for internal maintenance. You do not need to specify this parameter.
  // 
  // example:
  // 
  // SDK
  SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
}

func (s EnableAdditionalBandwidthRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAdditionalBandwidthRequest) GoString() string {
  return s.String()
}

func (s *EnableAdditionalBandwidthRequest) GetAutoPay() *bool  {
  return s.AutoPay
}

func (s *EnableAdditionalBandwidthRequest) GetAutoRenew() *bool  {
  return s.AutoRenew
}

func (s *EnableAdditionalBandwidthRequest) GetAutoRenewPeriod() *int32  {
  return s.AutoRenewPeriod
}

func (s *EnableAdditionalBandwidthRequest) GetBandWidthBurst() *bool  {
  return s.BandWidthBurst
}

func (s *EnableAdditionalBandwidthRequest) GetBandwidth() *string  {
  return s.Bandwidth
}

func (s *EnableAdditionalBandwidthRequest) GetChargeType() *string  {
  return s.ChargeType
}

func (s *EnableAdditionalBandwidthRequest) GetCouponNo() *string  {
  return s.CouponNo
}

func (s *EnableAdditionalBandwidthRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableAdditionalBandwidthRequest) GetNodeId() *string  {
  return s.NodeId
}

func (s *EnableAdditionalBandwidthRequest) GetOrderTimeLength() *string  {
  return s.OrderTimeLength
}

func (s *EnableAdditionalBandwidthRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableAdditionalBandwidthRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableAdditionalBandwidthRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableAdditionalBandwidthRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableAdditionalBandwidthRequest) GetSecurityToken() *string  {
  return s.SecurityToken
}

func (s *EnableAdditionalBandwidthRequest) GetSourceBiz() *string  {
  return s.SourceBiz
}

func (s *EnableAdditionalBandwidthRequest) SetAutoPay(v bool) *EnableAdditionalBandwidthRequest {
  s.AutoPay = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetAutoRenew(v bool) *EnableAdditionalBandwidthRequest {
  s.AutoRenew = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetAutoRenewPeriod(v int32) *EnableAdditionalBandwidthRequest {
  s.AutoRenewPeriod = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetBandWidthBurst(v bool) *EnableAdditionalBandwidthRequest {
  s.BandWidthBurst = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetBandwidth(v string) *EnableAdditionalBandwidthRequest {
  s.Bandwidth = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetChargeType(v string) *EnableAdditionalBandwidthRequest {
  s.ChargeType = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetCouponNo(v string) *EnableAdditionalBandwidthRequest {
  s.CouponNo = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetInstanceId(v string) *EnableAdditionalBandwidthRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetNodeId(v string) *EnableAdditionalBandwidthRequest {
  s.NodeId = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetOrderTimeLength(v string) *EnableAdditionalBandwidthRequest {
  s.OrderTimeLength = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetOwnerAccount(v string) *EnableAdditionalBandwidthRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetOwnerId(v int64) *EnableAdditionalBandwidthRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetResourceOwnerAccount(v string) *EnableAdditionalBandwidthRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetResourceOwnerId(v int64) *EnableAdditionalBandwidthRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetSecurityToken(v string) *EnableAdditionalBandwidthRequest {
  s.SecurityToken = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) SetSourceBiz(v string) *EnableAdditionalBandwidthRequest {
  s.SourceBiz = &v
  return s
}

func (s *EnableAdditionalBandwidthRequest) Validate() error {
  return dara.Validate(s)
}

