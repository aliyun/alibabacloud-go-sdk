// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountRealNameInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProfileInfo(v *QueryAccountRealNameInfoResponseBodyProfileInfo) *QueryAccountRealNameInfoResponseBody
	GetProfileInfo() *QueryAccountRealNameInfoResponseBodyProfileInfo
	SetRequestId(v string) *QueryAccountRealNameInfoResponseBody
	GetRequestId() *string
}

type QueryAccountRealNameInfoResponseBody struct {
	ProfileInfo *QueryAccountRealNameInfoResponseBodyProfileInfo `json:"ProfileInfo,omitempty" xml:"ProfileInfo,omitempty" type:"Struct"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccountRealNameInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountRealNameInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountRealNameInfoResponseBody) GetProfileInfo() *QueryAccountRealNameInfoResponseBodyProfileInfo {
	return s.ProfileInfo
}

func (s *QueryAccountRealNameInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountRealNameInfoResponseBody) SetProfileInfo(v *QueryAccountRealNameInfoResponseBodyProfileInfo) *QueryAccountRealNameInfoResponseBody {
	s.ProfileInfo = v
	return s
}

func (s *QueryAccountRealNameInfoResponseBody) SetRequestId(v string) *QueryAccountRealNameInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBody) Validate() error {
	if s.ProfileInfo != nil {
		if err := s.ProfileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountRealNameInfoResponseBodyProfileInfo struct {
	AccountCertifyType          *string `json:"AccountCertifyType,omitempty" xml:"AccountCertifyType,omitempty"`
	AuthAlipay                  *string `json:"AuthAlipay,omitempty" xml:"AuthAlipay,omitempty"`
	AuthBeiAnCid                *string `json:"AuthBeiAnCid,omitempty" xml:"AuthBeiAnCid,omitempty"`
	AuthDomain                  *string `json:"AuthDomain,omitempty" xml:"AuthDomain,omitempty"`
	CertifiedFrom               *string `json:"CertifiedFrom,omitempty" xml:"CertifiedFrom,omitempty"`
	CertifiedTime               *string `json:"CertifiedTime,omitempty" xml:"CertifiedTime,omitempty"`
	IsBankIDAuth                *string `json:"IsBankIDAuth,omitempty" xml:"IsBankIDAuth,omitempty"`
	IsCertified                 *string `json:"IsCertified,omitempty" xml:"IsCertified,omitempty"`
	LicenseNumber               *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	LicenseType                 *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	Name                        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessingEnterpriseCertify *bool   `json:"ProcessingEnterpriseCertify,omitempty" xml:"ProcessingEnterpriseCertify,omitempty"`
}

func (s QueryAccountRealNameInfoResponseBodyProfileInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountRealNameInfoResponseBodyProfileInfo) GoString() string {
	return s.String()
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetAccountCertifyType() *string {
	return s.AccountCertifyType
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetAuthAlipay() *string {
	return s.AuthAlipay
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetAuthBeiAnCid() *string {
	return s.AuthBeiAnCid
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetAuthDomain() *string {
	return s.AuthDomain
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetCertifiedFrom() *string {
	return s.CertifiedFrom
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetCertifiedTime() *string {
	return s.CertifiedTime
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetIsBankIDAuth() *string {
	return s.IsBankIDAuth
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetIsCertified() *string {
	return s.IsCertified
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetLicenseNumber() *string {
	return s.LicenseNumber
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetLicenseType() *string {
	return s.LicenseType
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetName() *string {
	return s.Name
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) GetProcessingEnterpriseCertify() *bool {
	return s.ProcessingEnterpriseCertify
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetAccountCertifyType(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.AccountCertifyType = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetAuthAlipay(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.AuthAlipay = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetAuthBeiAnCid(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.AuthBeiAnCid = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetAuthDomain(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.AuthDomain = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetCertifiedFrom(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.CertifiedFrom = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetCertifiedTime(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.CertifiedTime = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetIsBankIDAuth(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.IsBankIDAuth = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetIsCertified(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.IsCertified = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetLicenseNumber(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.LicenseNumber = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetLicenseType(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.LicenseType = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetName(v string) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.Name = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) SetProcessingEnterpriseCertify(v bool) *QueryAccountRealNameInfoResponseBodyProfileInfo {
	s.ProcessingEnterpriseCertify = &v
	return s
}

func (s *QueryAccountRealNameInfoResponseBodyProfileInfo) Validate() error {
	return dara.Validate(s)
}
