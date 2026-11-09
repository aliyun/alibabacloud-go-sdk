// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCaCertificateResponseBody
	GetRequestId() *string
	SetSn(v string) *DeleteCaCertificateResponseBody
	GetSn() *string
}

type DeleteCaCertificateResponseBody struct {
	// The common parameter. The ID of each request is unique and can be used for troubleshooting.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SN of the deleted CA certificate. The SN uniquely identifies a CA certificate.
	//
	// example:
	//
	// 007269004887******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s DeleteCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCaCertificateResponseBody) GetSn() *string {
	return s.Sn
}

func (s *DeleteCaCertificateResponseBody) SetRequestId(v string) *DeleteCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCaCertificateResponseBody) SetSn(v string) *DeleteCaCertificateResponseBody {
	s.Sn = &v
	return s
}

func (s *DeleteCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
