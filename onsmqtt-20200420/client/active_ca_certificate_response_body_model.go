// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActiveCaCertificateResponseBody
	GetRequestId() *string
	SetSn(v string) *ActiveCaCertificateResponseBody
	GetSn() *string
}

type ActiveCaCertificateResponseBody struct {
	// Common parameter. Each request has a unique ID that can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SN of the activated CA certificate. The SN uniquely identifies a CA certificate.
	//
	// example:
	//
	// 007269004887******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s ActiveCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveCaCertificateResponseBody) GetSn() *string {
	return s.Sn
}

func (s *ActiveCaCertificateResponseBody) SetRequestId(v string) *ActiveCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveCaCertificateResponseBody) SetSn(v string) *ActiveCaCertificateResponseBody {
	s.Sn = &v
	return s
}

func (s *ActiveCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
