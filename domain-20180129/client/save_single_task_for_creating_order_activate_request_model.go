// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderActivateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetAddress() *string
	SetAliyunDns(v bool) *SaveSingleTaskForCreatingOrderActivateRequest
	GetAliyunDns() *bool
	SetCity(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetCity() *string
	SetCountry(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetCountry() *string
	SetCouponNo(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetCouponNo() *string
	SetDns1(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetDns1() *string
	SetDns2(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetDns2() *string
	SetDomainName(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetDomainName() *string
	SetEmail(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetEmail() *string
	SetEnableDomainProxy(v bool) *SaveSingleTaskForCreatingOrderActivateRequest
	GetEnableDomainProxy() *bool
	SetLang(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetLang() *string
	SetPermitPremiumActivation(v bool) *SaveSingleTaskForCreatingOrderActivateRequest
	GetPermitPremiumActivation() *bool
	SetPostalCode(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetPostalCode() *string
	SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetPromotionNo() *string
	SetProvince(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetProvince() *string
	SetRegistrantName(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetRegistrantOrganization() *string
	SetRegistrantProfileId(v int64) *SaveSingleTaskForCreatingOrderActivateRequest
	GetRegistrantProfileId() *int64
	SetRegistrantType(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetRegistrantType() *string
	SetResourceGroupId(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetResourceGroupId() *string
	SetSubscriptionDuration(v int32) *SaveSingleTaskForCreatingOrderActivateRequest
	GetSubscriptionDuration() *int32
	SetTelArea(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetTelArea() *string
	SetTelExt(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetTelExt() *string
	SetTelephone(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetTelephone() *string
	SetTrademarkDomainActivation(v bool) *SaveSingleTaskForCreatingOrderActivateRequest
	GetTrademarkDomainActivation() *bool
	SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderActivateRequest
	GetUseCoupon() *bool
	SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderActivateRequest
	GetUsePromotion() *bool
	SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetUserClientIp() *string
	SetZhAddress(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetZhAddress() *string
	SetZhCity(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetZhCity() *string
	SetZhProvince(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetZhProvince() *string
	SetZhRegistrantName(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetZhRegistrantName() *string
	SetZhRegistrantOrganization(v string) *SaveSingleTaskForCreatingOrderActivateRequest
	GetZhRegistrantOrganization() *string
}

type SaveSingleTaskForCreatingOrderActivateRequest struct {
	// example:
	//
	// chao yang qu
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// true
	AliyunDns *bool `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	// example:
	//
	// bei jing shi
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// 123456
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// ns1.aliyun.com
	Dns1 *string `json:"Dns1,omitempty" xml:"Dns1,omitempty"`
	// example:
	//
	// ns2.aliyun.com
	Dns2 *string `json:"Dns2,omitempty" xml:"Dns2,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// false
	EnableDomainProxy *bool `json:"EnableDomainProxy,omitempty" xml:"EnableDomainProxy,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// false
	PermitPremiumActivation *bool `json:"PermitPremiumActivation,omitempty" xml:"PermitPremiumActivation,omitempty"`
	// example:
	//
	// 1234567
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// example:
	//
	// 123123
	PromotionNo *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	// example:
	//
	// bei jing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// ce shi
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// example:
	//
	// ce shi
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// example:
	//
	// 123
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// example:
	//
	// rg-XX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1
	SubscriptionDuration *int32 `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// example:
	//
	// 1234
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// example:
	//
	// 12345678
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// example:
	//
	// false
	TrademarkDomainActivation *bool `json:"TrademarkDomainActivation,omitempty" xml:"TrademarkDomainActivation,omitempty"`
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
	UserClientIp             *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhAddress                *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderActivateRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderActivateRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetAddress() *string {
	return s.Address
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetAliyunDns() *bool {
	return s.AliyunDns
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetCity() *string {
	return s.City
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetCountry() *string {
	return s.Country
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetDns1() *string {
	return s.Dns1
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetDns2() *string {
	return s.Dns2
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetEmail() *string {
	return s.Email
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetEnableDomainProxy() *bool {
	return s.EnableDomainProxy
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetPermitPremiumActivation() *bool {
	return s.PermitPremiumActivation
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetPostalCode() *string {
	return s.PostalCode
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetProvince() *string {
	return s.Province
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetSubscriptionDuration() *int32 {
	return s.SubscriptionDuration
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetTelArea() *string {
	return s.TelArea
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetTelExt() *string {
	return s.TelExt
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetTelephone() *string {
	return s.Telephone
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetTrademarkDomainActivation() *bool {
	return s.TrademarkDomainActivation
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetUseCoupon() *bool {
	return s.UseCoupon
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetUsePromotion() *bool {
	return s.UsePromotion
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetZhAddress() *string {
	return s.ZhAddress
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetZhCity() *string {
	return s.ZhCity
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetZhProvince() *string {
	return s.ZhProvince
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetAddress(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Address = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetAliyunDns(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.AliyunDns = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetCity(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.City = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetCountry(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Country = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetCouponNo(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetDns1(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Dns1 = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetDns2(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Dns2 = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetDomainName(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetEmail(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Email = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetEnableDomainProxy(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.EnableDomainProxy = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetLang(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetPermitPremiumActivation(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.PermitPremiumActivation = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetPostalCode(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetProvince(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Province = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetRegistrantName(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetRegistrantOrganization(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetRegistrantProfileId(v int64) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetRegistrantType(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetResourceGroupId(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetSubscriptionDuration(v int32) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.SubscriptionDuration = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetTelArea(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.TelArea = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetTelExt(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.TelExt = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetTelephone(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Telephone = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetTrademarkDomainActivation(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.TrademarkDomainActivation = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhAddress(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhCity(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhProvince(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhRegistrantName(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhRegistrantOrganization(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) Validate() error {
	return dara.Validate(s)
}
