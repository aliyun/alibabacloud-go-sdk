// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadRealNameInfoByPkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LoadRealNameInfoByPkResponseBody
	GetCode() *string
	SetData(v *LoadRealNameInfoByPkResponseBodyData) *LoadRealNameInfoByPkResponseBody
	GetData() *LoadRealNameInfoByPkResponseBodyData
	SetMsg(v string) *LoadRealNameInfoByPkResponseBody
	GetMsg() *string
	SetRequestId(v string) *LoadRealNameInfoByPkResponseBody
	GetRequestId() *string
}

type LoadRealNameInfoByPkResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *LoadRealNameInfoByPkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                               `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoadRealNameInfoByPkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoadRealNameInfoByPkResponseBody) GoString() string {
	return s.String()
}

func (s *LoadRealNameInfoByPkResponseBody) GetCode() *string {
	return s.Code
}

func (s *LoadRealNameInfoByPkResponseBody) GetData() *LoadRealNameInfoByPkResponseBodyData {
	return s.Data
}

func (s *LoadRealNameInfoByPkResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *LoadRealNameInfoByPkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoadRealNameInfoByPkResponseBody) SetCode(v string) *LoadRealNameInfoByPkResponseBody {
	s.Code = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBody) SetData(v *LoadRealNameInfoByPkResponseBodyData) *LoadRealNameInfoByPkResponseBody {
	s.Data = v
	return s
}

func (s *LoadRealNameInfoByPkResponseBody) SetMsg(v string) *LoadRealNameInfoByPkResponseBody {
	s.Msg = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBody) SetRequestId(v string) *LoadRealNameInfoByPkResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoadRealNameInfoByPkResponseBodyData struct {
	AccountCertifyType          *string `json:"AccountCertifyType,omitempty" xml:"AccountCertifyType,omitempty"`
	AuthAlipay                  *string `json:"AuthAlipay,omitempty" xml:"AuthAlipay,omitempty"`
	AuthAlipayDomain            *string `json:"AuthAlipayDomain,omitempty" xml:"AuthAlipayDomain,omitempty"`
	AuthAlipayLoginId           *string `json:"AuthAlipayLoginId,omitempty" xml:"AuthAlipayLoginId,omitempty"`
	AuthBeiAnCid                *string `json:"AuthBeiAnCid,omitempty" xml:"AuthBeiAnCid,omitempty"`
	AuthDomain                  *string `json:"AuthDomain,omitempty" xml:"AuthDomain,omitempty"`
	CertifiedFrom               *string `json:"CertifiedFrom,omitempty" xml:"CertifiedFrom,omitempty"`
	CertifiedTime               *string `json:"CertifiedTime,omitempty" xml:"CertifiedTime,omitempty"`
	CertifyStatus               *int32  `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	CicCertifyFrom              *int32  `json:"CicCertifyFrom,omitempty" xml:"CicCertifyFrom,omitempty"`
	CicCertifyProduct           *int64  `json:"CicCertifyProduct,omitempty" xml:"CicCertifyProduct,omitempty"`
	IsBankIDAuth                *bool   `json:"IsBankIDAuth,omitempty" xml:"IsBankIDAuth,omitempty"`
	IsCertified                 *bool   `json:"IsCertified,omitempty" xml:"IsCertified,omitempty"`
	LicenseNumber               *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	LicenseType                 *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	Name                        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewUnityRealNameAccount     *string `json:"NewUnityRealNameAccount,omitempty" xml:"NewUnityRealNameAccount,omitempty"`
	ProcessingEnterpriseCertify *bool   `json:"ProcessingEnterpriseCertify,omitempty" xml:"ProcessingEnterpriseCertify,omitempty"`
}

func (s LoadRealNameInfoByPkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LoadRealNameInfoByPkResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetAccountCertifyType() *string {
	return s.AccountCertifyType
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetAuthAlipay() *string {
	return s.AuthAlipay
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetAuthAlipayDomain() *string {
	return s.AuthAlipayDomain
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetAuthAlipayLoginId() *string {
	return s.AuthAlipayLoginId
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetAuthBeiAnCid() *string {
	return s.AuthBeiAnCid
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetAuthDomain() *string {
	return s.AuthDomain
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetCertifiedFrom() *string {
	return s.CertifiedFrom
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetCertifiedTime() *string {
	return s.CertifiedTime
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetCertifyStatus() *int32 {
	return s.CertifyStatus
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetCicCertifyFrom() *int32 {
	return s.CicCertifyFrom
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetCicCertifyProduct() *int64 {
	return s.CicCertifyProduct
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetIsBankIDAuth() *bool {
	return s.IsBankIDAuth
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetIsCertified() *bool {
	return s.IsCertified
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetLicenseNumber() *string {
	return s.LicenseNumber
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetLicenseType() *string {
	return s.LicenseType
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetName() *string {
	return s.Name
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetNewUnityRealNameAccount() *string {
	return s.NewUnityRealNameAccount
}

func (s *LoadRealNameInfoByPkResponseBodyData) GetProcessingEnterpriseCertify() *bool {
	return s.ProcessingEnterpriseCertify
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetAccountCertifyType(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.AccountCertifyType = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetAuthAlipay(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.AuthAlipay = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetAuthAlipayDomain(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.AuthAlipayDomain = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetAuthAlipayLoginId(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.AuthAlipayLoginId = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetAuthBeiAnCid(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.AuthBeiAnCid = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetAuthDomain(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.AuthDomain = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetCertifiedFrom(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.CertifiedFrom = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetCertifiedTime(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.CertifiedTime = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetCertifyStatus(v int32) *LoadRealNameInfoByPkResponseBodyData {
	s.CertifyStatus = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetCicCertifyFrom(v int32) *LoadRealNameInfoByPkResponseBodyData {
	s.CicCertifyFrom = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetCicCertifyProduct(v int64) *LoadRealNameInfoByPkResponseBodyData {
	s.CicCertifyProduct = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetIsBankIDAuth(v bool) *LoadRealNameInfoByPkResponseBodyData {
	s.IsBankIDAuth = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetIsCertified(v bool) *LoadRealNameInfoByPkResponseBodyData {
	s.IsCertified = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetLicenseNumber(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.LicenseNumber = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetLicenseType(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.LicenseType = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetName(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.Name = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetNewUnityRealNameAccount(v string) *LoadRealNameInfoByPkResponseBodyData {
	s.NewUnityRealNameAccount = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) SetProcessingEnterpriseCertify(v bool) *LoadRealNameInfoByPkResponseBodyData {
	s.ProcessingEnterpriseCertify = &v
	return s
}

func (s *LoadRealNameInfoByPkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
