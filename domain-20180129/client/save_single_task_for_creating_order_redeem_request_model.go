// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderRedeemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponNo(v string) *SaveSingleTaskForCreatingOrderRedeemRequest
	GetCouponNo() *string
	SetCurrentExpirationDate(v int64) *SaveSingleTaskForCreatingOrderRedeemRequest
	GetCurrentExpirationDate() *int64
	SetDomainName(v string) *SaveSingleTaskForCreatingOrderRedeemRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForCreatingOrderRedeemRequest
	GetLang() *string
	SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderRedeemRequest
	GetPromotionNo() *string
	SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderRedeemRequest
	GetUseCoupon() *bool
	SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderRedeemRequest
	GetUsePromotion() *bool
	SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderRedeemRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForCreatingOrderRedeemRequest struct {
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
	// 123123
	PromotionNo *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
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

func (s SaveSingleTaskForCreatingOrderRedeemRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRedeemRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) GetCurrentExpirationDate() *int64 {
	return s.CurrentExpirationDate
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) GetUsePromotion() *bool {
	return s.UsePromotion
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetCouponNo(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetCurrentExpirationDate(v int64) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.CurrentExpirationDate = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetDomainName(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetLang(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) Validate() error {
	return dara.Validate(s)
}
