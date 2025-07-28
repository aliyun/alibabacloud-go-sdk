// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *CreateCertsResponseBody
	GetCertIdentifier() *string
	SetRequestId(v string) *CreateCertsResponseBody
	GetRequestId() *string
}

type CreateCertsResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 123456-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5BBA38B1-07AE-559F-8766-AB50****C300
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCertsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertsResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *CreateCertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCertsResponseBody) SetCertIdentifier(v string) *CreateCertsResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *CreateCertsResponseBody) SetRequestId(v string) *CreateCertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertsResponseBody) Validate() error {
	return dara.Validate(s)
}
