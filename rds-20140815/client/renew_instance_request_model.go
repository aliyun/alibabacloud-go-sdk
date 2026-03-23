// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v string) *RenewInstanceRequest
	GetAutoPay() *string
	SetAutoRenew(v string) *RenewInstanceRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *RenewInstanceRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *RenewInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RenewInstanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *RenewInstanceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *RenewInstanceRequest
	GetPeriod() *int32
	SetPromotionCode(v string) *RenewInstanceRequest
	GetPromotionCode() *string
	SetResourceOwnerAccount(v string) *RenewInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewInstanceRequest
	GetResourceOwnerId() *int64
}

type RenewInstanceRequest struct {
	AutoPay       *string `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew     *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetAutoPay() *string {
	return s.AutoPay
}

func (s *RenewInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *RenewInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *RenewInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RenewInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *RenewInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewInstanceRequest) SetAutoPay(v string) *RenewInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewInstanceRequest) SetAutoRenew(v string) *RenewInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewInstanceRequest) SetAutoUseCoupon(v bool) *RenewInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetDBInstanceId(v string) *RenewInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v int32) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetPromotionCode(v string) *RenewInstanceRequest {
	s.PromotionCode = &v
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

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
