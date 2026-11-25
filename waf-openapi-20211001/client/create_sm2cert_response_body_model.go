// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSM2CertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *CreateSM2CertResponseBody
	GetCertIdentifier() *string
	SetRequestId(v string) *CreateSM2CertResponseBody
	GetRequestId() *string
}

type CreateSM2CertResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// ***-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSM2CertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSM2CertResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSM2CertResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *CreateSM2CertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSM2CertResponseBody) SetCertIdentifier(v string) *CreateSM2CertResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *CreateSM2CertResponseBody) SetRequestId(v string) *CreateSM2CertResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSM2CertResponseBody) Validate() error {
	return dara.Validate(s)
}
