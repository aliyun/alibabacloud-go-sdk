// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSMCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *DescribeDcdnSMCertificateDetailRequest
	GetCertIdentifier() *string
	SetOwnerId(v int64) *DescribeDcdnSMCertificateDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnSMCertificateDetailRequest
	GetSecurityToken() *string
}

type DescribeDcdnSMCertificateDetailRequest struct {
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 648****-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnSMCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSMCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateDetailRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeDcdnSMCertificateDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnSMCertificateDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnSMCertificateDetailRequest) SetCertIdentifier(v string) *DescribeDcdnSMCertificateDetailRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailRequest) SetOwnerId(v int64) *DescribeDcdnSMCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailRequest) SetSecurityToken(v string) *DescribeDcdnSMCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
