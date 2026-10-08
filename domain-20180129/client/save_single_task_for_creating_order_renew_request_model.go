// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderRenewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponNo(v string) *SaveSingleTaskForCreatingOrderRenewRequest
	GetCouponNo() *string
	SetCurrentExpirationDate(v int64) *SaveSingleTaskForCreatingOrderRenewRequest
	GetCurrentExpirationDate() *int64
	SetDomainName(v string) *SaveSingleTaskForCreatingOrderRenewRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForCreatingOrderRenewRequest
	GetLang() *string
	SetPermitPremiumRenew(v bool) *SaveSingleTaskForCreatingOrderRenewRequest
	GetPermitPremiumRenew() *bool
	SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderRenewRequest
	GetPromotionNo() *string
	SetSubscriptionDuration(v int32) *SaveSingleTaskForCreatingOrderRenewRequest
	GetSubscriptionDuration() *int32
	SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderRenewRequest
	GetUseCoupon() *bool
	SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderRenewRequest
	GetUsePromotion() *bool
	SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderRenewRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForCreatingOrderRenewRequest struct {
	// The coupon number.
	//
	// example:
	//
	// 123123
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The current expiration date of the domain name. This value is a Unix timestamp in milliseconds, representing the time elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1522080000000
	CurrentExpirationDate *int64 `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	// The domain name to renew.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of error messages returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// The default value is **en**.
	//
	// example:
	//
	// en
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PermitPremiumRenew *bool   `json:"PermitPremiumRenew,omitempty" xml:"PermitPremiumRenew,omitempty"`
	// The promotion number.
	//
	// example:
	//
	// 123132
	PromotionNo *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	// The renewal period, in years. The value must be an integer from **1*	- to **10**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SubscriptionDuration *int32 `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **false**: Do not use a coupon.
	//
	// - **true**: Use a coupon.
	//
	// example:
	//
	// false
	UseCoupon *bool `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	// Specifies whether to use a promotion. Valid values:
	//
	// - **false**: Do not use a promotion.
	//
	// - **true**: Use a promotion.
	//
	// example:
	//
	// false
	UsePromotion *bool `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	// The user\\"s IP address. You can set this parameter to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRenewRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRenewRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetCurrentExpirationDate() *int64 {
	return s.CurrentExpirationDate
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetPermitPremiumRenew() *bool {
	return s.PermitPremiumRenew
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetSubscriptionDuration() *int32 {
	return s.SubscriptionDuration
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetUsePromotion() *bool {
	return s.UsePromotion
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetCouponNo(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetCurrentExpirationDate(v int64) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.CurrentExpirationDate = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetDomainName(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetLang(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetPermitPremiumRenew(v bool) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.PermitPremiumRenew = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetSubscriptionDuration(v int32) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.SubscriptionDuration = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) Validate() error {
	return dara.Validate(s)
}
