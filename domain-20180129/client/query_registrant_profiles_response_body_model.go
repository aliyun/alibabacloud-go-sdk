// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegistrantProfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryRegistrantProfilesResponseBody
	GetCurrentPageNum() *int32
	SetNextPage(v bool) *QueryRegistrantProfilesResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryRegistrantProfilesResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryRegistrantProfilesResponseBody
	GetPrePage() *bool
	SetRegistrantProfiles(v *QueryRegistrantProfilesResponseBodyRegistrantProfiles) *QueryRegistrantProfilesResponseBody
	GetRegistrantProfiles() *QueryRegistrantProfilesResponseBodyRegistrantProfiles
	SetRequestId(v string) *QueryRegistrantProfilesResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryRegistrantProfilesResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryRegistrantProfilesResponseBody
	GetTotalPageNum() *int32
}

type QueryRegistrantProfilesResponseBody struct {
	// The page number returned.
	//
	// example:
	//
	// 1
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// Indicates whether the current page is followed by a page. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The number of entries returned on each page. Default value: **0**. Maximum value: **5000**.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the current page is preceded by a page. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// The list of registrant profiles.
	RegistrantProfiles *QueryRegistrantProfilesResponseBodyRegistrantProfiles `json:"RegistrantProfiles,omitempty" xml:"RegistrantProfiles,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 94053D79-7455-4F71-BF06-20EB2DEDE6BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// >  This parameter indicates the total number of queried registrant profiles. If multiple registrant profiles are queried, the information about these profiles is returned in sequence by profile.
	//
	// example:
	//
	// 9
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryRegistrantProfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryRegistrantProfilesResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryRegistrantProfilesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRegistrantProfilesResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryRegistrantProfilesResponseBody) GetRegistrantProfiles() *QueryRegistrantProfilesResponseBodyRegistrantProfiles {
	return s.RegistrantProfiles
}

func (s *QueryRegistrantProfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRegistrantProfilesResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryRegistrantProfilesResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryRegistrantProfilesResponseBody) SetCurrentPageNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetNextPage(v bool) *QueryRegistrantProfilesResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetPageSize(v int32) *QueryRegistrantProfilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetPrePage(v bool) *QueryRegistrantProfilesResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetRegistrantProfiles(v *QueryRegistrantProfilesResponseBodyRegistrantProfiles) *QueryRegistrantProfilesResponseBody {
	s.RegistrantProfiles = v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetRequestId(v string) *QueryRegistrantProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetTotalItemNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetTotalPageNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) Validate() error {
	if s.RegistrantProfiles != nil {
		if err := s.RegistrantProfiles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRegistrantProfilesResponseBodyRegistrantProfiles struct {
	RegistrantProfile []*QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile `json:"RegistrantProfile,omitempty" xml:"RegistrantProfile,omitempty" type:"Repeated"`
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfiles) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfiles) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfiles) GetRegistrantProfile() []*QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	return s.RegistrantProfile
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfiles) SetRegistrantProfile(v []*QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) *QueryRegistrantProfilesResponseBodyRegistrantProfiles {
	s.RegistrantProfile = v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfiles) Validate() error {
	if s.RegistrantProfile != nil {
		for _, item := range s.RegistrantProfile {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile struct {
	// The address of the domain name registrant.
	//
	// example:
	//
	// zhe jiang sheng hang zhou shi shi li qu shi li zhen shi li da sha 1001 hao
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The city where the domain name registrant is located, in English.
	//
	// example:
	//
	// hang zhou shi
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The code of the country or region where the domain name registrant is located, such as **CN*	- or **US**.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The time when the registrant profile was created.
	//
	// example:
	//
	// 2019-02-18 10:46:47
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The certificate number.
	//
	// example:
	//
	// 4****************1
	CredentialNo *string `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	// The certificate type.
	//
	// example:
	//
	// YYZZ
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// Indicates whether the template is the default template. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	DefaultRegistrantProfile *bool `json:"DefaultRegistrantProfile,omitempty" xml:"DefaultRegistrantProfile,omitempty"`
	// The email address of the domain name registrant.
	//
	// example:
	//
	// 82106****@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The status of the verification for the email address. Valid values:
	//
	// 	- **0**: not verified
	//
	// 	- **1**: verified
	//
	// example:
	//
	// 1
	EmailVerificationStatus *int32  `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	Params                  *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The postal code of the region where the domain name registrant is located.
	//
	// example:
	//
	// 310024
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// The province where the domain name registrant is located.
	//
	// example:
	//
	// zhe jiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The state of real-name verification for the domain name registrant. Valid values:
	//
	// 	- **FAILED**: Real-name verification for the domain name fails.
	//
	// 	- **SUCCEED**: Real-name verification for the domain name is successful.
	//
	// 	- **NONAUDIT**: Real-name verification for the domain name is not performed.
	//
	// 	- **AUDITING**: Real-name verification for the domain name is in progress.
	//
	// example:
	//
	// SUCCEED
	RealNameStatus *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	// The name of the domain name contact.
	//
	// example:
	//
	// li si
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// The name of the domain name registrant.
	//
	// example:
	//
	// li si
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// The ID of the registrant profile.
	//
	// example:
	//
	// 1000001
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// The type of the registrant profile. Valid values:
	//
	// 	- **common**: common profile.
	//
	// 	- **cnnic**: CNNIC profile.
	//
	// >  Only the Alibaba Cloud international site (alibabacloud.com) supports CNNIC profiles. To register domain names provided by CNNIC such as the .cn and . domain names on the Alibaba Cloud international site, you must use a CNNIC profile. To register other domain names, use a common profile.
	//
	// example:
	//
	// common
	RegistrantProfileType *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
	// The type of the domain name registrant. Valid values:
	//
	// 	- **1**: individual.
	//
	// 	- **2**: enterprise.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Test domain name
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The international dialing code of the country or region where the domain name contact is located. For example, the international dialing code of China is **86**.
	//
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// The extension of the phone number.
	//
	// example:
	//
	// 1234
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 1829756****
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	// The time when the registrant profile was updated.
	//
	// example:
	//
	// 2019-03-15 15:32:45
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The address of the domain name registrant, in Chinese.
	ZhAddress *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	// The city where the domain name registrant is located, in Chinese.
	ZhCity *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	// The province where the domain name registrant is located, in Chinese.
	ZhProvince *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	// The Chinese name of the domain name contact.
	ZhRegistrantName *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	// The Chinese name of the domain name registrant.
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetAddress() *string {
	return s.Address
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetCity() *string {
	return s.City
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetCountry() *string {
	return s.Country
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetCredentialNo() *string {
	return s.CredentialNo
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetCredentialType() *string {
	return s.CredentialType
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetDefaultRegistrantProfile() *bool {
	return s.DefaultRegistrantProfile
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetEmail() *string {
	return s.Email
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetEmailVerificationStatus() *int32 {
	return s.EmailVerificationStatus
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetParams() *string {
	return s.Params
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetPostalCode() *string {
	return s.PostalCode
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetProvince() *string {
	return s.Province
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRealNameStatus() *string {
	return s.RealNameStatus
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantProfileType() *string {
	return s.RegistrantProfileType
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRemark() *string {
	return s.Remark
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetTelArea() *string {
	return s.TelArea
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetTelExt() *string {
	return s.TelExt
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetTelephone() *string {
	return s.Telephone
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetZhAddress() *string {
	return s.ZhAddress
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetZhCity() *string {
	return s.ZhCity
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetZhProvince() *string {
	return s.ZhProvince
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetAddress(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Address = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCity(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.City = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCountry(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Country = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCreateTime(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.CreateTime = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCredentialNo(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.CredentialNo = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCredentialType(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.CredentialType = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetDefaultRegistrantProfile(v bool) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.DefaultRegistrantProfile = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetEmail(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Email = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetEmailVerificationStatus(v int32) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetParams(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Params = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetPostalCode(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.PostalCode = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetProvince(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Province = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRealNameStatus(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RealNameStatus = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantName(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantName = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantOrganization(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantProfileId(v int64) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantProfileType(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantProfileType = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantType(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantType = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRemark(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Remark = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelArea(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.TelArea = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelExt(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.TelExt = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelephone(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Telephone = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetUpdateTime(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.UpdateTime = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhAddress(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhAddress = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhCity(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhCity = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhProvince(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhProvince = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhRegistrantName(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhRegistrantName = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhRegistrantOrganization(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) Validate() error {
	return dara.Validate(s)
}
