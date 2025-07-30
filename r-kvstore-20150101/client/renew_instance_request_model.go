// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewInstanceRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewInstanceRequest
	GetAutoRenew() *bool
	SetBusinessInfo(v string) *RenewInstanceRequest
	GetBusinessInfo() *string
	SetCapacity(v string) *RenewInstanceRequest
	GetCapacity() *string
	SetClientToken(v string) *RenewInstanceRequest
	GetClientToken() *string
	SetCouponNo(v string) *RenewInstanceRequest
	GetCouponNo() *string
	SetFromApp(v string) *RenewInstanceRequest
	GetFromApp() *string
	SetInstanceClass(v string) *RenewInstanceRequest
	GetInstanceClass() *string
	SetInstanceId(v string) *RenewInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RenewInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RenewInstanceRequest
	GetOwnerId() *int64
	SetPeriod(v int64) *RenewInstanceRequest
	GetPeriod() *int64
	SetResourceOwnerAccount(v string) *RenewInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RenewInstanceRequest
	GetSecurityToken() *string
}

type RenewInstanceRequest struct {
	// Specifies whether to enable automatic payment. Default value: true. Valid values:
	//
	// 	- **true**: enables automatic payment.
	//
	// 	- **false**: disables automatic payment.
	//
	// If you select false, you must choose **Expenses*	- > **Renewal Management*	- in the top navigation bar. In the left-side navigation pane, click **Orders**. Find the specified order and pay for it.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**: enables auto-renewal. The instance is renewed based on the specified renewal duration. For example, if you set the renewal duration to three months, you are charged for three months of service each time the instance is automatically renewed.
	//
	// 	- **false*	- (default): disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the promotional event or business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The storage capacity of the instance. Unit: MB. When you renew the instance, you can specify this parameter to change specifications of the instance.
	//
	// > To change the specifications when you renew the instance, you must specify at least one of the `Capacity` and `InstanceClass` parameters.
	//
	// example:
	//
	// 1024
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token is case-sensitive. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// TF-ModifyInstanceSpec-1686645570-7dac7257-4a14-4811-939c-51a282f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The source of the request. The default value is **OpenAPI*	- and cannot be changed.
	//
	// example:
	//
	// OpenAPI
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The instance type code. For more information, see [Instance specifications overview](https://help.aliyun.com/document_detail/26350.html). When you renew the instance, you can specify this parameter to change specifications of the instance.
	//
	// > To change the specifications when you renew the instance, you must specify at least one of the `Capacity` and `InstanceClass` parameters.
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The renewal period. Valid values: **1**, 2, 3, 4, 5, 6, 7, 8, **9**, **12**, **24**, and **36**. Unit: months.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *RenewInstanceRequest) GetCapacity() *string {
	return s.Capacity
}

func (s *RenewInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewInstanceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *RenewInstanceRequest) GetFromApp() *string {
	return s.FromApp
}

func (s *RenewInstanceRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *RenewInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RenewInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewInstanceRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *RenewInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RenewInstanceRequest) SetAutoPay(v bool) *RenewInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewInstanceRequest) SetAutoRenew(v bool) *RenewInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewInstanceRequest) SetBusinessInfo(v string) *RenewInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *RenewInstanceRequest) SetCapacity(v string) *RenewInstanceRequest {
	s.Capacity = &v
	return s
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetCouponNo(v string) *RenewInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *RenewInstanceRequest) SetFromApp(v string) *RenewInstanceRequest {
	s.FromApp = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceClass(v string) *RenewInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerAccount(v string) *RenewInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v int64) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceOwnerAccount(v string) *RenewInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceOwnerId(v int64) *RenewInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetSecurityToken(v string) *RenewInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
