// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewDBInstanceRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewDBInstanceRequest
	GetAutoRenew() *bool
	SetBusinessInfo(v string) *RenewDBInstanceRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *RenewDBInstanceRequest
	GetClientToken() *string
	SetCouponNo(v string) *RenewDBInstanceRequest
	GetCouponNo() *string
	SetDBInstanceId(v string) *RenewDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *RenewDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RenewDBInstanceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *RenewDBInstanceRequest
	GetPeriod() *int32
	SetResourceOwnerAccount(v string) *RenewDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewDBInstanceRequest
	GetResourceOwnerId() *int64
}

type RenewDBInstanceRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: enables automatic payment. Make sure that you have sufficient balance within your account.
	//
	// 	- **false**: disables automatic payment. You must perform the following operations to pay for the instance: Log on to the ApsaraDB for MongoDB console. In the upper-right corner, click **Expenses*	- and select **User Center*	- from the drop-down list. The User Center page appears. In the left-side navigation pane, choose **Order Management*	- > Renew. On the Renewal tab, find the bill that you want to pay and then click Renew in the Actions column.
	//
	// Default value: **true**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The business information.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to use coupons. Default value: null. Valid values:
	//
	// 	- **default*	- or **null**: uses coupons.
	//
	// 	- **youhuiquan_promotion_option_id_for_blank**: does not use coupons.
	//
	// example:
	//
	// 1111111111111111
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription period of the instance. Unit: month. Valid values: **1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, and 36**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewDBInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewDBInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewDBInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *RenewDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewDBInstanceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *RenewDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RenewDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RenewDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewDBInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewDBInstanceRequest) SetAutoPay(v bool) *RenewDBInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewDBInstanceRequest) SetAutoRenew(v bool) *RenewDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewDBInstanceRequest) SetBusinessInfo(v string) *RenewDBInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *RenewDBInstanceRequest) SetClientToken(v string) *RenewDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewDBInstanceRequest) SetCouponNo(v string) *RenewDBInstanceRequest {
	s.CouponNo = &v
	return s
}

func (s *RenewDBInstanceRequest) SetDBInstanceId(v string) *RenewDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RenewDBInstanceRequest) SetOwnerAccount(v string) *RenewDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewDBInstanceRequest) SetOwnerId(v int64) *RenewDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewDBInstanceRequest) SetPeriod(v int32) *RenewDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewDBInstanceRequest) SetResourceOwnerAccount(v string) *RenewDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewDBInstanceRequest) SetResourceOwnerId(v int64) *RenewDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
