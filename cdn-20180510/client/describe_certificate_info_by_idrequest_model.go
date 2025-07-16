// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateInfoByIDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *DescribeCertificateInfoByIDRequest
	GetCertId() *string
}

type DescribeCertificateInfoByIDRequest struct {
	// The ID of the certificate. You can query only one certificate in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1644xx
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s DescribeCertificateInfoByIDRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateInfoByIDRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDRequest) GetCertId() *string {
	return s.CertId
}

func (s *DescribeCertificateInfoByIDRequest) SetCertId(v string) *DescribeCertificateInfoByIDRequest {
	s.CertId = &v
	return s
}

func (s *DescribeCertificateInfoByIDRequest) Validate() error {
	return dara.Validate(s)
}
