// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *DescribeCdnCertificateDetailRequest
	GetCertName() *string
	SetOwnerId(v int64) *DescribeCdnCertificateDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnCertificateDetailRequest
	GetSecurityToken() *string
}

type DescribeCdnCertificateDetailRequest struct {
	// The ID of the SSL certificate. You can query only one certificate at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// cert-15480655xxxx
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailRequest) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCdnCertificateDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnCertificateDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnCertificateDetailRequest) SetCertName(v string) *DescribeCdnCertificateDetailRequest {
	s.CertName = &v
	return s
}

func (s *DescribeCdnCertificateDetailRequest) SetOwnerId(v int64) *DescribeCdnCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnCertificateDetailRequest) SetSecurityToken(v string) *DescribeCdnCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
