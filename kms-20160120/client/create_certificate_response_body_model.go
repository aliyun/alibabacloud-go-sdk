// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *CreateCertificateResponseBody
	GetArn() *string
	SetCertificateId(v string) *CreateCertificateResponseBody
	GetCertificateId() *string
	SetCsr(v string) *CreateCertificateResponseBody
	GetCsr() *string
	SetRequestId(v string) *CreateCertificateResponseBody
	GetRequestId() *string
}

type CreateCertificateResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the certificate.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:certificate/98e85c94-52d0-40c9-b3b2-afda52f4****
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the certificate. It is the globally unique identifier (GUID) of the certificate in Certificates Manager.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The CSR in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----\\nMIIDADCCAegCAQAwgboxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhaaGVqaWFuZzER\\n****\\nmkj4rg==\\n-----END CERTIFICATE REQUEST-----\\n
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15a735a1-8fe6-45cc-a64c-3c4ff839334e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateResponseBody) GetArn() *string {
	return s.Arn
}

func (s *CreateCertificateResponseBody) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CreateCertificateResponseBody) GetCsr() *string {
	return s.Csr
}

func (s *CreateCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCertificateResponseBody) SetArn(v string) *CreateCertificateResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateCertificateResponseBody) SetCertificateId(v string) *CreateCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CreateCertificateResponseBody) SetCsr(v string) *CreateCertificateResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateCertificateResponseBody) SetRequestId(v string) *CreateCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
