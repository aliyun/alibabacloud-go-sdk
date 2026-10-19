// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformToPrePaidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *TransformToPrePaidRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *TransformToPrePaidRequest
	GetAutoRenew() *string
	SetAutoRenewPeriod(v int64) *TransformToPrePaidRequest
	GetAutoRenewPeriod() *int64
	SetCouponNo(v string) *TransformToPrePaidRequest
	GetCouponNo() *string
	SetInstanceId(v string) *TransformToPrePaidRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *TransformToPrePaidRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransformToPrePaidRequest
	GetOwnerId() *int64
	SetPeriod(v int64) *TransformToPrePaidRequest
	GetPeriod() *int64
	SetResourceOwnerAccount(v string) *TransformToPrePaidRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransformToPrePaidRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *TransformToPrePaidRequest
	GetSecurityToken() *string
}

type TransformToPrePaidRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: enables automatic payment.
	//
	// 	- **false**: disables automatic payment (default). You must manually renew the instance in the console. For more information, see [Manual renewal](https://help.aliyun.com/document_detail/26352.html).
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false**: disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal epoch. Unit: months. Valid values: **1**, **2**, **3**, **6**, and **12**.
	//
	// > This parameter is required when **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *int64  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	CouponNo        *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// Instance ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Unit: months. Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s TransformToPrePaidRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformToPrePaidRequest) GoString() string {
	return s.String()
}

func (s *TransformToPrePaidRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *TransformToPrePaidRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *TransformToPrePaidRequest) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *TransformToPrePaidRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *TransformToPrePaidRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransformToPrePaidRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransformToPrePaidRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransformToPrePaidRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *TransformToPrePaidRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransformToPrePaidRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransformToPrePaidRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *TransformToPrePaidRequest) SetAutoPay(v bool) *TransformToPrePaidRequest {
	s.AutoPay = &v
	return s
}

func (s *TransformToPrePaidRequest) SetAutoRenew(v string) *TransformToPrePaidRequest {
	s.AutoRenew = &v
	return s
}

func (s *TransformToPrePaidRequest) SetAutoRenewPeriod(v int64) *TransformToPrePaidRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *TransformToPrePaidRequest) SetCouponNo(v string) *TransformToPrePaidRequest {
	s.CouponNo = &v
	return s
}

func (s *TransformToPrePaidRequest) SetInstanceId(v string) *TransformToPrePaidRequest {
	s.InstanceId = &v
	return s
}

func (s *TransformToPrePaidRequest) SetOwnerAccount(v string) *TransformToPrePaidRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformToPrePaidRequest) SetOwnerId(v int64) *TransformToPrePaidRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformToPrePaidRequest) SetPeriod(v int64) *TransformToPrePaidRequest {
	s.Period = &v
	return s
}

func (s *TransformToPrePaidRequest) SetResourceOwnerAccount(v string) *TransformToPrePaidRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformToPrePaidRequest) SetResourceOwnerId(v int64) *TransformToPrePaidRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformToPrePaidRequest) SetSecurityToken(v string) *TransformToPrePaidRequest {
	s.SecurityToken = &v
	return s
}

func (s *TransformToPrePaidRequest) Validate() error {
	return dara.Validate(s)
}
