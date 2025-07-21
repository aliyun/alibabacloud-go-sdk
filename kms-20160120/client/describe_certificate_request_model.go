// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *DescribeCertificateRequest
	GetCertificateId() *string
}

type DescribeCertificateRequest struct {
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s DescribeCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificateRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DescribeCertificateRequest) SetCertificateId(v string) *DescribeCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *DescribeCertificateRequest) Validate() error {
	return dara.Validate(s)
}
