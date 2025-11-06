// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponNo(v string) *SaveBatchTaskForCreatingOrderTransferRequest
	GetCouponNo() *string
	SetLang(v string) *SaveBatchTaskForCreatingOrderTransferRequest
	GetLang() *string
	SetOrderTransferParam(v []*SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) *SaveBatchTaskForCreatingOrderTransferRequest
	GetOrderTransferParam() []*SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam
	SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderTransferRequest
	GetPromotionNo() *string
	SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderTransferRequest
	GetUseCoupon() *bool
	SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderTransferRequest
	GetUsePromotion() *bool
	SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderTransferRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForCreatingOrderTransferRequest struct {
	// example:
	//
	// 123123
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	OrderTransferParam []*SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam `json:"OrderTransferParam,omitempty" xml:"OrderTransferParam,omitempty" type:"Repeated"`
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

func (s SaveBatchTaskForCreatingOrderTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderTransferRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) GetOrderTransferParam() []*SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	return s.OrderTransferParam
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) GetUsePromotion() *bool {
	return s.UsePromotion
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetCouponNo(v string) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetLang(v string) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetOrderTransferParam(v []*SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.OrderTransferParam = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) Validate() error {
	if s.OrderTransferParam != nil {
		for _, item := range s.OrderTransferParam {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam struct {
	// example:
	//
	// testCode
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// false
	PermitPremiumTransfer *bool `json:"PermitPremiumTransfer,omitempty" xml:"PermitPremiumTransfer,omitempty"`
	// example:
	//
	// 123456
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) GetPermitPremiumTransfer() *bool {
	return s.PermitPremiumTransfer
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) SetAuthorizationCode(v string) *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	s.AuthorizationCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) SetDomainName(v string) *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) SetPermitPremiumTransfer(v bool) *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	s.PermitPremiumTransfer = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) SetRegistrantProfileId(v int64) *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) Validate() error {
	return dara.Validate(s)
}
