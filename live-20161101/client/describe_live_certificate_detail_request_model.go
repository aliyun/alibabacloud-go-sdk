// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *DescribeLiveCertificateDetailRequest
	GetCertName() *string
	SetOwnerId(v int64) *DescribeLiveCertificateDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveCertificateDetailRequest
	GetSecurityToken() *string
}

type DescribeLiveCertificateDetailRequest struct {
	// The name of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cert-****
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLiveCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateDetailRequest) GetCertName() *string {
	return s.CertName
}

func (s *DescribeLiveCertificateDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveCertificateDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveCertificateDetailRequest) SetCertName(v string) *DescribeLiveCertificateDetailRequest {
	s.CertName = &v
	return s
}

func (s *DescribeLiveCertificateDetailRequest) SetOwnerId(v int64) *DescribeLiveCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveCertificateDetailRequest) SetSecurityToken(v string) *DescribeLiveCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
