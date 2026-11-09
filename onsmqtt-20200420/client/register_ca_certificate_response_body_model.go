// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RegisterCaCertificateResponseBody
	GetRequestId() *string
	SetSn(v string) *RegisterCaCertificateResponseBody
	GetSn() *string
}

type RegisterCaCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the registered CA certificate. The serial number is the unique identifier of a CA certificate.
	//
	// example:
	//
	// 007269004887******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s RegisterCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterCaCertificateResponseBody) GetSn() *string {
	return s.Sn
}

func (s *RegisterCaCertificateResponseBody) SetRequestId(v string) *RegisterCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterCaCertificateResponseBody) SetSn(v string) *RegisterCaCertificateResponseBody {
	s.Sn = &v
	return s
}

func (s *RegisterCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
