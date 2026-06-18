// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificates(v []*ListCasCertificatesResponseBodyCertificates) *ListCasCertificatesResponseBody
	GetCertificates() []*ListCasCertificatesResponseBodyCertificates
	SetPageNumber(v int32) *ListCasCertificatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCasCertificatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCasCertificatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCasCertificatesResponseBody
	GetTotalCount() *int32
}

type ListCasCertificatesResponseBody struct {
	Certificates []*ListCasCertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCasCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCasCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCasCertificatesResponseBody) GetCertificates() []*ListCasCertificatesResponseBodyCertificates {
	return s.Certificates
}

func (s *ListCasCertificatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCasCertificatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCasCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCasCertificatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCasCertificatesResponseBody) SetCertificates(v []*ListCasCertificatesResponseBodyCertificates) *ListCasCertificatesResponseBody {
	s.Certificates = v
	return s
}

func (s *ListCasCertificatesResponseBody) SetPageNumber(v int32) *ListCasCertificatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCasCertificatesResponseBody) SetPageSize(v int32) *ListCasCertificatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCasCertificatesResponseBody) SetRequestId(v string) *ListCasCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCasCertificatesResponseBody) SetTotalCount(v int32) *ListCasCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCasCertificatesResponseBody) Validate() error {
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

type ListCasCertificatesResponseBodyCertificates struct {
	// example:
	//
	// example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// 0151xxxx
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// example:
	//
	// 30000145
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// DigiCert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1708423200000
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// example:
	//
	// ap-southeast-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListCasCertificatesResponseBodyCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListCasCertificatesResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *ListCasCertificatesResponseBodyCertificates) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCasCertificatesResponseBodyCertificates) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *ListCasCertificatesResponseBodyCertificates) GetId() *int64 {
	return s.Id
}

func (s *ListCasCertificatesResponseBodyCertificates) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCasCertificatesResponseBodyCertificates) GetName() *string {
	return s.Name
}

func (s *ListCasCertificatesResponseBodyCertificates) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListCasCertificatesResponseBodyCertificates) GetRegion() *string {
	return s.Region
}

func (s *ListCasCertificatesResponseBodyCertificates) SetCommonName(v string) *ListCasCertificatesResponseBodyCertificates {
	s.CommonName = &v
	return s
}

func (s *ListCasCertificatesResponseBodyCertificates) SetFingerprint(v string) *ListCasCertificatesResponseBodyCertificates {
	s.Fingerprint = &v
	return s
}

func (s *ListCasCertificatesResponseBodyCertificates) SetId(v int64) *ListCasCertificatesResponseBodyCertificates {
	s.Id = &v
	return s
}

func (s *ListCasCertificatesResponseBodyCertificates) SetIssuer(v string) *ListCasCertificatesResponseBodyCertificates {
	s.Issuer = &v
	return s
}

func (s *ListCasCertificatesResponseBodyCertificates) SetName(v string) *ListCasCertificatesResponseBodyCertificates {
	s.Name = &v
	return s
}

func (s *ListCasCertificatesResponseBodyCertificates) SetNotAfter(v string) *ListCasCertificatesResponseBodyCertificates {
	s.NotAfter = &v
	return s
}

func (s *ListCasCertificatesResponseBodyCertificates) SetRegion(v string) *ListCasCertificatesResponseBodyCertificates {
	s.Region = &v
	return s
}

func (s *ListCasCertificatesResponseBodyCertificates) Validate() error {
	return dara.Validate(s)
}
