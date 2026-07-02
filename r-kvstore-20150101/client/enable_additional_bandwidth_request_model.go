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
  // Specifies whether to enable auto-payment. Valid values:
  // 
  // - **true**: Enables auto-payment. This is the default. Ensure that your account has a sufficient balance.
  // 
  // - **false**: Disables auto-payment. You must manually complete the payment.
  // 
  // example:
  // 
  // true
  AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
  // Specifies whether to enable auto-renewal. Valid values:
  // 
  // - **true**: Enables auto-renewal.
  // 
  // - **false**: Disables auto-renewal. This is the default.
  // 
  // example:
  // 
  // false
  AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
  // The auto-renewal period, in months. Valid values: **1**, **2**, **3**, **4**, **5**, **6**, **7**, **8**, **9**, **12**, **24**, **36**, and **60**.
  // 
  // > - This parameter is required only when **AutoRenew*	- is set to **true**.
  // 
  // >
  // 
  // > - After you set this parameter, you cannot query its value by calling API operations. To check this setting, go to the console. In the top navigation bar, choose **Billing*	- > **Renewal Management**. Then, in the **Instance ID*	- field, enter the instance ID followed by the `-bw` suffix (for example, r-bp1zxszhcgatnx\\*\\*\\*\\*-bw).
  // 
  // example:
  // 
  // 1
  AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
  BandWidthBurst *bool `json:"BandWidthBurst,omitempty" xml:"BandWidthBurst,omitempty"`
  // The amount of bandwidth to add, in MB/s. The value must be an integer greater than or equal to **0**. The maximum value is six times the default bandwidth of the instance type or a single data shard, with an upper limit of 192 MB/s. For example, if the default bandwidth of an instance is 10 MB/s, the valid values for this parameter are **0*	- to **60**.
  // 
  // > - You can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/473782.html) operation and check the value of the **DefaultBandWidth*	- parameter in the response to get the default maximum bandwidth. For more information about instance types, see [Instance types](https://help.aliyun.com/document_detail/26350.html).
  // 
  // >
  // 
  // > - If you specified multiple data shard IDs for the **NodeId*	- parameter, the bandwidth values that you specify for this parameter must correspond to the order of the data shard IDs. Separate multiple bandwidth values with commas (,).
  // 
  // example:
  // 
  // 20
  Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
  // The billing method for the additional bandwidth. Valid values:
  // 
  // - **PrePaid**: subscription.
  // 
  // - **PostPaid**: pay-as-you-go. This is the only supported billing method.
  // 
  // example:
  // 
  // PrePaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // The coupon ID.
  // 
  // example:
  // 
  // youhuiquan_promotion_option_id_for_blank
  CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
  // The instance ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // r-bp1zxszhcgatnx****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the data shard. You can call the [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) operation to query data shard IDs. Separate multiple data shard IDs with commas (,). You can also set this parameter to **All*	- to specify all data shards.
  // 
  // > This parameter is required only when the instance uses the [cluster architecture](https://help.aliyun.com/document_detail/52228.html) or the [read/write splitting architecture](https://help.aliyun.com/document_detail/62870.html).
  // 
  // example:
  // 
  // r-bp1zxszhcgatnx****-db-0
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  // The subscription duration, in days. Valid values: **1**, **2**, **3**, **7**, **14**, **30**, **60**, **90**, **180**, **365**, **730**, **1095**, and **1825**.
  // 
  // > To continue using the purchased bandwidth, you must call the [RenewAdditionalBandwidth](https://help.aliyun.com/document_detail/473804.html) operation to renew the bandwidth before it expires.
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
  // The source of the call. This parameter is used for internal maintenance. Do not specify it.
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

