// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertList(v []*ListCertResponseBodyCertList) *ListCertResponseBody
	GetCertList() []*ListCertResponseBodyCertList
	SetCurrentPage(v int64) *ListCertResponseBody
	GetCurrentPage() *int64
	SetRequestId(v string) *ListCertResponseBody
	GetRequestId() *string
	SetShowSize(v int64) *ListCertResponseBody
	GetShowSize() *int64
	SetTotalCount(v int64) *ListCertResponseBody
	GetTotalCount() *int64
}

type ListCertResponseBody struct {
	// The list of certificates.
	CertList []*ListCertResponseBodyCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The page size. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertResponseBody) GetCertList() []*ListCertResponseBodyCertList {
	return s.CertList
}

func (s *ListCertResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertResponseBody) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListCertResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCertResponseBody) SetCertList(v []*ListCertResponseBodyCertList) *ListCertResponseBody {
	s.CertList = v
	return s
}

func (s *ListCertResponseBody) SetCurrentPage(v int64) *ListCertResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCertResponseBody) SetRequestId(v string) *ListCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertResponseBody) SetShowSize(v int64) *ListCertResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCertResponseBody) SetTotalCount(v int64) *ListCertResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertResponseBody) Validate() error {
	if s.CertList != nil {
		for _, item := range s.CertList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCertResponseBodyCertList struct {
	// The expiration date of the certificate. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1634283958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The encryption algorithm of the certificate. Valid values:
	//
	// - **RSA**: the RSA algorithm
	//
	// - **ECC**: the ECC algorithm
	//
	// - **SM2**: the SM2 algorithm
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The start date of the certificate\\"s validity period. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1665819958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate. Valid values:
	//
	// - **CA**: a Certificate Authority (CA) certificate
	//
	// - **CERT**: an issued certificate
	//
	// example:
	//
	// CERT
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The common name of the certificate. This is typically the primary domain name associated with the certificate.
	//
	// example:
	//
	// aliyun.alibaba.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// Indicates whether a private key is available for the certificate. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	ExistPrivateKey *bool `json:"ExistPrivateKey,omitempty" xml:"ExistPrivateKey,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 14dcc8afc7578e
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The issuer of the certificate.
	//
	// example:
	//
	// mySSL
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The Subject Alternative Names (SANs) associated with the certificate. Multiple domain names are separated by commas (,).
	//
	// example:
	//
	// *.alibaba.com,aliyun.alibaba.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate. This parameter is returned only if the `OrderType` request parameter is set to `CERT` or `UPLOAD`.
	//
	// example:
	//
	// 038abf4c27c33a7c11ad6658124135b52180
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The source of the certificate. Valid values:
	//
	// - **upload**: The certificate is uploaded.
	//
	// - **aliyun**: The certificate is from Alibaba Cloud.
	//
	// example:
	//
	// aliyun
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - **ISSUE**: The certificate is issued.
	//
	// - **REVOKE**: The certificate is revoked.
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The warehouse ID.
	//
	// example:
	//
	// 2
	WhId *int64 `json:"WhId,omitempty" xml:"WhId,omitempty"`
	// The warehouse instance ID.
	//
	// example:
	//
	// test_whInstanceId
	WhInstanceId *string `json:"WhInstanceId,omitempty" xml:"WhInstanceId,omitempty"`
}

func (s ListCertResponseBodyCertList) String() string {
	return dara.Prettify(s)
}

func (s ListCertResponseBodyCertList) GoString() string {
	return s.String()
}

func (s *ListCertResponseBodyCertList) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *ListCertResponseBodyCertList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCertResponseBodyCertList) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *ListCertResponseBodyCertList) GetCertType() *string {
	return s.CertType
}

func (s *ListCertResponseBodyCertList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCertResponseBodyCertList) GetExistPrivateKey() *bool {
	return s.ExistPrivateKey
}

func (s *ListCertResponseBodyCertList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListCertResponseBodyCertList) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCertResponseBodyCertList) GetSans() *string {
	return s.Sans
}

func (s *ListCertResponseBodyCertList) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListCertResponseBodyCertList) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *ListCertResponseBodyCertList) GetSourceType() *string {
	return s.SourceType
}

func (s *ListCertResponseBodyCertList) GetStatus() *string {
	return s.Status
}

func (s *ListCertResponseBodyCertList) GetWhId() *int64 {
	return s.WhId
}

func (s *ListCertResponseBodyCertList) GetWhInstanceId() *string {
	return s.WhInstanceId
}

func (s *ListCertResponseBodyCertList) SetAfterDate(v int64) *ListCertResponseBodyCertList {
	s.AfterDate = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetAlgorithm(v string) *ListCertResponseBodyCertList {
	s.Algorithm = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetBeforeDate(v int64) *ListCertResponseBodyCertList {
	s.BeforeDate = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetCertType(v string) *ListCertResponseBodyCertList {
	s.CertType = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetCommonName(v string) *ListCertResponseBodyCertList {
	s.CommonName = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetExistPrivateKey(v bool) *ListCertResponseBodyCertList {
	s.ExistPrivateKey = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetIdentifier(v string) *ListCertResponseBodyCertList {
	s.Identifier = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetIssuer(v string) *ListCertResponseBodyCertList {
	s.Issuer = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetSans(v string) *ListCertResponseBodyCertList {
	s.Sans = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetSerialNo(v string) *ListCertResponseBodyCertList {
	s.SerialNo = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetSignAlgorithm(v string) *ListCertResponseBodyCertList {
	s.SignAlgorithm = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetSourceType(v string) *ListCertResponseBodyCertList {
	s.SourceType = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetStatus(v string) *ListCertResponseBodyCertList {
	s.Status = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetWhId(v int64) *ListCertResponseBodyCertList {
	s.WhId = &v
	return s
}

func (s *ListCertResponseBodyCertList) SetWhInstanceId(v string) *ListCertResponseBodyCertList {
	s.WhInstanceId = &v
	return s
}

func (s *ListCertResponseBodyCertList) Validate() error {
	return dara.Validate(s)
}
