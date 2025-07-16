// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateDetailByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *DescribeCdnCertificateDetailByIdRequest
	GetCertId() *string
	SetCertRegion(v string) *DescribeCdnCertificateDetailByIdRequest
	GetCertRegion() *string
	SetOwnerId(v int64) *DescribeCdnCertificateDetailByIdRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnCertificateDetailByIdRequest
	GetSecurityToken() *string
}

type DescribeCdnCertificateDetailByIdRequest struct {
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The region of the certificate. Valid values:
	//
	// 	- **ap-southeast-1**: Singapore
	//
	// 	- **cn-hangzhou**: China (Hangzhou)
	//
	// Default value: **cn-hangzhou**
	//
	// example:
	//
	// cn-hangzhou
	CertRegion    *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnCertificateDetailByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailByIdRequest) GetCertId() *string {
	return s.CertId
}

func (s *DescribeCdnCertificateDetailByIdRequest) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeCdnCertificateDetailByIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnCertificateDetailByIdRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnCertificateDetailByIdRequest) SetCertId(v string) *DescribeCdnCertificateDetailByIdRequest {
	s.CertId = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdRequest) SetCertRegion(v string) *DescribeCdnCertificateDetailByIdRequest {
	s.CertRegion = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdRequest) SetOwnerId(v int64) *DescribeCdnCertificateDetailByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdRequest) SetSecurityToken(v string) *DescribeCdnCertificateDetailByIdRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdRequest) Validate() error {
	return dara.Validate(s)
}
