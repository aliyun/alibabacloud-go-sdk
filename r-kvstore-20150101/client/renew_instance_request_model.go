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
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: enables automatic payment. This is the default value.
	//
	// 	- **false**: disables automatic payment.
	//
	// <props="china">If you set this parameter to false, you must log on to the console, choose **Billing*	- > **Renewal Management*	- in the top navigation bar, and then choose **Order Management*	- > **My Orders*	- in the left-side navigation pane to find and pay for the order.
	//
	// <props="intl">If you set this parameter to false, you must log on to the console, choose **Billing*	- > **Renewal Management*	- in the top navigation bar, and then click **Order Management*	- in the left-side navigation pane to find and pay for the order.
	//
	//
	//
	// <props="partner">If you set this parameter to false, you must log on to the console, choose **Billing*	- > **Renewal Management*	- in the top navigation bar, and then click **Order Management*	- in the left-side navigation pane to find and pay for the order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: enables auto-renewal. The instance is renewed for the specified period. For example, if you set the renewal period to 3 months, the instance is automatically renewed for 3 months each time.
	//
	// 	- **false**: disables auto-renewal. This is the default value.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The activity ID and business information.
	//
	// example:
	//
	// 000000000
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// **[Deprecated]*	- You cannot change the storage capacity when you renew an instance. By default, the renewal order is placed based on the current specifications.
	//
	// example:
	//
	// -
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token is case-sensitive and can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 0.6823317918435086
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The source of the request. Default value: **OpenAPI**. You do not need to manually specify this parameter.
	//
	// example:
	//
	// OpenAPI
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// **[Deprecated]*	- You cannot change the instance type when you renew an instance. By default, the renewal order is placed based on the current specifications.
	//
	// example:
	//
	// -
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The renewal period. Unit: months. Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// >For example, if an instance expires in 3 months and you renew it for 3 months, the instance will expire in 6 months after the renewal.
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
