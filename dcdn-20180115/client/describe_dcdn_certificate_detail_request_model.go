// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *DescribeDcdnCertificateDetailRequest
	GetCertName() *string
	SetOwnerId(v int64) *DescribeDcdnCertificateDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnCertificateDetailRequest
	GetSecurityToken() *string
}

type DescribeDcdnCertificateDetailRequest struct {
	// The name of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateDetailRequest) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnCertificateDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnCertificateDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnCertificateDetailRequest) SetCertName(v string) *DescribeDcdnCertificateDetailRequest {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnCertificateDetailRequest) SetOwnerId(v int64) *DescribeDcdnCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnCertificateDetailRequest) SetSecurityToken(v string) *DescribeDcdnCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
