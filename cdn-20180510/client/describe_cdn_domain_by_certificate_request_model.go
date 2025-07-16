// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainByCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExact(v bool) *DescribeCdnDomainByCertificateRequest
	GetExact() *bool
	SetSSLPub(v string) *DescribeCdnDomainByCertificateRequest
	GetSSLPub() *string
	SetSSLStatus(v bool) *DescribeCdnDomainByCertificateRequest
	GetSSLStatus() *bool
}

type DescribeCdnDomainByCertificateRequest struct {
	// Specifies whether the domain name list to return match the SSL certificate.
	//
	// 	- true: The domain name list match the SSL certificate.
	//
	// 	- false: The domain name list do not match the SSL certificate.
	//
	// example:
	//
	// true
	Exact *bool `json:"Exact,omitempty" xml:"Exact,omitempty"`
	// The public key of the SSL certificate. You must encode the public key in Base64 and then call the encodeURIComponent function to encode the public key again.
	//
	// The public key must be in the PEM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// Specifies whether the domain name list to return contains only domain names with HTTPS enabled or disabled.
	//
	// 	- true: The domain name list contains only domain names with HTTPS enabled.
	//
	// 	- false: The domain name list contains only domain names with HTTPS disabled.
	//
	// example:
	//
	// true
	SSLStatus *bool `json:"SSLStatus,omitempty" xml:"SSLStatus,omitempty"`
}

func (s DescribeCdnDomainByCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainByCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateRequest) GetExact() *bool {
	return s.Exact
}

func (s *DescribeCdnDomainByCertificateRequest) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeCdnDomainByCertificateRequest) GetSSLStatus() *bool {
	return s.SSLStatus
}

func (s *DescribeCdnDomainByCertificateRequest) SetExact(v bool) *DescribeCdnDomainByCertificateRequest {
	s.Exact = &v
	return s
}

func (s *DescribeCdnDomainByCertificateRequest) SetSSLPub(v string) *DescribeCdnDomainByCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *DescribeCdnDomainByCertificateRequest) SetSSLStatus(v bool) *DescribeCdnDomainByCertificateRequest {
	s.SSLStatus = &v
	return s
}

func (s *DescribeCdnDomainByCertificateRequest) Validate() error {
	return dara.Validate(s)
}
