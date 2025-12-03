// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainExtensions(v *DescribeDomainExtensionsResponseBodyDomainExtensions) *DescribeDomainExtensionsResponseBody
	GetDomainExtensions() *DescribeDomainExtensionsResponseBodyDomainExtensions
	SetRequestId(v string) *DescribeDomainExtensionsResponseBody
	GetRequestId() *string
}

type DescribeDomainExtensionsResponseBody struct {
	// A list of additional certificates.
	DomainExtensions *DescribeDomainExtensionsResponseBodyDomainExtensions `json:"DomainExtensions,omitempty" xml:"DomainExtensions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 48C1B671-C6DB-4DDE-9B30-10557E36CDE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsResponseBody) GetDomainExtensions() *DescribeDomainExtensionsResponseBodyDomainExtensions {
	return s.DomainExtensions
}

func (s *DescribeDomainExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainExtensionsResponseBody) SetDomainExtensions(v *DescribeDomainExtensionsResponseBodyDomainExtensions) *DescribeDomainExtensionsResponseBody {
	s.DomainExtensions = v
	return s
}

func (s *DescribeDomainExtensionsResponseBody) SetRequestId(v string) *DescribeDomainExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainExtensionsResponseBody) Validate() error {
	if s.DomainExtensions != nil {
		if err := s.DomainExtensions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainExtensionsResponseBodyDomainExtensions struct {
	DomainExtension []*DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension `json:"DomainExtension,omitempty" xml:"DomainExtension,omitempty" type:"Repeated"`
}

func (s DescribeDomainExtensionsResponseBodyDomainExtensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainExtensionsResponseBodyDomainExtensions) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensions) GetDomainExtension() []*DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension {
	return s.DomainExtension
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensions) SetDomainExtension(v []*DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) *DescribeDomainExtensionsResponseBodyDomainExtensions {
	s.DomainExtension = v
	return s
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensions) Validate() error {
	if s.DomainExtension != nil {
		for _, item := range s.DomainExtension {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the additional certificate.
	//
	// example:
	//
	// de-bp1rp7ta1****
	DomainExtensionId *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	// The ID of the certificate used by the domain name.
	//
	// example:
	//
	// 1231579085529123_166f8204689_1714763408_70998****
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) GetDomainExtensionId() *string {
	return s.DomainExtensionId
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) SetDomain(v string) *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension {
	s.Domain = &v
	return s
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) SetDomainExtensionId(v string) *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) SetServerCertificateId(v string) *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeDomainExtensionsResponseBodyDomainExtensionsDomainExtension) Validate() error {
	return dara.Validate(s)
}
