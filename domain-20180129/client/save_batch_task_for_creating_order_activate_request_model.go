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
	// The voucher ID.
	//
	// example:
	//
	// 123456
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The language of the error message returned by the API operation. Valid values:
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
	// The list of task details.
	//
	// This parameter is required.
	OrderActivateParam []*SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam `json:"OrderActivateParam,omitempty" xml:"OrderActivateParam,omitempty" type:"Repeated"`
	// The coupon ID.
	//
	// example:
	//
	// 123124
	PromotionNo *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	// Specifies whether to use a voucher.
	//
	// example:
	//
	// false
	UseCoupon *bool `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	// Specifies whether to use a coupon.
	//
	// example:
	//
	// false
	UsePromotion *bool `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	// The IP address of the user.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
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
	// The mailing address in English.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// chao yan qu **	- dasha **	- hao
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Specifies whether to use Alibaba Cloud DNS. Valid values: **true*	- and **false**. Default value: **true**.
	//
	// > - If this parameter is set to **true**, you do not need to specify the **OrderActivateParam.N.Dns1*	- and **OrderActivateParam.N.Dns2*	- parameters. Otherwise, the specified **OrderActivateParam.N.Dns1*	- and **OrderActivateParam.N.Dns2*	- parameters do not take effect.
	//
	// - If this parameter is set to **false**, you must also specify the **OrderActivateParam.N.Dns1*	- and **OrderActivateParam.N.Dns2*	- parameters.
	//
	// example:
	//
	// true
	AliyunDns *bool `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	// The city name in English.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// bei jing shi
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The country code. For example, **CN*	- represents China, and **US*	- represents the United States.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The custom DNS server 1.
	//
	// > - This parameter is available and required only when the **OrderActivateParam.N.AliyunDns*	- parameter is set to **false**.
	//
	// - Make sure that the custom DNS server is correct. Otherwise, the registration may fail.
	//
	// example:
	//
	// ns2.aliyun.com
	Dns1 *string `json:"Dns1,omitempty" xml:"Dns1,omitempty"`
	// The custom DNS server 2.
	//
	// > - This parameter is available and required only when the **OrderActivateParam.N.AliyunDns*	- parameter is set to **false**.
	//
	// - Make sure that the custom DNS server is correct. Otherwise, the registration may fail.
	//
	// example:
	//
	// ns1.aliyun.com
	Dns2 *string `json:"Dns2,omitempty" xml:"Dns2,omitempty"`
	// The domain name to be registered.
	//
	// > When you register a domain name, you must specify the domain name registrant information. Otherwise, the domain name registration fails. You can specify the domain name registrant information by using the OrderActivateParam.N.RegistrantProfileId parameter to associate a domain name registrant profile.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The email address.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Specifies whether to enable the domain name privacy protection service. Default value: **true**.
	//
	// example:
	//
	// true
	EnableDomainProxy *bool `json:"EnableDomainProxy,omitempty" xml:"EnableDomainProxy,omitempty"`
	// The domain name in Punycode format. This parameter can be left empty.
	//
	// example:
	//
	// xn--fiqs8s.com
	ExpectedPunycode *string `json:"ExpectedPunycode,omitempty" xml:"ExpectedPunycode,omitempty"`
	// Specifies whether to allow the registration of premium domain names. Default value: **false**.
	//
	// example:
	//
	// true
	PermitPremiumActivation *bool `json:"PermitPremiumActivation,omitempty" xml:"PermitPremiumActivation,omitempty"`
	// The postal code.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 102629
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// The province name in English.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// bei jing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The domain name contact in English.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// zhang san
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// The name of the domain name registrant in English.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// zhang san
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// The ID of the domain name registrant profile. The profile contains information such as the name of the domain name registrant, the domain name contact, the phone number, and the email address. You can only use the ID of a real-name verified domain name registrant profile to register a domain name. If you have created a domain name registrant profile, you can call the [QueryRegistrantProfiles](https://help.aliyun.com/document_detail/67701.html) operation to query the profile ID.
	//
	// > After you specify this parameter, you do not need to specify the **OrderActivateParam.N.RegistrantType**, **OrderActivateParam.N.ZhRegistrantOrganization**, **OrderActivateParam.N.ZhRegistrantName**, **OrderActivateParam.N.ZhProvince**, **OrderActivateParam.N.ZhCity**, **OrderActivateParam.N.ZhAddress**, **OrderActivateParam.N.RegistrantOrganization**, **OrderActivateParam.N.RegistrantName**, **OrderActivateParam.N.Province**, **OrderActivateParam.N.City**, **OrderActivateParam.N.Address**, **OrderActivateParam.N.PostalCode**, **OrderActivateParam.N.Country**, **OrderActivateParam.N.TelArea**, **OrderActivateParam.N.Telephone**, **OrderActivateParam.N.TelExt**, and **OrderActivateParam.N.Email*	- parameters.
	//
	// example:
	//
	// 000000
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// The type of the domain name registrant. Valid values:
	//
	// - **1**: Individual.
	//
	// - **2**: Enterprise or organization.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// The resource group ID.
	//
	// > If this parameter is not specified or the specified resource group ID does not exist, the default resource group ID is used.
	//
	// example:
	//
	// rg-XX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The subscription duration. Unit: **year**. Default value: **1**.
	//
	// example:
	//
	// 1
	SubscriptionDuration *int32 `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
	// The country code for the phone number. For example, the country code for China is **86**.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// The extension number.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 1234
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// The phone number.
	//
	// > This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 1820000****
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// Specifies whether to allow the registration of trademark terms.
	//
	// example:
	//
	// false
	TrademarkDomainActivation *bool `json:"TrademarkDomainActivation,omitempty" xml:"TrademarkDomainActivation,omitempty"`
	// The mailing address in Chinese.
	//
	// > This parameter is applicable only to the China site. This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 朝阳区***大厦***号
	ZhAddress *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	// The city name in Chinese.
	//
	// > This parameter is applicable only to the China site. This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 北京市
	ZhCity *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	// The province name in Chinese.
	//
	// > This parameter is applicable only to the China site. This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 北京
	ZhProvince *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	// The domain name contact in Chinese.
	//
	// > This parameter is applicable only to the China site. This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 张三
	ZhRegistrantName *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	// The name of the domain name registrant in Chinese.
	//
	// > This parameter is applicable only to the China site. This parameter is available and required only when the **OrderActivateParam.N.RegistrantProfileId*	- parameter is not specified. If this parameter is not specified, the domain name registration fails.
	//
	// example:
	//
	// 张三
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
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

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetExpectedPunycode() *string {
	return s.ExpectedPunycode
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

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetZhAddress() *string {
	return s.ZhAddress
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetZhCity() *string {
	return s.ZhCity
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetZhProvince() *string {
	return s.ZhProvince
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
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

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetExpectedPunycode(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ExpectedPunycode = &v
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

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetResourceGroupId(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ResourceGroupId = &v
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

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhAddress(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhAddress = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhCity(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhCity = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhProvince(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhProvince = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhRegistrantName(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhRegistrantOrganization(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) Validate() error {
	return dara.Validate(s)
}
