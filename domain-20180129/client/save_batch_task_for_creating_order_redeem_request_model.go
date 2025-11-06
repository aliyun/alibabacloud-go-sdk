// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderRedeemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponNo(v string) *SaveBatchTaskForCreatingOrderRedeemRequest
	GetCouponNo() *string
	SetLang(v string) *SaveBatchTaskForCreatingOrderRedeemRequest
	GetLang() *string
	SetOrderRedeemParam(v []*SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) *SaveBatchTaskForCreatingOrderRedeemRequest
	GetOrderRedeemParam() []*SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam
	SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderRedeemRequest
	GetPromotionNo() *string
	SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderRedeemRequest
	GetUseCoupon() *bool
	SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderRedeemRequest
	GetUsePromotion() *bool
	SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderRedeemRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForCreatingOrderRedeemRequest struct {
	// example:
	//
	// 123123
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	OrderRedeemParam []*SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam `json:"OrderRedeemParam,omitempty" xml:"OrderRedeemParam,omitempty" type:"Repeated"`
	// example:
	//
	// 123213123
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

func (s SaveBatchTaskForCreatingOrderRedeemRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRedeemRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) GetOrderRedeemParam() []*SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam {
	return s.OrderRedeemParam
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) GetUsePromotion() *bool {
	return s.UsePromotion
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetCouponNo(v string) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetLang(v string) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetOrderRedeemParam(v []*SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.OrderRedeemParam = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) Validate() error {
	if s.OrderRedeemParam != nil {
		for _, item := range s.OrderRedeemParam {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam struct {
	// example:
	//
	// 000000
	CurrentExpirationDate *int64 `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	// example:
	//
	// Aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) GetCurrentExpirationDate() *int64 {
	return s.CurrentExpirationDate
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) SetCurrentExpirationDate(v int64) *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam {
	s.CurrentExpirationDate = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) SetDomainName(v string) *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) Validate() error {
	return dara.Validate(s)
}
