// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainByCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExact(v bool) *DescribeDcdnDomainByCertificateRequest
	GetExact() *bool
	SetSSLPub(v string) *DescribeDcdnDomainByCertificateRequest
	GetSSLPub() *string
	SetSSLStatus(v bool) *DescribeDcdnDomainByCertificateRequest
	GetSSLStatus() *bool
}

type DescribeDcdnDomainByCertificateRequest struct {
	// Specifies whether the domain name list to return matches the SSL certificate.
	//
	// - **true**: The domain name list matches the SSL certificate.
	//
	// - **false**: The domain name list does not match the SSL certificate.
	//
	// example:
	//
	// true
	Exact *bool `json:"Exact,omitempty" xml:"Exact,omitempty"`
	// The public key of the certificate.
	//
	// You must use Base64 encoding schemes and then the encodeURIComponent method to encode the public key. PEM files are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// Specifies whether the domain name list to return contains only domain names with HTTPS enabled or disabled.
	//
	// 	- true: The list contains only domain names with HTTPS enabled.
	//
	// 	- false: The list contains only domain names with HTTPS disabled.
	//
	// example:
	//
	// true
	SSLStatus *bool `json:"SSLStatus,omitempty" xml:"SSLStatus,omitempty"`
}

func (s DescribeDcdnDomainByCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateRequest) GetExact() *bool {
	return s.Exact
}

func (s *DescribeDcdnDomainByCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeDcdnDomainByCertificateRequest) GetSSLStatus() *bool {
	return s.SSLStatus
}

func (s *DescribeDcdnDomainByCertificateRequest) SetExact(v bool) *DescribeDcdnDomainByCertificateRequest {
	s.Exact = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateRequest) SetSSLPub(v string) *DescribeDcdnDomainByCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateRequest) SetSSLStatus(v bool) *DescribeDcdnDomainByCertificateRequest {
	s.SSLStatus = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateRequest) Validate() error {
	return dara.Validate(s)
}
