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
}

type DescribeCertsResponseBody struct {
	// The certificate information about the website.
	Certs []*DescribeCertsResponseBodyCerts `json:"Certs,omitempty" xml:"Certs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeCertsResponseBody) SetCerts(v []*DescribeCertsResponseBodyCerts) *DescribeCertsResponseBody {
	s.Certs = v
	return s
}

func (s *DescribeCertsResponseBody) SetRequestId(v string) *DescribeCertsResponseBody {
	s.RequestId = &v
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
	// The global certificate ID, which is in the certificate ID-cn-hangzhou format. If the ID of the certificate is 123, CertIdentifier is 123-cn-hangzhou.
	//
	// example:
	//
	// 126345-ap-southeast-1
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The domain name that is associated with the certificate.
	//
	// example:
	//
	// www.aliyun.com
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// Indicates whether the certificate is associated with the domain name. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DomainRelated *bool `json:"DomainRelated,omitempty" xml:"DomainRelated,omitempty"`
	// The expiration date of the certificate. The value is a string.
	//
	// example:
	//
	// 2021-09-12
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 81
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// Symantec
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// testcert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The issuance date of the certificate. The value is a string.
	//
	// example:
	//
	// 2019-09-12
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeCertsResponseBodyCerts) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertsResponseBodyCerts) GoString() string {
	return s.String()
}

func (s *DescribeCertsResponseBodyCerts) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeCertsResponseBodyCerts) GetCommon() *string {
	return s.Common
}

func (s *DescribeCertsResponseBodyCerts) GetDomainRelated() *bool {
	return s.DomainRelated
}

func (s *DescribeCertsResponseBodyCerts) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeCertsResponseBodyCerts) GetId() *int32 {
	return s.Id
}

func (s *DescribeCertsResponseBodyCerts) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeCertsResponseBodyCerts) GetName() *string {
	return s.Name
}

func (s *DescribeCertsResponseBodyCerts) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeCertsResponseBodyCerts) SetCertIdentifier(v string) *DescribeCertsResponseBodyCerts {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetCommon(v string) *DescribeCertsResponseBodyCerts {
	s.Common = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetDomainRelated(v bool) *DescribeCertsResponseBodyCerts {
	s.DomainRelated = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetEndDate(v string) *DescribeCertsResponseBodyCerts {
	s.EndDate = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetId(v int32) *DescribeCertsResponseBodyCerts {
	s.Id = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetIssuer(v string) *DescribeCertsResponseBodyCerts {
	s.Issuer = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetName(v string) *DescribeCertsResponseBodyCerts {
	s.Name = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetStartDate(v string) *DescribeCertsResponseBodyCerts {
	s.StartDate = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) Validate() error {
	return dara.Validate(s)
}
