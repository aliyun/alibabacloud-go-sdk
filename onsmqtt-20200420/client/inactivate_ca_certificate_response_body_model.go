// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInactivateCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InactivateCaCertificateResponseBody
	GetRequestId() *string
	SetSn(v string) *InactivateCaCertificateResponseBody
	GetSn() *string
}

type InactivateCaCertificateResponseBody struct {
	// The ID of the request. This ID is unique for each request. You can use it to troubleshoot and locate issues.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number (SN) of the unregistered CA certificate. The SN uniquely identifies a CA certificate.
	//
	// example:
	//
	// 007269004887******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s InactivateCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InactivateCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *InactivateCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InactivateCaCertificateResponseBody) GetSn() *string {
	return s.Sn
}

func (s *InactivateCaCertificateResponseBody) SetRequestId(v string) *InactivateCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *InactivateCaCertificateResponseBody) SetSn(v string) *InactivateCaCertificateResponseBody {
	s.Sn = &v
	return s
}

func (s *InactivateCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
