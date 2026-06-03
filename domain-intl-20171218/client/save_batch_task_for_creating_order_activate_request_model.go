// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderActivateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponNo(v string) *SaveBatchTaskForCreatingOrderActivateRequest
	GetCouponNo() *string
	SetLang(v string) *SaveBatchTaskForCreatingOrderActivateRequest
	GetLang() *string
	SetOrderActivateParam(v []*SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) *SaveBatchTaskForCreatingOrderActivateRequest
	GetOrderActivateParam() []*SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam
	SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderActivateRequest
	GetPromotionNo() *string
	SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderActivateRequest
	GetUseCoupon() *bool
	SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderActivateRequest
	GetUsePromotion() *bool
	SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderActivateRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForCreatingOrderActivateRequest struct {
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	OrderActivateParam []*SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam `json:"OrderActivateParam,omitempty" xml:"OrderActivateParam,omitempty" type:"Repeated"`
	PromotionNo        *string                                                           `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	UseCoupon          *bool                                                             `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion       *bool                                                             `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp       *string                                                           `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderActivateRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderActivateRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) GetOrderActivateParam() []*SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	return s.OrderActivateParam
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) GetUsePromotion() *bool {
	return s.UsePromotion
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetCouponNo(v string) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetLang(v string) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetOrderActivateParam(v []*SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.OrderActivateParam = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) Validate() error {
	if s.OrderActivateParam != nil {
		for _, item := range s.OrderActivateParam {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam struct {
	Address                   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AliyunDns                 *bool   `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	City                      *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                   *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Dns1                      *string `json:"Dns1,omitempty" xml:"Dns1,omitempty"`
	Dns2                      *string `json:"Dns2,omitempty" xml:"Dns2,omitempty"`
	DomainName                *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Email                     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EnableDomainProxy         *bool   `json:"EnableDomainProxy,omitempty" xml:"EnableDomainProxy,omitempty"`
	PermitPremiumActivation   *bool   `json:"PermitPremiumActivation,omitempty" xml:"PermitPremiumActivation,omitempty"`
	PostalCode                *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                  *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName            *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization    *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId       *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantType            *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	SubscriptionDuration      *int32  `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
	TelArea                   *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                    *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                 *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	TrademarkDomainActivation *bool   `json:"TrademarkDomainActivation,omitempty" xml:"TrademarkDomainActivation,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetAddress() *string {
	return s.Address
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetAliyunDns() *bool {
	return s.AliyunDns
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetCity() *string {
	return s.City
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetCountry() *string {
	return s.Country
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetDns1() *string {
	return s.Dns1
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetDns2() *string {
	return s.Dns2
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetEmail() *string {
	return s.Email
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetEnableDomainProxy() *bool {
	return s.EnableDomainProxy
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetPermitPremiumActivation() *bool {
	return s.PermitPremiumActivation
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetPostalCode() *string {
	return s.PostalCode
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetProvince() *string {
	return s.Province
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetSubscriptionDuration() *int32 {
	return s.SubscriptionDuration
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetTelArea() *string {
	return s.TelArea
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetTelExt() *string {
	return s.TelExt
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetTelephone() *string {
	return s.Telephone
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetTrademarkDomainActivation() *bool {
	return s.TrademarkDomainActivation
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetAddress(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Address = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetAliyunDns(v bool) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.AliyunDns = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetCity(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.City = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetCountry(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Country = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetDns1(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Dns1 = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetDns2(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Dns2 = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetDomainName(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetEmail(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Email = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetEnableDomainProxy(v bool) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.EnableDomainProxy = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetPermitPremiumActivation(v bool) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.PermitPremiumActivation = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetPostalCode(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.PostalCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetProvince(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Province = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetRegistrantName(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.RegistrantName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetRegistrantOrganization(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetRegistrantProfileId(v int64) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetRegistrantType(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.RegistrantType = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetSubscriptionDuration(v int32) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.SubscriptionDuration = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetTelArea(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.TelArea = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetTelExt(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.TelExt = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetTelephone(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Telephone = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetTrademarkDomainActivation(v bool) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.TrademarkDomainActivation = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) Validate() error {
	return dara.Validate(s)
}
