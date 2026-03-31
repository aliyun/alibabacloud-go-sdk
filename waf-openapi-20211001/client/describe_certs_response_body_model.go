// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCerts(v []*DescribeCertsResponseBodyCerts) *DescribeCertsResponseBody
	GetCerts() []*DescribeCertsResponseBodyCerts
	SetRequestId(v string) *DescribeCertsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeCertsResponseBody
	GetTotalCount() *int64
}

type DescribeCertsResponseBody struct {
	// The certificates.
	Certs []*DescribeCertsResponseBodyCerts `json:"Certs,omitempty" xml:"Certs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 19511B0D-5AE0-5600-BB8A-DC2C8345****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertsResponseBody) GetCerts() []*DescribeCertsResponseBodyCerts {
	return s.Certs
}

func (s *DescribeCertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCertsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCertsResponseBody) SetCerts(v []*DescribeCertsResponseBodyCerts) *DescribeCertsResponseBody {
	s.Certs = v
	return s
}

func (s *DescribeCertsResponseBody) SetRequestId(v string) *DescribeCertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertsResponseBody) SetTotalCount(v int64) *DescribeCertsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCertsResponseBody) Validate() error {
	if s.Certs != nil {
		for _, item := range s.Certs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCertsResponseBodyCerts struct {
	// The expiration time.
	//
	// example:
	//
	// 1976256736582
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The effective time.
	//
	// example:
	//
	// 1976256836582
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The globally unique ID of the certificate. The value follows a "\\<Certificate ID>-ap-southeast-1" format. For example, if the ID of the certificate is 123, the value of the CertIdentifier parameter is 123-ap-southeast-1.
	//
	// example:
	//
	// 123-ap-southeast-1
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// waf1234
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The common name.
	//
	// example:
	//
	// *.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The domain that is supported by the certificate.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the certificate chain is complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsChainCompleted *bool `json:"IsChainCompleted,omitempty" xml:"IsChainCompleted,omitempty"`
}

func (s DescribeCertsResponseBodyCerts) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertsResponseBodyCerts) GoString() string {
	return s.String()
}

func (s *DescribeCertsResponseBodyCerts) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *DescribeCertsResponseBodyCerts) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *DescribeCertsResponseBodyCerts) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeCertsResponseBodyCerts) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCertsResponseBodyCerts) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeCertsResponseBodyCerts) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCertsResponseBodyCerts) GetIsChainCompleted() *bool {
	return s.IsChainCompleted
}

func (s *DescribeCertsResponseBodyCerts) SetAfterDate(v int64) *DescribeCertsResponseBodyCerts {
	s.AfterDate = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetBeforeDate(v int64) *DescribeCertsResponseBodyCerts {
	s.BeforeDate = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetCertIdentifier(v string) *DescribeCertsResponseBodyCerts {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetCertName(v string) *DescribeCertsResponseBodyCerts {
	s.CertName = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetCommonName(v string) *DescribeCertsResponseBodyCerts {
	s.CommonName = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetDomain(v string) *DescribeCertsResponseBodyCerts {
	s.Domain = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetIsChainCompleted(v bool) *DescribeCertsResponseBodyCerts {
	s.IsChainCompleted = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) Validate() error {
	return dara.Validate(s)
}
