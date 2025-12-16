// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePcaAndExternalCACertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateList(v []*DescribePcaAndExternalCACertificateListResponseBodyCertificateList) *DescribePcaAndExternalCACertificateListResponseBody
	GetCertificateList() []*DescribePcaAndExternalCACertificateListResponseBodyCertificateList
	SetCurrentPage(v int32) *DescribePcaAndExternalCACertificateListResponseBody
	GetCurrentPage() *int32
	SetPageCount(v int32) *DescribePcaAndExternalCACertificateListResponseBody
	GetPageCount() *int32
	SetRequestId(v string) *DescribePcaAndExternalCACertificateListResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *DescribePcaAndExternalCACertificateListResponseBody
	GetShowSize() *int32
	SetTotalCount(v int32) *DescribePcaAndExternalCACertificateListResponseBody
	GetTotalCount() *int32
}

type DescribePcaAndExternalCACertificateListResponseBody struct {
	CertificateList []*DescribePcaAndExternalCACertificateListResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePcaAndExternalCACertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePcaAndExternalCACertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) GetCertificateList() []*DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	return s.CertificateList
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) GetPageCount() *int32 {
	return s.PageCount
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) SetCertificateList(v []*DescribePcaAndExternalCACertificateListResponseBodyCertificateList) *DescribePcaAndExternalCACertificateListResponseBody {
	s.CertificateList = v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) SetCurrentPage(v int32) *DescribePcaAndExternalCACertificateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) SetPageCount(v int32) *DescribePcaAndExternalCACertificateListResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) SetRequestId(v string) *DescribePcaAndExternalCACertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) SetShowSize(v int32) *DescribePcaAndExternalCACertificateListResponseBody {
	s.ShowSize = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) SetTotalCount(v int32) *DescribePcaAndExternalCACertificateListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBody) Validate() error {
	if s.CertificateList != nil {
		for _, item := range s.CertificateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePcaAndExternalCACertificateListResponseBodyCertificateList struct {
	// example:
	//
	// 2022-08-23T16:15Z
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// 2021-01-01T00:00Z
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// example:
	//
	// SUB_ROOT
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// example:
	//
	// aliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// example:
	//
	// 05e148d8d3ecc9976d9ecd2b2f25****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// example:
	//
	// 05e148d8d3ecc9976d9ecd2b2f25****
	Md5          *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// example:
	//
	// 1a83bcbb89e562885e40aa0108f5****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// example:
	//
	// [ {"Type": 7, "Value": "192.0.XX.XX"}, {"Type": 2, "Value": "www.aliyundoc.com"}, ]
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// example:
	//
	// 62b2b943a32d96883a6650e672ea0276****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// example:
	//
	// 14dcc8afc7578e1fcec36d658f7e20de18f6957bbac42b373a66bc9de4e9****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// ISSUE
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE----- …… -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	// example:
	//
	// 3
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s DescribePcaAndExternalCACertificateListResponseBodyCertificateList) String() string {
	return dara.Prettify(s)
}

func (s DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetCertificateType() *string {
	return s.CertificateType
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetKeySize() *int32 {
	return s.KeySize
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetLocality() *string {
	return s.Locality
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetMd5() *string {
	return s.Md5
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetOrganization() *string {
	return s.Organization
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetSans() *string {
	return s.Sans
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetSha2() *string {
	return s.Sha2
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetState() *string {
	return s.State
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetStatus() *string {
	return s.Status
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetSubjectDN() *string {
	return s.SubjectDN
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) GetYears() *int32 {
	return s.Years
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetAfterDate(v int64) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetAlgorithm(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetBeforeDate(v int64) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetCertificateType(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.CertificateType = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetCommonName(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetCountryCode(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetIdentifier(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetKeySize(v int32) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetLocality(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetMd5(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetOrganization(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetOrganizationUnit(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetParentIdentifier(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.ParentIdentifier = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetSans(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetSerialNumber(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetSha2(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetSignAlgorithm(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetState(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetStatus(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetSubjectDN(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetX509Certificate(v string) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.X509Certificate = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) SetYears(v int32) *DescribePcaAndExternalCACertificateListResponseBodyCertificateList {
	s.Years = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponseBodyCertificateList) Validate() error {
	return dara.Validate(s)
}
