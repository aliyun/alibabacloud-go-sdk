// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocatePublicIp(v bool) *ModifyInstanceNetworkSpecRequest
	GetAllocatePublicIp() *bool
	SetAutoPay(v bool) *ModifyInstanceNetworkSpecRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyInstanceNetworkSpecRequest
	GetClientToken() *string
	SetEndTime(v string) *ModifyInstanceNetworkSpecRequest
	GetEndTime() *string
	SetISP(v string) *ModifyInstanceNetworkSpecRequest
	GetISP() *string
	SetInstanceId(v string) *ModifyInstanceNetworkSpecRequest
	GetInstanceId() *string
	SetInternetMaxBandwidthIn(v int32) *ModifyInstanceNetworkSpecRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *ModifyInstanceNetworkSpecRequest
	GetInternetMaxBandwidthOut() *int32
	SetNetworkChargeType(v string) *ModifyInstanceNetworkSpecRequest
	GetNetworkChargeType() *string
	SetOwnerAccount(v string) *ModifyInstanceNetworkSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceNetworkSpecRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceNetworkSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceNetworkSpecRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *ModifyInstanceNetworkSpecRequest
	GetStartTime() *string
}

type ModifyInstanceNetworkSpecRequest struct {
	// Specifies whether to allocate a public IP address.
	//
	// - true: allocates a public IP address.
	//
	// - false: does not allocate a public IP address.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AllocatePublicIp *bool `json:"AllocatePublicIp,omitempty" xml:"AllocatePublicIp,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - true: After the bandwidth configuration is modified, the payment is automatically deducted. When you set AutoPay to true, make sure that your account balance is sufficient. If your account balance is insufficient, an abnormal order will be generated. This order cannot be paid through the ECS console and can only be voided.
	//
	// <props="china">
	//
	// - false: After the bandwidth configuration is modified, only an order is generated but not paid. If your payment method balance is insufficient, you can set AutoPay to false to cancel automatic payment. In this case, an unpaid order is generated when you call this operation. You can log on to the [ECS console](https://ecs.console.aliyun.com) to pay for the order.
	//
	//
	//
	// <props="intl">
	//
	// - false: After the bandwidth configuration is modified, only an order is generated but not paid. If your payment method balance is insufficient, you can set AutoPay to false to cancel automatic payment. In this case, an unpaid order is generated when you call this operation. You can log on to the [ECS console](https://ecs.console.aliyun.com) to pay for the order.
	//
	//
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end time of the temporary bandwidth upgrade. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in UTC+0 time in the format of yyyy-MM-ddTHHZ. The precision is down to **hours*	- (HH).
	//
	// > The interval between the end time and the start time of the temporary bandwidth upgrade must be greater than or equal to 3 hours.
	//
	// example:
	//
	// 2017-12-06T22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// > This parameter is in invitational preview and is not yet publicly available.
	//
	// example:
	//
	// null
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The ID of the instance for which you want to modify the network configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s (Megabit per second). Valid values:
	//
	// - When the purchased maximum outbound public bandwidth is less than or equal to 10 Mbit/s: 1 to 10. Default value: 10.
	//
	// - When the purchased maximum outbound public bandwidth is greater than 10 Mbit/s: 1 to the value of `InternetMaxBandwidthOut`. Default value: the value of `InternetMaxBandwidthOut`.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s (Megabit per second). Valid values:
	//
	// - Pay-by-traffic: 0 to 100.
	//
	// - Pay-by-bandwidth:
	//
	//   - Subscription instances: 0 to 200.
	//
	//   - Pay-as-you-go instances: 0 to 100.
	//
	//
	// > The maximum outbound bandwidth for a single instance is also limited by the **Network bandwidth baseline/burst (Gbit/s)*	- metric of the ECS instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The network billing method. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-traffic.
	//
	// > In **pay-by-traffic*	- mode, the inbound and outbound bandwidth peak values are both bandwidth upper limits and are not guaranteed as committed service metrics. When resource contention occurs, the bandwidth peak values may be limited. If your business requires guaranteed bandwidth, use the **pay-by-bandwidth*	- mode.
	//
	// example:
	//
	// PayByTraffic
	NetworkChargeType    *string `json:"NetworkChargeType,omitempty" xml:"NetworkChargeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time of the temporary bandwidth upgrade. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in UTC+0 time in the format of yyyy-MM-ddTHH:mmZ. The precision is down to **minutes*	- (mm).
	//
	// example:
	//
	// 2017-12-05T22:40Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyInstanceNetworkSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkSpecRequest) GetAllocatePublicIp() *bool {
	return s.AllocatePublicIp
}

func (s *ModifyInstanceNetworkSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceNetworkSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceNetworkSpecRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyInstanceNetworkSpecRequest) GetISP() *string {
	return s.ISP
}

func (s *ModifyInstanceNetworkSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyInstanceNetworkSpecRequest) GetNetworkChargeType() *string {
	return s.NetworkChargeType
}

func (s *ModifyInstanceNetworkSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceNetworkSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceNetworkSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceNetworkSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceNetworkSpecRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyInstanceNetworkSpecRequest) SetAllocatePublicIp(v bool) *ModifyInstanceNetworkSpecRequest {
	s.AllocatePublicIp = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetAutoPay(v bool) *ModifyInstanceNetworkSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetClientToken(v string) *ModifyInstanceNetworkSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetEndTime(v string) *ModifyInstanceNetworkSpecRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetISP(v string) *ModifyInstanceNetworkSpecRequest {
	s.ISP = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetInstanceId(v string) *ModifyInstanceNetworkSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthIn(v int32) *ModifyInstanceNetworkSpecRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthOut(v int32) *ModifyInstanceNetworkSpecRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetNetworkChargeType(v string) *ModifyInstanceNetworkSpecRequest {
	s.NetworkChargeType = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetOwnerAccount(v string) *ModifyInstanceNetworkSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetOwnerId(v int64) *ModifyInstanceNetworkSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetResourceOwnerAccount(v string) *ModifyInstanceNetworkSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetResourceOwnerId(v int64) *ModifyInstanceNetworkSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetStartTime(v string) *ModifyInstanceNetworkSpecRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) Validate() error {
	return dara.Validate(s)
}
