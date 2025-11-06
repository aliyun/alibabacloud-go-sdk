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
	// example:
	//
	// 123123
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0000
	CurrentExpirationDate *int64 `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 123132
	PromotionNo *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SubscriptionDuration *int32 `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
	// example:
	//
	// false
	UseCoupon *bool `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	// example:
	//
	// false
	UsePromotion *bool `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
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
