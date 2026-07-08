// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertDetail(v *DescribeCertDetailResponseBodyCertDetail) *DescribeCertDetailResponseBody
	GetCertDetail() *DescribeCertDetailResponseBodyCertDetail
	SetRequestId(v string) *DescribeCertDetailResponseBody
	GetRequestId() *string
}

type DescribeCertDetailResponseBody struct {
	// The SSL certificate details.
	CertDetail *DescribeCertDetailResponseBodyCertDetail `json:"CertDetail,omitempty" xml:"CertDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3C115DBE-8E53-5A12-9CAF-FD3F****CDF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCertDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertDetailResponseBody) GetCertDetail() *DescribeCertDetailResponseBodyCertDetail {
	return s.CertDetail
}

func (s *DescribeCertDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCertDetailResponseBody) SetCertDetail(v *DescribeCertDetailResponseBodyCertDetail) *DescribeCertDetailResponseBody {
	s.CertDetail = v
	return s
}

func (s *DescribeCertDetailResponseBody) SetRequestId(v string) *DescribeCertDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertDetailResponseBody) Validate() error {
	if s.CertDetail != nil {
		if err := s.CertDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCertDetailResponseBodyCertDetail struct {
	// The expiration date of the CA certificate. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1976256736582
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The issuance time of the certificate. Unit: milliseconds.
	//
	// example:
	//
	// 1976256736582
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// testCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// *.xxxaliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The domain name associated with the certificate.
	//
	// example:
	//
	// demo.xxxaliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The list of other domain names associated with the certificate.
	Sans []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
}

func (s DescribeCertDetailResponseBodyCertDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertDetailResponseBodyCertDetail) GoString() string {
	return s.String()
}

func (s *DescribeCertDetailResponseBodyCertDetail) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *DescribeCertDetailResponseBodyCertDetail) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *DescribeCertDetailResponseBodyCertDetail) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeCertDetailResponseBodyCertDetail) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCertDetailResponseBodyCertDetail) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeCertDetailResponseBodyCertDetail) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCertDetailResponseBodyCertDetail) GetSans() []*string {
	return s.Sans
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetAfterDate(v int64) *DescribeCertDetailResponseBodyCertDetail {
	s.AfterDate = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetBeforeDate(v int64) *DescribeCertDetailResponseBodyCertDetail {
	s.BeforeDate = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetCertIdentifier(v string) *DescribeCertDetailResponseBodyCertDetail {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetCertName(v string) *DescribeCertDetailResponseBodyCertDetail {
	s.CertName = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetCommonName(v string) *DescribeCertDetailResponseBodyCertDetail {
	s.CommonName = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetDomain(v string) *DescribeCertDetailResponseBodyCertDetail {
	s.Domain = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetSans(v []*string) *DescribeCertDetailResponseBodyCertDetail {
	s.Sans = v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) Validate() error {
	return dara.Validate(s)
}
