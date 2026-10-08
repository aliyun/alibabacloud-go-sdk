// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderRenewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponNo(v string) *SaveBatchTaskForCreatingOrderRenewRequest
	GetCouponNo() *string
	SetLang(v string) *SaveBatchTaskForCreatingOrderRenewRequest
	GetLang() *string
	SetOrderRenewParam(v []*SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) *SaveBatchTaskForCreatingOrderRenewRequest
	GetOrderRenewParam() []*SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam
	SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderRenewRequest
	GetPromotionNo() *string
	SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderRenewRequest
	GetUseCoupon() *bool
	SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderRenewRequest
	GetUsePromotion() *bool
	SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderRenewRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForCreatingOrderRenewRequest struct {
	// The coupon ID.
	//
	// example:
	//
	// 12312412
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The language of the error messages. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The parameters for each domain name to be renewed.
	//
	// This parameter is required.
	OrderRenewParam []*SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam `json:"OrderRenewParam,omitempty" xml:"OrderRenewParam,omitempty" type:"Repeated"`
	// The promotion ID.
	//
	// example:
	//
	// 123123123
	PromotionNo *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
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

func (s SaveBatchTaskForCreatingOrderRenewRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRenewRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) GetOrderRenewParam() []*SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	return s.OrderRenewParam
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) GetUsePromotion() *bool {
	return s.UsePromotion
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetCouponNo(v string) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetLang(v string) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetOrderRenewParam(v []*SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.OrderRenewParam = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) Validate() error {
	if s.OrderRenewParam != nil {
		for _, item := range s.OrderRenewParam {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam struct {
	// The current expiration date of the domain name, expressed in milliseconds since 00:00:00 UTC on January 1, 1970.
	//
	// example:
	//
	// 1522080000000
	CurrentExpirationDate *int64 `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	// The domain name that you want to renew. You can obtain a list of your domain names by calling the [QueryDomainList](https://help.aliyun.com/document_detail/67712.html) operation.
	//
	// example:
	//
	// Aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to allow the renewal of premium domain names. Default value: false.
	PermitPremiumRenew *bool `json:"PermitPremiumRenew,omitempty" xml:"PermitPremiumRenew,omitempty"`
	// The renewal duration, in years. Default value: **1**. Valid values: **1*	- to **10**.
	//
	// example:
	//
	// 1
	SubscriptionDuration *int32 `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) GetCurrentExpirationDate() *int64 {
	return s.CurrentExpirationDate
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) GetPermitPremiumRenew() *bool {
	return s.PermitPremiumRenew
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) GetSubscriptionDuration() *int32 {
	return s.SubscriptionDuration
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) SetCurrentExpirationDate(v int64) *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	s.CurrentExpirationDate = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) SetDomainName(v string) *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) SetPermitPremiumRenew(v bool) *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	s.PermitPremiumRenew = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) SetSubscriptionDuration(v int32) *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	s.SubscriptionDuration = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) Validate() error {
	return dara.Validate(s)
}
