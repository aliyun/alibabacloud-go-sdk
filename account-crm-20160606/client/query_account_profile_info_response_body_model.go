// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountProfileInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProfileInfo(v *QueryAccountProfileInfoResponseBodyProfileInfo) *QueryAccountProfileInfoResponseBody
	GetProfileInfo() *QueryAccountProfileInfoResponseBodyProfileInfo
	SetRequestId(v string) *QueryAccountProfileInfoResponseBody
	GetRequestId() *string
}

type QueryAccountProfileInfoResponseBody struct {
	ProfileInfo *QueryAccountProfileInfoResponseBodyProfileInfo `json:"ProfileInfo,omitempty" xml:"ProfileInfo,omitempty" type:"Struct"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccountProfileInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountProfileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountProfileInfoResponseBody) GetProfileInfo() *QueryAccountProfileInfoResponseBodyProfileInfo {
	return s.ProfileInfo
}

func (s *QueryAccountProfileInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountProfileInfoResponseBody) SetProfileInfo(v *QueryAccountProfileInfoResponseBodyProfileInfo) *QueryAccountProfileInfoResponseBody {
	s.ProfileInfo = v
	return s
}

func (s *QueryAccountProfileInfoResponseBody) SetRequestId(v string) *QueryAccountProfileInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBody) Validate() error {
	if s.ProfileInfo != nil {
		if err := s.ProfileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountProfileInfoResponseBodyProfileInfo struct {
	AccountAttr                    *string                                                 `json:"AccountAttr,omitempty" xml:"AccountAttr,omitempty"`
	AccountCertifyType             *string                                                 `json:"AccountCertifyType,omitempty" xml:"AccountCertifyType,omitempty"`
	ActiveNotSetMobile             *string                                                 `json:"ActiveNotSetMobile,omitempty" xml:"ActiveNotSetMobile,omitempty"`
	Address                        *string                                                 `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2                       *string                                                 `json:"Address2,omitempty" xml:"Address2,omitempty"`
	Address3                       *string                                                 `json:"Address3,omitempty" xml:"Address3,omitempty"`
	Address4                       *string                                                 `json:"Address4,omitempty" xml:"Address4,omitempty"`
	Address5                       *string                                                 `json:"Address5,omitempty" xml:"Address5,omitempty"`
	Address6                       *string                                                 `json:"Address6,omitempty" xml:"Address6,omitempty"`
	AlipayAccount                  *string                                                 `json:"AlipayAccount,omitempty" xml:"AlipayAccount,omitempty"`
	AlipayUid                      *string                                                 `json:"AlipayUid,omitempty" xml:"AlipayUid,omitempty"`
	AliyunID                       *string                                                 `json:"AliyunID,omitempty" xml:"AliyunID,omitempty"`
	AliyunPK                       *string                                                 `json:"AliyunPK,omitempty" xml:"AliyunPK,omitempty"`
	AuthAlipay                     *string                                                 `json:"AuthAlipay,omitempty" xml:"AuthAlipay,omitempty"`
	AuthDomainUserId               *string                                                 `json:"AuthDomainUserId,omitempty" xml:"AuthDomainUserId,omitempty"`
	B2bhid                         *string                                                 `json:"B2bhid,omitempty" xml:"B2bhid,omitempty"`
	BankId                         *string                                                 `json:"BankId,omitempty" xml:"BankId,omitempty"`
	BankName                       *string                                                 `json:"BankName,omitempty" xml:"BankName,omitempty"`
	BeiAnAuthCId                   *string                                                 `json:"BeiAnAuthCId,omitempty" xml:"BeiAnAuthCId,omitempty"`
	BeiAnIcpNumber                 *string                                                 `json:"BeiAnIcpNumber,omitempty" xml:"BeiAnIcpNumber,omitempty"`
	BeiAnMobile                    *string                                                 `json:"BeiAnMobile,omitempty" xml:"BeiAnMobile,omitempty"`
	BindAlipayNo                   *string                                                 `json:"BindAlipayNo,omitempty" xml:"BindAlipayNo,omitempty"`
	CertType                       *string                                                 `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertifiedFrom                  *string                                                 `json:"CertifiedFrom,omitempty" xml:"CertifiedFrom,omitempty"`
	CertifiedTime                  *string                                                 `json:"CertifiedTime,omitempty" xml:"CertifiedTime,omitempty"`
	City                           *QueryAccountProfileInfoResponseBodyProfileInfoCity     `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	ContactMethod                  *string                                                 `json:"ContactMethod,omitempty" xml:"ContactMethod,omitempty"`
	CreateTime                     *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	District                       *QueryAccountProfileInfoResponseBodyProfileInfoDistrict `json:"District,omitempty" xml:"District,omitempty" type:"Struct"`
	Eid                            *string                                                 `json:"Eid,omitempty" xml:"Eid,omitempty"`
	Email                          *string                                                 `json:"Email,omitempty" xml:"Email,omitempty"`
	Fax                            *string                                                 `json:"Fax,omitempty" xml:"Fax,omitempty"`
	FirstName                      *string                                                 `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	Fyl                            *string                                                 `json:"Fyl,omitempty" xml:"Fyl,omitempty"`
	HavanaId                       *string                                                 `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	Head                           *string                                                 `json:"Head,omitempty" xml:"Head,omitempty"`
	HeadUrl                        *string                                                 `json:"HeadUrl,omitempty" xml:"HeadUrl,omitempty"`
	IDNumber                       *string                                                 `json:"IDNumber,omitempty" xml:"IDNumber,omitempty"`
	IsBankIDAuth                   *string                                                 `json:"IsBankIDAuth,omitempty" xml:"IsBankIDAuth,omitempty"`
	IsCertified                    *string                                                 `json:"IsCertified,omitempty" xml:"IsCertified,omitempty"`
	LastName                       *string                                                 `json:"LastName,omitempty" xml:"LastName,omitempty"`
	Mobile                         *string                                                 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	NationalityCode                *string                                                 `json:"NationalityCode,omitempty" xml:"NationalityCode,omitempty"`
	NickName                       *string                                                 `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Own                            *string                                                 `json:"Own,omitempty" xml:"Own,omitempty"`
	Phone                          *string                                                 `json:"Phone,omitempty" xml:"Phone,omitempty"`
	PostCode                       *string                                                 `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	PreferredLanguage              *string                                                 `json:"PreferredLanguage,omitempty" xml:"PreferredLanguage,omitempty"`
	ProcessingEnterpriseCertify    *bool                                                   `json:"ProcessingEnterpriseCertify,omitempty" xml:"ProcessingEnterpriseCertify,omitempty"`
	Province                       *QueryAccountProfileInfoResponseBodyProfileInfoProvince `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	RegisterIP                     *string                                                 `json:"RegisterIP,omitempty" xml:"RegisterIP,omitempty"`
	SecurityMobile                 *string                                                 `json:"SecurityMobile,omitempty" xml:"SecurityMobile,omitempty"`
	SecurityQuestionExists         *bool                                                   `json:"SecurityQuestionExists,omitempty" xml:"SecurityQuestionExists,omitempty"`
	SelfServicingBusinessRegNum    *string                                                 `json:"SelfServicingBusinessRegNum,omitempty" xml:"SelfServicingBusinessRegNum,omitempty"`
	SelfServicingIdentificationNum *string                                                 `json:"SelfServicingIdentificationNum,omitempty" xml:"SelfServicingIdentificationNum,omitempty"`
	ShowNickName                   *string                                                 `json:"ShowNickName,omitempty" xml:"ShowNickName,omitempty"`
	Src                            *string                                                 `json:"Src,omitempty" xml:"Src,omitempty"`
	TaobaoAccount                  *string                                                 `json:"TaobaoAccount,omitempty" xml:"TaobaoAccount,omitempty"`
	TaobaoNickFromHavana           *string                                                 `json:"TaobaoNickFromHavana,omitempty" xml:"TaobaoNickFromHavana,omitempty"`
	Tbhid                          *string                                                 `json:"Tbhid,omitempty" xml:"Tbhid,omitempty"`
	TrueName                       *string                                                 `json:"TrueName,omitempty" xml:"TrueName,omitempty"`
	UpdateTime                     *string                                                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	YahooEmail                     *string                                                 `json:"YahooEmail,omitempty" xml:"YahooEmail,omitempty"`
}

func (s QueryAccountProfileInfoResponseBodyProfileInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountProfileInfoResponseBodyProfileInfo) GoString() string {
	return s.String()
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAccountAttr() *string {
	return s.AccountAttr
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAccountCertifyType() *string {
	return s.AccountCertifyType
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetActiveNotSetMobile() *string {
	return s.ActiveNotSetMobile
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAddress() *string {
	return s.Address
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAddress2() *string {
	return s.Address2
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAddress3() *string {
	return s.Address3
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAddress4() *string {
	return s.Address4
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAddress5() *string {
	return s.Address5
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAddress6() *string {
	return s.Address6
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAlipayAccount() *string {
	return s.AlipayAccount
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAlipayUid() *string {
	return s.AlipayUid
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAliyunID() *string {
	return s.AliyunID
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAliyunPK() *string {
	return s.AliyunPK
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAuthAlipay() *string {
	return s.AuthAlipay
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetAuthDomainUserId() *string {
	return s.AuthDomainUserId
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetB2bhid() *string {
	return s.B2bhid
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetBankId() *string {
	return s.BankId
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetBankName() *string {
	return s.BankName
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetBeiAnAuthCId() *string {
	return s.BeiAnAuthCId
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetBeiAnIcpNumber() *string {
	return s.BeiAnIcpNumber
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetBeiAnMobile() *string {
	return s.BeiAnMobile
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetBindAlipayNo() *string {
	return s.BindAlipayNo
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetCertType() *string {
	return s.CertType
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetCertifiedFrom() *string {
	return s.CertifiedFrom
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetCertifiedTime() *string {
	return s.CertifiedTime
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetCity() *QueryAccountProfileInfoResponseBodyProfileInfoCity {
	return s.City
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetContactMethod() *string {
	return s.ContactMethod
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetDistrict() *QueryAccountProfileInfoResponseBodyProfileInfoDistrict {
	return s.District
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetEid() *string {
	return s.Eid
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetEmail() *string {
	return s.Email
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetFax() *string {
	return s.Fax
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetFirstName() *string {
	return s.FirstName
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetFyl() *string {
	return s.Fyl
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetHead() *string {
	return s.Head
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetHeadUrl() *string {
	return s.HeadUrl
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetIDNumber() *string {
	return s.IDNumber
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetIsBankIDAuth() *string {
	return s.IsBankIDAuth
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetIsCertified() *string {
	return s.IsCertified
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetLastName() *string {
	return s.LastName
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetMobile() *string {
	return s.Mobile
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetNickName() *string {
	return s.NickName
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetOwn() *string {
	return s.Own
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetPhone() *string {
	return s.Phone
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetPostCode() *string {
	return s.PostCode
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetPreferredLanguage() *string {
	return s.PreferredLanguage
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetProcessingEnterpriseCertify() *bool {
	return s.ProcessingEnterpriseCertify
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetProvince() *QueryAccountProfileInfoResponseBodyProfileInfoProvince {
	return s.Province
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetRegisterIP() *string {
	return s.RegisterIP
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetSecurityMobile() *string {
	return s.SecurityMobile
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetSecurityQuestionExists() *bool {
	return s.SecurityQuestionExists
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetSelfServicingBusinessRegNum() *string {
	return s.SelfServicingBusinessRegNum
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetSelfServicingIdentificationNum() *string {
	return s.SelfServicingIdentificationNum
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetShowNickName() *string {
	return s.ShowNickName
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetSrc() *string {
	return s.Src
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetTaobaoAccount() *string {
	return s.TaobaoAccount
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetTaobaoNickFromHavana() *string {
	return s.TaobaoNickFromHavana
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetTbhid() *string {
	return s.Tbhid
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetTrueName() *string {
	return s.TrueName
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) GetYahooEmail() *string {
	return s.YahooEmail
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAccountAttr(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.AccountAttr = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAccountCertifyType(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.AccountCertifyType = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetActiveNotSetMobile(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.ActiveNotSetMobile = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAddress(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Address = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAddress2(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Address2 = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAddress3(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Address3 = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAddress4(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Address4 = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAddress5(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Address5 = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAddress6(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Address6 = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAlipayAccount(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.AlipayAccount = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAlipayUid(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.AlipayUid = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAliyunID(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.AliyunID = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAliyunPK(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.AliyunPK = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAuthAlipay(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.AuthAlipay = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetAuthDomainUserId(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.AuthDomainUserId = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetB2bhid(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.B2bhid = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetBankId(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.BankId = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetBankName(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.BankName = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetBeiAnAuthCId(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.BeiAnAuthCId = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetBeiAnIcpNumber(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.BeiAnIcpNumber = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetBeiAnMobile(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.BeiAnMobile = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetBindAlipayNo(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.BindAlipayNo = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetCertType(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.CertType = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetCertifiedFrom(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.CertifiedFrom = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetCertifiedTime(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.CertifiedTime = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetCity(v *QueryAccountProfileInfoResponseBodyProfileInfoCity) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.City = v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetContactMethod(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.ContactMethod = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetCreateTime(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.CreateTime = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetDistrict(v *QueryAccountProfileInfoResponseBodyProfileInfoDistrict) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.District = v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetEid(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Eid = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetEmail(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Email = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetFax(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Fax = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetFirstName(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.FirstName = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetFyl(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Fyl = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetHavanaId(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.HavanaId = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetHead(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Head = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetHeadUrl(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.HeadUrl = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetIDNumber(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.IDNumber = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetIsBankIDAuth(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.IsBankIDAuth = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetIsCertified(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.IsCertified = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetLastName(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.LastName = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetMobile(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Mobile = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetNationalityCode(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.NationalityCode = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetNickName(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.NickName = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetOwn(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Own = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetPhone(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Phone = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetPostCode(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.PostCode = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetPreferredLanguage(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.PreferredLanguage = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetProcessingEnterpriseCertify(v bool) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.ProcessingEnterpriseCertify = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetProvince(v *QueryAccountProfileInfoResponseBodyProfileInfoProvince) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Province = v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetRegisterIP(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.RegisterIP = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetSecurityMobile(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.SecurityMobile = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetSecurityQuestionExists(v bool) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.SecurityQuestionExists = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetSelfServicingBusinessRegNum(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.SelfServicingBusinessRegNum = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetSelfServicingIdentificationNum(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.SelfServicingIdentificationNum = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetShowNickName(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.ShowNickName = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetSrc(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Src = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetTaobaoAccount(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.TaobaoAccount = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetTaobaoNickFromHavana(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.TaobaoNickFromHavana = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetTbhid(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.Tbhid = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetTrueName(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.TrueName = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetUpdateTime(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.UpdateTime = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) SetYahooEmail(v string) *QueryAccountProfileInfoResponseBodyProfileInfo {
	s.YahooEmail = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfo) Validate() error {
	if s.City != nil {
		if err := s.City.Validate(); err != nil {
			return err
		}
	}
	if s.District != nil {
		if err := s.District.Validate(); err != nil {
			return err
		}
	}
	if s.Province != nil {
		if err := s.Province.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountProfileInfoResponseBodyProfileInfoCity struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAccountProfileInfoResponseBodyProfileInfoCity) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountProfileInfoResponseBodyProfileInfoCity) GoString() string {
	return s.String()
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoCity) GetId() *string {
	return s.Id
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoCity) GetName() *string {
	return s.Name
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoCity) SetId(v string) *QueryAccountProfileInfoResponseBodyProfileInfoCity {
	s.Id = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoCity) SetName(v string) *QueryAccountProfileInfoResponseBodyProfileInfoCity {
	s.Name = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoCity) Validate() error {
	return dara.Validate(s)
}

type QueryAccountProfileInfoResponseBodyProfileInfoDistrict struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAccountProfileInfoResponseBodyProfileInfoDistrict) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountProfileInfoResponseBodyProfileInfoDistrict) GoString() string {
	return s.String()
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoDistrict) GetId() *string {
	return s.Id
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoDistrict) GetName() *string {
	return s.Name
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoDistrict) SetId(v string) *QueryAccountProfileInfoResponseBodyProfileInfoDistrict {
	s.Id = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoDistrict) SetName(v string) *QueryAccountProfileInfoResponseBodyProfileInfoDistrict {
	s.Name = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoDistrict) Validate() error {
	return dara.Validate(s)
}

type QueryAccountProfileInfoResponseBodyProfileInfoProvince struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAccountProfileInfoResponseBodyProfileInfoProvince) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountProfileInfoResponseBodyProfileInfoProvince) GoString() string {
	return s.String()
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoProvince) GetId() *string {
	return s.Id
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoProvince) GetName() *string {
	return s.Name
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoProvince) SetId(v string) *QueryAccountProfileInfoResponseBodyProfileInfoProvince {
	s.Id = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoProvince) SetName(v string) *QueryAccountProfileInfoResponseBodyProfileInfoProvince {
	s.Name = &v
	return s
}

func (s *QueryAccountProfileInfoResponseBodyProfileInfoProvince) Validate() error {
	return dara.Validate(s)
}
