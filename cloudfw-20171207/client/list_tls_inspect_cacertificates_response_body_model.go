// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTlsInspectCACertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificates(v []*ListTlsInspectCACertificatesResponseBodyCertificates) *ListTlsInspectCACertificatesResponseBody
	GetCertificates() []*ListTlsInspectCACertificatesResponseBodyCertificates
	SetRequestId(v string) *ListTlsInspectCACertificatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListTlsInspectCACertificatesResponseBody
	GetTotalCount() *int64
}

type ListTlsInspectCACertificatesResponseBody struct {
	// The list of certificates.
	Certificates []*ListTlsInspectCACertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-******837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTlsInspectCACertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTlsInspectCACertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTlsInspectCACertificatesResponseBody) GetCertificates() []*ListTlsInspectCACertificatesResponseBodyCertificates {
	return s.Certificates
}

func (s *ListTlsInspectCACertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTlsInspectCACertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTlsInspectCACertificatesResponseBody) SetCertificates(v []*ListTlsInspectCACertificatesResponseBodyCertificates) *ListTlsInspectCACertificatesResponseBody {
	s.Certificates = v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBody) SetRequestId(v string) *ListTlsInspectCACertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBody) SetTotalCount(v int64) *ListTlsInspectCACertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBody) Validate() error {
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTlsInspectCACertificatesResponseBodyCertificates struct {
	// The encryption algorithm of the CA certificate. Valid values:
	//
	// - **RSA**: the RSA algorithm.
	//
	// - **ECC**: the ECC algorithm.
	//
	// - **SM2**: the SM2 algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The alias of the certificate.
	//
	// example:
	//
	// rsa_ml_***_root
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The ID of the CA certificate.
	//
	// example:
	//
	// 2732BB48-2969-5716-B5D9-******CA85
	CaCertId *string `json:"CaCertId,omitempty" xml:"CaCertId,omitempty"`
	// The type of the CA certificate. Valid values:
	//
	// - **ROOT**: a root CA certificate.
	//
	// - **SUB_ROOT**: a subordinate CA certificate.
	//
	// example:
	//
	// ROOT
	CaCertType *string `json:"CaCertType,omitempty" xml:"CaCertType,omitempty"`
	// The expiration timestamp.
	//
	// example:
	//
	// 1934***149
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The key length of the CA certificate.
	//
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The ID of the parent CA certificate.
	//
	// example:
	//
	// 340BB48-2969-5716-B5D9-****ACA85
	ParentCaCertId *string `json:"ParentCaCertId,omitempty" xml:"ParentCaCertId,omitempty"`
	// The signature algorithm of the CA certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - **ISSUE**: enabled.
	//
	// - **REVOKE**: revoked.
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTlsInspectCACertificatesResponseBodyCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListTlsInspectCACertificatesResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetAliasName() *string {
	return s.AliasName
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetCaCertId() *string {
	return s.CaCertId
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetCaCertType() *string {
	return s.CaCertType
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetKeySize() *int32 {
	return s.KeySize
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetParentCaCertId() *string {
	return s.ParentCaCertId
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) GetStatus() *string {
	return s.Status
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetAlgorithm(v string) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.Algorithm = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetAliasName(v string) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.AliasName = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetCaCertId(v string) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.CaCertId = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetCaCertType(v string) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.CaCertType = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetExpirationTime(v int64) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.ExpirationTime = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetKeySize(v int32) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.KeySize = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetParentCaCertId(v string) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.ParentCaCertId = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetSignAlgorithm(v string) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.SignAlgorithm = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) SetStatus(v string) *ListTlsInspectCACertificatesResponseBodyCertificates {
	s.Status = &v
	return s
}

func (s *ListTlsInspectCACertificatesResponseBodyCertificates) Validate() error {
	return dara.Validate(s)
}
