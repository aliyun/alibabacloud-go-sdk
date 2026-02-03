// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformInstanceChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *TransformInstanceChargeTypeRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *TransformInstanceChargeTypeRequest
	GetAutoRenew() *string
	SetAutoRenewPeriod(v int64) *TransformInstanceChargeTypeRequest
	GetAutoRenewPeriod() *int64
	SetChargeType(v string) *TransformInstanceChargeTypeRequest
	GetChargeType() *string
	SetCouponNo(v string) *TransformInstanceChargeTypeRequest
	GetCouponNo() *string
	SetInstanceId(v string) *TransformInstanceChargeTypeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *TransformInstanceChargeTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransformInstanceChargeTypeRequest
	GetOwnerId() *int64
	SetPeriod(v int64) *TransformInstanceChargeTypeRequest
	GetPeriod() *int64
	SetResourceOwnerAccount(v string) *TransformInstanceChargeTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransformInstanceChargeTypeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *TransformInstanceChargeTypeRequest
	GetSecurityToken() *string
}

type TransformInstanceChargeTypeRequest struct {
	// Specifies whether to enable automatic payment. Default value: true. Valid values:
	//
	// 	- **true**: Automatic payment is enabled.
	//
	// 	- **false**: Automatic payment is disabled. If automatic payment is disabled, you must perform the following steps to complete the payment: In the top navigation bar of the Tair (Redis OSS-compatible) console, choose **Expenses*	- > **Renewal Management**. In the left-side navigation pane of the Billing Management console, click **Orders**. On the **Orders*	- page, find the order and complete the payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false*	- (default): disables auto-renewal.
	//
	// Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The subscription duration that is supported by auto-renewal. Unit: month. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// >  This parameter is required if the **AutoRenew*	- parameter is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int64 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The new billing method. Valid values:
	//
	// 	- **PrePaid**: subscription. If you set this parameter to PrePaid, you must also specify the **Period*	- parameter.
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CouponNo   *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The ID of the instance. You can call the [DescribeInstances](~~DescribeInstances~~) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Unit: months. Valid values: **1**, 2, 3, 4, 5, 6, 7, 8, **9**, **12**, **24**, **36**.
	//
	// >  This parameter is valid and required only if you set the **ChargeType*	- parameter to **PrePaid**.
	//
	// example:
	//
	// 1
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s TransformInstanceChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *TransformInstanceChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *TransformInstanceChargeTypeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *TransformInstanceChargeTypeRequest) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *TransformInstanceChargeTypeRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *TransformInstanceChargeTypeRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *TransformInstanceChargeTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransformInstanceChargeTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransformInstanceChargeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransformInstanceChargeTypeRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *TransformInstanceChargeTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransformInstanceChargeTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransformInstanceChargeTypeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *TransformInstanceChargeTypeRequest) SetAutoPay(v bool) *TransformInstanceChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetAutoRenew(v string) *TransformInstanceChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetAutoRenewPeriod(v int64) *TransformInstanceChargeTypeRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetChargeType(v string) *TransformInstanceChargeTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetCouponNo(v string) *TransformInstanceChargeTypeRequest {
	s.CouponNo = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetInstanceId(v string) *TransformInstanceChargeTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetOwnerAccount(v string) *TransformInstanceChargeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetOwnerId(v int64) *TransformInstanceChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetPeriod(v int64) *TransformInstanceChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetResourceOwnerAccount(v string) *TransformInstanceChargeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetResourceOwnerId(v int64) *TransformInstanceChargeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) SetSecurityToken(v string) *TransformInstanceChargeTypeRequest {
	s.SecurityToken = &v
	return s
}

func (s *TransformInstanceChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
