// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAdditionalBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewAdditionalBandwidthRequest
	GetAutoPay() *bool
	SetCouponNo(v string) *RenewAdditionalBandwidthRequest
	GetCouponNo() *string
	SetInstanceId(v string) *RenewAdditionalBandwidthRequest
	GetInstanceId() *string
	SetOrderTimeLength(v string) *RenewAdditionalBandwidthRequest
	GetOrderTimeLength() *string
	SetOwnerAccount(v string) *RenewAdditionalBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RenewAdditionalBandwidthRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RenewAdditionalBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewAdditionalBandwidthRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RenewAdditionalBandwidthRequest
	GetSecurityToken() *string
	SetSourceBiz(v string) *RenewAdditionalBandwidthRequest
	GetSourceBiz() *string
}

type RenewAdditionalBandwidthRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: Automatic payment. This is the default value.
	//
	// 	- **false**: Manual payment. You must choose **Expenses and Costs*	- > **Renewals*	- in the top navigation bar of the console, and then choose **Order Management*	- > **My Orders*	- in the left-side navigation pane to find and pay for the order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// Instance ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The purchase duration. Unit: days. Valid values: **1**, **2**, **3**, **7**, **14**, **30**, **60**, **90**, **180**, **365**, **730**, **1095**, and **1825**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	OrderTimeLength      *string `json:"OrderTimeLength,omitempty" xml:"OrderTimeLength,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The source of the API call. This parameter is used only for internal maintenance and does not need to be specified.
	//
	// example:
	//
	// SDK
	SourceBiz *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
}

func (s RenewAdditionalBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAdditionalBandwidthRequest) GoString() string {
	return s.String()
}

func (s *RenewAdditionalBandwidthRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewAdditionalBandwidthRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *RenewAdditionalBandwidthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewAdditionalBandwidthRequest) GetOrderTimeLength() *string {
	return s.OrderTimeLength
}

func (s *RenewAdditionalBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RenewAdditionalBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewAdditionalBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewAdditionalBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewAdditionalBandwidthRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RenewAdditionalBandwidthRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *RenewAdditionalBandwidthRequest) SetAutoPay(v bool) *RenewAdditionalBandwidthRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetCouponNo(v string) *RenewAdditionalBandwidthRequest {
	s.CouponNo = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetInstanceId(v string) *RenewAdditionalBandwidthRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetOrderTimeLength(v string) *RenewAdditionalBandwidthRequest {
	s.OrderTimeLength = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetOwnerAccount(v string) *RenewAdditionalBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetOwnerId(v int64) *RenewAdditionalBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetResourceOwnerAccount(v string) *RenewAdditionalBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetResourceOwnerId(v int64) *RenewAdditionalBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetSecurityToken(v string) *RenewAdditionalBandwidthRequest {
	s.SecurityToken = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) SetSourceBiz(v string) *RenewAdditionalBandwidthRequest {
	s.SourceBiz = &v
	return s
}

func (s *RenewAdditionalBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
