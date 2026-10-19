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
  // Specifies whether to enable automatic payment. Valid values:
  // 
  // 	- **true**: enables automatic payment. This is the default value. Make sure that your account has sufficient balance.
  // 
  // 	- **false**: disables automatic payment. You must manually complete the payment.
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
  // The auto-renewal period, in months. Valid values: **1**, **2**, **3**, **4**, **5**, **6**, **7**, **8**, **9**, **12**, **24**, **36**, and **60**.
  // 
  // > 	- This parameter is available and required only when the **AutoRenew*	- parameter is set to **true**.
  // 
  // > 	- After this parameter is configured, it cannot be queried by calling API operations. To query the auto-renewal period, choose **Billing*	- > **Renewal Management*	- in the top navigation bar of the console, and then enter the instance ID with the `-bw` suffix (for example, r-bp1zxszhcgatnx****-bw) in the **Instance ID*	- text box.
  // 
  // example:
  // 
  // 1
  AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
  // Specifies whether to enable burstable (second-level) elastic bandwidth. After elastic scaling for bandwidth is enabled, burstable (second-level) elastic bandwidth cannot be enabled, and vice versa.
  // 
  // example:
  // 
  // false
  BandWidthBurst *bool `json:"BandWidthBurst,omitempty" xml:"BandWidthBurst,omitempty"`
  // The bandwidth to be increased, in MB/s. The value must be an integer greater than or equal to **0**. The maximum value is six times the default bandwidth of the instance (or a single shard) specifications, but the upper limit of the increase is 192 MB/s. For example, if the default bandwidth of the instance is 10 MB/s, the value range is **0*	- to **60**.
  // 
  // > 	- You can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/473782.html) operation and use the **DefaultBandWidth*	- response parameter to obtain the default maximum bandwidth. For more information about instance specifications, see [Overview of instance specifications](https://help.aliyun.com/document_detail/26350.html).
  // 
  // > 	- If multiple data shard IDs are specified in the **NodeId*	- request parameter, the bandwidth values specified in this parameter must correspond to the order of the shard IDs. Separate multiple bandwidth values with commas (,).
  // 
  // example:
  // 
  // 20
  Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
  // The billing method for the additional bandwidth. Valid values:
  // 
  // 	- **PrePaid**: subscription.
  // 
  // 	- **PostPaid**: pay-as-you-go. Currently, only this billing method is supported.
  // 
  // example:
  // 
  // PrePaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // The ID of the coupon.
  // 
  // example:
  // 
  // youhuiquan_promotion_option_id_for_blank
  CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
  // The ID of the instance. You can call [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) to obtain the instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // r-bp1zxszhcgatnx****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the data shard. You can call [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) to obtain the shard ID. Separate multiple shard IDs with commas (,). You can also set this parameter to **All*	- to specify all data shards.
  // 
  // > This parameter is available and required only when the instance uses the [cluster architecture](https://help.aliyun.com/document_detail/52228.html) or [read/write splitting architecture](https://help.aliyun.com/document_detail/62870.html).
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
  // The source of the call. This parameter is used only for internal maintenance and does not need to be specified.
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

