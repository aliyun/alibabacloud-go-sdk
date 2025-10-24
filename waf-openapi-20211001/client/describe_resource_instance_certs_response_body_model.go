// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceInstanceCertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCerts(v []*DescribeResourceInstanceCertsResponseBodyCerts) *DescribeResourceInstanceCertsResponseBody
	GetCerts() []*DescribeResourceInstanceCertsResponseBodyCerts
	SetRequestId(v string) *DescribeResourceInstanceCertsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeResourceInstanceCertsResponseBody
	GetTotalCount() *int64
}

type DescribeResourceInstanceCertsResponseBody struct {
	// The certificates.
	Certs []*DescribeResourceInstanceCertsResponseBodyCerts `json:"Certs,omitempty" xml:"Certs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D7861F61-***-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourceInstanceCertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceInstanceCertsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceInstanceCertsResponseBody) GetCerts() []*DescribeResourceInstanceCertsResponseBodyCerts {
	return s.Certs
}

func (s *DescribeResourceInstanceCertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceInstanceCertsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeResourceInstanceCertsResponseBody) SetCerts(v []*DescribeResourceInstanceCertsResponseBodyCerts) *DescribeResourceInstanceCertsResponseBody {
	s.Certs = v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBody) SetRequestId(v string) *DescribeResourceInstanceCertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBody) SetTotalCount(v int64) *DescribeResourceInstanceCertsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBody) Validate() error {
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

type DescribeResourceInstanceCertsResponseBodyCerts struct {
	// The time when the certificate expires.
	//
	// example:
	//
	// 1708415521211
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The time when the certificate was issued.
	//
	// example:
	//
	// 1708415521211
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The globally unique ID of the certificate. The value is in the "Certificate ID-cn-hangzhou" format. For example, if the ID of the certificate is 123, the value of CertIdentifier is 123-cn-hangzhou.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// demoCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The common name.
	//
	// example:
	//
	// *.aliyundemo.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The domain name for which the certificate is issued.
	//
	// example:
	//
	// waf.aliyundemo.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the certificate chain is complete.
	//
	// example:
	//
	// true
	IsChainCompleted *bool `json:"IsChainCompleted,omitempty" xml:"IsChainCompleted,omitempty"`
}

func (s DescribeResourceInstanceCertsResponseBodyCerts) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceInstanceCertsResponseBodyCerts) GoString() string {
	return s.String()
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) GetCertName() *string {
	return s.CertName
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) GetDomain() *string {
	return s.Domain
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) GetIsChainCompleted() *bool {
	return s.IsChainCompleted
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetAfterDate(v int64) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.AfterDate = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetBeforeDate(v int64) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.BeforeDate = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetCertIdentifier(v string) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetCertName(v string) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.CertName = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetCommonName(v string) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.CommonName = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetDomain(v string) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.Domain = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetIsChainCompleted(v bool) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.IsChainCompleted = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) Validate() error {
	return dara.Validate(s)
}
