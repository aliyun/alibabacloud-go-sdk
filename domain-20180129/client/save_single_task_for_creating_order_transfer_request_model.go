// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationCode(v string) *SaveSingleTaskForCreatingOrderTransferRequest
	GetAuthorizationCode() *string
	SetCouponNo(v string) *SaveSingleTaskForCreatingOrderTransferRequest
	GetCouponNo() *string
	SetDomainName(v string) *SaveSingleTaskForCreatingOrderTransferRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForCreatingOrderTransferRequest
	GetLang() *string
	SetPermitPremiumTransfer(v bool) *SaveSingleTaskForCreatingOrderTransferRequest
	GetPermitPremiumTransfer() *bool
	SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderTransferRequest
	GetPromotionNo() *string
	SetRegistrantProfileId(v int64) *SaveSingleTaskForCreatingOrderTransferRequest
	GetRegistrantProfileId() *int64
	SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderTransferRequest
	GetUseCoupon() *bool
	SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderTransferRequest
	GetUsePromotion() *bool
	SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderTransferRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForCreatingOrderTransferRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testCode
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// example:
	//
	// 123456
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
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
	// false
	PermitPremiumTransfer *bool `json:"PermitPremiumTransfer,omitempty" xml:"PermitPremiumTransfer,omitempty"`
	// example:
	//
	// 123456
	PromotionNo *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
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

func (s SaveSingleTaskForCreatingOrderTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderTransferRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetPermitPremiumTransfer() *bool {
	return s.PermitPremiumTransfer
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetUsePromotion() *bool {
	return s.UsePromotion
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetAuthorizationCode(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.AuthorizationCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetCouponNo(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetDomainName(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetLang(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetPermitPremiumTransfer(v bool) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.PermitPremiumTransfer = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetRegistrantProfileId(v int64) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) Validate() error {
	return dara.Validate(s)
}
