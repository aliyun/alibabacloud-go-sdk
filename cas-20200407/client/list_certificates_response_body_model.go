// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateList(v []*ListCertificatesResponseBodyCertificateList) *ListCertificatesResponseBody
	GetCertificateList() []*ListCertificatesResponseBodyCertificateList
	SetCurrentPage(v int32) *ListCertificatesResponseBody
	GetCurrentPage() *int32
	SetRequestId(v string) *ListCertificatesResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListCertificatesResponseBody
	GetShowSize() *int32
	SetTotalCount(v int64) *ListCertificatesResponseBody
	GetTotalCount() *int64
}

type ListCertificatesResponseBody struct {
	// The list of certificates.
	CertificateList []*ListCertificatesResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBody) GetCertificateList() []*ListCertificatesResponseBodyCertificateList {
	return s.CertificateList
}

func (s *ListCertificatesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertificatesResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCertificatesResponseBody) SetCertificateList(v []*ListCertificatesResponseBodyCertificateList) *ListCertificatesResponseBody {
	s.CertificateList = v
	return s
}

func (s *ListCertificatesResponseBody) SetCurrentPage(v int32) *ListCertificatesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCertificatesResponseBody) SetRequestId(v string) *ListCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertificatesResponseBody) SetShowSize(v int32) *ListCertificatesResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCertificatesResponseBody) SetTotalCount(v int64) *ListCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertificatesResponseBody) Validate() error {
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

type ListCertificatesResponseBodyCertificateList struct {
	// The encryption algorithm of the certificate. Valid values:
	//
	// - **RSA**
	//
	// - **ECC**
	//
	// - **SM2**
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The global ID of the certificate. This ID is used across Alibaba Cloud services and is in the `Certificate ID-Region ID` format. For example, if a certificate ID is `123`, the `CertIdentifier` is `123-cn-hangzhou` for the Alibaba Cloud China site and `123-ap-southeast-1` for the Alibaba Cloud International site (www\\.alibabacloud.com).
	//
	// - For the Alibaba Cloud China website, the format is certificate ID + "-cn-hangzhou".
	//
	// - For the Alibaba Cloud International website (www\\.alibabacloud.com), the format is certificate ID + "-ap-southeast-1".
	//
	// For example, if the certificate ID is 123, the CertIdentifier is "123-cn-hangzhou" for the China site and "123-ap-southeast-1" for the International site.
	//
	// example:
	//
	// 21589515-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 17281539
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// test
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// The source of the certificate.
	//
	// - BUY: A purchased certificate.
	//
	// - TEST: A test certificate.
	//
	// - UPLOAD: An uploaded certificate.
	//
	// example:
	//
	// BUY
	CertificateSource *string `json:"CertificateSource,omitempty" xml:"CertificateSource,omitempty"`
	// The status of the certificate.
	//
	// - **issued**: The certificate is issued.
	//
	// - **revoked**: The certificate is revoked.
	//
	// - **willExpire**: The certificate is about to expire.
	//
	// - **expired**: The certificate has expired.
	//
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// aliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The domain names that are bound to the certificate. Multiple domain names are separated by commas.
	//
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether a private key is available. Valid values:
	//
	// - **true**: A private key is available.
	//
	// - **false**: A private key is not available.
	//
	// example:
	//
	// true
	ExistPrivateKey *bool `json:"ExistPrivateKey,omitempty" xml:"ExistPrivateKey,omitempty"`
	// The fingerprint of the public key.
	//
	// example:
	//
	// 123
	FingerPrint *string `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
	// The ID of the certificate instance.
	//
	// example:
	//
	// cas-cn-v***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The certification authority.
	//
	// example:
	//
	// DigiCert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The key size, in bits.
	//
	// - For RSA keys, typical sizes are 2048, 3072, or 4096.
	//
	// - For ECC or SM2 keys, the typical size is 256.
	//
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The end of the certificate validity period.
	//
	// example:
	//
	// 1749580567000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the certificate validity period.
	//
	// example:
	//
	// 1760745600000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 123
	Serial *string `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// An array that contains the alternative domain names of the certificate. This parameter corresponds to the `Subject Alternative Name` field of the certificate.
	SubjectAlternativeNames []*string `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
	// The list of Alibaba Cloud products in which the certificate is deployed.
	UsingProductList []*string `json:"UsingProductList,omitempty" xml:"UsingProductList,omitempty" type:"Repeated"`
}

func (s ListCertificatesResponseBodyCertificateList) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBodyCertificateList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertificateName() *string {
	return s.CertificateName
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertificateSource() *string {
	return s.CertificateSource
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *ListCertificatesResponseBodyCertificateList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCertificatesResponseBodyCertificateList) GetDomain() *string {
	return s.Domain
}

func (s *ListCertificatesResponseBodyCertificateList) GetExistPrivateKey() *bool {
	return s.ExistPrivateKey
}

func (s *ListCertificatesResponseBodyCertificateList) GetFingerPrint() *string {
	return s.FingerPrint
}

func (s *ListCertificatesResponseBodyCertificateList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCertificatesResponseBodyCertificateList) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCertificatesResponseBodyCertificateList) GetKeySize() *int32 {
	return s.KeySize
}

func (s *ListCertificatesResponseBodyCertificateList) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *ListCertificatesResponseBodyCertificateList) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *ListCertificatesResponseBodyCertificateList) GetSerial() *string {
	return s.Serial
}

func (s *ListCertificatesResponseBodyCertificateList) GetSubjectAlternativeNames() []*string {
	return s.SubjectAlternativeNames
}

func (s *ListCertificatesResponseBodyCertificateList) GetUsingProductList() []*string {
	return s.UsingProductList
}

func (s *ListCertificatesResponseBodyCertificateList) SetAlgorithm(v string) *ListCertificatesResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertIdentifier(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertIdentifier = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertificateId(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertificateId = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertificateName(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertificateName = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertificateSource(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertificateSource = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertificateStatus(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertificateStatus = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCommonName(v string) *ListCertificatesResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetDomain(v string) *ListCertificatesResponseBodyCertificateList {
	s.Domain = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetExistPrivateKey(v bool) *ListCertificatesResponseBodyCertificateList {
	s.ExistPrivateKey = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetFingerPrint(v string) *ListCertificatesResponseBodyCertificateList {
	s.FingerPrint = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetInstanceId(v string) *ListCertificatesResponseBodyCertificateList {
	s.InstanceId = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetIssuer(v string) *ListCertificatesResponseBodyCertificateList {
	s.Issuer = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetKeySize(v int32) *ListCertificatesResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetNotAfter(v int64) *ListCertificatesResponseBodyCertificateList {
	s.NotAfter = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetNotBefore(v int64) *ListCertificatesResponseBodyCertificateList {
	s.NotBefore = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetSerial(v string) *ListCertificatesResponseBodyCertificateList {
	s.Serial = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetSubjectAlternativeNames(v []*string) *ListCertificatesResponseBodyCertificateList {
	s.SubjectAlternativeNames = v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetUsingProductList(v []*string) *ListCertificatesResponseBodyCertificateList {
	s.UsingProductList = v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) Validate() error {
	return dara.Validate(s)
}
