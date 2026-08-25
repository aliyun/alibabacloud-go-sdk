// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddExternalSAMLIdPCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *AddExternalSAMLIdPCertificateResponseBody
	GetCertificateId() *string
	SetRequestId(v string) *AddExternalSAMLIdPCertificateResponseBody
	GetRequestId() *string
}

type AddExternalSAMLIdPCertificateResponseBody struct {
	// The ID of the SAML signing certificate.
	//
	// example:
	//
	// idp-c-00wk2fb4foracls0****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12B3E332-DD16-515B-B695-39BA233AA172
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddExternalSAMLIdPCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddExternalSAMLIdPCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *AddExternalSAMLIdPCertificateResponseBody) GetCertificateId() *string {
	return s.CertificateId
}

func (s *AddExternalSAMLIdPCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddExternalSAMLIdPCertificateResponseBody) SetCertificateId(v string) *AddExternalSAMLIdPCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *AddExternalSAMLIdPCertificateResponseBody) SetRequestId(v string) *AddExternalSAMLIdPCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddExternalSAMLIdPCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
