// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInactivateDeviceCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceSn(v string) *InactivateDeviceCertificateResponseBody
	GetDeviceSn() *string
	SetRequestId(v string) *InactivateDeviceCertificateResponseBody
	GetRequestId() *string
}

type InactivateDeviceCertificateResponseBody struct {
	// The SN of the inactivated device certificate. The SN uniquely identifies a device certificate.
	//
	// example:
	//
	// 356217374433******
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// The ID of the request. This common parameter is unique to each request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InactivateDeviceCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InactivateDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *InactivateDeviceCertificateResponseBody) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *InactivateDeviceCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InactivateDeviceCertificateResponseBody) SetDeviceSn(v string) *InactivateDeviceCertificateResponseBody {
	s.DeviceSn = &v
	return s
}

func (s *InactivateDeviceCertificateResponseBody) SetRequestId(v string) *InactivateDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *InactivateDeviceCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
