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
	// example:
	//
	// 12312412
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	OrderRenewParam []*SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam `json:"OrderRenewParam,omitempty" xml:"OrderRenewParam,omitempty" type:"Repeated"`
	// example:
	//
	// 123123123
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
	// example:
	//
	// 1522080000000
	CurrentExpirationDate *int64 `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	// example:
	//
	// Aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
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

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) SetSubscriptionDuration(v int32) *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	s.SubscriptionDuration = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) Validate() error {
	return dara.Validate(s)
}
