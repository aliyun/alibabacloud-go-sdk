// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSMCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *DescribeCdnSMCertificateDetailRequest
	GetCertIdentifier() *string
	SetOwnerId(v int64) *DescribeCdnSMCertificateDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnSMCertificateDetailRequest
	GetSecurityToken() *string
}

type DescribeCdnSMCertificateDetailRequest struct {
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

func (s DescribeCdnSMCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateDetailRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeCdnSMCertificateDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnSMCertificateDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnSMCertificateDetailRequest) SetCertIdentifier(v string) *DescribeCdnSMCertificateDetailRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailRequest) SetOwnerId(v int64) *DescribeCdnSMCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailRequest) SetSecurityToken(v string) *DescribeCdnSMCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
